// chat — one completion.
//
// Operation: POST /v1/chat/completions (post_chat_completions), the gateway's
// own inference route. The document declares both halves now, so the prompt
// goes in as OpenaiChatCompletionRequest and choices[0].message.content comes
// back typed, both by regeneration with no decision taken in this file.
//
// Non-streaming. Streaming is SSE, a different transport that a generated
// client hands back as an opaque body, so demonstrating it here would teach the
// wrong shape.
//
//	HANZO_CLIENT_ID=... HANZO_CLIENT_SECRET=... go run ./examples/chat
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	hanzoai "github.com/hanzoai/go-sdk/v8"
)

func main() {
	model := os.Getenv("HANZO_MODEL")
	if model == "" {
		model = "enso"
	}

	client := hanzoai.New(hanzoai.Options{})

	completion, resp, err := client.AiAPI.PostChatCompletions(context.Background()).
		OpenaiChatCompletionRequest(hanzoai.OpenaiChatCompletionRequest{
			Model: &model,
			Messages: []hanzoai.OpenaiChatCompletionMessage{
				{Role: hanzoai.PtrString("user"), Content: hanzoai.PtrString("Say hello in one sentence.")},
			},
		}).Execute()
	if err != nil {
		log.Fatalf("chat: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("completion  %s  %s\n", resp.Status, completion.GetModel())
	for _, c := range completion.Choices {
		fmt.Println(c.Message.GetContent())
	}
}
