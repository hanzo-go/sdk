# ClusterView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmdGpu** | Pointer to **int64** | AmdGPU is the same count for &#x60;amd.com/gpu&#x60;: AMD accelerators across the BYO cluster&#39;s nodes, as of the attach. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when the cluster started existing: the earliest creation time among its pools for a managed cluster, and for a BYO one the RFC 3339 moment it was attached. Empty when the source states none. | [optional] 
**DoClusterId** | Pointer to **string** | DoClusterID carries the SAME id as DoksClusterID. Both names exist because the console&#39;s Cluster type reads either one; neither is a second identifier. | [optional] 
**DoksClusterId** | Pointer to **string** | DoksClusterID is the provider&#39;s own id for the cluster, and the value the /v1/compute/k8s/clusters/:id routes take. Empty for a BYO cluster: an attached kubeconfig was never provisioned, so there is no provider id to state. | [optional] 
**Kind** | Pointer to **string** | Kind says which of the two kinds of cluster this row is, and there are only two: \&quot;managed\&quot; — Visor provisioned it and Hanzo&#39;s account pays the provider — or \&quot;byo\&quot;, an existing cluster the org attached by kubeconfig. | [optional] 
**Name** | Pointer to **string** | Name is the cluster&#39;s name: the provider&#39;s for a managed cluster, and for a BYO one the lower-cased fleet name it was attached under — which is also how the detach route addresses it. | [optional] 
**NodeCount** | Pointer to **int64** | NodeCount is how many worker nodes the cluster has — the sum over its pools for a managed cluster, and for a BYO one the node count read off the cluster when it was attached. | [optional] 
**NodePools** | Pointer to [**[]NodePoolView**](NodePoolView.md) | NodePools is the authoritative node inventory — every pool, each with its own size and count. It is empty in two cases that are not \&quot;no pools\&quot;: a row from the /v1/compute/k8s/clusters LIST, which is deliberately lightweight and whose :id detail carries them, and a BYO cluster, whose pools were never read. | [optional] 
**NodeSize** | Pointer to **string** | NodeSize is a display convenience: the size slug of the FIRST pool. A cluster mixing sizes has more than one, and NodePools is where they all are. | [optional] 
**NvidiaGpu** | Pointer to **int64** | NvidiaGPU is how many NVIDIA accelerators the cluster&#39;s nodes advertise, the sum of &#x60;nvidia.com/gpu&#x60; allocatable across them. BYO only, and counted ONCE when the cluster was attached — it is an inventory, not live capacity. | [optional] 
**Region** | Pointer to **string** | Region is the provider region slug for a managed cluster. A BYO cluster has no region we can read, so it carries the free-form &#x60;provider&#x60; label the attach named it with (\&quot;gke\&quot;, \&quot;on-prem\&quot;) instead. | [optional] 
**Status** | Pointer to **string** | Status is the cluster&#39;s state: the provider&#39;s own word for a managed cluster (\&quot;running\&quot;, \&quot;provisioning\&quot;), \&quot;unknown\&quot; when the provider stated none, and always \&quot;attached\&quot; for a BYO cluster — that one says the kubeconfig is on file, not that the cluster is reachable this second. | [optional] 

## Methods

### NewClusterView

`func NewClusterView() *ClusterView`

NewClusterView instantiates a new ClusterView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterViewWithDefaults

`func NewClusterViewWithDefaults() *ClusterView`

NewClusterViewWithDefaults instantiates a new ClusterView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmdGpu

`func (o *ClusterView) GetAmdGpu() int64`

GetAmdGpu returns the AmdGpu field if non-nil, zero value otherwise.

### GetAmdGpuOk

`func (o *ClusterView) GetAmdGpuOk() (*int64, bool)`

GetAmdGpuOk returns a tuple with the AmdGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmdGpu

`func (o *ClusterView) SetAmdGpu(v int64)`

SetAmdGpu sets AmdGpu field to given value.

### HasAmdGpu

`func (o *ClusterView) HasAmdGpu() bool`

HasAmdGpu returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ClusterView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ClusterView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDoClusterId

`func (o *ClusterView) GetDoClusterId() string`

GetDoClusterId returns the DoClusterId field if non-nil, zero value otherwise.

### GetDoClusterIdOk

`func (o *ClusterView) GetDoClusterIdOk() (*string, bool)`

GetDoClusterIdOk returns a tuple with the DoClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoClusterId

`func (o *ClusterView) SetDoClusterId(v string)`

SetDoClusterId sets DoClusterId field to given value.

### HasDoClusterId

`func (o *ClusterView) HasDoClusterId() bool`

HasDoClusterId returns a boolean if a field has been set.

### GetDoksClusterId

`func (o *ClusterView) GetDoksClusterId() string`

GetDoksClusterId returns the DoksClusterId field if non-nil, zero value otherwise.

### GetDoksClusterIdOk

`func (o *ClusterView) GetDoksClusterIdOk() (*string, bool)`

GetDoksClusterIdOk returns a tuple with the DoksClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoksClusterId

`func (o *ClusterView) SetDoksClusterId(v string)`

SetDoksClusterId sets DoksClusterId field to given value.

### HasDoksClusterId

`func (o *ClusterView) HasDoksClusterId() bool`

HasDoksClusterId returns a boolean if a field has been set.

### GetKind

`func (o *ClusterView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ClusterView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeCount

`func (o *ClusterView) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterView) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterView) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterView) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodePools

`func (o *ClusterView) GetNodePools() []NodePoolView`

GetNodePools returns the NodePools field if non-nil, zero value otherwise.

### GetNodePoolsOk

`func (o *ClusterView) GetNodePoolsOk() (*[]NodePoolView, bool)`

GetNodePoolsOk returns a tuple with the NodePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePools

`func (o *ClusterView) SetNodePools(v []NodePoolView)`

SetNodePools sets NodePools field to given value.

### HasNodePools

`func (o *ClusterView) HasNodePools() bool`

HasNodePools returns a boolean if a field has been set.

### GetNodeSize

`func (o *ClusterView) GetNodeSize() string`

GetNodeSize returns the NodeSize field if non-nil, zero value otherwise.

### GetNodeSizeOk

`func (o *ClusterView) GetNodeSizeOk() (*string, bool)`

GetNodeSizeOk returns a tuple with the NodeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSize

`func (o *ClusterView) SetNodeSize(v string)`

SetNodeSize sets NodeSize field to given value.

### HasNodeSize

`func (o *ClusterView) HasNodeSize() bool`

HasNodeSize returns a boolean if a field has been set.

### GetNvidiaGpu

`func (o *ClusterView) GetNvidiaGpu() int64`

GetNvidiaGpu returns the NvidiaGpu field if non-nil, zero value otherwise.

### GetNvidiaGpuOk

`func (o *ClusterView) GetNvidiaGpuOk() (*int64, bool)`

GetNvidiaGpuOk returns a tuple with the NvidiaGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvidiaGpu

`func (o *ClusterView) SetNvidiaGpu(v int64)`

SetNvidiaGpu sets NvidiaGpu field to given value.

### HasNvidiaGpu

`func (o *ClusterView) HasNvidiaGpu() bool`

HasNvidiaGpu returns a boolean if a field has been set.

### GetRegion

`func (o *ClusterView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ClusterView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


