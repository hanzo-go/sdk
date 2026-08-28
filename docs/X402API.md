# \X402API

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetX402SettlementsById**](X402API.md#GetX402SettlementsById) | **Get** /v1/x402/settlements/{id} | Settlement reads one x402 payment receipt by id.



## GetX402SettlementsById

> Receipt GetX402SettlementsById(ctx, id).Execute()

Settlement reads one x402 payment receipt by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID is the settlement id from the URL — the deterministic keccak(from|nonce) key an x402 receipt is issued under (the `id` field of a Receipt, and the `transaction` of the SettlementResponse on the PAYMENT-RESPONSE header a paid request answers with).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.X402API.GetX402SettlementsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `X402API.GetX402SettlementsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetX402SettlementsById`: Receipt
	fmt.Fprintf(os.Stdout, "Response from `X402API.GetX402SettlementsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the settlement id from the URL — the deterministic keccak(from|nonce) key an x402 receipt is issued under (the &#x60;id&#x60; field of a Receipt, and the &#x60;transaction&#x60; of the SettlementResponse on the PAYMENT-RESPONSE header a paid request answers with). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetX402SettlementsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Receipt**](Receipt.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

