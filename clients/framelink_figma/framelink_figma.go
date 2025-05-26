package framelink_figma

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxFigmaMcpServer = "npx-figma-mcp-server"
)

type FramelinkFigmaParam struct {
	FigmaApiKey string
}

func InitFigmaMCPClient(p *FramelinkFigmaParam, options ...param.Option) *param.MCPClientConf {

	figmaMCPClient := &param.MCPClientConf{
		Name: NpxFigmaMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"-y",
				"figma-developer-mcp",
				"--figma-api-key=" + p.FigmaApiKey,
				"--stdio",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(figmaMCPClient)
	}

	if figmaMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		figmaMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if figmaMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		figmaMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/figma",
			Version: "0.1.0",
		}
	}

	return figmaMCPClient
}
