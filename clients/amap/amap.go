package amap

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxAmapMapsMcpServer = "npx-amap-maps-mcp-server"
)

type AmapParam struct {
	AmapApiKey string
}

func InitAmapMCPClient(p *AmapParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

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
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/amap-maps",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	amapMCPClient.StdioClientConf.InitReq = initRequest

	return amapMCPClient
}
