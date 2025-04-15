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
	conf := clients.InitStdioMCPClient("npx-amap-maps-mcp-server", "npx", []string{
		"AMAP_MAPS_API_KEY=" + "xxx",
	}, []string{
		"-y",
		"@amap/amap-maps-mcp-server",
	}, mcp.InitializeRequest{}, nil, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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
