# \AnalyticsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1AnalyticsHealth**](AnalyticsAPI.md#CloudGetV1AnalyticsHealth) | **Get** /v1/analytics/health | 
[**CloudGetV1AnalyticsOverview**](AnalyticsAPI.md#CloudGetV1AnalyticsOverview) | **Get** /v1/analytics/overview | Overview returns the caller org&#39;s analytics KPIs for one time window.
[**CloudGetV1AnalyticsTimeseries**](AnalyticsAPI.md#CloudGetV1AnalyticsTimeseries) | **Get** /v1/analytics/timeseries | Timeseries returns the caller org&#39;s LLM usage over time as an evenly-spaced series.
[**CloudGetV1AnalyticsTop**](AnalyticsAPI.md#CloudGetV1AnalyticsTop) | **Get** /v1/analytics/top | Top returns the caller org&#39;s ranked lenses for one window, five of them at once.
[**DnsGetZoneAnalytics**](AnalyticsAPI.md#DnsGetZoneAnalytics) | **Get** /v1/dns/zones/{zone}/analytics | Get query analytics



## CloudGetV1AnalyticsHealth

> CloudHealthReport CloudGetV1AnalyticsHealth(ctx).Execute()



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
	resp, r, err := apiClient.AnalyticsAPI.CloudGetV1AnalyticsHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.CloudGetV1AnalyticsHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AnalyticsHealth`: CloudHealthReport
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.CloudGetV1AnalyticsHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AnalyticsHealthRequest struct via the builder pattern


### Return type

[**CloudHealthReport**](CloudHealthReport.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AnalyticsOverview

> CloudOverview CloudGetV1AnalyticsOverview(ctx).Range_(range_).Start(start).End(end).Execute()

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
	range_ := "7d" // string | Range is a relative window: 24h, 7d or 30d. Default 24h. Ignored when both start and end are given. An unknown value is a 400. (optional)
	start := "start_example" // string | Start is the inclusive lower bound of a custom window, RFC3339. Requires end. (optional)
	end := "end_example" // string | End is the exclusive upper bound of a custom window, RFC3339. Requires start. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.CloudGetV1AnalyticsOverview(context.Background()).Range_(range_).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.CloudGetV1AnalyticsOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AnalyticsOverview`: CloudOverview
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.CloudGetV1AnalyticsOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AnalyticsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is a relative window: 24h, 7d or 30d. Default 24h. Ignored when both start and end are given. An unknown value is a 400. | 
 **start** | **string** | Start is the inclusive lower bound of a custom window, RFC3339. Requires end. | 
 **end** | **string** | End is the exclusive upper bound of a custom window, RFC3339. Requires start. | 

### Return type

[**CloudOverview**](CloudOverview.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AnalyticsTimeseries

> CloudTimeseries CloudGetV1AnalyticsTimeseries(ctx).Range_(range_).Start(start).End(end).Execute()

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
	range_ := "30d" // string | Range is a relative window: 24h, 7d or 30d. Default 24h. Ignored when both start and end are given. An unknown value is a 400. (optional)
	start := "start_example" // string | Start is the inclusive lower bound of a custom window, RFC3339. Requires end. (optional)
	end := "end_example" // string | End is the exclusive upper bound of a custom window, RFC3339. Requires start. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.CloudGetV1AnalyticsTimeseries(context.Background()).Range_(range_).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.CloudGetV1AnalyticsTimeseries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AnalyticsTimeseries`: CloudTimeseries
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.CloudGetV1AnalyticsTimeseries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AnalyticsTimeseriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is a relative window: 24h, 7d or 30d. Default 24h. Ignored when both start and end are given. An unknown value is a 400. | 
 **start** | **string** | Start is the inclusive lower bound of a custom window, RFC3339. Requires end. | 
 **end** | **string** | End is the exclusive upper bound of a custom window, RFC3339. Requires start. | 

### Return type

[**CloudTimeseries**](CloudTimeseries.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AnalyticsTop

> CloudTop CloudGetV1AnalyticsTop(ctx).Range_(range_).Start(start).End(end).Limit(limit).Execute()

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
	range_ := "7d" // string | Range is a relative window: 24h, 7d or 30d. Default 24h. Ignored when both start and end are given. An unknown value is a 400. (optional)
	start := "start_example" // string | Start is the inclusive lower bound of a custom window, RFC3339. Requires end. (optional)
	end := "end_example" // string | End is the exclusive upper bound of a custom window, RFC3339. Requires start. (optional)
	limit := int32(25) // int32 | Limit bounds every ranked lens in the response. Default 10, maximum 100; a value at or below zero, or one that is not a number, takes the default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.CloudGetV1AnalyticsTop(context.Background()).Range_(range_).Start(start).End(end).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.CloudGetV1AnalyticsTop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AnalyticsTop`: CloudTop
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.CloudGetV1AnalyticsTop`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AnalyticsTopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is a relative window: 24h, 7d or 30d. Default 24h. Ignored when both start and end are given. An unknown value is a 400. | 
 **start** | **string** | Start is the inclusive lower bound of a custom window, RFC3339. Requires end. | 
 **end** | **string** | End is the exclusive upper bound of a custom window, RFC3339. Requires start. | 
 **limit** | **int32** | Limit bounds every ranked lens in the response. Default 10, maximum 100; a value at or below zero, or one that is not a number, takes the default. | 

### Return type

[**CloudTop**](CloudTop.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsGetZoneAnalytics

> DnsQueryAnalytics DnsGetZoneAnalytics(ctx, zone).From(from).To(to).Granularity(granularity).Execute()

Get query analytics

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
	zone := "zone_example" // string | 
	from := time.Now() // time.Time |  (optional)
	to := time.Now() // time.Time |  (optional)
	granularity := "granularity_example" // string |  (optional) (default to "day")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.DnsGetZoneAnalytics(context.Background(), zone).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.DnsGetZoneAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsGetZoneAnalytics`: DnsQueryAnalytics
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.DnsGetZoneAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsGetZoneAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 
 **granularity** | **string** |  | [default to &quot;day&quot;]

### Return type

[**DnsQueryAnalytics**](DnsQueryAnalytics.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

