# \ChannelAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChannel**](ChannelAPI.md#GetChannel) | **Get** /v1/channel | Reports every chat channel this org can send through, and whether it can send through it right now.
[**GetChannelAgent**](ChannelAPI.md#GetChannelAgent) | **Get** /v1/channel/agent | Returns which agent answers the caller org&#39;s channel: the default and every room bound to another agent.
[**GetChannelAllowlist**](ChannelAPI.md#GetChannelAllowlist) | **Get** /v1/channel/allowlist | Returns the caller org&#39;s access policy for one channel: whether DMs are pairing-gated, allowlisted or open, whether group rooms are open, allowlisted or disabled, the config-managed DM and group allow entries, the senders approved through PAIRING (read-only here), and the org&#39;s named access groups.
[**GetChannelInbox**](ChannelAPI.md#GetChannelInbox) | **Get** /v1/channel/inbox | Returns the messages people have sent to the caller org&#39;s connected chat bots, oldest first, in the portable envelope shape every transport normalises into.
[**GetChannelPairing**](ChannelAPI.md#GetChannelPairing) | **Get** /v1/channel/pairing | Returns the pairing requests waiting for the caller org to approve — one per person who messaged a connected bot on a channel whose DM policy is \&quot;pairing\&quot; and who is not allowed yet.
[**PostChannelByChannelSend**](ChannelAPI.md#PostChannelByChannelSend) | **Post** /v1/channel/{channel}/send | Send a message from your org&#39;s bot to one chat room
[**PostChannelPairingApprove**](ChannelAPI.md#PostChannelPairingApprove) | **Post** /v1/channel/pairing/approve | Turns one pending pairing code into a standing allow entry, so that person can DM the org&#39;s bot on that channel from now on.
[**PutChannelAgent**](ChannelAPI.md#PutChannelAgent) | **Put** /v1/channel/agent | Binds agents to the caller org&#39;s channel and answers the bindings as GET would.
[**PutChannelAllowlist**](ChannelAPI.md#PutChannelAllowlist) | **Put** /v1/channel/allowlist | Edits the caller org&#39;s access policy for one channel and answers the policy as GET would, so both verbs return ONE shape.



## GetChannel

> ChannelChatChannels GetChannel(ctx).Execute()

Reports every chat channel this org can send through, and whether it can send through it right now.



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
	resp, r, err := apiClient.ChannelAPI.GetChannel(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.GetChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannel`: ChannelChatChannels
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.GetChannel`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelRequest struct via the builder pattern


### Return type

[**ChannelChatChannels**](ChannelChatChannels.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelAgent

> ChannelChannelAgents GetChannelAgent(ctx).Channel(channel).Execute()

Returns which agent answers the caller org's channel: the default and every room bound to another agent.



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
	channel := "channel_example" // string | Channel is the transport: discord, github, linear, slack, teams, telegram or whatsapp. Required; an unknown value is a 404. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.GetChannelAgent(context.Background()).Channel(channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.GetChannelAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelAgent`: ChannelChannelAgents
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.GetChannelAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **string** | Channel is the transport: discord, github, linear, slack, teams, telegram or whatsapp. Required; an unknown value is a 404. | 

### Return type

[**ChannelChannelAgents**](ChannelChannelAgents.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelAllowlist

> ChannelAllowlistView GetChannelAllowlist(ctx).Channel(channel).Execute()

Returns the caller org's access policy for one channel: whether DMs are pairing-gated, allowlisted or open, whether group rooms are open, allowlisted or disabled, the config-managed DM and group allow entries, the senders approved through PAIRING (read-only here), and the org's named access groups.



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
	channel := "slack" // string | Channel is the transport to read: discord, github, linear, slack, teams, telegram or whatsapp. Required; an unknown value is a 404. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.GetChannelAllowlist(context.Background()).Channel(channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.GetChannelAllowlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelAllowlist`: ChannelAllowlistView
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.GetChannelAllowlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelAllowlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **string** | Channel is the transport to read: discord, github, linear, slack, teams, telegram or whatsapp. Required; an unknown value is a 404. | 

### Return type

[**ChannelAllowlistView**](ChannelAllowlistView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelInbox

> ChannelInboxPage GetChannelInbox(ctx).Since(since).Limit(limit).Execute()

Returns the messages people have sent to the caller org's connected chat bots, oldest first, in the portable envelope shape every transport normalises into.



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
	resp, r, err := apiClient.ChannelAPI.GetChannelInbox(context.Background()).Since(since).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.GetChannelInbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelInbox`: ChannelInboxPage
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.GetChannelInbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelInboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **string** | Since is the exclusive cursor: only messages with a higher row id come back. Empty starts at the beginning. Must parse as an integer. | 
 **limit** | **string** | Limit caps how many messages come back. Empty or 0 uses the store&#39;s default page size. Must parse as an integer. | 

### Return type

[**ChannelInboxPage**](ChannelInboxPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelPairing

> ChannelPairingQueue GetChannelPairing(ctx).Execute()

Returns the pairing requests waiting for the caller org to approve — one per person who messaged a connected bot on a channel whose DM policy is \"pairing\" and who is not allowed yet.



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
	resp, r, err := apiClient.ChannelAPI.GetChannelPairing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.GetChannelPairing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelPairing`: ChannelPairingQueue
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.GetChannelPairing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelPairingRequest struct via the builder pattern


### Return type

[**ChannelPairingQueue**](ChannelPairingQueue.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostChannelByChannelSend

> PostChannelByChannelSend(ctx, channel).Execute()

Send a message from your org's bot to one chat room



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
	channel := "channel_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChannelAPI.PostChannelByChannelSend(context.Background(), channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.PostChannelByChannelSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostChannelByChannelSendRequest struct via the builder pattern


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


## PostChannelPairingApprove

> ChannelPairingApproved PostChannelPairingApprove(ctx).ChannelApprovePairingIn(channelApprovePairingIn).Execute()

Turns one pending pairing code into a standing allow entry, so that person can DM the org's bot on that channel from now on.



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
	channelApprovePairingIn := *openapiclient.NewChannelApprovePairingIn() // ChannelApprovePairingIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.PostChannelPairingApprove(context.Background()).ChannelApprovePairingIn(channelApprovePairingIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.PostChannelPairingApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostChannelPairingApprove`: ChannelPairingApproved
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.PostChannelPairingApprove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostChannelPairingApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelApprovePairingIn** | [**ChannelApprovePairingIn**](ChannelApprovePairingIn.md) |  | 

### Return type

[**ChannelPairingApproved**](ChannelPairingApproved.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutChannelAgent

> ChannelChannelAgents PutChannelAgent(ctx).ChannelChannelAgentsPut(channelChannelAgentsPut).Execute()

Binds agents to the caller org's channel and answers the bindings as GET would.



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
	channelChannelAgentsPut := *openapiclient.NewChannelChannelAgentsPut() // ChannelChannelAgentsPut | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.PutChannelAgent(context.Background()).ChannelChannelAgentsPut(channelChannelAgentsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.PutChannelAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutChannelAgent`: ChannelChannelAgents
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.PutChannelAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutChannelAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelChannelAgentsPut** | [**ChannelChannelAgentsPut**](ChannelChannelAgentsPut.md) |  | 

### Return type

[**ChannelChannelAgents**](ChannelChannelAgents.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutChannelAllowlist

> ChannelAllowlistView PutChannelAllowlist(ctx).ChannelAllowlistPutIn(channelAllowlistPutIn).Execute()

Edits the caller org's access policy for one channel and answers the policy as GET would, so both verbs return ONE shape.



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
	channelAllowlistPutIn := *openapiclient.NewChannelAllowlistPutIn() // ChannelAllowlistPutIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.PutChannelAllowlist(context.Background()).ChannelAllowlistPutIn(channelAllowlistPutIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.PutChannelAllowlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutChannelAllowlist`: ChannelAllowlistView
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.PutChannelAllowlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutChannelAllowlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelAllowlistPutIn** | [**ChannelAllowlistPutIn**](ChannelAllowlistPutIn.md) |  | 

### Return type

[**ChannelAllowlistView**](ChannelAllowlistView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

