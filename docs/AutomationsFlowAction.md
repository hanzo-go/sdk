# AutomationsFlowAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**Skip** | Pointer to **bool** |  | [optional] 
**Settings** | Pointer to [**AutomationsStepSettings**](AutomationsStepSettings.md) |  | [optional] 
**NextAction** | Pointer to [**AutomationsFlowAction**](AutomationsFlowAction.md) |  | [optional] 

## Methods

### NewAutomationsFlowAction

`func NewAutomationsFlowAction() *AutomationsFlowAction`

NewAutomationsFlowAction instantiates a new AutomationsFlowAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationsFlowActionWithDefaults

`func NewAutomationsFlowActionWithDefaults() *AutomationsFlowAction`

NewAutomationsFlowActionWithDefaults instantiates a new AutomationsFlowAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AutomationsFlowAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutomationsFlowAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutomationsFlowAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutomationsFlowAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AutomationsFlowAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutomationsFlowAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutomationsFlowAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AutomationsFlowAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayName

`func (o *AutomationsFlowAction) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AutomationsFlowAction) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AutomationsFlowAction) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AutomationsFlowAction) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetValid

`func (o *AutomationsFlowAction) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AutomationsFlowAction) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AutomationsFlowAction) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *AutomationsFlowAction) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetSkip

`func (o *AutomationsFlowAction) GetSkip() bool`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *AutomationsFlowAction) GetSkipOk() (*bool, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *AutomationsFlowAction) SetSkip(v bool)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *AutomationsFlowAction) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetSettings

`func (o *AutomationsFlowAction) GetSettings() AutomationsStepSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AutomationsFlowAction) GetSettingsOk() (*AutomationsStepSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AutomationsFlowAction) SetSettings(v AutomationsStepSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *AutomationsFlowAction) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetNextAction

`func (o *AutomationsFlowAction) GetNextAction() AutomationsFlowAction`

GetNextAction returns the NextAction field if non-nil, zero value otherwise.

### GetNextActionOk

`func (o *AutomationsFlowAction) GetNextActionOk() (*AutomationsFlowAction, bool)`

GetNextActionOk returns a tuple with the NextAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAction

`func (o *AutomationsFlowAction) SetNextAction(v AutomationsFlowAction)`

SetNextAction sets NextAction field to given value.

### HasNextAction

`func (o *AutomationsFlowAction) HasNextAction() bool`

HasNextAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


