# AutoFlowTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NextAction** | Pointer to [**AutoFlowAction**](AutoFlowAction.md) |  | [optional] 
**Settings** | Pointer to [**AutoStepSettings**](AutoStepSettings.md) |  | [optional] 
**Strategy** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | PIECE_TRIGGER | EMPTY | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 

## Methods

### NewAutoFlowTrigger

`func NewAutoFlowTrigger() *AutoFlowTrigger`

NewAutoFlowTrigger instantiates a new AutoFlowTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoFlowTriggerWithDefaults

`func NewAutoFlowTriggerWithDefaults() *AutoFlowTrigger`

NewAutoFlowTriggerWithDefaults instantiates a new AutoFlowTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *AutoFlowTrigger) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AutoFlowTrigger) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AutoFlowTrigger) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AutoFlowTrigger) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *AutoFlowTrigger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoFlowTrigger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoFlowTrigger) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoFlowTrigger) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextAction

`func (o *AutoFlowTrigger) GetNextAction() AutoFlowAction`

GetNextAction returns the NextAction field if non-nil, zero value otherwise.

### GetNextActionOk

`func (o *AutoFlowTrigger) GetNextActionOk() (*AutoFlowAction, bool)`

GetNextActionOk returns a tuple with the NextAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAction

`func (o *AutoFlowTrigger) SetNextAction(v AutoFlowAction)`

SetNextAction sets NextAction field to given value.

### HasNextAction

`func (o *AutoFlowTrigger) HasNextAction() bool`

HasNextAction returns a boolean if a field has been set.

### GetSettings

`func (o *AutoFlowTrigger) GetSettings() AutoStepSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AutoFlowTrigger) GetSettingsOk() (*AutoStepSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AutoFlowTrigger) SetSettings(v AutoStepSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *AutoFlowTrigger) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetStrategy

`func (o *AutoFlowTrigger) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *AutoFlowTrigger) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *AutoFlowTrigger) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *AutoFlowTrigger) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetType

`func (o *AutoFlowTrigger) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutoFlowTrigger) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutoFlowTrigger) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AutoFlowTrigger) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValid

`func (o *AutoFlowTrigger) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AutoFlowTrigger) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AutoFlowTrigger) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *AutoFlowTrigger) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


