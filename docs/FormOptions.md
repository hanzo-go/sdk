# FormOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to **string** | Identity is who a response is from when nobody is signed in. claimed (the default): an address typed into an email field is the respondent&#39;s CLAIM, filed on the CRM marked unverified, on the person it names only when the CRM already holds them and without changing them; it joins no list and is mailed no welcome, and a list&#39;s consent is asked by the double opt-in instead. verified: only a signed-in respondent (their token) is anyone; a typed address is an answer and nothing more. A signed-in respondent is verified under either. | [optional] 
**Language** | Pointer to **string** | Language is the form&#39;s language, a tag like en or pt-BR. Default en. | [optional] 
**Mode** | Pointer to **string** | Mode is one (a question at a time, the default) or all (every question at once). | [optional] 
**Progress** | Pointer to **bool** | Progress is whether the renderer shows a progress bar. | [optional] 
**Visibility** | Pointer to **string** | Visibility is public (answerable at its public address by anyone) or private (not served there). Default public. | [optional] 

## Methods

### NewFormOptions

`func NewFormOptions() *FormOptions`

NewFormOptions instantiates a new FormOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormOptionsWithDefaults

`func NewFormOptionsWithDefaults() *FormOptions`

NewFormOptionsWithDefaults instantiates a new FormOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *FormOptions) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *FormOptions) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *FormOptions) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *FormOptions) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetLanguage

`func (o *FormOptions) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *FormOptions) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *FormOptions) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *FormOptions) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMode

`func (o *FormOptions) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FormOptions) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FormOptions) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FormOptions) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetProgress

`func (o *FormOptions) GetProgress() bool`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *FormOptions) GetProgressOk() (*bool, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *FormOptions) SetProgress(v bool)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *FormOptions) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetVisibility

`func (o *FormOptions) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FormOptions) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FormOptions) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FormOptions) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


