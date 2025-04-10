package sentry

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	UvxSentryMcpServer    = "uvx-sentry-mcp-server"
	DockerSentryMcpServer = "docker-sentry-mcp-server"
)

func InitSentryMCPClient(sentryToken, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	sentryMCPClient := &param.MCPClientConf{
		Name:    UvxSentryMcpServer,
		Command: "npx",
		Env:     []string{},
		Args: []string{
			"mcp-server-sentry",
			"--auth-token",
			sentryToken,
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
		Name:    "mcp-server/sentry",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	sentryMCPClient.InitReq = initRequest

	return sentryMCPClient
}

func InitDockerSentryMCPClient(sentryToken, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	sentryMCPClient := &param.MCPClientConf{
		Name:    DockerSentryMcpServer,
		Command: "docker",
		Env:     []string{},
		Args: []string{
			"run",
			"-i",
			"--rm",
			"mcp/sentry",
			"--auth-token",
			sentryToken,
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
		Name:    "mcp-server/sentry",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	sentryMCPClient.InitReq = initRequest

	return sentryMCPClient
}
