package shopify

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxShopifyMcpServer = "npx-shopify-mcp-server"
)

type ShopifyParam struct{}

func InitShopifyMCPClient(p *ShopifyParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	ShopifyMCPClient := &param.MCPClientConf{
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
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/shopify",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	ShopifyMCPClient.StdioClientConf.InitReq = initRequest

	return ShopifyMCPClient
}
