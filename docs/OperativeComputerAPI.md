# \OperativeComputerAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OperativeComputerAction**](OperativeComputerAPI.md#OperativeComputerAction) | **Post** /v1/operative/sessions/{sessionId}/computer | Execute a computer action
[**OperativeTakeScreenshot**](OperativeComputerAPI.md#OperativeTakeScreenshot) | **Get** /v1/operative/sessions/{sessionId}/screenshot | Take a screenshot of the current desktop



## OperativeComputerAction

> OperativeToolResult OperativeComputerAction(ctx, sessionId).OperativeComputerActionRequest(operativeComputerActionRequest).Execute()

Execute a computer action



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
	operativeComputerActionRequest := *openapiclient.NewOperativeComputerActionRequest("Action_example") // OperativeComputerActionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperativeComputerAPI.OperativeComputerAction(context.Background(), sessionId).OperativeComputerActionRequest(operativeComputerActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperativeComputerAPI.OperativeComputerAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperativeComputerAction`: OperativeToolResult
	fmt.Fprintf(os.Stdout, "Response from `OperativeComputerAPI.OperativeComputerAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Unique session identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperativeComputerActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operativeComputerActionRequest** | [**OperativeComputerActionRequest**](OperativeComputerActionRequest.md) |  | 

### Return type

[**OperativeToolResult**](OperativeToolResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperativeTakeScreenshot

> OperativeScreenshotResult OperativeTakeScreenshot(ctx, sessionId).Execute()

Take a screenshot of the current desktop

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
	resp, r, err := apiClient.OperativeComputerAPI.OperativeTakeScreenshot(context.Background(), sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperativeComputerAPI.OperativeTakeScreenshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperativeTakeScreenshot`: OperativeScreenshotResult
	fmt.Fprintf(os.Stdout, "Response from `OperativeComputerAPI.OperativeTakeScreenshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Unique session identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperativeTakeScreenshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperativeScreenshotResult**](OperativeScreenshotResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

