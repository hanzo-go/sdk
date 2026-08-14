# PnL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expense** | Pointer to [**[]PnLLine**](PnLLine.md) |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Income** | Pointer to [**[]PnLLine**](PnLLine.md) |  | [optional] 
**NetIncome** | Pointer to **int32** | TotalIncome − TotalExpense | [optional] 
**To** | Pointer to **string** |  | [optional] 
**TotalExpense** | Pointer to **int32** |  | [optional] 
**TotalIncome** | Pointer to **int32** |  | [optional] 

## Methods

### NewPnL

`func NewPnL() *PnL`

NewPnL instantiates a new PnL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPnLWithDefaults

`func NewPnLWithDefaults() *PnL`

NewPnLWithDefaults instantiates a new PnL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpense

`func (o *PnL) GetExpense() []PnLLine`

GetExpense returns the Expense field if non-nil, zero value otherwise.

### GetExpenseOk

`func (o *PnL) GetExpenseOk() (*[]PnLLine, bool)`

GetExpenseOk returns a tuple with the Expense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpense

`func (o *PnL) SetExpense(v []PnLLine)`

SetExpense sets Expense field to given value.

### HasExpense

`func (o *PnL) HasExpense() bool`

HasExpense returns a boolean if a field has been set.

### GetFrom

`func (o *PnL) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PnL) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PnL) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PnL) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetIncome

`func (o *PnL) GetIncome() []PnLLine`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *PnL) GetIncomeOk() (*[]PnLLine, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *PnL) SetIncome(v []PnLLine)`

SetIncome sets Income field to given value.

### HasIncome

`func (o *PnL) HasIncome() bool`

HasIncome returns a boolean if a field has been set.

### GetNetIncome

`func (o *PnL) GetNetIncome() int32`

GetNetIncome returns the NetIncome field if non-nil, zero value otherwise.

### GetNetIncomeOk

`func (o *PnL) GetNetIncomeOk() (*int32, bool)`

GetNetIncomeOk returns a tuple with the NetIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIncome

`func (o *PnL) SetNetIncome(v int32)`

SetNetIncome sets NetIncome field to given value.

### HasNetIncome

`func (o *PnL) HasNetIncome() bool`

HasNetIncome returns a boolean if a field has been set.

### GetTo

`func (o *PnL) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PnL) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PnL) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *PnL) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTotalExpense

`func (o *PnL) GetTotalExpense() int32`

GetTotalExpense returns the TotalExpense field if non-nil, zero value otherwise.

### GetTotalExpenseOk

`func (o *PnL) GetTotalExpenseOk() (*int32, bool)`

GetTotalExpenseOk returns a tuple with the TotalExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExpense

`func (o *PnL) SetTotalExpense(v int32)`

SetTotalExpense sets TotalExpense field to given value.

### HasTotalExpense

`func (o *PnL) HasTotalExpense() bool`

HasTotalExpense returns a boolean if a field has been set.

### GetTotalIncome

`func (o *PnL) GetTotalIncome() int32`

GetTotalIncome returns the TotalIncome field if non-nil, zero value otherwise.

### GetTotalIncomeOk

`func (o *PnL) GetTotalIncomeOk() (*int32, bool)`

GetTotalIncomeOk returns a tuple with the TotalIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIncome

`func (o *PnL) SetTotalIncome(v int32)`

SetTotalIncome sets TotalIncome field to given value.

### HasTotalIncome

`func (o *PnL) HasTotalIncome() bool`

HasTotalIncome returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


