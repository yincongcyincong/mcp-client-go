package duckduckgo

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxDuckduckgoMcpServer = "npx-duckduckgo-mcp-server"
)

type DuckduckgoParam struct {
}

func InitDuckduckgoMCPClient(p *DuckduckgoParam, options ...param.Option) *param.MCPClientConf {

	duckduckgoMCPClient := &param.MCPClientConf{
		Name:       NpxDuckduckgoMcpServer,
		ClientType: param.StdioType,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uvx",
			Env:     []string{},
			Args: []string{
				"duckduckgo-mcp-server",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(duckduckgoMCPClient)
	}

	if duckduckgoMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		duckduckgoMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if duckduckgoMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		duckduckgoMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/duckduckgo",
			Version: "0.1.0",
		}
	}

	return duckduckgoMCPClient
}
