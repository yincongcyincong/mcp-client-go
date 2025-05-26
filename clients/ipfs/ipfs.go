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

func InitIpfsMCPClient(p *IpfsParam, options ...param.Option) *param.MCPClientConf {

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
	}

	for _, o := range options {
		o(ipfsMCPClient)
	}

	if ipfsMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		ipfsMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if ipfsMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		ipfsMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/ipfs",
			Version: "0.1.0",
		}
	}

	return ipfsMCPClient
}
