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

func InitTinyBirdMCPClient(p *TinyBirdParams, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

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
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/tiny-bird",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	awsMCPClient.StdioClientConf.InitReq = initRequest

	return awsMCPClient
}
