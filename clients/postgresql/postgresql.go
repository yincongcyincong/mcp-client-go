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

func InitPostgresqlMCPClient(p *PostgreSQLParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

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
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/postgresql",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	postgresqlMCPClient.StdioClientConf.InitReq = initRequest

	return postgresqlMCPClient
}

func InitDockerPostgresqlMCPClient(p *PostgreSQLParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

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
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/postgresql",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	postgresqlMCPClient.StdioClientConf.InitReq = initRequest

	return postgresqlMCPClient
}
