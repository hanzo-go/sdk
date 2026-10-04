# \SearchAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Search**](SearchAPI.md#Search) | **Post** /v1/search | Hybrid search over the org&#39;s own corpora



## Search

> SearchFusion Search(ctx).SearchRequest(searchRequest).Execute()

Hybrid search over the org's own corpora



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
	searchRequest := *openapiclient.NewSearchRequest() // SearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.Search(context.Background()).SearchRequest(searchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.Search``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Search`: SearchFusion
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchRequest** | [**SearchRequest**](SearchRequest.md) |  | 

### Return type

[**SearchFusion**](SearchFusion.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

