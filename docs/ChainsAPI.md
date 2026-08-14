# \ChainsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChains**](ChainsAPI.md#GetChains) | **Get** /v1/chains | Reports the chains this deployment can reach.
[**GetChainsByChain**](ChainsAPI.md#GetChainsByChain) | **Get** /v1/chains/{chain} | Reports one chain and whether its upstream is answering.



## GetChains

> ChainList GetChains(ctx).Execute()

Reports the chains this deployment can reach.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChainsAPI.GetChains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChainsAPI.GetChains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChains`: ChainList
	fmt.Fprintf(os.Stdout, "Response from `ChainsAPI.GetChains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChainsRequest struct via the builder pattern


### Return type

[**ChainList**](ChainList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChainsByChain

> ChainStatus GetChainsByChain(ctx, chain).Execute()

Reports one chain and whether its upstream is answering.



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
	chain := "chain_example" // string | Chain is the registry id, as in /v1/chains/lux.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChainsAPI.GetChainsByChain(context.Background(), chain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChainsAPI.GetChainsByChain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChainsByChain`: ChainStatus
	fmt.Fprintf(os.Stdout, "Response from `ChainsAPI.GetChainsByChain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chain** | **string** | Chain is the registry id, as in /v1/chains/lux. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChainsByChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChainStatus**](ChainStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

