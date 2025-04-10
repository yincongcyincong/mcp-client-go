package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/cohesion-org/deepseek-go"
	"github.com/cohesion-org/deepseek-go/constants"
	"github.com/yincongcyincong/mcp-client-go/clients"
	"github.com/yincongcyincong/mcp-client-go/clients/amap"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
	"github.com/yincongcyincong/mcp-client-go/utils"
	"io"
	"log"
)

func main() {
	mcpParams := make([]*param.MCPClientConf, 0)

	// todo add modify api key
	mcpParams = append(mcpParams,
		amap.InitAmapMCPClient("xxxx", "", nil, nil, nil))
	err := clients.RegisterMCPClient(context.Background(), mcpParams)
	if err != nil {
		log.Fatal("init amap fail", err)
	}

	mc, err := clients.GetMCPClient(amap.NpxAmapMapsMcpServer)
	if err != nil {
		log.Fatal("get mcp client fail", err)
	}

	deepseekTools := utils.TransToolsToDPFunctionCall(mc.Tools)

	// todo modify deepseek token
	client, err := deepseek.NewClientWithOptions("sk-xxxx")
	if err != nil {
		log.Fatal("Error creating deepseek client", "err", err)
	}

	request := &deepseek.StreamChatCompletionRequest{
		Model:  deepseek.DeepSeekChat,
		Stream: true,
		StreamOptions: deepseek.StreamOptions{
			IncludeUsage: true,
		},
		Tools: deepseekTools,
	}

	messages := []deepseek.ChatCompletionMessage{
		{
			Role:    constants.ChatMessageRoleUser,
			Content: "My IP address is 111.108.111.135. May I know which city I am in",
		},
	}

	request.Messages = messages

	ctx := context.Background()

	stream, err := client.CreateChatCompletionStream(ctx, request)
	if err != nil {
		log.Fatal("ChatCompletionStream error", "err", err)
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			log.Println("Stream finished")
			break
		}
		if err != nil {
			log.Fatal("Stream error", "err", err)
			break
		}
		for _, choice := range response.Choices {
			if len(choice.Delta.ToolCalls) > 0 {
				fmt.Println("tools", choice.Delta.ToolCalls)
				continue
			}

			// exceed max telegram one message length
			fmt.Println(choice.Delta.Content)
		}

	}
}
