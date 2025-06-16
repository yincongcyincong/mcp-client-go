package mongodb

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxMongodbMcpServer = "npx-mongodb-mcp-server"
)

type MongodbParam struct {
	MongodbURI      string
	MongodbReadOnly string
}

func InitMongodbMCPClient(p *MongodbParam, options ...param.Option) *param.MCPClientConf {

	mongodbMCPClient := &param.MCPClientConf{
		Name: NpxMongodbMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"MCP_MONGODB_URI=" + p.MongodbURI,
				"MCP_MONGODB_READONLY=" + p.MongodbReadOnly,
			},
			Args: []string{
				"-y",
				"mcp-mongo-server",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(mongodbMCPClient)
	}

	if mongodbMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		mongodbMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if mongodbMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		mongodbMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/mongodb",
			Version: "0.1.0",
		}
	}

	return mongodbMCPClient
}
