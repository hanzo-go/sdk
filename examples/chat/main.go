// chat — one completion.
//
// Operation: POST /v1/chat/completions (ai_createChatCompletion), the
// OpenAI-compatible surface.
//
//	HANZO_API_KEY=sk-... go run ./examples/chat
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	hanzoai "github.com/hanzoai/go-sdk"
)

func main() {
	model := os.Getenv("HANZO_MODEL")
	if model == "" {
		model = "zen-1"
	}

	client := hanzoai.NewClient("")

	req := hanzoai.AiChatCompletionRequest{
		Model: model,
		Messages: []hanzoai.AiChatMessage{
			{Role: "user", Content: "In one sentence: what is Hanzo?"},
		},
	}

	resp, _, err := client.OpenAICompatibleAPI.AiCreateChatCompletion(context.Background()).
		AiChatCompletionRequest(req).Execute()
	if err != nil {
		log.Fatalf("chat completion: %v", err)
	}

	for _, choice := range resp.Choices {
		msg := choice.GetMessage()
		// content is a string for plain replies and an array of parts for
		// multimodal ones, so the spec leaves it open.
		switch content := msg.Content.(type) {
		case string:
			fmt.Println(content)
		default:
			b, err := json.Marshal(content)
			if err != nil {
				log.Fatalf("encode content: %v", err)
			}
			fmt.Println(string(b))
		}
	}
	if u := resp.Usage; u != nil {
		fmt.Printf("\ntokens: %d prompt + %d completion\n", u.GetPromptTokens(), u.GetCompletionTokens())
	}
}
