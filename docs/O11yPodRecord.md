# O11yPodRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to **map[string]string** |  | [optional] 
**PodAge** | Pointer to **int32** |  | [optional] 
**PodCPU** | Pointer to **float32** |  | [optional] 
**PodCPULimit** | Pointer to **float32** |  | [optional] 
**PodCPURequest** | Pointer to **float32** |  | [optional] 
**PodCountsByPhase** | Pointer to [**O11yPodCountsByPhase**](O11yPodCountsByPhase.md) |  | [optional] 
**PodMemory** | Pointer to **float32** |  | [optional] 
**PodMemoryLimit** | Pointer to **float32** |  | [optional] 
**PodMemoryRequest** | Pointer to **float32** |  | [optional] 
**PodPhase** | Pointer to **interface{}** |  | [optional] 
**PodUID** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yPodRecord

`func NewO11yPodRecord() *O11yPodRecord`

NewO11yPodRecord instantiates a new O11yPodRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yPodRecordWithDefaults

`func NewO11yPodRecordWithDefaults() *O11yPodRecord`

NewO11yPodRecordWithDefaults instantiates a new O11yPodRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *O11yPodRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yPodRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yPodRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yPodRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPodAge

`func (o *O11yPodRecord) GetPodAge() int32`

GetPodAge returns the PodAge field if non-nil, zero value otherwise.

### GetPodAgeOk

`func (o *O11yPodRecord) GetPodAgeOk() (*int32, bool)`

GetPodAgeOk returns a tuple with the PodAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodAge

`func (o *O11yPodRecord) SetPodAge(v int32)`

SetPodAge sets PodAge field to given value.

### HasPodAge

`func (o *O11yPodRecord) HasPodAge() bool`

HasPodAge returns a boolean if a field has been set.

### GetPodCPU

`func (o *O11yPodRecord) GetPodCPU() float32`

GetPodCPU returns the PodCPU field if non-nil, zero value otherwise.

### GetPodCPUOk

`func (o *O11yPodRecord) GetPodCPUOk() (*float32, bool)`

GetPodCPUOk returns a tuple with the PodCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCPU

`func (o *O11yPodRecord) SetPodCPU(v float32)`

SetPodCPU sets PodCPU field to given value.

### HasPodCPU

`func (o *O11yPodRecord) HasPodCPU() bool`

HasPodCPU returns a boolean if a field has been set.

### GetPodCPULimit

`func (o *O11yPodRecord) GetPodCPULimit() float32`

GetPodCPULimit returns the PodCPULimit field if non-nil, zero value otherwise.

### GetPodCPULimitOk

`func (o *O11yPodRecord) GetPodCPULimitOk() (*float32, bool)`

GetPodCPULimitOk returns a tuple with the PodCPULimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCPULimit

`func (o *O11yPodRecord) SetPodCPULimit(v float32)`

SetPodCPULimit sets PodCPULimit field to given value.

### HasPodCPULimit

`func (o *O11yPodRecord) HasPodCPULimit() bool`

HasPodCPULimit returns a boolean if a field has been set.

### GetPodCPURequest

`func (o *O11yPodRecord) GetPodCPURequest() float32`

GetPodCPURequest returns the PodCPURequest field if non-nil, zero value otherwise.

### GetPodCPURequestOk

`func (o *O11yPodRecord) GetPodCPURequestOk() (*float32, bool)`

GetPodCPURequestOk returns a tuple with the PodCPURequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCPURequest

`func (o *O11yPodRecord) SetPodCPURequest(v float32)`

SetPodCPURequest sets PodCPURequest field to given value.

### HasPodCPURequest

`func (o *O11yPodRecord) HasPodCPURequest() bool`

HasPodCPURequest returns a boolean if a field has been set.

### GetPodCountsByPhase

`func (o *O11yPodRecord) GetPodCountsByPhase() O11yPodCountsByPhase`

GetPodCountsByPhase returns the PodCountsByPhase field if non-nil, zero value otherwise.

### GetPodCountsByPhaseOk

`func (o *O11yPodRecord) GetPodCountsByPhaseOk() (*O11yPodCountsByPhase, bool)`

