# \X402API

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetX402Settlements**](X402API.md#GetX402Settlements) | **Get** /v1/x402/settlements | Lists the caller&#39;s x402 receipts, newest first: as payer, what its ledger paid; as payee, what it was paid — each settled payment with what it bought, both parties, the exact amount and when.
[**GetX402SettlementsById**](X402API.md#GetX402SettlementsById) | **Get** /v1/x402/settlements/{id} | Reads one x402 payment receipt by id.



## GetX402Settlements

> X402SettlementList GetX402Settlements(ctx).Role(role).Year(year).Execute()

Lists the caller's x402 receipts, newest first: as payer, what its ledger paid; as payee, what it was paid — each settled payment with what it bought, both parties, the exact amount and when.



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
	role := "payee" // string | Role is payer — what the caller's org paid — or payee — what it was paid. Payer when empty. (optional)
	year := int64(2026) // int64 | Year keeps the calendar year (UTC) the payments settled in. Zero keeps every year. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.X402API.GetX402Settlements(context.Background()).Role(role).Year(year).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `X402API.GetX402Settlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetX402Settlements`: X402SettlementList
	fmt.Fprintf(os.Stdout, "Response from `X402API.GetX402Settlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetX402SettlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | **string** | Role is payer — what the caller&#39;s org paid — or payee — what it was paid. Payer when empty. | 
 **year** | **int64** | Year keeps the calendar year (UTC) the payments settled in. Zero keeps every year. | 

### Return type

[**X402SettlementList**](X402SettlementList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetX402SettlementsById

> X402Receipt GetX402SettlementsById(ctx, id).Execute()

Reads one x402 payment receipt by id.



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
	// response from `GetX402SettlementsById`: X402Receipt
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

[**X402Receipt**](X402Receipt.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

