# \AiClaudeCompatibleAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiCreateMessage**](AiClaudeCompatibleAPI.md#AiCreateMessage) | **Post** /v1/messages | Create a message (Anthropic-compatible)



## AiCreateMessage

> AiMessageResponse AiCreateMessage(ctx).AiMessageRequest(aiMessageRequest).Execute()

Create a message (Anthropic-compatible)



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
	aiMessageRequest := *openapiclient.NewAiMessageRequest("Model_example", []map[string]interface{}{map[string]interface{}(123)}, int32(123)) // AiMessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiClaudeCompatibleAPI.AiCreateMessage(context.Background()).AiMessageRequest(aiMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiClaudeCompatibleAPI.AiCreateMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiCreateMessage`: AiMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `AiClaudeCompatibleAPI.AiCreateMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiCreateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiMessageRequest** | [**AiMessageRequest**](AiMessageRequest.md) |  | 

### Return type

[**AiMessageResponse**](AiMessageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

