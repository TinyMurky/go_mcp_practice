// Package tools that register mcp tool to server
package tools

import (
	"mcpMathPractice/tools/calculator"

	"github.com/mark3labs/mcp-go/server"
)

// RegisterAllTools will register all tools that was created in tools package
func RegisterAllTools(s *server.MCPServer) {
	calculator.Register(s)
}
