// Package toolutils contain function that is common in tools package
package toolutils

import "github.com/mark3labs/mcp-go/mcp"

// NewToolResultError creates a new CallToolResult with an error message.
// Any errors that originate from the tool SHOULD be reported inside the result object.
// 這個function目前還沒被放在 mcp-go 0.17.0 版本
// 先從main複製過來
func NewToolResultError(text string) *mcp.CallToolResult {
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: text,
			},
		},
		IsError: true,
	}
}
