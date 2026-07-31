# \MultiSearchAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchMultiSearch**](MultiSearchAPI.md#SearchMultiSearch) | **Post** /v1/search/multi-search | Perform a multi-index search



## SearchMultiSearch

> SearchMultiSearch200Response SearchMultiSearch(ctx).SearchFederatedSearch(searchFederatedSearch).Execute()

Perform a multi-index search



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
	searchFederatedSearch := *openapiclient.NewSearchFederatedSearch([]openapiclient.SearchSearchQueryWithIndex{*openapiclient.NewSearchSearchQueryWithIndex("IndexUid_example")}) // SearchFederatedSearch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiSearchAPI.SearchMultiSearch(context.Background()).SearchFederatedSearch(searchFederatedSearch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiSearchAPI.SearchMultiSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchMultiSearch`: SearchMultiSearch200Response
	fmt.Fprintf(os.Stdout, "Response from `MultiSearchAPI.SearchMultiSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchMultiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchFederatedSearch** | [**SearchFederatedSearch**](SearchFederatedSearch.md) |  | 

### Return type

[**SearchMultiSearch200Response**](SearchMultiSearch200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

