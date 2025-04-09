package github

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/yincongcyincong/mcp-client-go/utils"
)

var (
	githubMCPClient *client.StdioMCPClient
)

func InitGithubMCPClient(ctx context.Context, githubAccessToken string) error {
	var err error
	githubMCPClient, err = client.NewStdioMCPClient(
		"npx",
		[]string{
			"GITHUB_PERSONAL_ACCESS_TOKEN=" + githubAccessToken,
		},
		"-y",
		"@modelcontextprotocol/server-github",
	)

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "modelcontextprotocol/server-github",
		Version: "0.2.0",
	}

	initResult, err := githubMCPClient.Initialize(ctx, initRequest)
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

func GetGithubMCPClient() *client.StdioMCPClient {
	return githubMCPClient
}

func GetGithubTools(ctx context.Context, cursor mcp.Cursor) ([]mcp.Tool, error) {
	toolsRequest := mcp.ListToolsRequest{}
	toolsRequest.Params.Cursor = cursor
	tools, err := githubMCPClient.ListTools(ctx, toolsRequest)
	if err != nil {
		return nil, err
	}
	return tools.Tools, nil
}

func CreateOrUpdateFile(ctx context.Context, owner, repo, path, content, message, branch, sha string) (string, error) {

	listDirRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	listDirRequest.Params.Name = "create_or_update_file"
	listDirRequest.Params.Arguments = map[string]interface{}{
		"owner":   owner,
		"repo":    repo,
		"path":    path,
		"message": message,
		"branch":  branch,
		"content": content,
		"sha":     sha,
	}

	result, err := githubMCPClient.CallTool(ctx, listDirRequest)
	if err != nil {
		return "", err
	}

	return utils.ReturnString(result), nil
}

func PushFiles(ctx context.Context, owner, repo, branch string, files []string, message string) (string, error) {
	// Validate required parameters
	if owner == "" || repo == "" || branch == "" || message == "" {
		return "", fmt.Errorf("owner, repo, branch and message cannot be empty")
	}

	if len(files) == 0 {
		return "", fmt.Errorf("files array cannot be empty")
	}

	pushRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}

	pushRequest.Params.Name = "push_files"
	pushRequest.Params.Arguments = map[string]interface{}{
		"owner":   owner,
		"repo":    repo,
		"branch":  branch,
		"files":   files,
		"message": message,
	}

	result, err := githubMCPClient.CallTool(ctx, pushRequest)
	if err != nil {
		return "", fmt.Errorf("failed to push files: %w", err)
	}

	return utils.ReturnString(result), nil
}

