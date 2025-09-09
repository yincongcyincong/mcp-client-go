package clients

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
	"io"
	"log"
	"runtime/debug"
	"sync"
)

var (
	_ _clients = (*Clients)(nil)
)

type _clients interface {
	ClearAllMCPClient()
	GetMCPClient(name string) (*MCPClient, error)
	GetMCPClientByToolName(toolName string) (*MCPClient, error)
	initMCPClient(ctx context.Context, c *client.Client, clientParam *param.MCPClientConf, initResult *mcp.InitializeResult) error
	RemoveMCPClient(name string) error
	RegisterMCPClient(ctx context.Context, params []*param.MCPClientConf) map[string]error
}

type Clients struct {
	mcpClients sync.Map
}

func (cs *Clients) RegisterMCPClient(ctx context.Context, params []*param.MCPClientConf) map[string]error {
	errs := make(map[string]error)

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	for _, clientParam := range params {
		wg.Add(1)
		go func(cp *param.MCPClientConf) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("panic recovered: %v\n%s", r, debug.Stack())
				}
				wg.Done()
			}()
			if cp.SSEClientConf != nil && cp.ClientType == param.SSEType {
				err := cs.createSSEMCPClient(ctx, cp)
				if err != nil {
					mutex.Lock()
					errs[clientParam.Name] = err
					mutex.Unlock()
				}
			} else if cp.HTTPStreamerConf != nil && cp.ClientType == param.HTTPStreamer {
				err := cs.createHTTPStreamCPClient(ctx, cp)
				if err != nil {
					mutex.Lock()
					errs[clientParam.Name] = err
					mutex.Unlock()
				}
			} else {
				err := cs.createStdioMCPClient(ctx, cp)
				if err != nil {
					mutex.Lock()
					errs[clientParam.Name] = err
					mutex.Unlock()
				}

			}
		}(clientParam)
	}
	wg.Wait()

	return errs
}

func (cs *Clients) ClearAllMCPClient() {
	cs.mcpClients.Range(func(key, value interface{}) bool {
		value.(*MCPClient).Close()
		return true
	})
	cs.mcpClients.Clear()
}

func (cs *Clients) GetMCPClient(name string) (*MCPClient, error) {
	v, ok := cs.mcpClients.Load(name)
	if !ok {
		return nil, fmt.Errorf("client %s not found", name)
	}
	return v.(*MCPClient), nil
}

func (cs *Clients) GetMCPClientByToolName(toolName string) (*MCPClient, error) {
	var res *MCPClient
	cs.mcpClients.Range(func(key, value interface{}) bool {
		mc := value.(*MCPClient)
		for _, tool := range mc.Tools {
			if tool.Name == toolName {
				res = mc
				return true
			}
		}
		return true
	})

	if res == nil {
		return nil, fmt.Errorf("tool %s not found", toolName)
	}

	return res, nil
}

func (cs *Clients) initMCPClient(ctx context.Context, c *client.Client, clientParam *param.MCPClientConf, initResult *mcp.InitializeResult) error {

	mcContext, cancel := context.WithCancel(context.Background())
	mc := &MCPClient{
		Conf:    clientParam,
		Client:  c,
		InitReq: initResult,

		ctx:    mcContext,
		cancel: cancel,

		parent: cs,
	}
	tools, err := mc.GetAllTools(ctx, "")
	if err != nil {
		return err
	}
	mc.Tools = tools

	if clientInter, ok := cs.mcpClients.Load(clientParam.Name); ok {
		clientInter.(*MCPClient).Close()
	}

	cs.mcpClients.Store(clientParam.Name, mc)

	go mc.handlePing()
	return nil
}

func (cs *Clients) RemoveMCPClient(name string) error {
	c, err := GetMCPClient(name)
	if err != nil {
		return err
	}
	c.Close()
	cs.mcpClients.Delete(name)
	return nil
}

// createStdioMCPClient create stdio MCP client
func (cs *Clients) createSSEMCPClient(ctx context.Context, clientParam *param.MCPClientConf) error {
	c, err := client.NewSSEMCPClient(
		clientParam.SSEClientConf.BaseUrl,
		clientParam.SSEClientConf.Options...,
	)
	if err != nil {
		return err
	}

	err = c.Start(context.Background())
	if err != nil {
		return err
	}

	initResult, err := c.Initialize(ctx, clientParam.SSEClientConf.InitReq)
	if err != nil {
		return err
	}

	return cs.initMCPClient(ctx, c, clientParam, initResult)
}

// createSSEMCPClient create SSE MCP client
func (cs *Clients) createHTTPStreamCPClient(ctx context.Context, clientParam *param.MCPClientConf) error {
	var c *client.Client
	var err error
	if clientParam.HTTPStreamerConf.Oauth != nil {
		c, err = client.NewOAuthStreamableHttpClient(
			clientParam.HTTPStreamerConf.BaseURL,
			*clientParam.HTTPStreamerConf.Oauth,
			clientParam.HTTPStreamerConf.Options...,
		)
	} else {
		c, err = client.NewStreamableHttpClient(
			clientParam.HTTPStreamerConf.BaseURL,
			clientParam.HTTPStreamerConf.Options...,
		)
	}

	if err != nil {
		return err
	}

	err = c.Start(context.Background())
	if err != nil {
		return err
	}

	initResult, err := c.Initialize(ctx, clientParam.HTTPStreamerConf.InitReq)
	if err != nil {
		return err
	}

	return cs.initMCPClient(ctx, c, clientParam, initResult)
}

// createStdioMCPClient create stdio MCP client
func (cs *Clients) createStdioMCPClient(ctx context.Context, clientParam *param.MCPClientConf) error {
	c, err := client.NewStdioMCPClient(
		clientParam.StdioClientConf.Command,
		clientParam.StdioClientConf.Env,
		clientParam.StdioClientConf.Args...,
	)
	if err != nil {
		return err
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic recovered: %v\n%s", r, debug.Stack())
			}
		}()
		stderr, ok := client.GetStderr(c)
		if !ok {
			return
		}

		buf := make([]byte, 1024) // 合理大小缓存区
		for {
			n, err := stderr.Read(buf)
			if err != nil {
				if err == io.EOF {
					log.Println("stderr closed:", clientParam.Name)
					break
				}
				log.Println("read stderr fail:", clientParam.Name, err)
				break
			}
			if n > 0 {
				log.Println("stderr:", clientParam.Name, string(buf[:n]))
			}
		}
	}()

	initResult, err := c.Initialize(ctx, clientParam.StdioClientConf.InitReq)
	if err != nil {
		return err
	}

	return cs.initMCPClient(ctx, c, clientParam, initResult)
}
