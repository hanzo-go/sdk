# EmailSyntaxCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | Domain is the domain-part of the email address. | [optional] 
**Reason** | Pointer to **string** | Reason describes the failure or skip reason when Status is not pass. | [optional] 
**Status** | Pointer to **string** | Status reports the check outcome: pass, fail, unknown, or skipped. | [optional] 
**Username** | Pointer to **string** | Username is the local-part of the email address. | [optional] 

## Methods

### NewEmailSyntaxCheck

`func NewEmailSyntaxCheck() *EmailSyntaxCheck`

NewEmailSyntaxCheck instantiates a new EmailSyntaxCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSyntaxCheckWithDefaults

`func NewEmailSyntaxCheckWithDefaults() *EmailSyntaxCheck`

NewEmailSyntaxCheckWithDefaults instantiates a new EmailSyntaxCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *EmailSyntaxCheck) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EmailSyntaxCheck) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EmailSyntaxCheck) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *EmailSyntaxCheck) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetReason

`func (o *EmailSyntaxCheck) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EmailSyntaxCheck) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EmailSyntaxCheck) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EmailSyntaxCheck) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *EmailSyntaxCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailSyntaxCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailSyntaxCheck) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailSyntaxCheck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsername

`func (o *EmailSyntaxCheck) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EmailSyntaxCheck) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EmailSyntaxCheck) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EmailSyntaxCheck) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


