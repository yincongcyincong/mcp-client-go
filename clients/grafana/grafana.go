package grafana

import (
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	SSEGrafanaMcpServer = "sse-grafana-mcp-server"
)

type GrafanaParam struct {
	BaseUrl string
	Options []transport.ClientOption
}

func InitGrafanaSSEMCPClient(p *GrafanaParam, options ...param.Option) *param.MCPClientConf {

	grafanaMCPClient := &param.MCPClientConf{
		Name:       SSEGrafanaMcpServer,
		ClientType: param.SSEType,
		SSEClientConf: &param.SSEClientConfig{
			BaseUrl: p.BaseUrl,
			Options: p.Options,
		},
	}

	for _, o := range options {
		o(grafanaMCPClient)
	}

	if grafanaMCPClient.SSEClientConf.InitReq.Params.ProtocolVersion == "" {
		grafanaMCPClient.SSEClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if grafanaMCPClient.SSEClientConf.InitReq.Params.ClientInfo.Name == "" {
		grafanaMCPClient.SSEClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/grafana",
			Version: "0.1.0",
		}
	}

	return grafanaMCPClient
}
