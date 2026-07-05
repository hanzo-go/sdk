# RegistryWebhookCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**EventTypes** | **[]string** |  | 
**Targets** | [**[]RegistryWebhookCreateTargetsInner**](RegistryWebhookCreateTargetsInner.md) |  | 

## Methods

### NewRegistryWebhookCreate

`func NewRegistryWebhookCreate(name string, eventTypes []string, targets []RegistryWebhookCreateTargetsInner, ) *RegistryWebhookCreate`

NewRegistryWebhookCreate instantiates a new RegistryWebhookCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryWebhookCreateWithDefaults

`func NewRegistryWebhookCreateWithDefaults() *RegistryWebhookCreate`

NewRegistryWebhookCreateWithDefaults instantiates a new RegistryWebhookCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RegistryWebhookCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistryWebhookCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistryWebhookCreate) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *RegistryWebhookCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RegistryWebhookCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RegistryWebhookCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RegistryWebhookCreate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventTypes

`func (o *RegistryWebhookCreate) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *RegistryWebhookCreate) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *RegistryWebhookCreate) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.


### GetTargets

`func (o *RegistryWebhookCreate) GetTargets() []RegistryWebhookCreateTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *RegistryWebhookCreate) GetTargetsOk() (*[]RegistryWebhookCreateTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *RegistryWebhookCreate) SetTargets(v []RegistryWebhookCreateTargetsInner)`

SetTargets sets Targets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


