package victoriametrics

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxVictoriaMetricsMcpServer = "npx-victoriametrics-mcp-server"
)

type VictoriaMetricsParam struct {
	VMUrl       string
	VMSelectUrl string
	VMInsertUrl string
}

func InitVictoriaMetricsMCPClient(p *VictoriaMetricsParam, options ...param.Option) *param.MCPClientConf {

	victoriametricsMCPClient := &param.MCPClientConf{
		Name: NpxVictoriaMetricsMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"VM_URL=" + p.VMUrl,
				"VM_SELECT_URL=" + p.VMSelectUrl,
				"VM_INSERT_URL=" + p.VMInsertUrl,
			},
			Args: []string{
				"-y",
				"@yincongcyincong/victoriametrics-mcp-server",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(victoriametricsMCPClient)
	}

	if victoriametricsMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		victoriametricsMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if victoriametricsMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		victoriametricsMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/victoriametrics",
			Version: "0.1.0",
		}
	}

	return victoriametricsMCPClient
}
