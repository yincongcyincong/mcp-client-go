package sentry

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	UvxSentryMcpServer    = "uvx-sentry-mcp-server"
	DockerSentryMcpServer = "docker-sentry-mcp-server"
)

type SentryParam struct {
	SentryToken string
}

func InitSentryMCPClient(p *SentryParam, options ...param.Option) *param.MCPClientConf {

	sentryMCPClient := &param.MCPClientConf{
		Name: UvxSentryMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"mcp-server-sentry",
				"--auth-token",
				p.SentryToken,
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(sentryMCPClient)
	}

	if sentryMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		sentryMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if sentryMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		sentryMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/sentry",
			Version: "0.1.0",
		}
	}

	return sentryMCPClient
}

func InitDockerSentryMCPClient(p *SentryParam, options ...param.Option) *param.MCPClientConf {

	sentryMCPClient := &param.MCPClientConf{
		Name: DockerSentryMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "docker",
			Env:     []string{},
			Args: []string{
				"run",
				"-i",
				"--rm",
				"mcp/sentry",
				"--auth-token",
				p.SentryToken,
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(sentryMCPClient)
	}

	if sentryMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		sentryMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if sentryMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		sentryMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/sentry",
			Version: "0.1.0",
		}
	}

	return sentryMCPClient
}
