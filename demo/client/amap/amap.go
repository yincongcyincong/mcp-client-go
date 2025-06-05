package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/yincongcyincong/mcp-client-go/clients"
	"github.com/yincongcyincong/mcp-client-go/clients/amap"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

func main() {

	httpStream()

}

func httpStream() {
	// todo modify token
	//mc := amap.InitHTTPAmapMCPClient(&amap.AmapHttpParam{
	//	BaseURL: "http://127.0.0.1:8000/mcp",
	//	Options: nil,
	//	Oauth:   nil,
	//})

	// execute `uvx amap-mcp-server streamable-http`
	mc := clients.InitHttpMCPClient(amap.UvxAmapMcpServer, "http://127.0.0.1:8000/mcp")

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	errs := clients.RegisterMCPClient(ctx, []*param.MCPClientConf{mc})
	if len(errs) > 0 {
		log.Fatal("InitMCPClient failed:", errs)
	}

	c, err := clients.GetMCPClient(amap.UvxAmapMcpServer)
	if err != nil {
		log.Fatal("GetMCPClient failed:", err)
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
}

func stdio() {
	// todo modify token
	mc := amap.InitAmapMCPClient(&amap.AmapParam{
		AmapApiKey: "xxx",
	})

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	errs := clients.RegisterMCPClient(ctx, []*param.MCPClientConf{mc})
	if len(errs) > 0 {
		log.Fatal("InitMCPClient failed:", errs)
	}

	c, err := clients.GetMCPClient(amap.NpxAmapMapsMcpServer)
	if err != nil {
		log.Fatal("GetMCPClient failed:", err)
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

func heartbeat() {
	mc := amap.InitAmapMCPClient(&amap.AmapParam{
		AmapApiKey: "xxx",
	})

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	errs := clients.RegisterMCPClient(ctx, []*param.MCPClientConf{mc})
	if len(errs) > 0 {
		log.Fatal("InitMCPClient failed:", errs)
	}

	for {
		c, err := clients.GetMCPClient(amap.NpxAmapMapsMcpServer)
		if err != nil {
			log.Fatal("GetMCPClient failed:", err)
		}

		err = c.Client.Ping(context.Background())
		if err != nil {
			fmt.Println("ping fail:", err)
		}

		time.Sleep(5 * time.Second)
	}
}
