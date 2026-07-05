# VisorClusterView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DoksClusterId** | Pointer to **string** |  | [optional] 
**DoClusterId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**NodePools** | Pointer to [**[]VisorNodePoolView**](VisorNodePoolView.md) |  | [optional] 
**NodeSize** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewVisorClusterView

`func NewVisorClusterView() *VisorClusterView`

NewVisorClusterView instantiates a new VisorClusterView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisorClusterViewWithDefaults

`func NewVisorClusterViewWithDefaults() *VisorClusterView`

NewVisorClusterViewWithDefaults instantiates a new VisorClusterView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoksClusterId

`func (o *VisorClusterView) GetDoksClusterId() string`

GetDoksClusterId returns the DoksClusterId field if non-nil, zero value otherwise.

### GetDoksClusterIdOk

`func (o *VisorClusterView) GetDoksClusterIdOk() (*string, bool)`

GetDoksClusterIdOk returns a tuple with the DoksClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoksClusterId

`func (o *VisorClusterView) SetDoksClusterId(v string)`

SetDoksClusterId sets DoksClusterId field to given value.

### HasDoksClusterId

`func (o *VisorClusterView) HasDoksClusterId() bool`

HasDoksClusterId returns a boolean if a field has been set.

### GetDoClusterId

`func (o *VisorClusterView) GetDoClusterId() string`

GetDoClusterId returns the DoClusterId field if non-nil, zero value otherwise.

### GetDoClusterIdOk

`func (o *VisorClusterView) GetDoClusterIdOk() (*string, bool)`

GetDoClusterIdOk returns a tuple with the DoClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoClusterId

`func (o *VisorClusterView) SetDoClusterId(v string)`

SetDoClusterId sets DoClusterId field to given value.

### HasDoClusterId

`func (o *VisorClusterView) HasDoClusterId() bool`

HasDoClusterId returns a boolean if a field has been set.

### GetName

`func (o *VisorClusterView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VisorClusterView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VisorClusterView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VisorClusterView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *VisorClusterView) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VisorClusterView) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VisorClusterView) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *VisorClusterView) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *VisorClusterView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VisorClusterView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VisorClusterView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VisorClusterView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNodePools

`func (o *VisorClusterView) GetNodePools() []VisorNodePoolView`

GetNodePools returns the NodePools field if non-nil, zero value otherwise.

### GetNodePoolsOk

`func (o *VisorClusterView) GetNodePoolsOk() (*[]VisorNodePoolView, bool)`

GetNodePoolsOk returns a tuple with the NodePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePools

`func (o *VisorClusterView) SetNodePools(v []VisorNodePoolView)`

SetNodePools sets NodePools field to given value.

### HasNodePools

`func (o *VisorClusterView) HasNodePools() bool`

HasNodePools returns a boolean if a field has been set.

### GetNodeSize

`func (o *VisorClusterView) GetNodeSize() string`

GetNodeSize returns the NodeSize field if non-nil, zero value otherwise.

### GetNodeSizeOk

`func (o *VisorClusterView) GetNodeSizeOk() (*string, bool)`

GetNodeSizeOk returns a tuple with the NodeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSize

`func (o *VisorClusterView) SetNodeSize(v string)`

SetNodeSize sets NodeSize field to given value.

### HasNodeSize

`func (o *VisorClusterView) HasNodeSize() bool`

HasNodeSize returns a boolean if a field has been set.

### GetNodeCount

`func (o *VisorClusterView) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *VisorClusterView) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *VisorClusterView) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *VisorClusterView) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VisorClusterView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VisorClusterView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VisorClusterView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VisorClusterView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


