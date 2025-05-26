package time

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	UvTimeMcpServer     = "uv-time-mcp-server"
	DockerTimeMcpServer = "docker-time-mcp-server"
)

type TimeParam struct {
	LocalTimezone string
}

func InitTimeMCPClient(p *TimeParam, options ...param.Option) *param.MCPClientConf {

	if p.LocalTimezone == "" {
		p.LocalTimezone = "Asia/Shanghai"
	}

	timeMCPClient := &param.MCPClientConf{
		Name: UvTimeMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uvx",
			Env:     []string{},
			Args: []string{
				"mcp-server-time",
				"--local-timezone=" + p.LocalTimezone,
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(timeMCPClient)
	}

	if timeMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		timeMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if timeMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		timeMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/time",
			Version: "0.1.0",
		}
	}

	return timeMCPClient
}

func InitDockerTimeMCPClient(p *TimeParam, options ...param.Option) *param.MCPClientConf {

	if p.LocalTimezone == "" {
		p.LocalTimezone = "Asia/Shanghai"
	}

	timeMCPClient := &param.MCPClientConf{
		Name: DockerTimeMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uvx",
			Env:     []string{},
			Args: []string{
				"run",
				"-i",
				"--rm",
				"mcp/time",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(timeMCPClient)
	}

	if timeMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		timeMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if timeMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		timeMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/time",
			Version: "0.1.0",
		}
	}

	return timeMCPClient
}
