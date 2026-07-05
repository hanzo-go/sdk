# \WebsearchSearchAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebsearchWebSearch**](WebsearchSearchAPI.md#WebsearchWebSearch) | **Get** /v1/websearch/search | Search the web (SearXNG JSON contract, proxied verbatim)



## WebsearchWebSearch

> WebsearchSearchResponse WebsearchWebSearch(ctx).Q(q).Format(format).Execute()

Search the web (SearXNG JSON contract, proxied verbatim)

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
	q := "q_example" // string | Search query
	format := "format_example" // string |  (optional) (default to "json")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebsearchSearchAPI.WebsearchWebSearch(context.Background()).Q(q).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebsearchSearchAPI.WebsearchWebSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebsearchWebSearch`: WebsearchSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `WebsearchSearchAPI.WebsearchWebSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebsearchWebSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Search query | 
 **format** | **string** |  | [default to &quot;json&quot;]

### Return type

[**WebsearchSearchResponse**](WebsearchSearchResponse.md)

### Authorization

[serviceKey](../README.md#serviceKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

