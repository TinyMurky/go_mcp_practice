// Package tools that register mcp tool to server
package calculator

import (
	"context"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Calculator is to mimic the real tool that need an struct to carry information
type Calculator struct{}

// NewTool describe how to use Handle to LLM Module
func NewTool() mcp.Tool {

	return mcp.NewTool(
		"calculate",
		mcp.WithDescription("Perform basic arithmetic operations"),

		mcp.WithString("operation",
			mcp.Required(),
			mcp.Description(),
		),
	)

}

func Handle(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	op := request.Params.Arguments["operation"].(string)
	x := request.Params.Arguments["x"].(float64)
	y := request.Params.Arguments["y"].(float64)

	var result float64
	switch op {
	case "add":
		result = x + y
	case "subtract":
		result = x - y
	case "multiply":
		result = x * y
	case "divide":
		if y == 0 {
			return nil, errors.New("Cannot divide by zero")
		}
		result = x / y
	}

	return mcp.NewToolResultText(fmt.Sprintf("%.2f", result)), nil
}

// Register tool to server
func Register(s *server.MCPServer) {
	s.AddTool(NewTool(), Handle)
}
