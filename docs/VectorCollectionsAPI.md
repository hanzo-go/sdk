# \VectorCollectionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VectorCreateCollection**](VectorCollectionsAPI.md#VectorCreateCollection) | **Put** /v1/vector/collections/{collection_name} | Create collection
[**VectorDeleteCollection**](VectorCollectionsAPI.md#VectorDeleteCollection) | **Delete** /v1/vector/collections/{collection_name} | Delete collection
[**VectorGetCollection**](VectorCollectionsAPI.md#VectorGetCollection) | **Get** /v1/vector/collections/{collection_name} | Get collection info
[**VectorListCollections**](VectorCollectionsAPI.md#VectorListCollections) | **Get** /v1/vector/collections | List collections



## VectorCreateCollection

> VectorCreateCollection200Response VectorCreateCollection(ctx, collectionName).VectorCreateCollectionRequest(vectorCreateCollectionRequest).Execute()

Create collection

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
	vectorCreateCollectionRequest := *openapiclient.NewVectorCreateCollectionRequest() // VectorCreateCollectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VectorCollectionsAPI.VectorCreateCollection(context.Background(), collectionName).VectorCreateCollectionRequest(vectorCreateCollectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorCollectionsAPI.VectorCreateCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VectorCreateCollection`: VectorCreateCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `VectorCollectionsAPI.VectorCreateCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVectorCreateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vectorCreateCollectionRequest** | [**VectorCreateCollectionRequest**](VectorCreateCollectionRequest.md) |  | 

### Return type

[**VectorCreateCollection200Response**](VectorCreateCollection200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VectorDeleteCollection

> VectorCreateCollection200Response VectorDeleteCollection(ctx, collectionName).Execute()

Delete collection

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VectorCollectionsAPI.VectorDeleteCollection(context.Background(), collectionName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorCollectionsAPI.VectorDeleteCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VectorDeleteCollection`: VectorCreateCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `VectorCollectionsAPI.VectorDeleteCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVectorDeleteCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VectorCreateCollection200Response**](VectorCreateCollection200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VectorGetCollection

> VectorGetCollection200Response VectorGetCollection(ctx, collectionName).Execute()

Get collection info

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VectorCollectionsAPI.VectorGetCollection(context.Background(), collectionName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorCollectionsAPI.VectorGetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VectorGetCollection`: VectorGetCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `VectorCollectionsAPI.VectorGetCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVectorGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VectorGetCollection200Response**](VectorGetCollection200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VectorListCollections

> VectorListCollections200Response VectorListCollections(ctx).Execute()

List collections

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
	resp, r, err := apiClient.VectorCollectionsAPI.VectorListCollections(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VectorCollectionsAPI.VectorListCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VectorListCollections`: VectorListCollections200Response
	fmt.Fprintf(os.Stdout, "Response from `VectorCollectionsAPI.VectorListCollections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVectorListCollectionsRequest struct via the builder pattern


### Return type

[**VectorListCollections200Response**](VectorListCollections200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

