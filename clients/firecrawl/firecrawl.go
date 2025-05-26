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

func InitFirecrawlMCPClient(p *FireCrawlParam, options ...param.Option) *param.MCPClientConf {

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
	}

	for _, o := range options {
		o(firecrawlMCPClient)
	}

	if firecrawlMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		firecrawlMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if firecrawlMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		firecrawlMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/firecrawl",
			Version: "0.1.0",
		}
	}

	return firecrawlMCPClient
}
