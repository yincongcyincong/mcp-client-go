package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"
	"sync"
	"time"
	
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
	"github.com/yincongcyincong/mcp-client-go/utils"
)

var (
	mcpClients sync.Map
)

type MCPClient struct {
	Conf    *param.MCPClientConf
	Client  *client.Client
	InitReq *mcp.InitializeResult
	Tools   []mcp.Tool
	
	ctx    context.Context
	cancel context.CancelFunc
}

// InitByConfFile Initialize multiple MCP clients from configuration files
func InitByConfFile(configFilePath string) ([]*param.MCPClientConf, error) {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	
	config := new(param.McpClientGoConfig)
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	
	mcs := make([]*param.MCPClientConf, 0)
	for mcpName, mcpConf := range config.McpServers {
		mcpType := getMcpType(mcpConf)
		switch mcpType {
		case param.StdioType:
			mcs = append(mcs, InitStdioMCPClient(mcpName, mcpConf.Command,
				utils.ChangeEnvMapToSlice(mcpConf.Env), mcpConf.Args,
				param.WithDescription(mcpConf.Description)))
		case param.HTTPConfigType:
			httpType, err := utils.CheckSSEOrHTTP(mcpConf.Url)
			if err != nil {
				log.Println("CheckSSEOrHTTP fail, err:", err)
				continue
			}
			
			if httpType == param.SSEType {
				mcs = append(mcs, InitSSEMCPClient(mcpName, mcpConf.Url,
					param.WithSSEOptions(transport.WithHeaders(mcpConf.Headers)),
					param.WithDescription(mcpConf.Description)))
			} else {
				mcs = append(mcs, InitHttpMCPClient(mcpName, mcpConf.Url,
					param.WithHttpOptions(transport.WithHTTPHeaders(mcpConf.Headers)),
					param.WithDescription(mcpConf.Description),
					param.WithHttpOauth(mcpConf.OAuth)))
			}
		
		default:
			log.Println("mcp type not exist, mcpType:", mcpType)
		}
	}
	
	return mcs, nil
	
}

// RegisterMCPClient register multiple MCP clients
func RegisterMCPClient(ctx context.Context, params []*param.MCPClientConf) map[string]error {
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
				err := createSSEMCPClient(ctx, cp)
				if err != nil {
					mutex.Lock()
					errs[clientParam.Name] = err
					mutex.Unlock()
				}
			} else if cp.HTTPStreamerConf != nil && cp.ClientType == param.HTTPStreamer {
				err := createHTTPStreamCPClient(ctx, cp)
				if err != nil {
					mutex.Lock()
					errs[clientParam.Name] = err
					mutex.Unlock()
				}
			} else {
				err := createStdioMCPClient(ctx, cp)
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

// createStdioMCPClient create stdio MCP client
func createSSEMCPClient(ctx context.Context, clientParam *param.MCPClientConf) error {
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
	
	return initMCPClient(ctx, c, clientParam, initResult)
}

// createSSEMCPClient create SSE MCP client
func createHTTPStreamCPClient(ctx context.Context, clientParam *param.MCPClientConf) error {
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
	
	return initMCPClient(ctx, c, clientParam, initResult)
}

// createStdioMCPClient create stdio MCP client
func createStdioMCPClient(ctx context.Context, clientParam *param.MCPClientConf) error {
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
	
	return initMCPClient(ctx, c, clientParam, initResult)
}

func initMCPClient(ctx context.Context, c *client.Client,
	clientParam *param.MCPClientConf, initResult *mcp.InitializeResult) error {
	
	mcContext, cancel := context.WithCancel(context.Background())
	mc := &MCPClient{
		Conf:    clientParam,
		Client:  c,
		InitReq: initResult,
		
		ctx:    mcContext,
		cancel: cancel,
	}
	tools, err := mc.GetAllTools(ctx, "")
	if err != nil {
		return err
	}
	mc.Tools = tools
	
	if clientInter, ok := mcpClients.Load(clientParam.Name); ok {
		clientInter.(*MCPClient).Close()
	}
	
	mcpClients.Store(clientParam.Name, mc)
	
	go mc.handlePing()
	return nil
}

// InitStdioMCPClient Initialize the stdio MCP client
func InitStdioMCPClient(name, command string, env, args []string, options ...param.Option) *param.MCPClientConf {
	
	mcpClient := &param.MCPClientConf{
		Name:       name,
		ClientType: param.StdioType,
		StdioClientConf: &param.StdioClientConfig{
			Command: command,
			Env:     env,
			Args:    args,
		},
	}
	
	for _, o := range options {
		o(mcpClient)
	}
	
	if mcpClient.StdioClientConf.InitReq.Params.ProtocolVersion == "" {
		mcpClient.StdioClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}
	
	if mcpClient.StdioClientConf.InitReq.Params.ClientInfo.Name == "" {
		mcpClient.StdioClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/unknown",
			Version: "0.1.0",
		}
	}
	
	return mcpClient
}

