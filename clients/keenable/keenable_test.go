package keenable_test

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
	"github.com/yincongcyincong/mcp-client-go/clients/keenable"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

// routeTransport checks what would go to Keenable and sends it to a local fixture instead.
type routeTransport struct {
	t      *testing.T
	target *url.URL
	base   http.RoundTripper
	url    string
	key    string
}

func (r routeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.String() != r.url {
		r.t.Errorf("unexpected endpoint: %s", req.URL)
	}
	if got := req.Header.Get("X-Keenable-Title"); got != keenable.Title {
		r.t.Errorf("unexpected X-Keenable-Title: %q", got)
	}
	if got := req.Header.Get("X-API-Key"); got != r.key {
		r.t.Errorf("unexpected X-API-Key: %q", got)
	}
	if req.Header.Get("Authorization") != "" {
		r.t.Error("client sent an Authorization header")
	}
	clone := req.Clone(req.Context())
	clone.URL = r.target
	return r.base.RoundTrip(clone)
}

func fixture(t *testing.T, seen map[string]int, mu *sync.Mutex) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			result = map[string]interface{}{"protocolVersion": "2025-03-26", "capabilities": map[string]interface{}{"tools": map[string]interface{}{}}, "serverInfo": map[string]string{"name": "keenable-fixture", "version": "1"}}
		case "notifications/initialized":
			w.WriteHeader(http.StatusAccepted)
			return
		case "tools/list":
			result = json.RawMessage(`{"tools":[{"name":"search_web_pages","inputSchema":{"type":"object","properties":{"query":{"type":"string"},"max_results":{"type":"integer"}},"required":["query"]}},{"name":"fetch_page_content","inputSchema":{"type":"object","properties":{"url":{"type":"string"}},"required":["url"]}}]}`)
		case "tools/call":
			switch req.Params.Name {
			case "search_web_pages":
				if req.Params.Arguments["query"] != "Go context cancellation" {
					t.Errorf("unexpected search arguments: %v", req.Params.Arguments)
				}
			case "fetch_page_content":
				if req.Params.Arguments["url"] != "https://go.dev/blog/context" {
					t.Errorf("unexpected fetch arguments: %v", req.Params.Arguments)
				}
			default:
				t.Errorf("unexpected tool: %s", req.Params.Name)
			}
			result = json.RawMessage(`{"content":[{"type":"text","text":"https://go.dev/blog/context: Context carries deadlines and cancellation signals."}]}`)
		default:
			t.Errorf("unexpected method: %s", req.Method)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"jsonrpc": "2.0", "id": req.ID, "result": result})
	}))
}

func run(t *testing.T, p *keenable.KeenableParam, wantURL, wantKey string) {
	var mu sync.Mutex
	seen := map[string]int{}
	srv := fixture(t, seen, &mu)
	defer srv.Close()
	endpoint, _ := url.Parse(srv.URL)
	previous := http.DefaultTransport
	http.DefaultTransport = routeTransport{t: t, target: endpoint, base: previous, url: wantURL, key: wantKey}
	defer func() { http.DefaultTransport = previous }()

	// Supplying a timeout must keep the helper's endpoint and headers.
	conf := keenable.InitKeenableMCPClient(p,
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
	c, err := cs.GetMCPClient(keenable.KeenableMCPServer)
	if err != nil {
		t.Fatal(err)
	}
	selected, err := cs.GetMCPClientByToolName("search_web_pages")
	if err != nil || selected != c {
		t.Fatalf("tool selection: %v", err)
	}
	for name, args := range map[string]map[string]interface{}{
		"search_web_pages":   {"query": "Go context cancellation", "max_results": 3},
		"fetch_page_content": {"url": "https://go.dev/blog/context"},
	} {
		output, err := c.ExecTools(ctx, name, args)
		if err != nil || !strings.Contains(output, "cancellation signals") {
			t.Fatalf("%s output %q, error %v", name, output, err)
		}
	}
	cancelled, stop := context.WithCancel(ctx)
	stop()
	_, err = c.ExecTools(cancelled, "fetch_page_content", map[string]interface{}{"url": "https://go.dev/blog/context"})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled call returned: %v", err)
	}
	mu.Lock()
	defer mu.Unlock()
	for _, key := range []string{"initialize:", "tools/list:", "tools/call:search_web_pages", "tools/call:fetch_page_content"} {
		if seen[key] != 1 {
			t.Errorf("%s requests: %d", key, seen[key])
		}
	}
}

func TestKeylessUsesPublicEndpoint(t *testing.T) {
	run(t, nil, keenable.PublicMCPURL, "")
	run(t, &keenable.KeenableParam{}, keenable.PublicMCPURL, "")
}

func TestApiKeyUsesKeyedEndpoint(t *testing.T) {
	run(t, &keenable.KeenableParam{ApiKey: "test-key"}, keenable.KeyedMCPURL, "test-key")
}
