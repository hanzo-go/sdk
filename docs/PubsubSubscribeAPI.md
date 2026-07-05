# \PubsubSubscribeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PubsubSubscribe**](PubsubSubscribeAPI.md#PubsubSubscribe) | **Get** /v1/pubsub/subscribe | Subscribe to a subject (SSE)



## PubsubSubscribe

> PubsubMessage PubsubSubscribe(ctx).Subject(subject).Queue(queue).Execute()

Subscribe to a subject (SSE)



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
	subject := "orders.>" // string | Subject to subscribe to (supports wildcards)
	queue := "queue_example" // string | Queue group name for load-balanced delivery (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubSubscribeAPI.PubsubSubscribe(context.Background()).Subject(subject).Queue(queue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubSubscribeAPI.PubsubSubscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubSubscribe`: PubsubMessage
	fmt.Fprintf(os.Stdout, "Response from `PubsubSubscribeAPI.PubsubSubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPubsubSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subject** | **string** | Subject to subscribe to (supports wildcards) | 
 **queue** | **string** | Queue group name for load-balanced delivery | 

### Return type

[**PubsubMessage**](PubsubMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

