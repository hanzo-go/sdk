# App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicConfigOptions** | Pointer to [**[]ApplicationConfigOption**](ApplicationConfigOption.md) |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**ApplicationView**](ApplicationView.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Manifest** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewApp

`func NewApp() *App`

NewApp instantiates a new App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWithDefaults

`func NewAppWithDefaults() *App`

NewAppWithDefaults instantiates a new App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicConfigOptions

`func (o *App) GetBasicConfigOptions() []ApplicationConfigOption`

GetBasicConfigOptions returns the BasicConfigOptions field if non-nil, zero value otherwise.

### GetBasicConfigOptionsOk

`func (o *App) GetBasicConfigOptionsOk() (*[]ApplicationConfigOption, bool)`

GetBasicConfigOptionsOk returns a tuple with the BasicConfigOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicConfigOptions

`func (o *App) SetBasicConfigOptions(v []ApplicationConfigOption)`

SetBasicConfigOptions sets BasicConfigOptions field to given value.

### HasBasicConfigOptions

`func (o *App) HasBasicConfigOptions() bool`

HasBasicConfigOptions returns a boolean if a field has been set.

### GetCreatedTime

`func (o *App) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *App) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *App) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *App) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *App) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *App) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *App) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *App) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetails

`func (o *App) GetDetails() ApplicationView`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *App) GetDetailsOk() (*ApplicationView, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *App) SetDetails(v ApplicationView)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *App) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDisplayName

`func (o *App) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *App) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *App) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *App) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetManifest

`func (o *App) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *App) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *App) SetManifest(v string)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *App) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### GetName

`func (o *App) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *App) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *App) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *App) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *App) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *App) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *App) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *App) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOwner

`func (o *App) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *App) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *App) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *App) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParameters

`func (o *App) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *App) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *App) SetParameters(v string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *App) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetStatus

`func (o *App) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *App) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *App) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *App) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplate

`func (o *App) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *App) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *App) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *App) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *App) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *App) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *App) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *App) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetUrl

`func (o *App) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *App) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *App) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *App) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


