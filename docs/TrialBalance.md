# TrialBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balanced** | Pointer to **bool** | Balanced is the proof this report exists to give: whether total debits equal total credits. It is computed from the rows above, never assumed, and false means the ledger itself is broken rather than that the report is wrong. | [optional] 
**From** | Pointer to **string** | From is the posting time the window opens at, as it was asked for. Absent means the report runs from the beginning of the ledger. | [optional] 
**Rows** | Pointer to [**[]TrialBalanceRow**](TrialBalanceRow.md) | Rows are the accounts that MOVED in one of the windows. An account that never moved is omitted rather than listed at zero, so this is shorter than the chart. | [optional] 
**To** | Pointer to **string** | To is the posting time the window closes at, inclusive. Absent means \&quot;up to now\&quot; — every posting the ledger holds. | [optional] 
**TotalCredit** | Pointer to **int32** | TotalCredit is the sum of every row&#39;s closing credit column, in cents. | [optional] 
**TotalDebit** | Pointer to **int32** | TotalDebit is the sum of every row&#39;s CLOSING debit column, in cents. | [optional] 

## Methods

### NewTrialBalance

`func NewTrialBalance() *TrialBalance`

NewTrialBalance instantiates a new TrialBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialBalanceWithDefaults

`func NewTrialBalanceWithDefaults() *TrialBalance`

NewTrialBalanceWithDefaults instantiates a new TrialBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanced

`func (o *TrialBalance) GetBalanced() bool`

GetBalanced returns the Balanced field if non-nil, zero value otherwise.

### GetBalancedOk

`func (o *TrialBalance) GetBalancedOk() (*bool, bool)`

GetBalancedOk returns a tuple with the Balanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanced

`func (o *TrialBalance) SetBalanced(v bool)`

SetBalanced sets Balanced field to given value.

### HasBalanced

`func (o *TrialBalance) HasBalanced() bool`

HasBalanced returns a boolean if a field has been set.

### GetFrom

`func (o *TrialBalance) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TrialBalance) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TrialBalance) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *TrialBalance) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetRows

`func (o *TrialBalance) GetRows() []TrialBalanceRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *TrialBalance) GetRowsOk() (*[]TrialBalanceRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *TrialBalance) SetRows(v []TrialBalanceRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *TrialBalance) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTo

`func (o *TrialBalance) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TrialBalance) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TrialBalance) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *TrialBalance) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTotalCredit

`func (o *TrialBalance) GetTotalCredit() int32`

GetTotalCredit returns the TotalCredit field if non-nil, zero value otherwise.

### GetTotalCreditOk

`func (o *TrialBalance) GetTotalCreditOk() (*int32, bool)`

GetTotalCreditOk returns a tuple with the TotalCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCredit

`func (o *TrialBalance) SetTotalCredit(v int32)`

SetTotalCredit sets TotalCredit field to given value.

### HasTotalCredit

`func (o *TrialBalance) HasTotalCredit() bool`

HasTotalCredit returns a boolean if a field has been set.

### GetTotalDebit

`func (o *TrialBalance) GetTotalDebit() int32`

GetTotalDebit returns the TotalDebit field if non-nil, zero value otherwise.

### GetTotalDebitOk

`func (o *TrialBalance) GetTotalDebitOk() (*int32, bool)`

GetTotalDebitOk returns a tuple with the TotalDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDebit

`func (o *TrialBalance) SetTotalDebit(v int32)`

SetTotalDebit sets TotalDebit field to given value.

### HasTotalDebit

`func (o *TrialBalance) HasTotalDebit() bool`

HasTotalDebit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


