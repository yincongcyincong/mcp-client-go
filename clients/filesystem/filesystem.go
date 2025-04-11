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

func InitFilesystemMCPClient(paths []string, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	filesystemMCPClient := &param.MCPClientConf{
		Name: NpxFilesystemMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"-y",
				"@modelcontextprotocol/server-filesystem",
			},
			InitReq: mcp.InitializeRequest{},
		},
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	filesystemMCPClient.StdioClientConf.Args = append(filesystemMCPClient.StdioClientConf.Args, paths...)

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
	filesystemMCPClient.StdioClientConf.InitReq = initRequest

	return filesystemMCPClient
}

func InitDockerFilesystemMCPClient(pathPair map[string]string, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	filesystemMCPClient := &param.MCPClientConf{
		Name: DockerFilesystemServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "docker",
			Env:     []string{},
			Args:    []string{},
			InitReq: mcp.InitializeRequest{},
		},
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	args := []string{
		"run",
		"-i",
		"--rm",
	}
	for srcPath, dstPath := range pathPair {
		args = append(args, "--mount", fmt.Sprintf("type=bind,src=%s,dst=%s", srcPath, dstPath))

	}
	args = append(args, "mcp/filesystem", "/projects")
	filesystemMCPClient.StdioClientConf.Args = args

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
	filesystemMCPClient.StdioClientConf.InitReq = initRequest

	return filesystemMCPClient
}
