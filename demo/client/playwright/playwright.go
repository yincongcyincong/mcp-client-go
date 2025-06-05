package main

import (
	"context"
	"log"
	"time"

	"github.com/yincongcyincong/mcp-client-go/clients"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
	"github.com/yincongcyincong/mcp-client-go/clients/playwright"
)

func main() {
	// execute `npx @playwright/mcp@latest --port 8931` first
	//mc := playwright.InitPlaywrightSSEMCPClient(&playwright.PlaywrightParam{
	//	BaseUrl: "http://localhost:8931/sse",
	//})
	//
	//// Create context with timeout
	//ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	//defer cancel()
	//
	//errs := clients.RegisterMCPClient(ctx, []*param.MCPClientConf{mc})
	//if len(errs) > 0 {
	//	log.Fatal("InitMCPClient failed:", errs)
	//}
	//
	//c, err := clients.GetMCPClient(playwright.SsePlaywrightMcpServer)
	//if err != nil {
	//	log.Fatal("GetMCPClient failed:", err)
	//}
	//
	//for _, tool := range c.Tools {
	//	toolByte, _ := json.Marshal(tool)
	//	fmt.Println(string(toolByte))
	//}
	//
	//data, err := c.ExecTools(ctx, "browser_navigate", map[string]interface{}{
	//	"url": "http://localhost:8931/sse",
	//})
	//if err != nil {
	//	log.Fatal("ExecTools failed:", err)
	//}
	//
	//fmt.Println(data)

	heartbeat()

}

func heartbeat() {
	mc := playwright.InitPlaywrightSSEMCPClient(&playwright.PlaywrightParam{
		BaseUrl: "http://localhost:8931/sse",
	})

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	errs := clients.RegisterMCPClient(ctx, []*param.MCPClientConf{mc})
	if len(errs) > 0 {
		log.Fatal("InitMCPClient failed:", errs)
	}

	for {
		c, err := clients.GetMCPClient(playwright.SsePlaywrightMcpServer)
		if err != nil {
			log.Println("GetMCPClient failed:", err)
		}

		err = c.Client.Ping(ctx)
		if err != nil {
			log.Println("Ping failed:", err)
		}

		time.Sleep(5 * time.Second)
	}
}
