package puppeteer

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxPuppeteerMcpServer    = "npx-puppeteer-mcp-server"
	DockerPuppeteerMcpServer = "docker-puppeteer-mcp-server"
)

type PuppeteerParam struct {
}

func InitPuppeteerMCPClient(p *PuppeteerParam, options ...param.Option) *param.MCPClientConf {

	puppeteerMCPClient := &param.MCPClientConf{
		Name: NpxPuppeteerMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env:     []string{},
			Args: []string{
				"-y",
				"@modelcontextprotocol/server-puppeteer",
			},
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(puppeteerMCPClient)
	}

	if puppeteerMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		puppeteerMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if puppeteerMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		puppeteerMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/puppeteer",
			Version: "0.1.0",
		}
	}

	return puppeteerMCPClient
}

func InitDockerPuppeteerMCPClient(p *PuppeteerParam, options ...param.Option) *param.MCPClientConf {

	puppeteerMCPClient := &param.MCPClientConf{
		Name: DockerPuppeteerMcpServer,
		StdioClientConf: &param.StdioClientConfig{
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
			InitReq: mcp.InitializeRequest{},
		},
	}

	for _, o := range options {
		o(puppeteerMCPClient)
	}

	if puppeteerMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		puppeteerMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if puppeteerMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		puppeteerMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/puppeteer",
			Version: "0.1.0",
		}
	}

	return puppeteerMCPClient
}
