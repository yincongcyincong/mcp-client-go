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

func InitGhidraMCPClient(p *GhidraParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

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
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/ghidra",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	ghidraMCPClient.StdioClientConf.InitReq = initRequest

	return ghidraMCPClient
}
