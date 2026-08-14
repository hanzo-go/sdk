# \RpcAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostRpcByChain**](RpcAPI.md#PostRpcByChain) | **Post** /v1/rpc/{chain} | Forwards a JSON-RPC call to the named chain and returns its answer unchanged.



## PostRpcByChain

> RpcOut PostRpcByChain(ctx, chain).RpcIn(rpcIn).Execute()

Forwards a JSON-RPC call to the named chain and returns its answer unchanged.



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
	chain := "chain_example" // string | Chain is the registry id, from the URL.
	rpcIn := *openapiclient.NewRpcIn() // RpcIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RpcAPI.PostRpcByChain(context.Background(), chain).RpcIn(rpcIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RpcAPI.PostRpcByChain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRpcByChain`: RpcOut
	fmt.Fprintf(os.Stdout, "Response from `RpcAPI.PostRpcByChain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** | Chain is the registry id, from the URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRpcByChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rpcIn** | [**RpcIn**](RpcIn.md) |  | 

### Return type

[**RpcOut**](RpcOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

