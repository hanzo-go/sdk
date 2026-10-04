# RunnerStateOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**RunnerState**](RunnerState.md) | State is the task&#39;s result AS THE FORGE HOLDS IT, which may differ from what was reported: that is how a runner learns it was cancelled. | [optional] 
**Stored** | Pointer to **[]string** | Stored names the outputs the forge has written down, so the runner stops resending them. | [optional] 

## Methods

### NewRunnerStateOut

`func NewRunnerStateOut() *RunnerStateOut`

NewRunnerStateOut instantiates a new RunnerStateOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerStateOutWithDefaults

`func NewRunnerStateOutWithDefaults() *RunnerStateOut`

NewRunnerStateOutWithDefaults instantiates a new RunnerStateOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *RunnerStateOut) GetState() RunnerState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RunnerStateOut) GetStateOk() (*RunnerState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RunnerStateOut) SetState(v RunnerState)`

SetState sets State field to given value.

### HasState

`func (o *RunnerStateOut) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStored

`func (o *RunnerStateOut) GetStored() []string`

GetStored returns the Stored field if non-nil, zero value otherwise.

### GetStoredOk

`func (o *RunnerStateOut) GetStoredOk() (*[]string, bool)`

GetStoredOk returns a tuple with the Stored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStored

`func (o *RunnerStateOut) SetStored(v []string)`

SetStored sets Stored field to given value.

### HasStored

`func (o *RunnerStateOut) HasStored() bool`

HasStored returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


