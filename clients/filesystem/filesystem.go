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

type FilesystemParam struct {
	Paths     []string
	PathPairs map[string]string
}

func InitFilesystemMCPClient(p *FilesystemParam, options ...param.Option) *param.MCPClientConf {

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
	}

	filesystemMCPClient.StdioClientConf.Args = append(filesystemMCPClient.StdioClientConf.Args, p.Paths...)

	for _, o := range options {
		o(filesystemMCPClient)
	}

	if filesystemMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		filesystemMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if filesystemMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		filesystemMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/filesystem",
			Version: "0.1.0",
		}
	}

	return filesystemMCPClient
}

func InitDockerFilesystemMCPClient(p *FilesystemParam, options ...param.Option) *param.MCPClientConf {

	filesystemMCPClient := &param.MCPClientConf{
		Name: DockerFilesystemServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "docker",
			Env:     []string{},
			Args:    []string{},
			InitReq: mcp.InitializeRequest{},
		},
	}

	args := []string{
		"run",
		"-i",
		"--rm",
	}
	for srcPath, dstPath := range p.PathPairs {
		args = append(args, "--mount", fmt.Sprintf("type=bind,src=%s,dst=%s", srcPath, dstPath))

	}
	args = append(args, "mcp/filesystem", "/projects")
	filesystemMCPClient.StdioClientConf.Args = args

	for _, o := range options {
		o(filesystemMCPClient)
	}

	if filesystemMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		filesystemMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if filesystemMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		filesystemMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/filesystem",
			Version: "0.1.0",
		}
	}

	return filesystemMCPClient
}
