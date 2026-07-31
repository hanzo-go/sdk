# \PubSubAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KvListChannels**](PubSubAPI.md#KvListChannels) | **Get** /v1/kv/pubsub/channels | List active channels
[**KvPublish**](PubSubAPI.md#KvPublish) | **Post** /v1/kv/pubsub/publish | Publish message to channel
[**KvSubscribe**](PubSubAPI.md#KvSubscribe) | **Get** /v1/kv/pubsub/subscribe | Subscribe to channels (SSE)



## KvListChannels

> KvListChannels200Response KvListChannels(ctx).Pattern(pattern).Execute()

List active channels

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
	pattern := "pattern_example" // string |  (optional) (default to "*")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubSubAPI.KvListChannels(context.Background()).Pattern(pattern).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.KvListChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvListChannels`: KvListChannels200Response
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.KvListChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKvListChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pattern** | **string** |  | [default to &quot;*&quot;]

### Return type

[**KvListChannels200Response**](KvListChannels200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvPublish

> KvPublish200Response KvPublish(ctx).KvPublishRequest(kvPublishRequest).Execute()

Publish message to channel

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
	kvPublishRequest := *openapiclient.NewKvPublishRequest("Channel_example", "Message_example") // KvPublishRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubSubAPI.KvPublish(context.Background()).KvPublishRequest(kvPublishRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.KvPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvPublish`: KvPublish200Response
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.KvPublish`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKvPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kvPublishRequest** | [**KvPublishRequest**](KvPublishRequest.md) |  | 

### Return type

[**KvPublish200Response**](KvPublish200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvSubscribe

> string KvSubscribe(ctx).Channels(channels).Pattern(pattern).Execute()

Subscribe to channels (SSE)

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
	channels := []string{"Inner_example"} // []string | Channels to subscribe to
	pattern := "pattern_example" // string | Pattern to subscribe to (e.g. user:*) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubSubAPI.KvSubscribe(context.Background()).Channels(channels).Pattern(pattern).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubSubAPI.KvSubscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvSubscribe`: string
	fmt.Fprintf(os.Stdout, "Response from `PubSubAPI.KvSubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKvSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channels** | **[]string** | Channels to subscribe to | 
 **pattern** | **string** | Pattern to subscribe to (e.g. user:*) | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

