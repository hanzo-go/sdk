# \ReportsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsCreateReport**](ReportsAPI.md#AnalyticsCreateReport) | **Post** /v1/analytics/reports | Create a new report
[**AnalyticsDeleteReport**](ReportsAPI.md#AnalyticsDeleteReport) | **Delete** /v1/analytics/reports/{reportId} | Delete a report
[**AnalyticsGetReport**](ReportsAPI.md#AnalyticsGetReport) | **Get** /v1/analytics/reports/{reportId} | Get a report by ID
[**AnalyticsGetRevenueValues**](ReportsAPI.md#AnalyticsGetRevenueValues) | **Get** /v1/analytics/reports/revenue | Get available revenue values for a website
[**AnalyticsGetWebsiteReports**](ReportsAPI.md#AnalyticsGetWebsiteReports) | **Get** /v1/analytics/websites/{websiteId}/reports | List reports for a specific website
[**AnalyticsListReports**](ReportsAPI.md#AnalyticsListReports) | **Get** /v1/analytics/reports | List reports, optionally filtered by website or team
[**AnalyticsRunAttributionReport**](ReportsAPI.md#AnalyticsRunAttributionReport) | **Post** /v1/analytics/reports/attribution | Run an attribution report
[**AnalyticsRunFunnelReport**](ReportsAPI.md#AnalyticsRunFunnelReport) | **Post** /v1/analytics/reports/funnel | Run a funnel report
[**AnalyticsRunGoalsReport**](ReportsAPI.md#AnalyticsRunGoalsReport) | **Post** /v1/analytics/reports/goals | Run a goals report
[**AnalyticsRunInsightsReport**](ReportsAPI.md#AnalyticsRunInsightsReport) | **Post** /v1/analytics/reports/insights | Run an insights report
[**AnalyticsRunJourneyReport**](ReportsAPI.md#AnalyticsRunJourneyReport) | **Post** /v1/analytics/reports/journey | Run a user journey report
[**AnalyticsRunRetentionReport**](ReportsAPI.md#AnalyticsRunRetentionReport) | **Post** /v1/analytics/reports/retention | Run a retention report
[**AnalyticsRunRevenueReport**](ReportsAPI.md#AnalyticsRunRevenueReport) | **Post** /v1/analytics/reports/revenue | Run a revenue report
[**AnalyticsRunUtmReport**](ReportsAPI.md#AnalyticsRunUtmReport) | **Post** /v1/analytics/reports/utm | Run a UTM report
[**AnalyticsUpdateReport**](ReportsAPI.md#AnalyticsUpdateReport) | **Post** /v1/analytics/reports/{reportId} | Update a report



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
	resp, r, err := apiClient.ReportsAPI.AnalyticsCreateReport(context.Background()).AnalyticsCreateReportRequest(analyticsCreateReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AnalyticsCreateReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsCreateReport`: AnalyticsReport
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AnalyticsCreateReport`: %v\n", resp)
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
	resp, r, err := apiClient.ReportsAPI.AnalyticsDeleteReport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AnalyticsDeleteReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsDeleteReport`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AnalyticsDeleteReport`: %v\n", resp)
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
	resp, r, err := apiClient.ReportsAPI.AnalyticsGetReport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AnalyticsGetReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetReport`: AnalyticsReport
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AnalyticsGetReport`: %v\n", resp)
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
	resp, r, err := apiClient.ReportsAPI.AnalyticsGetRevenueValues(context.Background()).WebsiteId(websiteId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AnalyticsGetRevenueValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetRevenueValues`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AnalyticsGetRevenueValues`: %v\n", resp)
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
	resp, r, err := apiClient.ReportsAPI.AnalyticsGetWebsiteReports(context.Background(), websiteId).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AnalyticsGetWebsiteReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsGetWebsiteReports`: []AnalyticsReport
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AnalyticsGetWebsiteReports`: %v\n", resp)
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
	resp, r, err := apiClient.ReportsAPI.AnalyticsListReports(context.Background()).WebsiteId(websiteId).TeamId(teamId).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AnalyticsListReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsListReports`: []AnalyticsReport
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AnalyticsListReports`: %v\n", resp)
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
	resp, r, err := apiClient.ReportsAPI.AnalyticsRunAttributionReport(context.Background()).AnalyticsRunAttributionReportRequest(analyticsRunAttributionReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AnalyticsRunAttributionReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRunAttributionReport`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AnalyticsRunAttributionReport`: %v\n", resp)
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
	resp, r, err := apiClient.ReportsAPI.AnalyticsRunFunnelReport(context.Background()).AnalyticsRunFunnelReportRequest(analyticsRunFunnelReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AnalyticsRunFunnelReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRunFunnelReport`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AnalyticsRunFunnelReport`: %v\n", resp)
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
	resp, r, err := apiClient.ReportsAPI.AnalyticsRunGoalsReport(context.Background()).AnalyticsRunGoalsReportRequest(analyticsRunGoalsReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AnalyticsRunGoalsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRunGoalsReport`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AnalyticsRunGoalsReport`: %v\n", resp)
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
	resp, r, err := apiClient.ReportsAPI.AnalyticsRunInsightsReport(context.Background()).AnalyticsRunInsightsReportRequest(analyticsRunInsightsReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AnalyticsRunInsightsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRunInsightsReport`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AnalyticsRunInsightsReport`: %v\n", resp)
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
	resp, r, err := apiClient.ReportsAPI.AnalyticsRunJourneyReport(context.Background()).AnalyticsRunJourneyReportRequest(analyticsRunJourneyReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AnalyticsRunJourneyReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRunJourneyReport`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AnalyticsRunJourneyReport`: %v\n", resp)
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
	resp, r, err := apiClient.ReportsAPI.AnalyticsRunRetentionReport(context.Background()).AnalyticsRunRetentionReportRequest(analyticsRunRetentionReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AnalyticsRunRetentionReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRunRetentionReport`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AnalyticsRunRetentionReport`: %v\n", resp)
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
	resp, r, err := apiClient.ReportsAPI.AnalyticsRunRevenueReport(context.Background()).AnalyticsRunRevenueReportRequest(analyticsRunRevenueReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AnalyticsRunRevenueReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRunRevenueReport`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AnalyticsRunRevenueReport`: %v\n", resp)
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
	resp, r, err := apiClient.ReportsAPI.AnalyticsRunUtmReport(context.Background()).AnalyticsRunUtmReportRequest(analyticsRunUtmReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AnalyticsRunUtmReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRunUtmReport`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AnalyticsRunUtmReport`: %v\n", resp)
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
	resp, r, err := apiClient.ReportsAPI.AnalyticsUpdateReport(context.Background(), reportId).AnalyticsCreateReportRequest(analyticsCreateReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.AnalyticsUpdateReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsUpdateReport`: AnalyticsReport
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.AnalyticsUpdateReport`: %v\n", resp)
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

