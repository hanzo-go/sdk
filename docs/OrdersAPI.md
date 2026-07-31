# \OrdersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommerceAuthorizeOrder**](OrdersAPI.md#CommerceAuthorizeOrder) | **Post** /v1/commerce/order/{orderid}/authorize | Authorize payment for order
[**CommerceCaptureOrder**](OrdersAPI.md#CommerceCaptureOrder) | **Post** /v1/commerce/order/{orderid}/capture | Capture authorized payment
[**CommerceChargeOrder**](OrdersAPI.md#CommerceChargeOrder) | **Post** /v1/commerce/order/{orderid}/charge | Authorize and capture payment (single step)
[**CommerceCreateOrder**](OrdersAPI.md#CommerceCreateOrder) | **Post** /v1/commerce/order | Create order
[**CommerceDeleteOrder**](OrdersAPI.md#CommerceDeleteOrder) | **Delete** /v1/commerce/order/{orderid} | Delete order
[**CommerceGetOrder**](OrdersAPI.md#CommerceGetOrder) | **Get** /v1/commerce/order/{orderid} | Get order
[**CommerceGetOrderPayments**](OrdersAPI.md#CommerceGetOrderPayments) | **Get** /v1/commerce/order/{orderid}/payments | Get order payments
[**CommerceGetOrderReturns**](OrdersAPI.md#CommerceGetOrderReturns) | **Get** /v1/commerce/order/{orderid}/returns | Get order returns
[**CommerceGetOrderStatus**](OrdersAPI.md#CommerceGetOrderStatus) | **Get** /v1/commerce/order/{orderid}/status | Get order status
[**CommerceListOrders**](OrdersAPI.md#CommerceListOrders) | **Get** /v1/commerce/order | List orders
[**CommercePatchOrder**](OrdersAPI.md#CommercePatchOrder) | **Patch** /v1/commerce/order/{orderid} | Partially update order
[**CommerceRefundOrder**](OrdersAPI.md#CommerceRefundOrder) | **Post** /v1/commerce/order/{orderid}/refund | Refund order
[**CommerceSendFulfillmentConfirmation**](OrdersAPI.md#CommerceSendFulfillmentConfirmation) | **Get** /v1/commerce/order/{orderid}/sendfulfillmentconfirmation | Send fulfillment confirmation email
[**CommerceSendOrderConfirmation**](OrdersAPI.md#CommerceSendOrderConfirmation) | **Get** /v1/commerce/order/{orderid}/sendorderconfirmation | Send order confirmation email
[**CommerceSendRefundConfirmation**](OrdersAPI.md#CommerceSendRefundConfirmation) | **Get** /v1/commerce/order/{orderid}/sendrefundconfirmation | Send refund confirmation email
[**CommerceUpdateOrder**](OrdersAPI.md#CommerceUpdateOrder) | **Put** /v1/commerce/order/{orderid} | Update order



## CommerceAuthorizeOrder

> CommerceOrder CommerceAuthorizeOrder(ctx, orderid).CommercePaymentRequest(commercePaymentRequest).Execute()

Authorize payment for order

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
	resp, r, err := apiClient.OrdersAPI.CommerceAuthorizeOrder(context.Background(), orderid).CommercePaymentRequest(commercePaymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CommerceAuthorizeOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceAuthorizeOrder`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CommerceAuthorizeOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceAuthorizeOrderRequest struct via the builder pattern


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


## CommerceCaptureOrder

> CommerceOrder CommerceCaptureOrder(ctx, orderid).Execute()

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
	resp, r, err := apiClient.OrdersAPI.CommerceCaptureOrder(context.Background(), orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CommerceCaptureOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCaptureOrder`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CommerceCaptureOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCaptureOrderRequest struct via the builder pattern


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


## CommerceChargeOrder

> CommerceOrder CommerceChargeOrder(ctx, orderid).CommercePaymentRequest(commercePaymentRequest).Execute()

Authorize and capture payment (single step)

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
	resp, r, err := apiClient.OrdersAPI.CommerceChargeOrder(context.Background(), orderid).CommercePaymentRequest(commercePaymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CommerceChargeOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceChargeOrder`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CommerceChargeOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceChargeOrderRequest struct via the builder pattern


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


## CommerceCreateOrder

> CommerceOrder CommerceCreateOrder(ctx).CommerceOrder(commerceOrder).Execute()

Create order

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
	commerceOrder := *openapiclient.NewCommerceOrder() // CommerceOrder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.CommerceCreateOrder(context.Background()).CommerceOrder(commerceOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CommerceCreateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateOrder`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CommerceCreateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceOrder** | [**CommerceOrder**](CommerceOrder.md) |  | 

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


## CommerceDeleteOrder

> CommerceDeleteOrder(ctx, orderid).Execute()

Delete order

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
	r, err := apiClient.OrdersAPI.CommerceDeleteOrder(context.Background(), orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CommerceDeleteOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceDeleteOrderRequest struct via the builder pattern


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


## CommerceGetOrder

> CommerceOrder CommerceGetOrder(ctx, orderid).Execute()

Get order

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
	resp, r, err := apiClient.OrdersAPI.CommerceGetOrder(context.Background(), orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CommerceGetOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetOrder`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CommerceGetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetOrderRequest struct via the builder pattern


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


## CommerceGetOrderPayments

> []CommercePayment CommerceGetOrderPayments(ctx, orderid).Execute()

Get order payments

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
	resp, r, err := apiClient.OrdersAPI.CommerceGetOrderPayments(context.Background(), orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CommerceGetOrderPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetOrderPayments`: []CommercePayment
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CommerceGetOrderPayments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetOrderPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CommercePayment**](CommercePayment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetOrderReturns

> []CommerceReturn CommerceGetOrderReturns(ctx, orderid).Execute()

Get order returns

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
	resp, r, err := apiClient.OrdersAPI.CommerceGetOrderReturns(context.Background(), orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CommerceGetOrderReturns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetOrderReturns`: []CommerceReturn
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CommerceGetOrderReturns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetOrderReturnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CommerceReturn**](CommerceReturn.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetOrderStatus

> CommerceGetOrderStatus200Response CommerceGetOrderStatus(ctx, orderid).Execute()

Get order status

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
	resp, r, err := apiClient.OrdersAPI.CommerceGetOrderStatus(context.Background(), orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CommerceGetOrderStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetOrderStatus`: CommerceGetOrderStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CommerceGetOrderStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetOrderStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceGetOrderStatus200Response**](CommerceGetOrderStatus200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceListOrders

> CommercePaginatedOrders CommerceListOrders(ctx).Page(page).Display(display).Sort(sort).Q(q).Execute()

List orders

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
	resp, r, err := apiClient.OrdersAPI.CommerceListOrders(context.Background()).Page(page).Display(display).Sort(sort).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CommerceListOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceListOrders`: CommercePaginatedOrders
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CommerceListOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceListOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (1-indexed) | [default to 1]
 **display** | **int32** | Number of items per page | [default to 20]
 **sort** | **string** | Sort field (prefix with - for descending) | [default to &quot;-UpdatedAt&quot;]
 **q** | **string** | Search query | 

### Return type

[**CommercePaginatedOrders**](CommercePaginatedOrders.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommercePatchOrder

> CommerceOrder CommercePatchOrder(ctx, orderid).CommerceOrder(commerceOrder).Execute()

Partially update order

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
	commerceOrder := *openapiclient.NewCommerceOrder() // CommerceOrder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.CommercePatchOrder(context.Background(), orderid).CommerceOrder(commerceOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CommercePatchOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommercePatchOrder`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CommercePatchOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommercePatchOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceOrder** | [**CommerceOrder**](CommerceOrder.md) |  | 

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


## CommerceRefundOrder

> CommerceOrder CommerceRefundOrder(ctx, orderid).CommerceRefundOrderRequest(commerceRefundOrderRequest).Execute()

Refund order

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
	commerceRefundOrderRequest := *openapiclient.NewCommerceRefundOrderRequest() // CommerceRefundOrderRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.CommerceRefundOrder(context.Background(), orderid).CommerceRefundOrderRequest(commerceRefundOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CommerceRefundOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceRefundOrder`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CommerceRefundOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceRefundOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceRefundOrderRequest** | [**CommerceRefundOrderRequest**](CommerceRefundOrderRequest.md) |  | 

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


## CommerceSendFulfillmentConfirmation

> map[string]interface{} CommerceSendFulfillmentConfirmation(ctx, orderid).Execute()

Send fulfillment confirmation email

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
	resp, r, err := apiClient.OrdersAPI.CommerceSendFulfillmentConfirmation(context.Background(), orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CommerceSendFulfillmentConfirmation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceSendFulfillmentConfirmation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CommerceSendFulfillmentConfirmation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceSendFulfillmentConfirmationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceSendOrderConfirmation

> map[string]interface{} CommerceSendOrderConfirmation(ctx, orderid).Execute()

Send order confirmation email

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
	resp, r, err := apiClient.OrdersAPI.CommerceSendOrderConfirmation(context.Background(), orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CommerceSendOrderConfirmation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceSendOrderConfirmation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CommerceSendOrderConfirmation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceSendOrderConfirmationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceSendRefundConfirmation

> map[string]interface{} CommerceSendRefundConfirmation(ctx, orderid).Execute()

Send refund confirmation email

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
	resp, r, err := apiClient.OrdersAPI.CommerceSendRefundConfirmation(context.Background(), orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CommerceSendRefundConfirmation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceSendRefundConfirmation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CommerceSendRefundConfirmation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceSendRefundConfirmationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceUpdateOrder

> CommerceOrder CommerceUpdateOrder(ctx, orderid).CommerceOrder(commerceOrder).Execute()

Update order

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
	commerceOrder := *openapiclient.NewCommerceOrder() // CommerceOrder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.CommerceUpdateOrder(context.Background(), orderid).CommerceOrder(commerceOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CommerceUpdateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceUpdateOrder`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CommerceUpdateOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceUpdateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceOrder** | [**CommerceOrder**](CommerceOrder.md) |  | 

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

