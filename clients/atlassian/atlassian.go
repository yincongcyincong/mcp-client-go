package atlassian

import (
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

func InitAtlassianMCPClient(p *AtlassianParam, options ...param.Option) *param.MCPClientConf {

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
	}

	for _, o := range options {
		o(atlassianMCPClient)
	}

	if atlassianMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		atlassianMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if atlassianMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		atlassianMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/atlassian",
			Version: "0.1.0",
		}
	}

	return atlassianMCPClient
}

type AtlassianSSEParam struct {
	BaseUrl string
	Option  []transport.ClientOption
}

func InitAtlassianSSEMCPClient(p *AtlassianSSEParam, options ...param.Option) *param.MCPClientConf {

	atlassianMCPClient := &param.MCPClientConf{
		Name:       SseAtlassianMcpServer,
		ClientType: param.SSEType,
		SSEClientConf: &param.SSEClientConfig{
			BaseUrl: p.BaseUrl,
			Options: p.Option,
		},
	}

	for _, o := range options {
		o(atlassianMCPClient)
	}

	if atlassianMCPClient.SSEClientConf.InitReq.Params.ProtocolVersion == "" {
		atlassianMCPClient.SSEClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if atlassianMCPClient.SSEClientConf.InitReq.Params.ClientInfo.Name == "" {
		atlassianMCPClient.SSEClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/atlassian",
			Version: "0.1.0",
		}
	}

	return atlassianMCPClient
}
