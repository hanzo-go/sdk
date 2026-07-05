# \MqSubscribeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MqSubscribe**](MqSubscribeAPI.md#MqSubscribe) | **Get** /v1/mq/subscribe/{subject} | Subscribe to subject via SSE



## MqSubscribe

> MqMessage MqSubscribe(ctx, subject).Queue(queue).Execute()

Subscribe to subject via SSE



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
	subject := "subject_example" // string | Subject to subscribe to. Supports wildcards (`events.*`, `logs.>`). 
	queue := "queue_example" // string | Queue group name for load-balanced delivery. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqSubscribeAPI.MqSubscribe(context.Background(), subject).Queue(queue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqSubscribeAPI.MqSubscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqSubscribe`: MqMessage
	fmt.Fprintf(os.Stdout, "Response from `MqSubscribeAPI.MqSubscribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subject** | **string** | Subject to subscribe to. Supports wildcards (&#x60;events.*&#x60;, &#x60;logs.&gt;&#x60;).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queue** | **string** | Queue group name for load-balanced delivery. | 

### Return type

[**MqMessage**](MqMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

