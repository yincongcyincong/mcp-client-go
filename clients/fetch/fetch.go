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

func InitFetchMCPClient(p *FetchParam, options ...param.Option) *param.MCPClientConf {

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
	}

	for _, o := range options {
		o(fetchMCPClient)
	}

	if fetchMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		fetchMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if fetchMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		fetchMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/fetch",
			Version: "0.1.0",
		}
	}

	return fetchMCPClient
}

func InitDockerFetchMCPClient(p *FetchParam, options ...param.Option) *param.MCPClientConf {

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
	}

	for _, o := range options {
		o(fetchMCPClient)
	}

	if fetchMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		fetchMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if fetchMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		fetchMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/fetch",
			Version: "0.1.0",
		}
	}

	return fetchMCPClient
}
