# CloudClusterView

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
**NodePools** | Pointer to [**[]CloudNodePoolView**](CloudNodePoolView.md) |  | [optional] 
**NodeSize** | Pointer to **string** |  | [optional] 
**NvidiaGpu** | Pointer to **int32** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudClusterView

`func NewCloudClusterView() *CloudClusterView`

NewCloudClusterView instantiates a new CloudClusterView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudClusterViewWithDefaults

`func NewCloudClusterViewWithDefaults() *CloudClusterView`

NewCloudClusterViewWithDefaults instantiates a new CloudClusterView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmdGpu

`func (o *CloudClusterView) GetAmdGpu() int32`

GetAmdGpu returns the AmdGpu field if non-nil, zero value otherwise.

### GetAmdGpuOk

`func (o *CloudClusterView) GetAmdGpuOk() (*int32, bool)`

GetAmdGpuOk returns a tuple with the AmdGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmdGpu

`func (o *CloudClusterView) SetAmdGpu(v int32)`

SetAmdGpu sets AmdGpu field to given value.

### HasAmdGpu

`func (o *CloudClusterView) HasAmdGpu() bool`

HasAmdGpu returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudClusterView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudClusterView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudClusterView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudClusterView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDoClusterId

`func (o *CloudClusterView) GetDoClusterId() string`

GetDoClusterId returns the DoClusterId field if non-nil, zero value otherwise.

### GetDoClusterIdOk

`func (o *CloudClusterView) GetDoClusterIdOk() (*string, bool)`

GetDoClusterIdOk returns a tuple with the DoClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoClusterId

`func (o *CloudClusterView) SetDoClusterId(v string)`

SetDoClusterId sets DoClusterId field to given value.

### HasDoClusterId

`func (o *CloudClusterView) HasDoClusterId() bool`

HasDoClusterId returns a boolean if a field has been set.

### GetDoksClusterId

`func (o *CloudClusterView) GetDoksClusterId() string`

GetDoksClusterId returns the DoksClusterId field if non-nil, zero value otherwise.

### GetDoksClusterIdOk

`func (o *CloudClusterView) GetDoksClusterIdOk() (*string, bool)`

GetDoksClusterIdOk returns a tuple with the DoksClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoksClusterId

`func (o *CloudClusterView) SetDoksClusterId(v string)`

SetDoksClusterId sets DoksClusterId field to given value.

### HasDoksClusterId

`func (o *CloudClusterView) HasDoksClusterId() bool`

HasDoksClusterId returns a boolean if a field has been set.

### GetKind

`func (o *CloudClusterView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudClusterView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudClusterView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudClusterView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *CloudClusterView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudClusterView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudClusterView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudClusterView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeCount

`func (o *CloudClusterView) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *CloudClusterView) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *CloudClusterView) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *CloudClusterView) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodePools

`func (o *CloudClusterView) GetNodePools() []CloudNodePoolView`

GetNodePools returns the NodePools field if non-nil, zero value otherwise.

### GetNodePoolsOk

`func (o *CloudClusterView) GetNodePoolsOk() (*[]CloudNodePoolView, bool)`

GetNodePoolsOk returns a tuple with the NodePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePools

`func (o *CloudClusterView) SetNodePools(v []CloudNodePoolView)`

SetNodePools sets NodePools field to given value.

### HasNodePools

`func (o *CloudClusterView) HasNodePools() bool`

HasNodePools returns a boolean if a field has been set.

### GetNodeSize

`func (o *CloudClusterView) GetNodeSize() string`

GetNodeSize returns the NodeSize field if non-nil, zero value otherwise.

### GetNodeSizeOk

`func (o *CloudClusterView) GetNodeSizeOk() (*string, bool)`

GetNodeSizeOk returns a tuple with the NodeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSize

`func (o *CloudClusterView) SetNodeSize(v string)`

SetNodeSize sets NodeSize field to given value.

### HasNodeSize

`func (o *CloudClusterView) HasNodeSize() bool`

HasNodeSize returns a boolean if a field has been set.

### GetNvidiaGpu

`func (o *CloudClusterView) GetNvidiaGpu() int32`

GetNvidiaGpu returns the NvidiaGpu field if non-nil, zero value otherwise.

### GetNvidiaGpuOk

`func (o *CloudClusterView) GetNvidiaGpuOk() (*int32, bool)`

GetNvidiaGpuOk returns a tuple with the NvidiaGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvidiaGpu

`func (o *CloudClusterView) SetNvidiaGpu(v int32)`

SetNvidiaGpu sets NvidiaGpu field to given value.

### HasNvidiaGpu

`func (o *CloudClusterView) HasNvidiaGpu() bool`

HasNvidiaGpu returns a boolean if a field has been set.

### GetRegion

`func (o *CloudClusterView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudClusterView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudClusterView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudClusterView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *CloudClusterView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudClusterView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudClusterView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudClusterView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


