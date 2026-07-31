# \CompletionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayCreateCompletion**](CompletionsAPI.md#GatewayCreateCompletion) | **Post** /v1/gateway/completions | Create completion



## GatewayCreateCompletion

> map[string]interface{} GatewayCreateCompletion(ctx).GatewayCreateCompletionRequest(gatewayCreateCompletionRequest).Execute()

Create completion

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
	gatewayCreateCompletionRequest := *openapiclient.NewGatewayCreateCompletionRequest("Model_example", openapiclient.gateway_createCompletion_request_prompt{ArrayOfString: new([]string)}) // GatewayCreateCompletionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompletionsAPI.GatewayCreateCompletion(context.Background()).GatewayCreateCompletionRequest(gatewayCreateCompletionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompletionsAPI.GatewayCreateCompletion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayCreateCompletion`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CompletionsAPI.GatewayCreateCompletion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateCompletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayCreateCompletionRequest** | [**GatewayCreateCompletionRequest**](GatewayCreateCompletionRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

