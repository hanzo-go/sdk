# ClusterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmdGpu** | Pointer to **int32** | AmdGPU is how many AMD GPUs those nodes advertise. | [optional] 
**Cluster** | Pointer to **string** | Cluster is the stable fleet name the cluster was folded under — the name /v1/clusters shows and a workload targets. | [optional] 
**Error** | Pointer to **string** | Error is why this cluster did not fold — a billing denial, an unsafe kubeconfig, or an unreachable apiserver. It never contains credential material. | [optional] 
**Folded** | Pointer to **bool** | Folded is whether it reached the fleet. False means Error says why, and this cluster alone was skipped. | [optional] 
**Nodes** | Pointer to **int32** | Nodes is how many nodes the fleet counted in it. | [optional] 
**NvidiaGpu** | Pointer to **int32** | NvidiaGPU is how many NVIDIA GPUs those nodes advertise. | [optional] 
**Region** | Pointer to **string** | Region is the provider region it runs in. | [optional] 
**Source** | Pointer to **string** | Source is the cluster&#39;s own name at the provider. | [optional] 

## Methods

### NewClusterResult

`func NewClusterResult() *ClusterResult`

NewClusterResult instantiates a new ClusterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterResultWithDefaults

`func NewClusterResultWithDefaults() *ClusterResult`

NewClusterResultWithDefaults instantiates a new ClusterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmdGpu

`func (o *ClusterResult) GetAmdGpu() int32`

GetAmdGpu returns the AmdGpu field if non-nil, zero value otherwise.

### GetAmdGpuOk

`func (o *ClusterResult) GetAmdGpuOk() (*int32, bool)`

GetAmdGpuOk returns a tuple with the AmdGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmdGpu

`func (o *ClusterResult) SetAmdGpu(v int32)`

SetAmdGpu sets AmdGpu field to given value.

### HasAmdGpu

`func (o *ClusterResult) HasAmdGpu() bool`

HasAmdGpu returns a boolean if a field has been set.

### GetCluster

`func (o *ClusterResult) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ClusterResult) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ClusterResult) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ClusterResult) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetError

`func (o *ClusterResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ClusterResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ClusterResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ClusterResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFolded

`func (o *ClusterResult) GetFolded() bool`

GetFolded returns the Folded field if non-nil, zero value otherwise.

### GetFoldedOk

`func (o *ClusterResult) GetFoldedOk() (*bool, bool)`

GetFoldedOk returns a tuple with the Folded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolded

`func (o *ClusterResult) SetFolded(v bool)`

SetFolded sets Folded field to given value.

### HasFolded

`func (o *ClusterResult) HasFolded() bool`

HasFolded returns a boolean if a field has been set.

### GetNodes

`func (o *ClusterResult) GetNodes() int32`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ClusterResult) GetNodesOk() (*int32, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ClusterResult) SetNodes(v int32)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ClusterResult) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetNvidiaGpu

`func (o *ClusterResult) GetNvidiaGpu() int32`

GetNvidiaGpu returns the NvidiaGpu field if non-nil, zero value otherwise.

### GetNvidiaGpuOk

`func (o *ClusterResult) GetNvidiaGpuOk() (*int32, bool)`

GetNvidiaGpuOk returns a tuple with the NvidiaGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvidiaGpu

`func (o *ClusterResult) SetNvidiaGpu(v int32)`

SetNvidiaGpu sets NvidiaGpu field to given value.

### HasNvidiaGpu

`func (o *ClusterResult) HasNvidiaGpu() bool`

HasNvidiaGpu returns a boolean if a field has been set.

### GetRegion

`func (o *ClusterResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterResult) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ClusterResult) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSource

`func (o *ClusterResult) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ClusterResult) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ClusterResult) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ClusterResult) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


