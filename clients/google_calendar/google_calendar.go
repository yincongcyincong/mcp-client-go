package google_calendar

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxGoogleCalendarMcpServer = "npx-googleCalendar-mcp-server"
)

type GoogleCalendarParam struct {
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURI  string
}

func InitGoogleCalendarMCPClient(p *GoogleCalendarParam, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	googleCalendarMCPClient := &param.MCPClientConf{
		Name: NpxGoogleCalendarMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "npx",
			Env: []string{
				"GOOGLE_CLIENT_ID=" + p.GoogleClientID,
				"GOOGLE_CLIENT_SECRET=" + p.GoogleClientSecret,
				"GOOGLE_REDIRECT_URI=" + p.GoogleRedirectURI,
			},
			Args: []string{
				"-y",
				"@takumi0706/google-calendar-mcp",
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
		Name:    "mcp-server/googleCalendar",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	googleCalendarMCPClient.StdioClientConf.InitReq = initRequest

	return googleCalendarMCPClient
}
