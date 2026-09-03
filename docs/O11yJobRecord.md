# O11yJobRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePods** | Pointer to **int64** |  | [optional] 
**DesiredSuccessfulPods** | Pointer to **int64** |  | [optional] 
**FailedPods** | Pointer to **int64** |  | [optional] 
**JobCPU** | Pointer to **float64** |  | [optional] 
**JobCPULimit** | Pointer to **float64** |  | [optional] 
**JobCPURequest** | Pointer to **float64** |  | [optional] 
**JobMemory** | Pointer to **float64** |  | [optional] 
**JobMemoryLimit** | Pointer to **float64** |  | [optional] 
**JobMemoryRequest** | Pointer to **float64** |  | [optional] 
**JobName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**PodCountsByPhase** | Pointer to [**O11yPodCountsByPhase**](O11yPodCountsByPhase.md) |  | [optional] 
**SuccessfulPods** | Pointer to **int64** |  | [optional] 

## Methods

### NewO11yJobRecord

`func NewO11yJobRecord() *O11yJobRecord`

NewO11yJobRecord instantiates a new O11yJobRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yJobRecordWithDefaults

`func NewO11yJobRecordWithDefaults() *O11yJobRecord`

NewO11yJobRecordWithDefaults instantiates a new O11yJobRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePods

`func (o *O11yJobRecord) GetActivePods() int64`

GetActivePods returns the ActivePods field if non-nil, zero value otherwise.

### GetActivePodsOk

`func (o *O11yJobRecord) GetActivePodsOk() (*int64, bool)`

GetActivePodsOk returns a tuple with the ActivePods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePods

`func (o *O11yJobRecord) SetActivePods(v int64)`

SetActivePods sets ActivePods field to given value.

### HasActivePods

`func (o *O11yJobRecord) HasActivePods() bool`

HasActivePods returns a boolean if a field has been set.

### GetDesiredSuccessfulPods

`func (o *O11yJobRecord) GetDesiredSuccessfulPods() int64`

GetDesiredSuccessfulPods returns the DesiredSuccessfulPods field if non-nil, zero value otherwise.

### GetDesiredSuccessfulPodsOk

`func (o *O11yJobRecord) GetDesiredSuccessfulPodsOk() (*int64, bool)`

GetDesiredSuccessfulPodsOk returns a tuple with the DesiredSuccessfulPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredSuccessfulPods

`func (o *O11yJobRecord) SetDesiredSuccessfulPods(v int64)`

SetDesiredSuccessfulPods sets DesiredSuccessfulPods field to given value.

### HasDesiredSuccessfulPods

`func (o *O11yJobRecord) HasDesiredSuccessfulPods() bool`

HasDesiredSuccessfulPods returns a boolean if a field has been set.

### GetFailedPods

`func (o *O11yJobRecord) GetFailedPods() int64`

GetFailedPods returns the FailedPods field if non-nil, zero value otherwise.

### GetFailedPodsOk

`func (o *O11yJobRecord) GetFailedPodsOk() (*int64, bool)`

GetFailedPodsOk returns a tuple with the FailedPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedPods

`func (o *O11yJobRecord) SetFailedPods(v int64)`

SetFailedPods sets FailedPods field to given value.

### HasFailedPods

`func (o *O11yJobRecord) HasFailedPods() bool`

HasFailedPods returns a boolean if a field has been set.

### GetJobCPU

`func (o *O11yJobRecord) GetJobCPU() float64`

GetJobCPU returns the JobCPU field if non-nil, zero value otherwise.

### GetJobCPUOk

`func (o *O11yJobRecord) GetJobCPUOk() (*float64, bool)`

GetJobCPUOk returns a tuple with the JobCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobCPU

`func (o *O11yJobRecord) SetJobCPU(v float64)`

SetJobCPU sets JobCPU field to given value.

### HasJobCPU

`func (o *O11yJobRecord) HasJobCPU() bool`

HasJobCPU returns a boolean if a field has been set.

### GetJobCPULimit

`func (o *O11yJobRecord) GetJobCPULimit() float64`

GetJobCPULimit returns the JobCPULimit field if non-nil, zero value otherwise.

### GetJobCPULimitOk

`func (o *O11yJobRecord) GetJobCPULimitOk() (*float64, bool)`

GetJobCPULimitOk returns a tuple with the JobCPULimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobCPULimit

`func (o *O11yJobRecord) SetJobCPULimit(v float64)`

SetJobCPULimit sets JobCPULimit field to given value.

### HasJobCPULimit

`func (o *O11yJobRecord) HasJobCPULimit() bool`

