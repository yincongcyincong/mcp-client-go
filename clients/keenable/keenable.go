// Package keenable configures the hosted Keenable search MCP service.
package keenable

import (
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/yincongcyincong/mcp-client-go/clients"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	KeenableMCPServer = "keenable-mcp"
	// PublicMCPURL serves search_web_pages and fetch_page_content without a key.
	PublicMCPURL = "https://api.keenable.ai/mcp-public"
	// KeyedMCPURL is used when an API key is set.
	KeyedMCPURL = "https://api.keenable.ai/mcp"
	Title       = "mcp-client-go"
)

// KeenableParam is optional. Leave ApiKey empty to use the keyless public tier.
type KeenableParam struct {
	ApiKey string
}

// InitKeenableMCPClient returns an opt-in Streamable HTTP client configuration.
// No local server is required. Options follow InitHttpMCPClient; custom HTTP
// headers replace the default ones.
func InitKeenableMCPClient(p *KeenableParam, options ...param.Option) *param.MCPClientConf {
	url := PublicMCPURL
	headers := map[string]string{"X-Keenable-Title": Title}
	if p != nil && p.ApiKey != "" {
		url = KeyedMCPURL
		headers["X-API-Key"] = p.ApiKey
	}
	conf := clients.InitHttpMCPClient(KeenableMCPServer, url, options...)
	conf.HTTPStreamerConf.Options = append([]transport.StreamableHTTPCOption{
		transport.WithHTTPHeaders(headers),
	}, conf.HTTPStreamerConf.Options...)
	return conf
}
