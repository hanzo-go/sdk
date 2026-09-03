# O11yNodeListRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountByCondition** | Pointer to [**O11yNodeCountByCondition**](O11yNodeCountByCondition.md) |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**NodeCPUAllocatable** | Pointer to **float64** |  | [optional] 
**NodeCPUUsage** | Pointer to **float64** |  | [optional] 
**NodeMemoryAllocatable** | Pointer to **float64** |  | [optional] 
**NodeMemoryUsage** | Pointer to **float64** |  | [optional] 
**NodeUID** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yNodeListRecord

`func NewO11yNodeListRecord() *O11yNodeListRecord`

NewO11yNodeListRecord instantiates a new O11yNodeListRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yNodeListRecordWithDefaults

`func NewO11yNodeListRecordWithDefaults() *O11yNodeListRecord`

NewO11yNodeListRecordWithDefaults instantiates a new O11yNodeListRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountByCondition

`func (o *O11yNodeListRecord) GetCountByCondition() O11yNodeCountByCondition`

GetCountByCondition returns the CountByCondition field if non-nil, zero value otherwise.

### GetCountByConditionOk

`func (o *O11yNodeListRecord) GetCountByConditionOk() (*O11yNodeCountByCondition, bool)`

GetCountByConditionOk returns a tuple with the CountByCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountByCondition

`func (o *O11yNodeListRecord) SetCountByCondition(v O11yNodeCountByCondition)`

SetCountByCondition sets CountByCondition field to given value.

### HasCountByCondition

`func (o *O11yNodeListRecord) HasCountByCondition() bool`

HasCountByCondition returns a boolean if a field has been set.

### GetMeta

`func (o *O11yNodeListRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yNodeListRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yNodeListRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yNodeListRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetNodeCPUAllocatable

`func (o *O11yNodeListRecord) GetNodeCPUAllocatable() float64`

GetNodeCPUAllocatable returns the NodeCPUAllocatable field if non-nil, zero value otherwise.

### GetNodeCPUAllocatableOk

`func (o *O11yNodeListRecord) GetNodeCPUAllocatableOk() (*float64, bool)`

GetNodeCPUAllocatableOk returns a tuple with the NodeCPUAllocatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCPUAllocatable

`func (o *O11yNodeListRecord) SetNodeCPUAllocatable(v float64)`

SetNodeCPUAllocatable sets NodeCPUAllocatable field to given value.

### HasNodeCPUAllocatable

`func (o *O11yNodeListRecord) HasNodeCPUAllocatable() bool`

HasNodeCPUAllocatable returns a boolean if a field has been set.

### GetNodeCPUUsage

`func (o *O11yNodeListRecord) GetNodeCPUUsage() float64`

GetNodeCPUUsage returns the NodeCPUUsage field if non-nil, zero value otherwise.

### GetNodeCPUUsageOk

`func (o *O11yNodeListRecord) GetNodeCPUUsageOk() (*float64, bool)`

GetNodeCPUUsageOk returns a tuple with the NodeCPUUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCPUUsage

`func (o *O11yNodeListRecord) SetNodeCPUUsage(v float64)`

SetNodeCPUUsage sets NodeCPUUsage field to given value.

### HasNodeCPUUsage

`func (o *O11yNodeListRecord) HasNodeCPUUsage() bool`

HasNodeCPUUsage returns a boolean if a field has been set.

### GetNodeMemoryAllocatable

`func (o *O11yNodeListRecord) GetNodeMemoryAllocatable() float64`

GetNodeMemoryAllocatable returns the NodeMemoryAllocatable field if non-nil, zero value otherwise.

### GetNodeMemoryAllocatableOk

`func (o *O11yNodeListRecord) GetNodeMemoryAllocatableOk() (*float64, bool)`

GetNodeMemoryAllocatableOk returns a tuple with the NodeMemoryAllocatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeMemoryAllocatable

`func (o *O11yNodeListRecord) SetNodeMemoryAllocatable(v float64)`

SetNodeMemoryAllocatable sets NodeMemoryAllocatable field to given value.

### HasNodeMemoryAllocatable

`func (o *O11yNodeListRecord) HasNodeMemoryAllocatable() bool`

HasNodeMemoryAllocatable returns a boolean if a field has been set.

### GetNodeMemoryUsage

`func (o *O11yNodeListRecord) GetNodeMemoryUsage() float64`

GetNodeMemoryUsage returns the NodeMemoryUsage field if non-nil, zero value otherwise.

### GetNodeMemoryUsageOk

`func (o *O11yNodeListRecord) GetNodeMemoryUsageOk() (*float64, bool)`

GetNodeMemoryUsageOk returns a tuple with the NodeMemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeMemoryUsage

`func (o *O11yNodeListRecord) SetNodeMemoryUsage(v float64)`

SetNodeMemoryUsage sets NodeMemoryUsage field to given value.

### HasNodeMemoryUsage

`func (o *O11yNodeListRecord) HasNodeMemoryUsage() bool`

HasNodeMemoryUsage returns a boolean if a field has been set.

### GetNodeUID

`func (o *O11yNodeListRecord) GetNodeUID() string`

GetNodeUID returns the NodeUID field if non-nil, zero value otherwise.

### GetNodeUIDOk

`func (o *O11yNodeListRecord) GetNodeUIDOk() (*string, bool)`

GetNodeUIDOk returns a tuple with the NodeUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeUID

`func (o *O11yNodeListRecord) SetNodeUID(v string)`

SetNodeUID sets NodeUID field to given value.

### HasNodeUID

`func (o *O11yNodeListRecord) HasNodeUID() bool`

HasNodeUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


