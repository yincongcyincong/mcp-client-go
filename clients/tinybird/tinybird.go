package tinybird

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	UvxTinyBirdServer = "uvx-tiny-bird-mcp-server"
)

type TinyBirdParams struct {
	TBAPIURL     string
	TBAdminToken string
}

func InitTinyBirdMCPClient(p *TinyBirdParams, options ...param.Option) *param.MCPClientConf {

	awsMCPClient := &param.MCPClientConf{
		Name: UvxTinyBirdServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uvx",
			Env: []string{
				"TB_API_URL=" + p.TBAPIURL,
				"TB_ADMIN_TOKEN=" + p.TBAdminToken,
			},
			Args: []string{
				"mcp-tinybird",
				"stdio",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(awsMCPClient)
	}

	if awsMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		awsMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if awsMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		awsMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/tinybird",
			Version: "0.1.0",
		}
	}
	return awsMCPClient
}
