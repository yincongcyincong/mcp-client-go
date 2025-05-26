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

func InitLeetcodeMCPClient(p *LeetcodeParam, options ...param.Option) *param.MCPClientConf {

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
	}

	for _, o := range options {
		o(leetcodeMCPClient)
	}

	if leetcodeMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		leetcodeMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if leetcodeMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		leetcodeMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/leetcode",
			Version: "0.1.0",
		}
	}

	return leetcodeMCPClient
}
