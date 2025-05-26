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

func InitYoutubeMCPClient(p *YoutubeParam, options ...param.Option) *param.MCPClientConf {

	youtubeMCPClient := &param.MCPClientConf{
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
	}

	for _, o := range options {
		o(youtubeMCPClient)
	}

	if youtubeMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		youtubeMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if youtubeMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		youtubeMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/youtube",
			Version: "0.1.0",
		}
	}

	return youtubeMCPClient
}
