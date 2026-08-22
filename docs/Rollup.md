# Rollup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to [**RollupBalance**](RollupBalance.md) |  | [optional] 
**ConsumedCents** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Included** | Pointer to [**RollupAllotment**](RollupAllotment.md) |  | [optional] 
**OverageCents** | Pointer to **int32** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Windows** | Pointer to [**[]Window**](Window.md) |  | [optional] 

## Methods

### NewRollup

`func NewRollup() *Rollup`

NewRollup instantiates a new Rollup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollupWithDefaults

`func NewRollupWithDefaults() *Rollup`

NewRollupWithDefaults instantiates a new Rollup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *Rollup) GetBalance() RollupBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Rollup) GetBalanceOk() (*RollupBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Rollup) SetBalance(v RollupBalance)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *Rollup) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetConsumedCents

`func (o *Rollup) GetConsumedCents() int32`

GetConsumedCents returns the ConsumedCents field if non-nil, zero value otherwise.

### GetConsumedCentsOk

`func (o *Rollup) GetConsumedCentsOk() (*int32, bool)`

GetConsumedCentsOk returns a tuple with the ConsumedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedCents

`func (o *Rollup) SetConsumedCents(v int32)`

SetConsumedCents sets ConsumedCents field to given value.

### HasConsumedCents

`func (o *Rollup) HasConsumedCents() bool`

HasConsumedCents returns a boolean if a field has been set.

### GetCurrency

`func (o *Rollup) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Rollup) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Rollup) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Rollup) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetIncluded

`func (o *Rollup) GetIncluded() RollupAllotment`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *Rollup) GetIncludedOk() (*RollupAllotment, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *Rollup) SetIncluded(v RollupAllotment)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *Rollup) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetOverageCents

`func (o *Rollup) GetOverageCents() int32`

GetOverageCents returns the OverageCents field if non-nil, zero value otherwise.

### GetOverageCentsOk

`func (o *Rollup) GetOverageCentsOk() (*int32, bool)`

GetOverageCentsOk returns a tuple with the OverageCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageCents

`func (o *Rollup) SetOverageCents(v int32)`

SetOverageCents sets OverageCents field to given value.

### HasOverageCents

`func (o *Rollup) HasOverageCents() bool`

HasOverageCents returns a boolean if a field has been set.

### GetPeriod

`func (o *Rollup) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Rollup) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Rollup) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *Rollup) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPlan

`func (o *Rollup) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Rollup) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Rollup) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Rollup) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetUser

`func (o *Rollup) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Rollup) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Rollup) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Rollup) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetWindows

`func (o *Rollup) GetWindows() []Window`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *Rollup) GetWindowsOk() (*[]Window, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *Rollup) SetWindows(v []Window)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *Rollup) HasWindows() bool`

HasWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


