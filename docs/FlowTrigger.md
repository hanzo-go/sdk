# FlowTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NextAction** | Pointer to [**FlowAction**](FlowAction.md) |  | [optional] 
**Settings** | Pointer to [**StepSettings**](StepSettings.md) |  | [optional] 
**Strategy** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | PIECE_TRIGGER | EMPTY | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 

## Methods

### NewFlowTrigger

`func NewFlowTrigger() *FlowTrigger`

NewFlowTrigger instantiates a new FlowTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowTriggerWithDefaults

`func NewFlowTriggerWithDefaults() *FlowTrigger`

NewFlowTriggerWithDefaults instantiates a new FlowTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *FlowTrigger) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FlowTrigger) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FlowTrigger) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FlowTrigger) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *FlowTrigger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowTrigger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowTrigger) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlowTrigger) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextAction

`func (o *FlowTrigger) GetNextAction() FlowAction`

GetNextAction returns the NextAction field if non-nil, zero value otherwise.

### GetNextActionOk

`func (o *FlowTrigger) GetNextActionOk() (*FlowAction, bool)`

GetNextActionOk returns a tuple with the NextAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAction

`func (o *FlowTrigger) SetNextAction(v FlowAction)`

SetNextAction sets NextAction field to given value.

### HasNextAction

`func (o *FlowTrigger) HasNextAction() bool`

HasNextAction returns a boolean if a field has been set.

### GetSettings

`func (o *FlowTrigger) GetSettings() StepSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *FlowTrigger) GetSettingsOk() (*StepSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *FlowTrigger) SetSettings(v StepSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *FlowTrigger) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetStrategy

`func (o *FlowTrigger) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *FlowTrigger) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *FlowTrigger) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *FlowTrigger) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetType

`func (o *FlowTrigger) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlowTrigger) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlowTrigger) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FlowTrigger) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValid

`func (o *FlowTrigger) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *FlowTrigger) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *FlowTrigger) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *FlowTrigger) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


