package time

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	UvTimeMcpServer     = "uv-time-mcp-server"
	DockerTimeMcpServer = "docker-time-mcp-server"
)

type TimeParma struct {
	LocalTimezone string
}

func InitTimeMCPClient(p *TimeParma, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

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
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/time",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	timeMCPClient.StdioClientConf.InitReq = initRequest

	return timeMCPClient
}

func InitDockerTimeMCPClient(p *TimeParma, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

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
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/time",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	timeMCPClient.StdioClientConf.InitReq = initRequest

	return timeMCPClient
}
