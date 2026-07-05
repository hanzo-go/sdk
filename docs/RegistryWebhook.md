# RegistryWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**EventTypes** | Pointer to **[]string** |  | [optional] 
**Targets** | Pointer to [**[]RegistryWebhookTargetsInner**](RegistryWebhookTargetsInner.md) |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRegistryWebhook

`func NewRegistryWebhook() *RegistryWebhook`

NewRegistryWebhook instantiates a new RegistryWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryWebhookWithDefaults

`func NewRegistryWebhookWithDefaults() *RegistryWebhook`

NewRegistryWebhookWithDefaults instantiates a new RegistryWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegistryWebhook) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistryWebhook) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistryWebhook) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RegistryWebhook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RegistryWebhook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistryWebhook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistryWebhook) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegistryWebhook) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *RegistryWebhook) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RegistryWebhook) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RegistryWebhook) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *RegistryWebhook) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetEnabled

`func (o *RegistryWebhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RegistryWebhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RegistryWebhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RegistryWebhook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventTypes

`func (o *RegistryWebhook) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *RegistryWebhook) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *RegistryWebhook) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *RegistryWebhook) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetTargets

`func (o *RegistryWebhook) GetTargets() []RegistryWebhookTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *RegistryWebhook) GetTargetsOk() (*[]RegistryWebhookTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *RegistryWebhook) SetTargets(v []RegistryWebhookTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *RegistryWebhook) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetCreationTime

`func (o *RegistryWebhook) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RegistryWebhook) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RegistryWebhook) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RegistryWebhook) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


