# ComputeClusterDetailView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmdGpu** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DoClusterId** | Pointer to **string** |  | [optional] 
**DoksClusterId** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int64** |  | [optional] 
**NodePools** | Pointer to [**[]ComputeNodePoolView**](ComputeNodePoolView.md) |  | [optional] 
**NodeSize** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to [**[]ComputeMachineView**](ComputeMachineView.md) | Nodes is every worker node in the cluster, each in the same shape the machines surface uses — a node IS a machine, addressable by its own id. This is the individual hardware behind the pool counts above. | [optional] 
**NvidiaGpu** | Pointer to **int64** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewComputeClusterDetailView

`func NewComputeClusterDetailView() *ComputeClusterDetailView`

NewComputeClusterDetailView instantiates a new ComputeClusterDetailView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeClusterDetailViewWithDefaults

`func NewComputeClusterDetailViewWithDefaults() *ComputeClusterDetailView`

NewComputeClusterDetailViewWithDefaults instantiates a new ComputeClusterDetailView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmdGpu

`func (o *ComputeClusterDetailView) GetAmdGpu() int64`

GetAmdGpu returns the AmdGpu field if non-nil, zero value otherwise.

### GetAmdGpuOk

`func (o *ComputeClusterDetailView) GetAmdGpuOk() (*int64, bool)`

GetAmdGpuOk returns a tuple with the AmdGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmdGpu

`func (o *ComputeClusterDetailView) SetAmdGpu(v int64)`

SetAmdGpu sets AmdGpu field to given value.

### HasAmdGpu

`func (o *ComputeClusterDetailView) HasAmdGpu() bool`

HasAmdGpu returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ComputeClusterDetailView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ComputeClusterDetailView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ComputeClusterDetailView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ComputeClusterDetailView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDoClusterId

`func (o *ComputeClusterDetailView) GetDoClusterId() string`

GetDoClusterId returns the DoClusterId field if non-nil, zero value otherwise.

### GetDoClusterIdOk

`func (o *ComputeClusterDetailView) GetDoClusterIdOk() (*string, bool)`

GetDoClusterIdOk returns a tuple with the DoClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoClusterId

`func (o *ComputeClusterDetailView) SetDoClusterId(v string)`

SetDoClusterId sets DoClusterId field to given value.

### HasDoClusterId

`func (o *ComputeClusterDetailView) HasDoClusterId() bool`

HasDoClusterId returns a boolean if a field has been set.

### GetDoksClusterId

`func (o *ComputeClusterDetailView) GetDoksClusterId() string`

GetDoksClusterId returns the DoksClusterId field if non-nil, zero value otherwise.

### GetDoksClusterIdOk

`func (o *ComputeClusterDetailView) GetDoksClusterIdOk() (*string, bool)`

GetDoksClusterIdOk returns a tuple with the DoksClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoksClusterId

`func (o *ComputeClusterDetailView) SetDoksClusterId(v string)`

SetDoksClusterId sets DoksClusterId field to given value.

### HasDoksClusterId

`func (o *ComputeClusterDetailView) HasDoksClusterId() bool`

HasDoksClusterId returns a boolean if a field has been set.

### GetKind

`func (o *ComputeClusterDetailView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ComputeClusterDetailView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ComputeClusterDetailView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ComputeClusterDetailView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ComputeClusterDetailView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputeClusterDetailView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputeClusterDetailView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputeClusterDetailView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeCount

`func (o *ComputeClusterDetailView) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ComputeClusterDetailView) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ComputeClusterDetailView) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ComputeClusterDetailView) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodePools

`func (o *ComputeClusterDetailView) GetNodePools() []ComputeNodePoolView`

GetNodePools returns the NodePools field if non-nil, zero value otherwise.

### GetNodePoolsOk

`func (o *ComputeClusterDetailView) GetNodePoolsOk() (*[]ComputeNodePoolView, bool)`

GetNodePoolsOk returns a tuple with the NodePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePools

`func (o *ComputeClusterDetailView) SetNodePools(v []ComputeNodePoolView)`

SetNodePools sets NodePools field to given value.

### HasNodePools

`func (o *ComputeClusterDetailView) HasNodePools() bool`

HasNodePools returns a boolean if a field has been set.

### GetNodeSize

`func (o *ComputeClusterDetailView) GetNodeSize() string`

GetNodeSize returns the NodeSize field if non-nil, zero value otherwise.

### GetNodeSizeOk

`func (o *ComputeClusterDetailView) GetNodeSizeOk() (*string, bool)`

GetNodeSizeOk returns a tuple with the NodeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSize

`func (o *ComputeClusterDetailView) SetNodeSize(v string)`

SetNodeSize sets NodeSize field to given value.

### HasNodeSize

`func (o *ComputeClusterDetailView) HasNodeSize() bool`

HasNodeSize returns a boolean if a field has been set.

### GetNodes

`func (o *ComputeClusterDetailView) GetNodes() []ComputeMachineView`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ComputeClusterDetailView) GetNodesOk() (*[]ComputeMachineView, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ComputeClusterDetailView) SetNodes(v []ComputeMachineView)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ComputeClusterDetailView) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetNvidiaGpu

`func (o *ComputeClusterDetailView) GetNvidiaGpu() int64`

GetNvidiaGpu returns the NvidiaGpu field if non-nil, zero value otherwise.

### GetNvidiaGpuOk

`func (o *ComputeClusterDetailView) GetNvidiaGpuOk() (*int64, bool)`

GetNvidiaGpuOk returns a tuple with the NvidiaGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvidiaGpu

`func (o *ComputeClusterDetailView) SetNvidiaGpu(v int64)`

SetNvidiaGpu sets NvidiaGpu field to given value.

### HasNvidiaGpu

`func (o *ComputeClusterDetailView) HasNvidiaGpu() bool`

HasNvidiaGpu returns a boolean if a field has been set.

### GetRegion

`func (o *ComputeClusterDetailView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ComputeClusterDetailView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ComputeClusterDetailView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ComputeClusterDetailView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *ComputeClusterDetailView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComputeClusterDetailView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComputeClusterDetailView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComputeClusterDetailView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


