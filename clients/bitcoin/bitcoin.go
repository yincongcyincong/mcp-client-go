package bitcoin

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxBitcoinMcpServer = "npx-bitcoin-mcp-server"
)

type BitcoinParam struct {
}

func InitBitcoinMCPClient(p *BitcoinParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	bitcoinMCPClient := &param.MCPClientConf{
		Name: NpxBitcoinMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"-y",
				"bitcoin-mcp@latest",
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
		Name:    "mcp-server/bitcoin",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	bitcoinMCPClient.StdioClientConf.InitReq = initRequest

	return bitcoinMCPClient
}
