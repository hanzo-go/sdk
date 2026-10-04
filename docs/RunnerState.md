# RunnerState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | ID is the task this is about. | [optional] 
**Result** | Pointer to **string** | Result is how it finished; empty means it has not. | [optional] 
**Started** | Pointer to **int64** | Started is when it began, in unix nanoseconds; 0 is unset. | [optional] 
**Steps** | Pointer to [**[]RunnerStep**](RunnerStep.md) | Steps is each step&#39;s own progress, in the order the workflow declares. | [optional] 
**Stopped** | Pointer to **int64** | Stopped is when it finished, in unix nanoseconds; 0 is unset. | [optional] 

## Methods

### NewRunnerState

`func NewRunnerState() *RunnerState`

NewRunnerState instantiates a new RunnerState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerStateWithDefaults

`func NewRunnerStateWithDefaults() *RunnerState`

NewRunnerStateWithDefaults instantiates a new RunnerState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RunnerState) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunnerState) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunnerState) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RunnerState) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResult

`func (o *RunnerState) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RunnerState) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RunnerState) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *RunnerState) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStarted

`func (o *RunnerState) GetStarted() int64`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *RunnerState) GetStartedOk() (*int64, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *RunnerState) SetStarted(v int64)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *RunnerState) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetSteps

`func (o *RunnerState) GetSteps() []RunnerStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *RunnerState) GetStepsOk() (*[]RunnerStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *RunnerState) SetSteps(v []RunnerStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *RunnerState) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetStopped

`func (o *RunnerState) GetStopped() int64`

GetStopped returns the Stopped field if non-nil, zero value otherwise.

### GetStoppedOk

`func (o *RunnerState) GetStoppedOk() (*int64, bool)`

GetStoppedOk returns a tuple with the Stopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopped

`func (o *RunnerState) SetStopped(v int64)`

SetStopped sets Stopped field to given value.

### HasStopped

`func (o *RunnerState) HasStopped() bool`

HasStopped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


