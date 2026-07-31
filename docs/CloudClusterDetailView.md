# CloudClusterDetailView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmdGpu** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DoClusterId** | Pointer to **string** |  | [optional] 
**DoksClusterId** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] 
**NodePools** | Pointer to [**[]CloudNodePoolView**](CloudNodePoolView.md) |  | [optional] 
**NodeSize** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to [**[]CloudMachineView**](CloudMachineView.md) |  | [optional] 
**NvidiaGpu** | Pointer to **int32** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudClusterDetailView

`func NewCloudClusterDetailView() *CloudClusterDetailView`

NewCloudClusterDetailView instantiates a new CloudClusterDetailView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudClusterDetailViewWithDefaults

`func NewCloudClusterDetailViewWithDefaults() *CloudClusterDetailView`

NewCloudClusterDetailViewWithDefaults instantiates a new CloudClusterDetailView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmdGpu

`func (o *CloudClusterDetailView) GetAmdGpu() int32`

GetAmdGpu returns the AmdGpu field if non-nil, zero value otherwise.

### GetAmdGpuOk

`func (o *CloudClusterDetailView) GetAmdGpuOk() (*int32, bool)`

GetAmdGpuOk returns a tuple with the AmdGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmdGpu

`func (o *CloudClusterDetailView) SetAmdGpu(v int32)`

SetAmdGpu sets AmdGpu field to given value.

### HasAmdGpu

`func (o *CloudClusterDetailView) HasAmdGpu() bool`

HasAmdGpu returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudClusterDetailView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudClusterDetailView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudClusterDetailView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudClusterDetailView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDoClusterId

`func (o *CloudClusterDetailView) GetDoClusterId() string`

GetDoClusterId returns the DoClusterId field if non-nil, zero value otherwise.

### GetDoClusterIdOk

`func (o *CloudClusterDetailView) GetDoClusterIdOk() (*string, bool)`

GetDoClusterIdOk returns a tuple with the DoClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoClusterId

`func (o *CloudClusterDetailView) SetDoClusterId(v string)`

SetDoClusterId sets DoClusterId field to given value.

### HasDoClusterId

`func (o *CloudClusterDetailView) HasDoClusterId() bool`

HasDoClusterId returns a boolean if a field has been set.

### GetDoksClusterId

`func (o *CloudClusterDetailView) GetDoksClusterId() string`

GetDoksClusterId returns the DoksClusterId field if non-nil, zero value otherwise.

### GetDoksClusterIdOk

`func (o *CloudClusterDetailView) GetDoksClusterIdOk() (*string, bool)`

GetDoksClusterIdOk returns a tuple with the DoksClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoksClusterId

`func (o *CloudClusterDetailView) SetDoksClusterId(v string)`

SetDoksClusterId sets DoksClusterId field to given value.

### HasDoksClusterId

`func (o *CloudClusterDetailView) HasDoksClusterId() bool`

HasDoksClusterId returns a boolean if a field has been set.

### GetKind

`func (o *CloudClusterDetailView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudClusterDetailView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudClusterDetailView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudClusterDetailView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *CloudClusterDetailView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudClusterDetailView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudClusterDetailView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudClusterDetailView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeCount

`func (o *CloudClusterDetailView) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *CloudClusterDetailView) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *CloudClusterDetailView) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *CloudClusterDetailView) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodePools

`func (o *CloudClusterDetailView) GetNodePools() []CloudNodePoolView`

GetNodePools returns the NodePools field if non-nil, zero value otherwise.

### GetNodePoolsOk

`func (o *CloudClusterDetailView) GetNodePoolsOk() (*[]CloudNodePoolView, bool)`

GetNodePoolsOk returns a tuple with the NodePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePools

`func (o *CloudClusterDetailView) SetNodePools(v []CloudNodePoolView)`

SetNodePools sets NodePools field to given value.

### HasNodePools

`func (o *CloudClusterDetailView) HasNodePools() bool`

HasNodePools returns a boolean if a field has been set.

### GetNodeSize

`func (o *CloudClusterDetailView) GetNodeSize() string`

GetNodeSize returns the NodeSize field if non-nil, zero value otherwise.

### GetNodeSizeOk

`func (o *CloudClusterDetailView) GetNodeSizeOk() (*string, bool)`

GetNodeSizeOk returns a tuple with the NodeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSize

`func (o *CloudClusterDetailView) SetNodeSize(v string)`

SetNodeSize sets NodeSize field to given value.

### HasNodeSize

`func (o *CloudClusterDetailView) HasNodeSize() bool`

HasNodeSize returns a boolean if a field has been set.

### GetNodes

`func (o *CloudClusterDetailView) GetNodes() []CloudMachineView`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *CloudClusterDetailView) GetNodesOk() (*[]CloudMachineView, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *CloudClusterDetailView) SetNodes(v []CloudMachineView)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *CloudClusterDetailView) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetNvidiaGpu

`func (o *CloudClusterDetailView) GetNvidiaGpu() int32`

GetNvidiaGpu returns the NvidiaGpu field if non-nil, zero value otherwise.

### GetNvidiaGpuOk

`func (o *CloudClusterDetailView) GetNvidiaGpuOk() (*int32, bool)`

GetNvidiaGpuOk returns a tuple with the NvidiaGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvidiaGpu

`func (o *CloudClusterDetailView) SetNvidiaGpu(v int32)`

SetNvidiaGpu sets NvidiaGpu field to given value.

### HasNvidiaGpu

`func (o *CloudClusterDetailView) HasNvidiaGpu() bool`

HasNvidiaGpu returns a boolean if a field has been set.

### GetRegion

`func (o *CloudClusterDetailView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudClusterDetailView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudClusterDetailView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudClusterDetailView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *CloudClusterDetailView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudClusterDetailView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudClusterDetailView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudClusterDetailView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


