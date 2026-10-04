# RunnerRegisterIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to **[]string** | Capabilities are the optional protocol behaviours this runner understands, e.g. \&quot;cancelling\&quot;. | [optional] 
**Ephemeral** | Pointer to **bool** | Ephemeral declares that this runner takes ONE job and exits. | [optional] 
**Labels** | Pointer to **[]string** | Labels are what a workflow&#39;s &#x60;runs-on:&#x60; will select this runner by. | [optional] 
**Name** | Pointer to **string** | Name is what to call this runner. Required. | [optional] 
**Token** | Pointer to **string** | Token is the REGISTRATION secret, a different secret from the runner token the reply carries. Required. | [optional] 
**Version** | Pointer to **string** | Version is the runner build asking to register. | [optional] 

## Methods

### NewRunnerRegisterIn

`func NewRunnerRegisterIn() *RunnerRegisterIn`

NewRunnerRegisterIn instantiates a new RunnerRegisterIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerRegisterInWithDefaults

`func NewRunnerRegisterInWithDefaults() *RunnerRegisterIn`

NewRunnerRegisterInWithDefaults instantiates a new RunnerRegisterIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *RunnerRegisterIn) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *RunnerRegisterIn) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *RunnerRegisterIn) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *RunnerRegisterIn) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetEphemeral

`func (o *RunnerRegisterIn) GetEphemeral() bool`

GetEphemeral returns the Ephemeral field if non-nil, zero value otherwise.

### GetEphemeralOk

`func (o *RunnerRegisterIn) GetEphemeralOk() (*bool, bool)`

GetEphemeralOk returns a tuple with the Ephemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral

`func (o *RunnerRegisterIn) SetEphemeral(v bool)`

SetEphemeral sets Ephemeral field to given value.

### HasEphemeral

`func (o *RunnerRegisterIn) HasEphemeral() bool`

HasEphemeral returns a boolean if a field has been set.

### GetLabels

`func (o *RunnerRegisterIn) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RunnerRegisterIn) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RunnerRegisterIn) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RunnerRegisterIn) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *RunnerRegisterIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RunnerRegisterIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RunnerRegisterIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RunnerRegisterIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *RunnerRegisterIn) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RunnerRegisterIn) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RunnerRegisterIn) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RunnerRegisterIn) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetVersion

`func (o *RunnerRegisterIn) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RunnerRegisterIn) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RunnerRegisterIn) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RunnerRegisterIn) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


