# \SearchSearchAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchSearchGet**](SearchSearchAPI.md#SearchSearchGet) | **Get** /v1/search/indexes/{indexUid}/search | Search documents (GET)
[**SearchSearchPost**](SearchSearchAPI.md#SearchSearchPost) | **Post** /v1/search/indexes/{indexUid}/search | Search documents (POST)



## SearchSearchGet

> SearchSearchResult SearchSearchGet(ctx, indexUid).Q(q).Offset(offset).Limit(limit).AttributesToRetrieve(attributesToRetrieve).AttributesToHighlight(attributesToHighlight).AttributesToCrop(attributesToCrop).CropLength(cropLength).Filter(filter).Sort(sort).Facets(facets).ShowMatchesPosition(showMatchesPosition).ShowRankingScore(showRankingScore).MatchingStrategy(matchingStrategy).Execute()

Search documents (GET)

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
	q := "q_example" // string | Search query (optional)
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 20)
	attributesToRetrieve := "attributesToRetrieve_example" // string | Comma-separated list of attributes to return (optional)
	attributesToHighlight := "attributesToHighlight_example" // string |  (optional)
	attributesToCrop := "attributesToCrop_example" // string |  (optional)
	cropLength := int32(56) // int32 |  (optional) (default to 10)
	filter := "filter_example" // string | Filter expression (optional)
	sort := "sort_example" // string | Comma-separated sort rules (optional)
	facets := "facets_example" // string | Comma-separated facet attributes (optional)
	showMatchesPosition := true // bool |  (optional)
	showRankingScore := true // bool |  (optional)
	matchingStrategy := "matchingStrategy_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchSearchAPI.SearchSearchGet(context.Background(), indexUid).Q(q).Offset(offset).Limit(limit).AttributesToRetrieve(attributesToRetrieve).AttributesToHighlight(attributesToHighlight).AttributesToCrop(attributesToCrop).CropLength(cropLength).Filter(filter).Sort(sort).Facets(facets).ShowMatchesPosition(showMatchesPosition).ShowRankingScore(showRankingScore).MatchingStrategy(matchingStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchSearchAPI.SearchSearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSearchGet`: SearchSearchResult
	fmt.Fprintf(os.Stdout, "Response from `SearchSearchAPI.SearchSearchGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Search query | 
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]
 **attributesToRetrieve** | **string** | Comma-separated list of attributes to return | 
 **attributesToHighlight** | **string** |  | 
 **attributesToCrop** | **string** |  | 
 **cropLength** | **int32** |  | [default to 10]
 **filter** | **string** | Filter expression | 
 **sort** | **string** | Comma-separated sort rules | 
 **facets** | **string** | Comma-separated facet attributes | 
 **showMatchesPosition** | **bool** |  | 
 **showRankingScore** | **bool** |  | 
 **matchingStrategy** | **string** |  | 

### Return type

[**SearchSearchResult**](SearchSearchResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSearchPost

> SearchSearchResult SearchSearchPost(ctx, indexUid).SearchSearchQuery(searchSearchQuery).Execute()

Search documents (POST)

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
	searchSearchQuery := *openapiclient.NewSearchSearchQuery() // SearchSearchQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchSearchAPI.SearchSearchPost(context.Background(), indexUid).SearchSearchQuery(searchSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchSearchAPI.SearchSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSearchPost`: SearchSearchResult
	fmt.Fprintf(os.Stdout, "Response from `SearchSearchAPI.SearchSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchSearchQuery** | [**SearchSearchQuery**](SearchSearchQuery.md) |  | 

### Return type

[**SearchSearchResult**](SearchSearchResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

