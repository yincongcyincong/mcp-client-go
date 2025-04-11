package playwright

import (
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxPlaywrightMcpServer = "npx-playwright-mcp-server"
)

type PlaywrightParam struct {
	Args []string
}

func InitPlaywrightMCPClient(p *PlaywrightParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	playwrightMCPClient := &param.MCPClientConf{
		Name:       NpxPlaywrightMcpServer,
		ClientType: param.StdioType,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Args: []string{
				"-y",
				"@playwright/mcp@latest",
			},
			InitReq: mcp.InitializeRequest{},
		},
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	playwrightMCPClient.StdioClientConf.Args = append(playwrightMCPClient.StdioClientConf.Args, p.Args...)

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/playwright",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	playwrightMCPClient.StdioClientConf.InitReq = initRequest

	return playwrightMCPClient
}

func InitPlaywrightSSEMCPClient(baseUrl string, options []client.ClientOption,
	protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	playwrightMCPClient := &param.MCPClientConf{
		Name:       NpxPlaywrightMcpServer,
		ClientType: param.SSEType,
		SSEClientConf: &param.SSEClientConfig{
			BaseUrl: baseUrl,
			Options: options,
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
		Name:    "mcp-server/playwright",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	playwrightMCPClient.SSEClientConf.InitReq = initRequest

	return playwrightMCPClient
}
