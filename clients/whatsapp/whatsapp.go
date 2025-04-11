package whatsapp

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	UvWhatsAppMcpServer = "nv-whatsapp-mcp-server"
)

func InitWhatsappMCPClient(whatsappPath, pythonMainFile string, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	whatsappMCPClient := &param.MCPClientConf{
		Name: UvWhatsAppMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uv",
			Env:     []string{},
			Args: []string{
				"--directory",
				whatsappPath, // cd into the repo, run `pwd` and enter the output here + "/whatsapp-mcp-server"
				"run",
				pythonMainFile,
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
		Name:    "mcp-server/whatsapp",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	whatsappMCPClient.StdioClientConf.InitReq = initRequest

	return whatsappMCPClient
}
