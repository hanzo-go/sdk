# \X402API

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1X402SettlementsId**](X402API.md#CloudGetV1X402SettlementsId) | **Get** /v1/x402/settlements/{id} | Settlement reads one x402 payment receipt by id.



## CloudGetV1X402SettlementsId

> CloudReceipt CloudGetV1X402SettlementsId(ctx, id).Execute()

Settlement reads one x402 payment receipt by id.



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
	id := "id_example" // string | ID is the settlement id from the URL — the deterministic keccak(from|nonce) key an x402 receipt is issued under (the `id` field of a Receipt, and the value of the X-Payment-Response header a paid request answers with).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.X402API.CloudGetV1X402SettlementsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `X402API.CloudGetV1X402SettlementsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1X402SettlementsId`: CloudReceipt
	fmt.Fprintf(os.Stdout, "Response from `X402API.CloudGetV1X402SettlementsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the settlement id from the URL — the deterministic keccak(from|nonce) key an x402 receipt is issued under (the &#x60;id&#x60; field of a Receipt, and the value of the X-Payment-Response header a paid request answers with). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1X402SettlementsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudReceipt**](CloudReceipt.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

