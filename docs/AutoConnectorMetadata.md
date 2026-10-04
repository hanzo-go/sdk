# AutoConnectorMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]AutoConnectorAction**](AutoConnectorAction.md) |  | [optional] 
**Auth** | Pointer to [**AutoConnectorAuth**](AutoConnectorAuth.md) |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Runnable** | Pointer to **bool** | Runnable is whether this connector&#39;s actions run HERE, so a caller can invoke what the entry lists. It is the registry&#39;s answer — an action registered for the connector — never a stored flag, because a stored flag is a second answer to a question the registry already answers. | [optional] 
**Triggers** | Pointer to [**[]AutoConnectorTrigger**](AutoConnectorTrigger.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewAutoConnectorMetadata

`func NewAutoConnectorMetadata() *AutoConnectorMetadata`

NewAutoConnectorMetadata instantiates a new AutoConnectorMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoConnectorMetadataWithDefaults

`func NewAutoConnectorMetadataWithDefaults() *AutoConnectorMetadata`

NewAutoConnectorMetadataWithDefaults instantiates a new AutoConnectorMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *AutoConnectorMetadata) GetActions() []AutoConnectorAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AutoConnectorMetadata) GetActionsOk() (*[]AutoConnectorAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AutoConnectorMetadata) SetActions(v []AutoConnectorAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AutoConnectorMetadata) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAuth

`func (o *AutoConnectorMetadata) GetAuth() AutoConnectorAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *AutoConnectorMetadata) GetAuthOk() (*AutoConnectorAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *AutoConnectorMetadata) SetAuth(v AutoConnectorAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *AutoConnectorMetadata) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetCategories

`func (o *AutoConnectorMetadata) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AutoConnectorMetadata) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AutoConnectorMetadata) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *AutoConnectorMetadata) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetDescription

`func (o *AutoConnectorMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoConnectorMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoConnectorMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoConnectorMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *AutoConnectorMetadata) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AutoConnectorMetadata) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AutoConnectorMetadata) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AutoConnectorMetadata) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *AutoConnectorMetadata) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *AutoConnectorMetadata) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *AutoConnectorMetadata) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *AutoConnectorMetadata) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetName

`func (o *AutoConnectorMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoConnectorMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoConnectorMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoConnectorMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRunnable

`func (o *AutoConnectorMetadata) GetRunnable() bool`

GetRunnable returns the Runnable field if non-nil, zero value otherwise.

### GetRunnableOk

`func (o *AutoConnectorMetadata) GetRunnableOk() (*bool, bool)`

GetRunnableOk returns a tuple with the Runnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnable

`func (o *AutoConnectorMetadata) SetRunnable(v bool)`

SetRunnable sets Runnable field to given value.

### HasRunnable

`func (o *AutoConnectorMetadata) HasRunnable() bool`

HasRunnable returns a boolean if a field has been set.

### GetTriggers

`func (o *AutoConnectorMetadata) GetTriggers() []AutoConnectorTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *AutoConnectorMetadata) GetTriggersOk() (*[]AutoConnectorTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *AutoConnectorMetadata) SetTriggers(v []AutoConnectorTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *AutoConnectorMetadata) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetVersion

`func (o *AutoConnectorMetadata) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AutoConnectorMetadata) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AutoConnectorMetadata) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AutoConnectorMetadata) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


