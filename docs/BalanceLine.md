# BalanceLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the chart-of-accounts number this line reports on. ABSENT marks a DERIVED line that no account holds — retained earnings is the one such line, computed from cumulative income minus expense. | [optional] 
**Amount** | Pointer to **int32** | Amount is the balance as of the statement date, in whole cents, in its NATURAL sign: positive when the account behaved normally, on all three sides. Assets are debit-normal and shown as stored; liabilities and equity are credit-normal and flipped once here for display. A negative asset is a real overdraft, not a sign convention. | [optional] 
**Name** | Pointer to **string** | Name is the account&#39;s human name, or the derived line&#39;s own name. | [optional] 
**Type** | Pointer to **string** | Type is the account&#39;s fundamental class. Absent on a derived line, which belongs to no account and therefore has none. | [optional] 

## Methods

### NewBalanceLine

`func NewBalanceLine() *BalanceLine`

NewBalanceLine instantiates a new BalanceLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceLineWithDefaults

`func NewBalanceLineWithDefaults() *BalanceLine`

NewBalanceLineWithDefaults instantiates a new BalanceLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *BalanceLine) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *BalanceLine) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *BalanceLine) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *BalanceLine) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAmount

`func (o *BalanceLine) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BalanceLine) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BalanceLine) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BalanceLine) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetName

`func (o *BalanceLine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BalanceLine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BalanceLine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BalanceLine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *BalanceLine) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BalanceLine) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BalanceLine) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BalanceLine) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


