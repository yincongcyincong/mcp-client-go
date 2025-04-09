package amap

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/utils"
)

var (
	amapMCPClient *client.StdioMCPClient
)

func InitAmapMCPClient(ctx context.Context, AmapApiKey string) error {
	var err error
	amapMCPClient, err = client.NewStdioMCPClient(
		"npx",
		[]string{
			"AMAP_MAPS_API_KEY=" + AmapApiKey,
		},
		"-y",
		"@amap/amap-maps-mcp-server",
	)

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/amap-maps",
		Version: "0.1.0",
	}

	initResult, err := amapMCPClient.Initialize(ctx, initRequest)
	if err != nil {
		return err
	}

	fmt.Printf(
		"Initialized with server: %s %s\n\n",
		initResult.ServerInfo.Name,
		initResult.ServerInfo.Version,
	)

	return err
}

func GetAmapMCPClient() *client.StdioMCPClient {
	return amapMCPClient
}

func GetAmapTools(ctx context.Context, cursor mcp.Cursor) ([]mcp.Tool, error) {
	toolsRequest := mcp.ListToolsRequest{}
	toolsRequest.Params.Cursor = cursor
	tools, err := amapMCPClient.ListTools(ctx, toolsRequest)
	if err != nil {
		return nil, err
	}
	return tools.Tools, nil
}

func MapsRegeocode(ctx context.Context, location string) (string, error) {
	listDirRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	listDirRequest.Params.Name = "maps_regeocode"
	listDirRequest.Params.Arguments = map[string]interface{}{
		"location": location,
	}

	result, err := amapMCPClient.CallTool(ctx, listDirRequest)
	if err != nil {
		return "", err
	}

	return utils.ReturnString(result), nil
}

func MapsGeo(ctx context.Context, address, city string) (string, error) {
	listDirRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	listDirRequest.Params.Name = "maps_geo"
	listDirRequest.Params.Arguments = map[string]interface{}{
		"address": address,
		"city":    city,
	}

	result, err := amapMCPClient.CallTool(ctx, listDirRequest)
	if err != nil {
		return "", err
	}

	return utils.ReturnString(result), nil
}

func MapsIPLocation(ctx context.Context, IP string) (string, error) {
	listDirRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	listDirRequest.Params.Name = "maps_ip_location"
	listDirRequest.Params.Arguments = map[string]interface{}{
		"ip": IP,
	}

	result, err := amapMCPClient.CallTool(ctx, listDirRequest)
	if err != nil {
		return "", err
	}

	return utils.ReturnString(result), nil
}

func MapsWeather(ctx context.Context, city string) (string, error) {
	listDirRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	listDirRequest.Params.Name = "maps_weather"
	listDirRequest.Params.Arguments = map[string]interface{}{
		"city": city,
	}

	result, err := amapMCPClient.CallTool(ctx, listDirRequest)
	if err != nil {
		return "", err
	}

	return utils.ReturnString(result), nil
}

func MapsSearchDetail(ctx context.Context, id string) (string, error) {
	listDirRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	listDirRequest.Params.Name = "maps_search_detail"
	listDirRequest.Params.Arguments = map[string]interface{}{
		"id": id,
	}

	result, err := amapMCPClient.CallTool(ctx, listDirRequest)
	if err != nil {
		return "", err
	}

	return utils.ReturnString(result), nil
}

func MapsBicycling(ctx context.Context, destination, origin string) (string, error) {
	listDirRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	listDirRequest.Params.Name = "maps_regeocode"
	listDirRequest.Params.Arguments = map[string]interface{}{
		"destination": destination,
		"origin":      origin,
	}

	result, err := amapMCPClient.CallTool(ctx, listDirRequest)
	if err != nil {
		return "", err
	}

	return utils.ReturnString(result), nil
}

func MapsDirectionWalking(ctx context.Context, destination, origin string) (string, error) {
	listDirRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	listDirRequest.Params.Name = "maps_direction_walking"
	listDirRequest.Params.Arguments = map[string]interface{}{
		"destination": destination,
		"origin":      origin,
	}

	result, err := amapMCPClient.CallTool(ctx, listDirRequest)
	if err != nil {
		return "", err
	}

	return utils.ReturnString(result), nil
}

func MapsDirectionDriving(ctx context.Context, destination, origin string) (string, error) {
	listDirRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	listDirRequest.Params.Name = "maps_direction_driving"
	listDirRequest.Params.Arguments = map[string]interface{}{
		"destination": destination,
		"origin":      origin,
	}

	result, err := amapMCPClient.CallTool(ctx, listDirRequest)
	if err != nil {
		return "", err
	}

	return utils.ReturnString(result), nil
}

func MapsDirectionTransitIntegrated(ctx context.Context, city, cityd, destination, origin string) (string, error) {
	listDirRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	listDirRequest.Params.Name = "maps_direction_transit_integrated"
	listDirRequest.Params.Arguments = map[string]interface{}{
		"city":        city,
		"cityd":       cityd,
		"destination": destination,
		"origin":      origin,
	}

	result, err := amapMCPClient.CallTool(ctx, listDirRequest)
	if err != nil {
		return "", err
	}

	return utils.ReturnString(result), nil
}

func MapsDistance(ctx context.Context, destination, origin, ty string) (string, error) {
	listDirRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	listDirRequest.Params.Name = "maps_distance"
	listDirRequest.Params.Arguments = map[string]interface{}{
		"destination": destination,
		"origin":      origin,
		"type":        ty,
	}

	result, err := amapMCPClient.CallTool(ctx, listDirRequest)
	if err != nil {
		return "", err
	}

	return utils.ReturnString(result), nil
}

func MapsTextSearch(ctx context.Context, city, keywords, types string) (string, error) {
	listDirRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	listDirRequest.Params.Name = "maps_text_search"
	listDirRequest.Params.Arguments = map[string]interface{}{
		"city":     city,
		"keywords": keywords,
		"types":    types,
	}

	result, err := amapMCPClient.CallTool(ctx, listDirRequest)
	if err != nil {
		return "", err
	}

	return utils.ReturnString(result), nil
}

func MapsAroundSearch(ctx context.Context, keywords, location, radius string) (string, error) {
	listDirRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	listDirRequest.Params.Name = "maps_around_search"
	listDirRequest.Params.Arguments = map[string]interface{}{
		"keywords": keywords,
		"location": location,
		"radius":   radius,
	}

	result, err := amapMCPClient.CallTool(ctx, listDirRequest)
	if err != nil {
		return "", err
	}

	return utils.ReturnString(result), nil
}
