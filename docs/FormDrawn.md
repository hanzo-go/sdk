# FormDrawn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to **string** | Language is the form&#39;s language tag. | [optional] 
**Mode** | Pointer to **string** | Mode is one (a question at a time) or all. | [optional] 
**Progress** | Pointer to **bool** | Progress is whether a progress bar is shown. | [optional] 

## Methods

### NewFormDrawn

`func NewFormDrawn() *FormDrawn`

NewFormDrawn instantiates a new FormDrawn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormDrawnWithDefaults

`func NewFormDrawnWithDefaults() *FormDrawn`

NewFormDrawnWithDefaults instantiates a new FormDrawn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *FormDrawn) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *FormDrawn) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *FormDrawn) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *FormDrawn) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMode

`func (o *FormDrawn) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *FormDrawn) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *FormDrawn) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *FormDrawn) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetProgress

`func (o *FormDrawn) GetProgress() bool`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *FormDrawn) GetProgressOk() (*bool, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *FormDrawn) SetProgress(v bool)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *FormDrawn) HasProgress() bool`

HasProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


