# RunnerTaskOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queue** | Pointer to **int64** | Queue is the forge&#39;s current queue version, to ask against next time. | [optional] 
**Task** | Pointer to [**RunnerTask**](RunnerTask.md) | Task is the assigned job, or nil for \&quot;no work\&quot;. | [optional] 

## Methods

### NewRunnerTaskOut

`func NewRunnerTaskOut() *RunnerTaskOut`

NewRunnerTaskOut instantiates a new RunnerTaskOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerTaskOutWithDefaults

`func NewRunnerTaskOutWithDefaults() *RunnerTaskOut`

NewRunnerTaskOutWithDefaults instantiates a new RunnerTaskOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueue

`func (o *RunnerTaskOut) GetQueue() int64`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *RunnerTaskOut) GetQueueOk() (*int64, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *RunnerTaskOut) SetQueue(v int64)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *RunnerTaskOut) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetTask

`func (o *RunnerTaskOut) GetTask() RunnerTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *RunnerTaskOut) GetTaskOk() (*RunnerTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *RunnerTaskOut) SetTask(v RunnerTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *RunnerTaskOut) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


