# UsageSpend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when the commerce ledger was unconfigured or unreachable. Every number below is then an honest zero, NOT a measured one. | [optional] 
**AvailableCents** | Pointer to **int64** | AvailableCents is what of that balance is still spendable. | [optional] 
**BalanceCents** | Pointer to **int64** | BalanceCents is the prepaid wallet&#39;s balance, in US cents. | [optional] 
**ByCategory** | Pointer to [**[]UsageCategorySpend**](UsageCategorySpend.md) | ByCategory is the window&#39;s spend split by ledger category, largest first. | [optional] 
**MtdCents** | Pointer to **int64** | MTDCents is commerce&#39;s authoritative month-to-date consumed figure, which is a different period from the window and is not derived from it. | [optional] 
**OverageCents** | Pointer to **int64** | OverageCents is month-to-date consumption beyond the plan&#39;s allowance. | [optional] 
**Series** | Pointer to [**[]UsageSpendPoint**](UsageSpendPoint.md) | Series is the window&#39;s spend over time, gap-filled at the window&#39;s interval. | [optional] 
**Source** | Pointer to **string** | Source names where the roll-up came from. | [optional] 
**TotalCents** | Pointer to **int64** | TotalCents is consumption over the requested window, in US cents. It is self-consistent with ByCategory and Series. | [optional] 

## Methods

### NewUsageSpend

`func NewUsageSpend() *UsageSpend`

NewUsageSpend instantiates a new UsageSpend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageSpendWithDefaults

`func NewUsageSpendWithDefaults() *UsageSpend`

NewUsageSpendWithDefaults instantiates a new UsageSpend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *UsageSpend) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *UsageSpend) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *UsageSpend) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *UsageSpend) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetAvailableCents

`func (o *UsageSpend) GetAvailableCents() int64`

GetAvailableCents returns the AvailableCents field if non-nil, zero value otherwise.

### GetAvailableCentsOk

`func (o *UsageSpend) GetAvailableCentsOk() (*int64, bool)`

GetAvailableCentsOk returns a tuple with the AvailableCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCents

`func (o *UsageSpend) SetAvailableCents(v int64)`

SetAvailableCents sets AvailableCents field to given value.

### HasAvailableCents

`func (o *UsageSpend) HasAvailableCents() bool`

HasAvailableCents returns a boolean if a field has been set.

### GetBalanceCents

`func (o *UsageSpend) GetBalanceCents() int64`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *UsageSpend) GetBalanceCentsOk() (*int64, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *UsageSpend) SetBalanceCents(v int64)`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *UsageSpend) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### GetByCategory

`func (o *UsageSpend) GetByCategory() []UsageCategorySpend`

GetByCategory returns the ByCategory field if non-nil, zero value otherwise.

### GetByCategoryOk

`func (o *UsageSpend) GetByCategoryOk() (*[]UsageCategorySpend, bool)`

GetByCategoryOk returns a tuple with the ByCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByCategory

`func (o *UsageSpend) SetByCategory(v []UsageCategorySpend)`

SetByCategory sets ByCategory field to given value.

### HasByCategory

`func (o *UsageSpend) HasByCategory() bool`

HasByCategory returns a boolean if a field has been set.

### GetMtdCents

`func (o *UsageSpend) GetMtdCents() int64`

GetMtdCents returns the MtdCents field if non-nil, zero value otherwise.

### GetMtdCentsOk

`func (o *UsageSpend) GetMtdCentsOk() (*int64, bool)`

GetMtdCentsOk returns a tuple with the MtdCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtdCents

`func (o *UsageSpend) SetMtdCents(v int64)`

SetMtdCents sets MtdCents field to given value.

### HasMtdCents

`func (o *UsageSpend) HasMtdCents() bool`

HasMtdCents returns a boolean if a field has been set.

### GetOverageCents

`func (o *UsageSpend) GetOverageCents() int64`

GetOverageCents returns the OverageCents field if non-nil, zero value otherwise.

### GetOverageCentsOk

`func (o *UsageSpend) GetOverageCentsOk() (*int64, bool)`

GetOverageCentsOk returns a tuple with the OverageCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageCents

`func (o *UsageSpend) SetOverageCents(v int64)`

SetOverageCents sets OverageCents field to given value.

### HasOverageCents

`func (o *UsageSpend) HasOverageCents() bool`

HasOverageCents returns a boolean if a field has been set.

### GetSeries

`func (o *UsageSpend) GetSeries() []UsageSpendPoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *UsageSpend) GetSeriesOk() (*[]UsageSpendPoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *UsageSpend) SetSeries(v []UsageSpendPoint)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *UsageSpend) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetSource

`func (o *UsageSpend) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UsageSpend) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UsageSpend) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *UsageSpend) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTotalCents

`func (o *UsageSpend) GetTotalCents() int64`

GetTotalCents returns the TotalCents field if non-nil, zero value otherwise.

### GetTotalCentsOk

`func (o *UsageSpend) GetTotalCentsOk() (*int64, bool)`

GetTotalCentsOk returns a tuple with the TotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCents

`func (o *UsageSpend) SetTotalCents(v int64)`

SetTotalCents sets TotalCents field to given value.

### HasTotalCents

`func (o *UsageSpend) HasTotalCents() bool`

HasTotalCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


