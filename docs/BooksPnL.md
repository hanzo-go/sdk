# BooksPnL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expense** | Pointer to [**[]BooksPnLLine**](BooksPnLLine.md) | Expense is the cost lines that moved in the period, one per account. | [optional] 
**From** | Pointer to **string** | From opens the period and is EXCLUSIVE — movement strictly after it, matching the trial balance&#39;s opening boundary so the two reports agree on what belongs to a period. Absent means from the beginning of the ledger. | [optional] 
**Income** | Pointer to [**[]BooksPnLLine**](BooksPnLLine.md) | Income is the revenue lines that moved in the period, one per account. Accounts that did not move are omitted rather than listed at zero. | [optional] 
**NetIncome** | Pointer to **int64** | NetIncome is totalIncome minus totalExpense, in cents. Negative is a loss. | [optional] 
**To** | Pointer to **string** | To closes the period and is inclusive. Absent means up to now. | [optional] 
**TotalExpense** | Pointer to **int64** | TotalExpense is cost MATCHED to that revenue, in cents, including accrued infrastructure that has not been billed yet. | [optional] 
**TotalIncome** | Pointer to **int64** | TotalIncome is revenue RECOGNIZED in the period, in cents — accrual, not cash, so a prepaid top-up is not in it until the credit is consumed. | [optional] 

## Methods

### NewBooksPnL

`func NewBooksPnL() *BooksPnL`

NewBooksPnL instantiates a new BooksPnL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksPnLWithDefaults

`func NewBooksPnLWithDefaults() *BooksPnL`

NewBooksPnLWithDefaults instantiates a new BooksPnL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpense

`func (o *BooksPnL) GetExpense() []BooksPnLLine`

GetExpense returns the Expense field if non-nil, zero value otherwise.

### GetExpenseOk

`func (o *BooksPnL) GetExpenseOk() (*[]BooksPnLLine, bool)`

GetExpenseOk returns a tuple with the Expense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpense

`func (o *BooksPnL) SetExpense(v []BooksPnLLine)`

SetExpense sets Expense field to given value.

### HasExpense

`func (o *BooksPnL) HasExpense() bool`

HasExpense returns a boolean if a field has been set.

### GetFrom

`func (o *BooksPnL) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BooksPnL) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BooksPnL) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *BooksPnL) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetIncome

`func (o *BooksPnL) GetIncome() []BooksPnLLine`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *BooksPnL) GetIncomeOk() (*[]BooksPnLLine, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *BooksPnL) SetIncome(v []BooksPnLLine)`

SetIncome sets Income field to given value.

### HasIncome

`func (o *BooksPnL) HasIncome() bool`

HasIncome returns a boolean if a field has been set.

### GetNetIncome

`func (o *BooksPnL) GetNetIncome() int64`

GetNetIncome returns the NetIncome field if non-nil, zero value otherwise.

### GetNetIncomeOk

`func (o *BooksPnL) GetNetIncomeOk() (*int64, bool)`

GetNetIncomeOk returns a tuple with the NetIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIncome

`func (o *BooksPnL) SetNetIncome(v int64)`

SetNetIncome sets NetIncome field to given value.

### HasNetIncome

`func (o *BooksPnL) HasNetIncome() bool`

HasNetIncome returns a boolean if a field has been set.

### GetTo

`func (o *BooksPnL) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BooksPnL) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BooksPnL) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *BooksPnL) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTotalExpense

`func (o *BooksPnL) GetTotalExpense() int64`

GetTotalExpense returns the TotalExpense field if non-nil, zero value otherwise.

### GetTotalExpenseOk

`func (o *BooksPnL) GetTotalExpenseOk() (*int64, bool)`

GetTotalExpenseOk returns a tuple with the TotalExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExpense

`func (o *BooksPnL) SetTotalExpense(v int64)`

SetTotalExpense sets TotalExpense field to given value.

### HasTotalExpense

`func (o *BooksPnL) HasTotalExpense() bool`

HasTotalExpense returns a boolean if a field has been set.

### GetTotalIncome

`func (o *BooksPnL) GetTotalIncome() int64`

GetTotalIncome returns the TotalIncome field if non-nil, zero value otherwise.

### GetTotalIncomeOk

`func (o *BooksPnL) GetTotalIncomeOk() (*int64, bool)`

GetTotalIncomeOk returns a tuple with the TotalIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIncome

`func (o *BooksPnL) SetTotalIncome(v int64)`

SetTotalIncome sets TotalIncome field to given value.

### HasTotalIncome

`func (o *BooksPnL) HasTotalIncome() bool`

HasTotalIncome returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


