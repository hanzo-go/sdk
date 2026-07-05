# \EdgeLogsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeGetFunctionLogs**](EdgeLogsAPI.md#EdgeGetFunctionLogs) | **Get** /v1/edge/functions/{slug}/logs | Get function logs



## EdgeGetFunctionLogs

> []EdgeLogEntry EdgeGetFunctionLogs(ctx, slug).Since(since).Until(until).Level(level).Limit(limit).Execute()

Get function logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	slug := "slug_example" // string | 
	since := time.Now() // time.Time | Return logs after this timestamp (optional)
	until := time.Now() // time.Time |  (optional)
	level := "level_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeLogsAPI.EdgeGetFunctionLogs(context.Background(), slug).Since(since).Until(until).Level(level).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeLogsAPI.EdgeGetFunctionLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeGetFunctionLogs`: []EdgeLogEntry
	fmt.Fprintf(os.Stdout, "Response from `EdgeLogsAPI.EdgeGetFunctionLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeGetFunctionLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **time.Time** | Return logs after this timestamp | 
 **until** | **time.Time** |  | 
 **level** | **string** |  | 
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]EdgeLogEntry**](EdgeLogEntry.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

