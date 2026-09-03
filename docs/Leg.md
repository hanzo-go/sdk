# Leg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the chart-of-accounts number this side posts to, e.g. \&quot;5300\&quot;. | [optional] 
**Credit** | Pointer to **int64** | Credit is the leg&#39;s credit in exact cents. Set this or Debit, not both. | [optional] 
**Debit** | Pointer to **int64** | Debit is the leg&#39;s debit in exact cents. Set this or Credit, not both. | [optional] 

## Methods

### NewLeg

`func NewLeg() *Leg`

NewLeg instantiates a new Leg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegWithDefaults

`func NewLegWithDefaults() *Leg`

NewLegWithDefaults instantiates a new Leg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Leg) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Leg) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Leg) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Leg) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCredit

`func (o *Leg) GetCredit() int64`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *Leg) GetCreditOk() (*int64, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *Leg) SetCredit(v int64)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *Leg) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetDebit

`func (o *Leg) GetDebit() int64`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *Leg) GetDebitOk() (*int64, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *Leg) SetDebit(v int64)`

SetDebit sets Debit field to given value.

### HasDebit

`func (o *Leg) HasDebit() bool`

HasDebit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


