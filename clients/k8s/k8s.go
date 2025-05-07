package k8s

import (
    "github.com/mark3labs/mcp-go/mcp"
    "github.com/yincongcyincong/mcp-client-go/clients/param"
    "strconv"
)

const (
    DockerK8sMcpServer    = "docker-k8s-mcp-server"
    DockerAwsK8sMcpServer = "docker-aws-k8s-mcp-server"
    DockerGcloudMcpServer = "docker-gcloud-k8s-mcp-server"
    DockerAzureMcpServer  = "docker-azure-k8s-mcp-server"
)

type K8sParam struct {
    KubConfPath   string // /Users/YOUR_USER_NAME/.kube:/home/appuser/.kube:ro
    K8sContext    string
    K8sNameSpace  string
    K8sMcpTimeout int
}

func InitDockerK8sMCPClient(p *K8sParam, protocolVersion string, clientInfo *mcp.Implementation,
    toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
    toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

    k8sMCPClient := &param.MCPClientConf{
        Name: DockerK8sMcpServer,
        StdioClientConf: &param.StdioClientConfig{
            Command: "docker",
            Env:     []string{},
            Args: []string{
                "run", "-i", "--rm",
                "-v", p.KubConfPath,
                "-e", "K8S_CONTEXT=" + p.K8sContext,
                "-e", "K8S_NAMESPACE=" + p.K8sNameSpace,
                "-e", "K8S_MCP_TIMEOUT=" + strconv.Itoa(p.K8sMcpTimeout),
                "ghcr.io/alexei-led/k8s-mcp-server:latest",
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
        Name:    "mcp-server/k8s",
        Version: "0.1.0",
    }
    if clientInfo != nil {
        initRequest.Params.ClientInfo = *clientInfo
    }
    k8sMCPClient.StdioClientConf.InitReq = initRequest

    return k8sMCPClient
}

type AwsK8sParam struct {
    KubConfPath string // /Users/YOUR_USER_NAME/.kube:/home/appuser/.kube:ro
    AwsConfPath string // /Users/YOUR_USER_NAME/.aws:/home/appuser/.aws:ro
    AwsProfile  string
    AwsRegion   string
}

func InitDockerAwsK8sMCPClient(p *AwsK8sParam, protocolVersion string, clientInfo *mcp.Implementation,
    toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
    toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

    k8sMCPClient := &param.MCPClientConf{
        Name: DockerAwsK8sMcpServer,
        StdioClientConf: &param.StdioClientConfig{
            Command: "docker",
            Env:     []string{},
            Args: []string{
                "run", "-i", "--rm",
                "-v", p.KubConfPath,
                "-v", p.AwsConfPath,
                "-e", "AWS_PROFILE=" + p.AwsProfile,
                "-e", "AWS_REGION=" + p.AwsRegion,
                "ghcr.io/alexei-led/k8s-mcp-server:latest",
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
        Name:    "mcp-server/aws-k8s",
        Version: "0.1.0",
    }
    if clientInfo != nil {
        initRequest.Params.ClientInfo = *clientInfo
    }
    k8sMCPClient.StdioClientConf.InitReq = initRequest

    return k8sMCPClient
}

type GcloudK8sParam struct {
    KubConfPath         string // /Users/YOUR_USER_NAME/.kube:/home/appuser/.kube:ro
    GcloudConfPath      string // /Users/YOUR_USER_NAME/.config/gcloud:/home/appuser/.config/gcloud:ro
    CloudSdkCoreProject string
    CloudComputeRegion  string
}

func InitDockerGcloudK8sMCPClient(p *GcloudK8sParam, protocolVersion string, clientInfo *mcp.Implementation,
    toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
    toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

    k8sMCPClient := &param.MCPClientConf{
        Name: DockerGcloudMcpServer,
        StdioClientConf: &param.StdioClientConfig{
            Command: "docker",
            Env:     []string{},
            Args: []string{
                "run", "-i", "--rm",
                "-v", p.KubConfPath,
                "-v", p.GcloudConfPath,
                "-e", "CLOUDSDK_CORE_PROJECT=" + p.CloudSdkCoreProject,
                "-e", "CLOUDSDK_COMPUTE_REGION=" + p.CloudComputeRegion,
                "ghcr.io/alexei-led/k8s-mcp-server:latest",
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
        Name:    "mcp-server/gcloud-k8s",
        Version: "0.1.0",
    }
    if clientInfo != nil {
        initRequest.Params.ClientInfo = *clientInfo
    }
    k8sMCPClient.StdioClientConf.InitReq = initRequest

    return k8sMCPClient
}

type AzureK8sParam struct {
    KubConfPath       string // /Users/YOUR_USER_NAME/.kube:/home/appuser/.kube:ro
    AzureConfPath     string // /Users/YOUR_USER_NAME/.azure:/home/appuser/.azure:ro
    AzureSubscription string
}

func InitDockerAzureK8sMCPClient(p *AzureK8sParam, protocolVersion string, clientInfo *mcp.Implementation,
    toolsBeforeFunc map[string]func(req *mcp.CallToolRequest) error,
    toolsAfterFunc map[string]func(req *mcp.CallToolResult) (string, error)) *param.MCPClientConf {

    k8sMCPClient := &param.MCPClientConf{
        Name: DockerAzureMcpServer,
        StdioClientConf: &param.StdioClientConfig{
            Command: "docker",
            Env:     []string{},
            Args: []string{
                "run", "-i", "--rm",
                "-v", p.KubConfPath,
                "-v", p.AzureConfPath,
                "-e", "AZURE_SUBSCRIPTION=" + p.AzureSubscription,
                "ghcr.io/alexei-led/k8s-mcp-server:latest",
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
        Name:    "mcp-server/azure-k8s",
        Version: "0.1.0",
    }
    if clientInfo != nil {
        initRequest.Params.ClientInfo = *clientInfo
    }
    k8sMCPClient.StdioClientConf.InitReq = initRequest

    return k8sMCPClient
}
