package opgg

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxOpggMcpServer = "npx-opgg-mcp-server"
)

type OpggParam struct {
}

func InitOpggMCPClient(p *OpggParam, options ...param.Option) *param.MCPClientConf {
	opggMCPClient := &param.MCPClientConf{
		Name:       NpxOpggMcpServer,
		ClientType: param.StdioType,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"-y",
				"supergateway",
				"--streamableHttp",
				"https://mcp-api.op.gg/mcp",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(opggMCPClient)
	}

	if opggMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		opggMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if opggMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		opggMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/opgg",
			Version: "0.1.0",
		}
	}

	return opggMCPClient
}
