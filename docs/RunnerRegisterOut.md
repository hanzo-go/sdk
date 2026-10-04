# RunnerRegisterOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Runner** | Pointer to [**RunnerIdentity**](RunnerIdentity.md) | Runner is the minted identity, carrying the token in cleartext once. | [optional] 

## Methods

### NewRunnerRegisterOut

`func NewRunnerRegisterOut() *RunnerRegisterOut`

NewRunnerRegisterOut instantiates a new RunnerRegisterOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerRegisterOutWithDefaults

`func NewRunnerRegisterOutWithDefaults() *RunnerRegisterOut`

NewRunnerRegisterOutWithDefaults instantiates a new RunnerRegisterOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunner

`func (o *RunnerRegisterOut) GetRunner() RunnerIdentity`

GetRunner returns the Runner field if non-nil, zero value otherwise.

### GetRunnerOk

`func (o *RunnerRegisterOut) GetRunnerOk() (*RunnerIdentity, bool)`

GetRunnerOk returns a tuple with the Runner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunner

`func (o *RunnerRegisterOut) SetRunner(v RunnerIdentity)`

SetRunner sets Runner field to given value.

### HasRunner

`func (o *RunnerRegisterOut) HasRunner() bool`

HasRunner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


