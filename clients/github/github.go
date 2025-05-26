package github

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxModelContextProtocolGithubServer = "npx-github-mcp-server"
	DockerGithubServer                  = "docker-github-mcp-server"
)

type GithubParam struct {
	GithubAccessToken string
}

func InitModelContextProtocolGithubMCPClient(p *GithubParam, options ...param.Option) *param.MCPClientConf {

	githubMCPClient := &param.MCPClientConf{
		Name: NpxModelContextProtocolGithubServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"GITHUB_PERSONAL_ACCESS_TOKEN=" + p.GithubAccessToken,
			},
			Args: []string{
				"-y",
				"@modelcontextprotocol/server-github",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(githubMCPClient)
	}

	if githubMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		githubMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if githubMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		githubMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/server-github",
			Version: "0.2.0",
		}
	}

	return githubMCPClient
}

func InitDockerGithubMCPClient(p *GithubParam, options ...param.Option) *param.MCPClientConf {

	githubMCPClient := &param.MCPClientConf{
		Name: DockerGithubServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "docker",
			Env: []string{
				"GITHUB_PERSONAL_ACCESS_TOKEN=" + p.GithubAccessToken,
			},
			Args: []string{
				"run",
				"-i",
				"--rm",
				"-e",
				"GITHUB_PERSONAL_ACCESS_TOKEN",
				"ghcr.io/github/github-mcp-server",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(githubMCPClient)
	}

	if githubMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		githubMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if githubMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		githubMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/github",
			Version: "0.1.0",
		}
	}

	return githubMCPClient
}
