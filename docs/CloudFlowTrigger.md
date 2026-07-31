# CloudFlowTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NextAction** | Pointer to [**CloudFlowAction**](CloudFlowAction.md) |  | [optional] 
**Settings** | Pointer to [**CloudStepSettings**](CloudStepSettings.md) |  | [optional] 
**Strategy** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | PIECE_TRIGGER | EMPTY | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 

## Methods

### NewCloudFlowTrigger

`func NewCloudFlowTrigger() *CloudFlowTrigger`

NewCloudFlowTrigger instantiates a new CloudFlowTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFlowTriggerWithDefaults

`func NewCloudFlowTriggerWithDefaults() *CloudFlowTrigger`

NewCloudFlowTriggerWithDefaults instantiates a new CloudFlowTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CloudFlowTrigger) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudFlowTrigger) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudFlowTrigger) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudFlowTrigger) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *CloudFlowTrigger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudFlowTrigger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudFlowTrigger) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudFlowTrigger) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextAction

`func (o *CloudFlowTrigger) GetNextAction() CloudFlowAction`

GetNextAction returns the NextAction field if non-nil, zero value otherwise.

### GetNextActionOk

`func (o *CloudFlowTrigger) GetNextActionOk() (*CloudFlowAction, bool)`

GetNextActionOk returns a tuple with the NextAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAction

`func (o *CloudFlowTrigger) SetNextAction(v CloudFlowAction)`

SetNextAction sets NextAction field to given value.

### HasNextAction

`func (o *CloudFlowTrigger) HasNextAction() bool`

HasNextAction returns a boolean if a field has been set.

### GetSettings

`func (o *CloudFlowTrigger) GetSettings() CloudStepSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CloudFlowTrigger) GetSettingsOk() (*CloudStepSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CloudFlowTrigger) SetSettings(v CloudStepSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CloudFlowTrigger) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetStrategy

`func (o *CloudFlowTrigger) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *CloudFlowTrigger) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *CloudFlowTrigger) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *CloudFlowTrigger) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetType

`func (o *CloudFlowTrigger) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudFlowTrigger) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudFlowTrigger) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudFlowTrigger) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValid

`func (o *CloudFlowTrigger) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *CloudFlowTrigger) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *CloudFlowTrigger) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *CloudFlowTrigger) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


