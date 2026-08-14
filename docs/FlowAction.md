# FlowAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NextAction** | Pointer to [**FlowAction**](FlowAction.md) |  | [optional] 
**Settings** | Pointer to [**StepSettings**](StepSettings.md) |  | [optional] 
**Skip** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** | PIECE | CODE | ROUTER | LOOP_ON_ITEMS | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 

## Methods

### NewFlowAction

`func NewFlowAction() *FlowAction`

NewFlowAction instantiates a new FlowAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowActionWithDefaults

`func NewFlowActionWithDefaults() *FlowAction`

NewFlowActionWithDefaults instantiates a new FlowAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *FlowAction) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FlowAction) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FlowAction) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FlowAction) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *FlowAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlowAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextAction

`func (o *FlowAction) GetNextAction() FlowAction`

GetNextAction returns the NextAction field if non-nil, zero value otherwise.

### GetNextActionOk

`func (o *FlowAction) GetNextActionOk() (*FlowAction, bool)`

GetNextActionOk returns a tuple with the NextAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAction

`func (o *FlowAction) SetNextAction(v FlowAction)`

SetNextAction sets NextAction field to given value.

### HasNextAction

`func (o *FlowAction) HasNextAction() bool`

HasNextAction returns a boolean if a field has been set.

### GetSettings

`func (o *FlowAction) GetSettings() StepSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *FlowAction) GetSettingsOk() (*StepSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *FlowAction) SetSettings(v StepSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *FlowAction) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSkip

`func (o *FlowAction) GetSkip() bool`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *FlowAction) GetSkipOk() (*bool, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *FlowAction) SetSkip(v bool)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *FlowAction) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetType

`func (o *FlowAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlowAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlowAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FlowAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValid

`func (o *FlowAction) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *FlowAction) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *FlowAction) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *FlowAction) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


