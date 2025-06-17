package utils

import (
	"encoding/json"
	"reflect"
	"sort"
	"testing"

	"github.com/cohesion-org/deepseek-go"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/revrost/go-openrouter"
	"github.com/sashabaranov/go-openai"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"google.golang.org/genai"
)

// Helper function to create a sample JSON Schema properties map
func createSampleProperties() map[string]interface{} {
	return map[string]interface{}{
		"query": map[string]interface{}{
			"type":        "string",
			"description": "The search query",
		},
		"count": map[string]interface{}{
			"type":        "number",
			"description": "Number of results",
		},
	}
}

// Helper function to create a sample MCP Tool
func createSampleMCPTool(name, description string) mcp.Tool {
	return mcp.Tool{
		Name:        name,
		Description: description,
		InputSchema: mcp.ToolInputSchema{
			Type:       "object",
			Properties: createSampleProperties(),
			Required:   []string{"query"},
		},
	}
}

func TestTransToolsToDPFunctionCall(t *testing.T) {
	mcpTools := []mcp.Tool{
		createSampleMCPTool("search_web", "Search the internet"),
		createSampleMCPTool("get_time", "Get current time"),
	}

	wantTools := []deepseek.Tool{
		{
			Type: "function",
			Function: deepseek.Function{
				Name:        "search_web",
				Description: "Search the internet",
				Parameters: &deepseek.FunctionParameters{
					Type:       "object",
					Properties: createSampleProperties(),
					Required:   []string{"query"},
				},
			},
		},
		{
			Type: "function",
			Function: deepseek.Function{
				Name:        "get_time",
				Description: "Get current time",
				Parameters: &deepseek.FunctionParameters{
					Type:       "object",
					Properties: createSampleProperties(),
					Required:   []string{"query"},
				},
			},
		},
	}

	gotTools := TransToolsToDPFunctionCall(mcpTools)

	if len(gotTools) != len(wantTools) {
		t.Fatalf("TransToolsToDPFunctionCall() got %d tools, want %d", len(gotTools), len(wantTools))
	}

	for i := range gotTools {
		if !reflect.DeepEqual(gotTools[i], wantTools[i]) {
			// Print detailed difference for easier debugging
			gotJSON, _ := json.MarshalIndent(gotTools[i], "", "  ")
			wantJSON, _ := json.MarshalIndent(wantTools[i], "", "  ")
			t.Errorf("TransToolsToDPFunctionCall() tool at index %d mismatch:\nGot:\n%s\nWant:\n%s", i, gotJSON, wantJSON)
		}
	}

	// Test with empty input
	gotEmpty := TransToolsToDPFunctionCall([]mcp.Tool{})
	if len(gotEmpty) != 0 {
		t.Errorf("TransToolsToDPFunctionCall() with empty input got %d tools, want 0", len(gotEmpty))
	}
}

func TestTransToolsToGeminiFunctionCall(t *testing.T) {
	mcpTools := []mcp.Tool{
		createSampleMCPTool("calendar_event", "Create a calendar event"),
	}

	// Gemini's Schema requires Properties to be map[string]*genai.Schema
	// We need to unmarshal the sample properties to this type first for comparison
	var expectedProperties map[string]*genai.Schema
	propBytes, _ := json.Marshal(createSampleProperties())
	json.Unmarshal(propBytes, &expectedProperties)

	wantTools := []*genai.Tool{
		{
			FunctionDeclarations: []*genai.FunctionDeclaration{
				{
					Name:        "calendar_event",
					Description: "Create a calendar event",
					Parameters: &genai.Schema{
						Type:       genai.TypeObject,
						Properties: expectedProperties,
						Required:   []string{"query"},
					},
				},
			},
		},
	}

	gotTools := TransToolsToGeminiFunctionCall(mcpTools)

	if len(gotTools) != len(wantTools) {
		t.Fatalf("TransToolsToGeminiFunctionCall() got %d tools, want %d", len(gotTools), len(wantTools))
	}

	// DeepEqual for pointers to structs needs care.
	// Marshal to JSON for robust comparison of nested structures.
	gotJSON, _ := json.MarshalIndent(gotTools, "", "  ")
	wantJSON, _ := json.MarshalIndent(wantTools, "", "  ")

	if string(gotJSON) != string(wantJSON) {
		t.Errorf("TransToolsToGeminiFunctionCall() mismatch:\nGot:\n%s\nWant:\n%s", gotJSON, wantJSON)
	}

	// Test with empty input
	gotEmpty := TransToolsToGeminiFunctionCall([]mcp.Tool{})
	if len(gotEmpty) != 0 {
		t.Errorf("TransToolsToGeminiFunctionCall() with empty input got %d tools, want 0", len(gotEmpty))
	}
}

func TestTransToolsToVolFunctionCall(t *testing.T) {
	mcpTools := []mcp.Tool{
		createSampleMCPTool("image_generation", "Generate an image"),
	}

	wantTools := []*model.Tool{
		{
			Type: "function",
			Function: &model.FunctionDefinition{
				Name:        "image_generation",
				Description: "Generate an image",
				Parameters: map[string]interface{}{
					"type":       "object",
					"properties": createSampleProperties(),
					"required":   []string{"query"},
				},
			},
		},
	}

	gotTools := TransToolsToVolFunctionCall(mcpTools)

	if len(gotTools) != len(wantTools) {
		t.Fatalf("TransToolsToVolFunctionCall() got %d tools, want %d", len(gotTools), len(wantTools))
	}

	// Compare by marshaling to JSON due to interface{} in Parameters and pointer types
	gotJSON, _ := json.MarshalIndent(gotTools, "", "  ")
	wantJSON, _ := json.MarshalIndent(wantTools, "", "  ")

	if string(gotJSON) != string(wantJSON) {
		t.Errorf("TransToolsToVolFunctionCall() mismatch:\nGot:\n%s\nWant:\n%s", gotJSON, wantJSON)
	}

	// Test with empty input
	gotEmpty := TransToolsToVolFunctionCall([]mcp.Tool{})
	if len(gotEmpty) != 0 {
		t.Errorf("TransToolsToVolFunctionCall() with empty input got %d tools, want 0", len(gotEmpty))
	}
}

