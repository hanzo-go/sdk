# BillingRollup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to [**BillingRollupBalance**](BillingRollupBalance.md) |  | [optional] 
**ConsumedCents** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Included** | Pointer to [**BillingRollupAllotment**](BillingRollupAllotment.md) |  | [optional] 
**OverageCents** | Pointer to **int64** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingRollup

`func NewBillingRollup() *BillingRollup`

NewBillingRollup instantiates a new BillingRollup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingRollupWithDefaults

`func NewBillingRollupWithDefaults() *BillingRollup`

NewBillingRollupWithDefaults instantiates a new BillingRollup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *BillingRollup) GetBalance() BillingRollupBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BillingRollup) GetBalanceOk() (*BillingRollupBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BillingRollup) SetBalance(v BillingRollupBalance)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *BillingRollup) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetConsumedCents

`func (o *BillingRollup) GetConsumedCents() int64`

GetConsumedCents returns the ConsumedCents field if non-nil, zero value otherwise.

### GetConsumedCentsOk

`func (o *BillingRollup) GetConsumedCentsOk() (*int64, bool)`

GetConsumedCentsOk returns a tuple with the ConsumedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedCents

`func (o *BillingRollup) SetConsumedCents(v int64)`

SetConsumedCents sets ConsumedCents field to given value.

### HasConsumedCents

`func (o *BillingRollup) HasConsumedCents() bool`

HasConsumedCents returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingRollup) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingRollup) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingRollup) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingRollup) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetIncluded

`func (o *BillingRollup) GetIncluded() BillingRollupAllotment`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *BillingRollup) GetIncludedOk() (*BillingRollupAllotment, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *BillingRollup) SetIncluded(v BillingRollupAllotment)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *BillingRollup) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetOverageCents

`func (o *BillingRollup) GetOverageCents() int64`

GetOverageCents returns the OverageCents field if non-nil, zero value otherwise.

### GetOverageCentsOk

`func (o *BillingRollup) GetOverageCentsOk() (*int64, bool)`

GetOverageCentsOk returns a tuple with the OverageCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageCents

`func (o *BillingRollup) SetOverageCents(v int64)`

SetOverageCents sets OverageCents field to given value.

### HasOverageCents

`func (o *BillingRollup) HasOverageCents() bool`

HasOverageCents returns a boolean if a field has been set.

### GetPeriod

`func (o *BillingRollup) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *BillingRollup) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *BillingRollup) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *BillingRollup) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPlan

`func (o *BillingRollup) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *BillingRollup) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *BillingRollup) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *BillingRollup) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetUser

`func (o *BillingRollup) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BillingRollup) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BillingRollup) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *BillingRollup) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


