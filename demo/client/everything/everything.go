package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
	"log"
	"time"
)

func main() {
	StdioClient()

	SSEClient()

}

func StdioClient() {
	conf := clients.InitStdioMCPClient("npx-amap-maps-mcp-server", "npx", []string{
		"AMAP_MAPS_API_KEY=" + "xxx",
	}, []string{
		"-y",
		"@amap/amap-maps-mcp-server",
	}, mcp.InitializeRequest{}, nil, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	clients.RegisterMCPClient(ctx, []*param.MCPClientConf{conf})

	c, err := clients.GetMCPClient("npx-amap-maps-mcp-server")
	if err != nil {
		log.Fatal("get client fail:", err)
	}

	for _, tool := range c.Tools {
		toolByte, _ := json.Marshal(tool)
		fmt.Println(string(toolByte))
	}

	data, err := c.ExecTools(ctx, "maps_regeocode", map[string]interface{}{
		"location": "117.1935, 39.1425",
	})
	if err != nil {
		log.Fatal("ExecTools failed:", err)
	}
	fmt.Println(data)

	data, err = c.ExecTools(ctx, "maps_ip_location", map[string]interface{}{
		"ip": "220.181.3.151",
	})
	if err != nil {
		log.Fatal("ExecTools failed:", err)
	}

	fmt.Println(data)
}

func SSEClient() {
	// execute npx @playwright/mcp@latest --port 8931
	mc := clients.InitSSEMCPClient("playwright", "http://localhost:8931/sse", nil, nil, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	errs := clients.RegisterMCPClient(ctx, []*param.MCPClientConf{mc})
	if len(errs) > 0 {
		log.Fatal("InitMCPClient failed:", errs)
	}

	c, err := clients.GetMCPClient("playwright")
	if err != nil {
		log.Fatal("GetMCPClient failed:", err)
	}

	for _, tool := range c.Tools {
		toolByte, _ := json.Marshal(tool)
		fmt.Println(string(toolByte))
	}

	data, err := c.ExecTools(ctx, "browser_navigate", map[string]interface{}{
		"url": "http://localhost:8931/sse",
	})
	if err != nil {
		log.Fatal("ExecTools failed:", err)
	}

	fmt.Println(data)
}
