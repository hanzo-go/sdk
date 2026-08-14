# \TokensAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTokensByChainByAddress**](TokensAPI.md#GetTokensByChainByAddress) | **Get** /v1/tokens/{chain}/{address} | Reads an address&#39;s native balance on a chain.



## GetTokensByChainByAddress

> Balances GetTokensByChainByAddress(ctx, chain, address).Execute()

Reads an address's native balance on a chain.



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
	chain := "chain_example" // string | Chain is the registry id.
	address := "address_example" // string | Address is the account, 0x-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.GetTokensByChainByAddress(context.Background(), chain, address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.GetTokensByChainByAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokensByChainByAddress`: Balances
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.GetTokensByChainByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** | Chain is the registry id. | 
**address** | **string** | Address is the account, 0x-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensByChainByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Balances**](Balances.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

