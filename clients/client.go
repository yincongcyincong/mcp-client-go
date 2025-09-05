package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"time"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
	"github.com/yincongcyincong/mcp-client-go/utils"
)

var (
	defaultClients Clients
)

type MCPClient struct {
	Conf    *param.MCPClientConf
	Client  *client.Client
	InitReq *mcp.InitializeResult
	Tools   []mcp.Tool

	ctx    context.Context
	cancel context.CancelFunc
	parent *Clients
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
		mcpClientConf := GetOneMCPClient(mcpName, mcpConf)
		if mcpClientConf != nil {
			mcs = append(mcs, mcpClientConf)
		}
	}

	return mcs, nil

}

func GetOneMCPClient(mcpName string, mcpConf *param.MCPConfig) *param.MCPClientConf {
	if mcpConf.Disabled {
		return nil
	}

	mcpType := getMcpType(mcpConf)
	switch mcpType {
	case param.StdioType:
		return InitStdioMCPClient(mcpName, mcpConf.Command,
			utils.ChangeEnvMapToSlice(mcpConf.Env), mcpConf.Args,
			param.WithDescription(mcpConf.Description))
	case param.HTTPConfigType:
		httpType, err := utils.CheckSSEOrHTTP(mcpConf.Url)
		if err != nil {
			log.Println("CheckSSEOrHTTP fail, err:", err)
			return nil
		}

		if httpType == param.SSEType {
			return InitSSEMCPClient(mcpName, mcpConf.Url,
				param.WithSSEOptions(transport.WithHeaders(mcpConf.Headers)),
				param.WithDescription(mcpConf.Description))
		} else {
			return InitHttpMCPClient(mcpName, mcpConf.Url,
				param.WithHttpOptions(transport.WithHTTPHeaders(mcpConf.Headers)),
				param.WithDescription(mcpConf.Description),
				param.WithHttpOauth(mcpConf.OAuth))
		}

	default:
		log.Println("mcp type not exist, mcpType:", mcpType)
	}

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
			err = m.parent.createSSEMCPClient(ctx, m.Conf)
		} else if m.Conf.HTTPStreamerConf != nil && m.Conf.ClientType == param.HTTPStreamer {
			err = m.parent.createHTTPStreamCPClient(ctx, m.Conf)
		} else {
			err = m.parent.createStdioMCPClient(ctx, m.Conf)
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
