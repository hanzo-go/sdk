# \EdgeVersionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeListFunctionVersions**](EdgeVersionsAPI.md#EdgeListFunctionVersions) | **Get** /v1/edge/functions/{slug}/versions | List function versions
[**EdgeRollbackFunction**](EdgeVersionsAPI.md#EdgeRollbackFunction) | **Post** /v1/edge/functions/{slug}/rollback | Rollback to version



## EdgeListFunctionVersions

> []EdgeFunctionVersion EdgeListFunctionVersions(ctx, slug).Execute()

List function versions

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
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeVersionsAPI.EdgeListFunctionVersions(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeVersionsAPI.EdgeListFunctionVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeListFunctionVersions`: []EdgeFunctionVersion
	fmt.Fprintf(os.Stdout, "Response from `EdgeVersionsAPI.EdgeListFunctionVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeListFunctionVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EdgeFunctionVersion**](EdgeFunctionVersion.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeRollbackFunction

> EdgeFunction EdgeRollbackFunction(ctx, slug).EdgeRollbackFunctionRequest(edgeRollbackFunctionRequest).Execute()

Rollback to version

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
	slug := "slug_example" // string | 
	edgeRollbackFunctionRequest := *openapiclient.NewEdgeRollbackFunctionRequest(int32(123)) // EdgeRollbackFunctionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeVersionsAPI.EdgeRollbackFunction(context.Background(), slug).EdgeRollbackFunctionRequest(edgeRollbackFunctionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeVersionsAPI.EdgeRollbackFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeRollbackFunction`: EdgeFunction
	fmt.Fprintf(os.Stdout, "Response from `EdgeVersionsAPI.EdgeRollbackFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeRollbackFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeRollbackFunctionRequest** | [**EdgeRollbackFunctionRequest**](EdgeRollbackFunctionRequest.md) |  | 

### Return type

[**EdgeFunction**](EdgeFunction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

