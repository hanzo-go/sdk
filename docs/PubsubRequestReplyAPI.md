# \PubsubRequestReplyAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PubsubRequestReply**](PubsubRequestReplyAPI.md#PubsubRequestReply) | **Post** /v1/pubsub/request | Request/reply



## PubsubRequestReply

> PubsubMessage PubsubRequestReply(ctx).PubsubPublishRequest(pubsubPublishRequest).Timeout(timeout).Execute()

Request/reply



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
	pubsubPublishRequest := *openapiclient.NewPubsubPublishRequest("orders.created", "{"id":"order-123","total":59.99}") // PubsubPublishRequest | 
	timeout := int32(56) // int32 | Timeout in milliseconds (optional) (default to 5000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubRequestReplyAPI.PubsubRequestReply(context.Background()).PubsubPublishRequest(pubsubPublishRequest).Timeout(timeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubRequestReplyAPI.PubsubRequestReply``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubRequestReply`: PubsubMessage
	fmt.Fprintf(os.Stdout, "Response from `PubsubRequestReplyAPI.PubsubRequestReply`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPubsubRequestReplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pubsubPublishRequest** | [**PubsubPublishRequest**](PubsubPublishRequest.md) |  | 
 **timeout** | **int32** | Timeout in milliseconds | [default to 5000]

### Return type

[**PubsubMessage**](PubsubMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

