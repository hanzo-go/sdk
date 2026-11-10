# EmailSuggestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is the suggested alternative address. | [optional] 
**Reason** | Pointer to **string** | Reason describes why the suggestion was made (e.g. domain_typo). | [optional] 

## Methods

### NewEmailSuggestion

`func NewEmailSuggestion() *EmailSuggestion`

NewEmailSuggestion instantiates a new EmailSuggestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSuggestionWithDefaults

`func NewEmailSuggestionWithDefaults() *EmailSuggestion`

NewEmailSuggestionWithDefaults instantiates a new EmailSuggestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EmailSuggestion) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailSuggestion) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailSuggestion) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmailSuggestion) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetReason

`func (o *EmailSuggestion) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EmailSuggestion) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EmailSuggestion) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EmailSuggestion) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


