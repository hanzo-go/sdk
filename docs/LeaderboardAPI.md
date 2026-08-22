# \LeaderboardAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLeaderboard**](LeaderboardAPI.md#GetLeaderboard) | **Get** /v1/leaderboard | Leaderboard ranks AI usage over a window, either the users of the caller&#39;s own org or organizations against each other, and always reports the caller&#39;s own standing even when it falls outside the returned page.
[**GetLeaderboardActivity**](LeaderboardAPI.md#GetLeaderboardActivity) | **Get** /v1/leaderboard/activity | Activity returns the per-day usage series for ONE authorized subject — the points a contribution heatmap and a timeline are drawn from, gap-filled so every day in the range is present.
[**GetLeaderboardOptin**](LeaderboardAPI.md#GetLeaderboardOptin) | **Get** /v1/leaderboard/optin | Returns the caller&#39;s own public-listing preference and their org&#39;s, each with whether the caller may change it.
[**PutLeaderboardOptin**](LeaderboardAPI.md#PutLeaderboardOptin) | **Put** /v1/leaderboard/optin | Sets the CALLER&#39;s own public-listing preference on the leaderboard.
[**PutLeaderboardOptinOrg**](LeaderboardAPI.md#PutLeaderboardOptinOrg) | **Put** /v1/leaderboard/optin/org | Sets the ORG&#39;s listing on the cross-org global board.



## GetLeaderboard

> LeaderboardView GetLeaderboard(ctx).Scope(scope).Metric(metric).Period(period).Limit(limit).Execute()

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
	resp, r, err := apiClient.LeaderboardAPI.GetLeaderboard(context.Background()).Scope(scope).Metric(metric).Period(period).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardAPI.GetLeaderboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaderboard`: LeaderboardView
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
 **limit** | **int32** | Limit caps the rows returned, clamped to 100. Defaults to 10, which is also what a non-positive or unparseable value takes. | 

### Return type

[**LeaderboardView**](LeaderboardView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaderboardActivity

> ActivityView GetLeaderboardActivity(ctx).Subject(subject).Id(id).From(from).To(to).Execute()

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
	resp, r, err := apiClient.LeaderboardAPI.GetLeaderboardActivity(context.Background()).Subject(subject).Id(id).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardAPI.GetLeaderboardActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaderboardActivity`: ActivityView
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

[**ActivityView**](ActivityView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaderboardOptin

> OptinView GetLeaderboardOptin(ctx).Execute()

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
	resp, r, err := apiClient.LeaderboardAPI.GetLeaderboardOptin(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardAPI.GetLeaderboardOptin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaderboardOptin`: OptinView
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardAPI.GetLeaderboardOptin`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaderboardOptinRequest struct via the builder pattern


### Return type

[**OptinView**](OptinView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLeaderboardOptin

> UserOptinView PutLeaderboardOptin(ctx).UserOptinReq(userOptinReq).Execute()

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
	resp, r, err := apiClient.LeaderboardAPI.PutLeaderboardOptin(context.Background()).UserOptinReq(userOptinReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardAPI.PutLeaderboardOptin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLeaderboardOptin`: UserOptinView
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardAPI.PutLeaderboardOptin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutLeaderboardOptinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userOptinReq** | [**UserOptinReq**](UserOptinReq.md) |  | 

### Return type

[**UserOptinView**](UserOptinView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLeaderboardOptinOrg

> OrgOptinView PutLeaderboardOptinOrg(ctx).OrgOptinReq(orgOptinReq).Execute()

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
	resp, r, err := apiClient.LeaderboardAPI.PutLeaderboardOptinOrg(context.Background()).OrgOptinReq(orgOptinReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaderboardAPI.PutLeaderboardOptinOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLeaderboardOptinOrg`: OrgOptinView
	fmt.Fprintf(os.Stdout, "Response from `LeaderboardAPI.PutLeaderboardOptinOrg`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutLeaderboardOptinOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgOptinReq** | [**OrgOptinReq**](OrgOptinReq.md) |  | 

### Return type

[**OrgOptinView**](OrgOptinView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

