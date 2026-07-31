// hello — who am I?
//
// The smallest complete round trip: authenticate with HANZO_API_KEY and read
// back the identity the API resolved it to.
//
// Operation: GET /v1/bot/auth/me (bot_authMe)
//
// The route has to FAIL on a bad key — that is the whole point of the flow.
// This one answers 403 "no validated principal" with no key and with a bogus
// one, and the typed user with a real one.
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

	me, _, err := client.AuthAPI.BotAuthMe(context.Background()).Execute()
	if err != nil {
		log.Fatalf("whoami: %v", err)
	}

	fmt.Printf("id      %s\n", me.GetId())
	fmt.Printf("handle  %s\n", me.GetHandle())
	fmt.Printf("name    %s\n", me.GetDisplayName())
}
