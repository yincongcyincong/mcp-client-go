package bilibili

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxBilibiliMcpServer = "npx-bilibili-mcp-server"
)

type BilibiliParam struct {
}

func InitBilibiliMCPClient(p *BilibiliParam, options ...param.Option) *param.MCPClientConf {

	bilibiliMCPClient := &param.MCPClientConf{
		Name: NpxBilibiliMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"-y",
				"bilibili-mcp",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(bilibiliMCPClient)
	}

	if bilibiliMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		bilibiliMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if bilibiliMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		bilibiliMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/bilibili",
			Version: "0.1.0",
		}
	}

	return bilibiliMCPClient
}
