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

| MCP Server           | 	Description                                                                                                                                                                                                  | doc                                                                                          | demo                                                                                              |                                                                                            
|----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| redis	               | A Model Context Protocol server that provides access to Redis databases.                                                                                                                                      | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/redis)                   | -                                                                                                 |
| github	              | The GitHub MCP Server is a Model Context Protocol (MCP) server that provides seamless integration with GitHub APIs                                                                                            | [doc](https://github.com/github/github-mcp-server)                                           | [github](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/github/github.go) |
| aws	                 | An MCP server implementation for retrieving information from the AWS Knowledge Base using the Bedrock Agent Runtime.                                                                                          | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/aws-kb-retrieval-server) | -                                                                                                 |
| sequential_thinking	 | An MCP server implementation that provides a tool for dynamic and reflective problem-solving through a structured thinking process.                                                                           | [doc](https://github.com/modelcontextprotocol/servers/tree/HEAD/src/sequentialthinking)      | -                                                                                                 |
| firecrawl	           | A Model Context Protocol (MCP) server implementation that integrates with Firecrawl for web scraping capabilities.                                                                                            | [doc](https://github.com/mendableai/firecrawl-mcp-server)                                    | -                                                                                                 |
| postgresql	          | A Model Context Protocol server that provides read-only access to PostgreSQL databases.                                                                                                                       | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/postgres)                | -                                                                                                 |
| gitlab	              | MCP Server for the GitLab API, enabling project management, file operations, and more.                                                                                                                        | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/gitlab)                  | -                                                                                                 |
| slack	               | MCP Server for the Slack API, enabling Claude to interact with Slack workspaces.                                                                                                                              | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/slack)                   | -                                                                                                 |
| puppeteer	           | A Model Context Protocol server that provides browser automation capabilities using Puppeteer.                                                                                                                | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/puppeteer)               | -                                                                                                 |
| everart	             | Image generation server for Claude Desktop using EverArt's API.                                                                                                                                               | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/everart)                 | -                                                                                                 |
| sentry	              | A Model Context Protocol server for retrieving and analyzing issues from Sentry.io                                                                                                                            | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/sentry)                  | -                                                                                                 |
| filesystem	          | Node.js server implementing Model Context Protocol (MCP) for filesystem operations.                                                                                                                           | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/filesystem)              | -                                                                                                 |
| fetch	               | A Model Context Protocol server that provides web content fetching capabilities. This server enables LLMs to retrieve and process content from web pages, converting HTML to markdown for easier consumption. | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/fetch)                   | -                                                                                                 |
| googlemap	           | MCP Server for the Google Maps API.                                                                                                                                                                           | [doc](https://github.com/modelcontextprotocol/servers/tree/main/src/google-maps)             | -                                                                                                 |
| flomo	               | This is a TypeScript-based MCP server help you write notes to Flomo.                                                                                                                                          | [doc](https://github.com/chatmcp/mcp-server-flomo)                                           | -                                                                                                 |
| chatsum	             | This MCP Server is used to summarize your chat messages.                                                                                                                                                      | [doc](https://github.com/chatmcp/mcp-server-chatsum)                                         | -                                                                                                 |
| amap	                | This repository is a collection of reference implementations for the Model Context Protocol (MCP), as well as references to community built servers and additional resources.                                 | [doc](https://github.com/modelcontextprotocol/servers)                                       | [amap](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/amap/amap.go)       |
| baidumap	            | This MCP Server is used to baidumap                                                                                                                                                                           | [doc](https://github.com/baidu-maps/mcp)                                                     | -                                                                                                 |
| blender	             | BlenderMCP connects Blender to Claude AI through the Model Context Protocol (MCP)                                                                                                                             | [doc](https://github.com/ahujasid/blender-mcp)                                               | -                                                                                                 |
| framelink	           | Give Cursor, Windsurf, Cline, and other AI-powered coding tools access to your Figma files with this Model Context Protocol server.                                                                           | [doc](https://github.com/GLips/Figma-Context-MCP)                                            | -                                                                                                 |
| playwright	          | A Model Context Protocol (MCP) server that provides browser automation capabilities using Playwright.                                                                                                         | [doc]()                                                                                      | -                                                                                                 |
| tavily	              | The Model Context Protocol (MCP) is an open standard that enables AI systems to interact seamlessly with various data sources and tools, facilitating secure, two-way connections.                            | [doc](https://github.com/tavily-ai/tavily-mcp)                                               | -                                                                                                 |

## 🚀 Getting Started

### Install

```bash
go get github.com/yincongcyincong/mcp-client-go
```

---

## 🧪 Example

[client demo](https://github.com/yincongcyincong/mcp-client-go/tree/main/demo/client)    
[app demo](https://github.com/yincongcyincong/mcp-client-go/tree/main/demo/app)

---

## 🧱 API Overview

### Initialize Clients

```go
conf := clients.InitMCPClient("npx-amap-maps-mcp-server", "npx", []string{
"AMAP_MAPS_API_KEY=" + AmapApiKey,
}, []string{
"-y",
"@amap/amap-maps-mcp-server",
}, nil, nil, nil)
```

### Register MCP Clients

```go
clients.RegisterMCPClient(ctx, []*param.MCPClientConf{conf})
```

### Get MCP Client

```go
client, err := clients.GetMCPClient(amap.NpxAmapMapsMcpServer)
```

### Execute Tools

```go
client.ExecTools(ctx, "tool_name", map[string]interface{}{...})
```

#### Common Tool Names (Amap)

- `"maps_regeocode"` – Get detailed address from coordinates
- `"maps_ip_location"` – Get location info based on IP address

---

## 🔐 Authentication

An **access token** is required to authenticate MCP clients. Please contact your service administrator to obtain a valid
token.

---

## 📄 License

This project is licensed under the [MIT License](./LICENSE).
