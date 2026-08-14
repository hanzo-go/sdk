# GpuJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | Pointer to **int32** |  | [optional] 
**CloseTime** | Pointer to **string** |  | [optional] 
**FailureCause** | Pointer to **string** |  | [optional] 
**Gpu** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**LastHeartbeat** | Pointer to **string** |  | [optional] 
**LeaseExpiry** | Pointer to **string** |  | [optional] 
**RunId** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | queued|running|completed|failed|canceled | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Worker** | Pointer to **string** |  | [optional] 

## Methods

### NewGpuJob

`func NewGpuJob() *GpuJob`

NewGpuJob instantiates a new GpuJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpuJobWithDefaults

`func NewGpuJobWithDefaults() *GpuJob`

NewGpuJobWithDefaults instantiates a new GpuJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *GpuJob) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *GpuJob) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *GpuJob) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *GpuJob) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetCloseTime

`func (o *GpuJob) GetCloseTime() string`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *GpuJob) GetCloseTimeOk() (*string, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *GpuJob) SetCloseTime(v string)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *GpuJob) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetFailureCause

`func (o *GpuJob) GetFailureCause() string`

GetFailureCause returns the FailureCause field if non-nil, zero value otherwise.

### GetFailureCauseOk

`func (o *GpuJob) GetFailureCauseOk() (*string, bool)`

GetFailureCauseOk returns a tuple with the FailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCause

`func (o *GpuJob) SetFailureCause(v string)`

SetFailureCause sets FailureCause field to given value.

### HasFailureCause

`func (o *GpuJob) HasFailureCause() bool`

HasFailureCause returns a boolean if a field has been set.

### GetGpu

`func (o *GpuJob) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *GpuJob) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *GpuJob) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *GpuJob) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetId

`func (o *GpuJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GpuJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GpuJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GpuJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GpuJob) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GpuJob) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GpuJob) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GpuJob) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLastHeartbeat

`func (o *GpuJob) GetLastHeartbeat() string`

GetLastHeartbeat returns the LastHeartbeat field if non-nil, zero value otherwise.

### GetLastHeartbeatOk

`func (o *GpuJob) GetLastHeartbeatOk() (*string, bool)`

GetLastHeartbeatOk returns a tuple with the LastHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeat

`func (o *GpuJob) SetLastHeartbeat(v string)`

SetLastHeartbeat sets LastHeartbeat field to given value.

### HasLastHeartbeat

`func (o *GpuJob) HasLastHeartbeat() bool`

HasLastHeartbeat returns a boolean if a field has been set.

### GetLeaseExpiry

`func (o *GpuJob) GetLeaseExpiry() string`

GetLeaseExpiry returns the LeaseExpiry field if non-nil, zero value otherwise.

### GetLeaseExpiryOk

`func (o *GpuJob) GetLeaseExpiryOk() (*string, bool)`

GetLeaseExpiryOk returns a tuple with the LeaseExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseExpiry

`func (o *GpuJob) SetLeaseExpiry(v string)`

SetLeaseExpiry sets LeaseExpiry field to given value.

### HasLeaseExpiry

`func (o *GpuJob) HasLeaseExpiry() bool`

HasLeaseExpiry returns a boolean if a field has been set.

### GetRunId

`func (o *GpuJob) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *GpuJob) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *GpuJob) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *GpuJob) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetStartTime

`func (o *GpuJob) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GpuJob) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GpuJob) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GpuJob) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *GpuJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GpuJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GpuJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GpuJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *GpuJob) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GpuJob) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GpuJob) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GpuJob) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorker

`func (o *GpuJob) GetWorker() string`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *GpuJob) GetWorkerOk() (*string, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *GpuJob) SetWorker(v string)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *GpuJob) HasWorker() bool`

HasWorker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


