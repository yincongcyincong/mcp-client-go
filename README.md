
# 📦 mcp-client-go

`mcp-client-go` is a Golang client library for the **Model Context Protocol (MCP)**. It allows developers to register and interact with various MCP-based services such as Amap (Gaode Maps) using a unified API.

---

## ✨ Features

- Easy integration with MCP-compatible services
- Modular support for service types (e.g., Amap, Github, GoogleMap)
- Unified registration and client management
- Simple and intuitive tool execution interface

---

## 📋 Supported Services 

## 🚀 Getting Started

### Install

```bash
go get github.com/yincongcyincong/mcp-client-go
```

---

## 🧪 Example

[client demo]()
[app demo]()

---

## 🧱 API Overview

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

An **access token** is required to authenticate MCP clients. Please contact your service administrator to obtain a valid token.

---

## 📄 License

This project is licensed under the [MIT License](./LICENSE).
