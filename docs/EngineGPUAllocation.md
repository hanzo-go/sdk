# EngineGPUAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuType** | Pointer to **string** |  | [optional] 
**Allocated** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**UtilizationPercent** | Pointer to **float32** |  | [optional] 
**Allocations** | Pointer to [**[]EngineGPUAllocationAllocationsInner**](EngineGPUAllocationAllocationsInner.md) |  | [optional] 

## Methods

### NewEngineGPUAllocation

`func NewEngineGPUAllocation() *EngineGPUAllocation`

NewEngineGPUAllocation instantiates a new EngineGPUAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineGPUAllocationWithDefaults

`func NewEngineGPUAllocationWithDefaults() *EngineGPUAllocation`

NewEngineGPUAllocationWithDefaults instantiates a new EngineGPUAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuType

`func (o *EngineGPUAllocation) GetGpuType() string`

GetGpuType returns the GpuType field if non-nil, zero value otherwise.

### GetGpuTypeOk

`func (o *EngineGPUAllocation) GetGpuTypeOk() (*string, bool)`

GetGpuTypeOk returns a tuple with the GpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuType

`func (o *EngineGPUAllocation) SetGpuType(v string)`

SetGpuType sets GpuType field to given value.

### HasGpuType

`func (o *EngineGPUAllocation) HasGpuType() bool`

HasGpuType returns a boolean if a field has been set.

### GetAllocated

`func (o *EngineGPUAllocation) GetAllocated() int32`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *EngineGPUAllocation) GetAllocatedOk() (*int32, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *EngineGPUAllocation) SetAllocated(v int32)`

SetAllocated sets Allocated field to given value.

### HasAllocated

`func (o *EngineGPUAllocation) HasAllocated() bool`

HasAllocated returns a boolean if a field has been set.

### GetTotal

`func (o *EngineGPUAllocation) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EngineGPUAllocation) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EngineGPUAllocation) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *EngineGPUAllocation) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUtilizationPercent

`func (o *EngineGPUAllocation) GetUtilizationPercent() float32`

GetUtilizationPercent returns the UtilizationPercent field if non-nil, zero value otherwise.

### GetUtilizationPercentOk

`func (o *EngineGPUAllocation) GetUtilizationPercentOk() (*float32, bool)`

GetUtilizationPercentOk returns a tuple with the UtilizationPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationPercent

`func (o *EngineGPUAllocation) SetUtilizationPercent(v float32)`

SetUtilizationPercent sets UtilizationPercent field to given value.

### HasUtilizationPercent

`func (o *EngineGPUAllocation) HasUtilizationPercent() bool`

HasUtilizationPercent returns a boolean if a field has been set.

### GetAllocations

`func (o *EngineGPUAllocation) GetAllocations() []EngineGPUAllocationAllocationsInner`

GetAllocations returns the Allocations field if non-nil, zero value otherwise.

### GetAllocationsOk

`func (o *EngineGPUAllocation) GetAllocationsOk() (*[]EngineGPUAllocationAllocationsInner, bool)`

GetAllocationsOk returns a tuple with the Allocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocations

`func (o *EngineGPUAllocation) SetAllocations(v []EngineGPUAllocationAllocationsInner)`

SetAllocations sets Allocations field to given value.

### HasAllocations

`func (o *EngineGPUAllocation) HasAllocations() bool`

HasAllocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


