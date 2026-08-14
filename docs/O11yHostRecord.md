# O11yHostRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveHostCount** | Pointer to **int32** |  | [optional] 
**Cpu** | Pointer to **float32** |  | [optional] 
**DiskUsage** | Pointer to **float32** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**InactiveHostCount** | Pointer to **int32** |  | [optional] 
**Load15** | Pointer to **float32** |  | [optional] 
**Memory** | Pointer to **float32** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**Status** | Pointer to **interface{}** |  | [optional] 
**Wait** | Pointer to **float32** |  | [optional] 

## Methods

### NewO11yHostRecord

`func NewO11yHostRecord() *O11yHostRecord`

NewO11yHostRecord instantiates a new O11yHostRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yHostRecordWithDefaults

`func NewO11yHostRecordWithDefaults() *O11yHostRecord`

NewO11yHostRecordWithDefaults instantiates a new O11yHostRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveHostCount

`func (o *O11yHostRecord) GetActiveHostCount() int32`

GetActiveHostCount returns the ActiveHostCount field if non-nil, zero value otherwise.

### GetActiveHostCountOk

`func (o *O11yHostRecord) GetActiveHostCountOk() (*int32, bool)`

GetActiveHostCountOk returns a tuple with the ActiveHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveHostCount

`func (o *O11yHostRecord) SetActiveHostCount(v int32)`

SetActiveHostCount sets ActiveHostCount field to given value.

### HasActiveHostCount

`func (o *O11yHostRecord) HasActiveHostCount() bool`

HasActiveHostCount returns a boolean if a field has been set.

### GetCpu

`func (o *O11yHostRecord) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *O11yHostRecord) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *O11yHostRecord) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *O11yHostRecord) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDiskUsage

`func (o *O11yHostRecord) GetDiskUsage() float32`

GetDiskUsage returns the DiskUsage field if non-nil, zero value otherwise.

### GetDiskUsageOk

`func (o *O11yHostRecord) GetDiskUsageOk() (*float32, bool)`

GetDiskUsageOk returns a tuple with the DiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsage

`func (o *O11yHostRecord) SetDiskUsage(v float32)`

SetDiskUsage sets DiskUsage field to given value.

### HasDiskUsage

`func (o *O11yHostRecord) HasDiskUsage() bool`

HasDiskUsage returns a boolean if a field has been set.

### GetHostName

`func (o *O11yHostRecord) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *O11yHostRecord) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *O11yHostRecord) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *O11yHostRecord) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetInactiveHostCount

`func (o *O11yHostRecord) GetInactiveHostCount() int32`

GetInactiveHostCount returns the InactiveHostCount field if non-nil, zero value otherwise.

### GetInactiveHostCountOk

`func (o *O11yHostRecord) GetInactiveHostCountOk() (*int32, bool)`

GetInactiveHostCountOk returns a tuple with the InactiveHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveHostCount

`func (o *O11yHostRecord) SetInactiveHostCount(v int32)`

SetInactiveHostCount sets InactiveHostCount field to given value.

### HasInactiveHostCount

`func (o *O11yHostRecord) HasInactiveHostCount() bool`

HasInactiveHostCount returns a boolean if a field has been set.

### GetLoad15

`func (o *O11yHostRecord) GetLoad15() float32`

GetLoad15 returns the Load15 field if non-nil, zero value otherwise.

### GetLoad15Ok

`func (o *O11yHostRecord) GetLoad15Ok() (*float32, bool)`

GetLoad15Ok returns a tuple with the Load15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad15

`func (o *O11yHostRecord) SetLoad15(v float32)`

SetLoad15 sets Load15 field to given value.

### HasLoad15

`func (o *O11yHostRecord) HasLoad15() bool`

HasLoad15 returns a boolean if a field has been set.

### GetMemory

`func (o *O11yHostRecord) GetMemory() float32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *O11yHostRecord) GetMemoryOk() (*float32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *O11yHostRecord) SetMemory(v float32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *O11yHostRecord) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMeta

`func (o *O11yHostRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yHostRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yHostRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yHostRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetStatus

`func (o *O11yHostRecord) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yHostRecord) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yHostRecord) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yHostRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *O11yHostRecord) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *O11yHostRecord) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetWait

`func (o *O11yHostRecord) GetWait() float32`

GetWait returns the Wait field if non-nil, zero value otherwise.

### GetWaitOk

`func (o *O11yHostRecord) GetWaitOk() (*float32, bool)`

GetWaitOk returns a tuple with the Wait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWait

`func (o *O11yHostRecord) SetWait(v float32)`

SetWait sets Wait field to given value.

### HasWait

`func (o *O11yHostRecord) HasWait() bool`

HasWait returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


