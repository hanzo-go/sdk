# \UsageAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsageActivity**](UsageAPI.md#GetUsageActivity) | **Get** /v1/usage/activity | Activity returns the per-day usage series for ONE authorized subject — the points a contribution heatmap and a timeline are drawn from, gap-filled so every day in the range is present.
[**GetUsageAnalytics**](UsageAPI.md#GetUsageAnalytics) | **Get** /v1/usage/analytics | Is the entitlement-GATED per-provider breakdown of the caller org&#39;s LLM usage — the paid lens over the same warehouse ledger GET /v1/usage/summary reads its totals from.
[**GetUsageAnalyticsAccess**](UsageAPI.md#GetUsageAnalyticsAccess) | **Get** /v1/usage/analytics/access | Echoes a plan&#39;s resolved analytics entitlement so a dashboard can configure itself against the LIVE catalog instead of hardcoding tier numbers.
[**GetUsageLeaderboard**](UsageAPI.md#GetUsageLeaderboard) | **Get** /v1/usage/leaderboard | Leaderboard ranks AI usage over a window, either the users of the caller&#39;s own org or organizations against each other, and always reports the caller&#39;s own standing even when it falls outside the returned page.
[**GetUsageLeaderboardOptin**](UsageAPI.md#GetUsageLeaderboardOptin) | **Get** /v1/usage/leaderboard/optin | Returns the caller&#39;s own public-listing preference and their org&#39;s, each with whether the caller may change it.
[**GetUsageSamples**](UsageAPI.md#GetUsageSamples) | **Get** /v1/usage/samples | Is the PER-PROVIDER view: one connected account&#39;s own consumption of its own plan — \&quot;my plan is 47% through its 6h window, resets at 14:20\&quot;.
[**GetUsageSummary**](UsageAPI.md#GetUsageSummary) | **Get** /v1/usage/summary | Answers GET /v1/usage/summary: the caller&#39;s own usage footprint over one window — the categorized spend roll-up from the commerce ledger, the org&#39;s LLM usage totals from the warehouse, and the caller&#39;s OWN linked provider accounts beside the org&#39;s Hanzo-routed usage.
[**PostUsage**](UsageAPI.md#PostUsage) | **Post** /v1/usage | Ingests a batch of account-usage samples — what a developer&#39;s OWN AI accounts have consumed of their OWN plans, metered from each provider&#39;s own login — and appends them to the warehouse series.
[**PostUsageRollupBackfill**](UsageAPI.md#PostUsageRollupBackfill) | **Post** /v1/usage/rollup/backfill | Backfill seeds the derived usage rollup from ledger history — the rows written before the incremental view existed, which that view can never capture.
[**PutUsageLeaderboardOptin**](UsageAPI.md#PutUsageLeaderboardOptin) | **Put** /v1/usage/leaderboard/optin | Sets the CALLER&#39;s own public-listing preference on the leaderboard.
[**PutUsageLeaderboardOptinOrg**](UsageAPI.md#PutUsageLeaderboardOptinOrg) | **Put** /v1/usage/leaderboard/optin/org | Sets the ORG&#39;s listing on the cross-org global board.



## GetUsageActivity

> ActivityView GetUsageActivity(ctx).Subject(subject).Id(id).From(from).To(to).Execute()

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
	resp, r, err := apiClient.UsageAPI.GetUsageActivity(context.Background()).Subject(subject).Id(id).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetUsageActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageActivity`: ActivityView
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.GetUsageActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subject** | **string** | Subject is what the series is about: \&quot;user\&quot; (default), \&quot;org\&quot; or \&quot;project\&quot;. | 
 **id** | **string** | ID names the subject within what the caller is entitled to see. Omitted (or \&quot;me\&quot;) it is the caller themselves, or their own org. Another user requires org admin and must belong to the caller&#39;s org; another org requires a SuperAdmin. | 
 **from** | **string** | From is the first day of the range, \&quot;2006-01-02\&quot;. Defaults to 90 days back. | 
 **to** | **string** | To is the last day of the range, \&quot;2006-01-02\&quot;. Defaults to today; the span is clamped to 366 days. | 

### Return type

[**ActivityView**](ActivityView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/hanzoai/go-sdk"
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

No authorization required

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
	openapiclient "github.com/hanzoai/go-sdk"
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageLeaderboard

> LeaderboardView GetUsageLeaderboard(ctx).Scope(scope).Metric(metric).Period(period).Limit(limit).Execute()

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
	resp, r, err := apiClient.UsageAPI.GetUsageLeaderboard(context.Background()).Scope(scope).Metric(metric).Period(period).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetUsageLeaderboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageLeaderboard`: LeaderboardView
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.GetUsageLeaderboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageLeaderboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | Scope picks the board: \&quot;personal\&quot; (default) ranks the caller among their own org&#39;s users, \&quot;org\&quot; is that same org board named for an admin, \&quot;global\&quot; ranks organizations against each other. | 
 **metric** | **string** | Metric is the value ranked: tokens (default), requests, or cost. | 
 **period** | **string** | Period is the window ranked: day, week, month (default) or all. | 
 **limit** | **int32** | Limit caps the rows returned, clamped to 100. Defaults to 10, which is also what a non-positive or unparseable value takes. | 

### Return type

[**LeaderboardView**](LeaderboardView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsageLeaderboardOptin

> OptinView GetUsageLeaderboardOptin(ctx).Execute()

Returns the caller's own public-listing preference and their org's, each with whether the caller may change it.



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
	resp, r, err := apiClient.UsageAPI.GetUsageLeaderboardOptin(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.GetUsageLeaderboardOptin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsageLeaderboardOptin`: OptinView
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.GetUsageLeaderboardOptin`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsageLeaderboardOptinRequest struct via the builder pattern


### Return type

[**OptinView**](OptinView.md)

### Authorization

No authorization required

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
	openapiclient "github.com/hanzoai/go-sdk"
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

No authorization required

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
	openapiclient "github.com/hanzoai/go-sdk"
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

No authorization required

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
	openapiclient "github.com/hanzoai/go-sdk"
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsageRollupBackfill

> BackfillResult PostUsageRollupBackfill(ctx).BackfillQuery(backfillQuery).Execute()

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
	backfillQuery := *openapiclient.NewBackfillQuery() // BackfillQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.PostUsageRollupBackfill(context.Background()).BackfillQuery(backfillQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.PostUsageRollupBackfill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsageRollupBackfill`: BackfillResult
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.PostUsageRollupBackfill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsageRollupBackfillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backfillQuery** | [**BackfillQuery**](BackfillQuery.md) |  | 

### Return type

[**BackfillResult**](BackfillResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUsageLeaderboardOptin

> UserOptinView PutUsageLeaderboardOptin(ctx).UserOptinReq(userOptinReq).Execute()

Sets the CALLER's own public-listing preference on the leaderboard.



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
	userOptinReq := *openapiclient.NewUserOptinReq() // UserOptinReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.PutUsageLeaderboardOptin(context.Background()).UserOptinReq(userOptinReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.PutUsageLeaderboardOptin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutUsageLeaderboardOptin`: UserOptinView
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.PutUsageLeaderboardOptin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutUsageLeaderboardOptinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userOptinReq** | [**UserOptinReq**](UserOptinReq.md) |  | 

### Return type

[**UserOptinView**](UserOptinView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUsageLeaderboardOptinOrg

> OrgOptinView PutUsageLeaderboardOptinOrg(ctx).OrgOptinReq(orgOptinReq).Execute()

Sets the ORG's listing on the cross-org global board.



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
	orgOptinReq := *openapiclient.NewOrgOptinReq() // OrgOptinReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.PutUsageLeaderboardOptinOrg(context.Background()).OrgOptinReq(orgOptinReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.PutUsageLeaderboardOptinOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutUsageLeaderboardOptinOrg`: OrgOptinView
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.PutUsageLeaderboardOptinOrg`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutUsageLeaderboardOptinOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgOptinReq** | [**OrgOptinReq**](OrgOptinReq.md) |  | 

### Return type

[**OrgOptinView**](OrgOptinView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

