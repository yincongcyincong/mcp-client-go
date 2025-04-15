package aws

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/clients/param"
)

const (
	NpxAwsCoreMcpServer             = "npx-aws-core-mcp-server"
	NpxAwsCanvasMcpServer           = "npx-aws-canvas-mcp-server"
	NpxAwsBedrockKbRetrievalsServer = "npx-aws-bedrock-kms-server"
	NpxAwsAnalysisServer            = "npx-aws-analysis-server"
	NpxAwsCDKServer                 = "npx-aws-cdk-server"
	NpxAwsDocumentServer            = "npx-aws-document-server"
	NpxAwsLambdaServer              = "npx-aws-lambda-server"
)

type AwsCoreParams struct {
	FastMCPLogLevel string
	MCPSettingPath  string
}

func InitAwsCoreMCPClient(p *AwsCoreParams, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	awsMCPClient := &param.MCPClientConf{
		Name: NpxAwsCoreMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uvx",
			Env: []string{
				"FASTMCP_LOG_LEVEL=" + p.FastMCPLogLevel,
				"MCP_SETTINGS_PATH=" + p.MCPSettingPath,
			},
			Args: []string{
				"awslabs.core-mcp-server@latest",
			},
			InitReq: mcp.InitializeRequest{},
		},
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/aws-core",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	awsMCPClient.StdioClientConf.InitReq = initRequest

	return awsMCPClient
}

type AwsNovaCanvasParams struct {
	FastMCPLogLevel string
	AwsProfile      string
	AwsRegion       string
}

func InitAwsNovaCanvasMCPClient(p *AwsNovaCanvasParams, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	awsMCPClient := &param.MCPClientConf{
		Name: NpxAwsCanvasMcpServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uvx",
			Env: []string{
				"FASTMCP_LOG_LEVEL=" + p.FastMCPLogLevel,
				"AWS_PROFILE=" + p.AwsProfile,
				"AWS_REGION=" + p.AwsRegion,
			},
			Args: []string{
				"awslabs.nova-canvas-mcp-server@latest",
			},
			InitReq: mcp.InitializeRequest{},
		},
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/aws-nova-canvas",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	awsMCPClient.StdioClientConf.InitReq = initRequest

	return awsMCPClient
}

type AwsBedrockKbRetrievalsParams struct {
	FastMCPLogLevel string
	AwsProfile      string
	AwsRegion       string
}

func InitAwsBedrockKbRetrievalMCPClient(p *AwsBedrockKbRetrievalsParams, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	awsMCPClient := &param.MCPClientConf{
		Name: NpxAwsBedrockKbRetrievalsServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uvx",
			Env: []string{
				"FASTMCP_LOG_LEVEL=" + p.FastMCPLogLevel,
				"AWS_PROFILE=" + p.AwsProfile,
				"AWS_REGION=" + p.AwsRegion,
			},
			Args: []string{
				"awslabs.bedrock-kb-retrieval-mcp-server@latest",
			},
			InitReq: mcp.InitializeRequest{},
		},
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/aws-bedrock-kb-retrieval",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	awsMCPClient.StdioClientConf.InitReq = initRequest

	return awsMCPClient
}

type AwsAnalysisParams struct {
	FastMCPLogLevel string
	AwsProfile      string
}

func InitAwsAnalysisMCPClient(p *AwsAnalysisParams, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	awsMCPClient := &param.MCPClientConf{
		Name: NpxAwsAnalysisServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uvx",
			Env: []string{
				"FASTMCP_LOG_LEVEL=" + p.FastMCPLogLevel,
				"AWS_PROFILE=" + p.AwsProfile,
			},
			Args: []string{
				"awslabs.cost-analysis-mcp-server@latest",
			},
			InitReq: mcp.InitializeRequest{},
		},
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/aws-analysis",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	awsMCPClient.StdioClientConf.InitReq = initRequest

	return awsMCPClient
}

type AwsCDKParams struct {
	FastMCPLogLevel string
}

func InitAwsCDKMCPClient(p *AwsCDKParams, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	awsMCPClient := &param.MCPClientConf{
		Name: NpxAwsCDKServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uvx",
			Env: []string{
				"FASTMCP_LOG_LEVEL=" + p.FastMCPLogLevel,
			},
			Args: []string{
				"awslabs.cdk-mcp-server@latest",
			},
			InitReq: mcp.InitializeRequest{},
		},
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/aws-cdk",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	awsMCPClient.StdioClientConf.InitReq = initRequest

	return awsMCPClient
}

type AwsDocumentationParams struct {
	FastMCPLogLevel string
}

func InitAwsDocumentationMCPClient(p *AwsDocumentationParams, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	awsMCPClient := &param.MCPClientConf{
		Name: NpxAwsDocumentServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uvx",
			Env: []string{
				"FASTMCP_LOG_LEVEL=" + p.FastMCPLogLevel,
			},
			Args: []string{
				"awslabs.aws-documentation-mcp-server@latest",
			},
			InitReq: mcp.InitializeRequest{},
		},
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/aws-documentation",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	awsMCPClient.StdioClientConf.InitReq = initRequest

	return awsMCPClient
}

type AwsLambdaParams struct {
	AwsProfile       string
	AwsRegion        string
	FunctionPrefix   string
	FunctionList     string
	FunctionTagKey   string
	FunctionTagValue string
}

func InitAwsLambdaMCPClient(p *AwsLambdaParams, protocolVersion string, clientInfo *mcp.Implementation,
	toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
	toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

	awsMCPClient := &param.MCPClientConf{
		Name: NpxAwsLambdaServer,
		StdioClientConf: &param.StdioClientConfig{
			Command: "uvx",
			Env: []string{
				"AWS_PROFILE=" + p.AwsProfile,
				"AWS_REGION=" + p.AwsRegion,
				"FUNCTION_PREFIX=" + p.FunctionPrefix,
				"FUNCTION_LIST=" + p.FunctionList,
				"FUNCTION_TAG_KEY=" + p.FunctionTagKey,
				"FUNCTION_TAG_VALUE=" + p.FunctionTagValue,
			},
			Args: []string{
				"awslabs.lambda-mcp-server@latest",
			},
			InitReq: mcp.InitializeRequest{},
		},
		ToolsBeforeFunc: toolsBeforeFunc,
		ToolsAfterFunc:  toolsAfterFunc,
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	if protocolVersion != "" {
		initRequest.Params.ProtocolVersion = protocolVersion
	}
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcp-server/aws-lambda",
		Version: "0.1.0",
	}
	if clientInfo != nil {
		initRequest.Params.ClientInfo = *clientInfo
	}
	awsMCPClient.StdioClientConf.InitReq = initRequest

	return awsMCPClient
}
