# FormTheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Background** | Pointer to **string** | Background is an https address of an image drawn behind the form. Optional. | [optional] 
**Colors** | Pointer to [**FormColors**](FormColors.md) | Colors are the theme&#39;s colours, each #rrggbb. | [optional] 
**Created** | Pointer to **string** | Created is when the theme was made, RFC 3339. | [optional] 
**Font** | Pointer to **string** | Font is the font family the form is set in: letters, digits, spaces and dashes, at most 64. Empty is the renderer&#39;s own. | [optional] 
**Id** | Pointer to **string** | ID is the theme&#39;s id. | [optional] 
**Name** | Pointer to **string** | Name is the theme&#39;s name in its org. At most 100 characters. | [optional] 
**Updated** | Pointer to **string** | Updated is when it last changed, RFC 3339. | [optional] 

## Methods

### NewFormTheme

`func NewFormTheme() *FormTheme`

NewFormTheme instantiates a new FormTheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormThemeWithDefaults

`func NewFormThemeWithDefaults() *FormTheme`

NewFormThemeWithDefaults instantiates a new FormTheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackground

`func (o *FormTheme) GetBackground() string`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *FormTheme) GetBackgroundOk() (*string, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *FormTheme) SetBackground(v string)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *FormTheme) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### GetColors

`func (o *FormTheme) GetColors() FormColors`

GetColors returns the Colors field if non-nil, zero value otherwise.

### GetColorsOk

`func (o *FormTheme) GetColorsOk() (*FormColors, bool)`

GetColorsOk returns a tuple with the Colors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColors

`func (o *FormTheme) SetColors(v FormColors)`

SetColors sets Colors field to given value.

### HasColors

`func (o *FormTheme) HasColors() bool`

HasColors returns a boolean if a field has been set.

### GetCreated

`func (o *FormTheme) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FormTheme) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FormTheme) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FormTheme) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFont

`func (o *FormTheme) GetFont() string`

GetFont returns the Font field if non-nil, zero value otherwise.

### GetFontOk

`func (o *FormTheme) GetFontOk() (*string, bool)`

GetFontOk returns a tuple with the Font field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFont

`func (o *FormTheme) SetFont(v string)`

SetFont sets Font field to given value.

### HasFont

`func (o *FormTheme) HasFont() bool`

HasFont returns a boolean if a field has been set.

### GetId

`func (o *FormTheme) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormTheme) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormTheme) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormTheme) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FormTheme) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormTheme) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormTheme) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormTheme) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdated

`func (o *FormTheme) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FormTheme) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FormTheme) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FormTheme) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


