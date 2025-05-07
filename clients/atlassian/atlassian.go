package atlassian

import (
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	UvxAtlassianMcpServer = "npx-atlassian-mcp-server"
	SseAtlassianMcpServer = "sse-atlassian-mcp-server"
)

type AtlassianParam struct {
	ConfluenceUrl      string
	ConfluenceUsername string
	ConfluenceApiToken string
	JiraUrl            string
	JiraUsername       string
	JiraApiToken       string
}

func InitAtlassianMCPClient(p *AtlassianParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	atlassianMCPClient := &param.MCPClientConf{
		Name:       UvxAtlassianMcpServer,
		ClientType: param.StdioType,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uvx",
			Env: []string{
				"CONFLUENCE_URL=" + p.ConfluenceUrl,
				"CONFLUENCE_USERNAME=" + p.ConfluenceUsername,
				"CONFLUENCE_API_TOKEN=" + p.ConfluenceApiToken,
				"JIRA_URL=" + p.JiraUrl,
				"JIRA_USERNAME=" + p.JiraUsername,
				"JIRA_API_TOKEN=" + p.JiraApiToken,
			},
			Args: []string{
				"mcp-atlassian",
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
		Name:    "mcp-server/atlassian",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	atlassianMCPClient.StdioClientConf.InitReq = initRequest

	return atlassianMCPClient
}

func InitAtlassianSSEMCPClient(baseUrl string, options []transport.ClientOption,
	protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	atlassianMCPClient := &param.MCPClientConf{
		Name:       SseAtlassianMcpServer,
		ClientType: param.SSEType,
		SSEClientConf: &param.SSEClientConfig{
			BaseUrl: baseUrl,
			Options: options,
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
		Name:    "mcp-server/atlassian",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	atlassianMCPClient.SSEClientConf.InitReq = initRequest

	return atlassianMCPClient
}
