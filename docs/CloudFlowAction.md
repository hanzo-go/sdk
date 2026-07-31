# CloudFlowAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NextAction** | Pointer to [**CloudFlowAction**](CloudFlowAction.md) |  | [optional] 
**Settings** | Pointer to [**CloudStepSettings**](CloudStepSettings.md) |  | [optional] 
**Skip** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** | PIECE | CODE | ROUTER | LOOP_ON_ITEMS | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 

## Methods

### NewCloudFlowAction

`func NewCloudFlowAction() *CloudFlowAction`

NewCloudFlowAction instantiates a new CloudFlowAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFlowActionWithDefaults

`func NewCloudFlowActionWithDefaults() *CloudFlowAction`

NewCloudFlowActionWithDefaults instantiates a new CloudFlowAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CloudFlowAction) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudFlowAction) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudFlowAction) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudFlowAction) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *CloudFlowAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudFlowAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudFlowAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudFlowAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextAction

`func (o *CloudFlowAction) GetNextAction() CloudFlowAction`

GetNextAction returns the NextAction field if non-nil, zero value otherwise.

### GetNextActionOk

`func (o *CloudFlowAction) GetNextActionOk() (*CloudFlowAction, bool)`

GetNextActionOk returns a tuple with the NextAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAction

`func (o *CloudFlowAction) SetNextAction(v CloudFlowAction)`

SetNextAction sets NextAction field to given value.

### HasNextAction

`func (o *CloudFlowAction) HasNextAction() bool`

HasNextAction returns a boolean if a field has been set.

### GetSettings

`func (o *CloudFlowAction) GetSettings() CloudStepSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CloudFlowAction) GetSettingsOk() (*CloudStepSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CloudFlowAction) SetSettings(v CloudStepSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CloudFlowAction) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSkip

`func (o *CloudFlowAction) GetSkip() bool`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *CloudFlowAction) GetSkipOk() (*bool, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *CloudFlowAction) SetSkip(v bool)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *CloudFlowAction) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetType

`func (o *CloudFlowAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudFlowAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudFlowAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudFlowAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValid

`func (o *CloudFlowAction) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *CloudFlowAction) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *CloudFlowAction) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *CloudFlowAction) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


