# \PubsubPublishAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PubsubPublishMessage**](PubsubPublishAPI.md#PubsubPublishMessage) | **Post** /v1/pubsub/publish | Publish a message



## PubsubPublishMessage

> PubsubPublishResponse PubsubPublishMessage(ctx).PubsubPublishRequest(pubsubPublishRequest).Execute()

Publish a message



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubPublishAPI.PubsubPublishMessage(context.Background()).PubsubPublishRequest(pubsubPublishRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubPublishAPI.PubsubPublishMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubPublishMessage`: PubsubPublishResponse
	fmt.Fprintf(os.Stdout, "Response from `PubsubPublishAPI.PubsubPublishMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPubsubPublishMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pubsubPublishRequest** | [**PubsubPublishRequest**](PubsubPublishRequest.md) |  | 

### Return type

[**PubsubPublishResponse**](PubsubPublishResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

