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

func InitAliyunMCPClient(p *AliyunParams, options ...param.Option) *param.MCPClientConf {

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
	}

	for _, o := range options {
		o(awsMCPClient)
	}

	if awsMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		awsMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if awsMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		awsMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/aws-core",
			Version: "0.1.0",
		}
	}

	return awsMCPClient
}
