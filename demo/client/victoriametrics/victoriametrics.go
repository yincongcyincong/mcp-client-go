package main

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients"
)

func main() {
	conf := clients.InitStdioMCPClient("npx-victoria-metrics-mcp-server", "npx", []string{
		"AMAP_MAPS_API_KEY=" + AmapApiKey,
	}, []string{
		"-y",
		"@amap/amap-maps-mcp-server",
	}, mcp.InitializeRequest{}, nil, nil)

}
