package duckduckgo

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxDuckduckgoMcpServer = "npx-duckduckgo-mcp-server"
)

type DuckduckgoParam struct {
}

func InitDuckduckgoMCPClient(p *DuckduckgoParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	duckduckgoMCPClient := &param.MCPClientConf{
		Name:       NpxDuckduckgoMcpServer,
		ClientType: param.StdioType,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uvx",
			Env:     []string{},
			Args: []string{
				"duckduckgo-mcp-server",
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
		Name:    "mcp-server/duckduckgo",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	duckduckgoMCPClient.StdioClientConf.InitReq = initRequest

	return duckduckgoMCPClient
}
