# AdminDoCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configured** | Pointer to **bool** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**CreditRemainingCents** | Pointer to **int64** |  | [optional] 
**MonthToDateSpendCents** | Pointer to **int64** |  | [optional] 
**AvgDailyBurnCents** | Pointer to **int64** |  | [optional] 
**AccountBalanceCents** | Pointer to **int64** |  | [optional] 
**GeneratedAt** | Pointer to **string** |  | [optional] 
**History** | Pointer to [**[]AdminDoHistoryPoint**](AdminDoHistoryPoint.md) |  | [optional] 

## Methods

### NewAdminDoCost

`func NewAdminDoCost() *AdminDoCost`

NewAdminDoCost instantiates a new AdminDoCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminDoCostWithDefaults

`func NewAdminDoCostWithDefaults() *AdminDoCost`

NewAdminDoCostWithDefaults instantiates a new AdminDoCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigured

`func (o *AdminDoCost) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *AdminDoCost) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *AdminDoCost) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *AdminDoCost) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetError

`func (o *AdminDoCost) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AdminDoCost) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AdminDoCost) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AdminDoCost) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCreditRemainingCents

`func (o *AdminDoCost) GetCreditRemainingCents() int64`

GetCreditRemainingCents returns the CreditRemainingCents field if non-nil, zero value otherwise.

### GetCreditRemainingCentsOk

`func (o *AdminDoCost) GetCreditRemainingCentsOk() (*int64, bool)`

GetCreditRemainingCentsOk returns a tuple with the CreditRemainingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditRemainingCents

`func (o *AdminDoCost) SetCreditRemainingCents(v int64)`

SetCreditRemainingCents sets CreditRemainingCents field to given value.

### HasCreditRemainingCents

`func (o *AdminDoCost) HasCreditRemainingCents() bool`

HasCreditRemainingCents returns a boolean if a field has been set.

### GetMonthToDateSpendCents

`func (o *AdminDoCost) GetMonthToDateSpendCents() int64`

GetMonthToDateSpendCents returns the MonthToDateSpendCents field if non-nil, zero value otherwise.

### GetMonthToDateSpendCentsOk

`func (o *AdminDoCost) GetMonthToDateSpendCentsOk() (*int64, bool)`

GetMonthToDateSpendCentsOk returns a tuple with the MonthToDateSpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthToDateSpendCents

`func (o *AdminDoCost) SetMonthToDateSpendCents(v int64)`

SetMonthToDateSpendCents sets MonthToDateSpendCents field to given value.

### HasMonthToDateSpendCents

`func (o *AdminDoCost) HasMonthToDateSpendCents() bool`

HasMonthToDateSpendCents returns a boolean if a field has been set.

### GetAvgDailyBurnCents

`func (o *AdminDoCost) GetAvgDailyBurnCents() int64`

GetAvgDailyBurnCents returns the AvgDailyBurnCents field if non-nil, zero value otherwise.

### GetAvgDailyBurnCentsOk

`func (o *AdminDoCost) GetAvgDailyBurnCentsOk() (*int64, bool)`

GetAvgDailyBurnCentsOk returns a tuple with the AvgDailyBurnCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDailyBurnCents

`func (o *AdminDoCost) SetAvgDailyBurnCents(v int64)`

SetAvgDailyBurnCents sets AvgDailyBurnCents field to given value.

### HasAvgDailyBurnCents

`func (o *AdminDoCost) HasAvgDailyBurnCents() bool`

HasAvgDailyBurnCents returns a boolean if a field has been set.

### GetAccountBalanceCents

`func (o *AdminDoCost) GetAccountBalanceCents() int64`

GetAccountBalanceCents returns the AccountBalanceCents field if non-nil, zero value otherwise.

### GetAccountBalanceCentsOk

`func (o *AdminDoCost) GetAccountBalanceCentsOk() (*int64, bool)`

GetAccountBalanceCentsOk returns a tuple with the AccountBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBalanceCents

`func (o *AdminDoCost) SetAccountBalanceCents(v int64)`

SetAccountBalanceCents sets AccountBalanceCents field to given value.

### HasAccountBalanceCents

`func (o *AdminDoCost) HasAccountBalanceCents() bool`

HasAccountBalanceCents returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *AdminDoCost) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *AdminDoCost) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *AdminDoCost) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *AdminDoCost) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetHistory

`func (o *AdminDoCost) GetHistory() []AdminDoHistoryPoint`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *AdminDoCost) GetHistoryOk() (*[]AdminDoHistoryPoint, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *AdminDoCost) SetHistory(v []AdminDoHistoryPoint)`

SetHistory sets History field to given value.

### HasHistory

`func (o *AdminDoCost) HasHistory() bool`

HasHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


