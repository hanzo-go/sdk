# RunnerIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ephemeral** | Pointer to **bool** | Ephemeral means the runner takes ONE job and exits, so the forge must not expect it back. | [optional] 
**Id** | Pointer to **int64** | ID is the forge&#39;s own row id for this runner. | [optional] 
**Labels** | Pointer to **[]string** | Labels are what a workflow&#39;s &#x60;runs-on:&#x60; selects this runner by. | [optional] 
**Name** | Pointer to **string** | Name is what a person calls this runner in the forge&#39;s UI. | [optional] 
**Token** | Pointer to **string** | Token is the secret that authenticates that handle. It is in cleartext here and nowhere else; the forge stores only a digest of it. | [optional] 
**Uuid** | Pointer to **string** | UUID is the handle the runner presents on every later op. | [optional] 
**Version** | Pointer to **string** | Version is the runner build the forge last heard from. | [optional] 

## Methods

### NewRunnerIdentity

`func NewRunnerIdentity() *RunnerIdentity`

NewRunnerIdentity instantiates a new RunnerIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerIdentityWithDefaults

`func NewRunnerIdentityWithDefaults() *RunnerIdentity`

NewRunnerIdentityWithDefaults instantiates a new RunnerIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEphemeral

`func (o *RunnerIdentity) GetEphemeral() bool`

GetEphemeral returns the Ephemeral field if non-nil, zero value otherwise.

### GetEphemeralOk

`func (o *RunnerIdentity) GetEphemeralOk() (*bool, bool)`

GetEphemeralOk returns a tuple with the Ephemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral

`func (o *RunnerIdentity) SetEphemeral(v bool)`

SetEphemeral sets Ephemeral field to given value.

### HasEphemeral

`func (o *RunnerIdentity) HasEphemeral() bool`

HasEphemeral returns a boolean if a field has been set.

### GetId

`func (o *RunnerIdentity) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunnerIdentity) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunnerIdentity) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RunnerIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabels

`func (o *RunnerIdentity) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RunnerIdentity) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RunnerIdentity) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RunnerIdentity) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *RunnerIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RunnerIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RunnerIdentity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RunnerIdentity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *RunnerIdentity) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RunnerIdentity) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RunnerIdentity) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RunnerIdentity) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUuid

`func (o *RunnerIdentity) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RunnerIdentity) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RunnerIdentity) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RunnerIdentity) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVersion

`func (o *RunnerIdentity) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RunnerIdentity) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RunnerIdentity) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RunnerIdentity) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


