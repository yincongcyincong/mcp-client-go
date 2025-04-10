package everart

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxEverartMcpServer = "npx-everart-mcp-server"
)

func InitEverartMCPClient(everartApiKey string, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	everartMCPClient := &param.MCPClientConf{
		Name:    NpxEverartMcpServer,
		Command: "npx",
		Env: []string{
			"EVERART_API_KEY=" + everartApiKey,
		},
		Args: []string{
			"-y",
			"@modelcontextprotocol/server-everart",
		},
		InitReq:         mcp.InitializeRequest{},
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/everart",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	everartMCPClient.InitReq = initRequest

	return everartMCPClient
}
