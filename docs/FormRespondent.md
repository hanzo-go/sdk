# FormRespondent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is the address the response is from: the account&#39;s, else the email field&#39;s. | [optional] 
**Kind** | Pointer to **string** | Kind is account (a signed-in person, from their token), email (an email field&#39;s answer), visitor (the page&#39;s visitor id) or none. | [optional] 
**PersonId** | Pointer to **string** | PersonID is the CRM person the address is: the one key a response, a list membership, an email event and an experiment share. Every event and timeline entry the service states carries it, never the address. | [optional] 
**Subject** | Pointer to **string** | Subject is the signed-in person&#39;s IAM subject. | [optional] 
**Verified** | Pointer to **bool** | Verified is whether the respondent is who the response says: true only for a signed-in person, from their token. A typed address is a claim. | [optional] 
**Visitor** | Pointer to **string** | Visitor is the page&#39;s visitor id. | [optional] 

## Methods

### NewFormRespondent

`func NewFormRespondent() *FormRespondent`

NewFormRespondent instantiates a new FormRespondent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormRespondentWithDefaults

`func NewFormRespondentWithDefaults() *FormRespondent`

NewFormRespondentWithDefaults instantiates a new FormRespondent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *FormRespondent) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FormRespondent) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FormRespondent) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FormRespondent) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetKind

`func (o *FormRespondent) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FormRespondent) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FormRespondent) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FormRespondent) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPersonId

`func (o *FormRespondent) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *FormRespondent) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *FormRespondent) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *FormRespondent) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetSubject

`func (o *FormRespondent) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *FormRespondent) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *FormRespondent) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *FormRespondent) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetVerified

`func (o *FormRespondent) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *FormRespondent) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *FormRespondent) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *FormRespondent) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetVisitor

`func (o *FormRespondent) GetVisitor() string`

GetVisitor returns the Visitor field if non-nil, zero value otherwise.

### GetVisitorOk

`func (o *FormRespondent) GetVisitorOk() (*string, bool)`

GetVisitorOk returns a tuple with the Visitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitor

`func (o *FormRespondent) SetVisitor(v string)`

SetVisitor sets Visitor field to given value.

### HasVisitor

`func (o *FormRespondent) HasVisitor() bool`

HasVisitor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


