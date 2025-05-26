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

func InitDockerJiraMCPClient(p *JiraParams, options ...param.Option) *param.MCPClientConf {

	jiraMCPClient := &param.MCPClientConf{
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
	}

	for _, o := range options {
		o(jiraMCPClient)
	}

	if jiraMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		jiraMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if jiraMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		jiraMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/jira",
			Version: "0.1.0",
		}
	}

	return jiraMCPClient
}
