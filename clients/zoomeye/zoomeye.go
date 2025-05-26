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

func InitZoomeyeMCPClient(p *ZoomeyeParam, options ...param.Option) *param.MCPClientConf {

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
	}

	for _, o := range options {
		o(ZoomeyeMCPClient)
	}

	if ZoomeyeMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		ZoomeyeMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if ZoomeyeMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		ZoomeyeMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/zoomeye",
			Version: "0.1.0",
		}
	}

	return ZoomeyeMCPClient
}
