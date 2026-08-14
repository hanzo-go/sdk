# \PaymentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPayment**](PaymentsAPI.md#GetPayment) | **Get** /v1/payments/{id} | Read one settled payment by its id
[**TakePayment**](PaymentsAPI.md#TakePayment) | **Post** /v1/payments | Take a card payment and credit the org&#39;s balance



## GetPayment

> PaymentRecord GetPayment(ctx, id).Execute()

Read one settled payment by its id



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
	id := "id_example" // string | ID is the ledger transaction id a payment returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.GetPayment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.GetPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPayment`: PaymentRecord
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.GetPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the ledger transaction id a payment returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentRecord**](PaymentRecord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TakePayment

> PaymentOut TakePayment(ctx).PaymentIn(paymentIn).Execute()

Take a card payment and credit the org's balance



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
	paymentIn := *openapiclient.NewPaymentIn() // PaymentIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.TakePayment(context.Background()).PaymentIn(paymentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.TakePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TakePayment`: PaymentOut
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.TakePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTakePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentIn** | [**PaymentIn**](PaymentIn.md) |  | 

### Return type

[**PaymentOut**](PaymentOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

