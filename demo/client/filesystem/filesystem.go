package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/yincongcyincong/mcp-client-go/clients"
	"github.com/yincongcyincong/mcp-client-go/clients/filesystem"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
	"log"
	"time"
)

func main() {
	mc := filesystem.InitFilesystemMCPClient([]string{"/Users/yincong/"}, "", nil, nil, nil)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := clients.RegisterMCPClient(ctx, []*param.MCPClientConf{mc})
	if err != nil {
		log.Fatal("InitMCPClient failed:", err)
	}

	c, err := clients.GetMCPClient(filesystem.NpxFilesystemMcpServer)
	if err != nil {
		log.Fatal("GetMCPClient failed:", err)
	}

	for _, tool := range c.Tools {
		toolByte, _ := json.Marshal(tool)
		fmt.Println(string(toolByte))
	}

	data, err := c.ExecTools(ctx, "list_directory", map[string]interface{}{
		"path": "/Users/yincong/go/src/github.com/yincongcyincong",
	})
	if err != nil {
		log.Fatal("ExecTools failed:", err)
	}

	fmt.Println(data)

}
