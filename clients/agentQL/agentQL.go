package agentQL

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxAgentQLMcpServer = "npx-agent-ql-mcp-server"
)

type AgentQLParam struct {
	AgentQLApiKey string
}

func InitAgentQLMCPClient(p *AgentQLParam, options ...param.Option) *param.MCPClientConf {
	agentQLMCPClient := &param.MCPClientConf{
		Name: NpxAgentQLMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"AGENTQL_API_KEY=" + p.AgentQLApiKey,
			},
			Args: []string{
				"-y",
				"agentql-mcp",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(agentQLMCPClient)
	}

	if agentQLMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		agentQLMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if agentQLMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		agentQLMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/agent-ql",
			Version: "0.1.0",
		}
	}
	return agentQLMCPClient
}
