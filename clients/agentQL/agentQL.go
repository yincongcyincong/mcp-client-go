package agentQL

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxAgentQLMcpServer = "npx-agent-ql-mcp-server"
)

func InitAmapMCPClient(AgentQLApiKey string, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	agentQLMCPClient := &param.MCPClientConf{
		Name:    NpxAgentQLMcpServer,
		Command: "npx",
		Env: []string{
			"AGENTQL_API_KEY=" + AgentQLApiKey,
		},
		Args: []string{
			"-y",
			"agentql-mcp",
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
		Name:    "mcp-server/agent-ql",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	agentQLMCPClient.InitReq = initRequest

	return agentQLMCPClient
}
