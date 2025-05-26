package shopify

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxShopifyMcpServer = "npx-shopify-mcp-server"
)

type ShopifyParam struct{}

func InitShopifyMCPClient(p *ShopifyParam, options ...param.Option) *param.MCPClientConf {

	shopifyMCPClient := &param.MCPClientConf{
		Name: NpxShopifyMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"-y",
				"@shopify/dev-mcp@latest",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(shopifyMCPClient)
	}

	if shopifyMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		shopifyMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if shopifyMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		shopifyMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/shopify",
			Version: "0.1.0",
		}
	}

	return shopifyMCPClient
}
