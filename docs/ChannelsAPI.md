# \ChannelsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChannels**](ChannelsAPI.md#GetChannels) | **Get** /v1/channels | Reports every chat channel this org can send through, and whether it can send through it right now.
[**GetChannelsAllowlist**](ChannelsAPI.md#GetChannelsAllowlist) | **Get** /v1/channels/allowlist | Returns the caller org&#39;s access policy for one channel: whether DMs are pairing-gated, allowlisted or open, whether group rooms are open, allowlisted or disabled, the config-managed DM and group allow entries, the senders approved through PAIRING (read-only here), and the org&#39;s named access groups.
[**GetChannelsInbox**](ChannelsAPI.md#GetChannelsInbox) | **Get** /v1/channels/inbox | Returns the messages people have sent to the caller org&#39;s connected chat bots, oldest first, in the portable envelope shape every transport normalises into.
[**GetChannelsPairing**](ChannelsAPI.md#GetChannelsPairing) | **Get** /v1/channels/pairing | Returns the pairing requests waiting for the caller org to approve — one per person who messaged a connected bot on a channel whose DM policy is \&quot;pairing\&quot; and who is not allowed yet.
[**PostChannelsByChannelSend**](ChannelsAPI.md#PostChannelsByChannelSend) | **Post** /v1/channels/{channel}/send | Send a message from your org&#39;s bot to one chat room
[**PostChannelsPairingApprove**](ChannelsAPI.md#PostChannelsPairingApprove) | **Post** /v1/channels/pairing/approve | Turns one pending pairing code into a standing allow entry, so that person can DM the org&#39;s bot on that channel from now on.
[**PutChannelsAllowlist**](ChannelsAPI.md#PutChannelsAllowlist) | **Put** /v1/channels/allowlist | Edits the caller org&#39;s access policy for one channel and answers the policy as GET would, so both verbs return ONE shape.



## GetChannels

> ChatChannels GetChannels(ctx).Execute()

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
	resp, r, err := apiClient.ChannelsAPI.GetChannels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannels`: ChatChannels
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetChannels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelsRequest struct via the builder pattern


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


## GetChannelsAllowlist

> AllowlistView GetChannelsAllowlist(ctx).Channel(channel).Execute()

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
	channel := "slack" // string | Channel is the transport to read: discord, slack, teams, telegram or whatsapp. Required; an unknown value is a 404. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.GetChannelsAllowlist(context.Background()).Channel(channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetChannelsAllowlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelsAllowlist`: AllowlistView
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetChannelsAllowlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelsAllowlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **string** | Channel is the transport to read: discord, slack, teams, telegram or whatsapp. Required; an unknown value is a 404. | 

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


## GetChannelsInbox

> InboxPage GetChannelsInbox(ctx).Since(since).Limit(limit).Execute()

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
	resp, r, err := apiClient.ChannelsAPI.GetChannelsInbox(context.Background()).Since(since).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetChannelsInbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelsInbox`: InboxPage
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetChannelsInbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelsInboxRequest struct via the builder pattern


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


## GetChannelsPairing

> PairingQueue GetChannelsPairing(ctx).Execute()

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
	resp, r, err := apiClient.ChannelsAPI.GetChannelsPairing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetChannelsPairing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelsPairing`: PairingQueue
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetChannelsPairing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelsPairingRequest struct via the builder pattern


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


## PostChannelsByChannelSend

> PostChannelsByChannelSend(ctx, channel).Execute()

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
	r, err := apiClient.ChannelsAPI.PostChannelsByChannelSend(context.Background(), channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.PostChannelsByChannelSend``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostChannelsByChannelSendRequest struct via the builder pattern


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


## PostChannelsPairingApprove

> PairingApproved PostChannelsPairingApprove(ctx).ApprovePairingIn(approvePairingIn).Execute()

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
	resp, r, err := apiClient.ChannelsAPI.PostChannelsPairingApprove(context.Background()).ApprovePairingIn(approvePairingIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.PostChannelsPairingApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostChannelsPairingApprove`: PairingApproved
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.PostChannelsPairingApprove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostChannelsPairingApproveRequest struct via the builder pattern


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


## PutChannelsAllowlist

> AllowlistView PutChannelsAllowlist(ctx).AllowlistPutIn(allowlistPutIn).Execute()

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
	resp, r, err := apiClient.ChannelsAPI.PutChannelsAllowlist(context.Background()).AllowlistPutIn(allowlistPutIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.PutChannelsAllowlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutChannelsAllowlist`: AllowlistView
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.PutChannelsAllowlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutChannelsAllowlistRequest struct via the builder pattern


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

