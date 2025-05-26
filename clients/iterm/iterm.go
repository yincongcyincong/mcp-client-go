package iterm

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxItermMcpServer = "npx-iterm-mcp-server"
)

type ItermParam struct {
}

func InitItermMCPClient(p *ItermParam, options ...param.Option) *param.MCPClientConf {

	itermMCPClient := &param.MCPClientConf{
		Name: NpxItermMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"-y",
				"iterm-mcp",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(itermMCPClient)
	}

	if itermMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		itermMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if itermMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		itermMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/iterm",
			Version: "0.1.0",
		}
	}

	return itermMCPClient
}
