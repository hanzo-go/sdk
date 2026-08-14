# ClusterView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmdGpu** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DoClusterId** | Pointer to **string** |  | [optional] 
**DoksClusterId** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** | Fleet fields (additive): \&quot;managed\&quot; (Visor-provisioned) vs \&quot;byo\&quot; (attached kubeconfig), and the live GPU inventory a BYO cluster reports. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] 
**NodePools** | Pointer to [**[]NodePoolView**](NodePoolView.md) |  | [optional] 
**NodeSize** | Pointer to **string** |  | [optional] 
**NvidiaGpu** | Pointer to **int32** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

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

`func (o *ClusterView) GetAmdGpu() int32`

GetAmdGpu returns the AmdGpu field if non-nil, zero value otherwise.

### GetAmdGpuOk

`func (o *ClusterView) GetAmdGpuOk() (*int32, bool)`

GetAmdGpuOk returns a tuple with the AmdGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmdGpu

`func (o *ClusterView) SetAmdGpu(v int32)`

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

`func (o *ClusterView) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterView) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterView) SetNodeCount(v int32)`

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

`func (o *ClusterView) GetNvidiaGpu() int32`

GetNvidiaGpu returns the NvidiaGpu field if non-nil, zero value otherwise.

### GetNvidiaGpuOk

`func (o *ClusterView) GetNvidiaGpuOk() (*int32, bool)`

GetNvidiaGpuOk returns a tuple with the NvidiaGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvidiaGpu

`func (o *ClusterView) SetNvidiaGpu(v int32)`

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


