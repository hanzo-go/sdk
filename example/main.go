package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hanzoai/go-sdk"
	"github.com/hanzoai/go-sdk/option"
)

func main() {
	// Create a new client with API key from environment
	client := hanzoai.NewClient(
		option.WithAPIKey("your-api-key-here"),
		// Optionally override the base URL
		// option.WithBaseURL("https://api.hanzo.ai"),
	)

	// Example: List models
	_ = context.Background() // ctx would be used for API calls

	fmt.Println("Hanzo AI Go SDK initialized successfully!")
	fmt.Printf("Client configured with options: %+v\n", client.Options)

	// Example of using the Models service
	// models, err := client.Models.List(ctx, hanzoai.ModelListParams{})
	// if err != nil {
	//     log.Fatal(err)
	// }
	// fmt.Printf("Models: %+v\n", models)

	log.Println("SDK is ready to use!")
}