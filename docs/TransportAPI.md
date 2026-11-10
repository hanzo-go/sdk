# \TransportAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransport**](TransportAPI.md#GetTransport) | **Get** /v1/transport | Reports every chat transport this org can send through, and whether it can send through it right now.
[**GetTransportAgent**](TransportAPI.md#GetTransportAgent) | **Get** /v1/transport/agent | Returns which agent answers the caller org&#39;s transport: the default and every channel bound to another agent.
[**GetTransportAllowlist**](TransportAPI.md#GetTransportAllowlist) | **Get** /v1/transport/allowlist | Returns the caller org&#39;s access policy for one transport: whether DMs are pairing-gated, allowlisted or open, whether group channels are open, allowlisted or disabled, the config-managed DM and group allow entries, the senders approved through PAIRING (read-only here), and the org&#39;s named access groups.
[**GetTransportInbox**](TransportAPI.md#GetTransportInbox) | **Get** /v1/transport/inbox | Returns the caller&#39;s own messages to the org&#39;s connected chat bots — those recorded on the caller&#39;s linked account — oldest first, in the portable envelope shape every transport normalises into.
[**GetTransportPairing**](TransportAPI.md#GetTransportPairing) | **Get** /v1/transport/pairing | Returns the pairing requests waiting for the caller org to approve — one per person who messaged a connected bot on a transport whose DM policy is \&quot;pairing\&quot; and who is not allowed yet.
[**PostTransportByTransportSend**](TransportAPI.md#PostTransportByTransportSend) | **Post** /v1/transport/{transport}/send | Send a message from your org&#39;s bot to one chat channel
[**PostTransportPairingApprove**](TransportAPI.md#PostTransportPairingApprove) | **Post** /v1/transport/pairing/approve | Turns one pending pairing code into a standing allow entry, so that person can DM the org&#39;s bot on that transport from now on.
[**PutTransportAgent**](TransportAPI.md#PutTransportAgent) | **Put** /v1/transport/agent | Binds agents to the caller org&#39;s transport and answers the bindings as GET would.
[**PutTransportAllowlist**](TransportAPI.md#PutTransportAllowlist) | **Put** /v1/transport/allowlist | Edits the caller org&#39;s access policy for one transport and answers the policy as GET would, so both verbs return ONE shape.



## GetTransport

> TransportChatTransports GetTransport(ctx).Execute()

Reports every chat transport this org can send through, and whether it can send through it right now.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportAPI.GetTransport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportAPI.GetTransport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransport`: TransportChatTransports
	fmt.Fprintf(os.Stdout, "Response from `TransportAPI.GetTransport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportRequest struct via the builder pattern


### Return type

[**TransportChatTransports**](TransportChatTransports.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportAgent

> TransportTransportAgents GetTransportAgent(ctx).Transport(transport).Execute()

Returns which agent answers the caller org's transport: the default and every channel bound to another agent.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	transport := "transport_example" // string | Transport is the transport: discord, github, linear, slack, teams, telegram or whatsapp. Required; an unknown value is a 404. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportAPI.GetTransportAgent(context.Background()).Transport(transport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportAPI.GetTransportAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportAgent`: TransportTransportAgents
	fmt.Fprintf(os.Stdout, "Response from `TransportAPI.GetTransportAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transport** | **string** | Transport is the transport: discord, github, linear, slack, teams, telegram or whatsapp. Required; an unknown value is a 404. | 

### Return type

[**TransportTransportAgents**](TransportTransportAgents.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportAllowlist

> TransportAllowlistView GetTransportAllowlist(ctx).Transport(transport).Execute()

Returns the caller org's access policy for one transport: whether DMs are pairing-gated, allowlisted or open, whether group channels are open, allowlisted or disabled, the config-managed DM and group allow entries, the senders approved through PAIRING (read-only here), and the org's named access groups.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	transport := "slack" // string | Transport is the transport to read: discord, github, linear, slack, teams, telegram or whatsapp. Required; an unknown value is a 404. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportAPI.GetTransportAllowlist(context.Background()).Transport(transport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportAPI.GetTransportAllowlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportAllowlist`: TransportAllowlistView
	fmt.Fprintf(os.Stdout, "Response from `TransportAPI.GetTransportAllowlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportAllowlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transport** | **string** | Transport is the transport to read: discord, github, linear, slack, teams, telegram or whatsapp. Required; an unknown value is a 404. | 

### Return type

[**TransportAllowlistView**](TransportAllowlistView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportInbox

> TransportInboxPage GetTransportInbox(ctx).Since(since).Limit(limit).Execute()

Returns the caller's own messages to the org's connected chat bots — those recorded on the caller's linked account — oldest first, in the portable envelope shape every transport normalises into.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	since := "1042" // string | Since is the exclusive cursor: only messages with a higher row id come back. Empty starts at the beginning. Must parse as an integer. (optional)
	limit := "100" // string | Limit caps how many messages come back. Empty or 0 uses the store's default page size. Must parse as an integer. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportAPI.GetTransportInbox(context.Background()).Since(since).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportAPI.GetTransportInbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportInbox`: TransportInboxPage
	fmt.Fprintf(os.Stdout, "Response from `TransportAPI.GetTransportInbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportInboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **string** | Since is the exclusive cursor: only messages with a higher row id come back. Empty starts at the beginning. Must parse as an integer. | 
 **limit** | **string** | Limit caps how many messages come back. Empty or 0 uses the store&#39;s default page size. Must parse as an integer. | 

### Return type

[**TransportInboxPage**](TransportInboxPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransportPairing

> TransportPairingQueue GetTransportPairing(ctx).Execute()

Returns the pairing requests waiting for the caller org to approve — one per person who messaged a connected bot on a transport whose DM policy is \"pairing\" and who is not allowed yet.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportAPI.GetTransportPairing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportAPI.GetTransportPairing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransportPairing`: TransportPairingQueue
	fmt.Fprintf(os.Stdout, "Response from `TransportAPI.GetTransportPairing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransportPairingRequest struct via the builder pattern


### Return type

[**TransportPairingQueue**](TransportPairingQueue.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransportByTransportSend

> PostTransportByTransportSend(ctx, transport).Execute()

Send a message from your org's bot to one chat channel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	transport := "transport_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransportAPI.PostTransportByTransportSend(context.Background(), transport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportAPI.PostTransportByTransportSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transport** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransportByTransportSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransportPairingApprove

> TransportPairingApproved PostTransportPairingApprove(ctx).TransportApprovePairingIn(transportApprovePairingIn).Execute()

Turns one pending pairing code into a standing allow entry, so that person can DM the org's bot on that transport from now on.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	transportApprovePairingIn := *openapiclient.NewTransportApprovePairingIn() // TransportApprovePairingIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportAPI.PostTransportPairingApprove(context.Background()).TransportApprovePairingIn(transportApprovePairingIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportAPI.PostTransportPairingApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTransportPairingApprove`: TransportPairingApproved
	fmt.Fprintf(os.Stdout, "Response from `TransportAPI.PostTransportPairingApprove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTransportPairingApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transportApprovePairingIn** | [**TransportApprovePairingIn**](TransportApprovePairingIn.md) |  | 

### Return type

[**TransportPairingApproved**](TransportPairingApproved.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTransportAgent

> TransportTransportAgents PutTransportAgent(ctx).TransportTransportAgentsPut(transportTransportAgentsPut).Execute()

Binds agents to the caller org's transport and answers the bindings as GET would.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	transportTransportAgentsPut := *openapiclient.NewTransportTransportAgentsPut() // TransportTransportAgentsPut | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportAPI.PutTransportAgent(context.Background()).TransportTransportAgentsPut(transportTransportAgentsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportAPI.PutTransportAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTransportAgent`: TransportTransportAgents
	fmt.Fprintf(os.Stdout, "Response from `TransportAPI.PutTransportAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutTransportAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transportTransportAgentsPut** | [**TransportTransportAgentsPut**](TransportTransportAgentsPut.md) |  | 

### Return type

[**TransportTransportAgents**](TransportTransportAgents.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTransportAllowlist

> TransportAllowlistView PutTransportAllowlist(ctx).TransportAllowlistPutIn(transportAllowlistPutIn).Execute()

Edits the caller org's access policy for one transport and answers the policy as GET would, so both verbs return ONE shape.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	transportAllowlistPutIn := *openapiclient.NewTransportAllowlistPutIn() // TransportAllowlistPutIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportAPI.PutTransportAllowlist(context.Background()).TransportAllowlistPutIn(transportAllowlistPutIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportAPI.PutTransportAllowlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTransportAllowlist`: TransportAllowlistView
	fmt.Fprintf(os.Stdout, "Response from `TransportAPI.PutTransportAllowlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutTransportAllowlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transportAllowlistPutIn** | [**TransportAllowlistPutIn**](TransportAllowlistPutIn.md) |  | 

### Return type

[**TransportAllowlistView**](TransportAllowlistView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

