package apple_shortcut

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxAppleShortcutMapsMcpServer = "npx-apple-shortcut-maps-mcp-server"
)

type AppleShortcutParam struct {
}

func InitAppleShortcutMCPClient(p *AppleShortcutParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	appleShortcutMCPClient := &param.MCPClientConf{
		Name: NpxAppleShortcutMapsMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"-y",
				"mcp-server-apple-shortcuts\"",
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
		Name:    "mcp-server/apple-shortcut",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	appleShortcutMCPClient.StdioClientConf.InitReq = initRequest

	return appleShortcutMCPClient
}
