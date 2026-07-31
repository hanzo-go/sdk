# \SwapAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchSwapIndexes**](SwapAPI.md#SearchSwapIndexes) | **Post** /v1/search/swap-indexes | Swap two index identifiers



## SearchSwapIndexes

> SearchSummarizedTaskView SearchSwapIndexes(ctx).SearchSwapIndexesRequestInner(searchSwapIndexesRequestInner).Execute()

Swap two index identifiers

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
	searchSwapIndexesRequestInner := []openapiclient.SearchSwapIndexesRequestInner{*openapiclient.NewSearchSwapIndexesRequestInner([]string{"Indexes_example"})} // []SearchSwapIndexesRequestInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwapAPI.SearchSwapIndexes(context.Background()).SearchSwapIndexesRequestInner(searchSwapIndexesRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwapAPI.SearchSwapIndexes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSwapIndexes`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `SwapAPI.SearchSwapIndexes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSwapIndexesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSwapIndexesRequestInner** | [**[]SearchSwapIndexesRequestInner**](SearchSwapIndexesRequestInner.md) |  | 

### Return type

[**SearchSummarizedTaskView**](SearchSummarizedTaskView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

