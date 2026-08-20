# \VectorAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVectorCollections**](VectorAPI.md#GetVectorCollections) | **Get** /v1/vector/collections | Lists the vector collections with their size and geometry.
[**GetVectorStats**](VectorAPI.md#GetVectorStats) | **Get** /v1/vector/stats | Totals the collections, vectors and storage across the vector store.



## GetVectorCollections

> VectorCollectionList GetVectorCollections(ctx).Authorization(authorization).Execute()

Lists the vector collections with their size and geometry.



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
	resp, r, err := apiClient.VectorAPI.GetVectorCollections(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorAPI.GetVectorCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVectorCollections`: VectorCollectionList
	fmt.Fprintf(os.Stdout, "Response from `VectorAPI.GetVectorCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVectorCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization carries the surface&#39;s bearer key (&#x60;Bearer &lt;key&gt;&#x60;); the bare key is accepted too. Search and vector are two surfaces with two keys. It is not &#x60;validate:\&quot;required\&quot;&#x60; on purpose: requireKey answers absence itself, so an unconfigured surface 503s and a missing bearer 401s — a validation refusal would rewrite both statuses. | 

### Return type

[**VectorCollectionList**](VectorCollectionList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVectorStats

> VectorStats GetVectorStats(ctx).Authorization(authorization).Execute()

Totals the collections, vectors and storage across the vector store.



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
	resp, r, err := apiClient.VectorAPI.GetVectorStats(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorAPI.GetVectorStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVectorStats`: VectorStats
	fmt.Fprintf(os.Stdout, "Response from `VectorAPI.GetVectorStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVectorStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Authorization carries the surface&#39;s bearer key (&#x60;Bearer &lt;key&gt;&#x60;); the bare key is accepted too. Search and vector are two surfaces with two keys. It is not &#x60;validate:\&quot;required\&quot;&#x60; on purpose: requireKey answers absence itself, so an unconfigured surface 503s and a missing bearer 401s — a validation refusal would rewrite both statuses. | 

### Return type

[**VectorStats**](VectorStats.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

