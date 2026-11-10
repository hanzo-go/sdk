# MarketingMailIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body is the mail as plain text, at most 20000 characters; a blank line separates paragraphs. | [optional] 
**Person** | Pointer to **string** | Person is the CRM person the mail is written to. Empty, the mail goes on the timeline of the person the org holds for the address, if any. | [optional] 
**Subject** | Pointer to **string** | Subject is the subject line, at most 200 characters on one line. | [optional] 
**To** | Pointer to **string** | To is the one address the mail goes to. | [optional] 

## Methods

### NewMarketingMailIn

`func NewMarketingMailIn() *MarketingMailIn`

NewMarketingMailIn instantiates a new MarketingMailIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingMailInWithDefaults

`func NewMarketingMailInWithDefaults() *MarketingMailIn`

NewMarketingMailInWithDefaults instantiates a new MarketingMailIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *MarketingMailIn) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MarketingMailIn) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *MarketingMailIn) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *MarketingMailIn) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetPerson

`func (o *MarketingMailIn) GetPerson() string`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *MarketingMailIn) GetPersonOk() (*string, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *MarketingMailIn) SetPerson(v string)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *MarketingMailIn) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetSubject

`func (o *MarketingMailIn) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MarketingMailIn) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MarketingMailIn) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MarketingMailIn) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTo

`func (o *MarketingMailIn) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MarketingMailIn) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MarketingMailIn) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *MarketingMailIn) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


