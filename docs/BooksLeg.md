# BooksLeg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the chart-of-accounts number this side posts to, e.g. \&quot;5300\&quot;. | [optional] 
**Credit** | Pointer to **int64** | Credit is the leg&#39;s credit in exact cents. Set this or Debit, not both. | [optional] 
**Debit** | Pointer to **int64** | Debit is the leg&#39;s debit in exact cents. Set this or Credit, not both. | [optional] 

## Methods

### NewBooksLeg

`func NewBooksLeg() *BooksLeg`

NewBooksLeg instantiates a new BooksLeg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksLegWithDefaults

`func NewBooksLegWithDefaults() *BooksLeg`

NewBooksLegWithDefaults instantiates a new BooksLeg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *BooksLeg) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *BooksLeg) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *BooksLeg) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *BooksLeg) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCredit

`func (o *BooksLeg) GetCredit() int64`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *BooksLeg) GetCreditOk() (*int64, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *BooksLeg) SetCredit(v int64)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *BooksLeg) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetDebit

`func (o *BooksLeg) GetDebit() int64`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *BooksLeg) GetDebitOk() (*int64, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *BooksLeg) SetDebit(v int64)`

SetDebit sets Debit field to given value.

### HasDebit

`func (o *BooksLeg) HasDebit() bool`

HasDebit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


