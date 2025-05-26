package slack

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxSlackMcpServer    = "npx-slack-mcp-server"
	DockerSlackMcpServer = "docker-slack-mcp-server"
)

type SlackParam struct {
	SlackBotToken string
	SlackTeamID   string
}

func InitSlackMCPClient(p *SlackParam, options ...param.Option) *param.MCPClientConf {

	slackMCPClient := &param.MCPClientConf{
		Name: NpxSlackMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"SLACK_BOT_TOKEN=" + p.SlackBotToken,
				"SLACK_TEAM_ID=" + p.SlackTeamID,
			},
			Args: []string{
				"-y",
				"@modelcontextprotocol/server-slack",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(slackMCPClient)
	}

	if slackMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		slackMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if slackMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		slackMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/slack",
			Version: "0.1.0",
		}
	}

	return slackMCPClient
}

func InitDockerSlackMCPClient(p *SlackParam, options ...param.Option) *param.MCPClientConf {

	slackMCPClient := &param.MCPClientConf{
		Name: DockerSlackMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "docker",
			Env: []string{
				"SLACK_BOT_TOKEN=" + p.SlackBotToken,
				"SLACK_TEAM_ID=" + p.SlackTeamID,
			},
			Args: []string{
				"run",
				"-i",
				"--rm",
				"-e",
				"SLACK_BOT_TOKEN",
				"-e",
				"SLACK_TEAM_ID",
				"mcp/slack",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(slackMCPClient)
	}

	if slackMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		slackMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if slackMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		slackMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/slack",
			Version: "0.1.0",
		}
	}

	return slackMCPClient
}
