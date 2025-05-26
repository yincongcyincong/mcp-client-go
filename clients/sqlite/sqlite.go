package sqlite

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxSqliteMcpServer = "npx-sqlite-mcp-server"
)

type SqliteParam struct {
	SqliteDBPath string
}

func InitSqliteMCPClient(p *SqliteParam, options ...param.Option) *param.MCPClientConf {

	sqliteMCPClient := &param.MCPClientConf{
		Name: NpxSqliteMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"-y",
				"mcp-sqlite",
				p.SqliteDBPath,
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(sqliteMCPClient)
	}

	if sqliteMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		sqliteMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if sqliteMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		sqliteMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/sqlite",
			Version: "0.1.0",
		}
	}

	return sqliteMCPClient
}
