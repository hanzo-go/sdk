# \StoreAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteStoreByStoreid**](StoreAPI.md#DeleteStoreByStoreid) | **Delete** /v1/store/{storeid} | Delete a storefront, keeping a recoverable copy
[**DeleteStoreByStoreidListingByKey**](StoreAPI.md#DeleteStoreByStoreidListingByKey) | **Delete** /v1/store/{storeid}/listing/{key} | Remove a listing override
[**GetStore**](StoreAPI.md#GetStore) | **Get** /v1/store/ | List your org&#39;s storefronts as a page
[**GetStoreAccess**](StoreAPI.md#GetStoreAccess) | **Get** /v1/store/access | Whether a store is entitled to trade, and why
[**GetStoreByStoreid**](StoreAPI.md#GetStoreByStoreid) | **Get** /v1/store/{storeid} | Fetch one storefront
[**GetStoreByStoreidBundleByKey**](StoreAPI.md#GetStoreByStoreidBundleByKey) | **Get** /v1/store/{storeid}/bundle/{key} | Fetch a bundle as this storefront sells it
[**GetStoreByStoreidListing**](StoreAPI.md#GetStoreByStoreidListing) | **Get** /v1/store/{storeid}/listing | The storefront&#39;s whole listing override map
[**GetStoreByStoreidListingByKey**](StoreAPI.md#GetStoreByStoreidListingByKey) | **Get** /v1/store/{storeid}/listing/{key} | Fetch one listing override, by item id or by its slug or SKU
[**GetStoreByStoreidProductByKey**](StoreAPI.md#GetStoreByStoreidProductByKey) | **Get** /v1/store/{storeid}/product/{key} | Fetch a product as this storefront sells it
[**GetStoreByStoreidVariantByKey**](StoreAPI.md#GetStoreByStoreidVariantByKey) | **Get** /v1/store/{storeid}/variant/{key} | Fetch a variant as this storefront sells it
[**GetStoreCurrent**](StoreAPI.md#GetStoreCurrent) | **Get** /v1/store/current | Resolve your org&#39;s active storefront without naming an id
[**PatchStoreByStoreid**](StoreAPI.md#PatchStoreByStoreid) | **Patch** /v1/store/{storeid} | Change part of a storefront
[**PatchStoreByStoreidListingByKey**](StoreAPI.md#PatchStoreByStoreidListingByKey) | **Patch** /v1/store/{storeid}/listing/{key} | Confirm a listing override exists and re-save the store
[**PostStore**](StoreAPI.md#PostStore) | **Post** /v1/store/ | Create a storefront
[**PostStoreByStoreid**](StoreAPI.md#PostStoreByStoreid) | **Post** /v1/store/{storeid} | Method-override tunnel for clients that cannot send PUT, PATCH or DELETE
[**PostStoreByStoreidAuthorize**](StoreAPI.md#PostStoreByStoreidAuthorize) | **Post** /v1/store/{storeid}/authorize | Authorize a new order against a storefront, holding the funds without settling them
[**PostStoreByStoreidAuthorizeByOrderid**](StoreAPI.md#PostStoreByStoreidAuthorizeByOrderid) | **Post** /v1/store/{storeid}/authorize/{orderid} | Authorize an order that already exists, holding the funds without settling them
[**PostStoreByStoreidCaptureByOrderid**](StoreAPI.md#PostStoreByStoreidCaptureByOrderid) | **Post** /v1/store/{storeid}/capture/{orderid} | Capture a previously authorized order and settle the payment
[**PostStoreByStoreidCharge**](StoreAPI.md#PostStoreByStoreidCharge) | **Post** /v1/store/{storeid}/charge | Authorize and capture a new order in one call
[**PostStoreByStoreidCheckoutAuthorize**](StoreAPI.md#PostStoreByStoreidCheckoutAuthorize) | **Post** /v1/store/{storeid}/checkout/authorize | Authorize a new order against a storefront, holding the funds — the checkout spelling
[**PostStoreByStoreidCheckoutAuthorizeByOrderid**](StoreAPI.md#PostStoreByStoreidCheckoutAuthorizeByOrderid) | **Post** /v1/store/{storeid}/checkout/authorize/{orderid} | Authorize an existing order, holding the funds — the checkout spelling
[**PostStoreByStoreidCheckoutCaptureByOrderid**](StoreAPI.md#PostStoreByStoreidCheckoutCaptureByOrderid) | **Post** /v1/store/{storeid}/checkout/capture/{orderid} | Capture a previously authorized order and settle it — the checkout spelling
[**PostStoreByStoreidCheckoutCharge**](StoreAPI.md#PostStoreByStoreidCheckoutCharge) | **Post** /v1/store/{storeid}/checkout/charge | Authorize and capture a new order in one call — the checkout spelling
[**PostStoreByStoreidCheckoutPaypalCancelByPaykey**](StoreAPI.md#PostStoreByStoreidCheckoutPaypalCancelByPaykey) | **Post** /v1/store/{storeid}/checkout/paypal/cancel/{payKey} | PayPal cancel by pay key — refuses, exactly as the unprefixed address does
[**PostStoreByStoreidCheckoutPaypalConfirmByPaykey**](StoreAPI.md#PostStoreByStoreidCheckoutPaypalConfirmByPaykey) | **Post** /v1/store/{storeid}/checkout/paypal/confirm/{payKey} | PayPal confirm by pay key — refuses, exactly as the unprefixed address does
[**PostStoreByStoreidCheckoutPaypalPay**](StoreAPI.md#PostStoreByStoreidCheckoutPaypalPay) | **Post** /v1/store/{storeid}/checkout/paypal/pay | Start a PayPal authorization for a new order — the checkout spelling
[**PostStoreByStoreidListingByKey**](StoreAPI.md#PostStoreByStoreidListingByKey) | **Post** /v1/store/{storeid}/listing/{key} | Add a listing override under a new key
[**PostStoreByStoreidPaypalCancelByPaykey**](StoreAPI.md#PostStoreByStoreidPaypalCancelByPaykey) | **Post** /v1/store/{storeid}/paypal/cancel/{payKey} | PayPal cancel by pay key — refuses, because a pay key alone does not identify the order
[**PostStoreByStoreidPaypalConfirmByPaykey**](StoreAPI.md#PostStoreByStoreidPaypalConfirmByPaykey) | **Post** /v1/store/{storeid}/paypal/confirm/{payKey} | PayPal confirm by pay key — refuses, because a pay key alone does not identify the order
[**PostStoreByStoreidPaypalPay**](StoreAPI.md#PostStoreByStoreidPaypalPay) | **Post** /v1/store/{storeid}/paypal/pay | Start a PayPal authorization for a new order
[**PostStoreByStoreidTrial**](StoreAPI.md#PostStoreByStoreidTrial) | **Post** /v1/store/{storeid}/trial | Start this store&#39;s no-card trial on the entry plan
[**PostStoreToken**](StoreAPI.md#PostStoreToken) | **Post** /v1/store/token | Mint your org&#39;s least-privilege storefront read key
[**PutStoreByStoreid**](StoreAPI.md#PutStoreByStoreid) | **Put** /v1/store/{storeid} | Replace a storefront outright
[**PutStoreByStoreidListingByKey**](StoreAPI.md#PutStoreByStoreidListingByKey) | **Put** /v1/store/{storeid}/listing/{key} | Upsert a listing override



## DeleteStoreByStoreid

> DeleteStoreByStoreid(ctx, storeid).Execute()

Delete a storefront, keeping a recoverable copy



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.DeleteStoreByStoreid(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.DeleteStoreByStoreid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStoreByStoreidRequest struct via the builder pattern


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


## DeleteStoreByStoreidListingByKey

> DeleteStoreByStoreidListingByKey(ctx, storeid, key).Execute()

Remove a listing override



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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.DeleteStoreByStoreidListingByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.DeleteStoreByStoreidListingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStoreByStoreidListingByKeyRequest struct via the builder pattern


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


## GetStore

> GetStore(ctx).Execute()

List your org's storefronts as a page



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
	r, err := apiClient.StoreAPI.GetStore(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreRequest struct via the builder pattern


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


## GetStoreAccess

> GetStoreAccess(ctx).Execute()

Whether a store is entitled to trade, and why



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
	r, err := apiClient.StoreAPI.GetStoreAccess(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetStoreAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreAccessRequest struct via the builder pattern


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


## GetStoreByStoreid

> GetStoreByStoreid(ctx, storeid).Execute()

Fetch one storefront



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.GetStoreByStoreid(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetStoreByStoreid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreByStoreidRequest struct via the builder pattern


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


## GetStoreByStoreidBundleByKey

> GetStoreByStoreidBundleByKey(ctx, storeid, key).Execute()

Fetch a bundle as this storefront sells it



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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.GetStoreByStoreidBundleByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetStoreByStoreidBundleByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreByStoreidBundleByKeyRequest struct via the builder pattern


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


## GetStoreByStoreidListing

> GetStoreByStoreidListing(ctx, storeid).Execute()

The storefront's whole listing override map



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.GetStoreByStoreidListing(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetStoreByStoreidListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreByStoreidListingRequest struct via the builder pattern


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


## GetStoreByStoreidListingByKey

> GetStoreByStoreidListingByKey(ctx, storeid, key).Execute()

Fetch one listing override, by item id or by its slug or SKU



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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.GetStoreByStoreidListingByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetStoreByStoreidListingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreByStoreidListingByKeyRequest struct via the builder pattern


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


## GetStoreByStoreidProductByKey

> GetStoreByStoreidProductByKey(ctx, storeid, key).Execute()

Fetch a product as this storefront sells it



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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.GetStoreByStoreidProductByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetStoreByStoreidProductByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreByStoreidProductByKeyRequest struct via the builder pattern


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


## GetStoreByStoreidVariantByKey

> GetStoreByStoreidVariantByKey(ctx, storeid, key).Execute()

Fetch a variant as this storefront sells it



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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.GetStoreByStoreidVariantByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetStoreByStoreidVariantByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreByStoreidVariantByKeyRequest struct via the builder pattern


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


## GetStoreCurrent

> GetStoreCurrent(ctx).Execute()

Resolve your org's active storefront without naming an id



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
	r, err := apiClient.StoreAPI.GetStoreCurrent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.GetStoreCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreCurrentRequest struct via the builder pattern


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


## PatchStoreByStoreid

> PatchStoreByStoreid(ctx, storeid).Execute()

Change part of a storefront



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PatchStoreByStoreid(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PatchStoreByStoreid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchStoreByStoreidRequest struct via the builder pattern


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


## PatchStoreByStoreidListingByKey

> PatchStoreByStoreidListingByKey(ctx, storeid, key).Execute()

Confirm a listing override exists and re-save the store



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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PatchStoreByStoreidListingByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PatchStoreByStoreidListingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchStoreByStoreidListingByKeyRequest struct via the builder pattern


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


## PostStore

> PostStore(ctx).Execute()

Create a storefront



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
	r, err := apiClient.StoreAPI.PostStore(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreRequest struct via the builder pattern


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


## PostStoreByStoreid

> PostStoreByStoreid(ctx, storeid).Execute()

Method-override tunnel for clients that cannot send PUT, PATCH or DELETE



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PostStoreByStoreid(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreByStoreid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreByStoreidRequest struct via the builder pattern


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


## PostStoreByStoreidAuthorize

> PostStoreByStoreidAuthorize(ctx, storeid).Execute()

Authorize a new order against a storefront, holding the funds without settling them



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PostStoreByStoreidAuthorize(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreByStoreidAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreByStoreidAuthorizeRequest struct via the builder pattern


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


## PostStoreByStoreidAuthorizeByOrderid

> PostStoreByStoreidAuthorizeByOrderid(ctx, storeid, orderid).Execute()

Authorize an order that already exists, holding the funds without settling them



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
	storeid := "storeid_example" // string | 
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PostStoreByStoreidAuthorizeByOrderid(context.Background(), storeid, orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreByStoreidAuthorizeByOrderid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreByStoreidAuthorizeByOrderidRequest struct via the builder pattern


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


## PostStoreByStoreidCaptureByOrderid

> PostStoreByStoreidCaptureByOrderid(ctx, storeid, orderid).Execute()

Capture a previously authorized order and settle the payment



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
	storeid := "storeid_example" // string | 
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PostStoreByStoreidCaptureByOrderid(context.Background(), storeid, orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreByStoreidCaptureByOrderid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreByStoreidCaptureByOrderidRequest struct via the builder pattern


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


## PostStoreByStoreidCharge

> PostStoreByStoreidCharge(ctx, storeid).Execute()

Authorize and capture a new order in one call



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PostStoreByStoreidCharge(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreByStoreidCharge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreByStoreidChargeRequest struct via the builder pattern


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


## PostStoreByStoreidCheckoutAuthorize

> PostStoreByStoreidCheckoutAuthorize(ctx, storeid).Execute()

Authorize a new order against a storefront, holding the funds — the checkout spelling



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PostStoreByStoreidCheckoutAuthorize(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreByStoreidCheckoutAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreByStoreidCheckoutAuthorizeRequest struct via the builder pattern


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


## PostStoreByStoreidCheckoutAuthorizeByOrderid

> PostStoreByStoreidCheckoutAuthorizeByOrderid(ctx, storeid, orderid).Execute()

Authorize an existing order, holding the funds — the checkout spelling



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
	storeid := "storeid_example" // string | 
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PostStoreByStoreidCheckoutAuthorizeByOrderid(context.Background(), storeid, orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreByStoreidCheckoutAuthorizeByOrderid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreByStoreidCheckoutAuthorizeByOrderidRequest struct via the builder pattern


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


## PostStoreByStoreidCheckoutCaptureByOrderid

> PostStoreByStoreidCheckoutCaptureByOrderid(ctx, storeid, orderid).Execute()

Capture a previously authorized order and settle it — the checkout spelling



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
	storeid := "storeid_example" // string | 
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PostStoreByStoreidCheckoutCaptureByOrderid(context.Background(), storeid, orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreByStoreidCheckoutCaptureByOrderid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreByStoreidCheckoutCaptureByOrderidRequest struct via the builder pattern


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


## PostStoreByStoreidCheckoutCharge

> PostStoreByStoreidCheckoutCharge(ctx, storeid).Execute()

Authorize and capture a new order in one call — the checkout spelling



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PostStoreByStoreidCheckoutCharge(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreByStoreidCheckoutCharge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreByStoreidCheckoutChargeRequest struct via the builder pattern


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


## PostStoreByStoreidCheckoutPaypalCancelByPaykey

> PostStoreByStoreidCheckoutPaypalCancelByPaykey(ctx, storeid, payKey).Execute()

PayPal cancel by pay key — refuses, exactly as the unprefixed address does



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
	storeid := "storeid_example" // string | 
	payKey := "payKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PostStoreByStoreidCheckoutPaypalCancelByPaykey(context.Background(), storeid, payKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreByStoreidCheckoutPaypalCancelByPaykey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**payKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreByStoreidCheckoutPaypalCancelByPaykeyRequest struct via the builder pattern


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


## PostStoreByStoreidCheckoutPaypalConfirmByPaykey

> PostStoreByStoreidCheckoutPaypalConfirmByPaykey(ctx, storeid, payKey).Execute()

PayPal confirm by pay key — refuses, exactly as the unprefixed address does



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
	storeid := "storeid_example" // string | 
	payKey := "payKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PostStoreByStoreidCheckoutPaypalConfirmByPaykey(context.Background(), storeid, payKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreByStoreidCheckoutPaypalConfirmByPaykey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**payKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreByStoreidCheckoutPaypalConfirmByPaykeyRequest struct via the builder pattern


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


## PostStoreByStoreidCheckoutPaypalPay

> PostStoreByStoreidCheckoutPaypalPay(ctx, storeid).Execute()

Start a PayPal authorization for a new order — the checkout spelling



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PostStoreByStoreidCheckoutPaypalPay(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreByStoreidCheckoutPaypalPay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreByStoreidCheckoutPaypalPayRequest struct via the builder pattern


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


## PostStoreByStoreidListingByKey

> PostStoreByStoreidListingByKey(ctx, storeid, key).Execute()

Add a listing override under a new key



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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PostStoreByStoreidListingByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreByStoreidListingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreByStoreidListingByKeyRequest struct via the builder pattern


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


## PostStoreByStoreidPaypalCancelByPaykey

> PostStoreByStoreidPaypalCancelByPaykey(ctx, storeid, payKey).Execute()

PayPal cancel by pay key — refuses, because a pay key alone does not identify the order



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
	storeid := "storeid_example" // string | 
	payKey := "payKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PostStoreByStoreidPaypalCancelByPaykey(context.Background(), storeid, payKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreByStoreidPaypalCancelByPaykey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**payKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreByStoreidPaypalCancelByPaykeyRequest struct via the builder pattern


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


## PostStoreByStoreidPaypalConfirmByPaykey

> PostStoreByStoreidPaypalConfirmByPaykey(ctx, storeid, payKey).Execute()

PayPal confirm by pay key — refuses, because a pay key alone does not identify the order



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
	storeid := "storeid_example" // string | 
	payKey := "payKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PostStoreByStoreidPaypalConfirmByPaykey(context.Background(), storeid, payKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreByStoreidPaypalConfirmByPaykey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**payKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreByStoreidPaypalConfirmByPaykeyRequest struct via the builder pattern


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


## PostStoreByStoreidPaypalPay

> PostStoreByStoreidPaypalPay(ctx, storeid).Execute()

Start a PayPal authorization for a new order



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PostStoreByStoreidPaypalPay(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreByStoreidPaypalPay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreByStoreidPaypalPayRequest struct via the builder pattern


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


## PostStoreByStoreidTrial

> PostStoreByStoreidTrial(ctx, storeid).Execute()

Start this store's no-card trial on the entry plan



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PostStoreByStoreidTrial(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreByStoreidTrial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreByStoreidTrialRequest struct via the builder pattern


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


## PostStoreToken

> PostStoreToken(ctx).Execute()

Mint your org's least-privilege storefront read key



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
	r, err := apiClient.StoreAPI.PostStoreToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PostStoreToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostStoreTokenRequest struct via the builder pattern


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


## PutStoreByStoreid

> PutStoreByStoreid(ctx, storeid).Execute()

Replace a storefront outright



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PutStoreByStoreid(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PutStoreByStoreid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutStoreByStoreidRequest struct via the builder pattern


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


## PutStoreByStoreidListingByKey

> PutStoreByStoreidListingByKey(ctx, storeid, key).Execute()

Upsert a listing override



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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.PutStoreByStoreidListingByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.PutStoreByStoreidListingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutStoreByStoreidListingByKeyRequest struct via the builder pattern


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

