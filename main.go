// Package main will start server base on SSE or stdio
package main

import (
	"github.com/mark3labs/mcp-go/server"
)

func main() {

}

func run(transport string, addr string) error {}

func newServer() *server.MCPServer {
	s := server.NewMCPServer(
		"mcp-practice",
		"0.0.1",
	)
	return s
}
