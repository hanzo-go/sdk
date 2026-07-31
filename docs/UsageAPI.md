# \UsageAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1UsageActivity**](UsageAPI.md#CloudGetV1UsageActivity) | **Get** /v1/usage/activity | Activity returns the per-day usage series for ONE authorized subject — the points a contribution heatmap and a timeline are drawn from, gap-filled so every day in the range is present.
[**CloudGetV1UsageAnalytics**](UsageAPI.md#CloudGetV1UsageAnalytics) | **Get** /v1/usage/analytics | Is the entitlement-GATED per-provider breakdown of the caller org&#39;s LLM usage — the paid lens over the same warehouse ledger GET /v1/usage/summary reads its totals from.
[**CloudGetV1UsageAnalyticsAccess**](UsageAPI.md#CloudGetV1UsageAnalyticsAccess) | **Get** /v1/usage/analytics/access | Echoes a plan&#39;s resolved analytics entitlement so a dashboard can configure itself against the LIVE catalog instead of hardcoding tier numbers.
[**CloudGetV1UsageLeaderboard**](UsageAPI.md#CloudGetV1UsageLeaderboard) | **Get** /v1/usage/leaderboard | Leaderboard ranks AI usage over a window, either the users of the caller&#39;s own org or organizations against each other, and always reports the caller&#39;s own standing even when it falls outside the returned page.
[**CloudGetV1UsageLeaderboardOptin**](UsageAPI.md#CloudGetV1UsageLeaderboardOptin) | **Get** /v1/usage/leaderboard/optin | GetOptin returns the caller&#39;s own public-listing preference and their org&#39;s, each with whether the caller may change it.
[**CloudGetV1UsageSamples**](UsageAPI.md#CloudGetV1UsageSamples) | **Get** /v1/usage/samples | Is the PER-PROVIDER view: one connected account&#39;s own consumption of its own plan — \&quot;my Claude Max plan is 47% through its 6h window, resets at 14:20\&quot;.
[**CloudGetV1UsageSummary**](UsageAPI.md#CloudGetV1UsageSummary) | **Get** /v1/usage/summary | Answers GET /v1/usage/summary: the caller&#39;s own usage footprint over one window — the categorized spend roll-up from the commerce ledger, the org&#39;s LLM usage totals from the warehouse, and the caller&#39;s OWN linked provider accounts beside the org&#39;s Hanzo-routed usage.
[**CloudPostV1Usage**](UsageAPI.md#CloudPostV1Usage) | **Post** /v1/usage | Ingests a batch of account-usage samples — what a developer&#39;s OWN AI accounts have consumed of their OWN plans, metered from each provider&#39;s own login — and appends them to the warehouse series.
[**CloudPostV1UsageRollupBackfill**](UsageAPI.md#CloudPostV1UsageRollupBackfill) | **Post** /v1/usage/rollup/backfill | Backfill seeds the derived usage rollup from ledger history — the rows written before the incremental view existed, which that view can never capture.
[**CloudPutV1UsageLeaderboardOptin**](UsageAPI.md#CloudPutV1UsageLeaderboardOptin) | **Put** /v1/usage/leaderboard/optin | PutUserOptin sets the CALLER&#39;s own public-listing preference on the leaderboard.
[**CloudPutV1UsageLeaderboardOptinOrg**](UsageAPI.md#CloudPutV1UsageLeaderboardOptinOrg) | **Put** /v1/usage/leaderboard/optin/org | PutOrgOptin sets the ORG&#39;s listing on the cross-org global board.



## CloudGetV1UsageActivity

> CloudActivityView CloudGetV1UsageActivity(ctx).Subject(subject).Id(id).From(from).To(to).Execute()

Activity returns the per-day usage series for ONE authorized subject — the points a contribution heatmap and a timeline are drawn from, gap-filled so every day in the range is present.



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
	subject := "user" // string | Subject is what the series is about: \"user\" (default), \"org\" or \"project\". (optional)
	id := "id_example" // string | ID names the subject within what the caller is entitled to see. Omitted (or \"me\") it is the caller themselves, or their own org. Another user requires org admin and must belong to the caller's org; another org requires a SuperAdmin. (optional)
	from := "2026-01-01" // string | From is the first day of the range, \"2006-01-02\". Defaults to 90 days back. (optional)
	to := "2026-03-31" // string | To is the last day of the range, \"2006-01-02\". Defaults to today; the span is clamped to 366 days. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.CloudGetV1UsageActivity(context.Background()).Subject(subject).Id(id).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.CloudGetV1UsageActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1UsageActivity`: CloudActivityView
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.CloudGetV1UsageActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1UsageActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subject** | **string** | Subject is what the series is about: \&quot;user\&quot; (default), \&quot;org\&quot; or \&quot;project\&quot;. | 
 **id** | **string** | ID names the subject within what the caller is entitled to see. Omitted (or \&quot;me\&quot;) it is the caller themselves, or their own org. Another user requires org admin and must belong to the caller&#39;s org; another org requires a SuperAdmin. | 
 **from** | **string** | From is the first day of the range, \&quot;2006-01-02\&quot;. Defaults to 90 days back. | 
 **to** | **string** | To is the last day of the range, \&quot;2006-01-02\&quot;. Defaults to today; the span is clamped to 366 days. | 

### Return type

[**CloudActivityView**](CloudActivityView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1UsageAnalytics

> CloudUsageAnalyticsView CloudGetV1UsageAnalytics(ctx).End(end).Plan(plan).Range_(range_).Start(start).Execute()

Is the entitlement-GATED per-provider breakdown of the caller org's LLM usage — the paid lens over the same warehouse ledger GET /v1/usage/summary reads its totals from.



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
	end := "end_example" // string | End is the exclusive window end, RFC3339. Read only when Range is custom. (optional)
	plan := "plan_example" // string | Plan is the plan id whose entitlement decides access and retention. INTERIM: cloud has no org-to-plan resolver yet, so the caller names the plan; when that resolver lands this becomes the caller org's own plan. (optional)
	range_ := "range__example" // string | Range is the window: 24h, 7d, 30d, or custom. Empty means 24h. (optional)
	start := "start_example" // string | Start is the inclusive window start, RFC3339. Read only when Range is custom, and clamped forward to the plan's retention floor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.CloudGetV1UsageAnalytics(context.Background()).End(end).Plan(plan).Range_(range_).Start(start).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.CloudGetV1UsageAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1UsageAnalytics`: CloudUsageAnalyticsView
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.CloudGetV1UsageAnalytics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1UsageAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **end** | **string** | End is the exclusive window end, RFC3339. Read only when Range is custom. | 
 **plan** | **string** | Plan is the plan id whose entitlement decides access and retention. INTERIM: cloud has no org-to-plan resolver yet, so the caller names the plan; when that resolver lands this becomes the caller org&#39;s own plan. | 
 **range_** | **string** | Range is the window: 24h, 7d, 30d, or custom. Empty means 24h. | 
 **start** | **string** | Start is the inclusive window start, RFC3339. Read only when Range is custom, and clamped forward to the plan&#39;s retention floor. | 

### Return type

[**CloudUsageAnalyticsView**](CloudUsageAnalyticsView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1UsageAnalyticsAccess

> CloudUsageAnalyticsAccess CloudGetV1UsageAnalyticsAccess(ctx).Plan(plan).Execute()

Echoes a plan's resolved analytics entitlement so a dashboard can configure itself against the LIVE catalog instead of hardcoding tier numbers.



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
	plan := "plan_example" // string | Plan is a plan id from the live @hanzo/plans catalog. Empty resolves the free floor, and so does an id the catalog does not know — this never fails on an unknown plan. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.CloudGetV1UsageAnalyticsAccess(context.Background()).Plan(plan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.CloudGetV1UsageAnalyticsAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1UsageAnalyticsAccess`: CloudUsageAnalyticsAccess
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.CloudGetV1UsageAnalyticsAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1UsageAnalyticsAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plan** | **string** | Plan is a plan id from the live @hanzo/plans catalog. Empty resolves the free floor, and so does an id the catalog does not know — this never fails on an unknown plan. | 

### Return type

[**CloudUsageAnalyticsAccess**](CloudUsageAnalyticsAccess.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1UsageLeaderboard

> CloudLeaderboardView CloudGetV1UsageLeaderboard(ctx).Scope(scope).Metric(metric).Period(period).Limit(limit).Execute()

Leaderboard ranks AI usage over a window, either the users of the caller's own org or organizations against each other, and always reports the caller's own standing even when it falls outside the returned page.



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
	scope := "personal" // string | Scope picks the board: \"personal\" (default) ranks the caller among their own org's users, \"org\" is that same org board named for an admin, \"global\" ranks organizations against each other. (optional)
	metric := "tokens" // string | Metric is the value ranked: tokens (default), requests, or cost. (optional)
	period := "week" // string | Period is the window ranked: day, week, month (default) or all. (optional)
	limit := int32(10) // int32 | Limit caps the rows returned, clamped to 100. Defaults to 10, which is also what a non-positive or unparseable value takes. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.CloudGetV1UsageLeaderboard(context.Background()).Scope(scope).Metric(metric).Period(period).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.CloudGetV1UsageLeaderboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1UsageLeaderboard`: CloudLeaderboardView
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.CloudGetV1UsageLeaderboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1UsageLeaderboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | Scope picks the board: \&quot;personal\&quot; (default) ranks the caller among their own org&#39;s users, \&quot;org\&quot; is that same org board named for an admin, \&quot;global\&quot; ranks organizations against each other. | 
 **metric** | **string** | Metric is the value ranked: tokens (default), requests, or cost. | 
 **period** | **string** | Period is the window ranked: day, week, month (default) or all. | 
 **limit** | **int32** | Limit caps the rows returned, clamped to 100. Defaults to 10, which is also what a non-positive or unparseable value takes. | 

### Return type

[**CloudLeaderboardView**](CloudLeaderboardView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1UsageLeaderboardOptin

> CloudOptinView CloudGetV1UsageLeaderboardOptin(ctx).Execute()

GetOptin returns the caller's own public-listing preference and their org's, each with whether the caller may change it.



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
	resp, r, err := apiClient.UsageAPI.CloudGetV1UsageLeaderboardOptin(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.CloudGetV1UsageLeaderboardOptin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1UsageLeaderboardOptin`: CloudOptinView
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.CloudGetV1UsageLeaderboardOptin`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1UsageLeaderboardOptinRequest struct via the builder pattern


### Return type

[**CloudOptinView**](CloudOptinView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1UsageSamples

> CloudDashResp CloudGetV1UsageSamples(ctx).Account(account).Provider(provider).Range_(range_).Window(window).Execute()

Is the PER-PROVIDER view: one connected account's own consumption of its own plan — \"my Claude Max plan is 47% through its 6h window, resets at 14:20\".



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
	account := "account_example" // string | Account narrows to ONE linked account of that provider. Empty covers every account the caller has linked there. (optional)
	provider := "provider_example" // string | Provider is the upstream to read, e.g. anthropic. Required. (optional)
	range_ := "range__example" // string | Range is the window to read: 1h, 24h, 7d or 30d. Empty means 24h, and any other label is refused rather than silently replaced. (optional)
	window := "window_example" // string | Window narrows to ONE window class: 6h, day, week or month. Empty covers every class. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.CloudGetV1UsageSamples(context.Background()).Account(account).Provider(provider).Range_(range_).Window(window).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.CloudGetV1UsageSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1UsageSamples`: CloudDashResp
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.CloudGetV1UsageSamples`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1UsageSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | Account narrows to ONE linked account of that provider. Empty covers every account the caller has linked there. | 
 **provider** | **string** | Provider is the upstream to read, e.g. anthropic. Required. | 
 **range_** | **string** | Range is the window to read: 1h, 24h, 7d or 30d. Empty means 24h, and any other label is refused rather than silently replaced. | 
 **window** | **string** | Window narrows to ONE window class: 6h, day, week or month. Empty covers every class. | 

### Return type

[**CloudDashResp**](CloudDashResp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1UsageSummary

> CloudUsageSummary CloudGetV1UsageSummary(ctx).Range_(range_).Start(start).End(end).Execute()

Answers GET /v1/usage/summary: the caller's own usage footprint over one window — the categorized spend roll-up from the commerce ledger, the org's LLM usage totals from the warehouse, and the caller's OWN linked provider accounts beside the org's Hanzo-routed usage.



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
	range_ := "range__example" // string | Range is the window: 24h, 7d, 30d, or custom. Empty means 24h. A label this surface does not know is refused rather than silently replaced. (optional)
	start := "start_example" // string | Start is the inclusive window start, RFC3339. Read only when Range is custom. (optional)
	end := "end_example" // string | End is the exclusive window end, RFC3339. Read only when Range is custom. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.CloudGetV1UsageSummary(context.Background()).Range_(range_).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.CloudGetV1UsageSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1UsageSummary`: CloudUsageSummary
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.CloudGetV1UsageSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1UsageSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is the window: 24h, 7d, 30d, or custom. Empty means 24h. A label this surface does not know is refused rather than silently replaced. | 
 **start** | **string** | Start is the inclusive window start, RFC3339. Read only when Range is custom. | 
 **end** | **string** | End is the exclusive window end, RFC3339. Read only when Range is custom. | 

### Return type

[**CloudUsageSummary**](CloudUsageSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Usage

> CloudReportResp CloudPostV1Usage(ctx).CloudReportReq(cloudReportReq).Execute()

Ingests a batch of account-usage samples — what a developer's OWN AI accounts have consumed of their OWN plans, metered from each provider's own login — and appends them to the warehouse series.



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
	cloudReportReq := *openapiclient.NewCloudReportReq() // CloudReportReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.CloudPostV1Usage(context.Background()).CloudReportReq(cloudReportReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.CloudPostV1Usage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Usage`: CloudReportResp
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.CloudPostV1Usage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1UsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudReportReq** | [**CloudReportReq**](CloudReportReq.md) |  | 

### Return type

[**CloudReportResp**](CloudReportResp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1UsageRollupBackfill

> CloudBackfillResult CloudPostV1UsageRollupBackfill(ctx).CloudBackfillQuery(cloudBackfillQuery).Execute()

Backfill seeds the derived usage rollup from ledger history — the rows written before the incremental view existed, which that view can never capture.



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
	cloudBackfillQuery := *openapiclient.NewCloudBackfillQuery() // CloudBackfillQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.CloudPostV1UsageRollupBackfill(context.Background()).CloudBackfillQuery(cloudBackfillQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.CloudPostV1UsageRollupBackfill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1UsageRollupBackfill`: CloudBackfillResult
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.CloudPostV1UsageRollupBackfill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1UsageRollupBackfillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudBackfillQuery** | [**CloudBackfillQuery**](CloudBackfillQuery.md) |  | 

### Return type

[**CloudBackfillResult**](CloudBackfillResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1UsageLeaderboardOptin

> CloudUserOptinView CloudPutV1UsageLeaderboardOptin(ctx).CloudUserOptinReq(cloudUserOptinReq).Execute()

PutUserOptin sets the CALLER's own public-listing preference on the leaderboard.



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
	cloudUserOptinReq := *openapiclient.NewCloudUserOptinReq() // CloudUserOptinReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.CloudPutV1UsageLeaderboardOptin(context.Background()).CloudUserOptinReq(cloudUserOptinReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.CloudPutV1UsageLeaderboardOptin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1UsageLeaderboardOptin`: CloudUserOptinView
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.CloudPutV1UsageLeaderboardOptin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1UsageLeaderboardOptinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudUserOptinReq** | [**CloudUserOptinReq**](CloudUserOptinReq.md) |  | 

### Return type

[**CloudUserOptinView**](CloudUserOptinView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1UsageLeaderboardOptinOrg

> CloudOrgOptinView CloudPutV1UsageLeaderboardOptinOrg(ctx).CloudOrgOptinReq(cloudOrgOptinReq).Execute()

PutOrgOptin sets the ORG's listing on the cross-org global board.



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
	cloudOrgOptinReq := *openapiclient.NewCloudOrgOptinReq() // CloudOrgOptinReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.CloudPutV1UsageLeaderboardOptinOrg(context.Background()).CloudOrgOptinReq(cloudOrgOptinReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.CloudPutV1UsageLeaderboardOptinOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1UsageLeaderboardOptinOrg`: CloudOrgOptinView
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.CloudPutV1UsageLeaderboardOptinOrg`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1UsageLeaderboardOptinOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudOrgOptinReq** | [**CloudOrgOptinReq**](CloudOrgOptinReq.md) |  | 

### Return type

[**CloudOrgOptinView**](CloudOrgOptinView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

