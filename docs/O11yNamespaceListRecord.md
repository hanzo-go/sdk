# O11yNamespaceListRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountByPhase** | Pointer to [**O11yPodCountByPhase**](O11yPodCountByPhase.md) |  | [optional] 
**CpuUsage** | Pointer to **float64** |  | [optional] 
**MemoryUsage** | Pointer to **float64** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**NamespaceName** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yNamespaceListRecord

`func NewO11yNamespaceListRecord() *O11yNamespaceListRecord`

NewO11yNamespaceListRecord instantiates a new O11yNamespaceListRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yNamespaceListRecordWithDefaults

`func NewO11yNamespaceListRecordWithDefaults() *O11yNamespaceListRecord`

NewO11yNamespaceListRecordWithDefaults instantiates a new O11yNamespaceListRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountByPhase

`func (o *O11yNamespaceListRecord) GetCountByPhase() O11yPodCountByPhase`

GetCountByPhase returns the CountByPhase field if non-nil, zero value otherwise.

### GetCountByPhaseOk

`func (o *O11yNamespaceListRecord) GetCountByPhaseOk() (*O11yPodCountByPhase, bool)`

GetCountByPhaseOk returns a tuple with the CountByPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountByPhase

`func (o *O11yNamespaceListRecord) SetCountByPhase(v O11yPodCountByPhase)`

SetCountByPhase sets CountByPhase field to given value.

### HasCountByPhase

`func (o *O11yNamespaceListRecord) HasCountByPhase() bool`

HasCountByPhase returns a boolean if a field has been set.

### GetCpuUsage

`func (o *O11yNamespaceListRecord) GetCpuUsage() float64`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *O11yNamespaceListRecord) GetCpuUsageOk() (*float64, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *O11yNamespaceListRecord) SetCpuUsage(v float64)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *O11yNamespaceListRecord) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetMemoryUsage

`func (o *O11yNamespaceListRecord) GetMemoryUsage() float64`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *O11yNamespaceListRecord) GetMemoryUsageOk() (*float64, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *O11yNamespaceListRecord) SetMemoryUsage(v float64)`

SetMemoryUsage sets MemoryUsage field to given value.

### HasMemoryUsage

`func (o *O11yNamespaceListRecord) HasMemoryUsage() bool`

HasMemoryUsage returns a boolean if a field has been set.

### GetMeta

`func (o *O11yNamespaceListRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yNamespaceListRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yNamespaceListRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yNamespaceListRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetNamespaceName

`func (o *O11yNamespaceListRecord) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *O11yNamespaceListRecord) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *O11yNamespaceListRecord) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.

### HasNamespaceName

`func (o *O11yNamespaceListRecord) HasNamespaceName() bool`

HasNamespaceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


