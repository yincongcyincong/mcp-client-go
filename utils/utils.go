package utils

import (
	"encoding/json"
	"github.com/cohesion-org/deepseek-go"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"strings"
)

func ReturnString(result *mcp.CallToolResult) string {
	var res strings.Builder
	for _, content := range result.Content {
		if textContent, ok := content.(mcp.TextContent); ok {
			res.WriteString(textContent.Text)
		} else {
			jsonBytes, _ := json.MarshalIndent(content, "", "  ")
			res.Write(jsonBytes)
		}
	}

	return res.String()
}

func TransToolsToDPFunctionCall(tools []mcp.Tool) []deepseek.Tool {
	deepseekTools := make([]deepseek.Tool, 0)
	for _, tool := range tools {
		deepseekTool := deepseek.Tool{
			Type: "function",
			Function: deepseek.Function{
				Name:        tool.Name,
				Description: tool.Description,
				Parameters: &deepseek.FunctionParameters{
					Type:       "object",
					Properties: tool.InputSchema.Properties,
					Required:   tool.InputSchema.Required,
				},
			},
		}
		deepseekTools = append(deepseekTools, deepseekTool)
	}

	return deepseekTools
}

func TransToolsToVolFunctionCall(tools []mcp.Tool) []*model.Tool {
	deepseekTools := make([]*model.Tool, 0)
	for _, tool := range tools {
		deepseekTool := &model.Tool{
			Type: "function",
			Function: &model.FunctionDefinition{
				Name:        tool.Name,
				Description: tool.Description,
				Parameters: map[string]interface{}{
					"type":       "object",
					"properties": tool.InputSchema.Properties,
					"required":   tool.InputSchema.Required,
				},
			},
		}
		deepseekTools = append(deepseekTools, deepseekTool)
	}

	return deepseekTools
}
