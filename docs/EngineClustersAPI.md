# \EngineClustersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EngineDeleteCluster**](EngineClustersAPI.md#EngineDeleteCluster) | **Delete** /v1/engine/clusters/{id} | Deregister cluster
[**EngineGetCluster**](EngineClustersAPI.md#EngineGetCluster) | **Get** /v1/engine/clusters/{id} | Get cluster
[**EngineListClusterNodes**](EngineClustersAPI.md#EngineListClusterNodes) | **Get** /v1/engine/clusters/{id}/nodes | List cluster nodes
[**EngineListClusters**](EngineClustersAPI.md#EngineListClusters) | **Get** /v1/engine/clusters | List GPU clusters
[**EngineRegisterCluster**](EngineClustersAPI.md#EngineRegisterCluster) | **Post** /v1/engine/clusters | Register GPU cluster
[**EngineUpdateCluster**](EngineClustersAPI.md#EngineUpdateCluster) | **Put** /v1/engine/clusters/{id} | Update cluster



## EngineDeleteCluster

> map[string]interface{} EngineDeleteCluster(ctx, id).Execute()

Deregister cluster

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
	resp, r, err := apiClient.EngineClustersAPI.EngineDeleteCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineClustersAPI.EngineDeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineDeleteCluster`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineClustersAPI.EngineDeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineDeleteClusterRequest struct via the builder pattern


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


## EngineGetCluster

> EngineCluster EngineGetCluster(ctx, id).Execute()

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
	resp, r, err := apiClient.EngineClustersAPI.EngineGetCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineClustersAPI.EngineGetCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineGetCluster`: EngineCluster
	fmt.Fprintf(os.Stdout, "Response from `EngineClustersAPI.EngineGetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineGetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EngineCluster**](EngineCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineListClusterNodes

> EngineListClusterNodes200Response EngineListClusterNodes(ctx, id).Execute()

List cluster nodes

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
	resp, r, err := apiClient.EngineClustersAPI.EngineListClusterNodes(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineClustersAPI.EngineListClusterNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineListClusterNodes`: EngineListClusterNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `EngineClustersAPI.EngineListClusterNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineListClusterNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EngineListClusterNodes200Response**](EngineListClusterNodes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineListClusters

> EngineListClusters200Response EngineListClusters(ctx).Status(status).Provider(provider).Page(page).PageSize(pageSize).Execute()

List GPU clusters

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
	provider := "provider_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineClustersAPI.EngineListClusters(context.Background()).Status(status).Provider(provider).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineClustersAPI.EngineListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineListClusters`: EngineListClusters200Response
	fmt.Fprintf(os.Stdout, "Response from `EngineClustersAPI.EngineListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEngineListClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 
 **provider** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]

### Return type

[**EngineListClusters200Response**](EngineListClusters200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineRegisterCluster

> EngineCluster EngineRegisterCluster(ctx).EngineClusterCreate(engineClusterCreate).Execute()

Register GPU cluster

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
	engineClusterCreate := *openapiclient.NewEngineClusterCreate("Name_example", "Provider_example") // EngineClusterCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineClustersAPI.EngineRegisterCluster(context.Background()).EngineClusterCreate(engineClusterCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineClustersAPI.EngineRegisterCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineRegisterCluster`: EngineCluster
	fmt.Fprintf(os.Stdout, "Response from `EngineClustersAPI.EngineRegisterCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEngineRegisterClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engineClusterCreate** | [**EngineClusterCreate**](EngineClusterCreate.md) |  | 

### Return type

[**EngineCluster**](EngineCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineUpdateCluster

> EngineCluster EngineUpdateCluster(ctx, id).EngineUpdateClusterRequest(engineUpdateClusterRequest).Execute()

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
	engineUpdateClusterRequest := *openapiclient.NewEngineUpdateClusterRequest() // EngineUpdateClusterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineClustersAPI.EngineUpdateCluster(context.Background(), id).EngineUpdateClusterRequest(engineUpdateClusterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineClustersAPI.EngineUpdateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineUpdateCluster`: EngineCluster
	fmt.Fprintf(os.Stdout, "Response from `EngineClustersAPI.EngineUpdateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngineUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **engineUpdateClusterRequest** | [**EngineUpdateClusterRequest**](EngineUpdateClusterRequest.md) |  | 

### Return type

[**EngineCluster**](EngineCluster.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

