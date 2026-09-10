// tools — list the MCP tools this credential can reach.
//
// Operation: GET /v1/tool (get_tool).
//
// This is the served tool list: every tool the caller's org can see, which is
// what an MCP tools/list answers. To CALL one, the next route is
// POST /v1/tool/call (post_tool_call).
//
//	HANZO_CLIENT_ID=... HANZO_CLIENT_SECRET=... go run ./examples/tools
package main

import (
	"context"
	"fmt"
	"log"

	hanzoai "github.com/hanzoai/go-sdk/v8"
)

func main() {
	client := hanzoai.New(hanzoai.Options{})

	list, _, err := client.ToolAPI.GetTool(context.Background()).Execute()
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
