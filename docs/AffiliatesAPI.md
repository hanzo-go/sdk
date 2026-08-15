# \AffiliatesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAffiliates**](AffiliatesAPI.md#GetAffiliates) | **Get** /v1/affiliates | Answers the caller org&#39;s OWN affiliate standing: status, referral code and share link, commission rate, how many orgs it has referred, and its lifetime accrued, still-pending and already-paid commission in integer cents, with its payout history.
[**GetAffiliatesLeaderboard**](AffiliatesAPI.md#GetAffiliatesLeaderboard) | **Get** /v1/affiliates/leaderboard | Answers the top affiliates by lifetime accrued commission, shown by OPT-IN HANDLE with aggregate figures only, plus the caller&#39;s own exact rank.
[**GetAffiliatesMe**](AffiliatesAPI.md#GetAffiliatesMe) | **Get** /v1/affiliates/me | Answers the richer self-view: the same lifetime accrued, pending and paid commission and payout history, plus the caller&#39;s downline broken out by upline LEVEL — direct, second, third — each with the rate paid at that level and how many orgs sit there.
[**GetAffiliatesMeEarnings**](AffiliatesAPI.md#GetAffiliatesMeEarnings) | **Get** /v1/affiliates/me/earnings | Answers the caller&#39;s own commission ledger: per period, the margin it earned against and the commission taken from that margin; and per referred org, that referral&#39;s aggregate contribution.
[**GetAffiliatesMeLinks**](AffiliatesAPI.md#GetAffiliatesMeLinks) | **Get** /v1/affiliates/me/links | Answers the caller&#39;s share links, each with its URL and its funnel: clicks tracked, signups — orgs attributed with that code — and conversions, meaning how many of those signups have actually produced commission.
[**PostAffiliatesApply**](AffiliatesAPI.md#PostAffiliatesApply) | **Post** /v1/affiliates/apply | Enrolls the caller&#39;s OWN org as an affiliate at status &#x60;applied&#x60;, optionally requesting a vanity code, and answers the record — 201 on the first apply, 200 with &#x60;created:false&#x60; afterwards.
[**PostAffiliatesAttribute**](AffiliatesAPI.md#PostAffiliatesAttribute) | **Post** /v1/affiliates/attribute | Records the first-touch edge every later commission is computed from: the caller&#39;s org was referred by the affiliate that owns this code.
[**PostAffiliatesClick**](AffiliatesAPI.md#PostAffiliatesClick) | **Post** /v1/affiliates/click | Counts a click on a share link.
[**PostAffiliatesMeHandle**](AffiliatesAPI.md#PostAffiliatesMeHandle) | **Post** /v1/affiliates/me/handle | Sets the caller&#39;s public leaderboard display name, or clears it.
[**PostAffiliatesMeLinks**](AffiliatesAPI.md#PostAffiliatesMeLinks) | **Post** /v1/affiliates/me/links | Mints a new share link for the caller&#39;s own affiliate and answers it with its full URL, 201.



## GetAffiliates

> AffiliateStanding GetAffiliates(ctx).Execute()

Answers the caller org's OWN affiliate standing: status, referral code and share link, commission rate, how many orgs it has referred, and its lifetime accrued, still-pending and already-paid commission in integer cents, with its payout history.



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
	resp, r, err := apiClient.AffiliatesAPI.GetAffiliates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.GetAffiliates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAffiliates`: AffiliateStanding
	fmt.Fprintf(os.Stdout, "Response from `AffiliatesAPI.GetAffiliates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAffiliatesRequest struct via the builder pattern


### Return type

[**AffiliateStanding**](AffiliateStanding.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAffiliatesLeaderboard

> AffiliateBoard GetAffiliatesLeaderboard(ctx).Execute()

Answers the top affiliates by lifetime accrued commission, shown by OPT-IN HANDLE with aggregate figures only, plus the caller's own exact rank.



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
	resp, r, err := apiClient.AffiliatesAPI.GetAffiliatesLeaderboard(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.GetAffiliatesLeaderboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAffiliatesLeaderboard`: AffiliateBoard
	fmt.Fprintf(os.Stdout, "Response from `AffiliatesAPI.GetAffiliatesLeaderboard`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAffiliatesLeaderboardRequest struct via the builder pattern


### Return type

[**AffiliateBoard**](AffiliateBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAffiliatesMe

> AffiliateSelf GetAffiliatesMe(ctx).Execute()

Answers the richer self-view: the same lifetime accrued, pending and paid commission and payout history, plus the caller's downline broken out by upline LEVEL — direct, second, third — each with the rate paid at that level and how many orgs sit there.



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
	resp, r, err := apiClient.AffiliatesAPI.GetAffiliatesMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.GetAffiliatesMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAffiliatesMe`: AffiliateSelf
	fmt.Fprintf(os.Stdout, "Response from `AffiliatesAPI.GetAffiliatesMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAffiliatesMeRequest struct via the builder pattern


### Return type

[**AffiliateSelf**](AffiliateSelf.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAffiliatesMeEarnings

> AffiliateEarnings GetAffiliatesMeEarnings(ctx).Execute()

Answers the caller's own commission ledger: per period, the margin it earned against and the commission taken from that margin; and per referred org, that referral's aggregate contribution.



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
	resp, r, err := apiClient.AffiliatesAPI.GetAffiliatesMeEarnings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.GetAffiliatesMeEarnings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAffiliatesMeEarnings`: AffiliateEarnings
	fmt.Fprintf(os.Stdout, "Response from `AffiliatesAPI.GetAffiliatesMeEarnings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAffiliatesMeEarningsRequest struct via the builder pattern


### Return type

[**AffiliateEarnings**](AffiliateEarnings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAffiliatesMeLinks

> AffiliateLinks GetAffiliatesMeLinks(ctx).Execute()

Answers the caller's share links, each with its URL and its funnel: clicks tracked, signups — orgs attributed with that code — and conversions, meaning how many of those signups have actually produced commission.



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
	resp, r, err := apiClient.AffiliatesAPI.GetAffiliatesMeLinks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.GetAffiliatesMeLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAffiliatesMeLinks`: AffiliateLinks
	fmt.Fprintf(os.Stdout, "Response from `AffiliatesAPI.GetAffiliatesMeLinks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAffiliatesMeLinksRequest struct via the builder pattern


### Return type

[**AffiliateLinks**](AffiliateLinks.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAffiliatesApply

> Application PostAffiliatesApply(ctx).ApplyRequest(applyRequest).Execute()

Enrolls the caller's OWN org as an affiliate at status `applied`, optionally requesting a vanity code, and answers the record — 201 on the first apply, 200 with `created:false` afterwards.



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
	applyRequest := *openapiclient.NewApplyRequest() // ApplyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliatesAPI.PostAffiliatesApply(context.Background()).ApplyRequest(applyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.PostAffiliatesApply``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAffiliatesApply`: Application
	fmt.Fprintf(os.Stdout, "Response from `AffiliatesAPI.PostAffiliatesApply`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAffiliatesApplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applyRequest** | [**ApplyRequest**](ApplyRequest.md) |  | 

### Return type

[**Application**](Application.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAffiliatesAttribute

> Attribution PostAffiliatesAttribute(ctx).AttributeRequest(attributeRequest).Execute()

Records the first-touch edge every later commission is computed from: the caller's org was referred by the affiliate that owns this code.



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
	attributeRequest := *openapiclient.NewAttributeRequest() // AttributeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliatesAPI.PostAffiliatesAttribute(context.Background()).AttributeRequest(attributeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.PostAffiliatesAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAffiliatesAttribute`: Attribution
	fmt.Fprintf(os.Stdout, "Response from `AffiliatesAPI.PostAffiliatesAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAffiliatesAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributeRequest** | [**AttributeRequest**](AttributeRequest.md) |  | 

### Return type

[**Attribution**](Attribution.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAffiliatesClick

> ClickCount PostAffiliatesClick(ctx).ClickRequest(clickRequest).Execute()

Counts a click on a share link.



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
	clickRequest := *openapiclient.NewClickRequest() // ClickRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliatesAPI.PostAffiliatesClick(context.Background()).ClickRequest(clickRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.PostAffiliatesClick``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAffiliatesClick`: ClickCount
	fmt.Fprintf(os.Stdout, "Response from `AffiliatesAPI.PostAffiliatesClick`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAffiliatesClickRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clickRequest** | [**ClickRequest**](ClickRequest.md) |  | 

### Return type

[**ClickCount**](ClickCount.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAffiliatesMeHandle

> HandleSet PostAffiliatesMeHandle(ctx).HandleRequest(handleRequest).Execute()

Sets the caller's public leaderboard display name, or clears it.



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
	handleRequest := *openapiclient.NewHandleRequest() // HandleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliatesAPI.PostAffiliatesMeHandle(context.Background()).HandleRequest(handleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.PostAffiliatesMeHandle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAffiliatesMeHandle`: HandleSet
	fmt.Fprintf(os.Stdout, "Response from `AffiliatesAPI.PostAffiliatesMeHandle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAffiliatesMeHandleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **handleRequest** | [**HandleRequest**](HandleRequest.md) |  | 

### Return type

[**HandleSet**](HandleSet.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAffiliatesMeLinks

> LinkMint PostAffiliatesMeLinks(ctx).CreateLinkRequest(createLinkRequest).Execute()

Mints a new share link for the caller's own affiliate and answers it with its full URL, 201.



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
	createLinkRequest := *openapiclient.NewCreateLinkRequest() // CreateLinkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliatesAPI.PostAffiliatesMeLinks(context.Background()).CreateLinkRequest(createLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.PostAffiliatesMeLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAffiliatesMeLinks`: LinkMint
	fmt.Fprintf(os.Stdout, "Response from `AffiliatesAPI.PostAffiliatesMeLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAffiliatesMeLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLinkRequest** | [**CreateLinkRequest**](CreateLinkRequest.md) |  | 

### Return type

[**LinkMint**](LinkMint.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

