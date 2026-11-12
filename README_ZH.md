# 📦 mcp-client-go 中文文档

`mcp-client-go` 是一个用于 **模型上下文协议(Model Context Protocol, MCP)** 的 Golang 客户端库。它允许开发者通过统一API注册和交互各种基于MCP的服务，如高德地图(Amap)等。

---

## ✨ 功能特性

- 轻松集成MCP兼容服务
- 模块化支持多种服务类型(如高德地图、Github、谷歌地图等)
- 统一的注册和客户端管理
- 简单直观的工具执行接口

---

## 📋 支持的服务

| MCP 服务            | 	描述                                                                                                                                                                                                  | 文档                                                                                          | 示例                                                                                                                |
|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------|
| redis	              | 提供Redis数据库访问的MCP服务                                                                                                                                                                          | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/redis)                   | [redis示例](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/redis/redis.go)                  |
| github	             | 提供GitHub API集成的MCP服务                                                                                                                                                                           | [文档](https://github.com/github/github-mcp-server)                                           | [github示例](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/github/github.go)               |
| aws	                | 通过Bedrock Agent Runtime从AWS知识库获取信息的MCP服务                                                                                                                                                 | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/aws-kb-retrieval-server) | -                                                                                                                   |
| 顺序思维(sequential_thinking) | 提供结构化思维过程的动态问题解决工具                                                                                                                                                                   | [文档](https://github.com/modelcontextprotocol/servers/tree/HEAD/src/sequentialthinking)      | -                                                                                                                   |
| firecrawl	          | 集成Firecrawl网页抓取能力的MCP服务                                                                                                                                                                    | [文档](https://github.com/mendableai/firecrawl-mcp-server)                                    | -                                                                                                                   |
| postgresql	         | 提供PostgreSQL数据库只读访问的MCP服务                                                                                                                                                                 | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/postgres)                | -                                                                                                                   |
| gitlab	             | 提供GitLab API集成的MCP服务，支持项目管理、文件操作等功能                                                                                                                                             | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/gitlab)                  | -                                                                                                                   |
| slack	              | 提供Slack API集成的MCP服务，支持与Slack工作区交互                                                                                                                                                     | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/slack)                   | -                                                                                                                   |
| puppeteer	          | 使用Puppeteer提供浏览器自动化能力的MCP服务                                                                                                                                                           | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/puppeteer)               | -                                                                                                                   |
| everart	            | 使用EverArt API为Claude Desktop提供图像生成服务                                                                                                                                                       | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/everart)                 | -                                                                                                                   |
| sentry	             | 从Sentry.io获取和分析问题的MCP服务                                                                                                                                                                    | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/sentry)                  | -                                                                                                                   |
| 文件系统(filesystem) | 提供文件系统操作的Node.js MCP服务                                                                                                                                                                     | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/filesystem)              | [filesystem示例](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/filesystem/filesystem.go)  |
| fetch	              | 提供网页内容抓取能力的MCP服务，可将HTML转换为Markdown格式                                                                                                                                             | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/fetch)                   | -                                                                                                                   |
| 谷歌地图(googlemap) | 提供Google Maps API集成的MCP服务                                                                                                                                                                      | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/google-maps)             | [googlemap示例](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/googlemap/googlemap.go)      |
| flomo	              | 基于TypeScript的MCP服务，支持向Flomo写笔记                                                                                                                                                            | [文档](https://github.com/chatmcp/mcp-server-flomo)                                           | -                                                                                                                   |
| 聊天摘要(chatsum)   | 用于总结聊天消息的MCP服务                                                                                                                                                                             | [文档](https://github.com/chatmcp/mcp-server-chatsum)                                         | -                                                                                                                   |
| 高德地图(amap)      | 提供高德地图API集成的MCP服务                                                                                                                                                                          | [文档](https://github.com/modelcontextprotocol/servers)                                       | [amap示例](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/amap/amap.go)                     |
| 百度地图(baidumap)  | 提供百度地图API集成的MCP服务                                                                                                                                                                          | [文档](https://github.com/baidu-maps/mcp)                                                     | -                                                                                                                   |
| blender	            | 通过MCP连接Blender和Claude AI的服务                                                                                                                                                                   | [文档](https://github.com/ahujasid/blender-mcp)                                               | -                                                                                                                   |
| framelink	          | 为Cursor、Windsurf等AI编程工具提供Figma文件访问能力的MCP服务                                                                                                                                          | [文档](https://github.com/GLips/Figma-Context-MCP)                                            | -                                                                                                                   |
| playwright	         | 使用Playwright提供浏览器自动化能力的MCP服务                                                                                                                                                           | [文档](https://github.com/microsoft/playwright-mcp)                                           | [playwright示例](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/playwright/playwright.go)   |
| parallel | 通过 Streamable HTTP 提供无需 API 密钥的免费网页搜索和内容提取。 | [文档](https://docs.parallel.ai/integrations/mcp/search-mcp) | [parallel 示例](demo/client/parallel/parallel.go) |
| tavily	             | 支持与各种数据源和工具安全双向连接的MCP标准实现                                                                                                                                                       | [文档](https://github.com/tavily-ai/tavily-mcp)                                               | -                                                                                                                   |
| 时间(time)          | 提供时间和时区转换功能的MCP服务                                                                                                                                                                       | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/time)                    | [time示例](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/client/time/time.go)                     |
| victoriametrics	    | 提供VictoriaMetrics数据库访问的MCP服务                                                                                                                                                                | [文档](https://github.com/yincongcyincong/VictoriaMetrics-mcp-server)                         | -                                                                                                                   |

## 🚀 快速开始

### 安装

请先安装`npx`、`uvx`和`docker`，并确保它们已加入环境变量！

```bash
go get github.com/yincongcyincong/mcp-client-go
```

---

## 🧪 示例代码

[客户端示例](https://github.com/yincongcyincong/mcp-client-go/tree/main/demo/client)
[应用示例](https://github.com/yincongcyincong/mcp-client-go/tree/main/demo/app)
[DeepSeek集成示例](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/app/deepseek/deepseek.go)
[OpenAI集成示例](https://github.com/yincongcyincong/mcp-client-go/blob/main/demo/app/openai/openai.go)

---

## 🧱 API概览

### 初始化客户端

```go
// 标准IO模式初始化
conf := clients.InitStdioMCPClient("npx-amap-maps-mcp-server", "npx", []string{
		"AMAP_MAPS_API_KEY=" + AmapApiKey,
		}, []string{
		"-y",
		"@amap/amap-maps-mcp-server",
		})
// 标准SSE初始化
conf := clients.InitSSEMCPClient("npx-amap-maps-mcp-server", "http://127.0.0.1", nil)
```

### 注册MCP客户端

```go
clients.RegisterMCPClient(ctx, []*param.MCPClientConf{conf})
```

### 获取MCP客户端

```go
// 通过服务名获取
client, err := clients.GetMCPClient(amap.NpxAmapMapsMcpServer)

// 通过工具名获取
client, err := clients.GetMCPClientByToolName("geo_location")
```

### 执行工具

```go
client.ExecTools(ctx, "tool_name", map[string]interface{}{...})
```

---

## 📄 许可证

本项目采用 [MIT 许可证](./LICENSE)。
