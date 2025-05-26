package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cohesion-org/deepseek-go"
	"github.com/cohesion-org/deepseek-go/constants"
	"github.com/yincongcyincong/mcp-client-go/clients"
	"github.com/yincongcyincong/mcp-client-go/clients/amap"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
	"github.com/yincongcyincong/mcp-client-go/utils"
	"log"
)

func main() {
	mcpParams := make([]*param.MCPClientConf, 0)

	// todo add modify api key
	amapApiKey := "xxxx"
	mcpParams = append(mcpParams,
		amap.InitAmapMCPClient(&amap.AmapParam{
			AmapApiKey: amapApiKey,
		}))
	errs := clients.RegisterMCPClient(context.Background(), mcpParams)
	if len(errs) > 0 {
		log.Fatal("init amap fail", errs)
	}

	mc, err := clients.GetMCPClient(amap.NpxAmapMapsMcpServer)
	if err != nil {
		log.Fatal("get mcp client fail", err)
	}

	deepseekTools := utils.TransToolsToDPFunctionCall(mc.Tools)

	// todo modify deepseek token
	deepseekKey := "sk-xxxx"
	client, err := deepseek.NewClientWithOptions(deepseekKey)
	if err != nil {
		log.Fatal("Error creating deepseek client", "err", err)
	}

	request := &deepseek.ChatCompletionRequest{
		Model: deepseek.DeepSeekChat,
		Tools: deepseekTools,
	}

	messages := []deepseek.ChatCompletionMessage{
		{
			Role:    constants.ChatMessageRoleUser,
			Content: "My IP address is 220.181.3.151. May I know which city I am in",
		},
	}

	request.Messages = messages

	ctx := context.Background()

	response, err := client.CreateChatCompletion(ctx, request)
	if err != nil {
		log.Fatal("ChatCompletionStream error", "err", err)
	}

	// will be empty
	fmt.Println("response:", response.Choices[0].Message.Content)

	// one tool call request
	fmt.Println("tool calls:", response.Choices[0].Message.ToolCalls)

	msg := response.Choices[0].Message
	toolCalls := msg.ToolCalls

	p := make(map[string]interface{})
	err = json.Unmarshal([]byte(msg.ToolCalls[0].Function.Arguments), &p)
	if err != nil {
		log.Fatal("unmarshal fail", "err", err)
	}
	toolRes, err := mc.ExecTools(ctx, msg.ToolCalls[0].Function.Name, p)
	if err != nil {
		log.Fatal("exec fail", "err", err)
	}
	fmt.Println("toolRes:", toolRes)

	question := deepseek.ChatCompletionMessage{
		Role:      deepseek.ChatMessageRoleAssistant,
		Content:   msg.Content,
		ToolCalls: toolCalls,
	}
	answer := deepseek.ChatCompletionMessage{
		Role:       deepseek.ChatMessageRoleTool,
		Content:    toolRes,
		ToolCallID: toolCalls[0].ID,
	}

	messages = append(messages, question, answer)
	toolReq := &deepseek.ChatCompletionRequest{
		Model:    request.Model,
		Messages: messages,
	}

	response, err = client.CreateChatCompletion(ctx, toolReq)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// will return the current time
	fmt.Println("response:", response.Choices[0].Message.Content)
}
