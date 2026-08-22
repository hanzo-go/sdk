# \ExplorerAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExplorerIndexers**](ExplorerAPI.md#GetExplorerIndexers) | **Get** /v1/explorer/indexers | Reports the deployment&#39;s chain indexer(s) and how far each has indexed.
[**GetExplorerOracles**](ExplorerAPI.md#GetExplorerOracles) | **Get** /v1/explorer/oracles | Reports the on-chain price/data oracles from the graph&#39;s O-Chain PriceFeed registry.



## GetExplorerIndexers

> IndexersOut GetExplorerIndexers(ctx).Execute()

Reports the deployment's chain indexer(s) and how far each has indexed.



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
	resp, r, err := apiClient.ExplorerAPI.GetExplorerIndexers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExplorerAPI.GetExplorerIndexers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExplorerIndexers`: IndexersOut
	fmt.Fprintf(os.Stdout, "Response from `ExplorerAPI.GetExplorerIndexers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExplorerIndexersRequest struct via the builder pattern


### Return type

[**IndexersOut**](IndexersOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExplorerOracles

> OraclesOut GetExplorerOracles(ctx).Execute()

Reports the on-chain price/data oracles from the graph's O-Chain PriceFeed registry.



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
	resp, r, err := apiClient.ExplorerAPI.GetExplorerOracles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExplorerAPI.GetExplorerOracles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExplorerOracles`: OraclesOut
	fmt.Fprintf(os.Stdout, "Response from `ExplorerAPI.GetExplorerOracles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExplorerOraclesRequest struct via the builder pattern


### Return type

[**OraclesOut**](OraclesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

