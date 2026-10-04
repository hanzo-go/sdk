# \MarketplaceAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketplaceListingsById**](MarketplaceAPI.md#DeleteMarketplaceListingsById) | **Delete** /v1/marketplace/listings/{id} | Withdraws one of the caller org&#39;s listings from the marketplace and answers 204.
[**GetMarketplace**](MarketplaceAPI.md#GetMarketplace) | **Get** /v1/marketplace | Lists every tool and agent the caller can reach in their own org and project, enriched with any public listing&#39;s title, category and price, and with installed&#x3D;true on the ones already activated for that scope.
[**GetMarketplaceJobs**](MarketplaceAPI.md#GetMarketplaceJobs) | **Get** /v1/marketplace/jobs | Lists the jobs the caller&#39;s org is a party to, newest first: as buyer, the work it hired — with the quotes it has not paid and the jobs being funded, each with the attempt that made it, so a hire whose answer was lost is found here — and as seller, the work it was hired for, which never includes another org&#39;s quote.
[**GetMarketplaceJobsById**](MarketplaceAPI.md#GetMarketplaceJobsById) | **Get** /v1/marketplace/jobs/{id} | Reads one job the caller&#39;s org is a party to.
[**GetMarketplaceListings**](MarketplaceAPI.md#GetMarketplaceListings) | **Get** /v1/marketplace/listings | Returns the listings the caller&#39;s own org has published — what this org is offering, not what it can buy.
[**GetMarketplaceSeller**](MarketplaceAPI.md#GetMarketplaceSeller) | **Get** /v1/marketplace/seller | Answers where the caller&#39;s org stands as a seller, in one read: its founders&#39; identity verification (KYC) and legal entity (KYB), its tax form and whether it is certified and valid, its own sanctions screening — only what an org may see about itself — the payout wallet it proved, the TaxPrincipalCredentials signed for its agents, what it earned in the year from the economic events the rails stated, and the 1099s payers furnished it.
[**GetMarketplaceShop**](MarketplaceAPI.md#GetMarketplaceShop) | **Get** /v1/marketplace/shop | Searches every public listing of every kind — agents and personas to hire, apps, skills, MCP servers and tools — newest first, with facets by kind, category, price and rating, paged.
[**GetMarketplaceShopById**](MarketplaceAPI.md#GetMarketplaceShopById) | **Get** /v1/marketplace/shop/{id} | Reads one public listing as the shop shows it — its seller&#39;s public face, its reputation and the ways to buy it.
[**PatchMarketplaceListingsById**](MarketplaceAPI.md#PatchMarketplaceListingsById) | **Patch** /v1/marketplace/listings/{id} | Edits one of the caller org&#39;s listings — its copy, price, payout wallet, visibility and documentation — under the same rules publish applies, and answers the listing as it now stands.
[**PostMarketplaceInstall**](MarketplaceAPI.md#PostMarketplaceInstall) | **Post** /v1/marketplace/install | Activates one tool for the caller&#39;s own org and project.
[**PostMarketplaceJobs**](MarketplaceAPI.md#PostMarketplaceJobs) | **Post** /v1/marketplace/jobs | Hires another org for a piece of work, through a public listing or by a direct offer, and answers 201 with the job — open, waiting for the seller.
[**PostMarketplaceJobsByIdAccept**](MarketplaceAPI.md#PostMarketplaceJobsByIdAccept) | **Post** /v1/marketplace/jobs/{id}/accept | Accepts a job the caller&#39;s org was hired for: the seller takes the work on and the clock toward its deadline is the seller&#39;s.
[**PostMarketplaceJobsByIdCancel**](MarketplaceAPI.md#PostMarketplaceJobsByIdCancel) | **Post** /v1/marketplace/jobs/{id}/cancel | Takes back a job the caller&#39;s org opened, before the seller accepts it.
[**PostMarketplaceJobsByIdDecline**](MarketplaceAPI.md#PostMarketplaceJobsByIdDecline) | **Post** /v1/marketplace/jobs/{id}/decline | Declines a job the caller&#39;s org was hired for, before any work.
[**PostMarketplaceJobsByIdDeliver**](MarketplaceAPI.md#PostMarketplaceJobsByIdDeliver) | **Post** /v1/marketplace/jobs/{id}/deliver | Records delivery of a job the caller&#39;s org accepted, before its deadline, and starts the review window: the buyer releases or disputes within it, or it releases itself when it closes.
[**PostMarketplaceJobsByIdDispute**](MarketplaceAPI.md#PostMarketplaceJobsByIdDispute) | **Post** /v1/marketplace/jobs/{id}/dispute | Stops a job for a ruling, as either party: before delivery (and before the deadline), or within the review window after it.
[**PostMarketplaceJobsByIdFeedback**](MarketplaceAPI.md#PostMarketplaceJobsByIdFeedback) | **Post** /v1/marketplace/jobs/{id}/feedback | Rates the other party of a settled job — the seller when the caller&#39;s org bought, the buyer when it sold — once per party per job, and never edited.
[**PostMarketplaceJobsByIdRefund**](MarketplaceAPI.md#PostMarketplaceJobsByIdRefund) | **Post** /v1/marketplace/jobs/{id}/refund | Refunds a job the caller&#39;s org accepted, as the seller, at any point before it is paid: nothing was moved, so the buyer&#39;s authorization is given up, never settled, and the amount it set aside returns to the buyer&#39;s wallet.
[**PostMarketplaceJobsByIdRelease**](MarketplaceAPI.md#PostMarketplaceJobsByIdRelease) | **Post** /v1/marketplace/jobs/{id}/release | Releases a job the caller&#39;s org is paying for, paying the seller the whole amount: the buyer&#39;s authorization is settled on the rail, once, and the rail states the payment.
[**PostMarketplaceListings**](MarketplaceAPI.md#PostMarketplaceListings) | **Post** /v1/marketplace/listings | Offers one thing the caller&#39;s org owns on the marketplace — an agent, persona, app, skill or MCP server, or, for the platform, a tool sold per call — optionally monetized.
[**PostMarketplaceSellerPayout**](MarketplaceAPI.md#PostMarketplaceSellerPayout) | **Post** /v1/marketplace/seller/payout | Starts binding one of the caller org&#39;s wallets as the wallet it is paid into: answers a challenge naming the org, the wallet and its address, to be signed with that wallet within fifteen minutes and sent to POST /v1/marketplace/seller/payout/verify.
[**PostMarketplaceSellerPayoutVerify**](MarketplaceAPI.md#PostMarketplaceSellerPayoutVerify) | **Post** /v1/marketplace/seller/payout/verify | Binds the caller org&#39;s payout wallet: the signature over its outstanding challenge is recovered, and must recover to the wallet&#39;s address, before the challenge expires.
[**PostMarketplaceUninstall**](MarketplaceAPI.md#PostMarketplaceUninstall) | **Post** /v1/marketplace/uninstall | Deactivates one tool for the caller&#39;s own org and project, so it stops being dispatchable there.



## DeleteMarketplaceListingsById

> DeleteMarketplaceListingsById(ctx, id).Execute()

Withdraws one of the caller org's listings from the marketplace and answers 204.



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
	id := "lst_1" // string | ID is the listing, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketplaceAPI.DeleteMarketplaceListingsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.DeleteMarketplaceListingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the listing, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketplaceListingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketplace

> MarketplaceMarketCatalog GetMarketplace(ctx).Execute()

Lists every tool and agent the caller can reach in their own org and project, enriched with any public listing's title, category and price, and with installed=true on the ones already activated for that scope.



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
	resp, r, err := apiClient.MarketplaceAPI.GetMarketplace(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.GetMarketplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketplace`: MarketplaceMarketCatalog
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.GetMarketplace`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketplaceRequest struct via the builder pattern


### Return type

[**MarketplaceMarketCatalog**](MarketplaceMarketCatalog.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketplaceJobs

> MarketplaceJobPage GetMarketplaceJobs(ctx).Role(role).Status(status).Execute()

Lists the jobs the caller's org is a party to, newest first: as buyer, the work it hired — with the quotes it has not paid and the jobs being funded, each with the attempt that made it, so a hire whose answer was lost is found here — and as seller, the work it was hired for, which never includes another org's quote.



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
	role := "seller" // string | Role is buyer — the jobs the caller's org hired for — or seller, the jobs it was hired for. Buyer when empty. (optional)
	status := "open" // string | Status keeps one state. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.GetMarketplaceJobs(context.Background()).Role(role).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.GetMarketplaceJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketplaceJobs`: MarketplaceJobPage
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.GetMarketplaceJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketplaceJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | **string** | Role is buyer — the jobs the caller&#39;s org hired for — or seller, the jobs it was hired for. Buyer when empty. | 
 **status** | **string** | Status keeps one state. | 

### Return type

[**MarketplaceJobPage**](MarketplaceJobPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketplaceJobsById

> MarketplaceJob GetMarketplaceJobsById(ctx, id).Execute()

Reads one job the caller's org is a party to.



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
	id := "job_9a1f" // string | ID is the job, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.GetMarketplaceJobsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.GetMarketplaceJobsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketplaceJobsById`: MarketplaceJob
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.GetMarketplaceJobsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the job, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketplaceJobsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MarketplaceJob**](MarketplaceJob.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketplaceListings

> MarketplaceListingPage GetMarketplaceListings(ctx).Execute()

Returns the listings the caller's own org has published — what this org is offering, not what it can buy.



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
	resp, r, err := apiClient.MarketplaceAPI.GetMarketplaceListings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.GetMarketplaceListings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketplaceListings`: MarketplaceListingPage
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.GetMarketplaceListings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketplaceListingsRequest struct via the builder pattern


### Return type

[**MarketplaceListingPage**](MarketplaceListingPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketplaceSeller

> MarketplaceOnboarding GetMarketplaceSeller(ctx).Year(year).Execute()

Answers where the caller's org stands as a seller, in one read: its founders' identity verification (KYC) and legal entity (KYB), its tax form and whether it is certified and valid, its own sanctions screening — only what an org may see about itself — the payout wallet it proved, the TaxPrincipalCredentials signed for its agents, what it earned in the year from the economic events the rails stated, and the 1099s payers furnished it.



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
	year := int64(2026) // int64 | Year is the calendar year (UTC); the current one when zero. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.GetMarketplaceSeller(context.Background()).Year(year).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.GetMarketplaceSeller``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketplaceSeller`: MarketplaceOnboarding
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.GetMarketplaceSeller`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketplaceSellerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **int64** | Year is the calendar year (UTC); the current one when zero. | 

### Return type

[**MarketplaceOnboarding**](MarketplaceOnboarding.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketplaceShop

> MarketplaceShop GetMarketplaceShop(ctx).Q(q).Kind(kind).Category(category).Price(price).Rating(rating).Seller(seller).Limit(limit).Offset(offset).Execute()

Searches every public listing of every kind — agents and personas to hire, apps, skills, MCP servers and tools — newest first, with facets by kind, category, price and rating, paged.



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
	q := "research" // string | Q keeps listings whose title, description, category, thing or seller contains every word of it, case-insensitively. (optional)
	kind := "agent" // string | Kind keeps one kind: agent, persona, app, skill, mcp or tool. (optional)
	category := "category_example" // string | Category keeps one category, exactly. (optional)
	price := "price_example" // string | Price keeps free listings or priced ones. (optional)
	rating := int64(789) // int64 | Rating keeps listings rated at least this many stars, 1 to 5. (optional)
	seller := "seller_example" // string | Seller keeps one seller org's listings. (optional)
	limit := int64(48) // int64 | Limit is the page size: 48 by default, 200 at most. (optional)
	offset := int64(789) // int64 | Offset is where the page starts. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.GetMarketplaceShop(context.Background()).Q(q).Kind(kind).Category(category).Price(price).Rating(rating).Seller(seller).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.GetMarketplaceShop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketplaceShop`: MarketplaceShop
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.GetMarketplaceShop`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketplaceShopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Q keeps listings whose title, description, category, thing or seller contains every word of it, case-insensitively. | 
 **kind** | **string** | Kind keeps one kind: agent, persona, app, skill, mcp or tool. | 
 **category** | **string** | Category keeps one category, exactly. | 
 **price** | **string** | Price keeps free listings or priced ones. | 
 **rating** | **int64** | Rating keeps listings rated at least this many stars, 1 to 5. | 
 **seller** | **string** | Seller keeps one seller org&#39;s listings. | 
 **limit** | **int64** | Limit is the page size: 48 by default, 200 at most. | 
 **offset** | **int64** | Offset is where the page starts. | 

### Return type

[**MarketplaceShop**](MarketplaceShop.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketplaceShopById

> MarketplaceShopListing GetMarketplaceShopById(ctx, id).Execute()

Reads one public listing as the shop shows it — its seller's public face, its reputation and the ways to buy it.



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
	id := "lst_4f1c" // string | ID is the listing, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.GetMarketplaceShopById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.GetMarketplaceShopById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketplaceShopById`: MarketplaceShopListing
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.GetMarketplaceShopById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the listing, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketplaceShopByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MarketplaceShopListing**](MarketplaceShopListing.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMarketplaceListingsById

> MarketplaceListing PatchMarketplaceListingsById(ctx, id).MarketplacePatchReq(marketplacePatchReq).Execute()

Edits one of the caller org's listings — its copy, price, payout wallet, visibility and documentation — under the same rules publish applies, and answers the listing as it now stands.



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
	id := "lst_4f1c" // string | ID is the listing to edit, from the path.
	marketplacePatchReq := *openapiclient.NewMarketplacePatchReq() // MarketplacePatchReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PatchMarketplaceListingsById(context.Background(), id).MarketplacePatchReq(marketplacePatchReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PatchMarketplaceListingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMarketplaceListingsById`: MarketplaceListing
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PatchMarketplaceListingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the listing to edit, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketplaceListingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketplacePatchReq** | [**MarketplacePatchReq**](MarketplacePatchReq.md) |  | 

### Return type

[**MarketplaceListing**](MarketplaceListing.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketplaceInstall

> MarketplaceInstallState PostMarketplaceInstall(ctx).MarketplaceInstallReq(marketplaceInstallReq).Execute()

Activates one tool for the caller's own org and project.



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
	marketplaceInstallReq := *openapiclient.NewMarketplaceInstallReq() // MarketplaceInstallReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PostMarketplaceInstall(context.Background()).MarketplaceInstallReq(marketplaceInstallReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PostMarketplaceInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketplaceInstall`: MarketplaceInstallState
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PostMarketplaceInstall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketplaceInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceInstallReq** | [**MarketplaceInstallReq**](MarketplaceInstallReq.md) |  | 

### Return type

[**MarketplaceInstallState**](MarketplaceInstallState.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketplaceJobs

> MarketplaceJob PostMarketplaceJobs(ctx).MarketplaceHireIn(marketplaceHireIn).Execute()

Hires another org for a piece of work, through a public listing or by a direct offer, and answers 201 with the job — open, waiting for the seller.



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
	marketplaceHireIn := *openapiclient.NewMarketplaceHireIn() // MarketplaceHireIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PostMarketplaceJobs(context.Background()).MarketplaceHireIn(marketplaceHireIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PostMarketplaceJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketplaceJobs`: MarketplaceJob
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PostMarketplaceJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketplaceJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceHireIn** | [**MarketplaceHireIn**](MarketplaceHireIn.md) |  | 

### Return type

[**MarketplaceJob**](MarketplaceJob.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketplaceJobsByIdAccept

> MarketplaceJob PostMarketplaceJobsByIdAccept(ctx, id).Execute()

Accepts a job the caller's org was hired for: the seller takes the work on and the clock toward its deadline is the seller's.



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
	id := "job_9a1f" // string | ID is the job, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PostMarketplaceJobsByIdAccept(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PostMarketplaceJobsByIdAccept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketplaceJobsByIdAccept`: MarketplaceJob
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PostMarketplaceJobsByIdAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the job, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketplaceJobsByIdAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MarketplaceJob**](MarketplaceJob.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketplaceJobsByIdCancel

> MarketplaceJob PostMarketplaceJobsByIdCancel(ctx, id).Execute()

Takes back a job the caller's org opened, before the seller accepts it.



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
	id := "job_9a1f" // string | ID is the job, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PostMarketplaceJobsByIdCancel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PostMarketplaceJobsByIdCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketplaceJobsByIdCancel`: MarketplaceJob
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PostMarketplaceJobsByIdCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the job, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketplaceJobsByIdCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MarketplaceJob**](MarketplaceJob.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketplaceJobsByIdDecline

> MarketplaceJob PostMarketplaceJobsByIdDecline(ctx, id).MarketplaceDeclineIn(marketplaceDeclineIn).Execute()

Declines a job the caller's org was hired for, before any work.



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
	id := "job_9a1f" // string | ID is the job, from the path.
	marketplaceDeclineIn := *openapiclient.NewMarketplaceDeclineIn() // MarketplaceDeclineIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PostMarketplaceJobsByIdDecline(context.Background(), id).MarketplaceDeclineIn(marketplaceDeclineIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PostMarketplaceJobsByIdDecline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketplaceJobsByIdDecline`: MarketplaceJob
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PostMarketplaceJobsByIdDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the job, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketplaceJobsByIdDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketplaceDeclineIn** | [**MarketplaceDeclineIn**](MarketplaceDeclineIn.md) |  | 

### Return type

[**MarketplaceJob**](MarketplaceJob.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketplaceJobsByIdDeliver

> MarketplaceJob PostMarketplaceJobsByIdDeliver(ctx, id).MarketplaceDeliverIn(marketplaceDeliverIn).Execute()

Records delivery of a job the caller's org accepted, before its deadline, and starts the review window: the buyer releases or disputes within it, or it releases itself when it closes.



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
	id := "job_9a1f" // string | ID is the job, from the path.
	marketplaceDeliverIn := *openapiclient.NewMarketplaceDeliverIn() // MarketplaceDeliverIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PostMarketplaceJobsByIdDeliver(context.Background(), id).MarketplaceDeliverIn(marketplaceDeliverIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PostMarketplaceJobsByIdDeliver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketplaceJobsByIdDeliver`: MarketplaceJob
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PostMarketplaceJobsByIdDeliver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the job, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketplaceJobsByIdDeliverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketplaceDeliverIn** | [**MarketplaceDeliverIn**](MarketplaceDeliverIn.md) |  | 

### Return type

[**MarketplaceJob**](MarketplaceJob.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketplaceJobsByIdDispute

> MarketplaceJob PostMarketplaceJobsByIdDispute(ctx, id).MarketplaceDisputeIn(marketplaceDisputeIn).Execute()

Stops a job for a ruling, as either party: before delivery (and before the deadline), or within the review window after it.



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
	id := "job_9a1f" // string | ID is the job, from the path.
	marketplaceDisputeIn := *openapiclient.NewMarketplaceDisputeIn() // MarketplaceDisputeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PostMarketplaceJobsByIdDispute(context.Background(), id).MarketplaceDisputeIn(marketplaceDisputeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PostMarketplaceJobsByIdDispute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketplaceJobsByIdDispute`: MarketplaceJob
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PostMarketplaceJobsByIdDispute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the job, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketplaceJobsByIdDisputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketplaceDisputeIn** | [**MarketplaceDisputeIn**](MarketplaceDisputeIn.md) |  | 

### Return type

[**MarketplaceJob**](MarketplaceJob.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketplaceJobsByIdFeedback

> MarketplaceFeedback PostMarketplaceJobsByIdFeedback(ctx, id).MarketplaceFeedbackIn(marketplaceFeedbackIn).Execute()

Rates the other party of a settled job — the seller when the caller's org bought, the buyer when it sold — once per party per job, and never edited.



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
	id := "job_9a1f" // string | ID is the job, from the path.
	marketplaceFeedbackIn := *openapiclient.NewMarketplaceFeedbackIn() // MarketplaceFeedbackIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PostMarketplaceJobsByIdFeedback(context.Background(), id).MarketplaceFeedbackIn(marketplaceFeedbackIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PostMarketplaceJobsByIdFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketplaceJobsByIdFeedback`: MarketplaceFeedback
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PostMarketplaceJobsByIdFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the job, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketplaceJobsByIdFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketplaceFeedbackIn** | [**MarketplaceFeedbackIn**](MarketplaceFeedbackIn.md) |  | 

### Return type

[**MarketplaceFeedback**](MarketplaceFeedback.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketplaceJobsByIdRefund

> MarketplaceJob PostMarketplaceJobsByIdRefund(ctx, id).Execute()

Refunds a job the caller's org accepted, as the seller, at any point before it is paid: nothing was moved, so the buyer's authorization is given up, never settled, and the amount it set aside returns to the buyer's wallet.



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
	id := "job_9a1f" // string | ID is the job, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PostMarketplaceJobsByIdRefund(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PostMarketplaceJobsByIdRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketplaceJobsByIdRefund`: MarketplaceJob
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PostMarketplaceJobsByIdRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the job, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketplaceJobsByIdRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MarketplaceJob**](MarketplaceJob.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketplaceJobsByIdRelease

> MarketplaceJob PostMarketplaceJobsByIdRelease(ctx, id).Execute()

Releases a job the caller's org is paying for, paying the seller the whole amount: the buyer's authorization is settled on the rail, once, and the rail states the payment.



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
	id := "job_9a1f" // string | ID is the job, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PostMarketplaceJobsByIdRelease(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PostMarketplaceJobsByIdRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketplaceJobsByIdRelease`: MarketplaceJob
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PostMarketplaceJobsByIdRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the job, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketplaceJobsByIdReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MarketplaceJob**](MarketplaceJob.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketplaceListings

> MarketplaceListing PostMarketplaceListings(ctx).MarketplacePublishReq(marketplacePublishReq).Execute()

Offers one thing the caller's org owns on the marketplace — an agent, persona, app, skill or MCP server, or, for the platform, a tool sold per call — optionally monetized.



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
	marketplacePublishReq := *openapiclient.NewMarketplacePublishReq() // MarketplacePublishReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PostMarketplaceListings(context.Background()).MarketplacePublishReq(marketplacePublishReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PostMarketplaceListings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketplaceListings`: MarketplaceListing
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PostMarketplaceListings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketplaceListingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplacePublishReq** | [**MarketplacePublishReq**](MarketplacePublishReq.md) |  | 

### Return type

[**MarketplaceListing**](MarketplaceListing.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketplaceSellerPayout

> MarketplacePayoutChallenge PostMarketplaceSellerPayout(ctx).MarketplacePayoutIn(marketplacePayoutIn).Execute()

Starts binding one of the caller org's wallets as the wallet it is paid into: answers a challenge naming the org, the wallet and its address, to be signed with that wallet within fifteen minutes and sent to POST /v1/marketplace/seller/payout/verify.



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
	marketplacePayoutIn := *openapiclient.NewMarketplacePayoutIn() // MarketplacePayoutIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PostMarketplaceSellerPayout(context.Background()).MarketplacePayoutIn(marketplacePayoutIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PostMarketplaceSellerPayout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketplaceSellerPayout`: MarketplacePayoutChallenge
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PostMarketplaceSellerPayout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketplaceSellerPayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplacePayoutIn** | [**MarketplacePayoutIn**](MarketplacePayoutIn.md) |  | 

### Return type

[**MarketplacePayoutChallenge**](MarketplacePayoutChallenge.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketplaceSellerPayoutVerify

> MarketplacePayout PostMarketplaceSellerPayoutVerify(ctx).MarketplaceVerifyIn(marketplaceVerifyIn).Execute()

Binds the caller org's payout wallet: the signature over its outstanding challenge is recovered, and must recover to the wallet's address, before the challenge expires.



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
	marketplaceVerifyIn := *openapiclient.NewMarketplaceVerifyIn() // MarketplaceVerifyIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PostMarketplaceSellerPayoutVerify(context.Background()).MarketplaceVerifyIn(marketplaceVerifyIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PostMarketplaceSellerPayoutVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketplaceSellerPayoutVerify`: MarketplacePayout
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PostMarketplaceSellerPayoutVerify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketplaceSellerPayoutVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceVerifyIn** | [**MarketplaceVerifyIn**](MarketplaceVerifyIn.md) |  | 

### Return type

[**MarketplacePayout**](MarketplacePayout.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketplaceUninstall

> MarketplaceInstallState PostMarketplaceUninstall(ctx).MarketplaceInstallReq(marketplaceInstallReq).Execute()

Deactivates one tool for the caller's own org and project, so it stops being dispatchable there.



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
	marketplaceInstallReq := *openapiclient.NewMarketplaceInstallReq() // MarketplaceInstallReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PostMarketplaceUninstall(context.Background()).MarketplaceInstallReq(marketplaceInstallReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PostMarketplaceUninstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketplaceUninstall`: MarketplaceInstallState
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PostMarketplaceUninstall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketplaceUninstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketplaceInstallReq** | [**MarketplaceInstallReq**](MarketplaceInstallReq.md) |  | 

### Return type

[**MarketplaceInstallState**](MarketplaceInstallState.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

