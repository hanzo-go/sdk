# Spend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when the commerce ledger was unconfigured or unreachable. Every number below is then an honest zero, NOT a measured one. | [optional] 
**AvailableCents** | Pointer to **int32** | AvailableCents is what of that balance is still spendable. | [optional] 
**BalanceCents** | Pointer to **int32** | BalanceCents is the prepaid wallet&#39;s balance, in US cents. | [optional] 
**ByCategory** | Pointer to [**[]CategorySpend**](CategorySpend.md) | ByCategory is the window&#39;s spend split by ledger category, largest first. | [optional] 
**MtdCents** | Pointer to **int32** | MTDCents is commerce&#39;s authoritative month-to-date consumed figure, which is a different period from the window and is not derived from it. | [optional] 
**OverageCents** | Pointer to **int32** | OverageCents is month-to-date consumption beyond the plan&#39;s allowance. | [optional] 
**Series** | Pointer to [**[]SpendPoint**](SpendPoint.md) | Series is the window&#39;s spend over time, gap-filled at the window&#39;s interval. | [optional] 
**Source** | Pointer to **string** | Source names where the roll-up came from. | [optional] 
**TotalCents** | Pointer to **int32** | TotalCents is consumption over the requested window, in US cents. It is self-consistent with ByCategory and Series. | [optional] 

## Methods

### NewSpend

`func NewSpend() *Spend`

NewSpend instantiates a new Spend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendWithDefaults

`func NewSpendWithDefaults() *Spend`

NewSpendWithDefaults instantiates a new Spend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *Spend) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Spend) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Spend) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *Spend) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetAvailableCents

`func (o *Spend) GetAvailableCents() int32`

GetAvailableCents returns the AvailableCents field if non-nil, zero value otherwise.

### GetAvailableCentsOk

`func (o *Spend) GetAvailableCentsOk() (*int32, bool)`

GetAvailableCentsOk returns a tuple with the AvailableCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCents

`func (o *Spend) SetAvailableCents(v int32)`

SetAvailableCents sets AvailableCents field to given value.

### HasAvailableCents

`func (o *Spend) HasAvailableCents() bool`

HasAvailableCents returns a boolean if a field has been set.

### GetBalanceCents

`func (o *Spend) GetBalanceCents() int32`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *Spend) GetBalanceCentsOk() (*int32, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *Spend) SetBalanceCents(v int32)`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *Spend) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### GetByCategory

`func (o *Spend) GetByCategory() []CategorySpend`

GetByCategory returns the ByCategory field if non-nil, zero value otherwise.

### GetByCategoryOk

`func (o *Spend) GetByCategoryOk() (*[]CategorySpend, bool)`

GetByCategoryOk returns a tuple with the ByCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByCategory

`func (o *Spend) SetByCategory(v []CategorySpend)`

SetByCategory sets ByCategory field to given value.

### HasByCategory

`func (o *Spend) HasByCategory() bool`

HasByCategory returns a boolean if a field has been set.

### GetMtdCents

`func (o *Spend) GetMtdCents() int32`

GetMtdCents returns the MtdCents field if non-nil, zero value otherwise.

### GetMtdCentsOk

`func (o *Spend) GetMtdCentsOk() (*int32, bool)`

GetMtdCentsOk returns a tuple with the MtdCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtdCents

`func (o *Spend) SetMtdCents(v int32)`

SetMtdCents sets MtdCents field to given value.

### HasMtdCents

`func (o *Spend) HasMtdCents() bool`

HasMtdCents returns a boolean if a field has been set.

### GetOverageCents

`func (o *Spend) GetOverageCents() int32`

GetOverageCents returns the OverageCents field if non-nil, zero value otherwise.

### GetOverageCentsOk

`func (o *Spend) GetOverageCentsOk() (*int32, bool)`

GetOverageCentsOk returns a tuple with the OverageCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageCents

`func (o *Spend) SetOverageCents(v int32)`

SetOverageCents sets OverageCents field to given value.

### HasOverageCents

`func (o *Spend) HasOverageCents() bool`

HasOverageCents returns a boolean if a field has been set.

### GetSeries

`func (o *Spend) GetSeries() []SpendPoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *Spend) GetSeriesOk() (*[]SpendPoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *Spend) SetSeries(v []SpendPoint)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *Spend) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetSource

`func (o *Spend) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Spend) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Spend) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Spend) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTotalCents

`func (o *Spend) GetTotalCents() int32`

GetTotalCents returns the TotalCents field if non-nil, zero value otherwise.

### GetTotalCentsOk

`func (o *Spend) GetTotalCentsOk() (*int32, bool)`

GetTotalCentsOk returns a tuple with the TotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCents

`func (o *Spend) SetTotalCents(v int32)`

SetTotalCents sets TotalCents field to given value.

### HasTotalCents

`func (o *Spend) HasTotalCents() bool`

HasTotalCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


