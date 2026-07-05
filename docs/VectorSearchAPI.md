# \VectorSearchAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VectorRecommendPoints**](VectorSearchAPI.md#VectorRecommendPoints) | **Post** /v1/vector/collections/{collection_name}/points/recommend | Recommend points
[**VectorSearchBatch**](VectorSearchAPI.md#VectorSearchBatch) | **Post** /v1/vector/collections/{collection_name}/points/search/batch | Batch search
[**VectorSearchPoints**](VectorSearchAPI.md#VectorSearchPoints) | **Post** /v1/vector/collections/{collection_name}/points/search | Search points



## VectorRecommendPoints

> VectorRecommendPoints200Response VectorRecommendPoints(ctx, collectionName).VectorRecommendPointsRequest(vectorRecommendPointsRequest).Execute()

Recommend points

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
	collectionName := "collectionName_example" // string | 
	vectorRecommendPointsRequest := *openapiclient.NewVectorRecommendPointsRequest([]openapiclient.VectorDeletePointsRequestPointsInner{openapiclient.vector_deletePoints_request_points_inner{Int32: new(int32)}}) // VectorRecommendPointsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VectorSearchAPI.VectorRecommendPoints(context.Background(), collectionName).VectorRecommendPointsRequest(vectorRecommendPointsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorSearchAPI.VectorRecommendPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VectorRecommendPoints`: VectorRecommendPoints200Response
	fmt.Fprintf(os.Stdout, "Response from `VectorSearchAPI.VectorRecommendPoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVectorRecommendPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vectorRecommendPointsRequest** | [**VectorRecommendPointsRequest**](VectorRecommendPointsRequest.md) |  | 

### Return type

[**VectorRecommendPoints200Response**](VectorRecommendPoints200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VectorSearchBatch

> VectorSearchBatch200Response VectorSearchBatch(ctx, collectionName).VectorSearchBatchRequest(vectorSearchBatchRequest).Execute()

Batch search

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
	collectionName := "collectionName_example" // string | 
	vectorSearchBatchRequest := *openapiclient.NewVectorSearchBatchRequest([]openapiclient.VectorSearchRequest{*openapiclient.NewVectorSearchRequest(openapiclient.vector_SearchRequest_vector{VectorSearchRequestVectorOneOf: openapiclient.NewVectorSearchRequestVectorOneOf()})}) // VectorSearchBatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VectorSearchAPI.VectorSearchBatch(context.Background(), collectionName).VectorSearchBatchRequest(vectorSearchBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorSearchAPI.VectorSearchBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VectorSearchBatch`: VectorSearchBatch200Response
	fmt.Fprintf(os.Stdout, "Response from `VectorSearchAPI.VectorSearchBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVectorSearchBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vectorSearchBatchRequest** | [**VectorSearchBatchRequest**](VectorSearchBatchRequest.md) |  | 

### Return type

[**VectorSearchBatch200Response**](VectorSearchBatch200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VectorSearchPoints

> VectorRecommendPoints200Response VectorSearchPoints(ctx, collectionName).VectorSearchRequest(vectorSearchRequest).Execute()

Search points

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
	collectionName := "collectionName_example" // string | 
	vectorSearchRequest := *openapiclient.NewVectorSearchRequest(openapiclient.vector_SearchRequest_vector{VectorSearchRequestVectorOneOf: openapiclient.NewVectorSearchRequestVectorOneOf()}) // VectorSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VectorSearchAPI.VectorSearchPoints(context.Background(), collectionName).VectorSearchRequest(vectorSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorSearchAPI.VectorSearchPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VectorSearchPoints`: VectorRecommendPoints200Response
	fmt.Fprintf(os.Stdout, "Response from `VectorSearchAPI.VectorSearchPoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVectorSearchPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vectorSearchRequest** | [**VectorSearchRequest**](VectorSearchRequest.md) |  | 

### Return type

[**VectorRecommendPoints200Response**](VectorRecommendPoints200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

