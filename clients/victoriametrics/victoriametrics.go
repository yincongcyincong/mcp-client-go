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

func InitVictoriaMetricsMCPClient(p *VictoriaMetricsParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

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
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/victoriametrics",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	victoriametricsMCPClient.StdioClientConf.InitReq = initRequest

	return victoriametricsMCPClient
}
