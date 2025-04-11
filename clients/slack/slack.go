package slack

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxSlackMcpServer    = "npx-slack-mcp-server"
	DockerSlackMcpServer = "docker-slack-mcp-server"
)

func InitSlackMCPClient(slackBotToken, slackTeamID, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	slackMCPClient := &param.MCPClientConf{
		Name: NpxSlackMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"SLACK_BOT_TOKEN=" + slackBotToken,
				"SLACK_TEAM_ID=" + slackTeamID,
			},
			Args: []string{
				"-y",
				"@modelcontextprotocol/server-slack",
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
		Name:    "mcp-server/slack",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	slackMCPClient.StdioClientConf.InitReq = initRequest

	return slackMCPClient
}

func InitDockerSlackMCPClient(slackBotToken, slackTeamID, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	slackMCPClient := &param.MCPClientConf{
		Name: DockerSlackMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "docker",
			Env: []string{
				"SLACK_BOT_TOKEN=" + slackBotToken,
				"SLACK_TEAM_ID=" + slackTeamID,
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
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/slack",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	slackMCPClient.StdioClientConf.InitReq = initRequest

	return slackMCPClient
}
