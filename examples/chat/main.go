// chat — one completion.
//
// Operation: POST /v1/chat/completions (post_chat_completions), the gateway's
// own inference route.
//
// The document states that address and not its shape — no request body, no
// responses, one of the 716 operations cloud publishes as an address alone — so
// the generated method takes no prompt and hands back the raw *http.Response.
// The one thing this example must not do is invent the missing shape. A
// hand-rolled request body inside a generated client is an opinion about the
// API rather than a projection of it, and that second authority is what these
// SDKs exist to remove. So it calls the operation the document declares and
// prints what the route answered.
//
// When cloud's handler declares its In and Out types this becomes
// PostChatCompletions(ctx).<Request>(...) printing choices[0].message.content —
// a regeneration away, with no decision left in this file.
//
// Non-streaming. Streaming is SSE, a different transport that a generated
// client hands back as an opaque body, so demonstrating it here would teach the
// wrong shape.
//
//	HANZO_API_KEY=sk-... go run ./examples/chat
package main

import (
	"context"
	"fmt"
	"io"
	"log"

	hanzoai "github.com/hanzoai/go-sdk"
)

func main() {
	client := hanzoai.NewClient("")

	resp, err := client.AiAPI.PostChatCompletions(context.Background()).Execute()
	if err != nil {
		log.Fatalf("chat: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read body: %v", err)
	}
	fmt.Printf("completion  %s\n%s\n", resp.Status, body)
}
