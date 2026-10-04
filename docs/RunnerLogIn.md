# RunnerLogIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int64** | Index is the position of the first line in that log, so a resent batch overlaps rather than duplicates. | [optional] 
**Last** | Pointer to **bool** | Last seals the log and moves it to storage. Appending past a seal is an error; resending one is acknowledged. | [optional] 
**Lines** | Pointer to [**[]RunnerLine**](RunnerLine.md) | Lines is the batch itself. | [optional] 
**Task** | Pointer to **int64** | Task is whose log this is. | [optional] 

## Methods

### NewRunnerLogIn

`func NewRunnerLogIn() *RunnerLogIn`

NewRunnerLogIn instantiates a new RunnerLogIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerLogInWithDefaults

`func NewRunnerLogInWithDefaults() *RunnerLogIn`

NewRunnerLogInWithDefaults instantiates a new RunnerLogIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *RunnerLogIn) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *RunnerLogIn) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *RunnerLogIn) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *RunnerLogIn) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetLast

`func (o *RunnerLogIn) GetLast() bool`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *RunnerLogIn) GetLastOk() (*bool, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *RunnerLogIn) SetLast(v bool)`

SetLast sets Last field to given value.

### HasLast

`func (o *RunnerLogIn) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLines

`func (o *RunnerLogIn) GetLines() []RunnerLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *RunnerLogIn) GetLinesOk() (*[]RunnerLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *RunnerLogIn) SetLines(v []RunnerLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *RunnerLogIn) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetTask

`func (o *RunnerLogIn) GetTask() int64`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *RunnerLogIn) GetTaskOk() (*int64, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *RunnerLogIn) SetTask(v int64)`

SetTask sets Task field to given value.

### HasTask

`func (o *RunnerLogIn) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


