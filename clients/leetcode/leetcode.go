package leetcode

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxLeetcodeMcpServer = "npx-leetcode-mcp-server"
)

type LeetcodeParam struct {
	LeetcodeSession string
}

func InitLeetcodeMCPClient(p *LeetcodeParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	leetcodeMCPClient := &param.MCPClientConf{
		Name: NpxLeetcodeMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"LEETCODE_SITE=global",
				"LEETCODE_SESSION=" + p.LeetcodeSession,
			},
			Args: []string{
				"-y",
				"@jinzcdev/leetcode-mcp-server",
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
		Name:    "mcp-server/leetcode",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	leetcodeMCPClient.StdioClientConf.InitReq = initRequest

	return leetcodeMCPClient
}
