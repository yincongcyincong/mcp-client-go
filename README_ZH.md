# 📦 mcp-client-go

`mcp-client-go` 是基于 Model Context Protocol (MCP) 的 Golang 客户端，支持注册多种服务客户端，并通过统一接口调用工具方法，例如高德地图相关接口。

---

## ✨ 功能特性

- 快速接入 MCP 协议服务
- 支持多种服务模块，如 Amap（高德）,github, googleMap 等
- 统一注册与管理客户端
- 简洁的接口调用方式

---

## 📋 模块支持 

## 🚀 快速开始

### 安装依赖

```bash
go get github.com/yincongcyincong/mcp-client-go
```

---

## 🧪 示例代码

[client demo]()
[app demo]()
```

---

## 📘 接口说明

### 注册客户端

```go
clients.RegisterMCPClient(ctx, []*param.MCPClientConf{conf})
```

### 获取客户端实例

```go
client, err := clients.GetMCPClient(amap.NpxAmapMapsMcpServer)
```

### 执行工具方法

```go
client.ExecTools(ctx, "tool_name", map[string]interface{}{...})
```

常用工具方法名称（以 Amap 为例）：

- `"maps_regeocode"`：根据经纬度获取详细地址
- `"maps_ip_location"`：根据 IP 获取地理位置

---

## 🔐 鉴权说明

使用 MCP 客户端需提供有效的 `access token`。如需申请访问权限，请联系相关平台负责人。

---

## 📄 协议许可

本项目使用 MIT License。详情请参见 [LICENSE](./LICENSE)。

---
