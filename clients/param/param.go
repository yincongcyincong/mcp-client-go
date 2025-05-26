package param

import (
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
)

const (
	SSEType   = "sse"
	StdioType = "stdio"
)

type MCPClientConf struct {
	Name            string
	ClientType      string
	SSEClientConf   *SSEClientConfig
	StdioClientConf *StdioClientConfig

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
