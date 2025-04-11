package tavily

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxTavilyMcpServer = "npx-tavily-mcp-server"
)

type TavilyParam struct {
	TavilyApiKey string
}

func InitTavilyMCPClient(p *TavilyParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	tavilyMCPClient := &param.MCPClientConf{
		Name: NpxTavilyMcpServer,
		StdioClientConf: &param.StdioClientConfig{Command: "npx",
			Env: []string{
				"TAVILY_API_KEY=" + p.TavilyApiKey,
			},
			Args: []string{
				"-y",
				"tavily-mcp@0.1.4",
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
		Name:    "mcp-server/tavily",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	tavilyMCPClient.StdioClientConf.InitReq = initRequest

	return tavilyMCPClient
}
