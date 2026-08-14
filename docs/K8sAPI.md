# \K8sAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKubernetesCluster**](K8sAPI.md#CreateKubernetesCluster) | **Post** /v1/k8s/clusters | Provisions a DOKS cluster for the caller&#39;s org and answers 201.
[**DeleteKubernetesCluster**](K8sAPI.md#DeleteKubernetesCluster) | **Delete** /v1/k8s/clusters/{id} | Destroys a DOKS cluster by id and answers 204.
[**GetKubernetesCluster**](K8sAPI.md#GetKubernetesCluster) | **Get** /v1/k8s/clusters/{id} | Returns one cluster&#39;s detail: node pools + worker nodes.
[**ListKubernetesClusters**](K8sAPI.md#ListKubernetesClusters) | **Get** /v1/k8s/clusters | Lists the org&#39;s DOKS clusters (Visor, house account) folded with the org&#39;s BYO clusters — ONE fleet cluster view under the unified k8s noun.
[**ListKubernetesNodes**](K8sAPI.md#ListKubernetesNodes) | **Get** /v1/k8s/nodes | Returns every DOKS worker node in the org&#39;s clusters as a machine — the SAME set the fleet folds in (managedMachines), exposed directly under the k8s noun.



## CreateKubernetesCluster

> ClusterView CreateKubernetesCluster(ctx).CreateClusterReq(createClusterReq).Execute()

Provisions a DOKS cluster for the caller's org and answers 201.



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
	createClusterReq := *openapiclient.NewCreateClusterReq() // CreateClusterReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.CreateKubernetesCluster(context.Background()).CreateClusterReq(createClusterReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.CreateKubernetesCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKubernetesCluster`: ClusterView
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.CreateKubernetesCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKubernetesClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createClusterReq** | [**CreateClusterReq**](CreateClusterReq.md) |  | 

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


## DeleteKubernetesCluster

> DeleteKubernetesCluster(ctx, id).Execute()

Destroys a DOKS cluster by id and answers 204.



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
	id := "id_example" // string | ID is the provider's DOKS cluster id. Visor scopes the lookup to the caller's org, so another tenant's id resolves to not-found rather than their cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.K8sAPI.DeleteKubernetesCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.DeleteKubernetesCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the provider&#39;s DOKS cluster id. Visor scopes the lookup to the caller&#39;s org, so another tenant&#39;s id resolves to not-found rather than their cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKubernetesClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetKubernetesCluster

> ClusterDetailView GetKubernetesCluster(ctx, id).Execute()

Returns one cluster's detail: node pools + worker nodes.



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
	id := "id_example" // string | ID is the provider's DOKS cluster id. Visor scopes the lookup to the caller's org, so another tenant's id resolves to not-found rather than their cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.GetKubernetesCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.GetKubernetesCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesCluster`: ClusterDetailView
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.GetKubernetesCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the provider&#39;s DOKS cluster id. Visor scopes the lookup to the caller&#39;s org, so another tenant&#39;s id resolves to not-found rather than their cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterDetailView**](ClusterDetailView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKubernetesClusters

> ClusterList ListKubernetesClusters(ctx).Execute()

Lists the org's DOKS clusters (Visor, house account) folded with the org's BYO clusters — ONE fleet cluster view under the unified k8s noun.



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
	resp, r, err := apiClient.K8sAPI.ListKubernetesClusters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.ListKubernetesClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKubernetesClusters`: ClusterList
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.ListKubernetesClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKubernetesClustersRequest struct via the builder pattern


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


## ListKubernetesNodes

> NodeList ListKubernetesNodes(ctx).Execute()

Returns every DOKS worker node in the org's clusters as a machine — the SAME set the fleet folds in (managedMachines), exposed directly under the k8s noun.



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
	resp, r, err := apiClient.K8sAPI.ListKubernetesNodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.ListKubernetesNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKubernetesNodes`: NodeList
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.ListKubernetesNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKubernetesNodesRequest struct via the builder pattern


### Return type

[**NodeList**](NodeList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