GetPodCountsByPhaseOk returns a tuple with the PodCountsByPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCountsByPhase

`func (o *O11yPodRecord) SetPodCountsByPhase(v O11yPodCountsByPhase)`

SetPodCountsByPhase sets PodCountsByPhase field to given value.

### HasPodCountsByPhase

`func (o *O11yPodRecord) HasPodCountsByPhase() bool`

HasPodCountsByPhase returns a boolean if a field has been set.

### GetPodMemory

`func (o *O11yPodRecord) GetPodMemory() float32`

GetPodMemory returns the PodMemory field if non-nil, zero value otherwise.

### GetPodMemoryOk

`func (o *O11yPodRecord) GetPodMemoryOk() (*float32, bool)`

GetPodMemoryOk returns a tuple with the PodMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodMemory

`func (o *O11yPodRecord) SetPodMemory(v float32)`

SetPodMemory sets PodMemory field to given value.

### HasPodMemory

`func (o *O11yPodRecord) HasPodMemory() bool`

HasPodMemory returns a boolean if a field has been set.

### GetPodMemoryLimit

`func (o *O11yPodRecord) GetPodMemoryLimit() float32`

GetPodMemoryLimit returns the PodMemoryLimit field if non-nil, zero value otherwise.

### GetPodMemoryLimitOk

`func (o *O11yPodRecord) GetPodMemoryLimitOk() (*float32, bool)`

GetPodMemoryLimitOk returns a tuple with the PodMemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodMemoryLimit

`func (o *O11yPodRecord) SetPodMemoryLimit(v float32)`

SetPodMemoryLimit sets PodMemoryLimit field to given value.

### HasPodMemoryLimit

`func (o *O11yPodRecord) HasPodMemoryLimit() bool`

HasPodMemoryLimit returns a boolean if a field has been set.

### GetPodMemoryRequest

`func (o *O11yPodRecord) GetPodMemoryRequest() float32`

GetPodMemoryRequest returns the PodMemoryRequest field if non-nil, zero value otherwise.

### GetPodMemoryRequestOk

`func (o *O11yPodRecord) GetPodMemoryRequestOk() (*float32, bool)`

GetPodMemoryRequestOk returns a tuple with the PodMemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodMemoryRequest

`func (o *O11yPodRecord) SetPodMemoryRequest(v float32)`

SetPodMemoryRequest sets PodMemoryRequest field to given value.

### HasPodMemoryRequest

`func (o *O11yPodRecord) HasPodMemoryRequest() bool`

HasPodMemoryRequest returns a boolean if a field has been set.

### GetPodPhase

`func (o *O11yPodRecord) GetPodPhase() interface{}`

GetPodPhase returns the PodPhase field if non-nil, zero value otherwise.

### GetPodPhaseOk

`func (o *O11yPodRecord) GetPodPhaseOk() (*interface{}, bool)`

GetPodPhaseOk returns a tuple with the PodPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPhase

`func (o *O11yPodRecord) SetPodPhase(v interface{})`

SetPodPhase sets PodPhase field to given value.

### HasPodPhase

`func (o *O11yPodRecord) HasPodPhase() bool`

HasPodPhase returns a boolean if a field has been set.

### SetPodPhaseNil

`func (o *O11yPodRecord) SetPodPhaseNil(b bool)`

 SetPodPhaseNil sets the value for PodPhase to be an explicit nil

### UnsetPodPhase
`func (o *O11yPodRecord) UnsetPodPhase()`

UnsetPodPhase ensures that no value is present for PodPhase, not even an explicit nil
### GetPodUID

`func (o *O11yPodRecord) GetPodUID() string`

GetPodUID returns the PodUID field if non-nil, zero value otherwise.

### GetPodUIDOk

`func (o *O11yPodRecord) GetPodUIDOk() (*string, bool)`

GetPodUIDOk returns a tuple with the PodUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodUID

`func (o *O11yPodRecord) SetPodUID(v string)`

SetPodUID sets PodUID field to given value.

### HasPodUID

`func (o *O11yPodRecord) HasPodUID() bool`

HasPodUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


