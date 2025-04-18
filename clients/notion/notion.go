package notion

import (
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxNotionMcpServer = "npx-notion-mcp-server"
)

type NotionParam struct {
	Authorization string
	NotionVersion string
}

func InitNotionMCPClient(p *NotionParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	notionMCPClient := &param.MCPClientConf{
		Name:       NpxNotionMcpServer,
		ClientType: param.StdioType,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uvx",
			Env: []string{
				"OPENAPI_MCP_HEADERS=" + fmt.Sprintf(`{\"Authorization\": \"Bearer %s\", \"Notion-Version\": \"%s\" }`, p.Authorization, p.NotionVersion),
			},
			Args: []string{
				"mcp-notion",
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
		Name:    "mcp-server/notion",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	notionMCPClient.StdioClientConf.InitReq = initRequest

	return notionMCPClient
}
