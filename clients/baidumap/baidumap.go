package baidumap

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxBaidumapMcpServer = "npx-baidumap-mcp-server"
)

type BaidumapParam struct {
	BaidumapApiKey string
}

func InitBaidumapMCPClient(p *BaidumapParam, options ...param.Option) *param.MCPClientConf {

	baidumapMCPClient := &param.MCPClientConf{
		Name: NpxBaidumapMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"BAIDU_MAP_API_KEY=" + p.BaidumapApiKey,
			},
			Args: []string{
				"-y",
				"@baidumap/mcp-server-baidu-map",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(baidumapMCPClient)
	}

	if baidumapMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		baidumapMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if baidumapMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		baidumapMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/baidumap",
			Version: "0.1.0",
		}
	}

	return baidumapMCPClient
}
