package firecrawl

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxFirecrawlMcpServer = "npx-firecrawl-mcp-server"
)

type FireCrawlParam struct {
	FilecrawlApiKey string
}

func InitFirecrawlMCPClient(p *FireCrawlParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	firecrawlMCPClient := &param.MCPClientConf{
		Name: NpxFirecrawlMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"FIRECRAWL_API_KEY=" + p.FilecrawlApiKey,
			},
			Args: []string{
				"-y",
				"firecrawl-mcp",
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
		Name:    "mcp-server/firecrawl",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	firecrawlMCPClient.StdioClientConf.InitReq = initRequest

	return firecrawlMCPClient
}
