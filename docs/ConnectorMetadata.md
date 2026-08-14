# ConnectorMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]ConnectorAction**](ConnectorAction.md) |  | [optional] 
**Auth** | Pointer to [**ConnectorAuth**](ConnectorAuth.md) |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Triggers** | Pointer to [**[]ConnectorTrigger**](ConnectorTrigger.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewConnectorMetadata

`func NewConnectorMetadata() *ConnectorMetadata`

NewConnectorMetadata instantiates a new ConnectorMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorMetadataWithDefaults

`func NewConnectorMetadataWithDefaults() *ConnectorMetadata`

NewConnectorMetadataWithDefaults instantiates a new ConnectorMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ConnectorMetadata) GetActions() []ConnectorAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ConnectorMetadata) GetActionsOk() (*[]ConnectorAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ConnectorMetadata) SetActions(v []ConnectorAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ConnectorMetadata) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAuth

`func (o *ConnectorMetadata) GetAuth() ConnectorAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *ConnectorMetadata) GetAuthOk() (*ConnectorAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *ConnectorMetadata) SetAuth(v ConnectorAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *ConnectorMetadata) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetCategories

`func (o *ConnectorMetadata) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ConnectorMetadata) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ConnectorMetadata) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ConnectorMetadata) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetDescription

`func (o *ConnectorMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectorMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectorMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectorMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ConnectorMetadata) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConnectorMetadata) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConnectorMetadata) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ConnectorMetadata) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *ConnectorMetadata) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *ConnectorMetadata) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *ConnectorMetadata) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *ConnectorMetadata) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetName

`func (o *ConnectorMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTriggers

`func (o *ConnectorMetadata) GetTriggers() []ConnectorTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *ConnectorMetadata) GetTriggersOk() (*[]ConnectorTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *ConnectorMetadata) SetTriggers(v []ConnectorTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *ConnectorMetadata) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetVersion

`func (o *ConnectorMetadata) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConnectorMetadata) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConnectorMetadata) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConnectorMetadata) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


