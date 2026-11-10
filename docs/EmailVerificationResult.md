# EmailVerificationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checks** | Pointer to [**EmailVerificationChecks**](EmailVerificationChecks.md) | Checks contains the outcomes of all requested checks. | [optional] 
**Email** | Pointer to **string** | Email is the original submitted email address. | [optional] 
**Suggestions** | Pointer to [**[]EmailSuggestion**](EmailSuggestion.md) | Suggestions lists potential corrections when requested. | [optional] 

## Methods

### NewEmailVerificationResult

`func NewEmailVerificationResult() *EmailVerificationResult`

NewEmailVerificationResult instantiates a new EmailVerificationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailVerificationResultWithDefaults

`func NewEmailVerificationResultWithDefaults() *EmailVerificationResult`

NewEmailVerificationResultWithDefaults instantiates a new EmailVerificationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecks

`func (o *EmailVerificationResult) GetChecks() EmailVerificationChecks`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *EmailVerificationResult) GetChecksOk() (*EmailVerificationChecks, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *EmailVerificationResult) SetChecks(v EmailVerificationChecks)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *EmailVerificationResult) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetEmail

`func (o *EmailVerificationResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailVerificationResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailVerificationResult) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmailVerificationResult) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSuggestions

`func (o *EmailVerificationResult) GetSuggestions() []EmailSuggestion`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *EmailVerificationResult) GetSuggestionsOk() (*[]EmailSuggestion, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *EmailVerificationResult) SetSuggestions(v []EmailSuggestion)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *EmailVerificationResult) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


