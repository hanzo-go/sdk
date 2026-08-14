# O11yJobListRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePods** | Pointer to **int32** |  | [optional] 
**CpuLimit** | Pointer to **float32** |  | [optional] 
**CpuRequest** | Pointer to **float32** |  | [optional] 
**CpuUsage** | Pointer to **float32** |  | [optional] 
**DesiredSuccessfulPods** | Pointer to **int32** |  | [optional] 
**FailedPods** | Pointer to **int32** |  | [optional] 
**JobName** | Pointer to **string** |  | [optional] 
**MemoryLimit** | Pointer to **float32** |  | [optional] 
**MemoryRequest** | Pointer to **float32** |  | [optional] 
**MemoryUsage** | Pointer to **float32** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**Restarts** | Pointer to **int32** |  | [optional] 
**SuccessfulPods** | Pointer to **int32** |  | [optional] 

## Methods

### NewO11yJobListRecord

`func NewO11yJobListRecord() *O11yJobListRecord`

NewO11yJobListRecord instantiates a new O11yJobListRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yJobListRecordWithDefaults

`func NewO11yJobListRecordWithDefaults() *O11yJobListRecord`

NewO11yJobListRecordWithDefaults instantiates a new O11yJobListRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePods

`func (o *O11yJobListRecord) GetActivePods() int32`

GetActivePods returns the ActivePods field if non-nil, zero value otherwise.

### GetActivePodsOk

`func (o *O11yJobListRecord) GetActivePodsOk() (*int32, bool)`

GetActivePodsOk returns a tuple with the ActivePods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePods

`func (o *O11yJobListRecord) SetActivePods(v int32)`

SetActivePods sets ActivePods field to given value.

### HasActivePods

`func (o *O11yJobListRecord) HasActivePods() bool`

HasActivePods returns a boolean if a field has been set.

### GetCpuLimit

`func (o *O11yJobListRecord) GetCpuLimit() float32`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *O11yJobListRecord) GetCpuLimitOk() (*float32, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *O11yJobListRecord) SetCpuLimit(v float32)`

SetCpuLimit sets CpuLimit field to given value.

### HasCpuLimit

`func (o *O11yJobListRecord) HasCpuLimit() bool`

HasCpuLimit returns a boolean if a field has been set.

### GetCpuRequest

`func (o *O11yJobListRecord) GetCpuRequest() float32`

GetCpuRequest returns the CpuRequest field if non-nil, zero value otherwise.

### GetCpuRequestOk

`func (o *O11yJobListRecord) GetCpuRequestOk() (*float32, bool)`

GetCpuRequestOk returns a tuple with the CpuRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuRequest

`func (o *O11yJobListRecord) SetCpuRequest(v float32)`

SetCpuRequest sets CpuRequest field to given value.

### HasCpuRequest

`func (o *O11yJobListRecord) HasCpuRequest() bool`

HasCpuRequest returns a boolean if a field has been set.

### GetCpuUsage

`func (o *O11yJobListRecord) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *O11yJobListRecord) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *O11yJobListRecord) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *O11yJobListRecord) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetDesiredSuccessfulPods

`func (o *O11yJobListRecord) GetDesiredSuccessfulPods() int32`

GetDesiredSuccessfulPods returns the DesiredSuccessfulPods field if non-nil, zero value otherwise.

### GetDesiredSuccessfulPodsOk

`func (o *O11yJobListRecord) GetDesiredSuccessfulPodsOk() (*int32, bool)`

GetDesiredSuccessfulPodsOk returns a tuple with the DesiredSuccessfulPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredSuccessfulPods

`func (o *O11yJobListRecord) SetDesiredSuccessfulPods(v int32)`

SetDesiredSuccessfulPods sets DesiredSuccessfulPods field to given value.

### HasDesiredSuccessfulPods

`func (o *O11yJobListRecord) HasDesiredSuccessfulPods() bool`

HasDesiredSuccessfulPods returns a boolean if a field has been set.

### GetFailedPods

`func (o *O11yJobListRecord) GetFailedPods() int32`

GetFailedPods returns the FailedPods field if non-nil, zero value otherwise.

### GetFailedPodsOk

`func (o *O11yJobListRecord) GetFailedPodsOk() (*int32, bool)`

GetFailedPodsOk returns a tuple with the FailedPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedPods

`func (o *O11yJobListRecord) SetFailedPods(v int32)`

SetFailedPods sets FailedPods field to given value.

### HasFailedPods

`func (o *O11yJobListRecord) HasFailedPods() bool`

HasFailedPods returns a boolean if a field has been set.

### GetJobName

`func (o *O11yJobListRecord) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *O11yJobListRecord) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *O11yJobListRecord) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *O11yJobListRecord) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *O11yJobListRecord) GetMemoryLimit() float32`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *O11yJobListRecord) GetMemoryLimitOk() (*float32, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *O11yJobListRecord) SetMemoryLimit(v float32)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *O11yJobListRecord) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetMemoryRequest

`func (o *O11yJobListRecord) GetMemoryRequest() float32`

GetMemoryRequest returns the MemoryRequest field if non-nil, zero value otherwise.

### GetMemoryRequestOk

`func (o *O11yJobListRecord) GetMemoryRequestOk() (*float32, bool)`

GetMemoryRequestOk returns a tuple with the MemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequest

`func (o *O11yJobListRecord) SetMemoryRequest(v float32)`

SetMemoryRequest sets MemoryRequest field to given value.

### HasMemoryRequest

`func (o *O11yJobListRecord) HasMemoryRequest() bool`

HasMemoryRequest returns a boolean if a field has been set.

### GetMemoryUsage

`func (o *O11yJobListRecord) GetMemoryUsage() float32`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *O11yJobListRecord) GetMemoryUsageOk() (*float32, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *O11yJobListRecord) SetMemoryUsage(v float32)`

SetMemoryUsage sets MemoryUsage field to given value.

### HasMemoryUsage

`func (o *O11yJobListRecord) HasMemoryUsage() bool`

HasMemoryUsage returns a boolean if a field has been set.

### GetMeta

`func (o *O11yJobListRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yJobListRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yJobListRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yJobListRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRestarts

`func (o *O11yJobListRecord) GetRestarts() int32`

GetRestarts returns the Restarts field if non-nil, zero value otherwise.

### GetRestartsOk

`func (o *O11yJobListRecord) GetRestartsOk() (*int32, bool)`

GetRestartsOk returns a tuple with the Restarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestarts

`func (o *O11yJobListRecord) SetRestarts(v int32)`

SetRestarts sets Restarts field to given value.

### HasRestarts

`func (o *O11yJobListRecord) HasRestarts() bool`

HasRestarts returns a boolean if a field has been set.

### GetSuccessfulPods

`func (o *O11yJobListRecord) GetSuccessfulPods() int32`

GetSuccessfulPods returns the SuccessfulPods field if non-nil, zero value otherwise.

### GetSuccessfulPodsOk

`func (o *O11yJobListRecord) GetSuccessfulPodsOk() (*int32, bool)`

GetSuccessfulPodsOk returns a tuple with the SuccessfulPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulPods

`func (o *O11yJobListRecord) SetSuccessfulPods(v int32)`

SetSuccessfulPods sets SuccessfulPods field to given value.

### HasSuccessfulPods

`func (o *O11yJobListRecord) HasSuccessfulPods() bool`

HasSuccessfulPods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


