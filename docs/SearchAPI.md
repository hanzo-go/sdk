# \SearchAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSearchByName**](SearchAPI.md#DeleteSearchByName) | **Delete** /v1/search/{name} | Deletes one search index from the shared backend and removes its metadata row.
[**GetSearch**](SearchAPI.md#GetSearch) | **Get** /v1/search | Lists the caller org&#39;s search indexes.
[**GetSearchByName**](SearchAPI.md#GetSearchByName) | **Get** /v1/search/{name} | Returns one search index&#39;s metadata.
[**GetSearchIndexes**](SearchAPI.md#GetSearchIndexes) | **Get** /v1/search/indexes | Lists the search indexes with their document counts and timestamps.
[**GetSearchStats**](SearchAPI.md#GetSearchStats) | **Get** /v1/search/stats | Totals the documents across every search index.
[**PostSearch**](SearchAPI.md#PostSearch) | **Post** /v1/search | Provision a search index for your org



## DeleteSearchByName

> DeleteSearchByName(ctx, name).Execute()

Deletes one search index from the shared backend and removes its metadata row.



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
	name := "products" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SearchAPI.DeleteSearchByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.DeleteSearchByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSearchByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearch

> []ProvisionedSummary GetSearch(ctx).Execute()

Lists the caller org's search indexes.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.GetSearch(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.GetSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearch`: []ProvisionedSummary
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.GetSearch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchRequest struct via the builder pattern


### Return type

[**[]ProvisionedSummary**](ProvisionedSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchByName

> ProvisionedResource GetSearchByName(ctx, name).Execute()

Returns one search index's metadata.



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
	name := "products" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.GetSearchByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.GetSearchByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchByName`: ProvisionedResource
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.GetSearchByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProvisionedResource**](ProvisionedResource.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## PostSearch

> ProvisionResult PostSearch(ctx).ProvisionRequest(provisionRequest).Execute()

Provision a search index for your org



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
	provisionRequest := *openapiclient.NewProvisionRequest() // ProvisionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.PostSearch(context.Background()).ProvisionRequest(provisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.PostSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSearch`: ProvisionResult
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.PostSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisionRequest** | [**ProvisionRequest**](ProvisionRequest.md) |  | 

### Return type

[**ProvisionResult**](ProvisionResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

