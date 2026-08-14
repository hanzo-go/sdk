# \AnalyticsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnalyticsHealth**](AnalyticsAPI.md#GetAnalyticsHealth) | **Get** /v1/analytics/health | Health reports whether the event plane can take a write and the warehouse can answer a read.
[**GetAnalyticsOverview**](AnalyticsAPI.md#GetAnalyticsOverview) | **Get** /v1/analytics/overview | Overview returns the caller org&#39;s analytics KPIs for one time window.
[**GetAnalyticsTimeseries**](AnalyticsAPI.md#GetAnalyticsTimeseries) | **Get** /v1/analytics/timeseries | Timeseries returns the caller org&#39;s LLM usage over time as an evenly-spaced series.
[**GetAnalyticsTop**](AnalyticsAPI.md#GetAnalyticsTop) | **Get** /v1/analytics/top | Top returns the caller org&#39;s ranked lenses for one window, five of them at once.



## GetAnalyticsHealth

> HealthReport GetAnalyticsHealth(ctx).Execute()

Health reports whether the event plane can take a write and the warehouse can answer a read.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetAnalyticsHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetAnalyticsHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalyticsHealth`: HealthReport
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetAnalyticsHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsHealthRequest struct via the builder pattern


### Return type

[**HealthReport**](HealthReport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticsOverview

> Overview GetAnalyticsOverview(ctx).Range_(range_).Start(start).End(end).Execute()

Overview returns the caller org's analytics KPIs for one time window.



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
	range_ := "7d" // string | Range is a relative window: a count and a unit — 24h, 7d, 90d, any <N>h or <N>d — or day, week, month, all. Default 24h. Ignored when both start and end are given. An unknown value, or one past the 730-day horizon, is a 400. (optional)
	start := "start_example" // string | Start is the inclusive lower bound of a custom window, RFC3339. Requires end. (optional)
	end := "end_example" // string | End is the exclusive upper bound of a custom window, RFC3339. Requires start. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetAnalyticsOverview(context.Background()).Range_(range_).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetAnalyticsOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalyticsOverview`: Overview
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetAnalyticsOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is a relative window: a count and a unit — 24h, 7d, 90d, any &lt;N&gt;h or &lt;N&gt;d — or day, week, month, all. Default 24h. Ignored when both start and end are given. An unknown value, or one past the 730-day horizon, is a 400. | 
 **start** | **string** | Start is the inclusive lower bound of a custom window, RFC3339. Requires end. | 
 **end** | **string** | End is the exclusive upper bound of a custom window, RFC3339. Requires start. | 

### Return type

[**Overview**](Overview.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticsTimeseries

> Timeseries GetAnalyticsTimeseries(ctx).Range_(range_).Start(start).End(end).Execute()

Timeseries returns the caller org's LLM usage over time as an evenly-spaced series.



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
	range_ := "30d" // string | Range is a relative window: a count and a unit — 24h, 7d, 90d, any <N>h or <N>d — or day, week, month, all. Default 24h. Ignored when both start and end are given. An unknown value, or one past the 730-day horizon, is a 400. (optional)
	start := "start_example" // string | Start is the inclusive lower bound of a custom window, RFC3339. Requires end. (optional)
	end := "end_example" // string | End is the exclusive upper bound of a custom window, RFC3339. Requires start. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetAnalyticsTimeseries(context.Background()).Range_(range_).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetAnalyticsTimeseries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalyticsTimeseries`: Timeseries
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetAnalyticsTimeseries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsTimeseriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is a relative window: a count and a unit — 24h, 7d, 90d, any &lt;N&gt;h or &lt;N&gt;d — or day, week, month, all. Default 24h. Ignored when both start and end are given. An unknown value, or one past the 730-day horizon, is a 400. | 
 **start** | **string** | Start is the inclusive lower bound of a custom window, RFC3339. Requires end. | 
 **end** | **string** | End is the exclusive upper bound of a custom window, RFC3339. Requires start. | 

### Return type

[**Timeseries**](Timeseries.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticsTop

> Top GetAnalyticsTop(ctx).Range_(range_).Start(start).End(end).Limit(limit).Execute()

Top returns the caller org's ranked lenses for one window, five of them at once.



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
	range_ := "7d" // string | Range is a relative window: a count and a unit — 24h, 7d, 90d, any <N>h or <N>d — or day, week, month, all. Default 24h. Ignored when both start and end are given. An unknown value, or one past the 730-day horizon, is a 400. (optional)
	start := "start_example" // string | Start is the inclusive lower bound of a custom window, RFC3339. Requires end. (optional)
	end := "end_example" // string | End is the exclusive upper bound of a custom window, RFC3339. Requires start. (optional)
	limit := int32(25) // int32 | Limit bounds every ranked lens in the response. Default 10, maximum 100; a value at or below zero, or one that is not a number, takes the default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetAnalyticsTop(context.Background()).Range_(range_).Start(start).End(end).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetAnalyticsTop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalyticsTop`: Top
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetAnalyticsTop`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticsTopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is a relative window: a count and a unit — 24h, 7d, 90d, any &lt;N&gt;h or &lt;N&gt;d — or day, week, month, all. Default 24h. Ignored when both start and end are given. An unknown value, or one past the 730-day horizon, is a 400. | 
 **start** | **string** | Start is the inclusive lower bound of a custom window, RFC3339. Requires end. | 
 **end** | **string** | End is the exclusive upper bound of a custom window, RFC3339. Requires start. | 
 **limit** | **int32** | Limit bounds every ranked lens in the response. Default 10, maximum 100; a value at or below zero, or one that is not a number, takes the default. | 

### Return type

[**Top**](Top.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

