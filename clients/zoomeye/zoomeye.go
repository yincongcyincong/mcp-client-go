package zoomeye

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxZoomeyeMcpServer = "docker-zoomeye-mcp-server"
)

type ZoomeyeParam struct {
	ZoomeyeApiKey string
}

func InitZoomeyeMCPClient(p *ZoomeyeParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	ZoomeyeMCPClient := &param.MCPClientConf{
		Name: NpxZoomeyeMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "docker",
			Env: []string{
				"YOUTUBE_API_KEY=" + p.ZoomeyeApiKey,
			},
			Args: []string{
				"run", "-i",
				"--rm", "-e", "" +
					"ZOOMEYE_API_KEY=" + p.ZoomeyeApiKey,
				"zoomeyeteam/mcp-server-zoomeye:latest",
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
		Name:    "mcp-server/zoomeye",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	ZoomeyeMCPClient.StdioClientConf.InitReq = initRequest

	return ZoomeyeMCPClient
}
