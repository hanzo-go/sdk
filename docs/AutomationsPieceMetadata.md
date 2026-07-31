# AutomationsPieceMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**Auth** | Pointer to [**AutomationsPieceAuth**](AutomationsPieceAuth.md) |  | [optional] 
**Actions** | Pointer to [**[]AutomationsPieceAction**](AutomationsPieceAction.md) |  | [optional] 
**Triggers** | Pointer to [**[]AutomationsPieceTrigger**](AutomationsPieceTrigger.md) |  | [optional] 

## Methods

### NewAutomationsPieceMetadata

`func NewAutomationsPieceMetadata() *AutomationsPieceMetadata`

NewAutomationsPieceMetadata instantiates a new AutomationsPieceMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationsPieceMetadataWithDefaults

`func NewAutomationsPieceMetadataWithDefaults() *AutomationsPieceMetadata`

NewAutomationsPieceMetadataWithDefaults instantiates a new AutomationsPieceMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AutomationsPieceMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutomationsPieceMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutomationsPieceMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutomationsPieceMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *AutomationsPieceMetadata) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AutomationsPieceMetadata) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AutomationsPieceMetadata) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AutomationsPieceMetadata) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *AutomationsPieceMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutomationsPieceMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutomationsPieceMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutomationsPieceMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogoUrl

`func (o *AutomationsPieceMetadata) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *AutomationsPieceMetadata) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *AutomationsPieceMetadata) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *AutomationsPieceMetadata) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetVersion

`func (o *AutomationsPieceMetadata) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AutomationsPieceMetadata) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AutomationsPieceMetadata) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AutomationsPieceMetadata) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCategories

`func (o *AutomationsPieceMetadata) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AutomationsPieceMetadata) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AutomationsPieceMetadata) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *AutomationsPieceMetadata) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetAuth

`func (o *AutomationsPieceMetadata) GetAuth() AutomationsPieceAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *AutomationsPieceMetadata) GetAuthOk() (*AutomationsPieceAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *AutomationsPieceMetadata) SetAuth(v AutomationsPieceAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *AutomationsPieceMetadata) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetActions

`func (o *AutomationsPieceMetadata) GetActions() []AutomationsPieceAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AutomationsPieceMetadata) GetActionsOk() (*[]AutomationsPieceAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AutomationsPieceMetadata) SetActions(v []AutomationsPieceAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AutomationsPieceMetadata) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetTriggers

`func (o *AutomationsPieceMetadata) GetTriggers() []AutomationsPieceTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *AutomationsPieceMetadata) GetTriggersOk() (*[]AutomationsPieceTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *AutomationsPieceMetadata) SetTriggers(v []AutomationsPieceTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *AutomationsPieceMetadata) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


