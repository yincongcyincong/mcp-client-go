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
