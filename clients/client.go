package clients

import (
	"context"
	"fmt"
	"sync"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
	"github.com/yincongcyincong/mcp-client-go/utils"
)

var (
	mcpClients sync.Map
)

type MCPClient struct {
	Conf        *param.MCPClientConf
	StdioClient *client.StdioMCPClient
	SSEClient   *client.SSEMCPClient
	InitReq     *mcp.InitializeResult
	Tools       []mcp.Tool
}

const (
	SSEType   = "sse"
	StdioType = "stdio"
)

func RegisterMCPClient(ctx context.Context, params []*param.MCPClientConf) []error {
	errs := make([]error, 0)
	for _, clientParam := range params {
		if clientParam.SSEClientConf == nil || clientParam.ClientType != SSEType {
			err := createStdioMCPClient(ctx, clientParam)
			if err != nil {
				errs = append(errs, err)
			}
		} else {
			err := createSSEMCPClient(ctx, clientParam)
			if err != nil {
				errs = append(errs, err)
			}
		}

	}

	return errs
}

func createSSEMCPClient(ctx context.Context, clientParam *param.MCPClientConf) error {
	c, err := client.NewSSEMCPClient(
		clientParam.SSEClientConf.BaseUrl,
		clientParam.SSEClientConf.Options...,
	)
	if err != nil {
		return err
	}

	initResult, err := c.Initialize(ctx, clientParam.StdioClientConf.InitReq)
	if err != nil {
		return err
	}

	mc := &MCPClient{
		Conf:      clientParam,
		SSEClient: c,
		InitReq:   initResult,
	}
	tools, err := mc.GetAllTools(ctx, "")
	if err != nil {
		return err
	}
	mc.Tools = tools

	mcpClients.Store(clientParam.Name, mc)
	return nil
}

func createStdioMCPClient(ctx context.Context, clientParam *param.MCPClientConf) error {
	c, err := client.NewStdioMCPClient(
		clientParam.StdioClientConf.Command,
		clientParam.StdioClientConf.Env,
		clientParam.StdioClientConf.Args...,
	)
	if err != nil {
		return err
	}

	initResult, err := c.Initialize(ctx, clientParam.StdioClientConf.InitReq)
	if err != nil {
		return err
	}

	mc := &MCPClient{
		Conf:        clientParam,
		StdioClient: c,
		InitReq:     initResult,
	}
	tools, err := mc.GetAllTools(ctx, "")
	if err != nil {
		return err
	}
	mc.Tools = tools

	mcpClients.Store(clientParam.Name, mc)
	return nil
}

func InitStdioMCPClient(name, command string, env, args []string, initReq mcp.InitializeRequest,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	amapMCPClient := &param.MCPClientConf{
		Name:       name,
		ClientType: StdioType,
		StdioClientConf: &param.StdioClientConfig{
			Command: command,
			Env:     env,
			Args:    args,
			InitReq: initReq,
		},
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	return amapMCPClient
}

func InitSSEMCPClient(name, baseUrl string, options []client.ClientOption,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	amapMCPClient := &param.MCPClientConf{
		Name:       name,
		ClientType: SSEType,
		SSEClientConf: &param.SSEClientConfig{
			Options: options,
			BaseUrl: baseUrl,
		},
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	return amapMCPClient
}

func GetMCPClient(name string) (*MCPClient, error) {
	v, ok := mcpClients.Load(name)
	if !ok {
		return nil, fmt.Errorf("client %s not found", name)
	}
	return v.(*MCPClient), nil
}

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

func (m *MCPClient) GetAllTools(ctx context.Context, cursor mcp.Cursor) ([]mcp.Tool, error) {
	toolsRequest := mcp.ListToolsRequest{}
	toolsRequest.Params.Cursor = cursor

	var tools *mcp.ListToolsResult
	var err error
	if m.Conf.SSEClientConf == nil || m.Conf.Name != SSEType {
		tools, err = m.StdioClient.ListTools(ctx, toolsRequest)
	} else {
		tools, err = m.SSEClient.ListTools(ctx, toolsRequest)
	}

	if err != nil {
		return nil, err
	}
	return tools.Tools, nil
}

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

	if m.Conf.SSEClientConf == nil || m.Conf.Name != SSEType {
		result, err = m.StdioClient.CallTool(ctx, reqTool)
	} else {
		result, err = m.SSEClient.CallTool(ctx, reqTool)
	}
	if err != nil {
		return "", err
	}

	if _, ok := m.Conf.ToolsAfterFunc[name]; ok {
		return m.Conf.ToolsAfterFunc[name](result)
	}

	return utils.ReturnString(result), nil
}
