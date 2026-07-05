# \FunctionsMetricsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FunctionsFunctionMetrics**](FunctionsMetricsAPI.md#FunctionsFunctionMetrics) | **Get** /v1/functions/metrics | Invocation histogram + status breakdown



## FunctionsFunctionMetrics

> FunctionsMetrics FunctionsFunctionMetrics(ctx).Range_(range_).Execute()

Invocation histogram + status breakdown

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
	range_ := "range__example" // string |  (optional) (default to "24H")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsMetricsAPI.FunctionsFunctionMetrics(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsMetricsAPI.FunctionsFunctionMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsFunctionMetrics`: FunctionsMetrics
	fmt.Fprintf(os.Stdout, "Response from `FunctionsMetricsAPI.FunctionsFunctionMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsFunctionMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** |  | [default to &quot;24H&quot;]

### Return type

[**FunctionsMetrics**](FunctionsMetrics.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

