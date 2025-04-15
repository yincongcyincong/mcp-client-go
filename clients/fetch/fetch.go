package fetch

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxFetchMcpServer = "npx-fetch-mcp-server"
	DockerFetchServer = "docker-fetch-mcp-server"
)

type FetchParam struct {
}

func InitFetchMCPClient(p *FetchParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	fetchMCPClient := &param.MCPClientConf{
		Name: NpxFetchMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uvx",
			Env:     []string{},
			Args: []string{
				"mcp-server-fetch",
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

func InitDockerFetchMCPClient(p *FetchParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	fetchMCPClient := &param.MCPClientConf{
		Name: DockerFetchServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "docker",
			Env:     []string{},
			Args: []string{
				"run",
				"-i",
				"--rm",
				"mcp/fetch",
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
