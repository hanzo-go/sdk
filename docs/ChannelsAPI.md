# \ChannelsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Channels**](ChannelsAPI.md#CloudGetV1Channels) | **Get** /v1/channels | Returns every chat transport channels can talk to — Discord, Slack, Teams and Telegram — with the caller org&#39;s own facts on each: whether it is connected and to which account, what the transport supports, the org&#39;s DM and group access policies, and how many pairing requests are pending approval.
[**CloudGetV1ChannelsAllowlist**](ChannelsAPI.md#CloudGetV1ChannelsAllowlist) | **Get** /v1/channels/allowlist | Returns the caller org&#39;s access policy for one channel: whether DMs are pairing-gated, allowlisted or open, whether group rooms are open, allowlisted or disabled, the config-managed DM and group allow entries, the senders approved through PAIRING (read-only here), and the org&#39;s named access groups.
[**CloudGetV1ChannelsInbox**](ChannelsAPI.md#CloudGetV1ChannelsInbox) | **Get** /v1/channels/inbox | Returns the messages people have sent to the caller org&#39;s connected chat bots, oldest first, in the portable envelope shape every transport normalises into.
[**CloudGetV1ChannelsPairing**](ChannelsAPI.md#CloudGetV1ChannelsPairing) | **Get** /v1/channels/pairing | Returns the pairing requests waiting for the caller org to approve — one per person who messaged a connected bot on a channel whose DM policy is \&quot;pairing\&quot; and who is not allowed yet.
[**CloudPostV1ChannelsByChannelSend**](ChannelsAPI.md#CloudPostV1ChannelsByChannelSend) | **Post** /v1/channels/{channel}/send | 
[**CloudPostV1ChannelsPairingApprove**](ChannelsAPI.md#CloudPostV1ChannelsPairingApprove) | **Post** /v1/channels/pairing/approve | Turns one pending pairing code into a standing allow entry, so that person can DM the org&#39;s bot on that channel from now on.
[**CloudPutV1ChannelsAllowlist**](ChannelsAPI.md#CloudPutV1ChannelsAllowlist) | **Put** /v1/channels/allowlist | Edits the caller org&#39;s access policy for one channel and answers the policy as GET would, so both verbs return ONE shape.



## CloudGetV1Channels

> CloudChatChannels CloudGetV1Channels(ctx).Execute()

Returns every chat transport channels can talk to — Discord, Slack, Teams and Telegram — with the caller org's own facts on each: whether it is connected and to which account, what the transport supports, the org's DM and group access policies, and how many pairing requests are pending approval.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.CloudGetV1Channels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.CloudGetV1Channels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Channels`: CloudChatChannels
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.CloudGetV1Channels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ChannelsRequest struct via the builder pattern


### Return type

[**CloudChatChannels**](CloudChatChannels.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ChannelsAllowlist

> CloudAllowlistView CloudGetV1ChannelsAllowlist(ctx).Channel(channel).Execute()

Returns the caller org's access policy for one channel: whether DMs are pairing-gated, allowlisted or open, whether group rooms are open, allowlisted or disabled, the config-managed DM and group allow entries, the senders approved through PAIRING (read-only here), and the org's named access groups.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	channel := "slack" // string | Channel is the transport to read: discord, slack, teams or telegram. Required; an unknown value is a 404. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.CloudGetV1ChannelsAllowlist(context.Background()).Channel(channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.CloudGetV1ChannelsAllowlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ChannelsAllowlist`: CloudAllowlistView
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.CloudGetV1ChannelsAllowlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ChannelsAllowlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **string** | Channel is the transport to read: discord, slack, teams or telegram. Required; an unknown value is a 404. | 

### Return type

[**CloudAllowlistView**](CloudAllowlistView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ChannelsInbox

> CloudInboxPage CloudGetV1ChannelsInbox(ctx).Since(since).Limit(limit).Execute()

Returns the messages people have sent to the caller org's connected chat bots, oldest first, in the portable envelope shape every transport normalises into.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	since := "1042" // string | Since is the exclusive cursor: only messages with a higher row id come back. Empty starts at the beginning. Must parse as an integer. (optional)
	limit := "100" // string | Limit caps how many messages come back. Empty or 0 uses the store's default page size. Must parse as an integer. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.CloudGetV1ChannelsInbox(context.Background()).Since(since).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.CloudGetV1ChannelsInbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ChannelsInbox`: CloudInboxPage
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.CloudGetV1ChannelsInbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ChannelsInboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **string** | Since is the exclusive cursor: only messages with a higher row id come back. Empty starts at the beginning. Must parse as an integer. | 
 **limit** | **string** | Limit caps how many messages come back. Empty or 0 uses the store&#39;s default page size. Must parse as an integer. | 

### Return type

[**CloudInboxPage**](CloudInboxPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ChannelsPairing

> CloudPairingQueue CloudGetV1ChannelsPairing(ctx).Execute()

Returns the pairing requests waiting for the caller org to approve — one per person who messaged a connected bot on a channel whose DM policy is \"pairing\" and who is not allowed yet.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.CloudGetV1ChannelsPairing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.CloudGetV1ChannelsPairing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ChannelsPairing`: CloudPairingQueue
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.CloudGetV1ChannelsPairing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ChannelsPairingRequest struct via the builder pattern


### Return type

[**CloudPairingQueue**](CloudPairingQueue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ChannelsByChannelSend

> CloudPostV1ChannelsByChannelSend(ctx, channel).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	channel := "channel_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChannelsAPI.CloudPostV1ChannelsByChannelSend(context.Background(), channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.CloudPostV1ChannelsByChannelSend``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1ChannelsByChannelSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ChannelsPairingApprove

> CloudPairingApproved CloudPostV1ChannelsPairingApprove(ctx).CloudApprovePairingIn(cloudApprovePairingIn).Execute()

Turns one pending pairing code into a standing allow entry, so that person can DM the org's bot on that channel from now on.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	cloudApprovePairingIn := *openapiclient.NewCloudApprovePairingIn() // CloudApprovePairingIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.CloudPostV1ChannelsPairingApprove(context.Background()).CloudApprovePairingIn(cloudApprovePairingIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.CloudPostV1ChannelsPairingApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ChannelsPairingApprove`: CloudPairingApproved
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.CloudPostV1ChannelsPairingApprove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ChannelsPairingApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudApprovePairingIn** | [**CloudApprovePairingIn**](CloudApprovePairingIn.md) |  | 

### Return type

[**CloudPairingApproved**](CloudPairingApproved.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1ChannelsAllowlist

> CloudAllowlistView CloudPutV1ChannelsAllowlist(ctx).CloudAllowlistPutIn(cloudAllowlistPutIn).Execute()

Edits the caller org's access policy for one channel and answers the policy as GET would, so both verbs return ONE shape.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	cloudAllowlistPutIn := *openapiclient.NewCloudAllowlistPutIn() // CloudAllowlistPutIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.CloudPutV1ChannelsAllowlist(context.Background()).CloudAllowlistPutIn(cloudAllowlistPutIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.CloudPutV1ChannelsAllowlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1ChannelsAllowlist`: CloudAllowlistView
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.CloudPutV1ChannelsAllowlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1ChannelsAllowlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudAllowlistPutIn** | [**CloudAllowlistPutIn**](CloudAllowlistPutIn.md) |  | 

### Return type

[**CloudAllowlistView**](CloudAllowlistView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

