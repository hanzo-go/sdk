# \K8sAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudCreateKubernetesCluster**](K8sAPI.md#CloudCreateKubernetesCluster) | **Post** /v1/k8s/clusters | Provisions a DOKS cluster for the caller&#39;s org and answers 201.
[**CloudDeleteKubernetesCluster**](K8sAPI.md#CloudDeleteKubernetesCluster) | **Delete** /v1/k8s/clusters/{id} | Destroys a DOKS cluster by id and answers 204.
[**CloudGetKubernetesCluster**](K8sAPI.md#CloudGetKubernetesCluster) | **Get** /v1/k8s/clusters/{id} | Returns one cluster&#39;s detail: node pools + worker nodes.
[**CloudListKubernetesClusters**](K8sAPI.md#CloudListKubernetesClusters) | **Get** /v1/k8s/clusters | Lists the org&#39;s DOKS clusters (Visor, house account) folded with the org&#39;s BYO clusters — ONE fleet cluster view under the unified k8s noun.
[**CloudListKubernetesNodes**](K8sAPI.md#CloudListKubernetesNodes) | **Get** /v1/k8s/nodes | Returns every DOKS worker node in the org&#39;s clusters as a machine — the SAME set the fleet folds in (managedMachines), exposed directly under the k8s noun.



## CloudCreateKubernetesCluster

> CloudClusterView CloudCreateKubernetesCluster(ctx).CloudCreateClusterReq(cloudCreateClusterReq).Execute()

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
	cloudCreateClusterReq := *openapiclient.NewCloudCreateClusterReq() // CloudCreateClusterReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.K8sAPI.CloudCreateKubernetesCluster(context.Background()).CloudCreateClusterReq(cloudCreateClusterReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.CloudCreateKubernetesCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudCreateKubernetesCluster`: CloudClusterView
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.CloudCreateKubernetesCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudCreateKubernetesClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCreateClusterReq** | [**CloudCreateClusterReq**](CloudCreateClusterReq.md) |  | 

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


## CloudDeleteKubernetesCluster

> CloudDeleteKubernetesCluster(ctx, id).Execute()

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
	r, err := apiClient.K8sAPI.CloudDeleteKubernetesCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.CloudDeleteKubernetesCluster``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteKubernetesClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CloudGetKubernetesCluster

> CloudClusterDetailView CloudGetKubernetesCluster(ctx, id).Execute()

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
	resp, r, err := apiClient.K8sAPI.CloudGetKubernetesCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.CloudGetKubernetesCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetKubernetesCluster`: CloudClusterDetailView
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.CloudGetKubernetesCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the provider&#39;s DOKS cluster id. Visor scopes the lookup to the caller&#39;s org, so another tenant&#39;s id resolves to not-found rather than their cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetKubernetesClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudClusterDetailView**](CloudClusterDetailView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudListKubernetesClusters

> CloudClusterList CloudListKubernetesClusters(ctx).Execute()

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
	resp, r, err := apiClient.K8sAPI.CloudListKubernetesClusters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.CloudListKubernetesClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudListKubernetesClusters`: CloudClusterList
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.CloudListKubernetesClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudListKubernetesClustersRequest struct via the builder pattern


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


## CloudListKubernetesNodes

> CloudNodeList CloudListKubernetesNodes(ctx).Execute()

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
	resp, r, err := apiClient.K8sAPI.CloudListKubernetesNodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `K8sAPI.CloudListKubernetesNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudListKubernetesNodes`: CloudNodeList
	fmt.Fprintf(os.Stdout, "Response from `K8sAPI.CloudListKubernetesNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudListKubernetesNodesRequest struct via the builder pattern


### Return type

[**CloudNodeList**](CloudNodeList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

