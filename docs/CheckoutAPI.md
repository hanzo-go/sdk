# \CheckoutAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommerceAuthorizePayment**](CheckoutAPI.md#CommerceAuthorizePayment) | **Post** /v1/commerce/checkout/authorize | Authorize new payment
[**CommerceAuthorizePaymentForOrder**](CheckoutAPI.md#CommerceAuthorizePaymentForOrder) | **Post** /v1/commerce/checkout/authorize/{orderid} | Authorize payment for existing order
[**CommerceCancelOrder**](CheckoutAPI.md#CommerceCancelOrder) | **Post** /v1/commerce/checkout/cancel/{orderid} | Cancel order
[**CommerceCapturePayment**](CheckoutAPI.md#CommerceCapturePayment) | **Post** /v1/commerce/checkout/capture/{orderid} | Capture authorized payment
[**CommerceChargePayment**](CheckoutAPI.md#CommerceChargePayment) | **Post** /v1/commerce/checkout/charge | Authorize and capture in one step
[**CommerceConfirmOrder**](CheckoutAPI.md#CommerceConfirmOrder) | **Post** /v1/commerce/checkout/confirm/{orderid} | Confirm order
[**CommerceLookupEthereumProxy**](CheckoutAPI.md#CommerceLookupEthereumProxy) | **Get** /v1/commerce/checkout/ethereum/lookup/{proxyaddress} | Lookup Ethereum proxy address



## CommerceAuthorizePayment

> CommerceOrder CommerceAuthorizePayment(ctx).CommerceCheckoutRequest(commerceCheckoutRequest).Execute()

Authorize new payment

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
	commerceCheckoutRequest := *openapiclient.NewCommerceCheckoutRequest() // CommerceCheckoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutAPI.CommerceAuthorizePayment(context.Background()).CommerceCheckoutRequest(commerceCheckoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutAPI.CommerceAuthorizePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceAuthorizePayment`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `CheckoutAPI.CommerceAuthorizePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceAuthorizePaymentRequest struct via the builder pattern


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


## CommerceAuthorizePaymentForOrder

> CommerceOrder CommerceAuthorizePaymentForOrder(ctx, orderid).CommercePaymentRequest(commercePaymentRequest).Execute()

Authorize payment for existing order

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
	orderid := "orderid_example" // string | 
	commercePaymentRequest := *openapiclient.NewCommercePaymentRequest() // CommercePaymentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutAPI.CommerceAuthorizePaymentForOrder(context.Background(), orderid).CommercePaymentRequest(commercePaymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutAPI.CommerceAuthorizePaymentForOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceAuthorizePaymentForOrder`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `CheckoutAPI.CommerceAuthorizePaymentForOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceAuthorizePaymentForOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commercePaymentRequest** | [**CommercePaymentRequest**](CommercePaymentRequest.md) |  | 

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


## CommerceCancelOrder

> CommerceOrder CommerceCancelOrder(ctx, orderid).Execute()

Cancel order

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
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutAPI.CommerceCancelOrder(context.Background(), orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutAPI.CommerceCancelOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCancelOrder`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `CheckoutAPI.CommerceCancelOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCancelOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceOrder**](CommerceOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceCapturePayment

> CommerceOrder CommerceCapturePayment(ctx, orderid).Execute()

Capture authorized payment

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
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutAPI.CommerceCapturePayment(context.Background(), orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutAPI.CommerceCapturePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCapturePayment`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `CheckoutAPI.CommerceCapturePayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCapturePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceOrder**](CommerceOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceChargePayment

> CommerceOrder CommerceChargePayment(ctx).CommerceCheckoutRequest(commerceCheckoutRequest).Execute()

Authorize and capture in one step

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
	commerceCheckoutRequest := *openapiclient.NewCommerceCheckoutRequest() // CommerceCheckoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutAPI.CommerceChargePayment(context.Background()).CommerceCheckoutRequest(commerceCheckoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutAPI.CommerceChargePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceChargePayment`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `CheckoutAPI.CommerceChargePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceChargePaymentRequest struct via the builder pattern


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


## CommerceConfirmOrder

> CommerceOrder CommerceConfirmOrder(ctx, orderid).Execute()

Confirm order

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
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutAPI.CommerceConfirmOrder(context.Background(), orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutAPI.CommerceConfirmOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceConfirmOrder`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `CheckoutAPI.CommerceConfirmOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceConfirmOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceOrder**](CommerceOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceLookupEthereumProxy

> CommerceLookupEthereumProxy200Response CommerceLookupEthereumProxy(ctx, proxyaddress).Execute()

Lookup Ethereum proxy address

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
	proxyaddress := "proxyaddress_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutAPI.CommerceLookupEthereumProxy(context.Background(), proxyaddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutAPI.CommerceLookupEthereumProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceLookupEthereumProxy`: CommerceLookupEthereumProxy200Response
	fmt.Fprintf(os.Stdout, "Response from `CheckoutAPI.CommerceLookupEthereumProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proxyaddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceLookupEthereumProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceLookupEthereumProxy200Response**](CommerceLookupEthereumProxy200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