func SearchRepositories(ctx context.Context, query string, page, perPage int) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "search_repositories"
	req.Params.Arguments = map[string]interface{}{
		"query":   query,
		"page":    page,
		"perPage": perPage,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func CreateRepository(ctx context.Context, name, description string, private, autoInit bool) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "create_repository"
	req.Params.Arguments = map[string]interface{}{
		"name":        name,
		"description": description,
		"private":     private,
		"autoInit":    autoInit,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func GetFileContents(ctx context.Context, owner, repo, path, branch string) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "get_file_contents"
	req.Params.Arguments = map[string]interface{}{
		"owner":  owner,
		"repo":   repo,
		"path":   path,
		"branch": branch,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func CreateIssue(ctx context.Context, owner, repo, title, body string, assignees, labels []string, milestone int) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "create_issue"
	req.Params.Arguments = map[string]interface{}{
		"owner":     owner,
		"repo":      repo,
		"title":     title,
		"body":      body,
		"assignees": assignees,
		"labels":    labels,
		"milestone": milestone,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func CreatePullRequest(ctx context.Context, owner, repo, title, body, head, base string, draft, maintainerCanModify bool) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "create_pull_request"
	req.Params.Arguments = map[string]interface{}{
		"owner":                 owner,
		"repo":                  repo,
		"title":                 title,
		"body":                  body,
		"head":                  head,
		"base":                  base,
		"draft":                 draft,
		"maintainer_can_modify": maintainerCanModify,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func ForkRepository(ctx context.Context, owner, repo, organization string) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "fork_repository"
	req.Params.Arguments = map[string]interface{}{
		"owner":        owner,
		"repo":         repo,
		"organization": organization,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func CreateBranch(ctx context.Context, owner, repo, branch, fromBranch string) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "create_branch"
	req.Params.Arguments = map[string]interface{}{
		"owner":       owner,
		"repo":        repo,
		"branch":      branch,
		"from_branch": fromBranch,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func ListIssues(ctx context.Context, owner, repo, state string, labels []string, sort, direction, since string, page, perPage int) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "list_issues"
	req.Params.Arguments = map[string]interface{}{
		"owner":     owner,
		"repo":      repo,
		"state":     state,
		"labels":    labels,
		"sort":      sort,
		"direction": direction,
		"since":     since,
		"page":      page,
		"per_page":  perPage,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func UpdateIssue(ctx context.Context, owner, repo string, issueNumber int, title, body, state string, labels, assignees []string, milestone int) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "update_issue"
	req.Params.Arguments = map[string]interface{}{
		"owner":        owner,
		"repo":         repo,
		"issue_number": issueNumber,
		"title":        title,
		"body":         body,
		"state":        state,
		"labels":       labels,
		"assignees":    assignees,
		"milestone":    milestone,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func AddIssueComment(ctx context.Context, owner, repo string, issueNumber int, body string) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "add_issue_comment"
	req.Params.Arguments = map[string]interface{}{
		"owner":        owner,
		"repo":         repo,
		"issue_number": issueNumber,
		"body":         body,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func SearchCode(ctx context.Context, q, sort, order string, perPage, page int) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "search_code"
	req.Params.Arguments = map[string]interface{}{
		"q":        q,
		"sort":     sort,
		"order":    order,
		"per_page": perPage,
		"page":     page,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func SearchIssues(ctx context.Context, q, sort, order string, perPage, page int) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "search_issues"
	req.Params.Arguments = map[string]interface{}{
		"q":        q,
		"sort":     sort,
		"order":    order,
		"per_page": perPage,
		"page":     page,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func SearchUsers(ctx context.Context, q, sort, order string, perPage, page int) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "search_users"
	req.Params.Arguments = map[string]interface{}{
		"q":        q,
		"sort":     sort,
		"order":    order,
		"per_page": perPage,
		"page":     page,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func ListCommits(ctx context.Context, owner, repo, page, perPage, sha string) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "list_commits"
	req.Params.Arguments = map[string]interface{}{
		"owner":    owner,
		"repo":     repo,
		"page":     page,
		"per_page": perPage,
		"sha":      sha,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func GetIssue(ctx context.Context, owner, repo string, issueNumber int) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "get_issue"
	req.Params.Arguments = map[string]interface{}{
		"owner":        owner,
		"repo":         repo,
		"issue_number": issueNumber,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func GetPullRequest(ctx context.Context, owner, repo string, pullNumber int) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "get_pull_request"
	req.Params.Arguments = map[string]interface{}{
		"owner":       owner,
		"repo":        repo,
		"pull_number": pullNumber,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func ListPullRequests(ctx context.Context, owner, repo, state, head, base, sort, direction string, perPage, page int) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "list_pull_requests"
	req.Params.Arguments = map[string]interface{}{
		"owner":     owner,
		"repo":      repo,
		"state":     state,
		"head":      head,
		"base":      base,
		"sort":      sort,
		"direction": direction,
		"per_page":  perPage,
		"page":      page,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func CreatePullRequestReview(ctx context.Context, owner, repo string, pullNumber int, body, event, commitId string, comments []map[string]interface{}) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "create_pull_request_review"
	req.Params.Arguments = map[string]interface{}{
		"owner":       owner,
		"repo":        repo,
		"pull_number": pullNumber,
		"body":        body,
		"event":       event,
		"commit_id":   commitId,
		"comments":    comments,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func MergePullRequest(ctx context.Context, owner, repo string, pullNumber int, commitTitle, commitMessage, mergeMethod string) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "merge_pull_request"
	req.Params.Arguments = map[string]interface{}{
		"owner":          owner,
		"repo":           repo,
		"pull_number":    pullNumber,
		"commit_title":   commitTitle,
		"commit_message": commitMessage,
		"merge_method":   mergeMethod,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func GetPullRequestFiles(ctx context.Context, owner, repo string, pullNumber int) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "get_pull_request_files"
	req.Params.Arguments = map[string]interface{}{
		"owner":       owner,
		"repo":        repo,
		"pull_number": pullNumber,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func GetPullRequestStatus(ctx context.Context, owner, repo string, pullNumber int) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "get_pull_request_status"
	req.Params.Arguments = map[string]interface{}{
		"owner":       owner,
		"repo":        repo,
		"pull_number": pullNumber,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func UpdatePullRequestBranch(ctx context.Context, owner, repo string, pullNumber int, expectedHeadSha string) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "update_pull_request_branch"
	req.Params.Arguments = map[string]interface{}{
		"owner":             owner,
		"repo":              repo,
		"pull_number":       pullNumber,
		"expected_head_sha": expectedHeadSha,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func GetPullRequestComments(ctx context.Context, owner, repo string, pullNumber int) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "get_pull_request_comments"
	req.Params.Arguments = map[string]interface{}{
		"owner":       owner,
		"repo":        repo,
		"pull_number": pullNumber,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}

func GetPullRequestReviews(ctx context.Context, owner, repo string, pullNumber int) (string, error) {
	req := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	req.Params.Name = "get_pull_request_reviews"
	req.Params.Arguments = map[string]interface{}{
		"owner":       owner,
		"repo":        repo,
		"pull_number": pullNumber,
	}

	result, err := githubMCPClient.CallTool(ctx, req)
	if err != nil {
		return "", err
	}
	return utils.ReturnString(result), nil
}
