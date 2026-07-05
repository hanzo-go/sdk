# \AnalyticsReportsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsCreateReport**](AnalyticsReportsAPI.md#AnalyticsCreateReport) | **Post** /v1/analytics/reports | Create a new report
[**AnalyticsDeleteReport**](AnalyticsReportsAPI.md#AnalyticsDeleteReport) | **Delete** /v1/analytics/reports/{reportId} | Delete a report
[**AnalyticsGetReport**](AnalyticsReportsAPI.md#AnalyticsGetReport) | **Get** /v1/analytics/reports/{reportId} | Get a report by ID
[**AnalyticsGetRevenueValues**](AnalyticsReportsAPI.md#AnalyticsGetRevenueValues) | **Get** /v1/analytics/reports/revenue | Get available revenue values for a website
[**AnalyticsGetWebsiteReports**](AnalyticsReportsAPI.md#AnalyticsGetWebsiteReports) | **Get** /v1/analytics/websites/{websiteId}/reports | List reports for a specific website
[**AnalyticsListReports**](AnalyticsReportsAPI.md#AnalyticsListReports) | **Get** /v1/analytics/reports | List reports, optionally filtered by website or team
[**AnalyticsRunAttributionReport**](AnalyticsReportsAPI.md#AnalyticsRunAttributionReport) | **Post** /v1/analytics/reports/attribution | Run an attribution report
[**AnalyticsRunFunnelReport**](AnalyticsReportsAPI.md#AnalyticsRunFunnelReport) | **Post** /v1/analytics/reports/funnel | Run a funnel report
[**AnalyticsRunGoalsReport**](AnalyticsReportsAPI.md#AnalyticsRunGoalsReport) | **Post** /v1/analytics/reports/goals | Run a goals report
[**AnalyticsRunInsightsReport**](AnalyticsReportsAPI.md#AnalyticsRunInsightsReport) | **Post** /v1/analytics/reports/insights | Run an insights report
[**AnalyticsRunJourneyReport**](AnalyticsReportsAPI.md#AnalyticsRunJourneyReport) | **Post** /v1/analytics/reports/journey | Run a user journey report
[**AnalyticsRunRetentionReport**](AnalyticsReportsAPI.md#AnalyticsRunRetentionReport) | **Post** /v1/analytics/reports/retention | Run a retention report
[**AnalyticsRunRevenueReport**](AnalyticsReportsAPI.md#AnalyticsRunRevenueReport) | **Post** /v1/analytics/reports/revenue | Run a revenue report
[**AnalyticsRunUtmReport**](AnalyticsReportsAPI.md#AnalyticsRunUtmReport) | **Post** /v1/analytics/reports/utm | Run a UTM report
[**AnalyticsUpdateReport**](AnalyticsReportsAPI.md#AnalyticsUpdateReport) | **Post** /v1/analytics/reports/{reportId} | Update a report



## AnalyticsCreateReport

> AnalyticsReport AnalyticsCreateReport(ctx).AnalyticsCreateReportRequest(analyticsCreateReportRequest).Execute()

Create a new report

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
	analyticsCreateReportRequest := *openapiclient.NewAnalyticsCreateReportRequest("WebsiteId_example", "Name_example", "Type_example", "Description_example", map[string]interface{}(123)) // AnalyticsCreateReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportsAPI.AnalyticsCreateReport(context.Background()).AnalyticsCreateReportRequest(analyticsCreateReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportsAPI.AnalyticsCreateReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsCreateReport`: AnalyticsReport
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportsAPI.AnalyticsCreateReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsCreateReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsCreateReportRequest** | [**AnalyticsCreateReportRequest**](AnalyticsCreateReportRequest.md) |  | 

### Return type

[**AnalyticsReport**](AnalyticsReport.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsDeleteReport

> map[string]interface{} AnalyticsDeleteReport(ctx, reportId).Execute()

Delete a report

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
	reportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportsAPI.AnalyticsDeleteReport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportsAPI.AnalyticsDeleteReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsDeleteReport`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportsAPI.AnalyticsDeleteReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsDeleteReportRequest struct via the builder pattern


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


## AnalyticsGetReport

> AnalyticsReport AnalyticsGetReport(ctx, reportId).Execute()

Get a report by ID

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
	reportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportsAPI.AnalyticsGetReport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportsAPI.AnalyticsGetReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetReport`: AnalyticsReport
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportsAPI.AnalyticsGetReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalyticsReport**](AnalyticsReport.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsGetRevenueValues

> []map[string]interface{} AnalyticsGetRevenueValues(ctx).WebsiteId(websiteId).StartDate(startDate).EndDate(endDate).Execute()

Get available revenue values for a website

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
	websiteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	startDate := time.Now() // time.Time | 
	endDate := time.Now() // time.Time | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportsAPI.AnalyticsGetRevenueValues(context.Background()).WebsiteId(websiteId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportsAPI.AnalyticsGetRevenueValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetRevenueValues`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportsAPI.AnalyticsGetRevenueValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetRevenueValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **websiteId** | **string** |  | 
 **startDate** | **time.Time** |  | 
 **endDate** | **time.Time** |  | 

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


## AnalyticsGetWebsiteReports

> []AnalyticsReport AnalyticsGetWebsiteReports(ctx, websiteId).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

List reports for a specific website

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
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportsAPI.AnalyticsGetWebsiteReports(context.Background(), websiteId).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportsAPI.AnalyticsGetWebsiteReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetWebsiteReports`: []AnalyticsReport
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportsAPI.AnalyticsGetWebsiteReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsGetWebsiteReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **orderBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]AnalyticsReport**](AnalyticsReport.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsListReports

> []AnalyticsReport AnalyticsListReports(ctx).WebsiteId(websiteId).TeamId(teamId).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

List reports, optionally filtered by website or team

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
	websiteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	teamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportsAPI.AnalyticsListReports(context.Background()).WebsiteId(websiteId).TeamId(teamId).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportsAPI.AnalyticsListReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsListReports`: []AnalyticsReport
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportsAPI.AnalyticsListReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsListReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **websiteId** | **string** |  | 
 **teamId** | **string** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **orderBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]AnalyticsReport**](AnalyticsReport.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsRunAttributionReport

> []map[string]interface{} AnalyticsRunAttributionReport(ctx).AnalyticsRunAttributionReportRequest(analyticsRunAttributionReportRequest).Execute()

Run an attribution report

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
	analyticsRunAttributionReportRequest := *openapiclient.NewAnalyticsRunAttributionReportRequest("WebsiteId_example", *openapiclient.NewAnalyticsDateRange(time.Now(), time.Now()), "Model_example", []openapiclient.AnalyticsRunAttributionReportRequestStepsInner{*openapiclient.NewAnalyticsRunAttributionReportRequestStepsInner("Type_example", "Value_example")}) // AnalyticsRunAttributionReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportsAPI.AnalyticsRunAttributionReport(context.Background()).AnalyticsRunAttributionReportRequest(analyticsRunAttributionReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportsAPI.AnalyticsRunAttributionReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRunAttributionReport`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportsAPI.AnalyticsRunAttributionReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsRunAttributionReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsRunAttributionReportRequest** | [**AnalyticsRunAttributionReportRequest**](AnalyticsRunAttributionReportRequest.md) |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsRunFunnelReport

> []map[string]interface{} AnalyticsRunFunnelReport(ctx).AnalyticsRunFunnelReportRequest(analyticsRunFunnelReportRequest).Execute()

Run a funnel report

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
	analyticsRunFunnelReportRequest := *openapiclient.NewAnalyticsRunFunnelReportRequest("WebsiteId_example", *openapiclient.NewAnalyticsDateRange(time.Now(), time.Now()), float32(123), []openapiclient.AnalyticsRunAttributionReportRequestStepsInner{*openapiclient.NewAnalyticsRunAttributionReportRequestStepsInner("Type_example", "Value_example")}) // AnalyticsRunFunnelReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportsAPI.AnalyticsRunFunnelReport(context.Background()).AnalyticsRunFunnelReportRequest(analyticsRunFunnelReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportsAPI.AnalyticsRunFunnelReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRunFunnelReport`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportsAPI.AnalyticsRunFunnelReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsRunFunnelReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsRunFunnelReportRequest** | [**AnalyticsRunFunnelReportRequest**](AnalyticsRunFunnelReportRequest.md) |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsRunGoalsReport

> []map[string]interface{} AnalyticsRunGoalsReport(ctx).AnalyticsRunGoalsReportRequest(analyticsRunGoalsReportRequest).Execute()

Run a goals report

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
	analyticsRunGoalsReportRequest := *openapiclient.NewAnalyticsRunGoalsReportRequest("WebsiteId_example", *openapiclient.NewAnalyticsDateRange(time.Now(), time.Now()), []openapiclient.AnalyticsRunGoalsReportRequestGoalsInner{*openapiclient.NewAnalyticsRunGoalsReportRequestGoalsInner("Type_example", "Value_example", float32(123))}) // AnalyticsRunGoalsReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportsAPI.AnalyticsRunGoalsReport(context.Background()).AnalyticsRunGoalsReportRequest(analyticsRunGoalsReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportsAPI.AnalyticsRunGoalsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRunGoalsReport`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportsAPI.AnalyticsRunGoalsReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsRunGoalsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsRunGoalsReportRequest** | [**AnalyticsRunGoalsReportRequest**](AnalyticsRunGoalsReportRequest.md) |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsRunInsightsReport

> []map[string]interface{} AnalyticsRunInsightsReport(ctx).AnalyticsRunInsightsReportRequest(analyticsRunInsightsReportRequest).Execute()

Run an insights report

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
	analyticsRunInsightsReportRequest := *openapiclient.NewAnalyticsRunInsightsReportRequest("WebsiteId_example", *openapiclient.NewAnalyticsDateRange(time.Now(), time.Now()), []openapiclient.AnalyticsRunInsightsReportRequestFieldsInner{*openapiclient.NewAnalyticsRunInsightsReportRequestFieldsInner("Name_example", "Type_example", "Label_example")}, []openapiclient.AnalyticsRunInsightsReportRequestFiltersInner{*openapiclient.NewAnalyticsRunInsightsReportRequestFiltersInner("Name_example", "Type_example", "Operator_example", "Value_example")}) // AnalyticsRunInsightsReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportsAPI.AnalyticsRunInsightsReport(context.Background()).AnalyticsRunInsightsReportRequest(analyticsRunInsightsReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportsAPI.AnalyticsRunInsightsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRunInsightsReport`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportsAPI.AnalyticsRunInsightsReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsRunInsightsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsRunInsightsReportRequest** | [**AnalyticsRunInsightsReportRequest**](AnalyticsRunInsightsReportRequest.md) |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsRunJourneyReport

> []map[string]interface{} AnalyticsRunJourneyReport(ctx).AnalyticsRunJourneyReportRequest(analyticsRunJourneyReportRequest).Execute()

Run a user journey report

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
	analyticsRunJourneyReportRequest := *openapiclient.NewAnalyticsRunJourneyReportRequest("WebsiteId_example", *openapiclient.NewAnalyticsDateRange(time.Now(), time.Now()), int32(123)) // AnalyticsRunJourneyReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportsAPI.AnalyticsRunJourneyReport(context.Background()).AnalyticsRunJourneyReportRequest(analyticsRunJourneyReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportsAPI.AnalyticsRunJourneyReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRunJourneyReport`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportsAPI.AnalyticsRunJourneyReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsRunJourneyReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsRunJourneyReportRequest** | [**AnalyticsRunJourneyReportRequest**](AnalyticsRunJourneyReportRequest.md) |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsRunRetentionReport

> []map[string]interface{} AnalyticsRunRetentionReport(ctx).AnalyticsRunRetentionReportRequest(analyticsRunRetentionReportRequest).Execute()

Run a retention report

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
	analyticsRunRetentionReportRequest := *openapiclient.NewAnalyticsRunRetentionReportRequest("WebsiteId_example", *openapiclient.NewAnalyticsDateRange(time.Now(), time.Now()), "Timezone_example") // AnalyticsRunRetentionReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportsAPI.AnalyticsRunRetentionReport(context.Background()).AnalyticsRunRetentionReportRequest(analyticsRunRetentionReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportsAPI.AnalyticsRunRetentionReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRunRetentionReport`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportsAPI.AnalyticsRunRetentionReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsRunRetentionReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsRunRetentionReportRequest** | [**AnalyticsRunRetentionReportRequest**](AnalyticsRunRetentionReportRequest.md) |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsRunRevenueReport

> []map[string]interface{} AnalyticsRunRevenueReport(ctx).AnalyticsRunRevenueReportRequest(analyticsRunRevenueReportRequest).Execute()

Run a revenue report

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
	analyticsRunRevenueReportRequest := *openapiclient.NewAnalyticsRunRevenueReportRequest("WebsiteId_example", *openapiclient.NewAnalyticsDateRange(time.Now(), time.Now()), "Currency_example", "Timezone_example") // AnalyticsRunRevenueReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportsAPI.AnalyticsRunRevenueReport(context.Background()).AnalyticsRunRevenueReportRequest(analyticsRunRevenueReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportsAPI.AnalyticsRunRevenueReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRunRevenueReport`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportsAPI.AnalyticsRunRevenueReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsRunRevenueReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsRunRevenueReportRequest** | [**AnalyticsRunRevenueReportRequest**](AnalyticsRunRevenueReportRequest.md) |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsRunUtmReport

> []map[string]interface{} AnalyticsRunUtmReport(ctx).AnalyticsRunUtmReportRequest(analyticsRunUtmReportRequest).Execute()

Run a UTM report

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
	analyticsRunUtmReportRequest := *openapiclient.NewAnalyticsRunUtmReportRequest("WebsiteId_example", *openapiclient.NewAnalyticsDateRange(time.Now(), time.Now())) // AnalyticsRunUtmReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportsAPI.AnalyticsRunUtmReport(context.Background()).AnalyticsRunUtmReportRequest(analyticsRunUtmReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportsAPI.AnalyticsRunUtmReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRunUtmReport`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportsAPI.AnalyticsRunUtmReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsRunUtmReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analyticsRunUtmReportRequest** | [**AnalyticsRunUtmReportRequest**](AnalyticsRunUtmReportRequest.md) |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsUpdateReport

> AnalyticsReport AnalyticsUpdateReport(ctx, reportId).AnalyticsCreateReportRequest(analyticsCreateReportRequest).Execute()

Update a report

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
	reportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	analyticsCreateReportRequest := *openapiclient.NewAnalyticsCreateReportRequest("WebsiteId_example", "Name_example", "Type_example", "Description_example", map[string]interface{}(123)) // AnalyticsCreateReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsReportsAPI.AnalyticsUpdateReport(context.Background(), reportId).AnalyticsCreateReportRequest(analyticsCreateReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsReportsAPI.AnalyticsUpdateReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsUpdateReport`: AnalyticsReport
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsReportsAPI.AnalyticsUpdateReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsUpdateReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **analyticsCreateReportRequest** | [**AnalyticsCreateReportRequest**](AnalyticsCreateReportRequest.md) |  | 

### Return type

[**AnalyticsReport**](AnalyticsReport.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

