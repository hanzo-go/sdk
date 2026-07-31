# CloudTrialBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balanced** | Pointer to **bool** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Rows** | Pointer to [**[]CloudTrialBalanceRow**](CloudTrialBalanceRow.md) |  | [optional] 
**To** | Pointer to **string** |  | [optional] 
**TotalCredit** | Pointer to **int32** |  | [optional] 
**TotalDebit** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudTrialBalance

`func NewCloudTrialBalance() *CloudTrialBalance`

NewCloudTrialBalance instantiates a new CloudTrialBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTrialBalanceWithDefaults

`func NewCloudTrialBalanceWithDefaults() *CloudTrialBalance`

NewCloudTrialBalanceWithDefaults instantiates a new CloudTrialBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanced

`func (o *CloudTrialBalance) GetBalanced() bool`

GetBalanced returns the Balanced field if non-nil, zero value otherwise.

### GetBalancedOk

`func (o *CloudTrialBalance) GetBalancedOk() (*bool, bool)`

GetBalancedOk returns a tuple with the Balanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanced

`func (o *CloudTrialBalance) SetBalanced(v bool)`

SetBalanced sets Balanced field to given value.

### HasBalanced

`func (o *CloudTrialBalance) HasBalanced() bool`

HasBalanced returns a boolean if a field has been set.

### GetFrom

`func (o *CloudTrialBalance) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CloudTrialBalance) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CloudTrialBalance) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CloudTrialBalance) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetRows

`func (o *CloudTrialBalance) GetRows() []CloudTrialBalanceRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *CloudTrialBalance) GetRowsOk() (*[]CloudTrialBalanceRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *CloudTrialBalance) SetRows(v []CloudTrialBalanceRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *CloudTrialBalance) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTo

`func (o *CloudTrialBalance) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CloudTrialBalance) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CloudTrialBalance) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CloudTrialBalance) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTotalCredit

`func (o *CloudTrialBalance) GetTotalCredit() int32`

GetTotalCredit returns the TotalCredit field if non-nil, zero value otherwise.

### GetTotalCreditOk

`func (o *CloudTrialBalance) GetTotalCreditOk() (*int32, bool)`

GetTotalCreditOk returns a tuple with the TotalCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCredit

`func (o *CloudTrialBalance) SetTotalCredit(v int32)`

SetTotalCredit sets TotalCredit field to given value.

### HasTotalCredit

`func (o *CloudTrialBalance) HasTotalCredit() bool`

HasTotalCredit returns a boolean if a field has been set.

### GetTotalDebit

`func (o *CloudTrialBalance) GetTotalDebit() int32`

GetTotalDebit returns the TotalDebit field if non-nil, zero value otherwise.

### GetTotalDebitOk

`func (o *CloudTrialBalance) GetTotalDebitOk() (*int32, bool)`

GetTotalDebitOk returns a tuple with the TotalDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDebit

`func (o *CloudTrialBalance) SetTotalDebit(v int32)`

SetTotalDebit sets TotalDebit field to given value.

### HasTotalDebit

`func (o *CloudTrialBalance) HasTotalDebit() bool`

HasTotalDebit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


