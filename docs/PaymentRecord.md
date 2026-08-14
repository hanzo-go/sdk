# PaymentRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** | AmountCents is the credited amount in whole cents. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when the credit was written, RFC3339. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO 4217 code. | [optional] 
**Id** | Pointer to **string** | ID is the ledger transaction id. | [optional] 
**Notes** | Pointer to **string** | Notes is the ledger memo, carrying the processor and its reference. | [optional] 
**Status** | Pointer to **string** | Status is the payment&#39;s state. This ledger writes a deposit only AFTER the processor settled, so a payment that can be read is one that succeeded. | [optional] 
**Subject** | Pointer to **string** | Subject is the billing key this payment credited. | [optional] 
**Test** | Pointer to **bool** | Test reports whether this was a sandbox charge (test balance) or live money. | [optional] 

## Methods

### NewPaymentRecord

`func NewPaymentRecord() *PaymentRecord`

NewPaymentRecord instantiates a new PaymentRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRecordWithDefaults

`func NewPaymentRecordWithDefaults() *PaymentRecord`

NewPaymentRecordWithDefaults instantiates a new PaymentRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *PaymentRecord) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *PaymentRecord) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *PaymentRecord) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *PaymentRecord) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaymentRecord) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentRecord) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentRecord) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaymentRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentRecord) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentRecord) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentRecord) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentRecord) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetId

`func (o *PaymentRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotes

`func (o *PaymentRecord) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PaymentRecord) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PaymentRecord) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PaymentRecord) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *PaymentRecord) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PaymentRecord) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PaymentRecord) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PaymentRecord) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTest

`func (o *PaymentRecord) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *PaymentRecord) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *PaymentRecord) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *PaymentRecord) HasTest() bool`

HasTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


