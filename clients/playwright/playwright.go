package playwright

import (
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxPlaywrightMcpServer = "npx-playwright-mcp-server"
	SsePlaywrightMcpServer = "sse-playwright-mcp-server"
)

type PlaywrightParam struct {
	Args []string

	BaseUrl string
	Options []transport.ClientOption
}

func InitPlaywrightMCPClient(p *PlaywrightParam, options ...param.Option) *param.MCPClientConf {

	playwrightMCPClient := &param.MCPClientConf{
		Name:       NpxPlaywrightMcpServer,
		ClientType: param.StdioType,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Args: []string{
				"-y",
				"@playwright/mcp@latest",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	playwrightMCPClient.StdioClientConf.Args = append(playwrightMCPClient.StdioClientConf.Args, p.Args...)

	for _, o := range options {
		o(playwrightMCPClient)
	}

	if playwrightMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		playwrightMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if playwrightMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		playwrightMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/amap-maps",
			Version: "0.1.0",
		}
	}

	return playwrightMCPClient
}

func InitPlaywrightSSEMCPClient(p *PlaywrightParam, options ...param.Option) *param.MCPClientConf {

	playwrightMCPClient := &param.MCPClientConf{
		Name:       SsePlaywrightMcpServer,
		ClientType: param.SSEType,
		SSEClientConf: &param.SSEClientConfig{
			BaseUrl: p.BaseUrl,
			Options: p.Options,
		},
	}

	for _, o := range options {
		o(playwrightMCPClient)
	}

	if playwrightMCPClient.SSEClientConf.InitReq.Params.ProtocolVersion == "" {
		playwrightMCPClient.SSEClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if playwrightMCPClient.SSEClientConf.InitReq.Params.ClientInfo.Name == "" {
		playwrightMCPClient.SSEClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/playwright",
			Version: "0.1.0",
		}
	}

	return playwrightMCPClient
}
