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

func InitTavilyMCPClient(p *TavilyParam, options ...param.Option) *param.MCPClientConf {

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
	}

	for _, o := range options {
		o(tavilyMCPClient)
	}

	if tavilyMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		tavilyMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if tavilyMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		tavilyMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/tavily",
			Version: "0.1.0",
		}
	}

	return tavilyMCPClient
}
