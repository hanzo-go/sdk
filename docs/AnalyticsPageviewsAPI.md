# \AnalyticsPageviewsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsGetPageviews**](AnalyticsPageviewsAPI.md#AnalyticsGetPageviews) | **Get** /v1/analytics/websites/{websiteId}/pageviews | Get pageview and session time series



## AnalyticsGetPageviews

> AnalyticsGetPageviews200Response AnalyticsGetPageviews(ctx, websiteId).StartAt(startAt).EndAt(endAt).Unit(unit).Timezone(timezone).Compare(compare).Url(url).Referrer(referrer).Title(title).Os(os).Browser(browser).Device(device).Country(country).Region(region).City(city).Tag(tag).Host(host).Language(language).Event(event).Execute()

Get pageview and session time series

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
	unit := "unit_example" // string |  (optional)
	timezone := "America/Los_Angeles" // string |  (optional)
	compare := "compare_example" // string |  (optional)
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
	resp, r, err := apiClient.AnalyticsPageviewsAPI.AnalyticsGetPageviews(context.Background(), websiteId).StartAt(startAt).EndAt(endAt).Unit(unit).Timezone(timezone).Compare(compare).Url(url).Referrer(referrer).Title(title).Os(os).Browser(browser).Device(device).Country(country).Region(region).City(city).Tag(tag).Host(host).Language(language).Event(event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsPageviewsAPI.AnalyticsGetPageviews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetPageviews`: AnalyticsGetPageviews200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsPageviewsAPI.AnalyticsGetPageviews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetPageviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 
 **unit** | **string** |  | 
 **timezone** | **string** |  | 
 **compare** | **string** |  | 
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

[**AnalyticsGetPageviews200Response**](AnalyticsGetPageviews200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

