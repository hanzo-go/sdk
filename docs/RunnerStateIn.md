# RunnerStateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outputs** | Pointer to [**[]RunnerPair**](RunnerPair.md) | Outputs are values the job has published since the last report. | [optional] 
**State** | Pointer to [**RunnerState**](RunnerState.md) | State is the task&#39;s progress and that of each of its steps. | [optional] 

## Methods

### NewRunnerStateIn

`func NewRunnerStateIn() *RunnerStateIn`

NewRunnerStateIn instantiates a new RunnerStateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerStateInWithDefaults

`func NewRunnerStateInWithDefaults() *RunnerStateIn`

NewRunnerStateInWithDefaults instantiates a new RunnerStateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutputs

`func (o *RunnerStateIn) GetOutputs() []RunnerPair`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *RunnerStateIn) GetOutputsOk() (*[]RunnerPair, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *RunnerStateIn) SetOutputs(v []RunnerPair)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *RunnerStateIn) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetState

`func (o *RunnerStateIn) GetState() RunnerState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RunnerStateIn) GetStateOk() (*RunnerState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RunnerStateIn) SetState(v RunnerState)`

SetState sets State field to given value.

### HasState

`func (o *RunnerStateIn) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


