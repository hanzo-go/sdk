# \VectorPointsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VectorDeletePoints**](VectorPointsAPI.md#VectorDeletePoints) | **Post** /v1/vector/collections/{collection_name}/points/delete | Delete points
[**VectorGetPoint**](VectorPointsAPI.md#VectorGetPoint) | **Get** /v1/vector/collections/{collection_name}/points/{id} | Get point
[**VectorUpsertPoints**](VectorPointsAPI.md#VectorUpsertPoints) | **Put** /v1/vector/collections/{collection_name}/points | Upsert points



## VectorDeletePoints

> VectorDeletePoints200Response VectorDeletePoints(ctx, collectionName).VectorDeletePointsRequest(vectorDeletePointsRequest).Wait(wait).Execute()

Delete points

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
	vectorDeletePointsRequest := *openapiclient.NewVectorDeletePointsRequest() // VectorDeletePointsRequest | 
	wait := true // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VectorPointsAPI.VectorDeletePoints(context.Background(), collectionName).VectorDeletePointsRequest(vectorDeletePointsRequest).Wait(wait).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorPointsAPI.VectorDeletePoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VectorDeletePoints`: VectorDeletePoints200Response
	fmt.Fprintf(os.Stdout, "Response from `VectorPointsAPI.VectorDeletePoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVectorDeletePointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vectorDeletePointsRequest** | [**VectorDeletePointsRequest**](VectorDeletePointsRequest.md) |  | 
 **wait** | **bool** |  | [default to true]

### Return type

[**VectorDeletePoints200Response**](VectorDeletePoints200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VectorGetPoint

> VectorGetPoint200Response VectorGetPoint(ctx, collectionName, id).Execute()

Get point

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
	id := openapiclient.vector_deletePoints_request_points_inner{Int32: new(int32)} // VectorDeletePointsRequestPointsInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VectorPointsAPI.VectorGetPoint(context.Background(), collectionName, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorPointsAPI.VectorGetPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VectorGetPoint`: VectorGetPoint200Response
	fmt.Fprintf(os.Stdout, "Response from `VectorPointsAPI.VectorGetPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** |  | 
**id** | [**VectorDeletePointsRequestPointsInner**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVectorGetPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VectorGetPoint200Response**](VectorGetPoint200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VectorUpsertPoints

> VectorUpsertPoints200Response VectorUpsertPoints(ctx, collectionName).VectorUpsertPointsRequest(vectorUpsertPointsRequest).Wait(wait).Execute()

Upsert points

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
	vectorUpsertPointsRequest := *openapiclient.NewVectorUpsertPointsRequest([]openapiclient.VectorPointStruct{*openapiclient.NewVectorPointStruct(openapiclient.vector_PointStruct_id{Int32: new(int32)}, openapiclient.vector_PointStruct_vector{ArrayOfFloat32: new([]float32)})}) // VectorUpsertPointsRequest | 
	wait := true // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VectorPointsAPI.VectorUpsertPoints(context.Background(), collectionName).VectorUpsertPointsRequest(vectorUpsertPointsRequest).Wait(wait).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorPointsAPI.VectorUpsertPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VectorUpsertPoints`: VectorUpsertPoints200Response
	fmt.Fprintf(os.Stdout, "Response from `VectorPointsAPI.VectorUpsertPoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVectorUpsertPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vectorUpsertPointsRequest** | [**VectorUpsertPointsRequest**](VectorUpsertPointsRequest.md) |  | 
 **wait** | **bool** |  | [default to true]

### Return type

[**VectorUpsertPoints200Response**](VectorUpsertPoints200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

