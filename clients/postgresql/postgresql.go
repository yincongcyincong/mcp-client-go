package postgresql

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxPostgresqlMcpServer    = "npx-postgresql-mcp-server"
	DockerPostgresqlMcpServer = "docker-postgresql-mcp-server"
)

func InitPostgresqlMCPClient(postgresqlLink, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	postgresqlMCPClient := &param.MCPClientConf{
		Name:    NpxPostgresqlMcpServer,
		Command: "npx",
		Env:     []string{},
		Args: []string{
			"-y",
			"@modelcontextprotocol/server-postgres",
			postgresqlLink,
		},
		InitReq:         mcp.InitializeRequest{},
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
	postgresqlMCPClient.InitReq = initRequest

	return postgresqlMCPClient
}

func InitDockerPostgresqlMCPClient(postgresqlLink, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	postgresqlMCPClient := &param.MCPClientConf{
		Name:    DockerPostgresqlMcpServer,
		Command: "docker",
		Env:     []string{},
		Args: []string{
			"run",
			"-i",
			"--rm",
			"mcp/postgres",
			postgresqlLink,
		},
		InitReq:         mcp.InitializeRequest{},
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
	postgresqlMCPClient.InitReq = initRequest

	return postgresqlMCPClient
}
