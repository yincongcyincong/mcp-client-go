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

func InitRedisMCPClient(p *RedisParam, options ...param.Option) *param.MCPClientConf {

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
	}

	for _, o := range options {
		o(redisMCPClient)
	}

	if redisMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		redisMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if redisMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		redisMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/redis",
			Version: "0.1.0",
		}
	}

	return redisMCPClient
}

func InitDockerRedisMCPClient(p *RedisParam, options ...param.Option) *param.MCPClientConf {

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
	}

	for _, o := range options {
		o(redisMCPClient)
	}

	if redisMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		redisMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if redisMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		redisMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/redis",
			Version: "0.1.0",
		}
	}

	return redisMCPClient
}
