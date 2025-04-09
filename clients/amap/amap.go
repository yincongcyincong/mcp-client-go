package amap

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

func InitAmapMCPClient(AmapApiKey string,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	amapMCPClient := &param.MCPClientConf{
		Name:    "npx-amap",
		Command: "npx",
		Env: []string{
			"AMAP_MAPS_API_KEY=" + AmapApiKey,
		},
		Args: []string{
			"-y",
			"@amap/amap-maps-mcp-server",
		},
		InitReq:         mcp.InitializeRequest{},
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/amap-maps",
		Version: "0.1.0",
	}
	amapMCPClient.InitReq = initRequest

	return amapMCPClient
}
