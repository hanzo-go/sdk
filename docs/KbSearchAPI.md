# \KbSearchAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KbKbSearch**](KbSearchAPI.md#KbKbSearch) | **Post** /v1/kb/search | Semantic search over the org&#39;s knowledge



## KbKbSearch

> KbKbSearch200Response KbKbSearch(ctx).KbSearchRequest(kbSearchRequest).Execute()

Semantic search over the org's knowledge

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
	kbSearchRequest := *openapiclient.NewKbSearchRequest("Query_example") // KbSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KbSearchAPI.KbKbSearch(context.Background()).KbSearchRequest(kbSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbSearchAPI.KbKbSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbKbSearch`: KbKbSearch200Response
	fmt.Fprintf(os.Stdout, "Response from `KbSearchAPI.KbKbSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKbKbSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kbSearchRequest** | [**KbSearchRequest**](KbSearchRequest.md) |  | 

### Return type

[**KbKbSearch200Response**](KbKbSearch200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

