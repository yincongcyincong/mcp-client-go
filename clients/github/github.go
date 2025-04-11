package github

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxModelContextProtocolGithubServer = "npx-github-mcp-server"
	DockerGithubServer                  = "docker-github-mcp-server"
)

func InitModelContextProtocolGithubMCPClient(githubAccessToken string, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	amapMCPClient := &param.MCPClientConf{
		Name: NpxModelContextProtocolGithubServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"GITHUB_PERSONAL_ACCESS_TOKEN=" + githubAccessToken,
			},
			Args: []string{
				"-y",
				"@modelcontextprotocol/server-github",
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
		Name:    "modelcontextprotocol/server-github",
		Version: "0.2.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}

	amapMCPClient.StdioClientConf.InitReq = initRequest

	return amapMCPClient
}

func InitDockerGithubMCPClient(githubAccessToken string, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	amapMCPClient := &param.MCPClientConf{
		Name: DockerGithubServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "docker",
			Env: []string{
				"GITHUB_PERSONAL_ACCESS_TOKEN=" + githubAccessToken,
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
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}

	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "modelcontextprotocol/server-github",
		Version: "0.2.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}

	amapMCPClient.StdioClientConf.InitReq = initRequest

	return amapMCPClient
}
