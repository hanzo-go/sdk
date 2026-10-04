# DomainRegisterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quote** | Pointer to [**DomainOffer**](DomainOffer.md) | the price it was bought at | [optional] 
**Record** | Pointer to [**DomainHolding**](DomainHolding.md) | the ownership row this purchase issued | [optional] 

## Methods

### NewDomainRegisterResult

`func NewDomainRegisterResult() *DomainRegisterResult`

NewDomainRegisterResult instantiates a new DomainRegisterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRegisterResultWithDefaults

`func NewDomainRegisterResultWithDefaults() *DomainRegisterResult`

NewDomainRegisterResultWithDefaults instantiates a new DomainRegisterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuote

`func (o *DomainRegisterResult) GetQuote() DomainOffer`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *DomainRegisterResult) GetQuoteOk() (*DomainOffer, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *DomainRegisterResult) SetQuote(v DomainOffer)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *DomainRegisterResult) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetRecord

`func (o *DomainRegisterResult) GetRecord() DomainHolding`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *DomainRegisterResult) GetRecordOk() (*DomainHolding, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *DomainRegisterResult) SetRecord(v DomainHolding)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *DomainRegisterResult) HasRecord() bool`

HasRecord returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


