// Package calculator that register calculator to server
package calculator

import (
	"context"
	"fmt"
	"log"
	toolutils "mcpMathPractice/tools/utils"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// NewTool describe how to use Handle to LLM Module
func NewTool() mcp.Tool {

	return mcp.NewTool(
		"calculate",
		mcp.WithDescription("Perform basic arithmetic operations"),

		mcp.WithString("operation",
			mcp.Required(),
			mcp.Description("The operation to perform (add, subtract, multiply, divide)"),
			mcp.Enum("add", "substract", "multiply", "divide"),
		),

		mcp.WithNumber("x",
			mcp.Required(),
			mcp.Description("First number that the operation will take"),
		),

		mcp.WithNumber("y",
			mcp.Required(),
			mcp.Description("Second number that the operation will take"),
		),
	)

}

// Handle is the real function that mcp server will execute for  LLM
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
			// 下面這個funn
			return toolutils.NewToolResultError("Cannot divide by zero"), nil
		}
		result = x / y
	default:
		return toolutils.NewToolResultError("This Method not allow"), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("%.2f", result)), nil
}

// Register tool to server
func Register(s *server.MCPServer) {
	s.AddTool(NewTool(), Handle)
	log.Printf("Registered tool: calculate")
}
