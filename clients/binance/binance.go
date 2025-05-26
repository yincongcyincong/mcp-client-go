package binance

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxBinanceMcpServer = "npx-binance-mcp-server"
)

type BinanceParam struct {
	BinanceApiKey string
}

func InitBinanceMCPClient(p *BinanceParam, options ...param.Option) *param.MCPClientConf {

	binanceMCPClient := &param.MCPClientConf{
		Name: NpxBinanceMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"BINANCE_API_KEY=" + p.BinanceApiKey,
			},
			Args: []string{
				"-y",
				"@snjyor/binance-mcp@latest",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(binanceMCPClient)
	}

	if binanceMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		binanceMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if binanceMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		binanceMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/binance",
			Version: "0.1.0",
		}
	}

	return binanceMCPClient
}
