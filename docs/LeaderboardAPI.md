# \LeaderboardAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLeaderboard**](LeaderboardAPI.md#GetLeaderboard) | **Get** /v1/leaderboard | Ranks AI usage over a window, either the users of the caller&#39;s own org or organizations against each other, and always reports the caller&#39;s own standing even when it falls outside the returned page.
[**GetLeaderboardActivity**](LeaderboardAPI.md#GetLeaderboardActivity) | **Get** /v1/leaderboard/activity | Returns the per-day usage series for ONE authorized subject — the points a contribution heatmap and a timeline are drawn from, gap-filled so every day in the range is present.
[**GetLeaderboardBeta**](LeaderboardAPI.md#GetLeaderboardBeta) | **Get** /v1/leaderboard/beta | Ranks everyone on the hanzo.ai/beta waitlist by one score: 1 per star on GitHub verified for the record and 1 per founder link followed, plus 1 per org that signed up with its referral code, plus 20 per such org that has made metered spend, plus 1 per 100,000 tokens the person used, all time.
[**GetLeaderboardBounty**](LeaderboardAPI.md#GetLeaderboardBounty) | **Get** /v1/leaderboard/bounty | Returns the gamified bounty &amp; rewards rules, ranked hunters, and recent bug/bounty awards.
[**GetLeaderboardFeed**](LeaderboardAPI.md#GetLeaderboardFeed) | **Get** /v1/leaderboard/feed | Stream the board&#39;s live lines to anyone as Server-Sent Events
[**GetLeaderboardOptin**](LeaderboardAPI.md#GetLeaderboardOptin) | **Get** /v1/leaderboard/optin | Returns the caller&#39;s own public-listing preference and their org&#39;s, each with whether the caller may change it.
[**GetLeaderboardPulse**](LeaderboardAPI.md#GetLeaderboardPulse) | **Get** /v1/leaderboard/pulse | Returns how fast the platform is being used and is spreading, as counts by day for thirty days and by hour for the last day: tokens used, orgs active, orgs signed up, orgs that came through a referral link, and orgs using a model for the first time, with the models most tried.
[**GetLeaderboardStanding**](LeaderboardAPI.md#GetLeaderboardStanding) | **Get** /v1/leaderboard/standing | Returns the public board of who has used and built the most on Hanzo, shown as a window around the reader.
[**PutLeaderboardOptin**](LeaderboardAPI.md#PutLeaderboardOptin) | **Put** /v1/leaderboard/optin | Sets the CALLER&#39;s own public-listing preference on the leaderboard.
[**PutLeaderboardOptinOrg**](LeaderboardAPI.md#PutLeaderboardOptinOrg) | **Put** /v1/leaderboard/optin/org | Sets the ORG&#39;s listing on the cross-org global board.



## GetLeaderboard

> LeaderboardLeaderboardView GetLeaderboard(ctx).Scope(scope).Metric(metric).Period(period).Limit(limit).Execute()

Ranks AI usage over a window, either the users of the caller's own org or organizations against each other, and always reports the caller's own standing even when it falls outside the returned page.



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
	scope := "personal" // string | Scope picks the board: \"personal\" (default) ranks the caller among their own org's users, \"org\" is that same org board named for an admin, \"global\" ranks organizations against each other. (optional)
	metric := "tokens" // string | Metric is the value ranked: tokens (default), requests, or cost. (optional)
	period := "week" // string | Period is the window ranked: day, week, month (default) or all. (optional)
	limit := int64(10) // int64 | Limit caps the rows returned, clamped to 100. Defaults to 10, which is also what a non-positive or unparseable value takes. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardAPI.GetLeaderboard(context.Background()).Scope(scope).Metric(metric).Period(period).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardAPI.GetLeaderboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaderboard`: LeaderboardLeaderboardView
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardAPI.GetLeaderboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaderboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | Scope picks the board: \&quot;personal\&quot; (default) ranks the caller among their own org&#39;s users, \&quot;org\&quot; is that same org board named for an admin, \&quot;global\&quot; ranks organizations against each other. | 
 **metric** | **string** | Metric is the value ranked: tokens (default), requests, or cost. | 
 **period** | **string** | Period is the window ranked: day, week, month (default) or all. | 
 **limit** | **int64** | Limit caps the rows returned, clamped to 100. Defaults to 10, which is also what a non-positive or unparseable value takes. | 

### Return type

[**LeaderboardLeaderboardView**](LeaderboardLeaderboardView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaderboardActivity

> LeaderboardActivityView GetLeaderboardActivity(ctx).Subject(subject).Id(id).From(from).To(to).Execute()

Returns the per-day usage series for ONE authorized subject — the points a contribution heatmap and a timeline are drawn from, gap-filled so every day in the range is present.



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
	subject := "user" // string | Subject is what the series is about: \"user\" (default), \"org\" or \"project\". (optional)
	id := "id_example" // string | ID names the subject within what the caller is entitled to see. Omitted (or \"me\") it is the caller themselves, or their own org. Another user requires org admin and must belong to the caller's org; another org requires a SuperAdmin. (optional)
	from := "2026-01-01" // string | From is the first day of the range, \"2006-01-02\". Defaults to 90 days back. (optional)
	to := "2026-03-31" // string | To is the last day of the range, \"2006-01-02\". Defaults to today; the span is clamped to 366 days. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardAPI.GetLeaderboardActivity(context.Background()).Subject(subject).Id(id).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardAPI.GetLeaderboardActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaderboardActivity`: LeaderboardActivityView
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardAPI.GetLeaderboardActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaderboardActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subject** | **string** | Subject is what the series is about: \&quot;user\&quot; (default), \&quot;org\&quot; or \&quot;project\&quot;. | 
 **id** | **string** | ID names the subject within what the caller is entitled to see. Omitted (or \&quot;me\&quot;) it is the caller themselves, or their own org. Another user requires org admin and must belong to the caller&#39;s org; another org requires a SuperAdmin. | 
 **from** | **string** | From is the first day of the range, \&quot;2006-01-02\&quot;. Defaults to 90 days back. | 
 **to** | **string** | To is the last day of the range, \&quot;2006-01-02\&quot;. Defaults to today; the span is clamped to 366 days. | 

### Return type

[**LeaderboardActivityView**](LeaderboardActivityView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaderboardBeta

> LeaderboardBetaBoard GetLeaderboardBeta(ctx).Limit(limit).Execute()

Ranks everyone on the hanzo.ai/beta waitlist by one score: 1 per star on GitHub verified for the record and 1 per founder link followed, plus 1 per org that signed up with its referral code, plus 20 per such org that has made metered spend, plus 1 per 100,000 tokens the person used, all time.



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
	limit := int64(789) // int64 | Limit caps the rows, clamped to 1..100. Defaults to 25, which is also what a non-positive value takes. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardAPI.GetLeaderboardBeta(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardAPI.GetLeaderboardBeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaderboardBeta`: LeaderboardBetaBoard
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardAPI.GetLeaderboardBeta`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaderboardBetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit caps the rows, clamped to 1..100. Defaults to 25, which is also what a non-positive value takes. | 

### Return type

[**LeaderboardBetaBoard**](LeaderboardBetaBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaderboardBounty

> LeaderboardBountyBoard GetLeaderboardBounty(ctx).Limit(limit).Execute()

Returns the gamified bounty & rewards rules, ranked hunters, and recent bug/bounty awards.



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
	limit := int64(789) // int64 | Limit caps the hunter and award rows returned, clamped to 1..100. Defaults to 25. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardAPI.GetLeaderboardBounty(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardAPI.GetLeaderboardBounty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaderboardBounty`: LeaderboardBountyBoard
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardAPI.GetLeaderboardBounty`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaderboardBountyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit caps the hunter and award rows returned, clamped to 1..100. Defaults to 25. | 

### Return type

[**LeaderboardBountyBoard**](LeaderboardBountyBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaderboardFeed

> *os.File GetLeaderboardFeed(ctx).Execute()

Stream the board's live lines to anyone as Server-Sent Events



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardAPI.GetLeaderboardFeed(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardAPI.GetLeaderboardFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaderboardFeed`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardAPI.GetLeaderboardFeed`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaderboardFeedRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaderboardOptin

> LeaderboardOptinView GetLeaderboardOptin(ctx).Execute()

Returns the caller's own public-listing preference and their org's, each with whether the caller may change it.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardAPI.GetLeaderboardOptin(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardAPI.GetLeaderboardOptin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaderboardOptin`: LeaderboardOptinView
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardAPI.GetLeaderboardOptin`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaderboardOptinRequest struct via the builder pattern


### Return type

[**LeaderboardOptinView**](LeaderboardOptinView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaderboardPulse

> LeaderboardPulseView GetLeaderboardPulse(ctx).Execute()

Returns how fast the platform is being used and is spreading, as counts by day for thirty days and by hour for the last day: tokens used, orgs active, orgs signed up, orgs that came through a referral link, and orgs using a model for the first time, with the models most tried.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardAPI.GetLeaderboardPulse(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardAPI.GetLeaderboardPulse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaderboardPulse`: LeaderboardPulseView
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardAPI.GetLeaderboardPulse`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaderboardPulseRequest struct via the builder pattern


### Return type

[**LeaderboardPulseView**](LeaderboardPulseView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaderboardStanding

> LeaderboardStandingView GetLeaderboardStanding(ctx).Execute()

Returns the public board of who has used and built the most on Hanzo, shown as a window around the reader.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardAPI.GetLeaderboardStanding(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardAPI.GetLeaderboardStanding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaderboardStanding`: LeaderboardStandingView
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardAPI.GetLeaderboardStanding`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaderboardStandingRequest struct via the builder pattern


### Return type

[**LeaderboardStandingView**](LeaderboardStandingView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLeaderboardOptin

> LeaderboardUserOptinView PutLeaderboardOptin(ctx).LeaderboardUserOptinReq(leaderboardUserOptinReq).Execute()

Sets the CALLER's own public-listing preference on the leaderboard.



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
	leaderboardUserOptinReq := *openapiclient.NewLeaderboardUserOptinReq() // LeaderboardUserOptinReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardAPI.PutLeaderboardOptin(context.Background()).LeaderboardUserOptinReq(leaderboardUserOptinReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardAPI.PutLeaderboardOptin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLeaderboardOptin`: LeaderboardUserOptinView
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardAPI.PutLeaderboardOptin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutLeaderboardOptinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **leaderboardUserOptinReq** | [**LeaderboardUserOptinReq**](LeaderboardUserOptinReq.md) |  | 

### Return type

[**LeaderboardUserOptinView**](LeaderboardUserOptinView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLeaderboardOptinOrg

> LeaderboardOrgOptinView PutLeaderboardOptinOrg(ctx).LeaderboardOrgOptinReq(leaderboardOrgOptinReq).Execute()

Sets the ORG's listing on the cross-org global board.



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
	leaderboardOrgOptinReq := *openapiclient.NewLeaderboardOrgOptinReq() // LeaderboardOrgOptinReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaderboardAPI.PutLeaderboardOptinOrg(context.Background()).LeaderboardOrgOptinReq(leaderboardOrgOptinReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardAPI.PutLeaderboardOptinOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLeaderboardOptinOrg`: LeaderboardOrgOptinView
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardAPI.PutLeaderboardOptinOrg`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutLeaderboardOptinOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **leaderboardOrgOptinReq** | [**LeaderboardOrgOptinReq**](LeaderboardOrgOptinReq.md) |  | 

### Return type

[**LeaderboardOrgOptinView**](LeaderboardOrgOptinView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

