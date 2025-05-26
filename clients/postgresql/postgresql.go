package postgresql

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxPostgresqlMcpServer    = "npx-postgresql-mcp-server"
	DockerPostgresqlMcpServer = "docker-postgresql-mcp-server"
)

type PostgreSQLParam struct {
	PostgresqlLink string
}

func InitPostgresqlMCPClient(p *PostgreSQLParam, options ...param.Option) *param.MCPClientConf {

	postgresqlMCPClient := &param.MCPClientConf{
		Name: NpxPostgresqlMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"-y",
				"@modelcontextprotocol/server-postgres",
				p.PostgresqlLink,
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(postgresqlMCPClient)
	}

	if postgresqlMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		postgresqlMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if postgresqlMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		postgresqlMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/postgresql",
			Version: "0.1.0",
		}
	}
	return postgresqlMCPClient
}

func InitDockerPostgresqlMCPClient(p *PostgreSQLParam, options ...param.Option) *param.MCPClientConf {

	postgresqlMCPClient := &param.MCPClientConf{
		Name: DockerPostgresqlMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "docker",
			Env:     []string{},
			Args: []string{
				"run",
				"-i",
				"--rm",
				"mcp/postgres",
				p.PostgresqlLink,
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(postgresqlMCPClient)
	}

	if postgresqlMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		postgresqlMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if postgresqlMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		postgresqlMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/postgresql",
			Version: "0.1.0",
		}
	}

	return postgresqlMCPClient
}
