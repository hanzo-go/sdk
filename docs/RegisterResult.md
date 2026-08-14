# RegisterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quote** | Pointer to [**Offer**](Offer.md) | the price it was bought at | [optional] 
**Record** | Pointer to [**Holding**](Holding.md) | the ownership row this purchase issued | [optional] 

## Methods

### NewRegisterResult

`func NewRegisterResult() *RegisterResult`

NewRegisterResult instantiates a new RegisterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterResultWithDefaults

`func NewRegisterResultWithDefaults() *RegisterResult`

NewRegisterResultWithDefaults instantiates a new RegisterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuote

`func (o *RegisterResult) GetQuote() Offer`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *RegisterResult) GetQuoteOk() (*Offer, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *RegisterResult) SetQuote(v Offer)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *RegisterResult) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetRecord

`func (o *RegisterResult) GetRecord() Holding`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *RegisterResult) GetRecordOk() (*Holding, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *RegisterResult) SetRecord(v Holding)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *RegisterResult) HasRecord() bool`

HasRecord returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


