# O11yNamespaceRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to **map[string]string** |  | [optional] 
**NamespaceCPU** | Pointer to **float64** |  | [optional] 
**NamespaceMemory** | Pointer to **float64** |  | [optional] 
**NamespaceName** | Pointer to **string** |  | [optional] 
**PodCountsByPhase** | Pointer to [**O11yPodCountsByPhase**](O11yPodCountsByPhase.md) |  | [optional] 

## Methods

### NewO11yNamespaceRecord

`func NewO11yNamespaceRecord() *O11yNamespaceRecord`

NewO11yNamespaceRecord instantiates a new O11yNamespaceRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yNamespaceRecordWithDefaults

`func NewO11yNamespaceRecordWithDefaults() *O11yNamespaceRecord`

NewO11yNamespaceRecordWithDefaults instantiates a new O11yNamespaceRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *O11yNamespaceRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yNamespaceRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yNamespaceRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yNamespaceRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetNamespaceCPU

`func (o *O11yNamespaceRecord) GetNamespaceCPU() float64`

GetNamespaceCPU returns the NamespaceCPU field if non-nil, zero value otherwise.

### GetNamespaceCPUOk

`func (o *O11yNamespaceRecord) GetNamespaceCPUOk() (*float64, bool)`

GetNamespaceCPUOk returns a tuple with the NamespaceCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceCPU

`func (o *O11yNamespaceRecord) SetNamespaceCPU(v float64)`

SetNamespaceCPU sets NamespaceCPU field to given value.

### HasNamespaceCPU

`func (o *O11yNamespaceRecord) HasNamespaceCPU() bool`

HasNamespaceCPU returns a boolean if a field has been set.

### GetNamespaceMemory

`func (o *O11yNamespaceRecord) GetNamespaceMemory() float64`

GetNamespaceMemory returns the NamespaceMemory field if non-nil, zero value otherwise.

### GetNamespaceMemoryOk

`func (o *O11yNamespaceRecord) GetNamespaceMemoryOk() (*float64, bool)`

GetNamespaceMemoryOk returns a tuple with the NamespaceMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceMemory

`func (o *O11yNamespaceRecord) SetNamespaceMemory(v float64)`

SetNamespaceMemory sets NamespaceMemory field to given value.

### HasNamespaceMemory

`func (o *O11yNamespaceRecord) HasNamespaceMemory() bool`

HasNamespaceMemory returns a boolean if a field has been set.

### GetNamespaceName

`func (o *O11yNamespaceRecord) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *O11yNamespaceRecord) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *O11yNamespaceRecord) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.

### HasNamespaceName

`func (o *O11yNamespaceRecord) HasNamespaceName() bool`

HasNamespaceName returns a boolean if a field has been set.

### GetPodCountsByPhase

`func (o *O11yNamespaceRecord) GetPodCountsByPhase() O11yPodCountsByPhase`

GetPodCountsByPhase returns the PodCountsByPhase field if non-nil, zero value otherwise.

### GetPodCountsByPhaseOk

`func (o *O11yNamespaceRecord) GetPodCountsByPhaseOk() (*O11yPodCountsByPhase, bool)`

GetPodCountsByPhaseOk returns a tuple with the PodCountsByPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCountsByPhase

`func (o *O11yNamespaceRecord) SetPodCountsByPhase(v O11yPodCountsByPhase)`

SetPodCountsByPhase sets PodCountsByPhase field to given value.

### HasPodCountsByPhase

`func (o *O11yNamespaceRecord) HasPodCountsByPhase() bool`

HasPodCountsByPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


