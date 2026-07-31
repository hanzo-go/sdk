# CloudDoCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountBalanceCents** | Pointer to **int32** |  | [optional] 
**AvgDailyBurnCents** | Pointer to **int32** |  | [optional] 
**Configured** | Pointer to **bool** |  | [optional] 
**CreditRemainingCents** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**GeneratedAt** | Pointer to **string** |  | [optional] 
**History** | Pointer to [**[]CloudDoHistoryPoint**](CloudDoHistoryPoint.md) |  | [optional] 
**MonthToDateSpendCents** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudDoCost

`func NewCloudDoCost() *CloudDoCost`

NewCloudDoCost instantiates a new CloudDoCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDoCostWithDefaults

`func NewCloudDoCostWithDefaults() *CloudDoCost`

NewCloudDoCostWithDefaults instantiates a new CloudDoCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountBalanceCents

`func (o *CloudDoCost) GetAccountBalanceCents() int32`

GetAccountBalanceCents returns the AccountBalanceCents field if non-nil, zero value otherwise.

### GetAccountBalanceCentsOk

`func (o *CloudDoCost) GetAccountBalanceCentsOk() (*int32, bool)`

GetAccountBalanceCentsOk returns a tuple with the AccountBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBalanceCents

`func (o *CloudDoCost) SetAccountBalanceCents(v int32)`

SetAccountBalanceCents sets AccountBalanceCents field to given value.

### HasAccountBalanceCents

`func (o *CloudDoCost) HasAccountBalanceCents() bool`

HasAccountBalanceCents returns a boolean if a field has been set.

### GetAvgDailyBurnCents

`func (o *CloudDoCost) GetAvgDailyBurnCents() int32`

GetAvgDailyBurnCents returns the AvgDailyBurnCents field if non-nil, zero value otherwise.

### GetAvgDailyBurnCentsOk

`func (o *CloudDoCost) GetAvgDailyBurnCentsOk() (*int32, bool)`

GetAvgDailyBurnCentsOk returns a tuple with the AvgDailyBurnCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDailyBurnCents

`func (o *CloudDoCost) SetAvgDailyBurnCents(v int32)`

SetAvgDailyBurnCents sets AvgDailyBurnCents field to given value.

### HasAvgDailyBurnCents

`func (o *CloudDoCost) HasAvgDailyBurnCents() bool`

HasAvgDailyBurnCents returns a boolean if a field has been set.

### GetConfigured

`func (o *CloudDoCost) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *CloudDoCost) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *CloudDoCost) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *CloudDoCost) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetCreditRemainingCents

`func (o *CloudDoCost) GetCreditRemainingCents() int32`

GetCreditRemainingCents returns the CreditRemainingCents field if non-nil, zero value otherwise.

### GetCreditRemainingCentsOk

`func (o *CloudDoCost) GetCreditRemainingCentsOk() (*int32, bool)`

GetCreditRemainingCentsOk returns a tuple with the CreditRemainingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditRemainingCents

`func (o *CloudDoCost) SetCreditRemainingCents(v int32)`

SetCreditRemainingCents sets CreditRemainingCents field to given value.

### HasCreditRemainingCents

`func (o *CloudDoCost) HasCreditRemainingCents() bool`

HasCreditRemainingCents returns a boolean if a field has been set.

### GetError

`func (o *CloudDoCost) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CloudDoCost) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CloudDoCost) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CloudDoCost) HasError() bool`

HasError returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *CloudDoCost) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *CloudDoCost) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *CloudDoCost) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *CloudDoCost) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetHistory

`func (o *CloudDoCost) GetHistory() []CloudDoHistoryPoint`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *CloudDoCost) GetHistoryOk() (*[]CloudDoHistoryPoint, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *CloudDoCost) SetHistory(v []CloudDoHistoryPoint)`

SetHistory sets History field to given value.

### HasHistory

`func (o *CloudDoCost) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetMonthToDateSpendCents

`func (o *CloudDoCost) GetMonthToDateSpendCents() int32`

GetMonthToDateSpendCents returns the MonthToDateSpendCents field if non-nil, zero value otherwise.

### GetMonthToDateSpendCentsOk

`func (o *CloudDoCost) GetMonthToDateSpendCentsOk() (*int32, bool)`

GetMonthToDateSpendCentsOk returns a tuple with the MonthToDateSpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthToDateSpendCents

`func (o *CloudDoCost) SetMonthToDateSpendCents(v int32)`

SetMonthToDateSpendCents sets MonthToDateSpendCents field to given value.

### HasMonthToDateSpendCents

`func (o *CloudDoCost) HasMonthToDateSpendCents() bool`

HasMonthToDateSpendCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


