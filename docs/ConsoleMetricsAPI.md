# \ConsoleMetricsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleGetDailyMetrics**](ConsoleMetricsAPI.md#ConsoleGetDailyMetrics) | **Get** /v1/console/metrics/daily | Get daily metrics
[**ConsoleGetMetrics**](ConsoleMetricsAPI.md#ConsoleGetMetrics) | **Get** /v1/console/metrics | Get metrics from the project



## ConsoleGetDailyMetrics

> ConsoleGetDailyMetrics200Response ConsoleGetDailyMetrics(ctx).Page(page).Limit(limit).TraceName(traceName).UserId(userId).Tags(tags).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Environment(environment).Execute()

Get daily metrics

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
	page := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)
	traceName := "traceName_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)
	tags := []string{"Inner_example"} // []string |  (optional)
	fromTimestamp := time.Now() // time.Time |  (optional)
	toTimestamp := time.Now() // time.Time |  (optional)
	environment := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleMetricsAPI.ConsoleGetDailyMetrics(context.Background()).Page(page).Limit(limit).TraceName(traceName).UserId(userId).Tags(tags).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Environment(environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleMetricsAPI.ConsoleGetDailyMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetDailyMetrics`: ConsoleGetDailyMetrics200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleMetricsAPI.ConsoleGetDailyMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetDailyMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **limit** | **int32** |  | 
 **traceName** | **string** |  | 
 **userId** | **string** |  | 
 **tags** | **[]string** |  | 
 **fromTimestamp** | **time.Time** |  | 
 **toTimestamp** | **time.Time** |  | 
 **environment** | **[]string** |  | 

### Return type

[**ConsoleGetDailyMetrics200Response**](ConsoleGetDailyMetrics200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleGetMetrics

> ConsoleGetMetrics200Response ConsoleGetMetrics(ctx).Query(query).Execute()

Get metrics from the project

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
	query := "query_example" // string | JSON string containing the query parameters (view, dimensions, metrics, filters, timeDimension, fromTimestamp, toTimestamp)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleMetricsAPI.ConsoleGetMetrics(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleMetricsAPI.ConsoleGetMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetMetrics`: ConsoleGetMetrics200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleMetricsAPI.ConsoleGetMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | JSON string containing the query parameters (view, dimensions, metrics, filters, timeDimension, fromTimestamp, toTimestamp) | 

### Return type

[**ConsoleGetMetrics200Response**](ConsoleGetMetrics200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

