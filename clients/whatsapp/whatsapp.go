package whatsapp

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	UvWhatsAppMcpServer = "nv-whatsapp-mcp-server"
)

type WhaPsAppParam struct {
	WhatsappPath   string
	PythonMainFile string
}

func InitWhatsappMCPClient(p *WhaPsAppParam, options ...param.Option) *param.MCPClientConf {

	whatsappMCPClient := &param.MCPClientConf{
		Name: UvWhatsAppMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uv",
			Env:     []string{},
			Args: []string{
				"--directory",
				p.WhatsappPath, // cd into the repo, run `pwd` and enter the output here + "/whatsapp-mcp-server"
				"run",
				p.PythonMainFile,
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(whatsappMCPClient)
	}

	if whatsappMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		whatsappMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if whatsappMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		whatsappMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/whatsapp",
			Version: "0.1.0",
		}
	}

	return whatsappMCPClient
}
