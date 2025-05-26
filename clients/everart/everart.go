package everart

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxEverartMcpServer = "npx-everart-mcp-server"
)

type EverartParam struct {
	EverartApiKey string
}

func InitEverartMCPClient(p *EverartParam, options ...param.Option) *param.MCPClientConf {

	everartMCPClient := &param.MCPClientConf{
		Name: NpxEverartMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"EVERART_API_KEY=" + p.EverartApiKey,
			},
			Args: []string{
				"-y",
				"@modelcontextprotocol/server-everart",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(everartMCPClient)
	}

	if everartMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		everartMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if everartMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		everartMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/everart",
			Version: "0.1.0",
		}
	}

	return everartMCPClient
}
