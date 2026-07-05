# \OperativeToolsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OperativeEditFile**](OperativeToolsAPI.md#OperativeEditFile) | **Post** /v1/operative/sessions/{sessionId}/edit | Execute a file editing command
[**OperativeExecuteBash**](OperativeToolsAPI.md#OperativeExecuteBash) | **Post** /v1/operative/sessions/{sessionId}/bash | Execute a bash command



## OperativeEditFile

> OperativeToolResult OperativeEditFile(ctx, sessionId).OperativeEditRequest(operativeEditRequest).Execute()

Execute a file editing command

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
	operativeEditRequest := *openapiclient.NewOperativeEditRequest("Command_example", "Path_example") // OperativeEditRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperativeToolsAPI.OperativeEditFile(context.Background(), sessionId).OperativeEditRequest(operativeEditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperativeToolsAPI.OperativeEditFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperativeEditFile`: OperativeToolResult
	fmt.Fprintf(os.Stdout, "Response from `OperativeToolsAPI.OperativeEditFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Unique session identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperativeEditFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operativeEditRequest** | [**OperativeEditRequest**](OperativeEditRequest.md) |  | 

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


## OperativeExecuteBash

> OperativeToolResult OperativeExecuteBash(ctx, sessionId).OperativeBashRequest(operativeBashRequest).Execute()

Execute a bash command

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
	operativeBashRequest := *openapiclient.NewOperativeBashRequest() // OperativeBashRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperativeToolsAPI.OperativeExecuteBash(context.Background(), sessionId).OperativeBashRequest(operativeBashRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperativeToolsAPI.OperativeExecuteBash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperativeExecuteBash`: OperativeToolResult
	fmt.Fprintf(os.Stdout, "Response from `OperativeToolsAPI.OperativeExecuteBash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Unique session identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperativeExecuteBashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operativeBashRequest** | [**OperativeBashRequest**](OperativeBashRequest.md) |  | 

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

