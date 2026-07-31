# \StoreAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommerceCreateStore**](StoreAPI.md#CommerceCreateStore) | **Post** /v1/commerce/store | Create store
[**CommerceCreateStoreListing**](StoreAPI.md#CommerceCreateStoreListing) | **Post** /v1/commerce/store/{storeid}/listing/{key} | Create store listing
[**CommerceDeleteStoreListing**](StoreAPI.md#CommerceDeleteStoreListing) | **Delete** /v1/commerce/store/{storeid}/listing/{key} | Delete store listing
[**CommerceGetStore**](StoreAPI.md#CommerceGetStore) | **Get** /v1/commerce/store/{storeid} | Get store
[**CommerceGetStoreListing**](StoreAPI.md#CommerceGetStoreListing) | **Get** /v1/commerce/store/{storeid}/listing/{key} | Get store listing
[**CommerceGetStoreProduct**](StoreAPI.md#CommerceGetStoreProduct) | **Get** /v1/commerce/store/{storeid}/product/{key} | Get store product
[**CommerceGetStoreVariant**](StoreAPI.md#CommerceGetStoreVariant) | **Get** /v1/commerce/store/{storeid}/variant/{key} | Get store variant
[**CommerceListStoreListings**](StoreAPI.md#CommerceListStoreListings) | **Get** /v1/commerce/store/{storeid}/listing | List store listings
[**CommerceListStores**](StoreAPI.md#CommerceListStores) | **Get** /v1/commerce/store | List stores
[**CommercePatchStoreListing**](StoreAPI.md#CommercePatchStoreListing) | **Patch** /v1/commerce/store/{storeid}/listing/{key} | Partially update store listing
[**CommerceStoreAuthorize**](StoreAPI.md#CommerceStoreAuthorize) | **Post** /v1/commerce/store/{storeid}/checkout/authorize | Authorize payment via store
[**CommerceStoreCharge**](StoreAPI.md#CommerceStoreCharge) | **Post** /v1/commerce/store/{storeid}/checkout/charge | Charge payment via store
[**CommerceUpdateStoreListing**](StoreAPI.md#CommerceUpdateStoreListing) | **Put** /v1/commerce/store/{storeid}/listing/{key} | Update store listing



## CommerceCreateStore

> CommerceStore CommerceCreateStore(ctx).CommerceStore(commerceStore).Execute()

Create store

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
	commerceStore := *openapiclient.NewCommerceStore() // CommerceStore | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommerceCreateStore(context.Background()).CommerceStore(commerceStore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceCreateStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateStore`: CommerceStore
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceCreateStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceStore** | [**CommerceStore**](CommerceStore.md) |  | 

### Return type

[**CommerceStore**](CommerceStore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceCreateStoreListing

> CommerceListing CommerceCreateStoreListing(ctx, storeid, key).CommerceListing(commerceListing).Execute()

Create store listing

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
	commerceListing := *openapiclient.NewCommerceListing() // CommerceListing | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommerceCreateStoreListing(context.Background(), storeid, key).CommerceListing(commerceListing).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceCreateStoreListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateStoreListing`: CommerceListing
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceCreateStoreListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateStoreListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commerceListing** | [**CommerceListing**](CommerceListing.md) |  | 

### Return type

[**CommerceListing**](CommerceListing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceDeleteStoreListing

> CommerceDeleteStoreListing(ctx, storeid, key).Execute()

Delete store listing

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
	r, err := apiClient.StoreAPI.CommerceDeleteStoreListing(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceDeleteStoreListing``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCommerceDeleteStoreListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetStore

> CommerceStore CommerceGetStore(ctx, storeid).Execute()

Get store

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
	resp, r, err := apiClient.StoreAPI.CommerceGetStore(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceGetStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetStore`: CommerceStore
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceGetStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceStore**](CommerceStore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetStoreListing

> CommerceListing CommerceGetStoreListing(ctx, storeid, key).Execute()

Get store listing

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
	resp, r, err := apiClient.StoreAPI.CommerceGetStoreListing(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceGetStoreListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetStoreListing`: CommerceListing
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceGetStoreListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetStoreListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CommerceListing**](CommerceListing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetStoreProduct

> CommerceProduct CommerceGetStoreProduct(ctx, storeid, key).Execute()

Get store product

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
	resp, r, err := apiClient.StoreAPI.CommerceGetStoreProduct(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceGetStoreProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetStoreProduct`: CommerceProduct
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceGetStoreProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetStoreProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CommerceProduct**](CommerceProduct.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetStoreVariant

> CommerceVariant CommerceGetStoreVariant(ctx, storeid, key).Execute()

Get store variant

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
	resp, r, err := apiClient.StoreAPI.CommerceGetStoreVariant(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceGetStoreVariant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetStoreVariant`: CommerceVariant
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceGetStoreVariant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetStoreVariantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CommerceVariant**](CommerceVariant.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceListStoreListings

> []CommerceListing CommerceListStoreListings(ctx, storeid).Execute()

List store listings

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
	resp, r, err := apiClient.StoreAPI.CommerceListStoreListings(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceListStoreListings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceListStoreListings`: []CommerceListing
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceListStoreListings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceListStoreListingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CommerceListing**](CommerceListing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceListStores

> CommercePaginatedStores CommerceListStores(ctx).Page(page).Display(display).Execute()

List stores

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
	page := int32(56) // int32 | Page number (1-indexed) (optional) (default to 1)
	display := int32(56) // int32 | Number of items per page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommerceListStores(context.Background()).Page(page).Display(display).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceListStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceListStores`: CommercePaginatedStores
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceListStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceListStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (1-indexed) | [default to 1]
 **display** | **int32** | Number of items per page | [default to 20]

### Return type

[**CommercePaginatedStores**](CommercePaginatedStores.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommercePatchStoreListing

> CommerceListing CommercePatchStoreListing(ctx, storeid, key).CommerceListing(commerceListing).Execute()

Partially update store listing

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
	commerceListing := *openapiclient.NewCommerceListing() // CommerceListing | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommercePatchStoreListing(context.Background(), storeid, key).CommerceListing(commerceListing).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommercePatchStoreListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommercePatchStoreListing`: CommerceListing
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommercePatchStoreListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommercePatchStoreListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commerceListing** | [**CommerceListing**](CommerceListing.md) |  | 

### Return type

[**CommerceListing**](CommerceListing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceStoreAuthorize

> CommerceOrder CommerceStoreAuthorize(ctx, storeid).CommerceCheckoutRequest(commerceCheckoutRequest).Execute()

Authorize payment via store

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
	commerceCheckoutRequest := *openapiclient.NewCommerceCheckoutRequest() // CommerceCheckoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommerceStoreAuthorize(context.Background(), storeid).CommerceCheckoutRequest(commerceCheckoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceStoreAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceStoreAuthorize`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceStoreAuthorize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceStoreAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceCheckoutRequest** | [**CommerceCheckoutRequest**](CommerceCheckoutRequest.md) |  | 

### Return type

[**CommerceOrder**](CommerceOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceStoreCharge

> CommerceOrder CommerceStoreCharge(ctx, storeid).CommerceCheckoutRequest(commerceCheckoutRequest).Execute()

Charge payment via store

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
	commerceCheckoutRequest := *openapiclient.NewCommerceCheckoutRequest() // CommerceCheckoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommerceStoreCharge(context.Background(), storeid).CommerceCheckoutRequest(commerceCheckoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceStoreCharge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceStoreCharge`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceStoreCharge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceStoreChargeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceCheckoutRequest** | [**CommerceCheckoutRequest**](CommerceCheckoutRequest.md) |  | 

### Return type

[**CommerceOrder**](CommerceOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceUpdateStoreListing

> CommerceListing CommerceUpdateStoreListing(ctx, storeid, key).CommerceListing(commerceListing).Execute()

Update store listing

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
	commerceListing := *openapiclient.NewCommerceListing() // CommerceListing | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommerceUpdateStoreListing(context.Background(), storeid, key).CommerceListing(commerceListing).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceUpdateStoreListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceUpdateStoreListing`: CommerceListing
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceUpdateStoreListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceUpdateStoreListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commerceListing** | [**CommerceListing**](CommerceListing.md) |  | 

### Return type

[**CommerceListing**](CommerceListing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

