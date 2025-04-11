package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/yincongcyincong/mcp-client-go/clients"
	"github.com/yincongcyincong/mcp-client-go/clients/github"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
	"log"
	"time"
)

func main() {
	// todo modify token
	mc := github.InitModelContextProtocolGithubMCPClient("xxxx", "", nil, nil, nil)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err := clients.RegisterMCPClient(ctx, []*param.MCPClientConf{mc})
	if err != nil {
		log.Fatal("InitMCPClient failed:", err)
	}

	c, err := clients.GetMCPClient(github.NpxModelContextProtocolServerGithub)
	if err != nil {
		log.Fatal("GetMCPClient failed:", err)
	}

	for _, tool := range c.Tools {
		toolByte, _ := json.Marshal(tool)
		fmt.Println(string(toolByte))
	}

	data, err := c.ExecTools(ctx, "list_commits", map[string]interface{}{
		"owner": "yincongcyincong",
		"repo":  "telegram-deepseek-bot",
	})
	if err != nil {
		log.Fatal("ExecTools failed:", err)
	}

	fmt.Println(data)

}
