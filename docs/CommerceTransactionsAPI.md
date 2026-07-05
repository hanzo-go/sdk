# \CommerceTransactionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommerceCreateHold**](CommerceTransactionsAPI.md#CommerceCreateHold) | **Post** /v1/commerce/transaction/hold | Create hold
[**CommerceCreateTransaction**](CommerceTransactionsAPI.md#CommerceCreateTransaction) | **Post** /v1/commerce/transaction | Create transaction
[**CommerceListTransactions**](CommerceTransactionsAPI.md#CommerceListTransactions) | **Get** /v1/commerce/transaction/{kind}/{id} | List transactions for entity
[**CommerceRemoveHold**](CommerceTransactionsAPI.md#CommerceRemoveHold) | **Delete** /v1/commerce/transaction/hold/{id} | Remove hold



## CommerceCreateHold

> CommerceTransaction CommerceCreateHold(ctx).CommerceTransaction(commerceTransaction).Execute()

Create hold

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
	commerceTransaction := *openapiclient.NewCommerceTransaction() // CommerceTransaction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceTransactionsAPI.CommerceCreateHold(context.Background()).CommerceTransaction(commerceTransaction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceTransactionsAPI.CommerceCreateHold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateHold`: CommerceTransaction
	fmt.Fprintf(os.Stdout, "Response from `CommerceTransactionsAPI.CommerceCreateHold`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateHoldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceTransaction** | [**CommerceTransaction**](CommerceTransaction.md) |  | 

### Return type

[**CommerceTransaction**](CommerceTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceCreateTransaction

> CommerceTransaction CommerceCreateTransaction(ctx).CommerceTransaction(commerceTransaction).Execute()

Create transaction

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
	commerceTransaction := *openapiclient.NewCommerceTransaction() // CommerceTransaction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceTransactionsAPI.CommerceCreateTransaction(context.Background()).CommerceTransaction(commerceTransaction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceTransactionsAPI.CommerceCreateTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateTransaction`: CommerceTransaction
	fmt.Fprintf(os.Stdout, "Response from `CommerceTransactionsAPI.CommerceCreateTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceTransaction** | [**CommerceTransaction**](CommerceTransaction.md) |  | 

### Return type

[**CommerceTransaction**](CommerceTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceListTransactions

> []CommerceTransaction CommerceListTransactions(ctx, kind, id).Execute()

List transactions for entity

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
	kind := "kind_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceTransactionsAPI.CommerceListTransactions(context.Background(), kind, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceTransactionsAPI.CommerceListTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceListTransactions`: []CommerceTransaction
	fmt.Fprintf(os.Stdout, "Response from `CommerceTransactionsAPI.CommerceListTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceListTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]CommerceTransaction**](CommerceTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceRemoveHold

> CommerceRemoveHold(ctx, id).Execute()

Remove hold

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceTransactionsAPI.CommerceRemoveHold(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceTransactionsAPI.CommerceRemoveHold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceRemoveHoldRequest struct via the builder pattern


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

