# \WebsitesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsCreateWebsite**](WebsitesAPI.md#AnalyticsCreateWebsite) | **Post** /v1/analytics/websites | Create a new website
[**AnalyticsDeleteWebsite**](WebsitesAPI.md#AnalyticsDeleteWebsite) | **Delete** /v1/analytics/websites/{websiteId} | Delete website
[**AnalyticsGetActiveVisitors**](WebsitesAPI.md#AnalyticsGetActiveVisitors) | **Get** /v1/analytics/websites/{websiteId}/active | Get count of active visitors
[**AnalyticsGetWebsite**](WebsitesAPI.md#AnalyticsGetWebsite) | **Get** /v1/analytics/websites/{websiteId} | Get website by ID
[**AnalyticsGetWebsiteDateRange**](WebsitesAPI.md#AnalyticsGetWebsiteDateRange) | **Get** /v1/analytics/websites/{websiteId}/daterange | Get the date range of data available for a website
[**AnalyticsGetWebsiteValues**](WebsitesAPI.md#AnalyticsGetWebsiteValues) | **Get** /v1/analytics/websites/{websiteId}/values | Get distinct values for a given column type
[**AnalyticsListWebsites**](WebsitesAPI.md#AnalyticsListWebsites) | **Get** /v1/analytics/websites | List websites owned by current user
[**AnalyticsResetWebsite**](WebsitesAPI.md#AnalyticsResetWebsite) | **Post** /v1/analytics/websites/{websiteId}/reset | Reset all data for a website
[**AnalyticsTransferWebsite**](WebsitesAPI.md#AnalyticsTransferWebsite) | **Post** /v1/analytics/websites/{websiteId}/transfer | Transfer website ownership to another user or team
[**AnalyticsUpdateWebsite**](WebsitesAPI.md#AnalyticsUpdateWebsite) | **Post** /v1/analytics/websites/{websiteId} | Update website



## AnalyticsCreateWebsite

> AnalyticsWebsite AnalyticsCreateWebsite(ctx).AnalyticsCreateWebsiteRequest(analyticsCreateWebsiteRequest).Execute()

Create a new website

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
	analyticsCreateWebsiteRequest := *openapiclient.NewAnalyticsCreateWebsiteRequest("Name_example", "Domain_example") // AnalyticsCreateWebsiteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebsitesAPI.AnalyticsCreateWebsite(context.Background()).AnalyticsCreateWebsiteRequest(analyticsCreateWebsiteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsitesAPI.AnalyticsCreateWebsite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsCreateWebsite`: AnalyticsWebsite
	fmt.Fprintf(os.Stdout, "Response from `WebsitesAPI.AnalyticsCreateWebsite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsCreateWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsCreateWebsiteRequest** | [**AnalyticsCreateWebsiteRequest**](AnalyticsCreateWebsiteRequest.md) |  | 

### Return type

[**AnalyticsWebsite**](AnalyticsWebsite.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsDeleteWebsite

> map[string]interface{} AnalyticsDeleteWebsite(ctx, websiteId).Execute()

Delete website

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebsitesAPI.AnalyticsDeleteWebsite(context.Background(), websiteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsitesAPI.AnalyticsDeleteWebsite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsDeleteWebsite`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebsitesAPI.AnalyticsDeleteWebsite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsDeleteWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetActiveVisitors

> AnalyticsGetActiveVisitors200Response AnalyticsGetActiveVisitors(ctx, websiteId).Execute()

Get count of active visitors

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebsitesAPI.AnalyticsGetActiveVisitors(context.Background(), websiteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsitesAPI.AnalyticsGetActiveVisitors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetActiveVisitors`: AnalyticsGetActiveVisitors200Response
	fmt.Fprintf(os.Stdout, "Response from `WebsitesAPI.AnalyticsGetActiveVisitors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetActiveVisitorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalyticsGetActiveVisitors200Response**](AnalyticsGetActiveVisitors200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetWebsite

> AnalyticsWebsite AnalyticsGetWebsite(ctx, websiteId).Execute()

Get website by ID

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebsitesAPI.AnalyticsGetWebsite(context.Background(), websiteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsitesAPI.AnalyticsGetWebsite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetWebsite`: AnalyticsWebsite
	fmt.Fprintf(os.Stdout, "Response from `WebsitesAPI.AnalyticsGetWebsite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalyticsWebsite**](AnalyticsWebsite.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetWebsiteDateRange

> AnalyticsGetWebsiteDateRange200Response AnalyticsGetWebsiteDateRange(ctx, websiteId).Execute()

Get the date range of data available for a website

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebsitesAPI.AnalyticsGetWebsiteDateRange(context.Background(), websiteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsitesAPI.AnalyticsGetWebsiteDateRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetWebsiteDateRange`: AnalyticsGetWebsiteDateRange200Response
	fmt.Fprintf(os.Stdout, "Response from `WebsitesAPI.AnalyticsGetWebsiteDateRange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetWebsiteDateRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalyticsGetWebsiteDateRange200Response**](AnalyticsGetWebsiteDateRange200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetWebsiteValues

> []string AnalyticsGetWebsiteValues(ctx, websiteId).StartAt(startAt).EndAt(endAt).Type_(type_).Search(search).Execute()

Get distinct values for a given column type

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
	type_ := "type__example" // string | Column type (e.g. browser, os, device, country, url, referrer, title, event, tag, language, region, city, host)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebsitesAPI.AnalyticsGetWebsiteValues(context.Background(), websiteId).StartAt(startAt).EndAt(endAt).Type_(type_).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsitesAPI.AnalyticsGetWebsiteValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetWebsiteValues`: []string
	fmt.Fprintf(os.Stdout, "Response from `WebsitesAPI.AnalyticsGetWebsiteValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetWebsiteValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | Start timestamp in milliseconds | 
 **endAt** | **int64** | End timestamp in milliseconds | 
 **type_** | **string** | Column type (e.g. browser, os, device, country, url, referrer, title, event, tag, language, region, city, host) | 
 **search** | **string** |  | 

### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsListWebsites

> []AnalyticsWebsite AnalyticsListWebsites(ctx).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

List websites owned by current user

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
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebsitesAPI.AnalyticsListWebsites(context.Background()).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsitesAPI.AnalyticsListWebsites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsListWebsites`: []AnalyticsWebsite
	fmt.Fprintf(os.Stdout, "Response from `WebsitesAPI.AnalyticsListWebsites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsListWebsitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **orderBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]AnalyticsWebsite**](AnalyticsWebsite.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsResetWebsite

> map[string]interface{} AnalyticsResetWebsite(ctx, websiteId).Execute()

Reset all data for a website

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebsitesAPI.AnalyticsResetWebsite(context.Background(), websiteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsitesAPI.AnalyticsResetWebsite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsResetWebsite`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebsitesAPI.AnalyticsResetWebsite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsResetWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsTransferWebsite

> AnalyticsWebsite AnalyticsTransferWebsite(ctx, websiteId).AnalyticsTransferWebsiteRequest(analyticsTransferWebsiteRequest).Execute()

Transfer website ownership to another user or team

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
	analyticsTransferWebsiteRequest := *openapiclient.NewAnalyticsTransferWebsiteRequest() // AnalyticsTransferWebsiteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebsitesAPI.AnalyticsTransferWebsite(context.Background(), websiteId).AnalyticsTransferWebsiteRequest(analyticsTransferWebsiteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsitesAPI.AnalyticsTransferWebsite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsTransferWebsite`: AnalyticsWebsite
	fmt.Fprintf(os.Stdout, "Response from `WebsitesAPI.AnalyticsTransferWebsite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsTransferWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **analyticsTransferWebsiteRequest** | [**AnalyticsTransferWebsiteRequest**](AnalyticsTransferWebsiteRequest.md) |  | 

### Return type

[**AnalyticsWebsite**](AnalyticsWebsite.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsUpdateWebsite

> AnalyticsWebsite AnalyticsUpdateWebsite(ctx, websiteId).AnalyticsUpdateWebsiteRequest(analyticsUpdateWebsiteRequest).Execute()

Update website

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
	analyticsUpdateWebsiteRequest := *openapiclient.NewAnalyticsUpdateWebsiteRequest("Name_example", "Domain_example") // AnalyticsUpdateWebsiteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebsitesAPI.AnalyticsUpdateWebsite(context.Background(), websiteId).AnalyticsUpdateWebsiteRequest(analyticsUpdateWebsiteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsitesAPI.AnalyticsUpdateWebsite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsUpdateWebsite`: AnalyticsWebsite
	fmt.Fprintf(os.Stdout, "Response from `WebsitesAPI.AnalyticsUpdateWebsite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsUpdateWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **analyticsUpdateWebsiteRequest** | [**AnalyticsUpdateWebsiteRequest**](AnalyticsUpdateWebsiteRequest.md) |  | 

### Return type

[**AnalyticsWebsite**](AnalyticsWebsite.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

