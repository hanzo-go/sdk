# EmailVerificationChecks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classification** | Pointer to [**EmailClassificationCheck**](EmailClassificationCheck.md) | Classification details account attributes (disposable, free, role). | [optional] 
**Domain** | Pointer to [**EmailDomainCheck**](EmailDomainCheck.md) | Domain details DNS domain host resolution. | [optional] 
**Mx** | Pointer to [**EmailMXCheck**](EmailMXCheck.md) | MX details DNS mail exchanger lookup. | [optional] 
**Smtp** | Pointer to [**EmailSMTPCheck**](EmailSMTPCheck.md) | SMTP details explicit mailbox probing. | [optional] 
**Syntax** | Pointer to [**EmailSyntaxCheck**](EmailSyntaxCheck.md) | Syntax details RFC 5322 structure verification. | [optional] 

## Methods

### NewEmailVerificationChecks

`func NewEmailVerificationChecks() *EmailVerificationChecks`

NewEmailVerificationChecks instantiates a new EmailVerificationChecks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailVerificationChecksWithDefaults

`func NewEmailVerificationChecksWithDefaults() *EmailVerificationChecks`

NewEmailVerificationChecksWithDefaults instantiates a new EmailVerificationChecks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassification

`func (o *EmailVerificationChecks) GetClassification() EmailClassificationCheck`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *EmailVerificationChecks) GetClassificationOk() (*EmailClassificationCheck, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *EmailVerificationChecks) SetClassification(v EmailClassificationCheck)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *EmailVerificationChecks) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetDomain

`func (o *EmailVerificationChecks) GetDomain() EmailDomainCheck`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EmailVerificationChecks) GetDomainOk() (*EmailDomainCheck, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EmailVerificationChecks) SetDomain(v EmailDomainCheck)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *EmailVerificationChecks) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetMx

`func (o *EmailVerificationChecks) GetMx() EmailMXCheck`

GetMx returns the Mx field if non-nil, zero value otherwise.

### GetMxOk

`func (o *EmailVerificationChecks) GetMxOk() (*EmailMXCheck, bool)`

GetMxOk returns a tuple with the Mx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMx

`func (o *EmailVerificationChecks) SetMx(v EmailMXCheck)`

SetMx sets Mx field to given value.

### HasMx

`func (o *EmailVerificationChecks) HasMx() bool`

HasMx returns a boolean if a field has been set.

### GetSmtp

`func (o *EmailVerificationChecks) GetSmtp() EmailSMTPCheck`

GetSmtp returns the Smtp field if non-nil, zero value otherwise.

### GetSmtpOk

`func (o *EmailVerificationChecks) GetSmtpOk() (*EmailSMTPCheck, bool)`

GetSmtpOk returns a tuple with the Smtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtp

`func (o *EmailVerificationChecks) SetSmtp(v EmailSMTPCheck)`

SetSmtp sets Smtp field to given value.

### HasSmtp

`func (o *EmailVerificationChecks) HasSmtp() bool`

HasSmtp returns a boolean if a field has been set.

### GetSyntax

`func (o *EmailVerificationChecks) GetSyntax() EmailSyntaxCheck`

GetSyntax returns the Syntax field if non-nil, zero value otherwise.

### GetSyntaxOk

`func (o *EmailVerificationChecks) GetSyntaxOk() (*EmailSyntaxCheck, bool)`

GetSyntaxOk returns a tuple with the Syntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntax

`func (o *EmailVerificationChecks) SetSyntax(v EmailSyntaxCheck)`

SetSyntax sets Syntax field to given value.

### HasSyntax

`func (o *EmailVerificationChecks) HasSyntax() bool`

HasSyntax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


