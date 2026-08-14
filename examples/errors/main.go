// errors — read a refusal.
//
// Not one of the six cross-SDK flows. This is Go's own contract, and it is the
// part of a generated client a caller learns once: every Execute returns
// (value, *http.Response, error), and on a refusal all three carry something.
//
//   - The status code is on the RESPONSE, which comes back ALONGSIDE the error
//     rather than instead of it. The error's own Error() is the status line as
//     text, so there is no code to parse back out of it.
//   - The error is a *hanzoai.GenericOpenAPIError, and Body() is what the API
//     actually said. The client re-wraps the response body before returning, so
//     reading one does not consume the other.
//   - The response is nil when the request never went out — a bad base URL, a
//     transport failure, a cancelled context — and that is the case worth
//     separating, because retrying it means something different.
//
// This calls GET /v1/keys with a key the API refuses, so it walks the failure
// path rather than the happy one. It needs no credential of its own.
//
//	go run ./examples/errors
package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	hanzoai "github.com/hanzoai/go-sdk"
)

func main() {
	client := hanzoai.NewClient("sk-this-key-does-not-exist")

	listing, resp, err := client.KeysAPI.GetKeys(context.Background()).Execute()
	if err == nil {
		fmt.Printf("accepted  %d key(s)\n", len(listing.Keys))
		return
	}
	if resp == nil {
		// The API never answered, so there is no status and no body to read.
		log.Fatalf("unsent: %v", err)
	}

	fmt.Printf("status    %s\n", resp.Status)

	var apiErr *hanzoai.GenericOpenAPIError
	if errors.As(err, &apiErr) {
		fmt.Printf("refused   %s\n", apiErr.Body())
		return
	}
	fmt.Printf("failed    %v\n", err)
}
