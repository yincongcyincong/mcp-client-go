package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/yincongcyincong/mcp-client-go/clients"
    "github.com/yincongcyincong/mcp-client-go/clients/grafana"
    "github.com/yincongcyincong/mcp-client-go/clients/param"
    "log"
    "time"
)

func main() {
    // start grafana-mcp first
    mc := grafana.InitGrafanaSSEMCPClient("http://localhost:8000/sse", nil, "", nil, nil, nil)

    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
    defer cancel()

    errs := clients.RegisterMCPClient(ctx, []*param.MCPClientConf{mc})
    if len(errs) > 0 {
        log.Fatal("InitMCPClient failed:", errs)
    }

    c, err := clients.GetMCPClient(grafana.SSEGrafanaMcpServer)
    if err != nil {
        log.Fatal("GetMCPClient failed:", err)
    }

    for _, tool := range c.Tools {
        toolByte, _ := json.Marshal(tool)
        fmt.Println(string(toolByte))
    }

    data, err := c.ExecTools(ctx, "get_datasource_by_uid", map[string]interface{}{
        "uid": "aeatj835ymyv4e",
    })
    if err != nil {
        log.Fatal("ExecTools failed:", err)
    }

    fmt.Println(data)

}
