# ComputeNodePoolView

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

### NewComputeNodePoolView

`func NewComputeNodePoolView() *ComputeNodePoolView`

NewComputeNodePoolView instantiates a new ComputeNodePoolView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeNodePoolViewWithDefaults

`func NewComputeNodePoolViewWithDefaults() *ComputeNodePoolView`

NewComputeNodePoolViewWithDefaults instantiates a new ComputeNodePoolView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScale

`func (o *ComputeNodePoolView) GetAutoScale() bool`

GetAutoScale returns the AutoScale field if non-nil, zero value otherwise.

### GetAutoScaleOk

`func (o *ComputeNodePoolView) GetAutoScaleOk() (*bool, bool)`

GetAutoScaleOk returns a tuple with the AutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScale

`func (o *ComputeNodePoolView) SetAutoScale(v bool)`

SetAutoScale sets AutoScale field to given value.

### HasAutoScale

`func (o *ComputeNodePoolView) HasAutoScale() bool`

HasAutoScale returns a boolean if a field has been set.

### GetCount

`func (o *ComputeNodePoolView) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ComputeNodePoolView) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ComputeNodePoolView) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ComputeNodePoolView) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMaxNodes

`func (o *ComputeNodePoolView) GetMaxNodes() int64`

GetMaxNodes returns the MaxNodes field if non-nil, zero value otherwise.

### GetMaxNodesOk

`func (o *ComputeNodePoolView) GetMaxNodesOk() (*int64, bool)`

GetMaxNodesOk returns a tuple with the MaxNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodes

`func (o *ComputeNodePoolView) SetMaxNodes(v int64)`

SetMaxNodes sets MaxNodes field to given value.

### HasMaxNodes

`func (o *ComputeNodePoolView) HasMaxNodes() bool`

HasMaxNodes returns a boolean if a field has been set.

### GetMinNodes

`func (o *ComputeNodePoolView) GetMinNodes() int64`

GetMinNodes returns the MinNodes field if non-nil, zero value otherwise.

### GetMinNodesOk

`func (o *ComputeNodePoolView) GetMinNodesOk() (*int64, bool)`

GetMinNodesOk returns a tuple with the MinNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodes

`func (o *ComputeNodePoolView) SetMinNodes(v int64)`

SetMinNodes sets MinNodes field to given value.

### HasMinNodes

`func (o *ComputeNodePoolView) HasMinNodes() bool`

HasMinNodes returns a boolean if a field has been set.

### GetName

`func (o *ComputeNodePoolView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputeNodePoolView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputeNodePoolView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputeNodePoolView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPoolId

`func (o *ComputeNodePoolView) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *ComputeNodePoolView) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *ComputeNodePoolView) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *ComputeNodePoolView) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetSize

`func (o *ComputeNodePoolView) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ComputeNodePoolView) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ComputeNodePoolView) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ComputeNodePoolView) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


