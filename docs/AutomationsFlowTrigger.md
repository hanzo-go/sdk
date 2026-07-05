# AutomationsFlowTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**Strategy** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**AutomationsStepSettings**](AutomationsStepSettings.md) |  | [optional] 
**NextAction** | Pointer to [**AutomationsFlowAction**](AutomationsFlowAction.md) |  | [optional] 

## Methods

### NewAutomationsFlowTrigger

`func NewAutomationsFlowTrigger() *AutomationsFlowTrigger`

NewAutomationsFlowTrigger instantiates a new AutomationsFlowTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationsFlowTriggerWithDefaults

`func NewAutomationsFlowTriggerWithDefaults() *AutomationsFlowTrigger`

NewAutomationsFlowTriggerWithDefaults instantiates a new AutomationsFlowTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AutomationsFlowTrigger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutomationsFlowTrigger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutomationsFlowTrigger) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutomationsFlowTrigger) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AutomationsFlowTrigger) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutomationsFlowTrigger) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutomationsFlowTrigger) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AutomationsFlowTrigger) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayName

`func (o *AutomationsFlowTrigger) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AutomationsFlowTrigger) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AutomationsFlowTrigger) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AutomationsFlowTrigger) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetValid

`func (o *AutomationsFlowTrigger) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AutomationsFlowTrigger) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AutomationsFlowTrigger) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *AutomationsFlowTrigger) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetStrategy

`func (o *AutomationsFlowTrigger) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *AutomationsFlowTrigger) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *AutomationsFlowTrigger) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *AutomationsFlowTrigger) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetSettings

`func (o *AutomationsFlowTrigger) GetSettings() AutomationsStepSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AutomationsFlowTrigger) GetSettingsOk() (*AutomationsStepSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AutomationsFlowTrigger) SetSettings(v AutomationsStepSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *AutomationsFlowTrigger) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetNextAction

`func (o *AutomationsFlowTrigger) GetNextAction() AutomationsFlowAction`

GetNextAction returns the NextAction field if non-nil, zero value otherwise.

### GetNextActionOk

`func (o *AutomationsFlowTrigger) GetNextActionOk() (*AutomationsFlowAction, bool)`

GetNextActionOk returns a tuple with the NextAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAction

`func (o *AutomationsFlowTrigger) SetNextAction(v AutomationsFlowAction)`

SetNextAction sets NextAction field to given value.

### HasNextAction

`func (o *AutomationsFlowTrigger) HasNextAction() bool`

HasNextAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


