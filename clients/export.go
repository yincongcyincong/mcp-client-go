package clients

import (
	"context"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

func RemoveMCPClient(name string) error {
	return defaultClients.RemoveMCPClient(name)
}

func ClearAllMCPClient() {
	defaultClients.ClearAllMCPClient()
}

// RegisterMCPClient register multiple MCP clients
func RegisterMCPClient(ctx context.Context, params []*param.MCPClientConf) map[string]error {
	return defaultClients.RegisterMCPClient(ctx, params)

} // GetMCPClientByToolName Get MCP client by tool name
func GetMCPClientByToolName(toolName string) (*MCPClient, error) {
	return defaultClients.GetMCPClientByToolName(toolName)
}

// GetMCPClient Get MCP client by name
func GetMCPClient(name string) (*MCPClient, error) {
	return defaultClients.GetMCPClient(name)
}
