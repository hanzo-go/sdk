# \SearchAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSearchIndexes**](SearchAPI.md#GetSearchIndexes) | **Get** /v1/search/indexes | Lists the search indexes with their document counts and timestamps.
[**GetSearchStats**](SearchAPI.md#GetSearchStats) | **Get** /v1/search/stats | Totals the documents across every search index.
[**Search**](SearchAPI.md#Search) | **Post** /v1/search | Hybrid search over the org&#39;s own corpora



## GetSearchIndexes

> SearchIndexList GetSearchIndexes(ctx).Authorization(authorization).Execute()

Lists the search indexes with their document counts and timestamps.



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
	authorization := "authorization_example" // string | Authorization carries the surface's bearer key (`Bearer <key>`); the bare key is accepted too. Search and vector are two surfaces with two keys. It is not `validate:\"required\"` on purpose: requireKey answers absence itself, so an unconfigured surface 503s and a missing bearer 401s — a validation refusal would rewrite both statuses. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.GetSearchIndexes(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.GetSearchIndexes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchIndexes`: SearchIndexList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.GetSearchIndexes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchIndexesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization carries the surface&#39;s bearer key (&#x60;Bearer &lt;key&gt;&#x60;); the bare key is accepted too. Search and vector are two surfaces with two keys. It is not &#x60;validate:\&quot;required\&quot;&#x60; on purpose: requireKey answers absence itself, so an unconfigured surface 503s and a missing bearer 401s — a validation refusal would rewrite both statuses. | 

### Return type

[**SearchIndexList**](SearchIndexList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchStats

> SearchStats GetSearchStats(ctx).Authorization(authorization).Execute()

Totals the documents across every search index.



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
	authorization := "authorization_example" // string | Authorization carries the surface's bearer key (`Bearer <key>`); the bare key is accepted too. Search and vector are two surfaces with two keys. It is not `validate:\"required\"` on purpose: requireKey answers absence itself, so an unconfigured surface 503s and a missing bearer 401s — a validation refusal would rewrite both statuses. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.GetSearchStats(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.GetSearchStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchStats`: SearchStats
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.GetSearchStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization carries the surface&#39;s bearer key (&#x60;Bearer &lt;key&gt;&#x60;); the bare key is accepted too. Search and vector are two surfaces with two keys. It is not &#x60;validate:\&quot;required\&quot;&#x60; on purpose: requireKey answers absence itself, so an unconfigured surface 503s and a missing bearer 401s — a validation refusal would rewrite both statuses. | 

### Return type

[**SearchStats**](SearchStats.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> Response Search(ctx).Request(request).Execute()

Hybrid search over the org's own corpora



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
	request := *openapiclient.NewRequest() // Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.Search(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.Search``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Search`: Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**Request**](Request.md) |  | 

### Return type

[**Response**](Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

