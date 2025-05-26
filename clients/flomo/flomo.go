package flomo

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxFlomoMcpServer = "npx-flomo-mcp-server"
)

type FlomoParam struct {
	FilecrawlApiUrl string
}

func InitFlomoMCPClient(p *FlomoParam, options ...param.Option) *param.MCPClientConf {

	flomoMCPClient := &param.MCPClientConf{
		Name: NpxFlomoMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"FLOMO_API_URL=" + p.FilecrawlApiUrl,
			},
			Args: []string{
				"-y",
				"@chatmcp/mcp-server-flomo",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(flomoMCPClient)
	}

	if flomoMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		flomoMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if flomoMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		flomoMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/flomo",
			Version: "0.1.0",
		}
	}

	return flomoMCPClient
}
