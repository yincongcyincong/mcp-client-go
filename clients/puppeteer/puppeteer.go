package puppeteer

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxPuppeteerMcpServer    = "npx-puppeteer-mcp-server"
	DockerPuppeteerMcpServer = "docker-puppeteer-mcp-server"
)

func InitPuppeteerMCPClient(protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	puppeteerMCPClient := &param.MCPClientConf{
		Name:    NpxPuppeteerMcpServer,
		Command: "npx",
		Env:     []string{},
		Args: []string{
			"-y",
			"@modelcontextprotocol/server-puppeteer",
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
		Name:    "mcp-server/puppeteer",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	puppeteerMCPClient.InitReq = initRequest

	return puppeteerMCPClient
}

func InitDockerPuppeteerMCPClient(protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	puppeteerMCPClient := &param.MCPClientConf{
		Name:    DockerPuppeteerMcpServer,
		Command: "docker",
		Env:     []string{},
		Args: []string{
			"run",
			"-i",
			"--rm",
			"--init",
			"-e", "DOCKER_CONTAINER=true",
			"mcp/puppeteer",
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
		Name:    "mcp-server/puppeteer",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	puppeteerMCPClient.InitReq = initRequest

	return puppeteerMCPClient
}
