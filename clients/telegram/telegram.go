package telegram

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	DockerTelegramMcpServer = "docker-telegram-mcp-server"
)

type TelegramParam struct {
	TelegramApiId         string
	TelegramApiHash       string
	TelegramSessionString string
}

func InitDockerTelegramMCPClient(p *TelegramParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	telegramMCPClient := &param.MCPClientConf{
		Name: DockerTelegramMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "docker",
			Env: []string{
				"TELEGRAM_API_ID=" + p.TelegramApiId,
				"TELEGRAM_API_HASH=" + p.TelegramApiHash,
				"TELEGRAM_SESSION_STRING=" + p.TelegramSessionString,
				"SESSION_STRING=" + p.TelegramSessionString,
			},
			Args: []string{
				"run",
				"-i",
				"--rm",
				"telegram-mcp:latest",
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
		Name:    "mcp-server/telegram",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	telegramMCPClient.StdioClientConf.InitReq = initRequest

	return telegramMCPClient
}
