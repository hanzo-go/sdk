# O11yClusterListRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterUID** | Pointer to **string** |  | [optional] 
**CpuAllocatable** | Pointer to **float32** |  | [optional] 
**CpuUsage** | Pointer to **float32** |  | [optional] 
**MemoryAllocatable** | Pointer to **float32** |  | [optional] 
**MemoryUsage** | Pointer to **float32** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewO11yClusterListRecord

`func NewO11yClusterListRecord() *O11yClusterListRecord`

NewO11yClusterListRecord instantiates a new O11yClusterListRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yClusterListRecordWithDefaults

`func NewO11yClusterListRecordWithDefaults() *O11yClusterListRecord`

NewO11yClusterListRecordWithDefaults instantiates a new O11yClusterListRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterUID

`func (o *O11yClusterListRecord) GetClusterUID() string`

GetClusterUID returns the ClusterUID field if non-nil, zero value otherwise.

### GetClusterUIDOk

`func (o *O11yClusterListRecord) GetClusterUIDOk() (*string, bool)`

GetClusterUIDOk returns a tuple with the ClusterUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUID

`func (o *O11yClusterListRecord) SetClusterUID(v string)`

SetClusterUID sets ClusterUID field to given value.

### HasClusterUID

`func (o *O11yClusterListRecord) HasClusterUID() bool`

HasClusterUID returns a boolean if a field has been set.

### GetCpuAllocatable

`func (o *O11yClusterListRecord) GetCpuAllocatable() float32`

GetCpuAllocatable returns the CpuAllocatable field if non-nil, zero value otherwise.

### GetCpuAllocatableOk

`func (o *O11yClusterListRecord) GetCpuAllocatableOk() (*float32, bool)`

GetCpuAllocatableOk returns a tuple with the CpuAllocatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuAllocatable

`func (o *O11yClusterListRecord) SetCpuAllocatable(v float32)`

SetCpuAllocatable sets CpuAllocatable field to given value.

### HasCpuAllocatable

`func (o *O11yClusterListRecord) HasCpuAllocatable() bool`

HasCpuAllocatable returns a boolean if a field has been set.

### GetCpuUsage

`func (o *O11yClusterListRecord) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *O11yClusterListRecord) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *O11yClusterListRecord) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *O11yClusterListRecord) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetMemoryAllocatable

`func (o *O11yClusterListRecord) GetMemoryAllocatable() float32`

GetMemoryAllocatable returns the MemoryAllocatable field if non-nil, zero value otherwise.

### GetMemoryAllocatableOk

`func (o *O11yClusterListRecord) GetMemoryAllocatableOk() (*float32, bool)`

GetMemoryAllocatableOk returns a tuple with the MemoryAllocatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryAllocatable

`func (o *O11yClusterListRecord) SetMemoryAllocatable(v float32)`

SetMemoryAllocatable sets MemoryAllocatable field to given value.

### HasMemoryAllocatable

`func (o *O11yClusterListRecord) HasMemoryAllocatable() bool`

HasMemoryAllocatable returns a boolean if a field has been set.

### GetMemoryUsage

`func (o *O11yClusterListRecord) GetMemoryUsage() float32`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *O11yClusterListRecord) GetMemoryUsageOk() (*float32, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *O11yClusterListRecord) SetMemoryUsage(v float32)`

SetMemoryUsage sets MemoryUsage field to given value.

### HasMemoryUsage

`func (o *O11yClusterListRecord) HasMemoryUsage() bool`

HasMemoryUsage returns a boolean if a field has been set.

### GetMeta

`func (o *O11yClusterListRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yClusterListRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yClusterListRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yClusterListRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


