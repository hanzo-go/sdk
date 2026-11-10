# FormPublicForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**[]FormField**](FormField.md) | Fields are the questions, in order. | [optional] 
**Hidden** | Pointer to **[]string** | Hidden names the values the page passes in. | [optional] 
**Id** | Pointer to **string** | ID is the form&#39;s public id. | [optional] 
**Kind** | Pointer to **string** | Kind is form or subscribe. | [optional] 
**Settings** | Pointer to [**FormDrawn**](FormDrawn.md) | Settings are how the form is answered. | [optional] 
**Thanks** | Pointer to [**FormScreen**](FormScreen.md) | Thanks is the screen after a submission, or where the respondent is sent. | [optional] 
**Theme** | Pointer to [**FormTheme**](FormTheme.md) | Theme is the theme the form is drawn in, when it names one. | [optional] 
**Title** | Pointer to **string** | Title is the form&#39;s title. | [optional] 
**Welcome** | Pointer to [**FormScreen**](FormScreen.md) | Welcome is the screen before the first question. | [optional] 

## Methods

### NewFormPublicForm

`func NewFormPublicForm() *FormPublicForm`

NewFormPublicForm instantiates a new FormPublicForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormPublicFormWithDefaults

`func NewFormPublicFormWithDefaults() *FormPublicForm`

NewFormPublicFormWithDefaults instantiates a new FormPublicForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *FormPublicForm) GetFields() []FormField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FormPublicForm) GetFieldsOk() (*[]FormField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FormPublicForm) SetFields(v []FormField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *FormPublicForm) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetHidden

`func (o *FormPublicForm) GetHidden() []string`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *FormPublicForm) GetHiddenOk() (*[]string, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *FormPublicForm) SetHidden(v []string)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *FormPublicForm) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *FormPublicForm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormPublicForm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormPublicForm) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormPublicForm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *FormPublicForm) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FormPublicForm) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FormPublicForm) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FormPublicForm) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSettings

`func (o *FormPublicForm) GetSettings() FormDrawn`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *FormPublicForm) GetSettingsOk() (*FormDrawn, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *FormPublicForm) SetSettings(v FormDrawn)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *FormPublicForm) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetThanks

`func (o *FormPublicForm) GetThanks() FormScreen`

GetThanks returns the Thanks field if non-nil, zero value otherwise.

### GetThanksOk

`func (o *FormPublicForm) GetThanksOk() (*FormScreen, bool)`

GetThanksOk returns a tuple with the Thanks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThanks

`func (o *FormPublicForm) SetThanks(v FormScreen)`

SetThanks sets Thanks field to given value.

### HasThanks

`func (o *FormPublicForm) HasThanks() bool`

HasThanks returns a boolean if a field has been set.

### GetTheme

`func (o *FormPublicForm) GetTheme() FormTheme`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *FormPublicForm) GetThemeOk() (*FormTheme, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *FormPublicForm) SetTheme(v FormTheme)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *FormPublicForm) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetTitle

`func (o *FormPublicForm) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FormPublicForm) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FormPublicForm) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FormPublicForm) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWelcome

`func (o *FormPublicForm) GetWelcome() FormScreen`

GetWelcome returns the Welcome field if non-nil, zero value otherwise.

### GetWelcomeOk

`func (o *FormPublicForm) GetWelcomeOk() (*FormScreen, bool)`

GetWelcomeOk returns a tuple with the Welcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcome

`func (o *FormPublicForm) SetWelcome(v FormScreen)`

SetWelcome sets Welcome field to given value.

### HasWelcome

`func (o *FormPublicForm) HasWelcome() bool`

HasWelcome returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


