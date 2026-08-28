# Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicConfigOptions** | Pointer to [**[]TemplateConfigOption**](TemplateConfigOption.md) |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EnableBasicConfig** | Pointer to **bool** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Manifest** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Readme** | Pointer to **string** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewTemplate

`func NewTemplate() *Template`

NewTemplate instantiates a new Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateWithDefaults

`func NewTemplateWithDefaults() *Template`

NewTemplateWithDefaults instantiates a new Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicConfigOptions

`func (o *Template) GetBasicConfigOptions() []TemplateConfigOption`

GetBasicConfigOptions returns the BasicConfigOptions field if non-nil, zero value otherwise.

### GetBasicConfigOptionsOk

`func (o *Template) GetBasicConfigOptionsOk() (*[]TemplateConfigOption, bool)`

GetBasicConfigOptionsOk returns a tuple with the BasicConfigOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicConfigOptions

`func (o *Template) SetBasicConfigOptions(v []TemplateConfigOption)`

SetBasicConfigOptions sets BasicConfigOptions field to given value.

### HasBasicConfigOptions

`func (o *Template) HasBasicConfigOptions() bool`

HasBasicConfigOptions returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Template) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Template) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Template) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Template) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *Template) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Template) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Template) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Template) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *Template) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Template) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Template) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Template) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableBasicConfig

`func (o *Template) GetEnableBasicConfig() bool`

GetEnableBasicConfig returns the EnableBasicConfig field if non-nil, zero value otherwise.

### GetEnableBasicConfigOk

`func (o *Template) GetEnableBasicConfigOk() (*bool, bool)`

GetEnableBasicConfigOk returns a tuple with the EnableBasicConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBasicConfig

`func (o *Template) SetEnableBasicConfig(v bool)`

SetEnableBasicConfig sets EnableBasicConfig field to given value.

### HasEnableBasicConfig

`func (o *Template) HasEnableBasicConfig() bool`

HasEnableBasicConfig returns a boolean if a field has been set.

### GetIcon

`func (o *Template) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Template) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Template) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Template) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetManifest

`func (o *Template) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *Template) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *Template) SetManifest(v string)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *Template) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### GetName

`func (o *Template) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Template) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Template) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Template) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *Template) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Template) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Template) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Template) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetReadme

`func (o *Template) GetReadme() string`

GetReadme returns the Readme field if non-nil, zero value otherwise.

### GetReadmeOk

`func (o *Template) GetReadmeOk() (*string, bool)`

GetReadmeOk returns a tuple with the Readme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadme

`func (o *Template) SetReadme(v string)`

SetReadme sets Readme field to given value.

### HasReadme

`func (o *Template) HasReadme() bool`

HasReadme returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *Template) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *Template) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *Template) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *Template) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetVersion

`func (o *Template) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Template) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Template) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Template) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


