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
	errs := clients.RegisterMCPClient(context.Background(), mcpParams)
	if len(errs) > 0 {
		log.Fatal("init amap fail", errs)
	}

	mc, err := clients.GetMCPClient(amap.NpxAmapMapsMcpServer)
	if err != nil {
		log.Fatal("get mcp client fail", err)
	}

	openaiTools := utils.TransToolsToChatGPTFunctionCall(mc.Tools)

	// todo modify deepseek token
	openAIkey := "xxx"

	//proxy, err := url.Parse("http://127.0.0.1:7890")
	//if err != nil {
	//	log.Fatal("parse deepseek proxy error", "err", err)
	//}

	config := openai.DefaultConfig(openAIkey)
	//config.HTTPClient = &http.Client{
	//	Transport: &http.Transport{
	//		Proxy: http.ProxyURL(proxy),
	//	},
	//}
	config.BaseURL = "https://api.chatanywhere.org"
	client := openai.NewClientWithConfig(config)

	ctx := context.Background()

	userMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: "My IP address is 220.181.3.151. May I know which city I am in",
	}

	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo0125,
		Messages: []openai.ChatCompletionMessage{userMessage},
		Tools:    openaiTools,
	})
	if err != nil {
		log.Fatalf("ChatCompletion error: %v\n", err)
	}

	msg := resp.Choices[0].Message

	msgStr, _ := json.Marshal(msg)

	fmt.Println(string(msgStr))
	// 如果需要调用函数
	if len(msg.ToolCalls) > 0 {
		args := make(map[string]interface{})
		if err := json.Unmarshal([]byte(msg.ToolCalls[0].Function.Arguments), &args); err != nil {
			log.Fatalf("Failed to parse function args: %v\n", err)
		}

		// 假设我们调用本地函数 getWeather
		result, err := mc.ExecTools(ctx, msg.ToolCalls[0].Function.Name, args)
		if err != nil {
			log.Fatalf("Exec fail: %v\n", err)
		}

		fmt.Println("call tools res:", result)

		// 把函数调用的结果喂回去
		resp2, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo0125,
			Messages: []openai.ChatCompletionMessage{
				userMessage,
				{
					Role:      openai.ChatMessageRoleAssistant,
					Content:   msg.Content,
					ToolCalls: msg.ToolCalls,
				},
				{
					Role:       openai.ChatMessageRoleTool,
					Content:    result,
					ToolCallID: msg.ToolCalls[0].ID,
				},
			},
		})
		if err != nil {
			log.Fatalf("Final ChatCompletion error: %v\n", err)
		}

		fmt.Println("💬 ChatGPT Response:", resp2.Choices[0].Message.Content)
	}
}
