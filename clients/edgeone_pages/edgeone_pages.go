package edgeone_pages

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxEdgeoneMcpServer = "npx-edgeone-mcp-server"
)

type EdgeOneParam struct {
}

func InitEdgeoneMCPClient(p *EdgeOneParam, options ...param.Option) *param.MCPClientConf {

	edgeoneMCPClient := &param.MCPClientConf{
		Name: NpxEdgeoneMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"edgeone-pages-mcp",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(edgeoneMCPClient)
	}

	if edgeoneMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		edgeoneMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if edgeoneMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		edgeoneMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/edgeone",
			Version: "0.1.0",
		}
	}

	return edgeoneMCPClient
}
