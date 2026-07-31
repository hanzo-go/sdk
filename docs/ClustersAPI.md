# \ClustersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudAttachCluster**](ClustersAPI.md#CloudAttachCluster) | **Post** /v1/clusters | Attaches a BYO cluster to the caller&#39;s org — the kubeconfig is validated, KMS-sealed and added to the fleet — and answers 201 with the cluster as it now appears on GET /v1/clusters.
[**CloudCreateNodePool**](ClustersAPI.md#CloudCreateNodePool) | **Post** /v1/clusters/{clusterId}/pools | Adds a node pool to one of the caller org&#39;s clusters and answers 201 with the created pool.
[**CloudDeleteNodePool**](ClustersAPI.md#CloudDeleteNodePool) | **Delete** /v1/clusters/{clusterId}/pools/{poolId} | Removes a node pool from one of the caller org&#39;s clusters.
[**CloudDetachCluster**](ClustersAPI.md#CloudDetachCluster) | **Delete** /v1/clusters/{id} | Removes a BYO cluster from the caller org&#39;s fleet.
[**CloudListClusters**](ClustersAPI.md#CloudListClusters) | **Get** /v1/clusters | Returns the caller org&#39;s clusters from both sources: the managed clusters projected from Visor&#39;s node pools, and the BYO clusters attached to the caller&#39;s project.
[**CloudScaleNodePool**](ClustersAPI.md#CloudScaleNodePool) | **Post** /v1/clusters/{clusterId}/pools/{poolId}/scale | Resizes a node pool to an absolute node count and returns the pool as Visor reports it after the change.
[**KvCreateCluster**](ClustersAPI.md#KvCreateCluster) | **Post** /v1/kv/clusters | Create KV cluster
[**KvDeleteCluster**](ClustersAPI.md#KvDeleteCluster) | **Delete** /v1/kv/clusters/{id} | Delete cluster
[**KvGetCluster**](ClustersAPI.md#KvGetCluster) | **Get** /v1/kv/clusters/{id} | Get cluster
[**KvGetClusterStats**](ClustersAPI.md#KvGetClusterStats) | **Get** /v1/kv/clusters/{id}/stats | Get cluster stats
[**KvListClusters**](ClustersAPI.md#KvListClusters) | **Get** /v1/kv/clusters | List KV clusters
[**KvUpdateCluster**](ClustersAPI.md#KvUpdateCluster) | **Put** /v1/kv/clusters/{id} | Update cluster



## CloudAttachCluster

> CloudClusterView CloudAttachCluster(ctx).CloudClusterAttach(cloudClusterAttach).Execute()

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
	cloudClusterAttach := *openapiclient.NewCloudClusterAttach() // CloudClusterAttach | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.CloudAttachCluster(context.Background()).CloudClusterAttach(cloudClusterAttach).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.CloudAttachCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAttachCluster`: CloudClusterView
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.CloudAttachCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAttachClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudClusterAttach** | [**CloudClusterAttach**](CloudClusterAttach.md) |  | 

### Return type

[**CloudClusterView**](CloudClusterView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudCreateNodePool

> CloudNodePoolView CloudCreateNodePool(ctx, clusterId).CloudPoolCreate(cloudPoolCreate).Execute()

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
	cloudPoolCreate := *openapiclient.NewCloudPoolCreate() // CloudPoolCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.CloudCreateNodePool(context.Background(), clusterId).CloudPoolCreate(cloudPoolCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.CloudCreateNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudCreateNodePool`: CloudNodePoolView
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.CloudCreateNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | ClusterID is the cluster to add the pool to, from the URL path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudCreateNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudPoolCreate** | [**CloudPoolCreate**](CloudPoolCreate.md) |  | 

### Return type

[**CloudNodePoolView**](CloudNodePoolView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteNodePool

> CloudDeleteNodePool(ctx, clusterId, poolId).Provider(provider).Execute()

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
	r, err := apiClient.ClustersAPI.CloudDeleteNodePool(context.Background(), clusterId, poolId).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.CloudDeleteNodePool``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **provider** | **string** | Provider is the cloud the cluster lives on, from ?provider&#x3D;. Required. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDetachCluster

> CloudClusterDetached CloudDetachCluster(ctx, id).Execute()

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
	resp, r, err := apiClient.ClustersAPI.CloudDetachCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.CloudDetachCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDetachCluster`: CloudClusterDetached
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.CloudDetachCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the cluster&#39;s fleet name (the &#x60;name&#x60; it was attached under), matched lower-cased. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDetachClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudClusterDetached**](CloudClusterDetached.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudListClusters

> CloudClusterList CloudListClusters(ctx).Execute()

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
	resp, r, err := apiClient.ClustersAPI.CloudListClusters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.CloudListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudListClusters`: CloudClusterList
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.CloudListClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudListClustersRequest struct via the builder pattern


### Return type

[**CloudClusterList**](CloudClusterList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudScaleNodePool

> CloudNodePoolView CloudScaleNodePool(ctx, clusterId, poolId).CloudPoolScale(cloudPoolScale).Execute()

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
	cloudPoolScale := *openapiclient.NewCloudPoolScale() // CloudPoolScale | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.CloudScaleNodePool(context.Background(), clusterId, poolId).CloudPoolScale(cloudPoolScale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.CloudScaleNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudScaleNodePool`: CloudNodePoolView
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.CloudScaleNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | ClusterID and PoolID address the pool, from the URL path. | 
**poolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudScaleNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudPoolScale** | [**CloudPoolScale**](CloudPoolScale.md) |  | 

### Return type

[**CloudNodePoolView**](CloudNodePoolView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvCreateCluster

> KvCluster KvCreateCluster(ctx).KvClusterCreate(kvClusterCreate).Execute()

Create KV cluster

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
	kvClusterCreate := *openapiclient.NewKvClusterCreate("Name_example") // KvClusterCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.KvCreateCluster(context.Background()).KvClusterCreate(kvClusterCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.KvCreateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvCreateCluster`: KvCluster
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.KvCreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKvCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kvClusterCreate** | [**KvClusterCreate**](KvClusterCreate.md) |  | 

### Return type

[**KvCluster**](KvCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvDeleteCluster

> map[string]interface{} KvDeleteCluster(ctx, id).Execute()

Delete cluster

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.KvDeleteCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.KvDeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvDeleteCluster`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.KvDeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvDeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvGetCluster

> KvCluster KvGetCluster(ctx, id).Execute()

Get cluster

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.KvGetCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.KvGetCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvGetCluster`: KvCluster
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.KvGetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvGetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KvCluster**](KvCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvGetClusterStats

> KvGetClusterStats200Response KvGetClusterStats(ctx, id).Execute()

Get cluster stats

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.KvGetClusterStats(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.KvGetClusterStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvGetClusterStats`: KvGetClusterStats200Response
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.KvGetClusterStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvGetClusterStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KvGetClusterStats200Response**](KvGetClusterStats200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvListClusters

> KvListClusters200Response KvListClusters(ctx).Status(status).Execute()

List KV clusters

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
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.KvListClusters(context.Background()).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.KvListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvListClusters`: KvListClusters200Response
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.KvListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKvListClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 

### Return type

[**KvListClusters200Response**](KvListClusters200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvUpdateCluster

> KvCluster KvUpdateCluster(ctx, id).KvUpdateClusterRequest(kvUpdateClusterRequest).Execute()

Update cluster

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kvUpdateClusterRequest := *openapiclient.NewKvUpdateClusterRequest() // KvUpdateClusterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.KvUpdateCluster(context.Background(), id).KvUpdateClusterRequest(kvUpdateClusterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.KvUpdateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvUpdateCluster`: KvCluster
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.KvUpdateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvUpdateClusterRequest** | [**KvUpdateClusterRequest**](KvUpdateClusterRequest.md) |  | 

### Return type

[**KvCluster**](KvCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

