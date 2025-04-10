package clients

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
	"github.com/yincongcyincong/mcp-client-go/utils"
	"sync"
)

var (
	mcpClients sync.Map
)

type MCPClient struct {
	Conf    *param.MCPClientConf
	Client  *client.StdioMCPClient
	InitReq *mcp.InitializeResult
	Tools   []mcp.Tool
}

func RegisterMCPClient(ctx context.Context, params []*param.MCPClientConf) error {
	for _, clientParam := range params {
		c, err := client.NewStdioMCPClient(
			clientParam.Command,
			clientParam.Env,
			clientParam.Args...,
		)
		if err != nil {
			return err
		}

		initResult, err := c.Initialize(ctx, clientParam.InitReq)
		if err != nil {
			return err
		}

		mc := &MCPClient{
			Conf:    clientParam,
			Client:  c,
			InitReq: initResult,
		}
		tools, err := mc.GetAllTools(ctx, "")
		if err != nil {
			return err
		}
		mc.Tools = tools

		mcpClients.Store(clientParam.Name, mc)

	}

	return nil
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
	tools, err := m.Client.ListTools(ctx, toolsRequest)
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

	result, err := m.Client.CallTool(ctx, reqTool)
	if err != nil {
		return "", err
	}

	if _, ok := m.Conf.ToolsAfterFunc[name]; ok {
		return m.Conf.ToolsAfterFunc[name](result)
	}

	return utils.ReturnString(result), nil
}
