# AutoFlowAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NextAction** | Pointer to [**AutoFlowAction**](AutoFlowAction.md) |  | [optional] 
**Settings** | Pointer to [**AutoStepSettings**](AutoStepSettings.md) |  | [optional] 
**Skip** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** | PIECE | CODE | ROUTER | LOOP_ON_ITEMS | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 

## Methods

### NewAutoFlowAction

`func NewAutoFlowAction() *AutoFlowAction`

NewAutoFlowAction instantiates a new AutoFlowAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoFlowActionWithDefaults

`func NewAutoFlowActionWithDefaults() *AutoFlowAction`

NewAutoFlowActionWithDefaults instantiates a new AutoFlowAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *AutoFlowAction) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AutoFlowAction) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AutoFlowAction) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AutoFlowAction) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *AutoFlowAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoFlowAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoFlowAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoFlowAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextAction

`func (o *AutoFlowAction) GetNextAction() AutoFlowAction`

GetNextAction returns the NextAction field if non-nil, zero value otherwise.

### GetNextActionOk

`func (o *AutoFlowAction) GetNextActionOk() (*AutoFlowAction, bool)`

GetNextActionOk returns a tuple with the NextAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAction

`func (o *AutoFlowAction) SetNextAction(v AutoFlowAction)`

SetNextAction sets NextAction field to given value.

### HasNextAction

`func (o *AutoFlowAction) HasNextAction() bool`

HasNextAction returns a boolean if a field has been set.

### GetSettings

`func (o *AutoFlowAction) GetSettings() AutoStepSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AutoFlowAction) GetSettingsOk() (*AutoStepSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AutoFlowAction) SetSettings(v AutoStepSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *AutoFlowAction) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSkip

`func (o *AutoFlowAction) GetSkip() bool`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *AutoFlowAction) GetSkipOk() (*bool, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *AutoFlowAction) SetSkip(v bool)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *AutoFlowAction) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetType

`func (o *AutoFlowAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutoFlowAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutoFlowAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AutoFlowAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValid

`func (o *AutoFlowAction) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AutoFlowAction) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AutoFlowAction) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *AutoFlowAction) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


