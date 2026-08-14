# \ClustersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachCluster**](ClustersAPI.md#AttachCluster) | **Post** /v1/clusters | Attaches a BYO cluster to the caller&#39;s org — the kubeconfig is validated, KMS-sealed and added to the fleet — and answers 201 with the cluster as it now appears on GET /v1/clusters.
[**CreateNodePool**](ClustersAPI.md#CreateNodePool) | **Post** /v1/clusters/{clusterId}/pools | Adds a node pool to one of the caller org&#39;s clusters and answers 201 with the created pool.
[**DeleteNodePool**](ClustersAPI.md#DeleteNodePool) | **Delete** /v1/clusters/{clusterId}/pools/{poolId} | Removes a node pool from one of the caller org&#39;s clusters.
[**DetachCluster**](ClustersAPI.md#DetachCluster) | **Delete** /v1/clusters/{id} | Removes a BYO cluster from the caller org&#39;s fleet.
[**ListClusters**](ClustersAPI.md#ListClusters) | **Get** /v1/clusters | Returns the caller org&#39;s clusters from both sources: the managed clusters projected from Visor&#39;s node pools, and the BYO clusters attached to the caller&#39;s project.
[**ScaleNodePool**](ClustersAPI.md#ScaleNodePool) | **Post** /v1/clusters/{clusterId}/pools/{poolId}/scale | Resizes a node pool to an absolute node count and returns the pool as Visor reports it after the change.



## AttachCluster

> ClusterView AttachCluster(ctx).ClusterAttach(clusterAttach).Execute()

Attaches a BYO cluster to the caller's org — the kubeconfig is validated, KMS-sealed and added to the fleet — and answers 201 with the cluster as it now appears on GET /v1/clusters.



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
	clusterAttach := *openapiclient.NewClusterAttach() // ClusterAttach | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.AttachCluster(context.Background()).ClusterAttach(clusterAttach).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.AttachCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachCluster`: ClusterView
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.AttachCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterAttach** | [**ClusterAttach**](ClusterAttach.md) |  | 

### Return type

[**ClusterView**](ClusterView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNodePool

> NodePoolView CreateNodePool(ctx, clusterId).PoolCreate(poolCreate).Execute()

Adds a node pool to one of the caller org's clusters and answers 201 with the created pool.



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
	clusterId := "clusterId_example" // string | ClusterID is the cluster to add the pool to, from the URL path.
	poolCreate := *openapiclient.NewPoolCreate() // PoolCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.CreateNodePool(context.Background(), clusterId).PoolCreate(poolCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.CreateNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNodePool`: NodePoolView
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.CreateNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | ClusterID is the cluster to add the pool to, from the URL path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **poolCreate** | [**PoolCreate**](PoolCreate.md) |  | 

### Return type

[**NodePoolView**](NodePoolView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNodePool

> DeleteNodePool(ctx, clusterId, poolId).Provider(provider).Execute()

Removes a node pool from one of the caller org's clusters.



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
	clusterId := "clusterId_example" // string | ClusterID and PoolID address the pool, from the URL path.
	poolId := "poolId_example" // string | 
	provider := "provider_example" // string | Provider is the cloud the cluster lives on, from ?provider=. Required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClustersAPI.DeleteNodePool(context.Background(), clusterId, poolId).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.DeleteNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | ClusterID and PoolID address the pool, from the URL path. | 
**poolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **provider** | **string** | Provider is the cloud the cluster lives on, from ?provider&#x3D;. Required. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachCluster

> ClusterDetached DetachCluster(ctx, id).Execute()

Removes a BYO cluster from the caller org's fleet.



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
	id := "id_example" // string | ID is the cluster's fleet name (the `name` it was attached under), matched lower-cased.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.DetachCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.DetachCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachCluster`: ClusterDetached
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.DetachCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the cluster&#39;s fleet name (the &#x60;name&#x60; it was attached under), matched lower-cased. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterDetached**](ClusterDetached.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusters

> ClusterList ListClusters(ctx).Execute()

Returns the caller org's clusters from both sources: the managed clusters projected from Visor's node pools, and the BYO clusters attached to the caller's project.



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
	resp, r, err := apiClient.ClustersAPI.ListClusters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.ListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusters`: ClusterList
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.ListClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListClustersRequest struct via the builder pattern


### Return type

[**ClusterList**](ClusterList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScaleNodePool

> NodePoolView ScaleNodePool(ctx, clusterId, poolId).PoolScale(poolScale).Execute()

Resizes a node pool to an absolute node count and returns the pool as Visor reports it after the change.



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
	clusterId := "clusterId_example" // string | ClusterID and PoolID address the pool, from the URL path.
	poolId := "poolId_example" // string | 
	poolScale := *openapiclient.NewPoolScale() // PoolScale | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.ScaleNodePool(context.Background(), clusterId, poolId).PoolScale(poolScale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.ScaleNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScaleNodePool`: NodePoolView
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.ScaleNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | ClusterID and PoolID address the pool, from the URL path. | 
**poolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiScaleNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **poolScale** | [**PoolScale**](PoolScale.md) |  | 

### Return type

[**NodePoolView**](NodePoolView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

