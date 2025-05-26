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

func InitNotionMCPClient(p *NotionParam, options ...param.Option) *param.MCPClientConf {

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
	}

	for _, o := range options {
		o(notionMCPClient)
	}

	if notionMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		notionMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if notionMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		notionMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/notion",
			Version: "0.1.0",
		}
	}

	return notionMCPClient
}