HasJobCPULimit returns a boolean if a field has been set.

### GetJobCPURequest

`func (o *O11yJobRecord) GetJobCPURequest() float64`

GetJobCPURequest returns the JobCPURequest field if non-nil, zero value otherwise.

### GetJobCPURequestOk

`func (o *O11yJobRecord) GetJobCPURequestOk() (*float64, bool)`

GetJobCPURequestOk returns a tuple with the JobCPURequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobCPURequest

`func (o *O11yJobRecord) SetJobCPURequest(v float64)`

SetJobCPURequest sets JobCPURequest field to given value.

### HasJobCPURequest

`func (o *O11yJobRecord) HasJobCPURequest() bool`

HasJobCPURequest returns a boolean if a field has been set.

### GetJobMemory

`func (o *O11yJobRecord) GetJobMemory() float64`

GetJobMemory returns the JobMemory field if non-nil, zero value otherwise.

### GetJobMemoryOk

`func (o *O11yJobRecord) GetJobMemoryOk() (*float64, bool)`

GetJobMemoryOk returns a tuple with the JobMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobMemory

`func (o *O11yJobRecord) SetJobMemory(v float64)`

SetJobMemory sets JobMemory field to given value.

### HasJobMemory

`func (o *O11yJobRecord) HasJobMemory() bool`

HasJobMemory returns a boolean if a field has been set.

### GetJobMemoryLimit

`func (o *O11yJobRecord) GetJobMemoryLimit() float64`

GetJobMemoryLimit returns the JobMemoryLimit field if non-nil, zero value otherwise.

### GetJobMemoryLimitOk

`func (o *O11yJobRecord) GetJobMemoryLimitOk() (*float64, bool)`

GetJobMemoryLimitOk returns a tuple with the JobMemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobMemoryLimit

`func (o *O11yJobRecord) SetJobMemoryLimit(v float64)`

SetJobMemoryLimit sets JobMemoryLimit field to given value.

### HasJobMemoryLimit

`func (o *O11yJobRecord) HasJobMemoryLimit() bool`

HasJobMemoryLimit returns a boolean if a field has been set.

### GetJobMemoryRequest

`func (o *O11yJobRecord) GetJobMemoryRequest() float64`

GetJobMemoryRequest returns the JobMemoryRequest field if non-nil, zero value otherwise.

### GetJobMemoryRequestOk

`func (o *O11yJobRecord) GetJobMemoryRequestOk() (*float64, bool)`

GetJobMemoryRequestOk returns a tuple with the JobMemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobMemoryRequest

`func (o *O11yJobRecord) SetJobMemoryRequest(v float64)`

SetJobMemoryRequest sets JobMemoryRequest field to given value.

### HasJobMemoryRequest

`func (o *O11yJobRecord) HasJobMemoryRequest() bool`

HasJobMemoryRequest returns a boolean if a field has been set.

### GetJobName

`func (o *O11yJobRecord) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *O11yJobRecord) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *O11yJobRecord) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *O11yJobRecord) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### GetMeta

`func (o *O11yJobRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yJobRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yJobRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yJobRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPodCountsByPhase

`func (o *O11yJobRecord) GetPodCountsByPhase() O11yPodCountsByPhase`

GetPodCountsByPhase returns the PodCountsByPhase field if non-nil, zero value otherwise.

### GetPodCountsByPhaseOk

`func (o *O11yJobRecord) GetPodCountsByPhaseOk() (*O11yPodCountsByPhase, bool)`

GetPodCountsByPhaseOk returns a tuple with the PodCountsByPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCountsByPhase

`func (o *O11yJobRecord) SetPodCountsByPhase(v O11yPodCountsByPhase)`

SetPodCountsByPhase sets PodCountsByPhase field to given value.

### HasPodCountsByPhase

`func (o *O11yJobRecord) HasPodCountsByPhase() bool`

HasPodCountsByPhase returns a boolean if a field has been set.

### GetSuccessfulPods

`func (o *O11yJobRecord) GetSuccessfulPods() int64`

GetSuccessfulPods returns the SuccessfulPods field if non-nil, zero value otherwise.

### GetSuccessfulPodsOk

`func (o *O11yJobRecord) GetSuccessfulPodsOk() (*int64, bool)`

GetSuccessfulPodsOk returns a tuple with the SuccessfulPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulPods

`func (o *O11yJobRecord) SetSuccessfulPods(v int64)`

SetSuccessfulPods sets SuccessfulPods field to given value.

### HasSuccessfulPods

`func (o *O11yJobRecord) HasSuccessfulPods() bool`

HasSuccessfulPods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


