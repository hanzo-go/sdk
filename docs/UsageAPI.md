# \UsageAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsageAnalytics**](UsageAPI.md#GetUsageAnalytics) | **Get** /v1/usage/analytics | Is the entitlement-GATED per-provider breakdown of the caller org&#39;s LLM usage — the paid lens over the same warehouse ledger GET /v1/usage/summary reads its totals from.
[**GetUsageAnalyticsAccess**](UsageAPI.md#GetUsageAnalyticsAccess) | **Get** /v1/usage/analytics/access | Echoes a plan&#39;s resolved analytics entitlement so a dashboard can configure itself against the LIVE catalog instead of hardcoding tier numbers.
[**GetUsageSamples**](UsageAPI.md#GetUsageSamples) | **Get** /v1/usage/samples | Is the PER-PROVIDER view: one connected account&#39;s own consumption of its own plan — \&quot;my plan is 47% through its 6h window, resets at 14:20\&quot;.
[**GetUsageSummary**](UsageAPI.md#GetUsageSummary) | **Get** /v1/usage/summary | Answers GET /v1/usage/summary: the caller&#39;s own usage footprint over one window — the categorized spend roll-up from the commerce ledger, the org&#39;s LLM usage totals from the warehouse, and the caller&#39;s OWN linked provider accounts beside the org&#39;s Hanzo-routed usage.
[**PostUsage**](UsageAPI.md#PostUsage) | **Post** /v1/usage | Ingests a batch of account-usage samples — what a developer&#39;s OWN AI accounts have consumed of their OWN plans, metered from each provider&#39;s own login — and appends them to the warehouse series.



## GetUsageAnalytics

> UsageAnalyticsView GetUsageAnalytics(ctx).End(end).Plan(plan).Range_(range_).Start(start).Execute()

Is the entitlement-GATED per-provider breakdown of the caller org's LLM usage — the paid lens over the same warehouse ledger GET /v1/usage/summary reads its totals from.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	end := "end_example" // string | End is the exclusive window end, RFC3339. Read only when Range is custom. (optional)
	plan := "plan_example" // string | Plan is the plan id whose entitlement decides access and retention. INTERIM: cloud has no org-to-plan resolver yet, so the caller names the plan; when that resolver lands this becomes the caller org's own plan. (optional)
	range_ := "range__example" // string | Range is the window: a count and a unit — 24h, 7d, 90d, any <N>h or <N>d — or day, week, month, all, custom. Empty means 24h. The window is then clamped forward to the plan's retention entitlement. (optional)
	start := "start_example" // string | Start is the inclusive window start, RFC3339. Read only when Range is custom, and clamped forward to the plan's retention floor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.GetUsageAnalytics(context.Background()).End(end).Plan(plan).Range_(range_).Start(start).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetUsageAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageAnalytics`: UsageAnalyticsView
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.GetUsageAnalytics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **end** | **string** | End is the exclusive window end, RFC3339. Read only when Range is custom. | 
 **plan** | **string** | Plan is the plan id whose entitlement decides access and retention. INTERIM: cloud has no org-to-plan resolver yet, so the caller names the plan; when that resolver lands this becomes the caller org&#39;s own plan. | 
 **range_** | **string** | Range is the window: a count and a unit — 24h, 7d, 90d, any &lt;N&gt;h or &lt;N&gt;d — or day, week, month, all, custom. Empty means 24h. The window is then clamped forward to the plan&#39;s retention entitlement. | 
 **start** | **string** | Start is the inclusive window start, RFC3339. Read only when Range is custom, and clamped forward to the plan&#39;s retention floor. | 

### Return type

[**UsageAnalyticsView**](UsageAnalyticsView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageAnalyticsAccess

> UsageAnalyticsAccess GetUsageAnalyticsAccess(ctx).Plan(plan).Execute()

Echoes a plan's resolved analytics entitlement so a dashboard can configure itself against the LIVE catalog instead of hardcoding tier numbers.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	plan := "plan_example" // string | Plan is a plan id from the live @hanzo/plans catalog. Empty resolves the free floor, and so does an id the catalog does not know — this never fails on an unknown plan. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.GetUsageAnalyticsAccess(context.Background()).Plan(plan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetUsageAnalyticsAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageAnalyticsAccess`: UsageAnalyticsAccess
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.GetUsageAnalyticsAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageAnalyticsAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plan** | **string** | Plan is a plan id from the live @hanzo/plans catalog. Empty resolves the free floor, and so does an id the catalog does not know — this never fails on an unknown plan. | 

### Return type

[**UsageAnalyticsAccess**](UsageAnalyticsAccess.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageSamples

> DashResp GetUsageSamples(ctx).Account(account).Provider(provider).Range_(range_).Window(window).Execute()

Is the PER-PROVIDER view: one connected account's own consumption of its own plan — \"my plan is 47% through its 6h window, resets at 14:20\".



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	account := "account_example" // string | Account narrows to ONE linked account of that provider. Empty covers every account the caller has linked there. (optional)
	provider := "provider_example" // string | Provider is the upstream to read, e.g. anthropic. Required. (optional)
	range_ := "range__example" // string | Range is the window to read: a count and a unit — 1h, 24h, 90d, any <N>h or <N>d — or day, week, month, all. Empty means 24h. A label that is not a count, or one reaching past the 730-day horizon, is refused rather than silently replaced. (optional)
	window := "window_example" // string | Window narrows to ONE window class: 6h, day, week or month. Empty covers every class. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.GetUsageSamples(context.Background()).Account(account).Provider(provider).Range_(range_).Window(window).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetUsageSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageSamples`: DashResp
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.GetUsageSamples`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | Account narrows to ONE linked account of that provider. Empty covers every account the caller has linked there. | 
 **provider** | **string** | Provider is the upstream to read, e.g. anthropic. Required. | 
 **range_** | **string** | Range is the window to read: a count and a unit — 1h, 24h, 90d, any &lt;N&gt;h or &lt;N&gt;d — or day, week, month, all. Empty means 24h. A label that is not a count, or one reaching past the 730-day horizon, is refused rather than silently replaced. | 
 **window** | **string** | Window narrows to ONE window class: 6h, day, week or month. Empty covers every class. | 

### Return type

[**DashResp**](DashResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageSummary

> UsageSummary GetUsageSummary(ctx).Range_(range_).Start(start).End(end).Execute()

Answers GET /v1/usage/summary: the caller's own usage footprint over one window — the categorized spend roll-up from the commerce ledger, the org's LLM usage totals from the warehouse, and the caller's OWN linked provider accounts beside the org's Hanzo-routed usage.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	range_ := "range__example" // string | Range is the window: a count and a unit — 24h, 7d, 90d, any <N>h or <N>d — or day, week, month, all, custom. Empty means 24h. A label this surface does not know, or one reaching past the 730-day horizon, is refused rather than silently replaced. (optional)
	start := "start_example" // string | Start is the inclusive window start, RFC3339. Read only when Range is custom. (optional)
	end := "end_example" // string | End is the exclusive window end, RFC3339. Read only when Range is custom. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.GetUsageSummary(context.Background()).Range_(range_).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetUsageSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageSummary`: UsageSummary
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.GetUsageSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is the window: a count and a unit — 24h, 7d, 90d, any &lt;N&gt;h or &lt;N&gt;d — or day, week, month, all, custom. Empty means 24h. A label this surface does not know, or one reaching past the 730-day horizon, is refused rather than silently replaced. | 
 **start** | **string** | Start is the inclusive window start, RFC3339. Read only when Range is custom. | 
 **end** | **string** | End is the exclusive window end, RFC3339. Read only when Range is custom. | 

### Return type

[**UsageSummary**](UsageSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsage

> ReportResp PostUsage(ctx).ReportReq(reportReq).Execute()

Ingests a batch of account-usage samples — what a developer's OWN AI accounts have consumed of their OWN plans, metered from each provider's own login — and appends them to the warehouse series.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	reportReq := *openapiclient.NewReportReq() // ReportReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.PostUsage(context.Background()).ReportReq(reportReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.PostUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsage`: ReportResp
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.PostUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportReq** | [**ReportReq**](ReportReq.md) |  | 

### Return type

[**ReportResp**](ReportResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

