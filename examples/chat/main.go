// chat — one completion.
//
// Operation: POST /v1/chat/completions (post_chat_completions), the gateway's
// own inference route.
//
// The document declares what comes back but not what goes in, so the generated
// method still takes no prompt and this prints whatever the route answers with.
// The one thing this example must not do is invent the missing half. A
// hand-rolled request body inside a generated client is an opinion about the
// API rather than a projection of it, and that second authority is what these
// SDKs exist to remove.
//
// It used to read the raw body for the same reason. Now that the Out type is
// declared it reads choices[0].message.content, which arrived by regeneration
// with no decision taken in this file. The request body follows the same way.
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
	"log"

	hanzoai "github.com/hanzoai/go-sdk/v8"
)

func main() {
	client := hanzoai.NewClient("")

	completion, resp, err := client.AiAPI.PostChatCompletions(context.Background()).Execute()
	if err != nil {
		log.Fatalf("chat: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("completion  %s  %s\n", resp.Status, completion.GetModel())
	for _, c := range completion.Choices {
		fmt.Println(c.Message.GetContent())
	}
}
