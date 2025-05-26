package blender

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	UvxBlenderMcpServer = "uvx-blender-mcp-server"
)

type BlenderParam struct {
}

func InitBlenderMCPClient(p BlenderParam, options ...param.Option) *param.MCPClientConf {

	blenderMCPClient := &param.MCPClientConf{
		Name: UvxBlenderMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uvx",
			Env:     []string{},
			Args: []string{
				"blender-mcp",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(blenderMCPClient)
	}

	if blenderMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		blenderMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if blenderMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		blenderMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/blender",
			Version: "0.1.0",
		}
	}

	return blenderMCPClient
}
