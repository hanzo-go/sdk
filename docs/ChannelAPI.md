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

> ChatChannels GetChannel(ctx).Execute()

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
	// response from `GetChannel`: ChatChannels
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.GetChannel`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelRequest struct via the builder pattern


### Return type

[**ChatChannels**](ChatChannels.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelAgent

> ChannelAgents GetChannelAgent(ctx).Channel(channel).Execute()

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
	// response from `GetChannelAgent`: ChannelAgents
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

[**ChannelAgents**](ChannelAgents.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelAllowlist

> AllowlistView GetChannelAllowlist(ctx).Channel(channel).Execute()

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
	// response from `GetChannelAllowlist`: AllowlistView
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

[**AllowlistView**](AllowlistView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelInbox

> InboxPage GetChannelInbox(ctx).Since(since).Limit(limit).Execute()

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
	// response from `GetChannelInbox`: InboxPage
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

[**InboxPage**](InboxPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelPairing

> PairingQueue GetChannelPairing(ctx).Execute()

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
	// response from `GetChannelPairing`: PairingQueue
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.GetChannelPairing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelPairingRequest struct via the builder pattern


### Return type

[**PairingQueue**](PairingQueue.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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

> PairingApproved PostChannelPairingApprove(ctx).ApprovePairingIn(approvePairingIn).Execute()

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
	approvePairingIn := *openapiclient.NewApprovePairingIn() // ApprovePairingIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.PostChannelPairingApprove(context.Background()).ApprovePairingIn(approvePairingIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.PostChannelPairingApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostChannelPairingApprove`: PairingApproved
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.PostChannelPairingApprove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostChannelPairingApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **approvePairingIn** | [**ApprovePairingIn**](ApprovePairingIn.md) |  | 

### Return type

[**PairingApproved**](PairingApproved.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutChannelAgent

> ChannelAgents PutChannelAgent(ctx).ChannelAgentsPut(channelAgentsPut).Execute()

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
	channelAgentsPut := *openapiclient.NewChannelAgentsPut() // ChannelAgentsPut | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.PutChannelAgent(context.Background()).ChannelAgentsPut(channelAgentsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.PutChannelAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutChannelAgent`: ChannelAgents
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.PutChannelAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutChannelAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelAgentsPut** | [**ChannelAgentsPut**](ChannelAgentsPut.md) |  | 

### Return type

[**ChannelAgents**](ChannelAgents.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutChannelAllowlist

> AllowlistView PutChannelAllowlist(ctx).AllowlistPutIn(allowlistPutIn).Execute()

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
	allowlistPutIn := *openapiclient.NewAllowlistPutIn() // AllowlistPutIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.PutChannelAllowlist(context.Background()).AllowlistPutIn(allowlistPutIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.PutChannelAllowlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutChannelAllowlist`: AllowlistView
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.PutChannelAllowlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutChannelAllowlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowlistPutIn** | [**AllowlistPutIn**](AllowlistPutIn.md) |  | 

### Return type

[**AllowlistView**](AllowlistView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

