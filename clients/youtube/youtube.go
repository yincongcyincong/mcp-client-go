package youtube

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxYoutubeMcpServer = "npx-youtube-mcp-server"
)

type YoutubeParam struct {
	YoutubeApiKey string
}

func InitYoutubeMCPClient(p *YoutubeParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	YoutubeMCPClient := &param.MCPClientConf{
		Name: NpxYoutubeMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"YOUTUBE_API_KEY=" + p.YoutubeApiKey,
			},
			Args: []string{
				"-y",
				"@modelcontextprotocol/server-youtube",
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
		Name:    "mcp-server/youtube",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	YoutubeMCPClient.StdioClientConf.InitReq = initRequest

	return YoutubeMCPClient
}
