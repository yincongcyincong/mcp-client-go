package gitlab

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxGitlabMcpServer    = "npx-gitlab-mcp-server"
	DockerGitlabMcpServer = "docker-gitlab-mcp-server"
)

type GitlabParam struct {
	GitlabApiKey string
	GitlabUrl    string
}

func InitGitlabMCPClient(p *GitlabParam, options ...param.Option) *param.MCPClientConf {

	gitlabMCPClient := &param.MCPClientConf{
		Name: NpxGitlabMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"GITLAB_PERSONAL_ACCESS_TOKEN=" + p.GitlabApiKey,
				"GITLAB_API_URL=" + p.GitlabUrl, // Optional, for self-hosted instances
			},
			Args: []string{
				"-y",
				"@modelcontextprotocol/server-gitlab",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(gitlabMCPClient)
	}

	if gitlabMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		gitlabMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if gitlabMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		gitlabMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/gitlab",
			Version: "0.1.0",
		}
	}

	return gitlabMCPClient
}

func InitDockerGitlabMCPClient(p *GitlabParam, options ...param.Option) *param.MCPClientConf {

	gitlabMCPClient := &param.MCPClientConf{
		Name: DockerGitlabMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "docker",
			Env: []string{
				"GITLAB_PERSONAL_ACCESS_TOKEN=" + p.GitlabApiKey,
				"GITLAB_API_URL" + p.GitlabUrl, // Optional, for self-hosted instances
			},
			Args: []string{
				"run",
				"--rm",
				"-i",
				"-e",
				"GITLAB_PERSONAL_ACCESS_TOKEN",
				"-e",
				"GITLAB_API_URL",
				"mcp/gitlab",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(gitlabMCPClient)
	}

	if gitlabMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		gitlabMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if gitlabMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		gitlabMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/gitlab",
			Version: "0.1.0",
		}
	}

	return gitlabMCPClient
}
