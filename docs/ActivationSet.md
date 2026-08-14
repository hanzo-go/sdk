# ActivationSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **[]string** | Enabled is every tool name activated for the caller&#39;s org and project. | [optional] 

## Methods

### NewActivationSet

`func NewActivationSet() *ActivationSet`

NewActivationSet instantiates a new ActivationSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivationSetWithDefaults

`func NewActivationSetWithDefaults() *ActivationSet`

NewActivationSetWithDefaults instantiates a new ActivationSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ActivationSet) GetEnabled() []string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ActivationSet) GetEnabledOk() (*[]string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ActivationSet) SetEnabled(v []string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ActivationSet) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


