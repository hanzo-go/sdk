# \IamPaymentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamApiControllerAddOrder**](IamPaymentsAPI.md#IamApiControllerAddOrder) | **Post** /v1/iam/orders | Api Controller Add Order
[**IamApiControllerAddPayment**](IamPaymentsAPI.md#IamApiControllerAddPayment) | **Post** /v1/iam/payments | Api Controller Add Payment
[**IamApiControllerAddPlan**](IamPaymentsAPI.md#IamApiControllerAddPlan) | **Post** /v1/iam/plans | Api Controller Add Plan
[**IamApiControllerAddPricing**](IamPaymentsAPI.md#IamApiControllerAddPricing) | **Post** /v1/iam/pricings | Api Controller Add Pricing
[**IamApiControllerAddProduct**](IamPaymentsAPI.md#IamApiControllerAddProduct) | **Post** /v1/iam/products | Api Controller Add Product
[**IamApiControllerAddSubscription**](IamPaymentsAPI.md#IamApiControllerAddSubscription) | **Post** /v1/iam/subscriptions | Api Controller Add Subscription
[**IamApiControllerAddTransaction**](IamPaymentsAPI.md#IamApiControllerAddTransaction) | **Post** /v1/iam/transactions | Api Controller Add Transaction
[**IamApiControllerCancelOrder**](IamPaymentsAPI.md#IamApiControllerCancelOrder) | **Post** /v1/iam/orders/cancel | Api Controller Cancel Order
[**IamApiControllerDeleteOrder**](IamPaymentsAPI.md#IamApiControllerDeleteOrder) | **Delete** /v1/iam/orders/{id} | Api Controller Delete Order
[**IamApiControllerDeletePayment**](IamPaymentsAPI.md#IamApiControllerDeletePayment) | **Delete** /v1/iam/payments/{id} | Api Controller Delete Payment
[**IamApiControllerDeletePlan**](IamPaymentsAPI.md#IamApiControllerDeletePlan) | **Delete** /v1/iam/plans/{id} | Api Controller Delete Plan
[**IamApiControllerDeletePricing**](IamPaymentsAPI.md#IamApiControllerDeletePricing) | **Delete** /v1/iam/pricings/{id} | Api Controller Delete Pricing
[**IamApiControllerDeleteProduct**](IamPaymentsAPI.md#IamApiControllerDeleteProduct) | **Delete** /v1/iam/products/{id} | Api Controller Delete Product
[**IamApiControllerDeleteSubscription**](IamPaymentsAPI.md#IamApiControllerDeleteSubscription) | **Delete** /v1/iam/subscriptions/{id} | Api Controller Delete Subscription
[**IamApiControllerDeleteTransaction**](IamPaymentsAPI.md#IamApiControllerDeleteTransaction) | **Delete** /v1/iam/transactions/{id} | Api Controller Delete Transaction
[**IamApiControllerGetOrder**](IamPaymentsAPI.md#IamApiControllerGetOrder) | **Get** /v1/iam/orders/{id} | Api Controller Get Order
[**IamApiControllerGetOrders**](IamPaymentsAPI.md#IamApiControllerGetOrders) | **Get** /v1/iam/orders | Api Controller Get Orders
[**IamApiControllerGetPlan**](IamPaymentsAPI.md#IamApiControllerGetPlan) | **Get** /v1/iam/plans/{id} | Api Controller Get Plan
[**IamApiControllerGetPlans**](IamPaymentsAPI.md#IamApiControllerGetPlans) | **Get** /v1/iam/plans | Api Controller Get Plans
[**IamApiControllerGetPricing**](IamPaymentsAPI.md#IamApiControllerGetPricing) | **Get** /v1/iam/pricings/{id} | Api Controller Get Pricing
[**IamApiControllerGetPricings**](IamPaymentsAPI.md#IamApiControllerGetPricings) | **Get** /v1/iam/pricings | Api Controller Get Pricings
[**IamApiControllerGetProduct**](IamPaymentsAPI.md#IamApiControllerGetProduct) | **Get** /v1/iam/products/{id} | Api Controller Get Product
[**IamApiControllerGetProducts**](IamPaymentsAPI.md#IamApiControllerGetProducts) | **Get** /v1/iam/products | Api Controller Get Products
[**IamApiControllerGetSubscription**](IamPaymentsAPI.md#IamApiControllerGetSubscription) | **Get** /v1/iam/subscriptions/{id} | Api Controller Get Subscription
[**IamApiControllerGetSubscriptions**](IamPaymentsAPI.md#IamApiControllerGetSubscriptions) | **Get** /v1/iam/subscriptions | Api Controller Get Subscriptions
[**IamApiControllerGetTransaction**](IamPaymentsAPI.md#IamApiControllerGetTransaction) | **Get** /v1/iam/transactions/{id} | Api Controller Get Transaction
[**IamApiControllerGetTransactions**](IamPaymentsAPI.md#IamApiControllerGetTransactions) | **Get** /v1/iam/transactions | Api Controller Get Transactions
[**IamApiControllerGetUserOrders**](IamPaymentsAPI.md#IamApiControllerGetUserOrders) | **Get** /v1/iam/user-orders | Api Controller Get User Orders
[**IamApiControllerInvoicePayment**](IamPaymentsAPI.md#IamApiControllerInvoicePayment) | **Post** /v1/iam/invoice-payment | Api Controller Invoice Payment
[**IamApiControllerNotifyPayment**](IamPaymentsAPI.md#IamApiControllerNotifyPayment) | **Post** /v1/iam/payments/notify | Api Controller Notify Payment
[**IamApiControllerPayOrder**](IamPaymentsAPI.md#IamApiControllerPayOrder) | **Post** /v1/iam/pay-order | Api Controller Pay Order
[**IamApiControllerPlaceOrder**](IamPaymentsAPI.md#IamApiControllerPlaceOrder) | **Post** /v1/iam/place-order | Api Controller Place Order
[**IamApiControllerUpdateOrder**](IamPaymentsAPI.md#IamApiControllerUpdateOrder) | **Put** /v1/iam/orders/{id} | Api Controller Update Order
[**IamApiControllerUpdatePayment**](IamPaymentsAPI.md#IamApiControllerUpdatePayment) | **Put** /v1/iam/payments/{id} | Api Controller Update Payment
[**IamApiControllerUpdatePlan**](IamPaymentsAPI.md#IamApiControllerUpdatePlan) | **Put** /v1/iam/plans/{id} | Api Controller Update Plan
[**IamApiControllerUpdatePricing**](IamPaymentsAPI.md#IamApiControllerUpdatePricing) | **Put** /v1/iam/pricings/{id} | Api Controller Update Pricing
[**IamApiControllerUpdateProduct**](IamPaymentsAPI.md#IamApiControllerUpdateProduct) | **Put** /v1/iam/products/{id} | Api Controller Update Product
[**IamApiControllerUpdateSubscription**](IamPaymentsAPI.md#IamApiControllerUpdateSubscription) | **Put** /v1/iam/subscriptions/{id} | Api Controller Update Subscription
[**IamApiControllerUpdateTransaction**](IamPaymentsAPI.md#IamApiControllerUpdateTransaction) | **Put** /v1/iam/transactions/{id} | Api Controller Update Transaction



## IamApiControllerAddOrder

> IamControllersResponse IamApiControllerAddOrder(ctx).IamObjectOrder(iamObjectOrder).Execute()

Api Controller Add Order



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
	iamObjectOrder := *openapiclient.NewIamObjectOrder() // IamObjectOrder | The details of the order

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerAddOrder(context.Background()).IamObjectOrder(iamObjectOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerAddOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddOrder`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerAddOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectOrder** | [**IamObjectOrder**](IamObjectOrder.md) | The details of the order | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddPayment

> IamControllersResponse IamApiControllerAddPayment(ctx).IamObjectPayment(iamObjectPayment).Execute()

Api Controller Add Payment



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
	iamObjectPayment := *openapiclient.NewIamObjectPayment() // IamObjectPayment | The details of the payment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerAddPayment(context.Background()).IamObjectPayment(iamObjectPayment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerAddPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddPayment`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerAddPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectPayment** | [**IamObjectPayment**](IamObjectPayment.md) | The details of the payment | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddPlan

> IamControllersResponse IamApiControllerAddPlan(ctx).IamObjectPlan(iamObjectPlan).Execute()

Api Controller Add Plan



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
	iamObjectPlan := *openapiclient.NewIamObjectPlan() // IamObjectPlan | The details of the plan

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerAddPlan(context.Background()).IamObjectPlan(iamObjectPlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerAddPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddPlan`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerAddPlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectPlan** | [**IamObjectPlan**](IamObjectPlan.md) | The details of the plan | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddPricing

> IamControllersResponse IamApiControllerAddPricing(ctx).IamObjectPricing(iamObjectPricing).Execute()

Api Controller Add Pricing



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
	iamObjectPricing := *openapiclient.NewIamObjectPricing() // IamObjectPricing | The details of the pricing

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerAddPricing(context.Background()).IamObjectPricing(iamObjectPricing).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerAddPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddPricing`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerAddPricing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddPricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectPricing** | [**IamObjectPricing**](IamObjectPricing.md) | The details of the pricing | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddProduct

> IamControllersResponse IamApiControllerAddProduct(ctx).IamObjectProduct(iamObjectProduct).Execute()

Api Controller Add Product



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
	iamObjectProduct := *openapiclient.NewIamObjectProduct() // IamObjectProduct | The details of the product

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerAddProduct(context.Background()).IamObjectProduct(iamObjectProduct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerAddProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddProduct`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerAddProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectProduct** | [**IamObjectProduct**](IamObjectProduct.md) | The details of the product | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddSubscription

> IamControllersResponse IamApiControllerAddSubscription(ctx).IamObjectSubscription(iamObjectSubscription).Execute()

Api Controller Add Subscription



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
	iamObjectSubscription := *openapiclient.NewIamObjectSubscription() // IamObjectSubscription | The details of the subscription

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerAddSubscription(context.Background()).IamObjectSubscription(iamObjectSubscription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerAddSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddSubscription`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerAddSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectSubscription** | [**IamObjectSubscription**](IamObjectSubscription.md) | The details of the subscription | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddTransaction

> IamControllersResponse IamApiControllerAddTransaction(ctx).IamObjectTransaction(iamObjectTransaction).DryRun(dryRun).Execute()

Api Controller Add Transaction



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
	iamObjectTransaction := *openapiclient.NewIamObjectTransaction() // IamObjectTransaction | The details of the transaction
	dryRun := "dryRun_example" // string | Dry run mode: set to 'true' or '1' to validate without committing (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerAddTransaction(context.Background()).IamObjectTransaction(iamObjectTransaction).DryRun(dryRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerAddTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddTransaction`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerAddTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectTransaction** | [**IamObjectTransaction**](IamObjectTransaction.md) | The details of the transaction | 
 **dryRun** | **string** | Dry run mode: set to &#39;true&#39; or &#39;1&#39; to validate without committing | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerCancelOrder

> IamControllersResponse IamApiControllerCancelOrder(ctx).Id(id).Execute()

Api Controller Cancel Order



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
	id := "id_example" // string | The id ( owner/name ) of the order

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerCancelOrder(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerCancelOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerCancelOrder`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerCancelOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerCancelOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the order | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteOrder

> IamControllersResponse IamApiControllerDeleteOrder(ctx, id).IamObjectOrder(iamObjectOrder).Execute()

Api Controller Delete Order



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectOrder := *openapiclient.NewIamObjectOrder() // IamObjectOrder | The details of the order

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerDeleteOrder(context.Background(), id).IamObjectOrder(iamObjectOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerDeleteOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteOrder`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerDeleteOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectOrder** | [**IamObjectOrder**](IamObjectOrder.md) | The details of the order | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeletePayment

> IamControllersResponse IamApiControllerDeletePayment(ctx, id).IamObjectPayment(iamObjectPayment).Execute()

Api Controller Delete Payment



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectPayment := *openapiclient.NewIamObjectPayment() // IamObjectPayment | The details of the payment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerDeletePayment(context.Background(), id).IamObjectPayment(iamObjectPayment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerDeletePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeletePayment`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerDeletePayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeletePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectPayment** | [**IamObjectPayment**](IamObjectPayment.md) | The details of the payment | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeletePlan

> IamControllersResponse IamApiControllerDeletePlan(ctx, id).IamObjectPlan(iamObjectPlan).Execute()

Api Controller Delete Plan



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectPlan := *openapiclient.NewIamObjectPlan() // IamObjectPlan | The details of the plan

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerDeletePlan(context.Background(), id).IamObjectPlan(iamObjectPlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerDeletePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeletePlan`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerDeletePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeletePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectPlan** | [**IamObjectPlan**](IamObjectPlan.md) | The details of the plan | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeletePricing

> IamControllersResponse IamApiControllerDeletePricing(ctx, id).IamObjectPricing(iamObjectPricing).Execute()

Api Controller Delete Pricing



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectPricing := *openapiclient.NewIamObjectPricing() // IamObjectPricing | The details of the pricing

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerDeletePricing(context.Background(), id).IamObjectPricing(iamObjectPricing).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerDeletePricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeletePricing`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerDeletePricing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeletePricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectPricing** | [**IamObjectPricing**](IamObjectPricing.md) | The details of the pricing | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteProduct

> IamControllersResponse IamApiControllerDeleteProduct(ctx, id).IamObjectProduct(iamObjectProduct).Execute()

Api Controller Delete Product



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectProduct := *openapiclient.NewIamObjectProduct() // IamObjectProduct | The details of the product

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerDeleteProduct(context.Background(), id).IamObjectProduct(iamObjectProduct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerDeleteProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteProduct`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerDeleteProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectProduct** | [**IamObjectProduct**](IamObjectProduct.md) | The details of the product | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteSubscription

> IamControllersResponse IamApiControllerDeleteSubscription(ctx, id).IamObjectSubscription(iamObjectSubscription).Execute()

Api Controller Delete Subscription



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectSubscription := *openapiclient.NewIamObjectSubscription() // IamObjectSubscription | The details of the subscription

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerDeleteSubscription(context.Background(), id).IamObjectSubscription(iamObjectSubscription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerDeleteSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteSubscription`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerDeleteSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectSubscription** | [**IamObjectSubscription**](IamObjectSubscription.md) | The details of the subscription | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteTransaction

> IamControllersResponse IamApiControllerDeleteTransaction(ctx, id).IamObjectTransaction(iamObjectTransaction).Execute()

Api Controller Delete Transaction



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectTransaction := *openapiclient.NewIamObjectTransaction() // IamObjectTransaction | The details of the transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerDeleteTransaction(context.Background(), id).IamObjectTransaction(iamObjectTransaction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerDeleteTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteTransaction`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerDeleteTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectTransaction** | [**IamObjectTransaction**](IamObjectTransaction.md) | The details of the transaction | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetOrder

> IamObjectOrder IamApiControllerGetOrder(ctx, id).Execute()

Api Controller Get Order



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
	id := "id_example" // string | The id ( owner/name ) of the order

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerGetOrder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerGetOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetOrder`: IamObjectOrder
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerGetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the order | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectOrder**](IamObjectOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetOrders

> []IamObjectOrder IamApiControllerGetOrders(ctx).Owner(owner).Execute()

Api Controller Get Orders



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
	owner := "owner_example" // string | The owner of orders

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerGetOrders(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerGetOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetOrders`: []IamObjectOrder
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerGetOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of orders | 

### Return type

[**[]IamObjectOrder**](IamObjectOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetPlan

> IamObjectPlan IamApiControllerGetPlan(ctx, id).IncludeOption(includeOption).Execute()

Api Controller Get Plan



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
	id := "id_example" // string | The id ( owner/name ) of the plan
	includeOption := true // bool | Should include plan's option (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerGetPlan(context.Background(), id).IncludeOption(includeOption).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerGetPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetPlan`: IamObjectPlan
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerGetPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the plan | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeOption** | **bool** | Should include plan&#39;s option | 

### Return type

[**IamObjectPlan**](IamObjectPlan.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetPlans

> []IamObjectPlan IamApiControllerGetPlans(ctx).Owner(owner).Execute()

Api Controller Get Plans



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
	owner := "owner_example" // string | The owner of plans

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerGetPlans(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerGetPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetPlans`: []IamObjectPlan
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerGetPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of plans | 

### Return type

[**[]IamObjectPlan**](IamObjectPlan.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetPricing

> IamObjectPricing IamApiControllerGetPricing(ctx, id).Execute()

Api Controller Get Pricing



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
	id := "id_example" // string | The id ( owner/name ) of the pricing

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerGetPricing(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerGetPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetPricing`: IamObjectPricing
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerGetPricing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the pricing | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetPricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectPricing**](IamObjectPricing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetPricings

> []IamObjectPricing IamApiControllerGetPricings(ctx).Owner(owner).Execute()

Api Controller Get Pricings



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
	owner := "owner_example" // string | The owner of pricings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerGetPricings(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerGetPricings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetPricings`: []IamObjectPricing
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerGetPricings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetPricingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of pricings | 

### Return type

[**[]IamObjectPricing**](IamObjectPricing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetProduct

> IamObjectProduct IamApiControllerGetProduct(ctx, id).Execute()

Api Controller Get Product



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
	id := "id_example" // string | The id ( owner/name ) of the product

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerGetProduct(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerGetProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetProduct`: IamObjectProduct
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerGetProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the product | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectProduct**](IamObjectProduct.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetProducts

> []IamObjectProduct IamApiControllerGetProducts(ctx).Owner(owner).Execute()

Api Controller Get Products



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
	owner := "owner_example" // string | The owner of products

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerGetProducts(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerGetProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetProducts`: []IamObjectProduct
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerGetProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of products | 

### Return type

[**[]IamObjectProduct**](IamObjectProduct.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetSubscription

> IamObjectSubscription IamApiControllerGetSubscription(ctx, id).Execute()

Api Controller Get Subscription



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
	id := "id_example" // string | The id ( owner/name ) of the subscription

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerGetSubscription(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerGetSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetSubscription`: IamObjectSubscription
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerGetSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectSubscription**](IamObjectSubscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetSubscriptions

> []IamObjectSubscription IamApiControllerGetSubscriptions(ctx).Owner(owner).Execute()

Api Controller Get Subscriptions



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
	owner := "owner_example" // string | The owner of subscriptions

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerGetSubscriptions(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerGetSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetSubscriptions`: []IamObjectSubscription
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerGetSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of subscriptions | 

### Return type

[**[]IamObjectSubscription**](IamObjectSubscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetTransaction

> IamObjectTransaction IamApiControllerGetTransaction(ctx, id).Execute()

Api Controller Get Transaction



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
	id := "id_example" // string | The id ( owner/name ) of the transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerGetTransaction(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerGetTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetTransaction`: IamObjectTransaction
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerGetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectTransaction**](IamObjectTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetTransactions

> []IamObjectTransaction IamApiControllerGetTransactions(ctx).Owner(owner).Execute()

Api Controller Get Transactions



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
	owner := "owner_example" // string | The owner of transactions

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerGetTransactions(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerGetTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetTransactions`: []IamObjectTransaction
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerGetTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of transactions | 

### Return type

[**[]IamObjectTransaction**](IamObjectTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetUserOrders

> []IamObjectOrder IamApiControllerGetUserOrders(ctx).Owner(owner).User(user).Execute()

Api Controller Get User Orders



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
	owner := "owner_example" // string | The owner of orders
	user := "user_example" // string | The username of the user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerGetUserOrders(context.Background()).Owner(owner).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerGetUserOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetUserOrders`: []IamObjectOrder
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerGetUserOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetUserOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of orders | 
 **user** | **string** | The username of the user | 

### Return type

[**[]IamObjectOrder**](IamObjectOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerInvoicePayment

> IamControllersResponse IamApiControllerInvoicePayment(ctx).Id(id).Execute()

Api Controller Invoice Payment



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
	id := "id_example" // string | The id ( owner/name ) of the payment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerInvoicePayment(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerInvoicePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerInvoicePayment`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerInvoicePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerInvoicePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the payment | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerNotifyPayment

> IamControllersResponse IamApiControllerNotifyPayment(ctx).IamObjectPayment(iamObjectPayment).Execute()

Api Controller Notify Payment



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
	iamObjectPayment := *openapiclient.NewIamObjectPayment() // IamObjectPayment | The details of the payment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerNotifyPayment(context.Background()).IamObjectPayment(iamObjectPayment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerNotifyPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerNotifyPayment`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerNotifyPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerNotifyPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectPayment** | [**IamObjectPayment**](IamObjectPayment.md) | The details of the payment | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerPayOrder

> IamControllersResponse IamApiControllerPayOrder(ctx).Id(id).ProviderName(providerName).Execute()

Api Controller Pay Order



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
	id := "id_example" // string | The id ( owner/name ) of the order
	providerName := "providerName_example" // string | The name of the provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerPayOrder(context.Background()).Id(id).ProviderName(providerName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerPayOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerPayOrder`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerPayOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerPayOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the order | 
 **providerName** | **string** | The name of the provider | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerPlaceOrder

> IamObjectOrder IamApiControllerPlaceOrder(ctx).ProductId(productId).PricingName(pricingName).PlanName(planName).CustomPrice(customPrice).UserName(userName).Execute()

Api Controller Place Order



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
	productId := "productId_example" // string | The id ( owner/name ) of the product
	pricingName := "pricingName_example" // string | The name of the pricing (for subscription) (optional)
	planName := "planName_example" // string | The name of the plan (for subscription) (optional)
	customPrice := float32(8.14) // float32 | Custom price for recharge products (optional)
	userName := "userName_example" // string | The username to place order for (admin only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerPlaceOrder(context.Background()).ProductId(productId).PricingName(pricingName).PlanName(planName).CustomPrice(customPrice).UserName(userName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerPlaceOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerPlaceOrder`: IamObjectOrder
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerPlaceOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerPlaceOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | The id ( owner/name ) of the product | 
 **pricingName** | **string** | The name of the pricing (for subscription) | 
 **planName** | **string** | The name of the plan (for subscription) | 
 **customPrice** | **float32** | Custom price for recharge products | 
 **userName** | **string** | The username to place order for (admin only) | 

### Return type

[**IamObjectOrder**](IamObjectOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateOrder

> IamControllersResponse IamApiControllerUpdateOrder(ctx, id).IamObjectOrder(iamObjectOrder).Execute()

Api Controller Update Order



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
	id := "id_example" // string | The id ( owner/name ) of the order
	iamObjectOrder := *openapiclient.NewIamObjectOrder() // IamObjectOrder | The details of the order

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerUpdateOrder(context.Background(), id).IamObjectOrder(iamObjectOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerUpdateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateOrder`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerUpdateOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the order | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectOrder** | [**IamObjectOrder**](IamObjectOrder.md) | The details of the order | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdatePayment

> IamControllersResponse IamApiControllerUpdatePayment(ctx, id).IamObjectPayment(iamObjectPayment).Execute()

Api Controller Update Payment



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
	id := "id_example" // string | The id ( owner/name ) of the payment
	iamObjectPayment := *openapiclient.NewIamObjectPayment() // IamObjectPayment | The details of the payment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerUpdatePayment(context.Background(), id).IamObjectPayment(iamObjectPayment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerUpdatePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdatePayment`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerUpdatePayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the payment | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectPayment** | [**IamObjectPayment**](IamObjectPayment.md) | The details of the payment | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdatePlan

> IamControllersResponse IamApiControllerUpdatePlan(ctx, id).IamObjectPlan(iamObjectPlan).Execute()

Api Controller Update Plan



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
	id := "id_example" // string | The id ( owner/name ) of the plan
	iamObjectPlan := *openapiclient.NewIamObjectPlan() // IamObjectPlan | The details of the plan

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerUpdatePlan(context.Background(), id).IamObjectPlan(iamObjectPlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerUpdatePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdatePlan`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerUpdatePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the plan | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectPlan** | [**IamObjectPlan**](IamObjectPlan.md) | The details of the plan | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdatePricing

> IamControllersResponse IamApiControllerUpdatePricing(ctx, id).IamObjectPricing(iamObjectPricing).Execute()

Api Controller Update Pricing



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
	id := "id_example" // string | The id ( owner/name ) of the pricing
	iamObjectPricing := *openapiclient.NewIamObjectPricing() // IamObjectPricing | The details of the pricing

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerUpdatePricing(context.Background(), id).IamObjectPricing(iamObjectPricing).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerUpdatePricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdatePricing`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerUpdatePricing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the pricing | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdatePricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectPricing** | [**IamObjectPricing**](IamObjectPricing.md) | The details of the pricing | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateProduct

> IamControllersResponse IamApiControllerUpdateProduct(ctx, id).IamObjectProduct(iamObjectProduct).Execute()

Api Controller Update Product



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
	id := "id_example" // string | The id ( owner/name ) of the product
	iamObjectProduct := *openapiclient.NewIamObjectProduct() // IamObjectProduct | The details of the product

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerUpdateProduct(context.Background(), id).IamObjectProduct(iamObjectProduct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerUpdateProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateProduct`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerUpdateProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the product | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectProduct** | [**IamObjectProduct**](IamObjectProduct.md) | The details of the product | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateSubscription

> IamControllersResponse IamApiControllerUpdateSubscription(ctx, id).IamObjectSubscription(iamObjectSubscription).Execute()

Api Controller Update Subscription



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
	id := "id_example" // string | The id ( owner/name ) of the subscription
	iamObjectSubscription := *openapiclient.NewIamObjectSubscription() // IamObjectSubscription | The details of the subscription

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerUpdateSubscription(context.Background(), id).IamObjectSubscription(iamObjectSubscription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerUpdateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateSubscription`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerUpdateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectSubscription** | [**IamObjectSubscription**](IamObjectSubscription.md) | The details of the subscription | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateTransaction

> IamControllersResponse IamApiControllerUpdateTransaction(ctx, id).IamObjectTransaction(iamObjectTransaction).Execute()

Api Controller Update Transaction



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
	id := "id_example" // string | The id ( owner/name ) of the transaction
	iamObjectTransaction := *openapiclient.NewIamObjectTransaction() // IamObjectTransaction | The details of the transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamPaymentsAPI.IamApiControllerUpdateTransaction(context.Background(), id).IamObjectTransaction(iamObjectTransaction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamPaymentsAPI.IamApiControllerUpdateTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateTransaction`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamPaymentsAPI.IamApiControllerUpdateTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectTransaction** | [**IamObjectTransaction**](IamObjectTransaction.md) | The details of the transaction | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

