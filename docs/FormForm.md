# FormForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | Created is when the form was made, RFC 3339. | [optional] 
**Fields** | Pointer to [**[]FormField**](FormField.md) | Fields are the questions, in the order they are asked. | [optional] 
**Hidden** | Pointer to **[]string** | Hidden names the values a page passes in when it embeds the form (a video id, a plan): lowercase letters, digits and underscores, at most 32 each. | [optional] 
**Id** | Pointer to **string** | ID is the form&#39;s id: 15 lowercase hex characters, derived from the slug when the form has one, so a slug names one form for good. | [optional] 
**Kind** | Pointer to **string** | Kind is form (every submission is a response, the default) or subscribe (a mailing list: an email field is required and one address is one response). | [optional] 
**PublicId** | Pointer to **string** | PublicID is the address the form is answered at: \&quot;&lt;org&gt;.&lt;slug or id&gt;\&quot;, as GET /v1/form/public/{id} and POST /v1/form/public/{id}/responses take it. | [optional] 
**Settings** | Pointer to [**FormOptions**](FormOptions.md) | Settings is how the form is answered. | [optional] 
**Sinks** | Pointer to [**FormSinks**](FormSinks.md) | Sinks are where each completed response goes. Absent on a write, the platform&#39;s defaults are taken. | [optional] 
**Slug** | Pointer to **string** | Slug is the form&#39;s own name within its org: lowercase letters, digits and dashes, at most 64. Optional; once set it is the form&#39;s id, so it cannot change. | [optional] 
**Tags** | Pointer to **[]string** | Tags label the form for its org: at most 20, each at most 64 characters. | [optional] 
**Thanks** | Pointer to [**FormScreen**](FormScreen.md) | Thanks is the screen shown after a submission, or where the respondent is sent. Optional. | [optional] 
**Theme** | Pointer to **string** | Theme is the id of the theme the form is drawn in. Optional. | [optional] 
**Title** | Pointer to **string** | Title is the form&#39;s name, as a respondent and the CRM timeline read it. At most 200 characters. | [optional] 
**Updated** | Pointer to **string** | Updated is when it last changed, RFC 3339. | [optional] 
**Welcome** | Pointer to [**FormScreen**](FormScreen.md) | Welcome is the screen shown before the first question. Optional. | [optional] 

## Methods

### NewFormForm

`func NewFormForm() *FormForm`

NewFormForm instantiates a new FormForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFormWithDefaults

`func NewFormFormWithDefaults() *FormForm`

NewFormFormWithDefaults instantiates a new FormForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *FormForm) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FormForm) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FormForm) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FormForm) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFields

`func (o *FormForm) GetFields() []FormField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FormForm) GetFieldsOk() (*[]FormField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FormForm) SetFields(v []FormField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *FormForm) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetHidden

`func (o *FormForm) GetHidden() []string`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *FormForm) GetHiddenOk() (*[]string, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *FormForm) SetHidden(v []string)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *FormForm) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *FormForm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormForm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormForm) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormForm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *FormForm) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FormForm) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FormForm) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FormForm) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPublicId

`func (o *FormForm) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *FormForm) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *FormForm) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *FormForm) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetSettings

`func (o *FormForm) GetSettings() FormOptions`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *FormForm) GetSettingsOk() (*FormOptions, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *FormForm) SetSettings(v FormOptions)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *FormForm) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSinks

`func (o *FormForm) GetSinks() FormSinks`

GetSinks returns the Sinks field if non-nil, zero value otherwise.

### GetSinksOk

`func (o *FormForm) GetSinksOk() (*FormSinks, bool)`

GetSinksOk returns a tuple with the Sinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinks

`func (o *FormForm) SetSinks(v FormSinks)`

SetSinks sets Sinks field to given value.

### HasSinks

`func (o *FormForm) HasSinks() bool`

HasSinks returns a boolean if a field has been set.

### GetSlug

`func (o *FormForm) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *FormForm) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *FormForm) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *FormForm) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTags

`func (o *FormForm) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FormForm) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FormForm) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FormForm) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThanks

`func (o *FormForm) GetThanks() FormScreen`

GetThanks returns the Thanks field if non-nil, zero value otherwise.

### GetThanksOk

`func (o *FormForm) GetThanksOk() (*FormScreen, bool)`

GetThanksOk returns a tuple with the Thanks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThanks

`func (o *FormForm) SetThanks(v FormScreen)`

SetThanks sets Thanks field to given value.

### HasThanks

`func (o *FormForm) HasThanks() bool`

HasThanks returns a boolean if a field has been set.

### GetTheme

`func (o *FormForm) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *FormForm) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *FormForm) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *FormForm) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetTitle

`func (o *FormForm) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FormForm) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FormForm) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FormForm) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *FormForm) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FormForm) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FormForm) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FormForm) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetWelcome

`func (o *FormForm) GetWelcome() FormScreen`

GetWelcome returns the Welcome field if non-nil, zero value otherwise.

### GetWelcomeOk

`func (o *FormForm) GetWelcomeOk() (*FormScreen, bool)`

GetWelcomeOk returns a tuple with the Welcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcome

`func (o *FormForm) SetWelcome(v FormScreen)`

SetWelcome sets Welcome field to given value.

### HasWelcome

`func (o *FormForm) HasWelcome() bool`

HasWelcome returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


