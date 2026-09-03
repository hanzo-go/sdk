# NodePoolView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScale** | Pointer to **bool** | AutoScale reports whether the provider&#39;s cluster autoscaler owns this pool&#39;s size, moving Count between MinNodes and MaxNodes as workloads demand. False means Count changes only when someone scales the pool. | [optional] 
**Count** | Pointer to **int64** | Count is how many nodes the pool has right now. Always present, so 0 means a pool that is genuinely empty rather than a figure the provider withheld. | [optional] 
**MaxNodes** | Pointer to **int64** | MaxNodes is the ceiling the autoscaler will not grow the pool past, and so the bound on what this pool can cost. Read it only with AutoScale set. | [optional] 
**MinNodes** | Pointer to **int64** | MinNodes is the floor the autoscaler will not shrink the pool below. Read it only with AutoScale set — the provider ignores it otherwise. | [optional] 
**Name** | Pointer to **string** | Name is the pool&#39;s name as the provider knows it. | [optional] 
**PoolId** | Pointer to **string** | PoolID is the provider&#39;s id for the pool — the value the scale and delete routes address it by. It falls back to the pool&#39;s name when the provider answered without one, so it is always something the routes accept. | [optional] 
**Size** | Pointer to **string** | Size is the provider size slug every node in the pool runs at (\&quot;s-4vcpu-8gb\&quot;, \&quot;gpu-h100x8-640gb\&quot;). One pool is one size — a mixed cluster is several pools. | [optional] 

## Methods

### NewNodePoolView

`func NewNodePoolView() *NodePoolView`

NewNodePoolView instantiates a new NodePoolView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePoolViewWithDefaults

`func NewNodePoolViewWithDefaults() *NodePoolView`

NewNodePoolViewWithDefaults instantiates a new NodePoolView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScale

`func (o *NodePoolView) GetAutoScale() bool`

GetAutoScale returns the AutoScale field if non-nil, zero value otherwise.

### GetAutoScaleOk

`func (o *NodePoolView) GetAutoScaleOk() (*bool, bool)`

GetAutoScaleOk returns a tuple with the AutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScale

`func (o *NodePoolView) SetAutoScale(v bool)`

SetAutoScale sets AutoScale field to given value.

### HasAutoScale

`func (o *NodePoolView) HasAutoScale() bool`

HasAutoScale returns a boolean if a field has been set.

### GetCount

`func (o *NodePoolView) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NodePoolView) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NodePoolView) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *NodePoolView) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMaxNodes

`func (o *NodePoolView) GetMaxNodes() int64`

GetMaxNodes returns the MaxNodes field if non-nil, zero value otherwise.

### GetMaxNodesOk

`func (o *NodePoolView) GetMaxNodesOk() (*int64, bool)`

GetMaxNodesOk returns a tuple with the MaxNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodes

`func (o *NodePoolView) SetMaxNodes(v int64)`

SetMaxNodes sets MaxNodes field to given value.

### HasMaxNodes

`func (o *NodePoolView) HasMaxNodes() bool`

HasMaxNodes returns a boolean if a field has been set.

### GetMinNodes

`func (o *NodePoolView) GetMinNodes() int64`

GetMinNodes returns the MinNodes field if non-nil, zero value otherwise.

### GetMinNodesOk

`func (o *NodePoolView) GetMinNodesOk() (*int64, bool)`

GetMinNodesOk returns a tuple with the MinNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodes

`func (o *NodePoolView) SetMinNodes(v int64)`

SetMinNodes sets MinNodes field to given value.

### HasMinNodes

`func (o *NodePoolView) HasMinNodes() bool`

HasMinNodes returns a boolean if a field has been set.

### GetName

`func (o *NodePoolView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodePoolView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodePoolView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodePoolView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPoolId

`func (o *NodePoolView) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *NodePoolView) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *NodePoolView) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *NodePoolView) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetSize

`func (o *NodePoolView) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *NodePoolView) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *NodePoolView) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *NodePoolView) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


