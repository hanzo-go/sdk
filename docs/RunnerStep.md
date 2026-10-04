# RunnerStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | ID is the step&#39;s position in the job, counting from 0. | [optional] 
**LogIndex** | Pointer to **int64** | LogIndex is the first line of the task log this step wrote. | [optional] 
**LogLength** | Pointer to **int64** | LogLength is how many lines it wrote from there. | [optional] 
**Result** | Pointer to **string** | Result is how the step finished; empty means it has not. | [optional] 
**Started** | Pointer to **int64** | Started is when it began, in unix nanoseconds; 0 is unset. | [optional] 
**Stopped** | Pointer to **int64** | Stopped is when it finished, in unix nanoseconds; 0 is unset. | [optional] 

## Methods

### NewRunnerStep

`func NewRunnerStep() *RunnerStep`

NewRunnerStep instantiates a new RunnerStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerStepWithDefaults

`func NewRunnerStepWithDefaults() *RunnerStep`

NewRunnerStepWithDefaults instantiates a new RunnerStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RunnerStep) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunnerStep) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunnerStep) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RunnerStep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogIndex

`func (o *RunnerStep) GetLogIndex() int64`

GetLogIndex returns the LogIndex field if non-nil, zero value otherwise.

### GetLogIndexOk

`func (o *RunnerStep) GetLogIndexOk() (*int64, bool)`

GetLogIndexOk returns a tuple with the LogIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogIndex

`func (o *RunnerStep) SetLogIndex(v int64)`

SetLogIndex sets LogIndex field to given value.

### HasLogIndex

`func (o *RunnerStep) HasLogIndex() bool`

HasLogIndex returns a boolean if a field has been set.

### GetLogLength

`func (o *RunnerStep) GetLogLength() int64`

GetLogLength returns the LogLength field if non-nil, zero value otherwise.

### GetLogLengthOk

`func (o *RunnerStep) GetLogLengthOk() (*int64, bool)`

GetLogLengthOk returns a tuple with the LogLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLength

`func (o *RunnerStep) SetLogLength(v int64)`

SetLogLength sets LogLength field to given value.

### HasLogLength

`func (o *RunnerStep) HasLogLength() bool`

HasLogLength returns a boolean if a field has been set.

### GetResult

`func (o *RunnerStep) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RunnerStep) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RunnerStep) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *RunnerStep) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStarted

`func (o *RunnerStep) GetStarted() int64`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *RunnerStep) GetStartedOk() (*int64, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *RunnerStep) SetStarted(v int64)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *RunnerStep) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetStopped

`func (o *RunnerStep) GetStopped() int64`

GetStopped returns the Stopped field if non-nil, zero value otherwise.

### GetStoppedOk

`func (o *RunnerStep) GetStoppedOk() (*int64, bool)`

GetStoppedOk returns a tuple with the Stopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopped

`func (o *RunnerStep) SetStopped(v int64)`

SetStopped sets Stopped field to given value.

### HasStopped

`func (o *RunnerStep) HasStopped() bool`

HasStopped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


