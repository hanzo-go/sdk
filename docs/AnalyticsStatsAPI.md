# \AnalyticsStatsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsGetWebsiteMetrics**](AnalyticsStatsAPI.md#AnalyticsGetWebsiteMetrics) | **Get** /v1/analytics/websites/{websiteId}/metrics | Get breakdown metrics by type (url, referrer, browser, os, device, country, event, channel, etc.)
[**AnalyticsGetWebsiteStats**](AnalyticsStatsAPI.md#AnalyticsGetWebsiteStats) | **Get** /v1/analytics/websites/{websiteId}/stats | Get aggregate statistics for a website



## AnalyticsGetWebsiteMetrics

> []AnalyticsMetric AnalyticsGetWebsiteMetrics(ctx, websiteId).StartAt(startAt).EndAt(endAt).Type_(type_).Limit(limit).Offset(offset).Search(search).Url(url).Referrer(referrer).Title(title).Os(os).Browser(browser).Device(device).Country(country).Region(region).City(city).Tag(tag).Host(host).Language(language).Event(event).Execute()

Get breakdown metrics by type (url, referrer, browser, os, device, country, event, channel, etc.)

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
	websiteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	startAt := int64(789) // int64 | Start timestamp in milliseconds
	endAt := int64(789) // int64 | End timestamp in milliseconds
	type_ := "type__example" // string | Metric type to break down by
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)
	url := "url_example" // string |  (optional)
	referrer := "referrer_example" // string |  (optional)
	title := "title_example" // string |  (optional)
	os := "os_example" // string |  (optional)
	browser := "browser_example" // string |  (optional)
	device := "device_example" // string |  (optional)
	country := "country_example" // string |  (optional)
	region := "region_example" // string |  (optional)
	city := "city_example" // string |  (optional)
	tag := "tag_example" // string |  (optional)
	host := "host_example" // string |  (optional)
	language := "language_example" // string |  (optional)
	event := "event_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsStatsAPI.AnalyticsGetWebsiteMetrics(context.Background(), websiteId).StartAt(startAt).EndAt(endAt).Type_(type_).Limit(limit).Offset(offset).Search(search).Url(url).Referrer(referrer).Title(title).Os(os).Browser(browser).Device(device).Country(country).Region(region).City(city).Tag(tag).Host(host).Language(language).Event(event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsStatsAPI.AnalyticsGetWebsiteMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetWebsiteMetrics`: []AnalyticsMetric
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsStatsAPI.AnalyticsGetWebsiteMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetWebsiteMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 
 **type_** | **string** | Metric type to break down by | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **search** | **string** |  | 
 **url** | **string** |  | 
 **referrer** | **string** |  | 
 **title** | **string** |  | 
 **os** | **string** |  | 
 **browser** | **string** |  | 
 **device** | **string** |  | 
 **country** | **string** |  | 
 **region** | **string** |  | 
 **city** | **string** |  | 
 **tag** | **string** |  | 
 **host** | **string** |  | 
 **language** | **string** |  | 
 **event** | **string** |  | 

### Return type

[**[]AnalyticsMetric**](AnalyticsMetric.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetWebsiteStats

> AnalyticsGetWebsiteStats200Response AnalyticsGetWebsiteStats(ctx, websiteId).StartAt(startAt).EndAt(endAt).Compare(compare).Url(url).Referrer(referrer).Title(title).Os(os).Browser(browser).Device(device).Country(country).Region(region).City(city).Tag(tag).Host(host).Language(language).Event(event).Execute()

Get aggregate statistics for a website

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
	websiteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	startAt := int64(789) // int64 | Start timestamp in milliseconds
	endAt := int64(789) // int64 | End timestamp in milliseconds
	compare := "compare_example" // string | Compare period (e.g. \"previous_period\") (optional)
	url := "url_example" // string |  (optional)
	referrer := "referrer_example" // string |  (optional)
	title := "title_example" // string |  (optional)
	os := "os_example" // string |  (optional)
	browser := "browser_example" // string |  (optional)
	device := "device_example" // string |  (optional)
	country := "country_example" // string |  (optional)
	region := "region_example" // string |  (optional)
	city := "city_example" // string |  (optional)
	tag := "tag_example" // string |  (optional)
	host := "host_example" // string |  (optional)
	language := "language_example" // string |  (optional)
	event := "event_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsStatsAPI.AnalyticsGetWebsiteStats(context.Background(), websiteId).StartAt(startAt).EndAt(endAt).Compare(compare).Url(url).Referrer(referrer).Title(title).Os(os).Browser(browser).Device(device).Country(country).Region(region).City(city).Tag(tag).Host(host).Language(language).Event(event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsStatsAPI.AnalyticsGetWebsiteStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetWebsiteStats`: AnalyticsGetWebsiteStats200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsStatsAPI.AnalyticsGetWebsiteStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetWebsiteStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 
 **compare** | **string** | Compare period (e.g. \&quot;previous_period\&quot;) | 
 **url** | **string** |  | 
 **referrer** | **string** |  | 
 **title** | **string** |  | 
 **os** | **string** |  | 
 **browser** | **string** |  | 
 **device** | **string** |  | 
 **country** | **string** |  | 
 **region** | **string** |  | 
 **city** | **string** |  | 
 **tag** | **string** |  | 
 **host** | **string** |  | 
 **language** | **string** |  | 
 **event** | **string** |  | 

### Return type

[**AnalyticsGetWebsiteStats200Response**](AnalyticsGetWebsiteStats200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

