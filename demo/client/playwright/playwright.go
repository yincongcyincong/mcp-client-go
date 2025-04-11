package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/yincongcyincong/mcp-client-go/clients"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
	"github.com/yincongcyincong/mcp-client-go/clients/playwright"
	"log"
	"time"
)

func main() {
	// execute `npx @playwright/mcp@latest --port 8931` first
	mc := playwright.InitPlaywrightSSEMCPClient("http://localhost:8931/sse", nil, "", nil, nil, nil)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	errs := clients.RegisterMCPClient(ctx, []*param.MCPClientConf{mc})
	if len(errs) > 0 {
		log.Fatal("InitMCPClient failed:", errs)
	}

	c, err := clients.GetMCPClient(playwright.NpxPlaywrightMcpServer)
	if err != nil {
		log.Fatal("GetMCPClient failed:", err)
	}

	for _, tool := range c.Tools {
		toolByte, _ := json.Marshal(tool)
		fmt.Println(string(toolByte))
	}

	data, err := c.ExecTools(ctx, "browser_tab_list", map[string]interface{}{})
	if err != nil {
		log.Fatal("ExecTools failed:", err)
	}

	fmt.Println(data)

}
