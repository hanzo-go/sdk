# FormFormPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**[]FormField**](FormField.md) |  | [optional] 
**Hidden** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** | ID is the form&#39;s id or its slug, from the path. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**FormOptions**](FormOptions.md) |  | [optional] 
**Sinks** | Pointer to [**FormSinks**](FormSinks.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Thanks** | Pointer to [**FormScreen**](FormScreen.md) |  | [optional] 
**Theme** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Welcome** | Pointer to [**FormScreen**](FormScreen.md) |  | [optional] 

## Methods

### NewFormFormPatch

`func NewFormFormPatch() *FormFormPatch`

NewFormFormPatch instantiates a new FormFormPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFormPatchWithDefaults

`func NewFormFormPatchWithDefaults() *FormFormPatch`

NewFormFormPatchWithDefaults instantiates a new FormFormPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *FormFormPatch) GetFields() []FormField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FormFormPatch) GetFieldsOk() (*[]FormField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FormFormPatch) SetFields(v []FormField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *FormFormPatch) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetHidden

`func (o *FormFormPatch) GetHidden() []string`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *FormFormPatch) GetHiddenOk() (*[]string, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *FormFormPatch) SetHidden(v []string)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *FormFormPatch) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *FormFormPatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormFormPatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormFormPatch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormFormPatch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *FormFormPatch) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FormFormPatch) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FormFormPatch) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FormFormPatch) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSettings

`func (o *FormFormPatch) GetSettings() FormOptions`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *FormFormPatch) GetSettingsOk() (*FormOptions, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *FormFormPatch) SetSettings(v FormOptions)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *FormFormPatch) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSinks

`func (o *FormFormPatch) GetSinks() FormSinks`

GetSinks returns the Sinks field if non-nil, zero value otherwise.

### GetSinksOk

`func (o *FormFormPatch) GetSinksOk() (*FormSinks, bool)`

GetSinksOk returns a tuple with the Sinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinks

`func (o *FormFormPatch) SetSinks(v FormSinks)`

SetSinks sets Sinks field to given value.

### HasSinks

`func (o *FormFormPatch) HasSinks() bool`

HasSinks returns a boolean if a field has been set.

### GetTags

`func (o *FormFormPatch) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FormFormPatch) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FormFormPatch) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FormFormPatch) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThanks

`func (o *FormFormPatch) GetThanks() FormScreen`

GetThanks returns the Thanks field if non-nil, zero value otherwise.

### GetThanksOk

`func (o *FormFormPatch) GetThanksOk() (*FormScreen, bool)`

GetThanksOk returns a tuple with the Thanks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThanks

`func (o *FormFormPatch) SetThanks(v FormScreen)`

SetThanks sets Thanks field to given value.

### HasThanks

`func (o *FormFormPatch) HasThanks() bool`

HasThanks returns a boolean if a field has been set.

### GetTheme

`func (o *FormFormPatch) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *FormFormPatch) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *FormFormPatch) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *FormFormPatch) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetTitle

`func (o *FormFormPatch) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FormFormPatch) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FormFormPatch) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FormFormPatch) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWelcome

`func (o *FormFormPatch) GetWelcome() FormScreen`

GetWelcome returns the Welcome field if non-nil, zero value otherwise.

### GetWelcomeOk

`func (o *FormFormPatch) GetWelcomeOk() (*FormScreen, bool)`

GetWelcomeOk returns a tuple with the Welcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcome

`func (o *FormFormPatch) SetWelcome(v FormScreen)`

SetWelcome sets Welcome field to given value.

### HasWelcome

`func (o *FormFormPatch) HasWelcome() bool`

HasWelcome returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


