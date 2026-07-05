# \VisorClustersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VisorCreatePool**](VisorClustersAPI.md#VisorCreatePool) | **Post** /v1/clusters/{clusterId}/pools | Add a node pool to a cluster
[**VisorDeletePool**](VisorClustersAPI.md#VisorDeletePool) | **Delete** /v1/clusters/{clusterId}/pools/{poolId} | Delete a node pool
[**VisorListClusters**](VisorClustersAPI.md#VisorListClusters) | **Get** /v1/clusters | List DOKS clusters (projected from node pools)
[**VisorScalePool**](VisorClustersAPI.md#VisorScalePool) | **Post** /v1/clusters/{clusterId}/pools/{poolId}/scale | Scale a node pool



## VisorCreatePool

> VisorNodePoolView VisorCreatePool(ctx, clusterId).VisorPoolRequest(visorPoolRequest).Execute()

Add a node pool to a cluster

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
	clusterId := "clusterId_example" // string | 
	visorPoolRequest := *openapiclient.NewVisorPoolRequest("Provider_example") // VisorPoolRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VisorClustersAPI.VisorCreatePool(context.Background(), clusterId).VisorPoolRequest(visorPoolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VisorClustersAPI.VisorCreatePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VisorCreatePool`: VisorNodePoolView
	fmt.Fprintf(os.Stdout, "Response from `VisorClustersAPI.VisorCreatePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVisorCreatePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **visorPoolRequest** | [**VisorPoolRequest**](VisorPoolRequest.md) |  | 

### Return type

[**VisorNodePoolView**](VisorNodePoolView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VisorDeletePool

> VisorDeletePool(ctx, clusterId, poolId).Provider(provider).Execute()

Delete a node pool

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
	clusterId := "clusterId_example" // string | 
	poolId := "poolId_example" // string | 
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VisorClustersAPI.VisorDeletePool(context.Background(), clusterId, poolId).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VisorClustersAPI.VisorDeletePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**poolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVisorDeletePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **provider** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VisorListClusters

> VisorListClusters200Response VisorListClusters(ctx).Execute()

List DOKS clusters (projected from node pools)

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
	resp, r, err := apiClient.VisorClustersAPI.VisorListClusters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VisorClustersAPI.VisorListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VisorListClusters`: VisorListClusters200Response
	fmt.Fprintf(os.Stdout, "Response from `VisorClustersAPI.VisorListClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVisorListClustersRequest struct via the builder pattern


### Return type

[**VisorListClusters200Response**](VisorListClusters200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VisorScalePool

> VisorNodePoolView VisorScalePool(ctx, clusterId, poolId).VisorScaleRequest(visorScaleRequest).Execute()

Scale a node pool

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
	clusterId := "clusterId_example" // string | 
	poolId := "poolId_example" // string | 
	visorScaleRequest := *openapiclient.NewVisorScaleRequest("Provider_example") // VisorScaleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VisorClustersAPI.VisorScalePool(context.Background(), clusterId, poolId).VisorScaleRequest(visorScaleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VisorClustersAPI.VisorScalePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VisorScalePool`: VisorNodePoolView
	fmt.Fprintf(os.Stdout, "Response from `VisorClustersAPI.VisorScalePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**poolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVisorScalePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **visorScaleRequest** | [**VisorScaleRequest**](VisorScaleRequest.md) |  | 

### Return type

[**VisorNodePoolView**](VisorNodePoolView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

