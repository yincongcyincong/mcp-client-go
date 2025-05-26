package cloudflare

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxCloudflareMcpServer = "npx-cloudflare-mcp-server"
)

type CloudflareParam struct {
}

func InitCloudflareMCPClient(p *CloudflareParam, options ...param.Option) *param.MCPClientConf {

	cloudflareMCPClient := &param.MCPClientConf{
		Name:       NpxCloudflareMcpServer,
		ClientType: param.StdioType,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"mcp-remote",
				"https://observability.mcp.cloudflare.com/sse",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(cloudflareMCPClient)
	}

	if cloudflareMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		cloudflareMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if cloudflareMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		cloudflareMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/cloudflare",
			Version: "0.1.0",
		}
	}

	return cloudflareMCPClient
}
