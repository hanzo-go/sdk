# \VectorAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVectorByName**](VectorAPI.md#DeleteVectorByName) | **Delete** /v1/vector/{name} | Deletes one vector collection from the shared backend and removes its metadata row.
[**GetVector**](VectorAPI.md#GetVector) | **Get** /v1/vector | Lists the caller org&#39;s vector collections.
[**GetVectorByName**](VectorAPI.md#GetVectorByName) | **Get** /v1/vector/{name} | Returns one vector collection&#39;s metadata.
[**GetVectorCollections**](VectorAPI.md#GetVectorCollections) | **Get** /v1/vector/collections | Lists the vector collections with their size and geometry.
[**GetVectorStats**](VectorAPI.md#GetVectorStats) | **Get** /v1/vector/stats | Totals the collections, vectors and storage across the vector store.
[**PostVector**](VectorAPI.md#PostVector) | **Post** /v1/vector | Provision a vector collection for your org



## DeleteVectorByName

> DeleteVectorByName(ctx, name).Execute()

Deletes one vector collection from the shared backend and removes its metadata row.



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
	name := "embeddings" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VectorAPI.DeleteVectorByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorAPI.DeleteVectorByName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteVectorByNameRequest struct via the builder pattern


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


## GetVector

> []ProvisionedSummary GetVector(ctx).Execute()

Lists the caller org's vector collections.



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
	resp, r, err := apiClient.VectorAPI.GetVector(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorAPI.GetVector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVector`: []ProvisionedSummary
	fmt.Fprintf(os.Stdout, "Response from `VectorAPI.GetVector`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVectorRequest struct via the builder pattern


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


## GetVectorByName

> ProvisionedResource GetVectorByName(ctx, name).Execute()

Returns one vector collection's metadata.



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
	name := "embeddings" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VectorAPI.GetVectorByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorAPI.GetVectorByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVectorByName`: ProvisionedResource
	fmt.Fprintf(os.Stdout, "Response from `VectorAPI.GetVectorByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVectorByNameRequest struct via the builder pattern


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


## PostVector

> ProvisionResult PostVector(ctx).ProvisionRequest(provisionRequest).Execute()

Provision a vector collection for your org



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
	resp, r, err := apiClient.VectorAPI.PostVector(context.Background()).ProvisionRequest(provisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorAPI.PostVector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVector`: ProvisionResult
	fmt.Fprintf(os.Stdout, "Response from `VectorAPI.PostVector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostVectorRequest struct via the builder pattern


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

