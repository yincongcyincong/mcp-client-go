package utils

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
	
	"github.com/cohesion-org/deepseek-go"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/revrost/go-openrouter"
	"github.com/sashabaranov/go-openai"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
	"google.golang.org/genai"
)

func ReturnString(result *mcp.CallToolResult) string {
	if result == nil {
		return ""
	}
	
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
	openRouterTools := make([]deepseek.Tool, 0)
	for _, tool := range tools {
		openRouterTool := deepseek.Tool{
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
		openRouterTools = append(openRouterTools, openRouterTool)
	}
	
	return openRouterTools
}

func TransToolsToChatGPTFunctionCall(tools []mcp.Tool) []openai.Tool {
	openaiTools := make([]openai.Tool, 0)
	for _, tool := range tools {
		openaiTool := openai.Tool{
			Type: "function",
			Function: &openai.FunctionDefinition{
				Name:        tool.Name,
				Description: tool.Description,
				Parameters: &deepseek.FunctionParameters{
					Type:       "object",
					Properties: tool.InputSchema.Properties,
					Required:   tool.InputSchema.Required,
				},
			},
		}
		openaiTools = append(openaiTools, openaiTool)
	}
	
	return openaiTools
}

func TransToolsToGeminiFunctionCall(tools []mcp.Tool) []*genai.Tool {
	geminiTools := []*genai.Tool{}
	
	for _, tool := range tools {
		prop := make(map[string]*genai.Schema)
		propByte, err := json.Marshal(tool.InputSchema.Properties)
		if err != nil {
			continue
		}
		err = json.Unmarshal(propByte, &prop)
		if err != nil {
			continue
		}
		geminiTool := &genai.Tool{
			FunctionDeclarations: []*genai.FunctionDeclaration{
				{
					Name:        tool.Name,
					Description: tool.Description,
					Parameters: &genai.Schema{
						Type:       genai.TypeObject,
						Properties: prop,
						Required:   tool.InputSchema.Required,
					},
				},
			},
		}
		
		geminiTools = append(geminiTools, geminiTool)
	}
	
	return geminiTools
}

func TransToolsToOpenRouterFunctionCall(tools []mcp.Tool) []openrouter.Tool {
	deepseekTools := make([]openrouter.Tool, 0)
	for _, tool := range tools {
		deepseekTool := openrouter.Tool{
			Type: "function",
			Function: &openrouter.FunctionDefinition{
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

func ChangeEnvMapToSlice(env map[string]string) []string {
	res := make([]string, 0)
	for k, v := range env {
		res = append(res, k+"="+v)
	}
	return res
}

func CheckSSEOrHTTP(url string) (string, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	
	contentType := resp.Header.Get("Content-Type")
	if contentType == "text/event-stream" {
		return param.SSEType, nil
	}
	
	return param.HTTPStreamer, nil
}
