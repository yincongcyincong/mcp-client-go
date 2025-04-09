package github

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

func InitGithubMCPClient(githubAccessToken string,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	amapMCPClient := &param.MCPClientConf{
		Name:    "npx-amap",
		Command: "npx",
		Env: []string{
			"GITHUB_PERSONAL_ACCESS_TOKEN=" + githubAccessToken,
		},
		Args: []string{
			"-y",
			"@modelcontextprotocol/server-github",
		},
		InitReq:         mcp.InitializeRequest{},
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}
	
	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "modelcontextprotocol/server-github",
		Version: "0.2.0",
	}
	amapMCPClient.InitReq = initRequest

	return amapMCPClient
}
