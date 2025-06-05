package amap

import (
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxAmapMapsMcpServer = "npx-amap-maps-mcp-server"
	UvxAmapMcpServer     = "upx-http-amap-mcp-server"
)

type AmapParam struct {
	AmapApiKey string
}

func InitAmapMCPClient(p *AmapParam, options ...param.Option) *param.MCPClientConf {
	amapMCPClient := &param.MCPClientConf{
		Name: NpxAmapMapsMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"AMAP_MAPS_API_KEY=" + p.AmapApiKey,
			},
			Args: []string{
				"-y",
				"@amap/amap-maps-mcp-server",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(amapMCPClient)
	}

	if amapMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		amapMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if amapMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		amapMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/amap-maps",
			Version: "0.1.0",
		}
	}

	return amapMCPClient
}

type AmapHttpParam struct {
	BaseURL string
	Options []transport.StreamableHTTPCOption
	Oauth   *client.OAuthConfig
}

func InitHTTPAmapMCPClient(p *AmapHttpParam, options ...param.Option) *param.MCPClientConf {
	amapMCPClient := &param.MCPClientConf{
		ClientType: param.HTTPStreamer,
		Name:       UvxAmapMcpServer,
		HTTPStreamerConf: &param.HTTPStreamerConfig{
			Options: p.Options,
			BaseURL: p.BaseURL,
			Oauth:   p.Oauth,
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(amapMCPClient)
	}

	if amapMCPClient.HTTPStreamerConf.InitReq.Params.ProtocolVersion == "" {
		amapMCPClient.HTTPStreamerConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if amapMCPClient.HTTPStreamerConf.InitReq.Params.ClientInfo.Name == "" {
		amapMCPClient.HTTPStreamerConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/amap-maps",
			Version: "0.1.0",
		}
	}

	return amapMCPClient
}
