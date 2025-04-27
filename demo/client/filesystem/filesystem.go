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
	mc := filesystem.InitFilesystemMCPClient(&filesystem.FilesystemParam{
		Paths: []string{"/Users/yincong/"},
	}, "", nil, nil, nil)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	errs := clients.RegisterMCPClient(ctx, []*param.MCPClientConf{mc})
	if len(errs) > 0 {
		log.Fatal("InitMCPClient failed:", errs)
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
