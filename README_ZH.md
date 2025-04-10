# 📦 mcp-client-go

`mcp-client-go` 是一个用于 **Model Context Protocol（MCP）** 的 Golang 客户端库。它允许开发者通过统一的 API 注册和调用多个 MCP 协议服务，例如高德地图（Amap）。

---

## ✨ 特性

- 简单集成 MCP 协议服务
- 模块化支持多种服务类型（如 Amap、Github、GoogleMap 等）
- 支持统一注册和客户端管理
- 提供简单直观的工具执行接口

---

## 📋 支持的服务列表

| MCP 服务名称           | 描述                                                                                                                                                              | 文档链接                                                                                      |
|------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------|
| redis	                | 提供对 Redis 数据库的访问能力。                                                                                                                                   | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/redis)                   |
| github	              | 集成 GitHub API 的 MCP 服务。                                                                                                                                     | [文档](https://github.com/github/github-mcp-server)                                           |
| aws	                  | 使用 Bedrock Agent Runtime 从 AWS 知识库中检索信息。                                                                                                               | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/aws-kb-retrieval-server) |
| sequential_thinking	  | 提供结构化思维流程的动态反思与问题解决工具。                                                                                                                       | [文档](https://github.com/modelcontextprotocol/servers/tree/HEAD/src/sequentialthinking)      |
| firecrawl	            | 集成 Firecrawl 的网页抓取能力。                                                                                                                                   | [文档](https://github.com/mendableai/firecrawl-mcp-server)                                    |
| postgresql	          | 提供 PostgreSQL 数据库的只读访问能力。                                                                                                                              | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/postgres)                |
| gitlab	              | 集成 GitLab API，实现项目管理、文件操作等功能。                                                                                                                    | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/gitlab)                  |
| slack	                | 集成 Slack API，支持 Claude 与 Slack 工作区交互。                                                                                                                   | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/slack)                   |
| puppeteer	            | 使用 Puppeteer 实现浏览器自动化的 MCP 服务。                                                                                                                       | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/puppeteer)               |
| everart	              | 使用 EverArt API 进行图像生成，服务于 Claude Desktop。                                                                                                               | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/everart)                 |
| sentry	              | 集成 Sentry.io 的问题分析和获取能力。                                                                                                                              | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/sentry)                  |
| filesystem	          | 使用 Node.js 实现的文件系统操作 MCP 服务。                                                                                                                         | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/filesystem)              |
| fetch	                | 获取网页内容，将 HTML 转换为 markdown 以便 LLM 消化。                                                                                                               | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/fetch)                   |
| googlemap	            | 集成 Google Maps API 的 MCP 服务。                                                                                                                                | [文档](https://github.com/modelcontextprotocol/servers/tree/main/src/google-maps)             |
| flomo	                | 用于向 Flomo 写笔记的 MCP 服务，基于 TypeScript 实现。                                                                                                               | [文档](https://github.com/chatmcp/mcp-server-flomo)                                           |
| chatsum	              | 用于总结聊天记录的 MCP 服务。                                                                                                                                    | [文档](https://github.com/chatmcp/mcp-server-chatsum)                                         |
| amap	                  | MCP 协议服务参考实现集，包括高德地图等常用接口示例。                                                                                                               | [文档](https://github.com/modelcontextprotocol/servers)                                       |
| baidumap	            | 用于调用百度地图的 MCP 服务。                                                                                                                                    | [文档](https://github.com/baidu-maps/mcp)                                                     |
| blender	              | 让 Blender 通过 MCP 协议与 Claude AI 连接的服务。                                                                                                                  | [文档](https://github.com/ahujasid/blender-mcp)                                               |
| framelink	            | 让 Cursor、Windsurf、Cline 等 AI 编程工具访问你的 Figma 文件。                                                                                                       | [文档](https://github.com/GLips/Figma-Context-MCP)                                            |
| playwright	            | 使用 Playwright 实现浏览器自动化的 MCP 服务。                                                                                                                     | [文档]()                                                                                      |
| tavily	              | MCP 协议标准化接口服务，连接多种数据源与工具，支持双向安全通信。                                                                                                  | [文档](https://github.com/tavily-ai/tavily-mcp)                                               |

---

## 🚀 快速开始

### 安装

```bash
go get github.com/yincongcyincong/mcp-client-go
```

---

## 🧪 示例

[客户端示例]()  
[应用示例]()

---

## 🧱 API 总览

### 初始化客户端

```go
conf := clients.InitMCPClient("npx-amap-maps-mcp-server", "npx", []string{
  "AMAP_MAPS_API_KEY=" + AmapApiKey,
}, []string{
  "-y",
  "@amap/amap-maps-mcp-server",
}, nil, nil, nil)
```

### 注册 MCP 客户端

```go
clients.RegisterMCPClient(ctx, []*param.MCPClientConf{conf})
```

### 获取 MCP 客户端

```go
client, err := clients.GetMCPClient(amap.NpxAmapMapsMcpServer)
```

### 执行工具

```go
client.ExecTools(ctx, "tool_name", map[string]interface{}{...})
```

#### 常用工具名称（Amap）

- `"maps_regeocode"` – 通过坐标获取详细地址信息
- `"maps_ip_location"` – 根据 IP 地址获取定位信息

---

## 🔐 认证说明

使用 MCP 客户端前必须配置 **Access Token（访问令牌）**。请联系你的服务管理员以获取有效的令牌。

---

## 📄 许可证

本项目基于 [MIT License](./LICENSE) 开源许可协议发布。
