# O11yPodListRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountByPhase** | Pointer to [**O11yPodCountByPhase**](O11yPodCountByPhase.md) |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**PodCPU** | Pointer to **float32** |  | [optional] 
**PodCPULimit** | Pointer to **float32** |  | [optional] 
**PodCPURequest** | Pointer to **float32** |  | [optional] 
**PodMemory** | Pointer to **float32** |  | [optional] 
**PodMemoryLimit** | Pointer to **float32** |  | [optional] 
**PodMemoryRequest** | Pointer to **float32** |  | [optional] 
**PodUID** | Pointer to **string** |  | [optional] 
**RestartCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewO11yPodListRecord

`func NewO11yPodListRecord() *O11yPodListRecord`

NewO11yPodListRecord instantiates a new O11yPodListRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yPodListRecordWithDefaults

`func NewO11yPodListRecordWithDefaults() *O11yPodListRecord`

NewO11yPodListRecordWithDefaults instantiates a new O11yPodListRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountByPhase

`func (o *O11yPodListRecord) GetCountByPhase() O11yPodCountByPhase`

GetCountByPhase returns the CountByPhase field if non-nil, zero value otherwise.

### GetCountByPhaseOk

`func (o *O11yPodListRecord) GetCountByPhaseOk() (*O11yPodCountByPhase, bool)`

GetCountByPhaseOk returns a tuple with the CountByPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountByPhase

`func (o *O11yPodListRecord) SetCountByPhase(v O11yPodCountByPhase)`

SetCountByPhase sets CountByPhase field to given value.

### HasCountByPhase

`func (o *O11yPodListRecord) HasCountByPhase() bool`

HasCountByPhase returns a boolean if a field has been set.

### GetMeta

`func (o *O11yPodListRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yPodListRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yPodListRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yPodListRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPodCPU

`func (o *O11yPodListRecord) GetPodCPU() float32`

GetPodCPU returns the PodCPU field if non-nil, zero value otherwise.

### GetPodCPUOk

`func (o *O11yPodListRecord) GetPodCPUOk() (*float32, bool)`

GetPodCPUOk returns a tuple with the PodCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCPU

`func (o *O11yPodListRecord) SetPodCPU(v float32)`

SetPodCPU sets PodCPU field to given value.

### HasPodCPU

`func (o *O11yPodListRecord) HasPodCPU() bool`

HasPodCPU returns a boolean if a field has been set.

### GetPodCPULimit

`func (o *O11yPodListRecord) GetPodCPULimit() float32`

GetPodCPULimit returns the PodCPULimit field if non-nil, zero value otherwise.

### GetPodCPULimitOk

`func (o *O11yPodListRecord) GetPodCPULimitOk() (*float32, bool)`

GetPodCPULimitOk returns a tuple with the PodCPULimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCPULimit

`func (o *O11yPodListRecord) SetPodCPULimit(v float32)`

SetPodCPULimit sets PodCPULimit field to given value.

### HasPodCPULimit

`func (o *O11yPodListRecord) HasPodCPULimit() bool`

HasPodCPULimit returns a boolean if a field has been set.

### GetPodCPURequest

`func (o *O11yPodListRecord) GetPodCPURequest() float32`

GetPodCPURequest returns the PodCPURequest field if non-nil, zero value otherwise.

### GetPodCPURequestOk

`func (o *O11yPodListRecord) GetPodCPURequestOk() (*float32, bool)`

GetPodCPURequestOk returns a tuple with the PodCPURequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCPURequest

`func (o *O11yPodListRecord) SetPodCPURequest(v float32)`

SetPodCPURequest sets PodCPURequest field to given value.

### HasPodCPURequest

`func (o *O11yPodListRecord) HasPodCPURequest() bool`

HasPodCPURequest returns a boolean if a field has been set.

### GetPodMemory

`func (o *O11yPodListRecord) GetPodMemory() float32`

GetPodMemory returns the PodMemory field if non-nil, zero value otherwise.

### GetPodMemoryOk

`func (o *O11yPodListRecord) GetPodMemoryOk() (*float32, bool)`

GetPodMemoryOk returns a tuple with the PodMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodMemory

`func (o *O11yPodListRecord) SetPodMemory(v float32)`

SetPodMemory sets PodMemory field to given value.

### HasPodMemory

`func (o *O11yPodListRecord) HasPodMemory() bool`

HasPodMemory returns a boolean if a field has been set.

### GetPodMemoryLimit

`func (o *O11yPodListRecord) GetPodMemoryLimit() float32`

GetPodMemoryLimit returns the PodMemoryLimit field if non-nil, zero value otherwise.

### GetPodMemoryLimitOk

`func (o *O11yPodListRecord) GetPodMemoryLimitOk() (*float32, bool)`

GetPodMemoryLimitOk returns a tuple with the PodMemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodMemoryLimit

`func (o *O11yPodListRecord) SetPodMemoryLimit(v float32)`

SetPodMemoryLimit sets PodMemoryLimit field to given value.

### HasPodMemoryLimit

`func (o *O11yPodListRecord) HasPodMemoryLimit() bool`

HasPodMemoryLimit returns a boolean if a field has been set.

### GetPodMemoryRequest

`func (o *O11yPodListRecord) GetPodMemoryRequest() float32`

GetPodMemoryRequest returns the PodMemoryRequest field if non-nil, zero value otherwise.

### GetPodMemoryRequestOk

`func (o *O11yPodListRecord) GetPodMemoryRequestOk() (*float32, bool)`

GetPodMemoryRequestOk returns a tuple with the PodMemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodMemoryRequest

`func (o *O11yPodListRecord) SetPodMemoryRequest(v float32)`

SetPodMemoryRequest sets PodMemoryRequest field to given value.

### HasPodMemoryRequest

`func (o *O11yPodListRecord) HasPodMemoryRequest() bool`

HasPodMemoryRequest returns a boolean if a field has been set.

### GetPodUID

`func (o *O11yPodListRecord) GetPodUID() string`

GetPodUID returns the PodUID field if non-nil, zero value otherwise.

### GetPodUIDOk

`func (o *O11yPodListRecord) GetPodUIDOk() (*string, bool)`

GetPodUIDOk returns a tuple with the PodUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodUID

`func (o *O11yPodListRecord) SetPodUID(v string)`

SetPodUID sets PodUID field to given value.

### HasPodUID

`func (o *O11yPodListRecord) HasPodUID() bool`

HasPodUID returns a boolean if a field has been set.

### GetRestartCount

`func (o *O11yPodListRecord) GetRestartCount() int32`

GetRestartCount returns the RestartCount field if non-nil, zero value otherwise.

### GetRestartCountOk

`func (o *O11yPodListRecord) GetRestartCountOk() (*int32, bool)`

GetRestartCountOk returns a tuple with the RestartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount

`func (o *O11yPodListRecord) SetRestartCount(v int32)`

SetRestartCount sets RestartCount field to given value.

### HasRestartCount

`func (o *O11yPodListRecord) HasRestartCount() bool`

HasRestartCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


