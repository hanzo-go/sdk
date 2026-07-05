# \SearchSimilarAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchFindSimilar**](SearchSimilarAPI.md#SearchFindSimilar) | **Post** /v1/search/indexes/{indexUid}/similar | Find similar documents



## SearchFindSimilar

> SearchSimilarResult SearchFindSimilar(ctx, indexUid).SearchSimilarQuery(searchSimilarQuery).Execute()

Find similar documents

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
	indexUid := "indexUid_example" // string | Unique index identifier
	searchSimilarQuery := *openapiclient.NewSearchSimilarQuery(openapiclient.search_SimilarQuery_id{Int32: new(int32)}) // SearchSimilarQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchSimilarAPI.SearchFindSimilar(context.Background(), indexUid).SearchSimilarQuery(searchSimilarQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchSimilarAPI.SearchFindSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchFindSimilar`: SearchSimilarResult
	fmt.Fprintf(os.Stdout, "Response from `SearchSimilarAPI.SearchFindSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchFindSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchSimilarQuery** | [**SearchSimilarQuery**](SearchSimilarQuery.md) |  | 

### Return type

[**SearchSimilarResult**](SearchSimilarResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

