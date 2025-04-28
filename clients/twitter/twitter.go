package twitter

import (
    "github.com/mark3labs/mcp-go/mcp"
    "github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
    NpxTwitterMcpServer = "npx-twitter-mcp-server"
)

type TwitterParam struct {
    ApiKey            string
    ApiSecretKey      string
    AccessToken       string
    AccessTokenSecret string
}

func InitTwitterMCPClient(p *TwitterParam, protocolVersion string, clientInfo *mcp.Implementation,
    toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
    toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

    twitterMCPClient := &param.MCPClientConf{
        Name: NpxTwitterMcpServer,
        StdioClientConf: &param.StdioClientConfig{
            Command: "npx",
            Env: []string{
                "API_KEY=" + p.ApiKey,
                "API_SECRET_KEY=" + p.ApiSecretKey,
                "ACCESS_TOKEN=" + p.AccessToken,
                "ACCESS_TOKEN_SECRET=" + p.AccessTokenSecret,
            },
            Args: []string{
                "-y",
                "@enescinar/twitter-mcp",
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
        Name:    "mcp-server/twitter",
        Version: "0.1.0",
    }
    if clientInfo != nil {
        initRequest.Params.ClientInfo = *clientInfo
    }
    twitterMCPClient.StdioClientConf.InitReq = initRequest

    return twitterMCPClient
}
