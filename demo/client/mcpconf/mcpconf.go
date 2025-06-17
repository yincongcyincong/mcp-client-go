package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/yincongcyincong/mcp-client-go/clients"
)

// todo start `npx @playwright/mcp@latest --port 8931` and ` uvx amap-mcp-server streamable-http` first
func main() {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	mcs, err := clients.InitByConfFile(ctx,
		"./test.json")
	if err != nil {
		log.Fatal("get conf fail:", err)
	}

	errs := clients.RegisterMCPClient(ctx, mcs)
	if len(errs) > 0 {
		log.Fatal("InitMCPClient failed:", errs)
	}

	c, err := clients.GetMCPClient("github")
	if err != nil {
		log.Fatal("GetMCPClient failed:", err)
	}

	for _, tool := range c.Tools {
		toolByte, _ := json.Marshal(tool)
		fmt.Println(string(toolByte))
	}

	fmt.Println("-------------------------------------------------------")

	c, err = clients.GetMCPClient("playwright")
	if err != nil {
		log.Fatal("GetMCPClient failed:", err)
	}

	for _, tool := range c.Tools {
		toolByte, _ := json.Marshal(tool)
		fmt.Println(string(toolByte))
	}
	fmt.Println("-------------------------------------------------------")

	c, err = clients.GetMCPClient("amap-mcp-server")
	if err != nil {
		log.Fatal("GetMCPClient failed:", err)
	}

	for _, tool := range c.Tools {
		toolByte, _ := json.Marshal(tool)
		fmt.Println(string(toolByte))
	}
}
