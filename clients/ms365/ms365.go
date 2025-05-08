package ms365

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxMS365McpServer = "npx-ms365-mcp-server"
)

type MS365Param struct {
	MS365Session string
}

func InitMS365MCPClient(p *MS365Param, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	ms365MCPClient := &param.MCPClientConf{
		Name: NpxMS365McpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
			},
			Args: []string{
                "-y",
                "@softeria/ms-365-mcp-server"
			},
			InitReq: mcp.InitializeRequest{},
		},
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/ms365",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	ms365MCPClient.StdioClientConf.InitReq = initRequest

	return ms365MCPClient
}
