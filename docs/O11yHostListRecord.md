# O11yHostListRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**Cpu** | Pointer to **float32** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**Load15** | Pointer to **float32** |  | [optional] 
**Memory** | Pointer to **float32** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**Wait** | Pointer to **float32** |  | [optional] 

## Methods

### NewO11yHostListRecord

`func NewO11yHostListRecord() *O11yHostListRecord`

NewO11yHostListRecord instantiates a new O11yHostListRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yHostListRecordWithDefaults

`func NewO11yHostListRecordWithDefaults() *O11yHostListRecord`

NewO11yHostListRecordWithDefaults instantiates a new O11yHostListRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *O11yHostListRecord) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *O11yHostListRecord) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *O11yHostListRecord) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *O11yHostListRecord) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCpu

`func (o *O11yHostListRecord) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *O11yHostListRecord) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *O11yHostListRecord) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *O11yHostListRecord) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetHostName

`func (o *O11yHostListRecord) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *O11yHostListRecord) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *O11yHostListRecord) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *O11yHostListRecord) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetLoad15

`func (o *O11yHostListRecord) GetLoad15() float32`

GetLoad15 returns the Load15 field if non-nil, zero value otherwise.

### GetLoad15Ok

`func (o *O11yHostListRecord) GetLoad15Ok() (*float32, bool)`

GetLoad15Ok returns a tuple with the Load15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad15

`func (o *O11yHostListRecord) SetLoad15(v float32)`

SetLoad15 sets Load15 field to given value.

### HasLoad15

`func (o *O11yHostListRecord) HasLoad15() bool`

HasLoad15 returns a boolean if a field has been set.

### GetMemory

`func (o *O11yHostListRecord) GetMemory() float32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *O11yHostListRecord) GetMemoryOk() (*float32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *O11yHostListRecord) SetMemory(v float32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *O11yHostListRecord) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMeta

`func (o *O11yHostListRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yHostListRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yHostListRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yHostListRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetOs

`func (o *O11yHostListRecord) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *O11yHostListRecord) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *O11yHostListRecord) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *O11yHostListRecord) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetWait

`func (o *O11yHostListRecord) GetWait() float32`

GetWait returns the Wait field if non-nil, zero value otherwise.

### GetWaitOk

`func (o *O11yHostListRecord) GetWaitOk() (*float32, bool)`

GetWaitOk returns a tuple with the Wait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWait

`func (o *O11yHostListRecord) SetWait(v float32)`

SetWait sets Wait field to given value.

### HasWait

`func (o *O11yHostListRecord) HasWait() bool`

HasWait returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


