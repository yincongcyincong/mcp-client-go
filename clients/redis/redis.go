package redis

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxRedisMcpServer    = "npx-redis-mcp-server"
	DockerRedisMcpServer = "docker-redis-mcp-server"
)

type RedisParam struct {
	RedisPath string
}

func InitRedisMCPClient(p *RedisParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	redisMCPClient := &param.MCPClientConf{
		Name: NpxRedisMcpServer,

		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Args: []string{
				"-y",
				"@modelcontextprotocol/server-redis",
				p.RedisPath,
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
		Name:    "mcp-server/redis",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	redisMCPClient.StdioClientConf.InitReq = initRequest

	return redisMCPClient
}

func InitDockerRedisMCPClient(p *RedisParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	redisMCPClient := &param.MCPClientConf{
		Name: DockerRedisMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "docker",
			Args: []string{
				"run",
				"-i",
				"--rm",
				"mcp/redis",
				p.RedisPath,
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
		Name:    "mcp-server/redis",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	redisMCPClient.StdioClientConf.InitReq = initRequest

	return redisMCPClient
}
