# ComputeGpuJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | Pointer to **int64** | Attempt is which try this is, counting from 1. Above 1 means the job was retried after a failed or abandoned run. | [optional] 
**CloseTime** | Pointer to **string** | CloseTime is when the job reached a terminal state, RFC 3339. Empty means it is still live — queued, running or stalled. | [optional] 
**FailureCause** | Pointer to **string** | FailureCause is the engine&#39;s reason the job failed. Empty unless it did. | [optional] 
**Gpu** | Pointer to **string** | GPU is the node this job is aimed AT — the lane \&quot;gpu:&lt;node&gt;\&quot; it was submitted on. Empty means the shared any-GPU lane: it was not aimed anywhere and the first free worker takes it. | [optional] 
**Id** | Pointer to **string** | ID is the job&#39;s id, and the id the cancel route takes. The dispatcher sets it equal to the render&#39;s prompt id, so it is the same value the studio knows the job by. | [optional] 
**Label** | Pointer to **string** | Label is the cheap human name for the render — the output filename prefix lifted out of the submitted graph. Empty when the graph carried none. The graph itself is never in this list; the tasks describe endpoint serves it. | [optional] 
**LastHeartbeat** | Pointer to **string** | LastHeartbeat is the claiming worker&#39;s most recent beat on this job, RFC 3339 — the evidence a long render is still alive rather than wedged. | [optional] 
**LeaseExpiry** | Pointer to **string** | LeaseExpiry is when the worker&#39;s claim lapses, RFC 3339. Past it with the job still STARTED, the claimant is presumed dead and Status reads \&quot;stalled\&quot;. | [optional] 
**RunId** | Pointer to **string** | RunID identifies this execution of the job. It equals ID for a job the dispatcher submitted, which is why a cancel that omits it still works. | [optional] 
**StartTime** | Pointer to **string** | StartTime is when a worker began executing the job, RFC 3339. Empty while it is still queued. | [optional] 
**Status** | Pointer to **string** | Status is the job&#39;s lifecycle state: queued, running, completed, failed or canceled — plus \&quot;stalled\&quot;, which is this surface&#39;s own reading of a job that is STARTED whose worker died: its lease has elapsed and no reaper has taken it back yet. Without it such a job reads \&quot;running\&quot; forever. An engine state this surface does not recognize passes through lower-cased rather than being coerced into one of these. | [optional] 
**Type** | Pointer to **string** | Type is the work being done (\&quot;studio.render\&quot;) — what the claiming worker has to be able to execute. | [optional] 
**Worker** | Pointer to **string** | Worker is the node that actually CLAIMED the job, which is not always the one it was aimed at: a shared-lane job has no GPU but does have a Worker once picked up. Empty while the job is still waiting. | [optional] 

## Methods

### NewComputeGpuJob

`func NewComputeGpuJob() *ComputeGpuJob`

NewComputeGpuJob instantiates a new ComputeGpuJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeGpuJobWithDefaults

`func NewComputeGpuJobWithDefaults() *ComputeGpuJob`

NewComputeGpuJobWithDefaults instantiates a new ComputeGpuJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *ComputeGpuJob) GetAttempt() int64`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *ComputeGpuJob) GetAttemptOk() (*int64, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *ComputeGpuJob) SetAttempt(v int64)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *ComputeGpuJob) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetCloseTime

`func (o *ComputeGpuJob) GetCloseTime() string`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *ComputeGpuJob) GetCloseTimeOk() (*string, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *ComputeGpuJob) SetCloseTime(v string)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *ComputeGpuJob) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetFailureCause

`func (o *ComputeGpuJob) GetFailureCause() string`

GetFailureCause returns the FailureCause field if non-nil, zero value otherwise.

### GetFailureCauseOk

`func (o *ComputeGpuJob) GetFailureCauseOk() (*string, bool)`

GetFailureCauseOk returns a tuple with the FailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCause

`func (o *ComputeGpuJob) SetFailureCause(v string)`

SetFailureCause sets FailureCause field to given value.

### HasFailureCause

`func (o *ComputeGpuJob) HasFailureCause() bool`

HasFailureCause returns a boolean if a field has been set.

### GetGpu

`func (o *ComputeGpuJob) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *ComputeGpuJob) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *ComputeGpuJob) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *ComputeGpuJob) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetId

`func (o *ComputeGpuJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputeGpuJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputeGpuJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComputeGpuJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *ComputeGpuJob) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ComputeGpuJob) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ComputeGpuJob) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ComputeGpuJob) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLastHeartbeat

`func (o *ComputeGpuJob) GetLastHeartbeat() string`

GetLastHeartbeat returns the LastHeartbeat field if non-nil, zero value otherwise.

### GetLastHeartbeatOk

`func (o *ComputeGpuJob) GetLastHeartbeatOk() (*string, bool)`

GetLastHeartbeatOk returns a tuple with the LastHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeat

`func (o *ComputeGpuJob) SetLastHeartbeat(v string)`

SetLastHeartbeat sets LastHeartbeat field to given value.

### HasLastHeartbeat

`func (o *ComputeGpuJob) HasLastHeartbeat() bool`

HasLastHeartbeat returns a boolean if a field has been set.

### GetLeaseExpiry

`func (o *ComputeGpuJob) GetLeaseExpiry() string`

GetLeaseExpiry returns the LeaseExpiry field if non-nil, zero value otherwise.

### GetLeaseExpiryOk

`func (o *ComputeGpuJob) GetLeaseExpiryOk() (*string, bool)`

GetLeaseExpiryOk returns a tuple with the LeaseExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseExpiry

`func (o *ComputeGpuJob) SetLeaseExpiry(v string)`

SetLeaseExpiry sets LeaseExpiry field to given value.

### HasLeaseExpiry

`func (o *ComputeGpuJob) HasLeaseExpiry() bool`

HasLeaseExpiry returns a boolean if a field has been set.

### GetRunId

`func (o *ComputeGpuJob) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *ComputeGpuJob) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *ComputeGpuJob) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *ComputeGpuJob) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetStartTime

`func (o *ComputeGpuJob) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ComputeGpuJob) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ComputeGpuJob) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ComputeGpuJob) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ComputeGpuJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComputeGpuJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComputeGpuJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComputeGpuJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ComputeGpuJob) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComputeGpuJob) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComputeGpuJob) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ComputeGpuJob) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorker

`func (o *ComputeGpuJob) GetWorker() string`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *ComputeGpuJob) GetWorkerOk() (*string, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *ComputeGpuJob) SetWorker(v string)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *ComputeGpuJob) HasWorker() bool`

HasWorker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


