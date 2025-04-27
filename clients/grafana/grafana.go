package grafana

import (
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	SSEGrafanaMcpServer = "sse-grafana-mcp-server"
)

type GoogleMapParam struct{}

func InitGrafanaSSEMCPClient(baseUrl string, options []client.ClientOption,
	protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	grafanaMCPClient := &param.MCPClientConf{
		Name:       SSEGrafanaMcpServer,
		ClientType: param.SSEType,
		SSEClientConf: &param.SSEClientConfig{
			BaseUrl: baseUrl,
			Options: options,
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
		Name:    "mcp-server/grafana",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	grafanaMCPClient.SSEClientConf.InitReq = initRequest

	return grafanaMCPClient
}
