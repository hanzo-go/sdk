# RunnerDeclareOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to **[]string** | Capabilities are what THIS FORGE understands, so a runner learns what the other side supports from the reply body rather than from a header. | [optional] 
**Runner** | Pointer to [**RunnerIdentity**](RunnerIdentity.md) | Runner is the stored identity as the forge now holds it. Its Token is empty: the credential is minted once, by register. | [optional] 

## Methods

### NewRunnerDeclareOut

`func NewRunnerDeclareOut() *RunnerDeclareOut`

NewRunnerDeclareOut instantiates a new RunnerDeclareOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerDeclareOutWithDefaults

`func NewRunnerDeclareOutWithDefaults() *RunnerDeclareOut`

NewRunnerDeclareOutWithDefaults instantiates a new RunnerDeclareOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *RunnerDeclareOut) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *RunnerDeclareOut) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *RunnerDeclareOut) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *RunnerDeclareOut) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetRunner

`func (o *RunnerDeclareOut) GetRunner() RunnerIdentity`

GetRunner returns the Runner field if non-nil, zero value otherwise.

### GetRunnerOk

`func (o *RunnerDeclareOut) GetRunnerOk() (*RunnerIdentity, bool)`

GetRunnerOk returns a tuple with the Runner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunner

`func (o *RunnerDeclareOut) SetRunner(v RunnerIdentity)`

SetRunner sets Runner field to given value.

### HasRunner

`func (o *RunnerDeclareOut) HasRunner() bool`

HasRunner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


