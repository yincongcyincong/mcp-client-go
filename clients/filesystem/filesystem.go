package filesystem

import (
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxFilesystemMcpServer = "npx-filesystem-mcp-server"
	DockerFilesystemServer = "docker-filesystem-mcp-server"
)

func InitFilesystemMCPClient(srcPath, dstPath string, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	filesystemMCPClient := &param.MCPClientConf{
		Name:    NpxFilesystemMcpServer,
		Command: "npx",
		Env:     []string{},
		Args: []string{
			"-y",
			"@modelcontextprotocol/server-filesystem",
			srcPath,
			dstPath,
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
		Name:    "mcp-server/filesystem",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	filesystemMCPClient.InitReq = initRequest

	return filesystemMCPClient
}

func InitDockerFilesystemMCPClient(srcPath, dstPath string, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	filesystemMCPClient := &param.MCPClientConf{
		Name:    DockerFilesystemServer,
		Command: "docker",
		Env:     []string{},
		Args: []string{
			"run",
			"-i",
			"--rm",
			"--mount",
			fmt.Sprintf("type=bind,src=%s,dst=%s", srcPath, dstPath),
			"mcp/filesystem",
			"/projects",
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
		Name:    "mcp-server/filesystem",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	filesystemMCPClient.InitReq = initRequest

	return filesystemMCPClient
}
