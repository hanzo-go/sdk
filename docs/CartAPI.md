# \CartAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscardCart**](CartAPI.md#DiscardCart) | **Post** /v1/cart/{id}/discard | Discard a cart the shopper abandoned
[**GetCart**](CartAPI.md#GetCart) | **Get** /v1/cart/{id} | Read one cart with its lines and totals
[**OpenCart**](CartAPI.md#OpenCart) | **Post** /v1/cart | Open a cart for a shopper to fill
[**SetCartItem**](CartAPI.md#SetCartItem) | **Post** /v1/cart/{id}/item | Set one item&#39;s quantity in a cart; zero removes it



## DiscardCart

> Cart DiscardCart(ctx, id).Execute()

Discard a cart the shopper abandoned



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
	id := "id_example" // string | ID is the cart's id, as the open call answered it.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.DiscardCart(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.DiscardCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscardCart`: Cart
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.DiscardCart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the cart&#39;s id, as the open call answered it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscardCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Cart**](Cart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCart

> Cart GetCart(ctx, id).Execute()

Read one cart with its lines and totals



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
	id := "id_example" // string | ID is the cart's id, as the open call answered it.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.GetCart(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.GetCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCart`: Cart
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.GetCart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the cart&#39;s id, as the open call answered it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Cart**](Cart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenCart

> Cart OpenCart(ctx).CartOpen(cartOpen).Execute()

Open a cart for a shopper to fill



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
	cartOpen := *openapiclient.NewCartOpen() // CartOpen | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.OpenCart(context.Background()).CartOpen(cartOpen).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.OpenCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenCart`: Cart
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.OpenCart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cartOpen** | [**CartOpen**](CartOpen.md) |  | 

### Return type

[**Cart**](Cart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCartItem

> Cart SetCartItem(ctx, id).CartItemSet(cartItemSet).Execute()

Set one item's quantity in a cart; zero removes it



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
	id := "id_example" // string | ID is the cart to amend, from the path.
	cartItemSet := *openapiclient.NewCartItemSet() // CartItemSet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.SetCartItem(context.Background(), id).CartItemSet(cartItemSet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.SetCartItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetCartItem`: Cart
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.SetCartItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the cart to amend, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCartItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cartItemSet** | [**CartItemSet**](CartItemSet.md) |  | 

### Return type

[**Cart**](Cart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

