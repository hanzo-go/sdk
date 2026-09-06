# \ComputeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachCluster**](ComputeAPI.md#AttachCluster) | **Post** /v1/compute/clusters | Attaches a BYO cluster to the caller&#39;s org — the kubeconfig is validated, KMS-sealed and added to the fleet — and answers 201 with the cluster as it now appears on GET /v1/compute/clusters.
[**BindMachineAgent**](ComputeAPI.md#BindMachineAgent) | **Put** /v1/compute/machines/{id}/agent | Binds a cloud Agent to one of the caller org&#39;s machines: the machine is recorded as running that Agent&#39;s @hanzo/bot runtime.
[**CancelFleetJob**](ComputeAPI.md#CancelFleetJob) | **Post** /v1/compute/fleet/jobs/{id}/cancel | Cancels a queued or running render in the caller&#39;s org.
[**CreateKubernetesCluster**](ComputeAPI.md#CreateKubernetesCluster) | **Post** /v1/compute/k8s/clusters | Provisions a DOKS cluster for the caller&#39;s org and answers 201.
[**CreateNodePool**](ComputeAPI.md#CreateNodePool) | **Post** /v1/compute/clusters/{clusterId}/pools | Adds a node pool to one of the caller org&#39;s clusters and answers 201 with the created pool.
[**DeleteKubernetesCluster**](ComputeAPI.md#DeleteKubernetesCluster) | **Delete** /v1/compute/k8s/clusters/{id} | Destroys a DOKS cluster by id and answers 204.
[**DeleteMachine**](ComputeAPI.md#DeleteMachine) | **Delete** /v1/compute/machines/{id} | Terminates one of the caller org&#39;s machines.
[**DeleteNodePool**](ComputeAPI.md#DeleteNodePool) | **Delete** /v1/compute/clusters/{clusterId}/pools/{poolId} | Removes a node pool from one of the caller org&#39;s clusters.
[**DetachCluster**](ComputeAPI.md#DetachCluster) | **Delete** /v1/compute/clusters/{id} | Removes a BYO cluster from the caller org&#39;s fleet.
[**GetComputeRegions**](ComputeAPI.md#GetComputeRegions) | **Get** /v1/compute/regions | Regions lists the regions a machine can be launched in.
[**GetComputeSizes**](ComputeAPI.md#GetComputeSizes) | **Get** /v1/compute/sizes | Sizes lists the machine sizes available to launch, with their specifications.
[**GetKubernetesCluster**](ComputeAPI.md#GetKubernetesCluster) | **Get** /v1/compute/k8s/clusters/{id} | Returns one cluster&#39;s detail: node pools + worker nodes.
[**GetMachine**](ComputeAPI.md#GetMachine) | **Get** /v1/compute/machines/{id} | Returns one of the caller org&#39;s machines by its org-scoped name.
[**GetMachineAgent**](ComputeAPI.md#GetMachineAgent) | **Get** /v1/compute/machines/{id}/agent | Returns the agent binding of one of the caller org&#39;s machines, or 404 when the machine runs no bot runtime.
[**ListClusters**](ComputeAPI.md#ListClusters) | **Get** /v1/compute/clusters | Returns the caller org&#39;s clusters from both sources: the managed clusters projected from Visor&#39;s node pools, and the BYO clusters attached to the caller&#39;s project.
[**ListFleet**](ComputeAPI.md#ListFleet) | **Get** /v1/compute/fleet | Returns every compute unit the caller&#39;s org has, from every source, each carrying its latest utilization: agent run-targets, the BYO machines that dialed in, attached BYO clusters and Visor-provisioned machines.
[**ListFleetJobs**](ComputeAPI.md#ListFleetJobs) | **Get** /v1/compute/fleet/jobs | Returns the caller org&#39;s gpu-jobs render queue, each row tagged with the GPU it targets (empty &#x3D; the shared any-GPU lane) and the node claiming it, optionally narrowed to one GPU&#39;s queue and/or one status.
[**ListFleetSamples**](ComputeAPI.md#ListFleetSamples) | **Get** /v1/compute/fleet/samples | Returns the caller org&#39;s utilization series, oldest first.
[**ListFleetWorkers**](ComputeAPI.md#ListFleetWorkers) | **Get** /v1/compute/fleet/workers | Returns the caller org&#39;s BYO machines — the ones that dialed in via &#x60;hanzo link&#x60; — with everything each host reported about itself.
[**ListGpuAlerts**](ComputeAPI.md#ListGpuAlerts) | **Get** /v1/compute/gpus/alerts | Is an HONEST empty surface: Visor exposes no GPU alert inventory, so this returns [] rather than fabricating alerts.
[**ListGpus**](ComputeAPI.md#ListGpus) | **Get** /v1/compute/gpus | Returns one row per physical accelerator the caller&#39;s org has, derived from its real GPU machines (the size slug says how many cards a node holds) and from the accelerators BYO workers report through nvidia-smi.
[**ListKubernetesClusters**](ComputeAPI.md#ListKubernetesClusters) | **Get** /v1/compute/k8s/clusters | Lists the org&#39;s DOKS clusters (Visor, house account) folded with the org&#39;s BYO clusters — ONE fleet cluster view under the unified k8s noun.
[**ListKubernetesNodes**](ComputeAPI.md#ListKubernetesNodes) | **Get** /v1/compute/k8s/nodes | Returns every DOKS worker node in the org&#39;s clusters as a machine — the SAME set the fleet folds in (managedMachines), exposed directly under the k8s noun.
[**ListMachineAgents**](ComputeAPI.md#ListMachineAgents) | **Get** /v1/compute/machines/agents | Returns every agent↔machine binding in the caller&#39;s org — which machines are running which cloud Agent, with vm&#39;s own reconciled status.
[**ListMachines**](ComputeAPI.md#ListMachines) | **Get** /v1/compute/machines | Returns every machine the caller&#39;s org has — Visor&#39;s registry, the live DigitalOcean droplets and the DOKS worker nodes (deduped into one union), plus the BYO machines that dialed in via &#x60;hanzo link&#x60; (provider \&quot;byo\&quot;).
[**PostComputeMachines**](ComputeAPI.md#PostComputeMachines) | **Post** /v1/compute/machines | Launch a metered machine for your org, or price one first with dryRun
[**PostComputeMachinesByIdByAction**](ComputeAPI.md#PostComputeMachinesByIdByAction) | **Post** /v1/compute/machines/{id}/{action} | Message a bot, or stop it, by naming the action in the path
[**RecordFleetSample**](ComputeAPI.md#RecordFleetSample) | **Post** /v1/compute/fleet/samples | Records a BYO worker&#39;s live GPU utilization into the SAME series the fleet board overlays.
[**ScaleNodePool**](ComputeAPI.md#ScaleNodePool) | **Post** /v1/compute/clusters/{clusterId}/pools/{poolId}/scale | Resizes a node pool to an absolute node count and returns the pool as Visor reports it after the change.
[**UnbindMachineAgent**](ComputeAPI.md#UnbindMachineAgent) | **Delete** /v1/compute/machines/{id}/agent | Detaches the agent runtime from one of the caller org&#39;s machines.



## AttachCluster

> ClusterView AttachCluster(ctx).ClusterAttach(clusterAttach).Execute()

Attaches a BYO cluster to the caller's org — the kubeconfig is validated, KMS-sealed and added to the fleet — and answers 201 with the cluster as it now appears on GET /v1/compute/clusters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	clusterAttach := *openapiclient.NewClusterAttach() // ClusterAttach | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.AttachCluster(context.Background()).ClusterAttach(clusterAttach).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.AttachCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachCluster`: ClusterView
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.AttachCluster`: %v\n", resp)
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BindMachineAgent

> AgentBinding BindMachineAgent(ctx, id).BindAgentReq(bindAgentReq).Execute()

Binds a cloud Agent to one of the caller org's machines: the machine is recorded as running that Agent's @hanzo/bot runtime.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID is the machine to bind, from the URL path.
	bindAgentReq := *openapiclient.NewBindAgentReq() // BindAgentReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.BindMachineAgent(context.Background(), id).BindAgentReq(bindAgentReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.BindMachineAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BindMachineAgent`: AgentBinding
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.BindMachineAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the machine to bind, from the URL path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindMachineAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bindAgentReq** | [**BindAgentReq**](BindAgentReq.md) |  | 

### Return type

[**AgentBinding**](AgentBinding.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelFleetJob

> JobCanceled CancelFleetJob(ctx, id).JobCancel(jobCancel).Execute()

Cancels a queued or running render in the caller's org.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID is the job (activity) id, from the URL path.
	jobCancel := *openapiclient.NewJobCancel() // JobCancel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.CancelFleetJob(context.Background(), id).JobCancel(jobCancel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.CancelFleetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelFleetJob`: JobCanceled
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.CancelFleetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the job (activity) id, from the URL path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelFleetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobCancel** | [**JobCancel**](JobCancel.md) |  | 

### Return type

[**JobCanceled**](JobCanceled.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	createClusterReq := *openapiclient.NewCreateClusterReq() // CreateClusterReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.CreateKubernetesCluster(context.Background()).CreateClusterReq(createClusterReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.CreateKubernetesCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKubernetesCluster`: ClusterView
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.CreateKubernetesCluster`: %v\n", resp)
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	clusterId := "clusterId_example" // string | ClusterID is the cluster to add the pool to, from the URL path.
	poolCreate := *openapiclient.NewPoolCreate() // PoolCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.CreateNodePool(context.Background(), clusterId).PoolCreate(poolCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.CreateNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNodePool`: NodePoolView
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.CreateNodePool`: %v\n", resp)
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID is the provider's DOKS cluster id. Visor scopes the lookup to the caller's org, so another tenant's id resolves to not-found rather than their cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputeAPI.DeleteKubernetesCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.DeleteKubernetesCluster``: %v\n", err)
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMachine

> DeleteMachine(ctx, id).Execute()

Terminates one of the caller org's machines.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID is the machine's org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputeAPI.DeleteMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.DeleteMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the machine&#39;s org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	clusterId := "clusterId_example" // string | ClusterID and PoolID address the pool, from the URL path.
	poolId := "poolId_example" // string | 
	provider := "provider_example" // string | Provider is the cloud the cluster lives on, from ?provider=. Required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputeAPI.DeleteNodePool(context.Background(), clusterId, poolId).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.DeleteNodePool``: %v\n", err)
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID is the cluster's fleet name (the `name` it was attached under), matched lower-cased.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.DetachCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.DetachCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachCluster`: ClusterDetached
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.DetachCluster`: %v\n", resp)
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComputeRegions

> interface{} GetComputeRegions(ctx).Execute()

Regions lists the regions a machine can be launched in.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.GetComputeRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.GetComputeRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComputeRegions`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.GetComputeRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetComputeRegionsRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComputeSizes

> interface{} GetComputeSizes(ctx).Execute()

Sizes lists the machine sizes available to launch, with their specifications.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.GetComputeSizes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.GetComputeSizes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComputeSizes`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.GetComputeSizes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetComputeSizesRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID is the provider's DOKS cluster id. Visor scopes the lookup to the caller's org, so another tenant's id resolves to not-found rather than their cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.GetKubernetesCluster(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.GetKubernetesCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKubernetesCluster`: ClusterDetailView
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.GetKubernetesCluster`: %v\n", resp)
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMachine

> MachineView GetMachine(ctx, id).Execute()

Returns one of the caller org's machines by its org-scoped name.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID is the machine's org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.GetMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.GetMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMachine`: MachineView
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.GetMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the machine&#39;s org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MachineView**](MachineView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMachineAgent

> AgentBinding GetMachineAgent(ctx, id).Execute()

Returns the agent binding of one of the caller org's machines, or 404 when the machine runs no bot runtime.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID is the machine's org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.GetMachineAgent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.GetMachineAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMachineAgent`: AgentBinding
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.GetMachineAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the machine&#39;s org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMachineAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentBinding**](AgentBinding.md)

### Authorization

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.ListClusters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.ListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusters`: ClusterList
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.ListClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListClustersRequest struct via the builder pattern


### Return type

[**ClusterList**](ClusterList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFleet

> FleetBoard ListFleet(ctx).Execute()

Returns every compute unit the caller's org has, from every source, each carrying its latest utilization: agent run-targets, the BYO machines that dialed in, attached BYO clusters and Visor-provisioned machines.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.ListFleet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.ListFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFleet`: FleetBoard
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.ListFleet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFleetRequest struct via the builder pattern


### Return type

[**FleetBoard**](FleetBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFleetJobs

> JobList ListFleetJobs(ctx).Gpu(gpu).Status(status).Execute()

Returns the caller org's gpu-jobs render queue, each row tagged with the GPU it targets (empty = the shared any-GPU lane) and the node claiming it, optionally narrowed to one GPU's queue and/or one status.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	gpu := "gpu_example" // string | GPU selects one node's lane: jobs TARGETED at it (gpu:<node>) or CLAIMED by it. The literal \"shared\" selects the any-GPU lane — no target, no claimant. Matched case-insensitively. (optional)
	status := "status_example" // string | Status selects one lifecycle state: queued, running, stalled, completed, failed or canceled. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.ListFleetJobs(context.Background()).Gpu(gpu).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.ListFleetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFleetJobs`: JobList
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.ListFleetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFleetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gpu** | **string** | GPU selects one node&#39;s lane: jobs TARGETED at it (gpu:&lt;node&gt;) or CLAIMED by it. The literal \&quot;shared\&quot; selects the any-GPU lane — no target, no claimant. Matched case-insensitively. | 
 **status** | **string** | Status selects one lifecycle state: queued, running, stalled, completed, failed or canceled. | 

### Return type

[**JobList**](JobList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFleetSamples

> SampleList ListFleetSamples(ctx).Unit(unit).Source(source).Range_(range_).Execute()

Returns the caller org's utilization series, oldest first.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	unit := "unit_example" // string | Unit selects one compute unit's series by its source-local id. (optional)
	source := "source_example" // string | Source selects one plane: \"agent\", \"byo\" or \"visor\". (optional)
	range_ := "range__example" // string | Range is the lookback window (e.g. \"1h\", \"24h\", \"7d\"); empty takes the warehouse default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.ListFleetSamples(context.Background()).Unit(unit).Source(source).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.ListFleetSamples``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFleetSamples`: SampleList
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.ListFleetSamples`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFleetSamplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unit** | **string** | Unit selects one compute unit&#39;s series by its source-local id. | 
 **source** | **string** | Source selects one plane: \&quot;agent\&quot;, \&quot;byo\&quot; or \&quot;visor\&quot;. | 
 **range_** | **string** | Range is the lookback window (e.g. \&quot;1h\&quot;, \&quot;24h\&quot;, \&quot;7d\&quot;); empty takes the warehouse default. | 

### Return type

[**SampleList**](SampleList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFleetWorkers

> WorkerList ListFleetWorkers(ctx).Execute()

Returns the caller org's BYO machines — the ones that dialed in via `hanzo link` — with everything each host reported about itself.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.ListFleetWorkers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.ListFleetWorkers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFleetWorkers`: WorkerList
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.ListFleetWorkers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFleetWorkersRequest struct via the builder pattern


### Return type

[**WorkerList**](WorkerList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGpuAlerts

> GpuAlertList ListGpuAlerts(ctx).Execute()

Is an HONEST empty surface: Visor exposes no GPU alert inventory, so this returns [] rather than fabricating alerts.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.ListGpuAlerts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.ListGpuAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGpuAlerts`: GpuAlertList
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.ListGpuAlerts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGpuAlertsRequest struct via the builder pattern


### Return type

[**GpuAlertList**](GpuAlertList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGpus

> GpuList ListGpus(ctx).Execute()

Returns one row per physical accelerator the caller's org has, derived from its real GPU machines (the size slug says how many cards a node holds) and from the accelerators BYO workers report through nvidia-smi.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.ListGpus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.ListGpus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGpus`: GpuList
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.ListGpus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGpusRequest struct via the builder pattern


### Return type

[**GpuList**](GpuList.md)

### Authorization

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.ListKubernetesClusters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.ListKubernetesClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKubernetesClusters`: ClusterList
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.ListKubernetesClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKubernetesClustersRequest struct via the builder pattern


### Return type

[**ClusterList**](ClusterList.md)

### Authorization

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.ListKubernetesNodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.ListKubernetesNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKubernetesNodes`: NodeList
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.ListKubernetesNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKubernetesNodesRequest struct via the builder pattern


### Return type

[**NodeList**](NodeList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMachineAgents

> BindingList ListMachineAgents(ctx).Execute()

Returns every agent↔machine binding in the caller's org — which machines are running which cloud Agent, with vm's own reconciled status.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.ListMachineAgents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.ListMachineAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMachineAgents`: BindingList
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.ListMachineAgents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMachineAgentsRequest struct via the builder pattern


### Return type

[**BindingList**](BindingList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMachines

> MachineList ListMachines(ctx).Kind(kind).Execute()

Returns every machine the caller's org has — Visor's registry, the live DigitalOcean droplets and the DOKS worker nodes (deduped into one union), plus the BYO machines that dialed in via `hanzo link` (provider \"byo\").



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	kind := "kind_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.ListMachines(context.Background()).Kind(kind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.ListMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMachines`: MachineList
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.ListMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** |  | 

### Return type

[**MachineList**](MachineList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostComputeMachines

> PostComputeMachines(ctx).Execute()

Launch a metered machine for your org, or price one first with dryRun



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputeAPI.PostComputeMachines(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.PostComputeMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostComputeMachinesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostComputeMachinesByIdByAction

> PostComputeMachinesByIdByAction(ctx, id, action).Execute()

Message a bot, or stop it, by naming the action in the path



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | 
	action := "action_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputeAPI.PostComputeMachinesByIdByAction(context.Background(), id, action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.PostComputeMachinesByIdByAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**action** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostComputeMachinesByIdByActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordFleetSample

> SampleAccepted RecordFleetSample(ctx).SampleIngest(sampleIngest).Execute()

Records a BYO worker's live GPU utilization into the SAME series the fleet board overlays.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	sampleIngest := *openapiclient.NewSampleIngest() // SampleIngest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.RecordFleetSample(context.Background()).SampleIngest(sampleIngest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.RecordFleetSample``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordFleetSample`: SampleAccepted
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.RecordFleetSample`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecordFleetSampleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sampleIngest** | [**SampleIngest**](SampleIngest.md) |  | 

### Return type

[**SampleAccepted**](SampleAccepted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	clusterId := "clusterId_example" // string | ClusterID is the cluster holding the pool, from the URL path.
	poolId := "poolId_example" // string | PoolID is the pool to resize, from the URL path — the `poolId` a cluster read reports for it. Required.
	poolScale := *openapiclient.NewPoolScale() // PoolScale | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.ScaleNodePool(context.Background(), clusterId, poolId).PoolScale(poolScale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.ScaleNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScaleNodePool`: NodePoolView
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.ScaleNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | ClusterID is the cluster holding the pool, from the URL path. | 
**poolId** | **string** | PoolID is the pool to resize, from the URL path — the &#x60;poolId&#x60; a cluster read reports for it. Required. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScaleNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **poolScale** | [**PoolScale**](PoolScale.md) |  | 

### Return type

[**NodePoolView**](NodePoolView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnbindMachineAgent

> UnbindMachineAgent(ctx, id).Execute()

Detaches the agent runtime from one of the caller org's machines.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID is the machine's org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputeAPI.UnbindMachineAgent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.UnbindMachineAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the machine&#39;s org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbindMachineAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

