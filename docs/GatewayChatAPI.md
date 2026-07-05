# \GatewayChatAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayCreateChatCompletion**](GatewayChatAPI.md#GatewayCreateChatCompletion) | **Post** /v1/gateway/chat/completions | Create chat completion



## GatewayCreateChatCompletion

> GatewayChatCompletionResponse GatewayCreateChatCompletion(ctx).GatewayChatCompletionRequest(gatewayChatCompletionRequest).Execute()

Create chat completion

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
	gatewayChatCompletionRequest := *openapiclient.NewGatewayChatCompletionRequest("Model_example", []openapiclient.GatewayChatMessage{*openapiclient.NewGatewayChatMessage("Role_example", openapiclient.gateway_ChatMessage_content{ArrayOfMapmapOfStringAny: new([]map[string]interface{})})}) // GatewayChatCompletionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayChatAPI.GatewayCreateChatCompletion(context.Background()).GatewayChatCompletionRequest(gatewayChatCompletionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayChatAPI.GatewayCreateChatCompletion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayCreateChatCompletion`: GatewayChatCompletionResponse
	fmt.Fprintf(os.Stdout, "Response from `GatewayChatAPI.GatewayCreateChatCompletion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateChatCompletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayChatCompletionRequest** | [**GatewayChatCompletionRequest**](GatewayChatCompletionRequest.md) |  | 

### Return type

[**GatewayChatCompletionResponse**](GatewayChatCompletionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

