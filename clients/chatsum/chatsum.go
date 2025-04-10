package chatsum

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NodeChatsumMcpServer = "node-chatsum-mcp-server"
)

func InitChatsumMCPClient(chatDBPath, indexJsPath string, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	chatsumMCPClient := &param.MCPClientConf{
		Name:    NodeChatsumMcpServer,
		Command: "node",
		Env: []string{
			"CHAT_DB_PATH=" + chatDBPath,
		},
		Args: []string{
			indexJsPath,
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
		Name:    "mcp-server/chatsum",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	chatsumMCPClient.InitReq = initRequest

	return chatsumMCPClient
}
