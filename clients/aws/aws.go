package aws

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxAwsMcpServer    = "npx-aws-mcp-server"
	DockerAwsMcpServer = "docker-aws-mcp-server"
)

func InitAwsMCPClient(awsAccessKey, awsSecretKey, awsRegion string, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	awsMCPClient := &param.MCPClientConf{
		Name: NpxAwsMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"AWS_ACCESS_KEY_ID=" + awsAccessKey,
				"AWS_SECRET_ACCESS_KEY=" + awsSecretKey,
				"AWS_REGION=" + awsRegion,
			},
			Args: []string{
				"-y",
				"@modelcontextprotocol/server-aws-kb-retrieval",
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
		Name:    "mcp-server/aws",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	awsMCPClient.StdioClientConf.InitReq = initRequest

	return awsMCPClient
}

func InitDockerAwsMCPClient(awsAccessKey, awsSecretKey, awsRegion string, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	awsMCPClient := &param.MCPClientConf{
		Name: DockerAwsMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"AWS_ACCESS_KEY_ID=" + awsAccessKey,
				"AWS_SECRET_ACCESS_KEY=" + awsSecretKey,
				"AWS_REGION=" + awsRegion,
			},
			Args: []string{
				"run",
				"-i",
				"--rm",
				"-e", "AWS_ACCESS_KEY_ID",
				"-e", "AWS_SECRET_ACCESS_KEY",
				"-e", "AWS_REGION",
				"mcp/aws-kb-retrieval-server",
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
		Name:    "mcp-server/aws",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	awsMCPClient.StdioClientConf.InitReq = initRequest

	return awsMCPClient
}