func TestTransToolsToChatGPTFunctionCall(t *testing.T) {
	mcpTools := []mcp.Tool{
		createSampleMCPTool("weather_app", "Get current weather"),
	}

	wantTools := []openai.Tool{
		{
			Type: "function",
			Function: &openai.FunctionDefinition{
				Name:        "weather_app",
				Description: "Get current weather",
				Parameters: &deepseek.FunctionParameters{ // Note: deepseek.FunctionParameters is used here
					Type:       "object",
					Properties: createSampleProperties(),
					Required:   []string{"query"},
				},
			},
		},
	}

	gotTools := TransToolsToChatGPTFunctionCall(mcpTools)

	if len(gotTools) != len(wantTools) {
		t.Fatalf("TransToolsToChatGPTFunctionCall() got %d tools, want %d", len(gotTools), len(wantTools))
	}

	// Due to potential differences in internal struct representations (even if content is same)
	// it's better to marshal to JSON and compare for complex nested structs if DeepEqual fails
	// or specific fields are not comparable directly.
	// For this test, reflect.DeepEqual should work given the structure.
	if !reflect.DeepEqual(gotTools, wantTools) {
		gotJSON, _ := json.MarshalIndent(gotTools, "", "  ")
		wantJSON, _ := json.MarshalIndent(wantTools, "", "  ")
		t.Errorf("TransToolsToChatGPTFunctionCall() mismatch:\nGot:\n%s\nWant:\n%s", gotJSON, wantJSON)
	}

	// Test with empty input
	gotEmpty := TransToolsToChatGPTFunctionCall([]mcp.Tool{})
	if len(gotEmpty) != 0 {
		t.Errorf("TransToolsToChatGPTFunctionCall() with empty input got %d tools, want 0", len(gotEmpty))
	}
}

func TestTransToolsToOpenRouterFunctionCall(t *testing.T) {
	mcpTools := []mcp.Tool{
		createSampleMCPTool("blog_post", "Write a blog post"),
	}

	wantTools := []openrouter.Tool{
		{
			Type: "function",
			Function: &openrouter.FunctionDefinition{
				Name:        "blog_post",
				Description: "Write a blog post",
				Parameters: map[string]interface{}{
					"type":       "object",
					"properties": createSampleProperties(),
					"required":   []string{"query"},
				},
			},
		},
	}

	gotTools := TransToolsToOpenRouterFunctionCall(mcpTools)

	if len(gotTools) != len(wantTools) {
		t.Fatalf("TransToolsToOpenRouterFunctionCall() got %d tools, want %d", len(gotTools), len(wantTools))
	}

	if !reflect.DeepEqual(gotTools, wantTools) {
		gotJSON, _ := json.MarshalIndent(gotTools, "", "  ")
		wantJSON, _ := json.MarshalIndent(wantTools, "", "  ")
		t.Errorf("TransToolsToOpenRouterFunctionCall() mismatch:\nGot:\n%s\nWant:\n%s", gotJSON, wantJSON)
	}

	// Test with empty input
	gotEmpty := TransToolsToOpenRouterFunctionCall([]mcp.Tool{})
	if len(gotEmpty) != 0 {
		t.Errorf("TransToolsToOpenRouterFunctionCall() with empty input got %d tools, want 0", len(gotEmpty))
	}
}

func TestChangeEnvMapToSlice(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string
		want []string
	}{
		{
			name: "empty map",
			env:  map[string]string{},
			want: []string{},
		},
		{
			name: "single entry",
			env:  map[string]string{"KEY1": "VALUE1"},
			want: []string{"KEY1=VALUE1"},
		},
		{
			name: "multiple entries",
			env:  map[string]string{"KEY1": "VALUE1", "KEY2": "VALUE2"},
			// Order is not guaranteed for map iteration, so we sort both for comparison
			want: []string{"KEY1=VALUE1", "KEY2=VALUE2"},
		},
		{
			name: "map with empty string value",
			env:  map[string]string{"KEY_EMPTY": ""},
			want: []string{"KEY_EMPTY="},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ChangeEnvMapToSlice(tt.env)
			// Sort the slice to ensure consistent order for comparison, as map iteration order is not guaranteed.
			// This is important for testing functions that convert maps to slices.
			sort.Strings(got)
			sort.Strings(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangeEnvMapToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReturnString(t *testing.T) {
	tests := []struct {
		name   string
		result *mcp.CallToolResult
		want   string
	}{
		{
			name: "single text content",
			result: &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{Text: "Hello, world!"},
				},
			},
			want: "Hello, world!",
		},
		{
			name: "multiple text contents",
			result: &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{Text: "Part 1. "},
					mcp.TextContent{Text: "Part 2."},
				},
			},
			want: "Part 1. Part 2.",
		},
		{
			name:   "empty content",
			result: &mcp.CallToolResult{Content: []mcp.Content{}},
			want:   "",
		},
		{
			name:   "nil result",
			result: nil,
			want:   "", // Should not panic, but return empty string for nil receiver or content
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReturnString(tt.result)
			if got != tt.want {
				t.Errorf("ReturnString() = %q, want %q", got, tt.want)
			}
		})
	}
}
