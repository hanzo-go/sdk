# \MarketplaceAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1MarketplaceListingsId**](MarketplaceAPI.md#CloudDeleteV1MarketplaceListingsId) | **Delete** /v1/marketplace/listings/{id} | Unpublish withdraws one of the caller org&#39;s listings from the marketplace and answers 204.
[**CloudGetV1Marketplace**](MarketplaceAPI.md#CloudGetV1Marketplace) | **Get** /v1/marketplace | Discover lists every tool and agent the caller can reach in their own org and project, enriched with any public listing&#39;s title, category and price, and with installed&#x3D;true on the ones already activated for that scope.
[**CloudGetV1MarketplaceListings**](MarketplaceAPI.md#CloudGetV1MarketplaceListings) | **Get** /v1/marketplace/listings | ListListings returns the listings the caller&#39;s own org has published — what this org is offering, not what it can buy.
[**CloudPostV1MarketplaceInstall**](MarketplaceAPI.md#CloudPostV1MarketplaceInstall) | **Post** /v1/marketplace/install | Install activates one tool for the caller&#39;s own org and project.
[**CloudPostV1MarketplaceListings**](MarketplaceAPI.md#CloudPostV1MarketplaceListings) | **Post** /v1/marketplace/listings | Publish offers one tool on the marketplace, optionally monetized.
[**CloudPostV1MarketplaceUninstall**](MarketplaceAPI.md#CloudPostV1MarketplaceUninstall) | **Post** /v1/marketplace/uninstall | Uninstall deactivates one tool for the caller&#39;s own org and project, so it stops being dispatchable there.



## CloudDeleteV1MarketplaceListingsId

> CloudDeleteV1MarketplaceListingsId(ctx, id).Execute()

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
	r, err := apiClient.MarketplaceAPI.CloudDeleteV1MarketplaceListingsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.CloudDeleteV1MarketplaceListingsId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1MarketplaceListingsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1Marketplace

> CloudMarketCatalog CloudGetV1Marketplace(ctx).Execute()

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
	resp, r, err := apiClient.MarketplaceAPI.CloudGetV1Marketplace(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.CloudGetV1Marketplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Marketplace`: CloudMarketCatalog
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.CloudGetV1Marketplace`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketplaceRequest struct via the builder pattern


### Return type

[**CloudMarketCatalog**](CloudMarketCatalog.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MarketplaceListings

> CloudListingPage CloudGetV1MarketplaceListings(ctx).Execute()

ListListings returns the listings the caller's own org has published — what this org is offering, not what it can buy.



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
	resp, r, err := apiClient.MarketplaceAPI.CloudGetV1MarketplaceListings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.CloudGetV1MarketplaceListings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketplaceListings`: CloudListingPage
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.CloudGetV1MarketplaceListings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketplaceListingsRequest struct via the builder pattern


### Return type

[**CloudListingPage**](CloudListingPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MarketplaceInstall

> CloudInstallState CloudPostV1MarketplaceInstall(ctx).CloudInstallReq(cloudInstallReq).Execute()

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
	cloudInstallReq := *openapiclient.NewCloudInstallReq() // CloudInstallReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.CloudPostV1MarketplaceInstall(context.Background()).CloudInstallReq(cloudInstallReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.CloudPostV1MarketplaceInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MarketplaceInstall`: CloudInstallState
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.CloudPostV1MarketplaceInstall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MarketplaceInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudInstallReq** | [**CloudInstallReq**](CloudInstallReq.md) |  | 

### Return type

[**CloudInstallState**](CloudInstallState.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MarketplaceListings

> CloudListing CloudPostV1MarketplaceListings(ctx).CloudPublishReq(cloudPublishReq).Execute()

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
	cloudPublishReq := *openapiclient.NewCloudPublishReq() // CloudPublishReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.CloudPostV1MarketplaceListings(context.Background()).CloudPublishReq(cloudPublishReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.CloudPostV1MarketplaceListings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MarketplaceListings`: CloudListing
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.CloudPostV1MarketplaceListings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MarketplaceListingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudPublishReq** | [**CloudPublishReq**](CloudPublishReq.md) |  | 

### Return type

[**CloudListing**](CloudListing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MarketplaceUninstall

> CloudInstallState CloudPostV1MarketplaceUninstall(ctx).CloudInstallReq(cloudInstallReq).Execute()

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
	cloudInstallReq := *openapiclient.NewCloudInstallReq() // CloudInstallReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.CloudPostV1MarketplaceUninstall(context.Background()).CloudInstallReq(cloudInstallReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.CloudPostV1MarketplaceUninstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MarketplaceUninstall`: CloudInstallState
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.CloudPostV1MarketplaceUninstall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MarketplaceUninstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudInstallReq** | [**CloudInstallReq**](CloudInstallReq.md) |  | 

### Return type

[**CloudInstallState**](CloudInstallState.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

