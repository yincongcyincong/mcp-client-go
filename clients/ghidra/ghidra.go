package ghidra

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxGhidraMcpServer = "npx-ghidra-mcp-server"
)

type GhidraParam struct {
	PythonPath string
	RunCommand string
	ServerUrl  string
}

func InitGhidraMCPClient(p *GhidraParam, options ...param.Option) *param.MCPClientConf {

	ghidraMCPClient := &param.MCPClientConf{
		Name: NpxGhidraMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "python",
			Env:     []string{},
			Args: []string{
				p.PythonPath,
				p.RunCommand,
				p.ServerUrl,
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(ghidraMCPClient)
	}

	if ghidraMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		ghidraMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if ghidraMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		ghidraMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/ghidra",
			Version: "0.1.0",
		}
	}

	return ghidraMCPClient
}
