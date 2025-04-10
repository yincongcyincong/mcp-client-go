package main

import (
	"context"
	"encoding/json"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"github.com/yincongcyincong/mcp-client-go/clients"
	"github.com/yincongcyincong/mcp-client-go/clients/amap"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
	"github.com/yincongcyincong/mcp-client-go/utils"
	"log"
)

func main() {
	mcpParams := make([]*param.MCPClientConf, 0)

	// todo add modify api key
	amapApiKey := "xxx"
	mcpParams = append(mcpParams,
		amap.InitAmapMCPClient(amapApiKey, "", nil, nil, nil))
	err := clients.RegisterMCPClient(context.Background(), mcpParams)
	if err != nil {
		log.Fatal("init amap fail", err)
	}

	mc, err := clients.GetMCPClient(amap.NpxAmapMapsMcpServer)
	if err != nil {
		log.Fatal("get mcp client fail", err)
	}

	openaiTools := utils.TransToolsToChatGPTFunctionCall(mc.Tools)

	// todo modify deepseek token
	openAIkey := "xxx"
	client := openai.NewClient(openAIkey)

	ctx := context.Background()

	// 用户输入
	userMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: "What's the weather like in Tokyo today?",
	}

	// 定义函数 schema

	// 第一次请求：ChatGPT 判断是否调用函数
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    openai.GPT4,
		Messages: []openai.ChatCompletionMessage{userMessage},
		Tools:    openaiTools,
	})
	if err != nil {
		log.Fatalf("ChatCompletion error: %v\n", err)
	}

	msg := resp.Choices[0].Message

	// 如果需要调用函数
	if msg.FunctionCall != nil {
		args := make(map[string]interface{})
		if err := json.Unmarshal([]byte(msg.FunctionCall.Arguments), &args); err != nil {
			log.Fatalf("Failed to parse function args: %v\n", err)
		}

		// 假设我们调用本地函数 getWeather
		result, err := mc.ExecTools(ctx, msg.FunctionCall.Name, args)
		if err != nil {
			log.Fatalf("Exec fail: %v\n", err)
		}

		// 把函数调用的结果喂回去
		resp2, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				userMessage,
				msg,
				{
					Role:    openai.ChatMessageRoleFunction,
					Name:    msg.FunctionCall.Name,
					Content: result,
				},
			},
		})
		if err != nil {
			log.Fatalf("Final ChatCompletion error: %v\n", err)
		}

		fmt.Println("💬 ChatGPT Response:", resp2.Choices[0].Message.Content)
	}
}
