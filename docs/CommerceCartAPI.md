# \CommerceCartAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommerceCreateCart**](CommerceCartAPI.md#CommerceCreateCart) | **Post** /v1/commerce/cart | Create cart
[**CommerceDeleteCart**](CommerceCartAPI.md#CommerceDeleteCart) | **Delete** /v1/commerce/cart/{cartid} | Delete cart
[**CommerceDiscardCart**](CommerceCartAPI.md#CommerceDiscardCart) | **Post** /v1/commerce/cart/{cartid}/discard | Discard cart
[**CommerceGetCart**](CommerceCartAPI.md#CommerceGetCart) | **Get** /v1/commerce/cart/{cartid} | Get cart
[**CommerceListCarts**](CommerceCartAPI.md#CommerceListCarts) | **Get** /v1/commerce/cart | List carts
[**CommercePatchCart**](CommerceCartAPI.md#CommercePatchCart) | **Patch** /v1/commerce/cart/{cartid} | Partially update cart
[**CommerceSetCartItem**](CommerceCartAPI.md#CommerceSetCartItem) | **Post** /v1/commerce/cart/{cartid}/set | Set item in cart
[**CommerceUpdateCart**](CommerceCartAPI.md#CommerceUpdateCart) | **Put** /v1/commerce/cart/{cartid} | Update cart



## CommerceCreateCart

> CommerceCart CommerceCreateCart(ctx).CommerceCart(commerceCart).Execute()

Create cart

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
	commerceCart := *openapiclient.NewCommerceCart() // CommerceCart | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceCartAPI.CommerceCreateCart(context.Background()).CommerceCart(commerceCart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceCartAPI.CommerceCreateCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateCart`: CommerceCart
	fmt.Fprintf(os.Stdout, "Response from `CommerceCartAPI.CommerceCreateCart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceCart** | [**CommerceCart**](CommerceCart.md) |  | 

### Return type

[**CommerceCart**](CommerceCart.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceDeleteCart

> CommerceDeleteCart(ctx, cartid).Execute()

Delete cart

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
	cartid := "cartid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceCartAPI.CommerceDeleteCart(context.Background(), cartid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceCartAPI.CommerceDeleteCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceDeleteCartRequest struct via the builder pattern


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


## CommerceDiscardCart

> CommerceCart CommerceDiscardCart(ctx, cartid).Execute()

Discard cart

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
	cartid := "cartid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceCartAPI.CommerceDiscardCart(context.Background(), cartid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceCartAPI.CommerceDiscardCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceDiscardCart`: CommerceCart
	fmt.Fprintf(os.Stdout, "Response from `CommerceCartAPI.CommerceDiscardCart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceDiscardCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceCart**](CommerceCart.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetCart

> CommerceCart CommerceGetCart(ctx, cartid).Execute()

Get cart

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
	cartid := "cartid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceCartAPI.CommerceGetCart(context.Background(), cartid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceCartAPI.CommerceGetCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetCart`: CommerceCart
	fmt.Fprintf(os.Stdout, "Response from `CommerceCartAPI.CommerceGetCart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceCart**](CommerceCart.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceListCarts

> CommercePaginatedCarts CommerceListCarts(ctx).Page(page).Display(display).Execute()

List carts

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
	resp, r, err := apiClient.CommerceCartAPI.CommerceListCarts(context.Background()).Page(page).Display(display).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceCartAPI.CommerceListCarts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceListCarts`: CommercePaginatedCarts
	fmt.Fprintf(os.Stdout, "Response from `CommerceCartAPI.CommerceListCarts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceListCartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (1-indexed) | [default to 1]
 **display** | **int32** | Number of items per page | [default to 20]

### Return type

[**CommercePaginatedCarts**](CommercePaginatedCarts.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommercePatchCart

> CommerceCart CommercePatchCart(ctx, cartid).CommerceCart(commerceCart).Execute()

Partially update cart

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
	cartid := "cartid_example" // string | 
	commerceCart := *openapiclient.NewCommerceCart() // CommerceCart | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceCartAPI.CommercePatchCart(context.Background(), cartid).CommerceCart(commerceCart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceCartAPI.CommercePatchCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommercePatchCart`: CommerceCart
	fmt.Fprintf(os.Stdout, "Response from `CommerceCartAPI.CommercePatchCart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommercePatchCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceCart** | [**CommerceCart**](CommerceCart.md) |  | 

### Return type

[**CommerceCart**](CommerceCart.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceSetCartItem

> CommerceCart CommerceSetCartItem(ctx, cartid).CommerceSetCartItemRequest(commerceSetCartItemRequest).Execute()

Set item in cart

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
	cartid := "cartid_example" // string | 
	commerceSetCartItemRequest := *openapiclient.NewCommerceSetCartItemRequest() // CommerceSetCartItemRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceCartAPI.CommerceSetCartItem(context.Background(), cartid).CommerceSetCartItemRequest(commerceSetCartItemRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceCartAPI.CommerceSetCartItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceSetCartItem`: CommerceCart
	fmt.Fprintf(os.Stdout, "Response from `CommerceCartAPI.CommerceSetCartItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceSetCartItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceSetCartItemRequest** | [**CommerceSetCartItemRequest**](CommerceSetCartItemRequest.md) |  | 

### Return type

[**CommerceCart**](CommerceCart.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceUpdateCart

> CommerceCart CommerceUpdateCart(ctx, cartid).CommerceCart(commerceCart).Execute()

Update cart

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
	cartid := "cartid_example" // string | 
	commerceCart := *openapiclient.NewCommerceCart() // CommerceCart | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceCartAPI.CommerceUpdateCart(context.Background(), cartid).CommerceCart(commerceCart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceCartAPI.CommerceUpdateCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceUpdateCart`: CommerceCart
	fmt.Fprintf(os.Stdout, "Response from `CommerceCartAPI.CommerceUpdateCart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceUpdateCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceCart** | [**CommerceCart**](CommerceCart.md) |  | 

### Return type

[**CommerceCart**](CommerceCart.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

