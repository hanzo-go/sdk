# \AffiliateAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAffiliate**](AffiliateAPI.md#GetAffiliate) | **Get** /v1/affiliate | Answers the caller org&#39;s OWN affiliate standing: status, referral code and share link, commission rate, how many orgs it has referred, and its lifetime accrued, still-pending and already-paid commission in integer cents, with its payout history.
[**GetAffiliateLeaderboard**](AffiliateAPI.md#GetAffiliateLeaderboard) | **Get** /v1/affiliate/leaderboard | Answers the top affiliates by lifetime accrued commission, shown by OPT-IN HANDLE with aggregate figures only, plus the caller&#39;s own exact rank.
[**GetAffiliateMe**](AffiliateAPI.md#GetAffiliateMe) | **Get** /v1/affiliate/me | Answers the richer self-view: the same lifetime accrued, pending and paid commission and payout history, plus the caller&#39;s downline broken out by upline LEVEL — direct, second, third — each with the rate paid at that level and how many orgs sit there.
[**GetAffiliateMeEarnings**](AffiliateAPI.md#GetAffiliateMeEarnings) | **Get** /v1/affiliate/me/earnings | Answers the caller&#39;s own commission ledger: per period, the margin it earned against and the commission taken from that margin; and per referred org, that referral&#39;s aggregate contribution.
[**GetAffiliateMeLinks**](AffiliateAPI.md#GetAffiliateMeLinks) | **Get** /v1/affiliate/me/links | Answers the caller&#39;s share links, each with its URL and its funnel: clicks tracked, signups — orgs attributed with that code — and conversions, meaning how many of those signups have actually produced commission.
[**PostAffiliateApply**](AffiliateAPI.md#PostAffiliateApply) | **Post** /v1/affiliate/apply | Enrolls the caller&#39;s OWN org as an affiliate at status &#x60;applied&#x60;, optionally requesting a vanity code, and answers the record — 201 on the first apply, 200 with &#x60;created:false&#x60; afterwards.
[**PostAffiliateAttribute**](AffiliateAPI.md#PostAffiliateAttribute) | **Post** /v1/affiliate/attribute | Records the first-touch edge every later commission is computed from: the caller&#39;s org was referred by the affiliate that owns this code.
[**PostAffiliateClick**](AffiliateAPI.md#PostAffiliateClick) | **Post** /v1/affiliate/click | Counts a click on a share link.
[**PostAffiliateMeHandle**](AffiliateAPI.md#PostAffiliateMeHandle) | **Post** /v1/affiliate/me/handle | Sets the caller&#39;s public leaderboard display name, or clears it.
[**PostAffiliateMeLinks**](AffiliateAPI.md#PostAffiliateMeLinks) | **Post** /v1/affiliate/me/links | Mints a new share link for the caller&#39;s own affiliate and answers it with its full URL, 201.



## GetAffiliate

> AffiliateAffiliateStanding GetAffiliate(ctx).Execute()

Answers the caller org's OWN affiliate standing: status, referral code and share link, commission rate, how many orgs it has referred, and its lifetime accrued, still-pending and already-paid commission in integer cents, with its payout history.



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
	resp, r, err := apiClient.AffiliateAPI.GetAffiliate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliateAPI.GetAffiliate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAffiliate`: AffiliateAffiliateStanding
	fmt.Fprintf(os.Stdout, "Response from `AffiliateAPI.GetAffiliate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAffiliateRequest struct via the builder pattern


### Return type

[**AffiliateAffiliateStanding**](AffiliateAffiliateStanding.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAffiliateLeaderboard

> AffiliateAffiliateBoard GetAffiliateLeaderboard(ctx).Execute()

Answers the top affiliates by lifetime accrued commission, shown by OPT-IN HANDLE with aggregate figures only, plus the caller's own exact rank.



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
	resp, r, err := apiClient.AffiliateAPI.GetAffiliateLeaderboard(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliateAPI.GetAffiliateLeaderboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAffiliateLeaderboard`: AffiliateAffiliateBoard
	fmt.Fprintf(os.Stdout, "Response from `AffiliateAPI.GetAffiliateLeaderboard`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAffiliateLeaderboardRequest struct via the builder pattern


### Return type

[**AffiliateAffiliateBoard**](AffiliateAffiliateBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAffiliateMe

> AffiliateAffiliateSelf GetAffiliateMe(ctx).Execute()

Answers the richer self-view: the same lifetime accrued, pending and paid commission and payout history, plus the caller's downline broken out by upline LEVEL — direct, second, third — each with the rate paid at that level and how many orgs sit there.



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
	resp, r, err := apiClient.AffiliateAPI.GetAffiliateMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliateAPI.GetAffiliateMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAffiliateMe`: AffiliateAffiliateSelf
	fmt.Fprintf(os.Stdout, "Response from `AffiliateAPI.GetAffiliateMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAffiliateMeRequest struct via the builder pattern


### Return type

[**AffiliateAffiliateSelf**](AffiliateAffiliateSelf.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAffiliateMeEarnings

> AffiliateAffiliateEarnings GetAffiliateMeEarnings(ctx).Execute()

Answers the caller's own commission ledger: per period, the margin it earned against and the commission taken from that margin; and per referred org, that referral's aggregate contribution.



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
	resp, r, err := apiClient.AffiliateAPI.GetAffiliateMeEarnings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliateAPI.GetAffiliateMeEarnings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAffiliateMeEarnings`: AffiliateAffiliateEarnings
	fmt.Fprintf(os.Stdout, "Response from `AffiliateAPI.GetAffiliateMeEarnings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAffiliateMeEarningsRequest struct via the builder pattern


### Return type

[**AffiliateAffiliateEarnings**](AffiliateAffiliateEarnings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAffiliateMeLinks

> AffiliateAffiliateLinks GetAffiliateMeLinks(ctx).Execute()

Answers the caller's share links, each with its URL and its funnel: clicks tracked, signups — orgs attributed with that code — and conversions, meaning how many of those signups have actually produced commission.



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
	resp, r, err := apiClient.AffiliateAPI.GetAffiliateMeLinks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliateAPI.GetAffiliateMeLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAffiliateMeLinks`: AffiliateAffiliateLinks
	fmt.Fprintf(os.Stdout, "Response from `AffiliateAPI.GetAffiliateMeLinks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAffiliateMeLinksRequest struct via the builder pattern


### Return type

[**AffiliateAffiliateLinks**](AffiliateAffiliateLinks.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAffiliateApply

> AffiliateApplication PostAffiliateApply(ctx).AffiliateApplyRequest(affiliateApplyRequest).Execute()

Enrolls the caller's OWN org as an affiliate at status `applied`, optionally requesting a vanity code, and answers the record — 201 on the first apply, 200 with `created:false` afterwards.



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
	affiliateApplyRequest := *openapiclient.NewAffiliateApplyRequest() // AffiliateApplyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliateAPI.PostAffiliateApply(context.Background()).AffiliateApplyRequest(affiliateApplyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliateAPI.PostAffiliateApply``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAffiliateApply`: AffiliateApplication
	fmt.Fprintf(os.Stdout, "Response from `AffiliateAPI.PostAffiliateApply`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAffiliateApplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **affiliateApplyRequest** | [**AffiliateApplyRequest**](AffiliateApplyRequest.md) |  | 

### Return type

[**AffiliateApplication**](AffiliateApplication.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAffiliateAttribute

> AffiliateAttribution PostAffiliateAttribute(ctx).AffiliateAttributeRequest(affiliateAttributeRequest).Execute()

Records the first-touch edge every later commission is computed from: the caller's org was referred by the affiliate that owns this code.



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
	affiliateAttributeRequest := *openapiclient.NewAffiliateAttributeRequest() // AffiliateAttributeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliateAPI.PostAffiliateAttribute(context.Background()).AffiliateAttributeRequest(affiliateAttributeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliateAPI.PostAffiliateAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAffiliateAttribute`: AffiliateAttribution
	fmt.Fprintf(os.Stdout, "Response from `AffiliateAPI.PostAffiliateAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAffiliateAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **affiliateAttributeRequest** | [**AffiliateAttributeRequest**](AffiliateAttributeRequest.md) |  | 

### Return type

[**AffiliateAttribution**](AffiliateAttribution.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAffiliateClick

> AffiliateClickCount PostAffiliateClick(ctx).AffiliateClickRequest(affiliateClickRequest).Execute()

Counts a click on a share link.



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
	affiliateClickRequest := *openapiclient.NewAffiliateClickRequest() // AffiliateClickRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliateAPI.PostAffiliateClick(context.Background()).AffiliateClickRequest(affiliateClickRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliateAPI.PostAffiliateClick``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAffiliateClick`: AffiliateClickCount
	fmt.Fprintf(os.Stdout, "Response from `AffiliateAPI.PostAffiliateClick`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAffiliateClickRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **affiliateClickRequest** | [**AffiliateClickRequest**](AffiliateClickRequest.md) |  | 

### Return type

[**AffiliateClickCount**](AffiliateClickCount.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAffiliateMeHandle

> AffiliateHandleSet PostAffiliateMeHandle(ctx).AffiliateHandleRequest(affiliateHandleRequest).Execute()

Sets the caller's public leaderboard display name, or clears it.



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
	affiliateHandleRequest := *openapiclient.NewAffiliateHandleRequest() // AffiliateHandleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliateAPI.PostAffiliateMeHandle(context.Background()).AffiliateHandleRequest(affiliateHandleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliateAPI.PostAffiliateMeHandle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAffiliateMeHandle`: AffiliateHandleSet
	fmt.Fprintf(os.Stdout, "Response from `AffiliateAPI.PostAffiliateMeHandle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAffiliateMeHandleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **affiliateHandleRequest** | [**AffiliateHandleRequest**](AffiliateHandleRequest.md) |  | 

### Return type

[**AffiliateHandleSet**](AffiliateHandleSet.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAffiliateMeLinks

> AffiliateLinkMint PostAffiliateMeLinks(ctx).AffiliateCreateLinkRequest(affiliateCreateLinkRequest).Execute()

Mints a new share link for the caller's own affiliate and answers it with its full URL, 201.



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
	affiliateCreateLinkRequest := *openapiclient.NewAffiliateCreateLinkRequest() // AffiliateCreateLinkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliateAPI.PostAffiliateMeLinks(context.Background()).AffiliateCreateLinkRequest(affiliateCreateLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliateAPI.PostAffiliateMeLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAffiliateMeLinks`: AffiliateLinkMint
	fmt.Fprintf(os.Stdout, "Response from `AffiliateAPI.PostAffiliateMeLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAffiliateMeLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **affiliateCreateLinkRequest** | [**AffiliateCreateLinkRequest**](AffiliateCreateLinkRequest.md) |  | 

### Return type

[**AffiliateLinkMint**](AffiliateLinkMint.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

