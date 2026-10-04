# \PubsubAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostPubsubPublish**](PubsubAPI.md#PostPubsubPublish) | **Post** /v1/pubsub/publish | Puts one message on the org&#39;s bus.
[**PostPubsubRequest**](PubsubAPI.md#PostPubsubRequest) | **Post** /v1/pubsub/request | Sends one request on the org&#39;s bus and waits for one reply — the synchronous half of pub/sub, for callers speaking to a responder subscribed on the NATS port.



## PostPubsubPublish

> PubsubBusAck PostPubsubPublish(ctx).PubsubBusPublish(pubsubBusPublish).Execute()

Puts one message on the org's bus.



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
	pubsubBusPublish := *openapiclient.NewPubsubBusPublish() // PubsubBusPublish | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubAPI.PostPubsubPublish(context.Background()).PubsubBusPublish(pubsubBusPublish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.PostPubsubPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPubsubPublish`: PubsubBusAck
	fmt.Fprintf(os.Stdout, "Response from `PubsubAPI.PostPubsubPublish`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPubsubPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pubsubBusPublish** | [**PubsubBusPublish**](PubsubBusPublish.md) |  | 

### Return type

[**PubsubBusAck**](PubsubBusAck.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPubsubRequest

> PubsubBusMessage PostPubsubRequest(ctx).PubsubBusRequest(pubsubBusRequest).Execute()

Sends one request on the org's bus and waits for one reply — the synchronous half of pub/sub, for callers speaking to a responder subscribed on the NATS port.



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
	pubsubBusRequest := *openapiclient.NewPubsubBusRequest() // PubsubBusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubAPI.PostPubsubRequest(context.Background()).PubsubBusRequest(pubsubBusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubAPI.PostPubsubRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPubsubRequest`: PubsubBusMessage
	fmt.Fprintf(os.Stdout, "Response from `PubsubAPI.PostPubsubRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPubsubRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pubsubBusRequest** | [**PubsubBusRequest**](PubsubBusRequest.md) |  | 

### Return type

[**PubsubBusMessage**](PubsubBusMessage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

