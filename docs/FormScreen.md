# FormScreen

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Button** | Pointer to **string** | Button is the button&#39;s label. At most 64 characters. | [optional] 
**Description** | Pointer to **string** | Description is the text under it. At most 2000 characters. | [optional] 
**Redirect** | Pointer to **string** | Redirect is where a thank-you screen sends the respondent instead of showing itself: an https address. Ignored on a welcome screen. | [optional] 
**Title** | Pointer to **string** | Title is the screen&#39;s heading. At most 200 characters. | [optional] 

## Methods

### NewFormScreen

`func NewFormScreen() *FormScreen`

NewFormScreen instantiates a new FormScreen object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormScreenWithDefaults

`func NewFormScreenWithDefaults() *FormScreen`

NewFormScreenWithDefaults instantiates a new FormScreen object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetButton

`func (o *FormScreen) GetButton() string`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *FormScreen) GetButtonOk() (*string, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *FormScreen) SetButton(v string)`

SetButton sets Button field to given value.

### HasButton

`func (o *FormScreen) HasButton() bool`

HasButton returns a boolean if a field has been set.

### GetDescription

`func (o *FormScreen) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FormScreen) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FormScreen) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FormScreen) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRedirect

`func (o *FormScreen) GetRedirect() string`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *FormScreen) GetRedirectOk() (*string, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *FormScreen) SetRedirect(v string)`

SetRedirect sets Redirect field to given value.

### HasRedirect

`func (o *FormScreen) HasRedirect() bool`

HasRedirect returns a boolean if a field has been set.

### GetTitle

`func (o *FormScreen) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FormScreen) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FormScreen) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FormScreen) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


