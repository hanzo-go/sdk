# RunnerDeclareIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to **[]string** | Capabilities are the optional protocol behaviours this runner understands, e.g. \&quot;cancelling\&quot;. | [optional] 
**Labels** | Pointer to **[]string** | Labels replace what the forge holds, so a relabelled runner follows its configuration without re-registering. | [optional] 
**Version** | Pointer to **string** | Version is the runner build now running. | [optional] 

## Methods

### NewRunnerDeclareIn

`func NewRunnerDeclareIn() *RunnerDeclareIn`

NewRunnerDeclareIn instantiates a new RunnerDeclareIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerDeclareInWithDefaults

`func NewRunnerDeclareInWithDefaults() *RunnerDeclareIn`

NewRunnerDeclareInWithDefaults instantiates a new RunnerDeclareIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *RunnerDeclareIn) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *RunnerDeclareIn) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *RunnerDeclareIn) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *RunnerDeclareIn) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetLabels

`func (o *RunnerDeclareIn) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RunnerDeclareIn) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RunnerDeclareIn) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RunnerDeclareIn) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetVersion

`func (o *RunnerDeclareIn) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RunnerDeclareIn) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RunnerDeclareIn) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RunnerDeclareIn) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


