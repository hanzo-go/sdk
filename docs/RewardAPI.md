# \RewardAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRewardBadgeByHandle**](RewardAPI.md#GetRewardBadgeByHandle) | **Get** /v1/reward/badge/{handle} | The README badge of a person&#39;s level, as an SVG
[**GetRewardCatalog**](RewardAPI.md#GetRewardCatalog) | **Get** /v1/reward/catalog | The events, pillars, levels and rule operators the program is built from
[**GetRewardHealth**](RewardAPI.md#GetRewardHealth) | **Get** /v1/reward/health | Whether the rewards program can answer
[**GetRewardMe**](RewardAPI.md#GetRewardMe) | **Get** /v1/reward/me | Your level, lifetime points, redeemable points and badges
[**GetRewardMeClaims**](RewardAPI.md#GetRewardMeClaims) | **Get** /v1/reward/me/claims | Your claims, and where each stands
[**GetRewardMeExplore**](RewardAPI.md#GetRewardMeExplore) | **Get** /v1/reward/me/explore | What you could earn next, by pillar
[**GetRewardMePosts**](RewardAPI.md#GetRewardMePosts) | **Get** /v1/reward/me/posts | Lists your posts, newest first, each with its status and every look at it.
[**GetRewardMeProfile**](RewardAPI.md#GetRewardMeProfile) | **Get** /v1/reward/me/profile | Your rewards profile
[**GetRewardProfileByHandle**](RewardAPI.md#GetRewardProfileByHandle) | **Get** /v1/reward/profile/{handle} | A person&#39;s public rewards profile
[**PostRewardEvents**](RewardAPI.md#PostRewardEvents) | **Post** /v1/reward/events | State one event to the rules engine (the platform only)
[**PostRewardMeGithubBadge**](RewardAPI.md#PostRewardMeGithubBadge) | **Post** /v1/reward/me/github/badge | Verifies against GitHub&#39;s API that a project of yours carries the Hanzo badge in its README, and puts it to staff.
[**PostRewardMeGithubCheck**](RewardAPI.md#PostRewardMeGithubCheck) | **Post** /v1/reward/me/github/check | Reads the stars, follows and forks of the GitHub account you linked at sign-in, from GitHub&#39;s API, and puts to staff each that is a repository or account of ours, once each.
[**PostRewardMePosts**](RewardAPI.md#PostRewardMePosts) | **Post** /v1/reward/me/posts | Takes the address of a post of yours, reads it from its network, and holds it for a week if it is public, is yours, and carries your own referral link.
[**PostRewardMePostsByIdRecheck**](RewardAPI.md#PostRewardMePostsByIdRecheck) | **Post** /v1/reward/me/posts/{id}/recheck | Looks at one of your pending or removed posts now, with the same checks as the daily look, and adds the look to its history.
[**PostRewardMePostsScan**](RewardAPI.md#PostRewardMePostsScan) | **Post** /v1/reward/me/posts/scan | Reads your recent public posts on Bluesky, Mastodon, Reddit, Hacker News or DEV by the handle you give, and holds each that carries your own referral link, as pasting it would.
[**PutRewardMeProfile**](RewardAPI.md#PutRewardMeProfile) | **Put** /v1/reward/me/profile | Set your handle, whether your profile is public, and your links



## GetRewardBadgeByHandle

> GetRewardBadgeByHandle(ctx, handle).Execute()

The README badge of a person's level, as an SVG



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
	handle := "handle_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RewardAPI.GetRewardBadgeByHandle(context.Background(), handle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardAPI.GetRewardBadgeByHandle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRewardBadgeByHandleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRewardCatalog

> GetRewardCatalog(ctx).Execute()

The events, pillars, levels and rule operators the program is built from



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
	r, err := apiClient.RewardAPI.GetRewardCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardAPI.GetRewardCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRewardCatalogRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRewardHealth

> GetRewardHealth(ctx).Execute()

Whether the rewards program can answer



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
	r, err := apiClient.RewardAPI.GetRewardHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardAPI.GetRewardHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRewardHealthRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRewardMe

> GetRewardMe(ctx).Execute()

Your level, lifetime points, redeemable points and badges



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
	r, err := apiClient.RewardAPI.GetRewardMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardAPI.GetRewardMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRewardMeRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRewardMeClaims

> GetRewardMeClaims(ctx).Execute()

Your claims, and where each stands



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
	r, err := apiClient.RewardAPI.GetRewardMeClaims(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardAPI.GetRewardMeClaims``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRewardMeClaimsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRewardMeExplore

> GetRewardMeExplore(ctx).Execute()

What you could earn next, by pillar



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
	r, err := apiClient.RewardAPI.GetRewardMeExplore(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardAPI.GetRewardMeExplore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRewardMeExploreRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRewardMePosts

> RewardPostList GetRewardMePosts(ctx).Execute()

Lists your posts, newest first, each with its status and every look at it.



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
	resp, r, err := apiClient.RewardAPI.GetRewardMePosts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardAPI.GetRewardMePosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRewardMePosts`: RewardPostList
	fmt.Fprintf(os.Stdout, "Response from `RewardAPI.GetRewardMePosts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRewardMePostsRequest struct via the builder pattern


### Return type

[**RewardPostList**](RewardPostList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRewardMeProfile

> GetRewardMeProfile(ctx).Execute()

Your rewards profile



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
	r, err := apiClient.RewardAPI.GetRewardMeProfile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardAPI.GetRewardMeProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRewardMeProfileRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRewardProfileByHandle

> GetRewardProfileByHandle(ctx, handle).Execute()

A person's public rewards profile



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
	handle := "handle_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RewardAPI.GetRewardProfileByHandle(context.Background(), handle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardAPI.GetRewardProfileByHandle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRewardProfileByHandleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRewardEvents

> PostRewardEvents(ctx).Execute()

State one event to the rules engine (the platform only)



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
	r, err := apiClient.RewardAPI.PostRewardEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardAPI.PostRewardEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostRewardEventsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRewardMeGithubBadge

> RewardBadgeOut PostRewardMeGithubBadge(ctx).RewardBadgeIn(rewardBadgeIn).Execute()

Verifies against GitHub's API that a project of yours carries the Hanzo badge in its README, and puts it to staff.



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
	rewardBadgeIn := *openapiclient.NewRewardBadgeIn() // RewardBadgeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RewardAPI.PostRewardMeGithubBadge(context.Background()).RewardBadgeIn(rewardBadgeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardAPI.PostRewardMeGithubBadge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRewardMeGithubBadge`: RewardBadgeOut
	fmt.Fprintf(os.Stdout, "Response from `RewardAPI.PostRewardMeGithubBadge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRewardMeGithubBadgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rewardBadgeIn** | [**RewardBadgeIn**](RewardBadgeIn.md) |  | 

### Return type

[**RewardBadgeOut**](RewardBadgeOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRewardMeGithubCheck

> RewardCheckOut PostRewardMeGithubCheck(ctx).Execute()

Reads the stars, follows and forks of the GitHub account you linked at sign-in, from GitHub's API, and puts to staff each that is a repository or account of ours, once each.



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
	resp, r, err := apiClient.RewardAPI.PostRewardMeGithubCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardAPI.PostRewardMeGithubCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRewardMeGithubCheck`: RewardCheckOut
	fmt.Fprintf(os.Stdout, "Response from `RewardAPI.PostRewardMeGithubCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostRewardMeGithubCheckRequest struct via the builder pattern


### Return type

[**RewardCheckOut**](RewardCheckOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRewardMePosts

> RewardPostView PostRewardMePosts(ctx).RewardPostIn(rewardPostIn).Execute()

Takes the address of a post of yours, reads it from its network, and holds it for a week if it is public, is yours, and carries your own referral link.



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
	rewardPostIn := *openapiclient.NewRewardPostIn() // RewardPostIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RewardAPI.PostRewardMePosts(context.Background()).RewardPostIn(rewardPostIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardAPI.PostRewardMePosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRewardMePosts`: RewardPostView
	fmt.Fprintf(os.Stdout, "Response from `RewardAPI.PostRewardMePosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRewardMePostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rewardPostIn** | [**RewardPostIn**](RewardPostIn.md) |  | 

### Return type

[**RewardPostView**](RewardPostView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRewardMePostsByIdRecheck

> RewardPostView PostRewardMePostsByIdRecheck(ctx, id).Execute()

Looks at one of your pending or removed posts now, with the same checks as the daily look, and adds the look to its history.



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
	id := "post:9f8c1d2e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RewardAPI.PostRewardMePostsByIdRecheck(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardAPI.PostRewardMePostsByIdRecheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRewardMePostsByIdRecheck`: RewardPostView
	fmt.Fprintf(os.Stdout, "Response from `RewardAPI.PostRewardMePostsByIdRecheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRewardMePostsByIdRecheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RewardPostView**](RewardPostView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRewardMePostsScan

> RewardScanOut PostRewardMePostsScan(ctx).RewardScanIn(rewardScanIn).Execute()

Reads your recent public posts on Bluesky, Mastodon, Reddit, Hacker News or DEV by the handle you give, and holds each that carries your own referral link, as pasting it would.



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
	rewardScanIn := *openapiclient.NewRewardScanIn() // RewardScanIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RewardAPI.PostRewardMePostsScan(context.Background()).RewardScanIn(rewardScanIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardAPI.PostRewardMePostsScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRewardMePostsScan`: RewardScanOut
	fmt.Fprintf(os.Stdout, "Response from `RewardAPI.PostRewardMePostsScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRewardMePostsScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rewardScanIn** | [**RewardScanIn**](RewardScanIn.md) |  | 

### Return type

[**RewardScanOut**](RewardScanOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRewardMeProfile

> PutRewardMeProfile(ctx).Execute()

Set your handle, whether your profile is public, and your links



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
	r, err := apiClient.RewardAPI.PutRewardMeProfile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardAPI.PutRewardMeProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPutRewardMeProfileRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

