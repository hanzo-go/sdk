# O11yStatefulSetListRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailablePods** | Pointer to **int32** |  | [optional] 
**CpuLimit** | Pointer to **float32** |  | [optional] 
**CpuRequest** | Pointer to **float32** |  | [optional] 
**CpuUsage** | Pointer to **float32** |  | [optional] 
**DesiredPods** | Pointer to **int32** |  | [optional] 
**MemoryLimit** | Pointer to **float32** |  | [optional] 
**MemoryRequest** | Pointer to **float32** |  | [optional] 
**MemoryUsage** | Pointer to **float32** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**Restarts** | Pointer to **int32** |  | [optional] 
**StatefulSetName** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yStatefulSetListRecord

`func NewO11yStatefulSetListRecord() *O11yStatefulSetListRecord`

NewO11yStatefulSetListRecord instantiates a new O11yStatefulSetListRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yStatefulSetListRecordWithDefaults

`func NewO11yStatefulSetListRecordWithDefaults() *O11yStatefulSetListRecord`

NewO11yStatefulSetListRecordWithDefaults instantiates a new O11yStatefulSetListRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailablePods

`func (o *O11yStatefulSetListRecord) GetAvailablePods() int32`

GetAvailablePods returns the AvailablePods field if non-nil, zero value otherwise.

### GetAvailablePodsOk

`func (o *O11yStatefulSetListRecord) GetAvailablePodsOk() (*int32, bool)`

GetAvailablePodsOk returns a tuple with the AvailablePods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePods

`func (o *O11yStatefulSetListRecord) SetAvailablePods(v int32)`

SetAvailablePods sets AvailablePods field to given value.

### HasAvailablePods

`func (o *O11yStatefulSetListRecord) HasAvailablePods() bool`

HasAvailablePods returns a boolean if a field has been set.

### GetCpuLimit

`func (o *O11yStatefulSetListRecord) GetCpuLimit() float32`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *O11yStatefulSetListRecord) GetCpuLimitOk() (*float32, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *O11yStatefulSetListRecord) SetCpuLimit(v float32)`

SetCpuLimit sets CpuLimit field to given value.

### HasCpuLimit

`func (o *O11yStatefulSetListRecord) HasCpuLimit() bool`

HasCpuLimit returns a boolean if a field has been set.

### GetCpuRequest

`func (o *O11yStatefulSetListRecord) GetCpuRequest() float32`

GetCpuRequest returns the CpuRequest field if non-nil, zero value otherwise.

### GetCpuRequestOk

`func (o *O11yStatefulSetListRecord) GetCpuRequestOk() (*float32, bool)`

GetCpuRequestOk returns a tuple with the CpuRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuRequest

`func (o *O11yStatefulSetListRecord) SetCpuRequest(v float32)`

SetCpuRequest sets CpuRequest field to given value.

### HasCpuRequest

`func (o *O11yStatefulSetListRecord) HasCpuRequest() bool`

HasCpuRequest returns a boolean if a field has been set.

### GetCpuUsage

`func (o *O11yStatefulSetListRecord) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *O11yStatefulSetListRecord) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *O11yStatefulSetListRecord) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *O11yStatefulSetListRecord) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetDesiredPods

`func (o *O11yStatefulSetListRecord) GetDesiredPods() int32`

GetDesiredPods returns the DesiredPods field if non-nil, zero value otherwise.

### GetDesiredPodsOk

`func (o *O11yStatefulSetListRecord) GetDesiredPodsOk() (*int32, bool)`

GetDesiredPodsOk returns a tuple with the DesiredPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredPods

`func (o *O11yStatefulSetListRecord) SetDesiredPods(v int32)`

SetDesiredPods sets DesiredPods field to given value.

### HasDesiredPods

`func (o *O11yStatefulSetListRecord) HasDesiredPods() bool`

HasDesiredPods returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *O11yStatefulSetListRecord) GetMemoryLimit() float32`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *O11yStatefulSetListRecord) GetMemoryLimitOk() (*float32, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *O11yStatefulSetListRecord) SetMemoryLimit(v float32)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *O11yStatefulSetListRecord) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetMemoryRequest

`func (o *O11yStatefulSetListRecord) GetMemoryRequest() float32`

GetMemoryRequest returns the MemoryRequest field if non-nil, zero value otherwise.

### GetMemoryRequestOk

`func (o *O11yStatefulSetListRecord) GetMemoryRequestOk() (*float32, bool)`

GetMemoryRequestOk returns a tuple with the MemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequest

`func (o *O11yStatefulSetListRecord) SetMemoryRequest(v float32)`

SetMemoryRequest sets MemoryRequest field to given value.

### HasMemoryRequest

`func (o *O11yStatefulSetListRecord) HasMemoryRequest() bool`

HasMemoryRequest returns a boolean if a field has been set.

### GetMemoryUsage

`func (o *O11yStatefulSetListRecord) GetMemoryUsage() float32`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *O11yStatefulSetListRecord) GetMemoryUsageOk() (*float32, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *O11yStatefulSetListRecord) SetMemoryUsage(v float32)`

SetMemoryUsage sets MemoryUsage field to given value.

### HasMemoryUsage

`func (o *O11yStatefulSetListRecord) HasMemoryUsage() bool`

HasMemoryUsage returns a boolean if a field has been set.

### GetMeta

`func (o *O11yStatefulSetListRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yStatefulSetListRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yStatefulSetListRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yStatefulSetListRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRestarts

`func (o *O11yStatefulSetListRecord) GetRestarts() int32`

GetRestarts returns the Restarts field if non-nil, zero value otherwise.

### GetRestartsOk

`func (o *O11yStatefulSetListRecord) GetRestartsOk() (*int32, bool)`

GetRestartsOk returns a tuple with the Restarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestarts

`func (o *O11yStatefulSetListRecord) SetRestarts(v int32)`

SetRestarts sets Restarts field to given value.

### HasRestarts

`func (o *O11yStatefulSetListRecord) HasRestarts() bool`

HasRestarts returns a boolean if a field has been set.

### GetStatefulSetName

`func (o *O11yStatefulSetListRecord) GetStatefulSetName() string`

GetStatefulSetName returns the StatefulSetName field if non-nil, zero value otherwise.

### GetStatefulSetNameOk

`func (o *O11yStatefulSetListRecord) GetStatefulSetNameOk() (*string, bool)`

GetStatefulSetNameOk returns a tuple with the StatefulSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatefulSetName

`func (o *O11yStatefulSetListRecord) SetStatefulSetName(v string)`

SetStatefulSetName sets StatefulSetName field to given value.

### HasStatefulSetName

`func (o *O11yStatefulSetListRecord) HasStatefulSetName() bool`

HasStatefulSetName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


