package airbnb

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxAirbnbMcpServer = "npx-airbnb-mcp-server"
)

type AirbnbParam struct {
}

func InitAirbnbMCPClient(p *AirbnbParam, options ...param.Option) *param.MCPClientConf {

	airbnbMCPClient := &param.MCPClientConf{
		Name: NpxAirbnbMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"-y",
				"@openbnb/mcp-server-airbnb",
				"--ignore-robots-txt",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(airbnbMCPClient)
	}

	if airbnbMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		airbnbMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if airbnbMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		airbnbMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/airbnb",
			Version: "0.1.0",
		}
	}

	return airbnbMCPClient
}
