package jira

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	DockerJiraMcpServer = "docker-jira-mcp-server"
)

type JiraParams struct {
	AtlassianHost  string
	AtlassianEmail string
	AtlassianToken string
}

func InitDockerFetchMCPClient(p *JiraParams, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	fetchMCPClient := &param.MCPClientConf{
		Name: DockerJiraMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "docker",
			Env:     []string{},
			Args: []string{
				"run",
				"--rm",
				"-i",
				"-e", "ATLASSIAN_HOST=" + p.AtlassianHost,
				"-e", "ATLASSIAN_EMAIL=" + p.AtlassianEmail,
				"-e", "ATLASSIAN_TOKEN=" + p.AtlassianToken,
				"ghcr.io/nguyenvanduocit/jira-mcp:latest",
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
		Name:    "mcp-server/fetch",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	fetchMCPClient.StdioClientConf.InitReq = initRequest

	return fetchMCPClient
}
