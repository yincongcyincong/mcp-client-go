package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/yincongcyincong/mcp-client-go/clients"
	"github.com/yincongcyincong/mcp-client-go/clients/parallel"
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
	conf := parallel.InitParallelMCPClient()
	if errs := cs.RegisterMCPClient(ctx, []*param.MCPClientConf{conf}); len(errs) > 0 {
		return fmt.Errorf("register Parallel: %v", errs)
	}
	c, err := cs.GetMCPClient(parallel.ParallelMCPServer)
	if err != nil {
		return err
	}

	// Reuse a conversation identifier for related search and fetch calls.
	var sessionBytes [16]byte
	if _, err := rand.Read(sessionBytes[:]); err != nil {
		return fmt.Errorf("create session identifier: %w", err)
	}
	sessionID := hex.EncodeToString(sessionBytes[:])
	search, err := c.ExecTools(ctx, "web_search", map[string]interface{}{
		"objective":      "Find the official Go documentation for context cancellation.",
		"search_queries": []string{"Go context cancellation documentation"},
		"session_id":     sessionID,
	})
	if err != nil {
		return fmt.Errorf("web_search: %w", err)
	}
	fmt.Println(search)

	page, err := c.ExecTools(ctx, "web_fetch", map[string]interface{}{
		"urls":       []string{"https://go.dev/blog/context"},
		"objective":  "Explain how Go contexts support cancellation.",
		"session_id": sessionID,
	})
	if err != nil {
		return fmt.Errorf("web_fetch: %w", err)
	}
	fmt.Println(page)
	return nil
}
