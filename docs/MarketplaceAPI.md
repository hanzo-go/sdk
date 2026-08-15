# \MarketplaceAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketplaceListingsById**](MarketplaceAPI.md#DeleteMarketplaceListingsById) | **Delete** /v1/marketplace/listings/{id} | Unpublish withdraws one of the caller org&#39;s listings from the marketplace and answers 204.
[**GetMarketplace**](MarketplaceAPI.md#GetMarketplace) | **Get** /v1/marketplace | Discover lists every tool and agent the caller can reach in their own org and project, enriched with any public listing&#39;s title, category and price, and with installed&#x3D;true on the ones already activated for that scope.
[**GetMarketplaceListings**](MarketplaceAPI.md#GetMarketplaceListings) | **Get** /v1/marketplace/listings | Returns the listings the caller&#39;s own org has published — what this org is offering, not what it can buy.
[**PostMarketplaceInstall**](MarketplaceAPI.md#PostMarketplaceInstall) | **Post** /v1/marketplace/install | Install activates one tool for the caller&#39;s own org and project.
[**PostMarketplaceListings**](MarketplaceAPI.md#PostMarketplaceListings) | **Post** /v1/marketplace/listings | Publish offers one tool on the marketplace, optionally monetized.
[**PostMarketplaceUninstall**](MarketplaceAPI.md#PostMarketplaceUninstall) | **Post** /v1/marketplace/uninstall | Uninstall deactivates one tool for the caller&#39;s own org and project, so it stops being dispatchable there.



## DeleteMarketplaceListingsById

> DeleteMarketplaceListingsById(ctx, id).Execute()

Unpublish withdraws one of the caller org's listings from the marketplace and answers 204.



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
	id := "lst_1" // string | ID is the listing to unpublish, from the path.

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
**id** | **string** | ID is the listing to unpublish, from the path. | 

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketplace

> MarketCatalog GetMarketplace(ctx).Execute()

Discover lists every tool and agent the caller can reach in their own org and project, enriched with any public listing's title, category and price, and with installed=true on the ones already activated for that scope.



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
	resp, r, err := apiClient.MarketplaceAPI.GetMarketplace(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.GetMarketplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketplace`: MarketCatalog
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.GetMarketplace`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketplaceRequest struct via the builder pattern


### Return type

[**MarketCatalog**](MarketCatalog.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketplaceListings

> ListingPage GetMarketplaceListings(ctx).Execute()

Returns the listings the caller's own org has published — what this org is offering, not what it can buy.



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
	resp, r, err := apiClient.MarketplaceAPI.GetMarketplaceListings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.GetMarketplaceListings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketplaceListings`: ListingPage
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.GetMarketplaceListings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketplaceListingsRequest struct via the builder pattern


### Return type

[**ListingPage**](ListingPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketplaceInstall

> InstallState PostMarketplaceInstall(ctx).InstallReq(installReq).Execute()

Install activates one tool for the caller's own org and project.



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
	installReq := *openapiclient.NewInstallReq() // InstallReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PostMarketplaceInstall(context.Background()).InstallReq(installReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PostMarketplaceInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketplaceInstall`: InstallState
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PostMarketplaceInstall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketplaceInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **installReq** | [**InstallReq**](InstallReq.md) |  | 

### Return type

[**InstallState**](InstallState.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketplaceListings

> Listing PostMarketplaceListings(ctx).PublishReq(publishReq).Execute()

Publish offers one tool on the marketplace, optionally monetized.



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
	publishReq := *openapiclient.NewPublishReq() // PublishReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PostMarketplaceListings(context.Background()).PublishReq(publishReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PostMarketplaceListings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketplaceListings`: Listing
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PostMarketplaceListings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketplaceListingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publishReq** | [**PublishReq**](PublishReq.md) |  | 

### Return type

[**Listing**](Listing.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketplaceUninstall

> InstallState PostMarketplaceUninstall(ctx).InstallReq(installReq).Execute()

Uninstall deactivates one tool for the caller's own org and project, so it stops being dispatchable there.



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
	installReq := *openapiclient.NewInstallReq() // InstallReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.PostMarketplaceUninstall(context.Background()).InstallReq(installReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.PostMarketplaceUninstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketplaceUninstall`: InstallState
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.PostMarketplaceUninstall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketplaceUninstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **installReq** | [**InstallReq**](InstallReq.md) |  | 

### Return type

[**InstallState**](InstallState.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