// InitSSEMCPClient Initialize the SSE MCP client
func InitSSEMCPClient(name, baseUrl string, options ...param.Option) *param.MCPClientConf {
	
	mcpClient := &param.MCPClientConf{
		Name:       name,
		ClientType: param.SSEType,
		SSEClientConf: &param.SSEClientConfig{
			BaseUrl: baseUrl,
		},
	}
	
	for _, o := range options {
		o(mcpClient)
	}
	
	if mcpClient.SSEClientConf.InitReq.Params.ProtocolVersion == "" {
		mcpClient.SSEClientConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}
	
	if mcpClient.SSEClientConf.InitReq.Params.ClientInfo.Name == "" {
		mcpClient.SSEClientConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/unknown",
			Version: "0.1.0",
		}
	}
	
	return mcpClient
}

// InitHttpMCPClient Initialize the HTTP MCP client
func InitHttpMCPClient(name, baseUrl string, options ...param.Option) *param.MCPClientConf {
	mcpClient := &param.MCPClientConf{
		Name:       name,
		ClientType: param.HTTPStreamer,
		HTTPStreamerConf: &param.HTTPStreamerConfig{
			BaseURL: baseUrl,
		},
	}
	
	for _, o := range options {
		o(mcpClient)
	}
	
	if mcpClient.HTTPStreamerConf.InitReq.Params.ProtocolVersion == "" {
		mcpClient.HTTPStreamerConf.InitReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	}
	
	if mcpClient.HTTPStreamerConf.InitReq.Params.ClientInfo.Name == "" {
		mcpClient.HTTPStreamerConf.InitReq.Params.ClientInfo = mcp.Implementation{
			Name:    "mcp-server/unknown",
			Version: "0.1.0",
		}
	}
	
	return mcpClient
}

// GetMCPClient Get MCP client by name
func GetMCPClient(name string) (*MCPClient, error) {
	v, ok := mcpClients.Load(name)
	if !ok {
		return nil, fmt.Errorf("client %s not found", name)
	}
	return v.(*MCPClient), nil
}

// GetMCPClientByToolName Get MCP client by tool name
func GetMCPClientByToolName(toolName string) (*MCPClient, error) {
	var res *MCPClient
	mcpClients.Range(func(key, value interface{}) bool {
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

// GetAllTools Get all tools from MCP client
func (m *MCPClient) GetAllTools(ctx context.Context, cursor mcp.Cursor) ([]mcp.Tool, error) {
	toolsRequest := mcp.ListToolsRequest{}
	toolsRequest.Params.Cursor = cursor
	
	var tools *mcp.ListToolsResult
	var err error
	tools, err = m.Client.ListTools(ctx, toolsRequest)
	
	if err != nil {
		return nil, err
	}
	return tools.Tools, nil
}

// ExecTools Execute a tool
func (m *MCPClient) ExecTools(ctx context.Context, name string, params map[string]interface{}) (string, error) {
	var useTool *mcp.Tool
	
	for _, tool := range m.Tools {
		if tool.Name == name {
			useTool = &tool
			break
		}
	}
	
	if useTool == nil {
		return "", fmt.Errorf("tool %s not found", name)
	}
	
	for _, reqParam := range useTool.InputSchema.Required {
		if _, ok := params[reqParam]; !ok {
			return "", fmt.Errorf("required parameter %s not found", reqParam)
		}
	}
	
	reqTool := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	reqTool.Params.Name = name
	reqTool.Params.Arguments = params
	
	if _, ok := m.Conf.ToolsBeforeFunc[name]; ok {
		err := m.Conf.ToolsBeforeFunc[name](&reqTool)
		if err != nil {
			return "", err
		}
	}
	var result *mcp.CallToolResult
	var err error
	result, err = m.Client.CallTool(ctx, reqTool)
	if err != nil {
		return "", err
	}
	
	if _, ok := m.Conf.ToolsAfterFunc[name]; ok {
		return m.Conf.ToolsAfterFunc[name](result)
	}
	
	return utils.ReturnString(result), nil
}

// handlePing handle ping
func (m *MCPClient) handlePing() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic recovered: %v\n%s", r, debug.Stack())
		}
	}()
	
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			m.restartMCPServer()
		case <-m.ctx.Done():
			log.Println("mcp client stoped", m.Conf.Name)
			return
		}
	}
	
}

// restartMCPServer restart mcp server
func (m *MCPClient) restartMCPServer() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := m.Client.Ping(ctx)
	
	if err != nil {
		log.Println("mcp ping fail:", err)
		
		if m.Conf.SSEClientConf != nil && m.Conf.ClientType == param.SSEType {
			err = createSSEMCPClient(ctx, m.Conf)
		} else if m.Conf.HTTPStreamerConf != nil && m.Conf.ClientType == param.HTTPStreamer {
			err = createHTTPStreamCPClient(ctx, m.Conf)
		} else {
			err = createStdioMCPClient(ctx, m.Conf)
		}
		
		if err != nil {
			log.Println("create new mcp client fail:", err)
			return false
		}
		
		return true
	}
	
	return false
}

func (m *MCPClient) Close() {
	m.cancel()
	m.Client.Close()
}

// getMcpType get mcp type
func getMcpType(conf *param.MCPConfig) string {
	if conf.Type != "" {
		return conf.Type
	}
	
	if conf.Command != "" {
		return param.StdioConfigType
	}
	
	return param.HTTPConfigType
}
