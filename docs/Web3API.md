# \Web3API

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWeb3Chains**](Web3API.md#GetWeb3Chains) | **Get** /v1/web3/chains | Reports the chains this deployment can reach.
[**GetWeb3ChainsByChain**](Web3API.md#GetWeb3ChainsByChain) | **Get** /v1/web3/chains/{chain} | Reports one chain and whether its upstream is answering.
[**GetWeb3TokensByChainByAddress**](Web3API.md#GetWeb3TokensByChainByAddress) | **Get** /v1/web3/tokens/{chain}/{address} | Reads an address&#39;s native balance on a chain.
[**PostWeb3RpcByChain**](Web3API.md#PostWeb3RpcByChain) | **Post** /v1/web3/rpc/{chain} | Forwards a JSON-RPC call to the named chain and returns its answer unchanged.



## GetWeb3Chains

> ChainList GetWeb3Chains(ctx).Execute()

Reports the chains this deployment can reach.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Web3API.GetWeb3Chains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Web3API.GetWeb3Chains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWeb3Chains`: ChainList
	fmt.Fprintf(os.Stdout, "Response from `Web3API.GetWeb3Chains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3ChainsRequest struct via the builder pattern


### Return type

[**ChainList**](ChainList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWeb3ChainsByChain

> ChainStatus GetWeb3ChainsByChain(ctx, chain).Execute()

Reports one chain and whether its upstream is answering.



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
	chain := "chain_example" // string | Chain is the registry id, as in /v1/web3/chains/lux.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Web3API.GetWeb3ChainsByChain(context.Background(), chain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Web3API.GetWeb3ChainsByChain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWeb3ChainsByChain`: ChainStatus
	fmt.Fprintf(os.Stdout, "Response from `Web3API.GetWeb3ChainsByChain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** | Chain is the registry id, as in /v1/web3/chains/lux. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3ChainsByChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChainStatus**](ChainStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWeb3TokensByChainByAddress

> Balances GetWeb3TokensByChainByAddress(ctx, chain, address).Execute()

Reads an address's native balance on a chain.



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
	chain := "chain_example" // string | Chain is the registry id.
	address := "address_example" // string | Address is the account, 0x-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Web3API.GetWeb3TokensByChainByAddress(context.Background(), chain, address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Web3API.GetWeb3TokensByChainByAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWeb3TokensByChainByAddress`: Balances
	fmt.Fprintf(os.Stdout, "Response from `Web3API.GetWeb3TokensByChainByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** | Chain is the registry id. | 
**address** | **string** | Address is the account, 0x-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeb3TokensByChainByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Balances**](Balances.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWeb3RpcByChain

> RpcOut PostWeb3RpcByChain(ctx, chain).RpcIn(rpcIn).Execute()

Forwards a JSON-RPC call to the named chain and returns its answer unchanged.



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
	chain := "chain_example" // string | Chain is the registry id, from the URL.
	rpcIn := *openapiclient.NewRpcIn() // RpcIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Web3API.PostWeb3RpcByChain(context.Background(), chain).RpcIn(rpcIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Web3API.PostWeb3RpcByChain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWeb3RpcByChain`: RpcOut
	fmt.Fprintf(os.Stdout, "Response from `Web3API.PostWeb3RpcByChain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** | Chain is the registry id, from the URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWeb3RpcByChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rpcIn** | [**RpcIn**](RpcIn.md) |  | 

### Return type

[**RpcOut**](RpcOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

