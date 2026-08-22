# TierLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedModels** | Pointer to **[]string** |  | [optional] 
**DailyCreditsCents** | Pointer to **int32** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**MaxAgents** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UnlimitedAgents** | Pointer to **bool** | UnlimitedAgents reports that MaxAgents 0 means \&quot;no ceiling\&quot; rather than \&quot;no agents\&quot; — the reading a bare zero cannot carry. | [optional] 

## Methods

### NewTierLimits

`func NewTierLimits() *TierLimits`

NewTierLimits instantiates a new TierLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierLimitsWithDefaults

`func NewTierLimitsWithDefaults() *TierLimits`

NewTierLimitsWithDefaults instantiates a new TierLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedModels

`func (o *TierLimits) GetAllowedModels() []string`

GetAllowedModels returns the AllowedModels field if non-nil, zero value otherwise.

### GetAllowedModelsOk

`func (o *TierLimits) GetAllowedModelsOk() (*[]string, bool)`

GetAllowedModelsOk returns a tuple with the AllowedModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedModels

`func (o *TierLimits) SetAllowedModels(v []string)`

SetAllowedModels sets AllowedModels field to given value.

### HasAllowedModels

`func (o *TierLimits) HasAllowedModels() bool`

HasAllowedModels returns a boolean if a field has been set.

### GetDailyCreditsCents

`func (o *TierLimits) GetDailyCreditsCents() int32`

GetDailyCreditsCents returns the DailyCreditsCents field if non-nil, zero value otherwise.

### GetDailyCreditsCentsOk

`func (o *TierLimits) GetDailyCreditsCentsOk() (*int32, bool)`

GetDailyCreditsCentsOk returns a tuple with the DailyCreditsCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyCreditsCents

`func (o *TierLimits) SetDailyCreditsCents(v int32)`

SetDailyCreditsCents sets DailyCreditsCents field to given value.

### HasDailyCreditsCents

`func (o *TierLimits) HasDailyCreditsCents() bool`

HasDailyCreditsCents returns a boolean if a field has been set.

### GetDisplayName

`func (o *TierLimits) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TierLimits) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TierLimits) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TierLimits) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMaxAgents

`func (o *TierLimits) GetMaxAgents() int32`

GetMaxAgents returns the MaxAgents field if non-nil, zero value otherwise.

### GetMaxAgentsOk

`func (o *TierLimits) GetMaxAgentsOk() (*int32, bool)`

GetMaxAgentsOk returns a tuple with the MaxAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAgents

`func (o *TierLimits) SetMaxAgents(v int32)`

SetMaxAgents sets MaxAgents field to given value.

### HasMaxAgents

`func (o *TierLimits) HasMaxAgents() bool`

HasMaxAgents returns a boolean if a field has been set.

### GetName

`func (o *TierLimits) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TierLimits) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TierLimits) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TierLimits) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUnlimitedAgents

`func (o *TierLimits) GetUnlimitedAgents() bool`

GetUnlimitedAgents returns the UnlimitedAgents field if non-nil, zero value otherwise.

### GetUnlimitedAgentsOk

`func (o *TierLimits) GetUnlimitedAgentsOk() (*bool, bool)`

GetUnlimitedAgentsOk returns a tuple with the UnlimitedAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlimitedAgents

`func (o *TierLimits) SetUnlimitedAgents(v bool)`

SetUnlimitedAgents sets UnlimitedAgents field to given value.

### HasUnlimitedAgents

`func (o *TierLimits) HasUnlimitedAgents() bool`

HasUnlimitedAgents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


