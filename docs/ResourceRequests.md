# ResourceRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **string** |  | [optional] 
**Memory** | Pointer to **string** |  | [optional] 

## Methods

### NewResourceRequests

`func NewResourceRequests() *ResourceRequests`

NewResourceRequests instantiates a new ResourceRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRequestsWithDefaults

`func NewResourceRequestsWithDefaults() *ResourceRequests`

NewResourceRequestsWithDefaults instantiates a new ResourceRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *ResourceRequests) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ResourceRequests) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ResourceRequests) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ResourceRequests) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *ResourceRequests) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ResourceRequests) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ResourceRequests) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ResourceRequests) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


