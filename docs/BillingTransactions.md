# BillingTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Transactions** | Pointer to [**[]BillingTransaction**](BillingTransaction.md) |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingTransactions

`func NewBillingTransactions() *BillingTransactions`

NewBillingTransactions instantiates a new BillingTransactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingTransactionsWithDefaults

`func NewBillingTransactionsWithDefaults() *BillingTransactions`

NewBillingTransactionsWithDefaults instantiates a new BillingTransactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *BillingTransactions) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BillingTransactions) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BillingTransactions) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *BillingTransactions) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTransactions

`func (o *BillingTransactions) GetTransactions() []BillingTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *BillingTransactions) GetTransactionsOk() (*[]BillingTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *BillingTransactions) SetTransactions(v []BillingTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *BillingTransactions) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetUser

`func (o *BillingTransactions) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BillingTransactions) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BillingTransactions) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *BillingTransactions) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


