# RunnerTaskIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queue** | Pointer to **int64** | Queue is the queue version the runner last saw. When it equals the forge&#39;s, nothing has been queued since. | [optional] 

## Methods

### NewRunnerTaskIn

`func NewRunnerTaskIn() *RunnerTaskIn`

NewRunnerTaskIn instantiates a new RunnerTaskIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerTaskInWithDefaults

`func NewRunnerTaskInWithDefaults() *RunnerTaskIn`

NewRunnerTaskInWithDefaults instantiates a new RunnerTaskIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueue

`func (o *RunnerTaskIn) GetQueue() int64`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *RunnerTaskIn) GetQueueOk() (*int64, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *RunnerTaskIn) SetQueue(v int64)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *RunnerTaskIn) HasQueue() bool`

HasQueue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


