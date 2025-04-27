package bilibili

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxBilibiliMcpServer = "npx-bilibili-mcp-server"
)

type BilibiliParam struct {
}

func InitBilibiliMCPClient(p *BilibiliParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	bilibiliMCPClient := &param.MCPClientConf{
		Name: NpxBilibiliMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"-y",
				"bilibili-mcp",
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
		Name:    "mcp-server/bilibili",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	bilibiliMCPClient.StdioClientConf.InitReq = initRequest

	return bilibiliMCPClient
}
