# O11yO11yErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the machine-readable code. | [optional] 
**Errors** | Pointer to [**[]O11yO11yErrorItem**](O11yO11yErrorItem.md) | Errors are further details, one message and its suggestions each. | [optional] 
**Message** | Pointer to **string** | Message is the human-readable reason. | [optional] 
**Retry** | Pointer to [**O11yO11yRetry**](O11yO11yRetry.md) | Retry says when it is worth trying again, for errors that pass. | [optional] 
**Suggestions** | Pointer to **[]string** | Suggestions say what to try instead. | [optional] 
**Type** | Pointer to **string** | Type is the error&#39;s category, e.g. invalid_input, not_found. | [optional] 
**Url** | Pointer to **string** | Url points at documentation for the error, when there is any. | [optional] 

## Methods

### NewO11yO11yErrorDetail

`func NewO11yO11yErrorDetail() *O11yO11yErrorDetail`

NewO11yO11yErrorDetail instantiates a new O11yO11yErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yErrorDetailWithDefaults

`func NewO11yO11yErrorDetailWithDefaults() *O11yO11yErrorDetail`

NewO11yO11yErrorDetailWithDefaults instantiates a new O11yO11yErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *O11yO11yErrorDetail) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *O11yO11yErrorDetail) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *O11yO11yErrorDetail) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *O11yO11yErrorDetail) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetErrors

`func (o *O11yO11yErrorDetail) GetErrors() []O11yO11yErrorItem`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *O11yO11yErrorDetail) GetErrorsOk() (*[]O11yO11yErrorItem, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *O11yO11yErrorDetail) SetErrors(v []O11yO11yErrorItem)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *O11yO11yErrorDetail) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetMessage

`func (o *O11yO11yErrorDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *O11yO11yErrorDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *O11yO11yErrorDetail) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *O11yO11yErrorDetail) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRetry

`func (o *O11yO11yErrorDetail) GetRetry() O11yO11yRetry`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *O11yO11yErrorDetail) GetRetryOk() (*O11yO11yRetry, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *O11yO11yErrorDetail) SetRetry(v O11yO11yRetry)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *O11yO11yErrorDetail) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetSuggestions

`func (o *O11yO11yErrorDetail) GetSuggestions() []string`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *O11yO11yErrorDetail) GetSuggestionsOk() (*[]string, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *O11yO11yErrorDetail) SetSuggestions(v []string)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *O11yO11yErrorDetail) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yErrorDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yErrorDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yErrorDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yErrorDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *O11yO11yErrorDetail) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *O11yO11yErrorDetail) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *O11yO11yErrorDetail) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *O11yO11yErrorDetail) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


