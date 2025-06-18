package param

import (
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	SSEType      = "sse"
	StdioType    = "stdio"
	HTTPStreamer = "http-streamer"

	HTTPConfigType  = "http"
	StdioConfigType = "stdio"
)

type McpClientGoConfig struct {
	McpServers map[string]*MCPConfig `json:"mcpServers"`
}

type MCPConfig struct {
	Command string            `json:"command"`
	Args    []string          `json:"args"`
	Env     map[string]string `json:"env"`

	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`

	Type        string `json:"type"`
	Description string `json:"description"`
}

type MCPClientConf struct {
	Name             string
	Description      string
	ClientType       string
	SSEClientConf    *SSEClientConfig
	StdioClientConf  *StdioClientConfig
	HTTPStreamerConf *HTTPStreamerConfig

	ToolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error
	ToolsAfterFunc  map[string]func(req *mcp.CallToolResult) (string, error)
}

type StdioClientConfig struct {
	Command string
	Env     []string
	Args    []string
	InitReq mcp.InitializeRequest
}

type SSEClientConfig struct {
	Options []transport.ClientOption
	BaseUrl string
	InitReq mcp.InitializeRequest
}

type HTTPStreamerConfig struct {
	BaseURL string
	Options []transport.StreamableHTTPCOption
	InitReq mcp.InitializeRequest
	Oauth   *client.OAuthConfig
}

type Option func(p *MCPClientConf)

func WithProtocolVersion(protocolVersion string) Option {
	return func(p *MCPClientConf) {
		if p.StdioClientConf != nil {
			p.StdioClientConf.InitReq.Params.ProtocolVersion = protocolVersion
		}
		if p.SSEClientConf != nil {
			p.SSEClientConf.InitReq.Params.ProtocolVersion = protocolVersion
		}
	}
}

func WithClientInfo(clientInfo mcp.Implementation) Option {
	return func(p *MCPClientConf) {
		if p.StdioClientConf != nil {
			p.StdioClientConf.InitReq.Params.ClientInfo = clientInfo
		}
		if p.SSEClientConf != nil {
			p.SSEClientConf.InitReq.Params.ClientInfo = clientInfo
		}
	}
}

func WithToolsBeforeFunc(toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error) Option {
	return func(p *MCPClientConf) {
		p.ToolsBeforeFunc = toolsBeforeFunc
	}
}

func WithToolsAfterFunc(toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) Option {
	return func(p *MCPClientConf) {
		p.ToolsAfterFunc = toolsAfterFunc
	}
}

func WithHttpOptions(httpOptions []transport.StreamableHTTPCOption) Option {
	return func(p *MCPClientConf) {
		p.HTTPStreamerConf.Options = httpOptions
	}
}

func WithHttpOauth(oauth *client.OAuthConfig) Option {
	return func(p *MCPClientConf) {
		p.HTTPStreamerConf.Oauth = oauth
	}
}

func WithSSEOptions(sseOptions []transport.ClientOption) Option {
	return func(p *MCPClientConf) {
		p.SSEClientConf.Options = sseOptions
	}
}

func WithDescription(description string) Option {
	return func(p *MCPClientConf) {
		p.Description = description
	}
}
