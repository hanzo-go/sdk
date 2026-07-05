# \CommerceProductsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommerceCreateProduct**](CommerceProductsAPI.md#CommerceCreateProduct) | **Post** /v1/commerce/product | Create product
[**CommerceDeleteProduct**](CommerceProductsAPI.md#CommerceDeleteProduct) | **Delete** /v1/commerce/product/{productid} | Delete product
[**CommerceGetProduct**](CommerceProductsAPI.md#CommerceGetProduct) | **Get** /v1/commerce/product/{productid} | Get product
[**CommerceListProducts**](CommerceProductsAPI.md#CommerceListProducts) | **Get** /v1/commerce/product | List products
[**CommercePatchProduct**](CommerceProductsAPI.md#CommercePatchProduct) | **Patch** /v1/commerce/product/{productid} | Partially update product
[**CommerceUpdateProduct**](CommerceProductsAPI.md#CommerceUpdateProduct) | **Put** /v1/commerce/product/{productid} | Update product



## CommerceCreateProduct

> CommerceProduct CommerceCreateProduct(ctx).CommerceProduct(commerceProduct).Execute()

Create product

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
	commerceProduct := *openapiclient.NewCommerceProduct() // CommerceProduct | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceProductsAPI.CommerceCreateProduct(context.Background()).CommerceProduct(commerceProduct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceProductsAPI.CommerceCreateProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateProduct`: CommerceProduct
	fmt.Fprintf(os.Stdout, "Response from `CommerceProductsAPI.CommerceCreateProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceProduct** | [**CommerceProduct**](CommerceProduct.md) |  | 

### Return type

[**CommerceProduct**](CommerceProduct.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceDeleteProduct

> CommerceDeleteProduct(ctx, productid).Execute()

Delete product

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
	productid := "productid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceProductsAPI.CommerceDeleteProduct(context.Background(), productid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceProductsAPI.CommerceDeleteProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceDeleteProductRequest struct via the builder pattern


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


## CommerceGetProduct

> CommerceProduct CommerceGetProduct(ctx, productid).Execute()

Get product

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
	productid := "productid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceProductsAPI.CommerceGetProduct(context.Background(), productid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceProductsAPI.CommerceGetProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetProduct`: CommerceProduct
	fmt.Fprintf(os.Stdout, "Response from `CommerceProductsAPI.CommerceGetProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetProductRequest struct via the builder pattern


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


## CommerceListProducts

> CommercePaginatedProducts CommerceListProducts(ctx).Page(page).Display(display).Sort(sort).Q(q).Execute()

List products

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
	sort := "sort_example" // string | Sort field (prefix with - for descending) (optional) (default to "-UpdatedAt")
	q := "q_example" // string | Search query (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceProductsAPI.CommerceListProducts(context.Background()).Page(page).Display(display).Sort(sort).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceProductsAPI.CommerceListProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceListProducts`: CommercePaginatedProducts
	fmt.Fprintf(os.Stdout, "Response from `CommerceProductsAPI.CommerceListProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceListProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (1-indexed) | [default to 1]
 **display** | **int32** | Number of items per page | [default to 20]
 **sort** | **string** | Sort field (prefix with - for descending) | [default to &quot;-UpdatedAt&quot;]
 **q** | **string** | Search query | 

### Return type

[**CommercePaginatedProducts**](CommercePaginatedProducts.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommercePatchProduct

> CommerceProduct CommercePatchProduct(ctx, productid).CommerceProduct(commerceProduct).Execute()

Partially update product

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
	productid := "productid_example" // string | 
	commerceProduct := *openapiclient.NewCommerceProduct() // CommerceProduct | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceProductsAPI.CommercePatchProduct(context.Background(), productid).CommerceProduct(commerceProduct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceProductsAPI.CommercePatchProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommercePatchProduct`: CommerceProduct
	fmt.Fprintf(os.Stdout, "Response from `CommerceProductsAPI.CommercePatchProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommercePatchProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceProduct** | [**CommerceProduct**](CommerceProduct.md) |  | 

### Return type

[**CommerceProduct**](CommerceProduct.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceUpdateProduct

> CommerceProduct CommerceUpdateProduct(ctx, productid).CommerceProduct(commerceProduct).Execute()

Update product

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
	productid := "productid_example" // string | 
	commerceProduct := *openapiclient.NewCommerceProduct() // CommerceProduct | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceProductsAPI.CommerceUpdateProduct(context.Background(), productid).CommerceProduct(commerceProduct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceProductsAPI.CommerceUpdateProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceUpdateProduct`: CommerceProduct
	fmt.Fprintf(os.Stdout, "Response from `CommerceProductsAPI.CommerceUpdateProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceUpdateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceProduct** | [**CommerceProduct**](CommerceProduct.md) |  | 

### Return type

[**CommerceProduct**](CommerceProduct.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

