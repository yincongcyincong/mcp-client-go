package airbnb

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxAirbnbMcpServer = "npx-airbnb-mcp-server"
)

type AirbnbParam struct {
}

func InitAirbnbMCPClient(p *AirbnbParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	airbnbMCPClient := &param.MCPClientConf{
		Name: NpxAirbnbMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"-y",
				"@openbnb/mcp-server-airbnb",
				"--ignore-robots-txt",
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
		Name:    "mcp-server/airbnb",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	airbnbMCPClient.StdioClientConf.InitReq = initRequest

	return airbnbMCPClient
}
