package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yincongcyincong/mcp-client-go/clients"
	"github.com/yincongcyincong/mcp-client-go/clients/keenable"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// Keep this example's clients separate from any other registered services.
	cs := &clients.Clients{}
	defer cs.ClearAllMCPClient()
	// KEENABLE_API_KEY is optional; without it the keyless public tier is used.
	conf := keenable.InitKeenableMCPClient(&keenable.KeenableParam{ApiKey: os.Getenv("KEENABLE_API_KEY")})
	if errs := cs.RegisterMCPClient(ctx, []*param.MCPClientConf{conf}); len(errs) > 0 {
		return fmt.Errorf("register Keenable: %v", errs)
	}
	c, err := cs.GetMCPClient(keenable.KeenableMCPServer)
	if err != nil {
		return err
	}

	search, err := c.ExecTools(ctx, "search_web_pages", map[string]interface{}{
		"query":       "official Go blog post on context cancellation",
		"max_results": 3,
	})
	if err != nil {
		return fmt.Errorf("search_web_pages: %w", err)
	}
	fmt.Println(search)

	page, err := c.ExecTools(ctx, "fetch_page_content", map[string]interface{}{
		"url":       "https://go.dev/blog/context",
		"max_chars": 2000,
	})
	if err != nil {
		return fmt.Errorf("fetch_page_content: %w", err)
	}
	fmt.Println(page)
	return nil
}
