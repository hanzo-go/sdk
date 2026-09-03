# TierBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditsRemaining** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DailyRemaining** | Pointer to **int64** |  | [optional] 
**EffectiveAvailable** | Pointer to **int64** |  | [optional] 
**PrepaidAvailable** | Pointer to **int64** |  | [optional] 

## Methods

### NewTierBalance

`func NewTierBalance() *TierBalance`

NewTierBalance instantiates a new TierBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierBalanceWithDefaults

`func NewTierBalanceWithDefaults() *TierBalance`

NewTierBalanceWithDefaults instantiates a new TierBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditsRemaining

`func (o *TierBalance) GetCreditsRemaining() int64`

GetCreditsRemaining returns the CreditsRemaining field if non-nil, zero value otherwise.

### GetCreditsRemainingOk

`func (o *TierBalance) GetCreditsRemainingOk() (*int64, bool)`

GetCreditsRemainingOk returns a tuple with the CreditsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsRemaining

`func (o *TierBalance) SetCreditsRemaining(v int64)`

SetCreditsRemaining sets CreditsRemaining field to given value.

### HasCreditsRemaining

`func (o *TierBalance) HasCreditsRemaining() bool`

HasCreditsRemaining returns a boolean if a field has been set.

### GetCurrency

`func (o *TierBalance) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TierBalance) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TierBalance) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TierBalance) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDailyRemaining

`func (o *TierBalance) GetDailyRemaining() int64`

GetDailyRemaining returns the DailyRemaining field if non-nil, zero value otherwise.

### GetDailyRemainingOk

`func (o *TierBalance) GetDailyRemainingOk() (*int64, bool)`

GetDailyRemainingOk returns a tuple with the DailyRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyRemaining

`func (o *TierBalance) SetDailyRemaining(v int64)`

SetDailyRemaining sets DailyRemaining field to given value.

### HasDailyRemaining

`func (o *TierBalance) HasDailyRemaining() bool`

HasDailyRemaining returns a boolean if a field has been set.

### GetEffectiveAvailable

`func (o *TierBalance) GetEffectiveAvailable() int64`

GetEffectiveAvailable returns the EffectiveAvailable field if non-nil, zero value otherwise.

### GetEffectiveAvailableOk

`func (o *TierBalance) GetEffectiveAvailableOk() (*int64, bool)`

GetEffectiveAvailableOk returns a tuple with the EffectiveAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAvailable

`func (o *TierBalance) SetEffectiveAvailable(v int64)`

SetEffectiveAvailable sets EffectiveAvailable field to given value.

### HasEffectiveAvailable

`func (o *TierBalance) HasEffectiveAvailable() bool`

HasEffectiveAvailable returns a boolean if a field has been set.

### GetPrepaidAvailable

`func (o *TierBalance) GetPrepaidAvailable() int64`

GetPrepaidAvailable returns the PrepaidAvailable field if non-nil, zero value otherwise.

### GetPrepaidAvailableOk

`func (o *TierBalance) GetPrepaidAvailableOk() (*int64, bool)`

GetPrepaidAvailableOk returns a tuple with the PrepaidAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepaidAvailable

`func (o *TierBalance) SetPrepaidAvailable(v int64)`

SetPrepaidAvailable sets PrepaidAvailable field to given value.

### HasPrepaidAvailable

`func (o *TierBalance) HasPrepaidAvailable() bool`

HasPrepaidAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


