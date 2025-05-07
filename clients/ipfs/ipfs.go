package ipfs

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxIpfsMcpServer = "npx-ipfs-mcp-server"
)

type IpfsParam struct {
	W3LoginEmail string
}

func InitIpfsMCPClient(p *IpfsParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	ipfsMCPClient := &param.MCPClientConf{
		Name: NpxIpfsMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"W3_LOGIN_EMAIL=" + p.W3LoginEmail,
			},
			Args: []string{
				"-y",
				"ipfs-mcp",
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
		Name:    "mcp-server/ipfs",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	ipfsMCPClient.StdioClientConf.InitReq = initRequest

	return ipfsMCPClient
}
