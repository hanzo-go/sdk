# \OperativeTasksAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OperativeGetMessages**](OperativeTasksAPI.md#OperativeGetMessages) | **Get** /v1/operative/sessions/{sessionId}/messages | Get conversation history for a session
[**OperativeSendMessage**](OperativeTasksAPI.md#OperativeSendMessage) | **Post** /v1/operative/sessions/{sessionId}/messages | Send a message to the operative agent



## OperativeGetMessages

> OperativeGetMessages200Response OperativeGetMessages(ctx, sessionId).Execute()

Get conversation history for a session

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
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique session identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperativeTasksAPI.OperativeGetMessages(context.Background(), sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperativeTasksAPI.OperativeGetMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperativeGetMessages`: OperativeGetMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `OperativeTasksAPI.OperativeGetMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Unique session identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperativeGetMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperativeGetMessages200Response**](OperativeGetMessages200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperativeSendMessage

> OperativeAgentResponse OperativeSendMessage(ctx, sessionId).OperativeSendMessageRequest(operativeSendMessageRequest).Execute()

Send a message to the operative agent



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
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique session identifier
	operativeSendMessageRequest := *openapiclient.NewOperativeSendMessageRequest("Content_example") // OperativeSendMessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperativeTasksAPI.OperativeSendMessage(context.Background(), sessionId).OperativeSendMessageRequest(operativeSendMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperativeTasksAPI.OperativeSendMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperativeSendMessage`: OperativeAgentResponse
	fmt.Fprintf(os.Stdout, "Response from `OperativeTasksAPI.OperativeSendMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Unique session identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperativeSendMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operativeSendMessageRequest** | [**OperativeSendMessageRequest**](OperativeSendMessageRequest.md) |  | 

### Return type

[**OperativeAgentResponse**](OperativeAgentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

