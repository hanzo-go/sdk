// tools — list the MCP tools this key can reach.
//
// Operation: GET /v1/tools (get_tools).
//
// This is the served tool list: every tool the caller's org can see, which is
// what an MCP tools/list answers. To CALL one, the next route is
// POST /v1/tools/call (post_tools_call).
//
//	HANZO_API_KEY=sk-... go run ./examples/tools
package main

import (
	"context"
	"fmt"
	"log"

	hanzoai "github.com/hanzoai/go-sdk"
)

func main() {
	client := hanzoai.NewClient("")

	list, _, err := client.ToolsAPI.GetTools(context.Background()).Execute()
	if err != nil {
		log.Fatalf("tools: %v", err)
	}

	tools := list.Tools
	if len(tools) == 0 {
		log.Fatal("tools: the list is empty")
	}
	fmt.Printf("%d tools\n", len(tools))
	for _, tool := range tools[:min(3, len(tools))] {
		fmt.Printf("  %-32s %s\n", tool.GetName(), tool.GetDescription())
	}
}
