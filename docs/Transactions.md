# Transactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Transactions** | Pointer to [**[]Transaction**](Transaction.md) |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewTransactions

`func NewTransactions() *Transactions`

NewTransactions instantiates a new Transactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsWithDefaults

`func NewTransactionsWithDefaults() *Transactions`

NewTransactionsWithDefaults instantiates a new Transactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *Transactions) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Transactions) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Transactions) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Transactions) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTransactions

`func (o *Transactions) GetTransactions() []Transaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *Transactions) GetTransactionsOk() (*[]Transaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *Transactions) SetTransactions(v []Transaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *Transactions) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetUser

`func (o *Transactions) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Transactions) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Transactions) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Transactions) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


