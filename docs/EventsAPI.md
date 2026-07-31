# \EventsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsGetEventDataEvents**](EventsAPI.md#AnalyticsGetEventDataEvents) | **Get** /v1/analytics/websites/{websiteId}/event-data/events | Get event data grouped by event name
[**AnalyticsGetEventDataFields**](EventsAPI.md#AnalyticsGetEventDataFields) | **Get** /v1/analytics/websites/{websiteId}/event-data/fields | Get event data fields
[**AnalyticsGetEventDataProperties**](EventsAPI.md#AnalyticsGetEventDataProperties) | **Get** /v1/analytics/websites/{websiteId}/event-data/properties | Get event data properties
[**AnalyticsGetEventDataStats**](EventsAPI.md#AnalyticsGetEventDataStats) | **Get** /v1/analytics/websites/{websiteId}/event-data/stats | Get event data aggregate stats
[**AnalyticsGetEventDataValues**](EventsAPI.md#AnalyticsGetEventDataValues) | **Get** /v1/analytics/websites/{websiteId}/event-data/values | Get event data values for a property
[**AnalyticsGetEventSeries**](EventsAPI.md#AnalyticsGetEventSeries) | **Get** /v1/analytics/websites/{websiteId}/events/series | Get event metrics as a time series
[**AnalyticsGetEvents**](EventsAPI.md#AnalyticsGetEvents) | **Get** /v1/analytics/websites/{websiteId}/events | Get paginated list of events
[**S3GetBucketNotification**](EventsAPI.md#S3GetBucketNotification) | **Get** /v1/s3/{bucket}?notification | Get event notification config
[**S3PutBucketNotification**](EventsAPI.md#S3PutBucketNotification) | **Put** /v1/s3/{bucket}?notification | Set event notification config



## AnalyticsGetEventDataEvents

> []map[string]interface{} AnalyticsGetEventDataEvents(ctx, websiteId).StartAt(startAt).EndAt(endAt).Event(event).Execute()

Get event data grouped by event name

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
	event := "event_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.AnalyticsGetEventDataEvents(context.Background(), websiteId).StartAt(startAt).EndAt(endAt).Event(event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.AnalyticsGetEventDataEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetEventDataEvents`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.AnalyticsGetEventDataEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetEventDataEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 
 **event** | **string** |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetEventDataFields

> []AnalyticsGetEventDataFields200ResponseInner AnalyticsGetEventDataFields(ctx, websiteId).StartAt(startAt).EndAt(endAt).Execute()

Get event data fields

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.AnalyticsGetEventDataFields(context.Background(), websiteId).StartAt(startAt).EndAt(endAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.AnalyticsGetEventDataFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetEventDataFields`: []AnalyticsGetEventDataFields200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.AnalyticsGetEventDataFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetEventDataFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 

### Return type

[**[]AnalyticsGetEventDataFields200ResponseInner**](AnalyticsGetEventDataFields200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetEventDataProperties

> []map[string]interface{} AnalyticsGetEventDataProperties(ctx, websiteId).StartAt(startAt).EndAt(endAt).PropertyName(propertyName).Execute()

Get event data properties

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
	propertyName := "propertyName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.AnalyticsGetEventDataProperties(context.Background(), websiteId).StartAt(startAt).EndAt(endAt).PropertyName(propertyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.AnalyticsGetEventDataProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetEventDataProperties`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.AnalyticsGetEventDataProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetEventDataPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 
 **propertyName** | **string** |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetEventDataStats

> []map[string]interface{} AnalyticsGetEventDataStats(ctx, websiteId).StartAt(startAt).EndAt(endAt).PropertyName(propertyName).Execute()

Get event data aggregate stats

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
	propertyName := "propertyName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.AnalyticsGetEventDataStats(context.Background(), websiteId).StartAt(startAt).EndAt(endAt).PropertyName(propertyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.AnalyticsGetEventDataStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetEventDataStats`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.AnalyticsGetEventDataStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetEventDataStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 
 **propertyName** | **string** |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetEventDataValues

> []map[string]interface{} AnalyticsGetEventDataValues(ctx, websiteId).StartAt(startAt).EndAt(endAt).EventName(eventName).PropertyName(propertyName).Execute()

Get event data values for a property

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
	eventName := "eventName_example" // string |  (optional)
	propertyName := "propertyName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.AnalyticsGetEventDataValues(context.Background(), websiteId).StartAt(startAt).EndAt(endAt).EventName(eventName).PropertyName(propertyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.AnalyticsGetEventDataValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetEventDataValues`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.AnalyticsGetEventDataValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetEventDataValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 
 **eventName** | **string** |  | 
 **propertyName** | **string** |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetEventSeries

> []AnalyticsPageviewSeries AnalyticsGetEventSeries(ctx, websiteId).StartAt(startAt).EndAt(endAt).Unit(unit).Timezone(timezone).Url(url).Referrer(referrer).Title(title).Os(os).Browser(browser).Device(device).Country(country).Region(region).City(city).Tag(tag).Host(host).Language(language).Event(event).Execute()

Get event metrics as a time series

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
	resp, r, err := apiClient.EventsAPI.AnalyticsGetEventSeries(context.Background(), websiteId).StartAt(startAt).EndAt(endAt).Unit(unit).Timezone(timezone).Url(url).Referrer(referrer).Title(title).Os(os).Browser(browser).Device(device).Country(country).Region(region).City(city).Tag(tag).Host(host).Language(language).Event(event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.AnalyticsGetEventSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetEventSeries`: []AnalyticsPageviewSeries
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.AnalyticsGetEventSeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetEventSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 
 **unit** | **string** |  | 
 **timezone** | **string** |  | 
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

[**[]AnalyticsPageviewSeries**](AnalyticsPageviewSeries.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetEvents

> []map[string]interface{} AnalyticsGetEvents(ctx, websiteId).StartAt(startAt).EndAt(endAt).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

Get paginated list of events

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
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.AnalyticsGetEvents(context.Background(), websiteId).StartAt(startAt).EndAt(endAt).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.AnalyticsGetEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetEvents`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.AnalyticsGetEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **orderBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3GetBucketNotification

> S3GetBucketNotification200Response S3GetBucketNotification(ctx, bucket).Execute()

Get event notification config

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
	bucket := "bucket_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.S3GetBucketNotification(context.Background(), bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.S3GetBucketNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3GetBucketNotification`: S3GetBucketNotification200Response
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.S3GetBucketNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3GetBucketNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**S3GetBucketNotification200Response**](S3GetBucketNotification200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3PutBucketNotification

> S3PutBucketNotification(ctx, bucket).S3GetBucketNotification200Response(s3GetBucketNotification200Response).Execute()

Set event notification config

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
	bucket := "bucket_example" // string | 
	s3GetBucketNotification200Response := *openapiclient.NewS3GetBucketNotification200Response() // S3GetBucketNotification200Response | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.S3PutBucketNotification(context.Background(), bucket).S3GetBucketNotification200Response(s3GetBucketNotification200Response).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.S3PutBucketNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3PutBucketNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **s3GetBucketNotification200Response** | [**S3GetBucketNotification200Response**](S3GetBucketNotification200Response.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

