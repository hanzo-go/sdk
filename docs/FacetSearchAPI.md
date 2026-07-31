# \FacetSearchAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchFacetSearch**](FacetSearchAPI.md#SearchFacetSearch) | **Post** /v1/search/indexes/{indexUid}/facet-search | Search within facet values



## SearchFacetSearch

> SearchFacetSearch200Response SearchFacetSearch(ctx, indexUid).SearchFacetSearchRequest(searchFacetSearchRequest).Execute()

Search within facet values

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
	searchFacetSearchRequest := *openapiclient.NewSearchFacetSearchRequest("FacetName_example") // SearchFacetSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FacetSearchAPI.SearchFacetSearch(context.Background(), indexUid).SearchFacetSearchRequest(searchFacetSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacetSearchAPI.SearchFacetSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchFacetSearch`: SearchFacetSearch200Response
	fmt.Fprintf(os.Stdout, "Response from `FacetSearchAPI.SearchFacetSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchFacetSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchFacetSearchRequest** | [**SearchFacetSearchRequest**](SearchFacetSearchRequest.md) |  | 

### Return type

[**SearchFacetSearch200Response**](SearchFacetSearch200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

