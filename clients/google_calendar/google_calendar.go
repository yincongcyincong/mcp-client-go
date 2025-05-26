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

func InitGoogleCalendarMCPClient(p *GoogleCalendarParam, options ...param.Option) *param.MCPClientConf {

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
	}

	for _, o := range options {
		o(googleCalendarMCPClient)
	}

	if googleCalendarMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		googleCalendarMCPClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}

	if googleCalendarMCPClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		googleCalendarMCPClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/googleCalendar",
			Version: "0.1.0",
		}
	}

	return googleCalendarMCPClient
}
