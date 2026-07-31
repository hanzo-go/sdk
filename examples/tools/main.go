// tools — list the MCP tools this key can call.
//
// The Hanzo MCP surface is JSON-RPC 2.0 (HIP-0300) over a single endpoint.
// `tools/list` is the discovery call; every connector action shows up as a tool
// named <connector>_<action>.
//
// Operation: POST /v1/automations/mcp (automations_mcp), method "tools/list".
//
//	HANZO_API_KEY=sk-... go run ./examples/tools
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	hanzoai "github.com/hanzoai/go-sdk"
)

func main() {
	client := hanzoai.NewClient("")

	resp, _, err := client.MCPAPI.AutomationsMcp(context.Background()).
		AutomationsMcpRequest(hanzoai.AutomationsMcpRequest{
			Jsonrpc: "2.0",
			Id:      1,
			Method:  "tools/list",
		}).Execute()
	if err != nil {
		log.Fatalf("tools/list: %v", err)
	}

	// JSON-RPC reports call errors in the body, with HTTP 200.
	if e := resp.Error; e != nil {
		log.Fatalf("tools/list: %s", e.GetMessage())
	}

	b, err := json.MarshalIndent(resp.Result, "", "  ")
	if err != nil {
		log.Fatalf("encode result: %v", err)
	}
	fmt.Println(string(b))
}
