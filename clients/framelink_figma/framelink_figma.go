package framelink_figma

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxFigmaMcpServer = "npx-figma-mcp-server"
)

func InitFigmaMCPClient(figmaApiKey, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	figmaMCPClient := &param.MCPClientConf{
		Name:    NpxFigmaMcpServer,
		Command: "npx",
		Env:     []string{},
		Args: []string{
			"-y",
			"figma-developer-mcp",
			"--figma-api-key=" + figmaApiKey,
			"--stdio",
		},
		InitReq:         mcp.InitializeRequest{},
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/figma",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	figmaMCPClient.InitReq = initRequest

	return figmaMCPClient
}
