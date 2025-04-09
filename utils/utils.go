package utils

import (
	"encoding/json"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

func ReturnString(result *mcp.CallToolResult) string {
	var res strings.Builder
	for _, content := range result.Content {
		if textContent, ok := content.(mcp.TextContent); ok {
			res.WriteString(textContent.Text)
		} else {
			jsonBytes, _ := json.MarshalIndent(content, "", "  ")
			res.Write(jsonBytes)
		}
	}

	return res.String()
}
