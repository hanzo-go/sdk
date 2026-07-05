# \KvClustersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KvCreateCluster**](KvClustersAPI.md#KvCreateCluster) | **Post** /v1/kv/clusters | Create KV cluster
[**KvDeleteCluster**](KvClustersAPI.md#KvDeleteCluster) | **Delete** /v1/kv/clusters/{id} | Delete cluster
[**KvGetCluster**](KvClustersAPI.md#KvGetCluster) | **Get** /v1/kv/clusters/{id} | Get cluster
[**KvGetClusterStats**](KvClustersAPI.md#KvGetClusterStats) | **Get** /v1/kv/clusters/{id}/stats | Get cluster stats
[**KvListClusters**](KvClustersAPI.md#KvListClusters) | **Get** /v1/kv/clusters | List KV clusters
[**KvUpdateCluster**](KvClustersAPI.md#KvUpdateCluster) | **Put** /v1/kv/clusters/{id} | Update cluster



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
	resp, r, err := apiClient.KvClustersAPI.KvCreateCluster(context.Background()).KvClusterCreate(kvClusterCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvClustersAPI.KvCreateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvCreateCluster`: KvCluster
	fmt.Fprintf(os.Stdout, "Response from `KvClustersAPI.KvCreateCluster`: %v\n", resp)
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
	resp, r, err := apiClient.KvClustersAPI.KvDeleteCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvClustersAPI.KvDeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvDeleteCluster`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KvClustersAPI.KvDeleteCluster`: %v\n", resp)
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
	resp, r, err := apiClient.KvClustersAPI.KvGetCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvClustersAPI.KvGetCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvGetCluster`: KvCluster
	fmt.Fprintf(os.Stdout, "Response from `KvClustersAPI.KvGetCluster`: %v\n", resp)
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
	resp, r, err := apiClient.KvClustersAPI.KvGetClusterStats(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvClustersAPI.KvGetClusterStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvGetClusterStats`: KvGetClusterStats200Response
	fmt.Fprintf(os.Stdout, "Response from `KvClustersAPI.KvGetClusterStats`: %v\n", resp)
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
	resp, r, err := apiClient.KvClustersAPI.KvListClusters(context.Background()).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvClustersAPI.KvListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvListClusters`: KvListClusters200Response
	fmt.Fprintf(os.Stdout, "Response from `KvClustersAPI.KvListClusters`: %v\n", resp)
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
	resp, r, err := apiClient.KvClustersAPI.KvUpdateCluster(context.Background(), id).KvUpdateClusterRequest(kvUpdateClusterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvClustersAPI.KvUpdateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvUpdateCluster`: KvCluster
	fmt.Fprintf(os.Stdout, "Response from `KvClustersAPI.KvUpdateCluster`: %v\n", resp)
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

