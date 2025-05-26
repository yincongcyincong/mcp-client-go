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

func InitBitcoinMCPClient(p *BitcoinParam, options ...param.Option) *param.MCPClientConf {

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
	}

	for _, o := range options {
		o(bitcoinMCPClient)
	}

	if bitcoinMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		bitcoinMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if bitcoinMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		bitcoinMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/bitcoin",
			Version: "0.1.0",
		}
	}

	return bitcoinMCPClient
}
