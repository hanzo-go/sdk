// hello — who am I?
//
// The smallest complete round trip: authenticate with HANZO_API_KEY and read
// back the identity the API resolved it to.
//
// Operation: GET /v1/bot/whoami (bot_whoami)
//
//	HANZO_API_KEY=sk-... go run ./examples/hello
package main

import (
	"context"
	"fmt"
	"log"

	hanzoai "github.com/hanzoai/go-sdk"
)

func main() {
	client := hanzoai.NewClient("")

	me, _, err := client.AuthAPI.BotWhoami(context.Background()).Execute()
	if err != nil {
		log.Fatalf("whoami: %v", err)
	}

	fmt.Printf("id      %s\n", me.GetId())
	fmt.Printf("handle  %s\n", me.GetHandle())
	fmt.Printf("email   %s\n", me.GetEmail())
}
