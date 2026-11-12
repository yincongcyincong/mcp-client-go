package parallel_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/yincongcyincong/mcp-client-go/clients"
	"github.com/yincongcyincong/mcp-client-go/clients/parallel"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

type routeTransport struct {
	t      *testing.T
	target *url.URL
	base   http.RoundTripper
}

func (r routeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.String() != parallel.SearchMCPURL {
		r.t.Errorf("unexpected endpoint: %s", req.URL)
	}
	if req.Header.Get("User-Agent") != parallel.UserAgent {
		r.t.Errorf("unexpected User-Agent: %q", req.Header.Get("User-Agent"))
	}
	if req.Header.Get("Authorization") != "" || req.Header.Get("x-api-key") != "" {
		r.t.Error("anonymous client sent credentials")
	}
	clone := req.Clone(req.Context())
	clone.URL = r.target
	return r.base.RoundTrip(clone)
}

func TestRegisteredParallelTools(t *testing.T) {
	var mu sync.Mutex
	seen := map[string]int{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID     json.RawMessage `json:"id"`
			Method string          `json:"method"`
			Params struct {
				Name      string                 `json:"name"`
				Arguments map[string]interface{} `json:"arguments"`
			} `json:"params"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		mu.Lock()
		seen[req.Method+":"+req.Params.Name]++
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		var result interface{}
		switch req.Method {
		case "initialize":
			result = map[string]interface{}{"protocolVersion": "2025-03-26", "capabilities": map[string]interface{}{"tools": map[string]interface{}{}}, "serverInfo": map[string]string{"name": "parallel-fixture", "version": "1"}}
		case "notifications/initialized":
			w.WriteHeader(http.StatusAccepted)
			return
		case "tools/list":
			result = json.RawMessage(`{"tools":[{"name":"web_search","inputSchema":{"type":"object","properties":{"objective":{"type":"string"},"search_queries":{"type":"array","items":{"type":"string"}}},"required":["objective","search_queries"]}},{"name":"web_fetch","inputSchema":{"type":"object","properties":{"urls":{"type":"array","items":{"type":"string"}}},"required":["urls"]}}]}`)
		case "tools/call":
			if req.Params.Name == "web_search" {
				if req.Params.Arguments["objective"] != "Go context cancellation" {
					t.Errorf("unexpected search arguments: %v", req.Params.Arguments)
				}
				if req.Params.Arguments["search_queries"] == nil {
					t.Error("missing search queries")
				}
			} else if req.Params.Name == "web_fetch" {
				urls, _ := req.Params.Arguments["urls"].([]interface{})
				if len(urls) != 1 || urls[0] != "https://go.dev/blog/context" {
					t.Errorf("unexpected fetch URLs: %v", urls)
				}
			} else {
				t.Errorf("unexpected tool: %s", req.Params.Name)
			}
			result = json.RawMessage(`{"content":[{"type":"text","text":"https://go.dev/blog/context: Context carries deadlines and cancellation signals."}]}`)
		default:
			t.Errorf("unexpected method: %s", req.Method)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"jsonrpc": "2.0", "id": req.ID, "result": result})
	}))
	defer srv.Close()
	endpoint, _ := url.Parse(srv.URL)
	previous := http.DefaultTransport
	http.DefaultTransport = routeTransport{t, endpoint, previous}
	defer func() { http.DefaultTransport = previous }()

	// Supplying a timeout must retain the helper's anonymous endpoint and headers.
	conf := parallel.InitParallelMCPClient(
		param.WithDescription("optional web search"),
		param.WithHttpOptions(transport.WithHTTPTimeout(time.Second)),
	)
	cs := &clients.Clients{}
	defer cs.ClearAllMCPClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if errs := cs.RegisterMCPClient(ctx, []*param.MCPClientConf{conf}); len(errs) != 0 {
		t.Fatalf("registration: %v", errs)
	}
	c, err := cs.GetMCPClient(parallel.ParallelMCPServer)
	if err != nil {
		t.Fatal(err)
	}
	selected, err := cs.GetMCPClientByToolName("web_search")
	if err != nil || selected != c {
		t.Fatalf("tool selection: %v", err)
	}
	for name, args := range map[string]map[string]interface{}{
		"web_search": {"objective": "Go context cancellation", "search_queries": []string{"Go context cancellation documentation"}},
		"web_fetch":  {"urls": []string{"https://go.dev/blog/context"}},
	} {
		output, err := c.ExecTools(ctx, name, args)
		if err != nil || !strings.Contains(output, "cancellation signals") {
			t.Fatalf("%s output %q, error %v", name, output, err)
		}
	}
	cancelled, stop := context.WithCancel(ctx)
	stop()
	_, err = c.ExecTools(cancelled, "web_fetch", map[string]interface{}{"urls": []string{"https://go.dev/blog/context"}})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled call returned: %v", err)
	}
	mu.Lock()
	defer mu.Unlock()
	for _, key := range []string{"initialize:", "tools/list:", "tools/call:web_search", "tools/call:web_fetch"} {
		if seen[key] != 1 {
			t.Errorf("%s requests: %d", key, seen[key])
		}
	}
}
