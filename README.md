# 📦 mcp-client-go

`mcp-client-go` is a Golang client library for the **Model Context Protocol (MCP)**. It allows developers to register
and interact with various MCP-based services such as Amap (Gaode Maps) using a unified API.

---

## ✨ Features

- Easy integration with MCP-compatible services
- Modular support for service types (e.g., Amap, Github, GoogleMap)
- Unified registration and client management
- Simple and intuitive tool execution interface

---

## 📋 Supported Services

| MCP Server           | 	Description                                                                                                                                                                                                  | doc                                                                                     | demo                                                                                                                         |
|----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------|
| redis	               | A Model Context Protocol server that provides access to Redis databases.                                                                                                                                      | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/redis)              | [redis](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/redis/redis.go)                               |
| github	              | The GitHub MCP Server is a Model Context Protocol (MCP) server that provides seamless integration with GitHub APIs                                                                                            | [doc](https://github.com/github/github-mcp-server)                                      | [github](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/github/github.go)                            |
| aws	                 | An MCP server implementation for retrieving information from the AWS Knowledge Base using the Bedrock Agent Runtime.                                                                                          | [doc](https://github.com/awslabs/mcp)                                                   | [aws](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/aws/aws.go)                                     |
| sequential_thinking	 | An MCP server implementation that provides a tool for dynamic and reflective problem-solving through a structured thinking process.                                                                           | [doc](https://github.com/modelcontextprotocol/servers/tree/HEAD/src/sequentialthinking) | -                                                                                                                            |
| firecrawl	           | A Model Context Protocol (MCP) server implementation that integrates with Firecrawl for web scraping capabilities.                                                                                            | [doc](https://github.com/mendableai/firecrawl-mcp-server)                               | [firecrawl](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/firecrawl/firecrawl.go)                   |
| postgresql	          | A Model Context Protocol server that provides read-only access to PostgreSQL databases.                                                                                                                       | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/postgres)           | -                                                                                                                            |
| gitlab	              | MCP Server for the GitLab API, enabling project management, file operations, and more.                                                                                                                        | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/gitlab)             | -                                                                                                                            |
| slack	               | MCP Server for the Slack API, enabling Claude to interact with Slack workspaces.                                                                                                                              | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/slack)              | -                                                                                                                            |
| puppeteer	           | A Model Context Protocol server that provides browser automation capabilities using Puppeteer.                                                                                                                | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/puppeteer)          | -                                                                                                                            |
| everart	             | Image generation server for Claude Desktop using EverArt's API.                                                                                                                                               | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/everart)            | -                                                                                                                            |
| sentry	              | A Model Context Protocol server for retrieving and analyzing issues from Sentry.io                                                                                                                            | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/sentry)             | -                                                                                                                            |
| filesystem	          | Node.js server implementing Model Context Protocol (MCP) for filesystem operations.                                                                                                                           | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/filesystem)         | [filesystem](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/filesystem/filesystem.go)                |
| fetch	               | A Model Context Protocol server that provides web content fetching capabilities. This server enables LLMs to retrieve and process content from web pages, converting HTML to markdown for easier consumption. | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/fetch)              | -                                                                                                                            |
| googlemap	           | MCP Server for the Google Maps API.                                                                                                                                                                           | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/google-maps)        | [googlemap](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/googlemap/googlemap.go)                   |
| flomo	               | This is a TypeScript-based MCP server help you write notes to Flomo.                                                                                                                                          | [doc](https://github.com/chatmcp/mcp-server-flomo)                                      | -                                                                                                                            |
| chatsum	             | This MCP Server is used to summarize your chat messages.                                                                                                                                                      | [doc](https://github.com/chatmcp/mcp-server-chatsum)                                    | -                                                                                                                            |
| amap	                | This repository is a collection of reference implementations for the Model Context Protocol (MCP), as well as references to community built servers and additional resources.                                 | [doc](https://github.com/modelcontextprotocol/servers)                                  | [amap](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/amap/amap.go)                                  |
| baidumap	            | This MCP Server is used to baidumap                                                                                                                                                                           | [doc](https://github.com/baidu-maps/mcp)                                                | -                                                                                                                            |
| blender	             | BlenderMCP connects Blender to Claude AI through the Model Context Protocol (MCP)                                                                                                                             | [doc](https://github.com/ahujasid/blender-mcp)                                          | -                                                                                                                            |
| framelink	           | Give Cursor, Windsurf, Cline, and other AI-powered coding tools access to your Figma files with this Model Context Protocol server.                                                                           | [doc](https://github.com/GLips/Figma-Context-MCP)                                       | -                                                                                                                            |
| playwright	          | A Model Context Protocol (MCP) server that provides browser automation capabilities using Playwright.                                                                                                         | [doc](https://github.com/microsoft/playwright-mcp)                                      | [playwright](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/playwright/playwright.go)   sse          |
| tavily	              | The Model Context Protocol (MCP) is an open standard that enables AI systems to interact seamlessly with various data sources and tools, facilitating secure, two-way connections.                            | [doc](https://github.com/tavily-ai/tavily-mcp)                                          | [tavily](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/tavily/tavily.go)                            |
| time	                | A Model Context Protocol server that provides time and timezone conversion capabilities.                                                                                                                      | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/time)               | [time](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/time/time.go)   uvx                            |
| victoriametrics	     | A Model Context Protocol server that provide access to victoria metrics databases.                                                                                                                            | [doc](https://github.com/yincongcyincong/VictoriaMetrics-mcp-server)                    | [victoriametrics](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/victoriametrics/victoriametrics.go) |
| atlassian	           | Model Context Protocol (MCP) server for Atlassian products (Confluence and Jira)                                                                                                                              | [doc](https://github.com/sooperset/mcp-atlassian)                                       |                                                                                                                              |
| notion	              | MCP server for the Notion API.                                                                                                                                                                                | [doc](https://github.com/makenotion/notion-mcp-server)                                  |                                                                                                                              |
| cloudflare	          | MCP Server for Cloudflare's API.                                                                                                                                                                              | [doc](https://github.com/cloudflare/mcp-server-cloudflare)                              |                                                                                                                              |
| binance	             | MCP Server for Binance's API.                                                                                                                                                                                 | [doc](https://github.com/snjyor/binance-mcp)                                            |                                                                                                                              |
| youtube	             | MCP Server for Youtube's API.                                                                                                                                                                                 | [doc](https://github.com/ZubeidHendricks/youtube-mcp-server)                            |                                                                                                                              |
| shopify	             | MCP Server for Shopify's API.                                                                                                                                                                                 | [doc](https://github.com/Shopify/dev-mcp)                                               |                                                                                                                              |
| duckduckgo	          | MCP Server for Duckduckgo's API.                                                                                                                                                                              | [doc](https://github.com/nickclyde/duckduckgo-mcp-server)                               |                                                                                                                              |
| aliyun	              | MCP Server for Aliyun's API.                                                                                                                                                                                  | [doc](https://github.com/aliyun/alibaba-cloud-ops-mcp-server)                           | [aliyun](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/aliyun/aliyun.go)                            |
| bilibili	            | MCP Server for Bilibili's API.                                                                                                                                                                                | [doc](https://github.com/34892002/bilibili-mcp-js)                                      | [bilibili](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/bilibili/bilibili.go)                      |
| bitcoin	             | MCP Server for Bitcoin's API.                                                                                                                                                                                 | [doc](https://github.com/AbdelStark/bitcoin-mcp)                                        | [bitcoin](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/bitcoin/bitcoin.go)                         |
| airbnb	              | MCP Server for Airbnb's API.                                                                                                                                                                                  | [doc](https://github.com/openbnb-org/mcp-server-airbnb)                                 | [airbnb](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/airbnb/airbnb.go)                            |
| jira	                | MCP Server for Jira's API.                                                                                                                                                                                    | [doc](https://github.com/nguyenvanduocit/jira-mcp)                                      |                                                                                                                              |
| twitter	             | MCP Server for Twitter's API.                                                                                                                                                                                 | [doc](https://github.com/EnesCinr/twitter-mcp)                                          |                                                                                                                              |
| leetcode	            | MCP Server for Leetcode's API.                                                                                                                                                                                | [doc](https://github.com/jinzcdev/leetcode-mcp-server)                                  |                                                                                                                              |
| iterm	               | MCP Server for Iterm's API.                                                                                                                                                                                   | [doc](https://github.com/ferrislucas/iterm-mcp)                                         | [iterm](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/iterm/iterm.go)                               |
| telegram	            | MCP Server for Telegram's API.                                                                                                                                                                                | [doc](https://github.com/chigwell/telegram-mcp)                                         | [telegram](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/telegram/telegram.go)                      |
| zoomeye	             | MCP Server for Zoomeye's API.                                                                                                                                                                                 | [doc](https://github.com/zoomeye-ai/mcp_zoomeye)                                        |                                                                                                                              |
| ipfs	                | MCP Server for IPFS's API.                                                                                                                                                                                    | [doc](https://github.com/alexbakers/mcp-ipfs)                                           |                                                                                                                              |
| k8s	                 | MCP Server for K8s's API.                                                                                                                                                                                     | [doc](https://github.com/alexei-led/k8s-mcp-server)                                     |                                                                                                                              |
| apple-shortcut	      | MCP Server for Apple Shortcut's API.                                                                                                                                                                          | [doc](https://github.com/recursechat/mcp-server-apple-shortcuts)                        |                                                                                                                              |
| ms-365	              | MCP Server for MS356's API.                                                                                                                                                                                   | [doc](https://github.com/softeria/ms-365-mcp-server)                                    |                                                                                                                              |

## 🚀 Getting Started

### Install

install `npx`, `uvx`, `docker`, and put them to env!

```bash
go get github.com/yincongcyincong/mcp-client-go
```

---

## 🧪 Example

[client demo](https://github.com/yincongcyincong/mcp-client-go/tree/main/demo/client)
[app demo](https://github.com/yincongcyincong/mcp-client-go/tree/main/demo/app)
[deepseek demo](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/app/deepseek/deepseek.go)
[openai demo](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/app/openai/openai.go)

---

## 🧱 API Overview

### Initialize Clients

```go
conf := clients.InitStdioMCPClient("npx-amap-maps-mcp-server", "npx", []string{
"AMAP_MAPS_API_KEY=" + AmapApiKey,
}, []string{
"-y",
"@amap/amap-maps-mcp-server",
}, mcp.InitializeRequest{}, nil, nil)
/
conf := clients.InitSSEMCPClient("npx-amap-maps-mcp-server", "http://127.0.0.1", nil, nil, nil)
```

### Register MCP Clients

```go
clients.RegisterMCPClient(context.Background(), []*param.MCPClientConf{conf})
```

### Get MCP Client

```go
client, err := clients.GetMCPClient("npx-amap-maps-mcp-server")
/
client, err := clients.GetMCPClientByToolName("geo_location")
```

### Execute Tools

```go
client.ExecTools(ctx, "tool_name", map[string]interface{}{...})
```

---

## 📄 License

This project is licensed under the [MIT License](./LICENSE).
