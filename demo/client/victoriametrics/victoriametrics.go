package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/yincongcyincong/mcp-client-go/clients"
    "github.com/yincongcyincong/mcp-client-go/clients/param"
    "github.com/yincongcyincong/mcp-client-go/clients/victoriametrics"
    "log"
    "time"
)

func main() {
    mc := victoriametrics.InitVictoriaMetricsMCPClient(&victoriametrics.VictoriaMetricsParam{
        VMUrl: "http://localhost:8428",
    }, "", nil, nil, nil)

    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
    defer cancel()

    errs := clients.RegisterMCPClient(ctx, []*param.MCPClientConf{mc})
    if len(errs) > 0 {
        log.Fatal("InitMCPClient failed:", errs)
    }

    c, err := clients.GetMCPClient(victoriametrics.NpxVictoriaMetricsMcpServer)
    if err != nil {
        log.Fatal("GetMCPClient failed:", err)
    }

    for _, tool := range c.Tools {
        toolByte, _ := json.Marshal(tool)
        fmt.Println(string(toolByte))
    }

    data, err := c.ExecTools(ctx, "vm_query_range", map[string]interface{}{
        "query": "cpu_usage",
        "start": 1744596752,
        "end":   1744598552,
    })
    if err != nil {
        log.Fatal("ExecTools failed:", err)
    }

    fmt.Println(data)

}
