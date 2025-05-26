package chatsum

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NodeChatsumMcpServer = "node-chatsum-mcp-server"
)

type ChatsumParam struct {
	chatDBPath  string
	indexJsPath string
}

func InitChatsumMCPClient(p *ChatsumParam, options ...param.Option) *param.MCPClientConf {

	chatsumMCPClient := &param.MCPClientConf{
		Name: NodeChatsumMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "node",
			Env: []string{
				"CHAT_DB_PATH=" + p.chatDBPath,
			},
			Args: []string{
				p.indexJsPath,
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(chatsumMCPClient)
	}

	if chatsumMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		chatsumMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if chatsumMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		chatsumMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/chatsum",
			Version: "0.1.0",
		}
	}

	return chatsumMCPClient
}
