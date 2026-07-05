# \MqPublishAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MqPublishMessage**](MqPublishAPI.md#MqPublishMessage) | **Post** /v1/mq/publish | Publish message to subject
[**MqRequestReply**](MqPublishAPI.md#MqRequestReply) | **Post** /v1/mq/request | Request/reply pattern



## MqPublishMessage

> MqPublishResponse MqPublishMessage(ctx).MqPublishRequest(mqPublishRequest).Execute()

Publish message to subject



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
	mqPublishRequest := *openapiclient.NewMqPublishRequest("Subject_example", "Data_example") // MqPublishRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqPublishAPI.MqPublishMessage(context.Background()).MqPublishRequest(mqPublishRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqPublishAPI.MqPublishMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqPublishMessage`: MqPublishResponse
	fmt.Fprintf(os.Stdout, "Response from `MqPublishAPI.MqPublishMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMqPublishMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mqPublishRequest** | [**MqPublishRequest**](MqPublishRequest.md) |  | 

### Return type

[**MqPublishResponse**](MqPublishResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqRequestReply

> MqMessage MqRequestReply(ctx).MqRequestReply(mqRequestReply).Execute()

Request/reply pattern



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
	mqRequestReply := *openapiclient.NewMqRequestReply("Subject_example", "Data_example") // MqRequestReply | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqPublishAPI.MqRequestReply(context.Background()).MqRequestReply(mqRequestReply).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqPublishAPI.MqRequestReply``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqRequestReply`: MqMessage
	fmt.Fprintf(os.Stdout, "Response from `MqPublishAPI.MqRequestReply`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMqRequestReplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mqRequestReply** | [**MqRequestReply**](MqRequestReply.md) |  | 

### Return type

[**MqMessage**](MqMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

