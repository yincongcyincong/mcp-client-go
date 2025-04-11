package googlemap

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxGooglemapMcpServer    = "npx-googlemap-mcp-server"
	DockerGooglemapMcpServer = "docker-googlemap-mcp-server"
)

type GoogleMapParam struct {
	GooglemapApiKey string
}

func InitGooglemapMCPClient(p *GoogleMapParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	googlemapMCPClient := &param.MCPClientConf{
		Name: NpxGooglemapMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"GOOGLE_MAPS_API_KEY=" + p.GooglemapApiKey,
			},
			Args: []string{
				"-y",
				"@modelcontextprotocol/server-google-maps",
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
		Name:    "mcp-server/googlemap",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	googlemapMCPClient.StdioClientConf.InitReq = initRequest

	return googlemapMCPClient
}

func InitDockerGooglemapMCPClient(p *GoogleMapParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	googlemapMCPClient := &param.MCPClientConf{
		Name: DockerGooglemapMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "docker",
			Env: []string{
				"GOOGLE_MAPS_API_KEY=" + p.GooglemapApiKey,
			},
			Args: []string{
				"run",
				"-i",
				"--rm",
				"-e",
				"GOOGLE_MAPS_API_KEY",
				"mcp/google-maps",
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
		Name:    "mcp-server/googlemap",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	googlemapMCPClient.StdioClientConf.InitReq = initRequest

	return googlemapMCPClient
}
