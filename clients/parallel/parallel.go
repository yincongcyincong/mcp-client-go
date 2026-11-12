// Package parallel configures the keyless Parallel Search MCP service.
package parallel

import (
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/yincongcyincong/mcp-client-go/clients"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	ParallelMCPServer = "parallel-search-mcp"
	SearchMCPURL      = "https://search.parallel.ai/mcp"
	UserAgent         = "mcp-client-go/parallel"
)

// InitParallelMCPClient returns an opt-in Streamable HTTP client configuration.
// No API key or local server is required. Options follow InitHttpMCPClient;
// custom HTTP headers replace the default User-Agent header.
func InitParallelMCPClient(options ...param.Option) *param.MCPClientConf {
	conf := clients.InitHttpMCPClient(ParallelMCPServer, SearchMCPURL, options...)
	conf.HTTPStreamerConf.Options = append([]transport.StreamableHTTPCOption{
		transport.WithHTTPHeaders(map[string]string{"User-Agent": UserAgent}),
	}, conf.HTTPStreamerConf.Options...)
	return conf
}
