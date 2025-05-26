package ms365

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxMS365McpServer = "npx-ms365-mcp-server"
)

type MS365Param struct {
	MS365Session string
}

func InitMS365MCPClient(p *MS365Param, options ...param.Option) *param.MCPClientConf {

	ms365MCPClient := &param.MCPClientConf{
		Name: NpxMS365McpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"-y",
				"@softeria/ms-365-mcp-server",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(ms365MCPClient)
	}

	if ms365MCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		ms365MCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if ms365MCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		ms365MCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/m365",
			Version: "0.1.0",
		}
	}

	return ms365MCPClient
}
