// Package main will start server base on SSE or stdio
//
// Ref:
//
// 1. https://github.com/grafana/mcp-grafana/blob/9287a51cdcdceb84768e6c57f598f9a525aee427/cmd/mcp-grafana/main.go#L5
//
// 2. https://ganhua.wang/mcp-go#heading-hello-world
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mcpMathPractice/tools"
	"os"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// 用來決定要用stdio還是sse來和LLM連接
	// 預設 stdio
	var transport string
	flag.StringVar(&transport, "t", "stdio", "Transport type (stdio or sse)")

	flag.StringVar(
		&transport,
		"transport",
		"stdio",
		"Transport type (stdio or sse)",
	)
	addr := flag.String("sse-address", "localhost:8000", "The host and port to start the sse server on")
	flag.Parse()

	if err := run(transport, *addr); err != nil {
		panic(err)
	}
}

func newServer() *server.MCPServer {
	s := server.NewMCPServer(
		"mcp-practice",
		"0.0.1",
	)

	tools.RegisterAllTools(s)
	return s
}

func run(transport string, addr string) error {
	s := newServer()

	// MCP 可以選擇 透過 os.stdio, os.stdout 和LLM連接 ("stdio")
	// 或是 透過http 連接 ("sse")
	// Ref: https://ganhua.wang/mcp-go#heading-mcp-lifecycle
	switch transport {
	case "stdio":
		srv := server.NewStdioServer(s)

		// 這邊os.Stdin 和 os.Stdout 其實只要吃 io.Reader和io.Writer就好了
		return srv.Listen(context.Background(), os.Stdin, os.Stdout)
	case "sse":
		srv := server.NewSSEServer(s)

		log.Printf("SSE server listening on %s", addr)
		if err := srv.Start(addr); err != nil {
			return fmt.Errorf("Server error: %v", err)
		}
	default:
		return fmt.Errorf(
			"Invalid transport type: %s. Must be 'stdio' or 'sse'",
			transport,
		)
	}

	return nil
}
