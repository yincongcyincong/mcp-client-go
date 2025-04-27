package aliyun

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	UvxAliyunMcpServer = "uvx-aliyun-mcp-server"
)

type AliyunParams struct {
	AliyunAccessKeyID     string
	AliyunAccessKeySecret string
}

func InitAliyunMCPClient(p *AliyunParams, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	awsMCPClient := &param.MCPClientConf{
		Name: UvxAliyunMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uvx",
			Env: []string{
				"ALIBABA_CLOUD_ACCESS_KEY_ID=" + p.AliyunAccessKeyID,
				"ALIBABA_CLOUD_ACCESS_KEY_SECRET=" + p.AliyunAccessKeySecret,
			},
			Args: []string{
				"alibaba-cloud-ops-mcp-server@latest",
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
		Name:    "mcp-server/aws-core",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	awsMCPClient.StdioClientConf.InitReq = initRequest

	return awsMCPClient
}
