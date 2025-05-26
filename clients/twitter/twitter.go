package twitter

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxTwitterMcpServer = "npx-twitter-mcp-server"
)

type TwitterParam struct {
	ApiKey            string
	ApiSecretKey      string
	AccessToken       string
	AccessTokenSecret string
}

func InitTwitterMCPClient(p *TwitterParam, options ...param.Option) *param.MCPClientConf {

	twitterMCPClient := &param.MCPClientConf{
		Name: NpxTwitterMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"API_KEY=" + p.ApiKey,
				"API_SECRET_KEY=" + p.ApiSecretKey,
				"ACCESS_TOKEN=" + p.AccessToken,
				"ACCESS_TOKEN_SECRET=" + p.AccessTokenSecret,
			},
			Args: []string{
				"-y",
				"@enescinar/twitter-mcp",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(twitterMCPClient)
	}

	if twitterMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		twitterMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if twitterMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		twitterMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/twitter",
			Version: "0.1.0",
		}
	}

	return twitterMCPClient
}
