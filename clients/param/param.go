package param

import (
	"github.com/mark3labs/mcp-go/mcp"
)

type MCPClientConf struct {
	Name    string
	Command string
	Env     []string
	Args    []string
	InitReq mcp.InitializeRequest
	
	ToolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error
	ToolsAfterFunc  map[string]func(req *mcp.CallToolResult) (string, error)
}
