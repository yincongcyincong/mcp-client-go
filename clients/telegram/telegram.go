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

func InitDockerTelegramMCPClient(p *TelegramParam, options ...param.Option) *param.MCPClientConf {

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
	}

	for _, o := range options {
		o(telegramMCPClient)
	}

	if telegramMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		telegramMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if telegramMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		telegramMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/telegram",
			Version: "0.1.0",
		}
	}

	return telegramMCPClient
}
