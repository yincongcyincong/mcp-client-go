package main

import (
	"context"
	"fmt"
	"github.com/yincongcyincong/mcp-client-go/clients"
	"github.com/yincongcyincong/mcp-client-go/clients/amap"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
	"log"
	"time"
)

func main() {
	// todo modify token
	mc := amap.InitAmapMCPClient("xxxx", "", nil, nil, nil)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := clients.RegisterMCPClient(ctx, []*param.MCPClientConf{mc})
	if err != nil {
		log.Fatal("InitMCPClient failed:", err)
	}

	c, err := clients.GetMCPClient(amap.NpxAmapMapsMcpServer)
	if err != nil {
		log.Fatal("GetMCPClient failed:", err)
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
