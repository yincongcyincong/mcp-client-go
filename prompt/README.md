Sure! Here's the **English version** of the `README.md` based on your project, describing a collection of prompts that trigger function calls, organized by service:

---

# 📚 Function Call Prompt Collection (for MCP Server Integration)

This repository contains a curated collection of **natural language prompts** designed to trigger OpenAI Function Calls. Each prompt corresponds to a specific backend function and is intended to work seamlessly with an MCP Server.

Each prompt is written in Markdown and organized by third-party services such as **Aliyun**, **Amap**, **GitHub**, **Google Maps**, and **Grafana**.

---

## 📁 Project Structure

```
prompt/
├── aliyun/aliyun.md
├── amap/amap.md
├── github/github.md
├── googlemap/googlemap.md
├── grafana/grafana.md
```

Each `.md` file contains prompt examples that are mapped to specific backend functions via Function Call.

---

## 🔍 Prompt Example Format

Here’s a sample from `prompt/amap/amap.md`:

### ✅ 1. `maps_regeocode` – Convert coordinates to address

**Prompt**:

> What is the address of the location at `116.481488,39.990464`?

---

### ✅ 2. `maps_geo` – Convert address to coordinates

**Prompt**:

> Can you get the coordinates for "Wangjing SOHO, Chaoyang District, Beijing"?

---

## 🧩 Supported Services and Functions

| Service     | Function Examples                | Description                               |
| ----------- | -------------------------------- | ----------------------------------------- |
| Amap        | `maps_geo`, `maps_regeocode`     | Coordinate ↔️ Address conversion, weather |
| GitHub      | `list_repos`, `create_issue`     | Repository listing, issue creation        |
| Aliyun      | `list_ecs`, `query_logs`         | ECS management, log queries               |
| Google Maps | `search_place`, `get_directions` | Place search, route planning              |
| Grafana     | `query_dashboard`, `get_alerts`  | Dashboard data, alert status              |

---

## 🛠 How to Add New Prompts

To add new prompts:

1. Go to the corresponding service folder (e.g., `prompt/github/`)
2. Edit or create the `.md` file
3. Follow this format:

```md
`function_name` – Brief Description
- Example of a natural language request
```

---

## 📎 License

This project is licensed under the MIT License. Feel free to use and modify it.

