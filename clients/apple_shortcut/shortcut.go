package apple_shortcut

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxAppleShortcutMapsMcpServer = "npx-apple-shortcut-maps-mcp-server"
)

type AppleShortcutParam struct {
}

func InitAppleShortcutMCPClient(p *AppleShortcutParam, options ...param.Option) *param.MCPClientConf {

	appleShortcutMCPClient := &param.MCPClientConf{
		Name: NpxAppleShortcutMapsMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"-y",
				"mcp-server-apple-shortcuts\"",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(appleShortcutMCPClient)
	}

	if appleShortcutMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		appleShortcutMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if appleShortcutMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		appleShortcutMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/shortcut",
			Version: "0.1.0",
		}
	}

	return appleShortcutMCPClient
}
