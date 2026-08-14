# DoCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountBalanceCents** | Pointer to **int32** |  | [optional] 
**AvgDailyBurnCents** | Pointer to **int32** |  | [optional] 
**Configured** | Pointer to **bool** |  | [optional] 
**CreditRemainingCents** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**GeneratedAt** | Pointer to **string** |  | [optional] 
**History** | Pointer to [**[]DoHistoryPoint**](DoHistoryPoint.md) |  | [optional] 
**MonthToDateSpendCents** | Pointer to **int32** |  | [optional] 

## Methods

### NewDoCost

`func NewDoCost() *DoCost`

NewDoCost instantiates a new DoCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDoCostWithDefaults

`func NewDoCostWithDefaults() *DoCost`

NewDoCostWithDefaults instantiates a new DoCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountBalanceCents

`func (o *DoCost) GetAccountBalanceCents() int32`

GetAccountBalanceCents returns the AccountBalanceCents field if non-nil, zero value otherwise.

### GetAccountBalanceCentsOk

`func (o *DoCost) GetAccountBalanceCentsOk() (*int32, bool)`

GetAccountBalanceCentsOk returns a tuple with the AccountBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBalanceCents

`func (o *DoCost) SetAccountBalanceCents(v int32)`

SetAccountBalanceCents sets AccountBalanceCents field to given value.

### HasAccountBalanceCents

`func (o *DoCost) HasAccountBalanceCents() bool`

HasAccountBalanceCents returns a boolean if a field has been set.

### GetAvgDailyBurnCents

`func (o *DoCost) GetAvgDailyBurnCents() int32`

GetAvgDailyBurnCents returns the AvgDailyBurnCents field if non-nil, zero value otherwise.

### GetAvgDailyBurnCentsOk

`func (o *DoCost) GetAvgDailyBurnCentsOk() (*int32, bool)`

GetAvgDailyBurnCentsOk returns a tuple with the AvgDailyBurnCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDailyBurnCents

`func (o *DoCost) SetAvgDailyBurnCents(v int32)`

SetAvgDailyBurnCents sets AvgDailyBurnCents field to given value.

### HasAvgDailyBurnCents

`func (o *DoCost) HasAvgDailyBurnCents() bool`

HasAvgDailyBurnCents returns a boolean if a field has been set.

### GetConfigured

`func (o *DoCost) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *DoCost) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *DoCost) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *DoCost) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetCreditRemainingCents

`func (o *DoCost) GetCreditRemainingCents() int32`

GetCreditRemainingCents returns the CreditRemainingCents field if non-nil, zero value otherwise.

### GetCreditRemainingCentsOk

`func (o *DoCost) GetCreditRemainingCentsOk() (*int32, bool)`

GetCreditRemainingCentsOk returns a tuple with the CreditRemainingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditRemainingCents

`func (o *DoCost) SetCreditRemainingCents(v int32)`

SetCreditRemainingCents sets CreditRemainingCents field to given value.

### HasCreditRemainingCents

`func (o *DoCost) HasCreditRemainingCents() bool`

HasCreditRemainingCents returns a boolean if a field has been set.

### GetError

`func (o *DoCost) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DoCost) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DoCost) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DoCost) HasError() bool`

HasError returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *DoCost) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *DoCost) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *DoCost) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *DoCost) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetHistory

`func (o *DoCost) GetHistory() []DoHistoryPoint`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *DoCost) GetHistoryOk() (*[]DoHistoryPoint, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *DoCost) SetHistory(v []DoHistoryPoint)`

SetHistory sets History field to given value.

### HasHistory

`func (o *DoCost) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetMonthToDateSpendCents

`func (o *DoCost) GetMonthToDateSpendCents() int32`

GetMonthToDateSpendCents returns the MonthToDateSpendCents field if non-nil, zero value otherwise.

### GetMonthToDateSpendCentsOk

`func (o *DoCost) GetMonthToDateSpendCentsOk() (*int32, bool)`

GetMonthToDateSpendCentsOk returns a tuple with the MonthToDateSpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthToDateSpendCents

`func (o *DoCost) SetMonthToDateSpendCents(v int32)`

SetMonthToDateSpendCents sets MonthToDateSpendCents field to given value.

### HasMonthToDateSpendCents

`func (o *DoCost) HasMonthToDateSpendCents() bool`

HasMonthToDateSpendCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


