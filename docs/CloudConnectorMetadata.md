# CloudConnectorMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]CloudConnectorAction**](CloudConnectorAction.md) |  | [optional] 
**Auth** | Pointer to [**CloudConnectorAuth**](CloudConnectorAuth.md) |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Triggers** | Pointer to [**[]CloudConnectorTrigger**](CloudConnectorTrigger.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudConnectorMetadata

`func NewCloudConnectorMetadata() *CloudConnectorMetadata`

NewCloudConnectorMetadata instantiates a new CloudConnectorMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudConnectorMetadataWithDefaults

`func NewCloudConnectorMetadataWithDefaults() *CloudConnectorMetadata`

NewCloudConnectorMetadataWithDefaults instantiates a new CloudConnectorMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *CloudConnectorMetadata) GetActions() []CloudConnectorAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CloudConnectorMetadata) GetActionsOk() (*[]CloudConnectorAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CloudConnectorMetadata) SetActions(v []CloudConnectorAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *CloudConnectorMetadata) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAuth

`func (o *CloudConnectorMetadata) GetAuth() CloudConnectorAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *CloudConnectorMetadata) GetAuthOk() (*CloudConnectorAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *CloudConnectorMetadata) SetAuth(v CloudConnectorAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *CloudConnectorMetadata) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetCategories

`func (o *CloudConnectorMetadata) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CloudConnectorMetadata) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CloudConnectorMetadata) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *CloudConnectorMetadata) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetDescription

`func (o *CloudConnectorMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudConnectorMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudConnectorMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudConnectorMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudConnectorMetadata) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudConnectorMetadata) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudConnectorMetadata) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudConnectorMetadata) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *CloudConnectorMetadata) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *CloudConnectorMetadata) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *CloudConnectorMetadata) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *CloudConnectorMetadata) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetName

`func (o *CloudConnectorMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudConnectorMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudConnectorMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudConnectorMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTriggers

`func (o *CloudConnectorMetadata) GetTriggers() []CloudConnectorTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *CloudConnectorMetadata) GetTriggersOk() (*[]CloudConnectorTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *CloudConnectorMetadata) SetTriggers(v []CloudConnectorTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *CloudConnectorMetadata) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetVersion

`func (o *CloudConnectorMetadata) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudConnectorMetadata) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudConnectorMetadata) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudConnectorMetadata) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


