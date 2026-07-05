# AdminAnalyticsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | Pointer to **string** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**GeneratedAt** | Pointer to **string** |  | [optional] 
**Signups** | Pointer to [**[]AdminSeriesPoint**](AdminSeriesPoint.md) |  | [optional] 
**CumulativeCustomers** | Pointer to [**[]AdminSeriesPoint**](AdminSeriesPoint.md) |  | [optional] 
**TotalCustomers** | Pointer to **int32** |  | [optional] 
**NewCustomers** | Pointer to **int32** |  | [optional] 
**GrowthRatePct** | Pointer to **float64** |  | [optional] 
**ActiveCustomers** | Pointer to [**[]AdminSeriesPoint**](AdminSeriesPoint.md) |  | [optional] 
**Dau** | Pointer to **int32** |  | [optional] 
**Wau** | Pointer to **int32** |  | [optional] 
**Mau** | Pointer to **int32** |  | [optional] 
**Retention** | Pointer to [**AdminRetentionGrid**](AdminRetentionGrid.md) |  | [optional] 
**Churn** | Pointer to [**[]AdminSeriesPoint**](AdminSeriesPoint.md) |  | [optional] 
**ChurnRatePct** | Pointer to **float64** |  | [optional] 
**MrrCents** | Pointer to **int64** |  | [optional] 
**Revenue** | Pointer to [**[]AdminSeriesPoint**](AdminSeriesPoint.md) |  | [optional] 
**ArpuCents** | Pointer to **int64** |  | [optional] 
**LtvCents** | Pointer to **int64** |  | [optional] 
**NrrPct** | Pointer to **float64** |  | [optional] 
**Usage** | Pointer to [**[]AdminSeriesPoint**](AdminSeriesPoint.md) |  | [optional] 
**TopCustomers** | Pointer to [**[]AdminAnalyticsSlice**](AdminAnalyticsSlice.md) |  | [optional] 
**Computed** | Pointer to **map[string]bool** |  | [optional] 
**Sources** | Pointer to [**[]AdminSourceStatus**](AdminSourceStatus.md) |  | [optional] 

## Methods

### NewAdminAnalyticsData

`func NewAdminAnalyticsData() *AdminAnalyticsData`

NewAdminAnalyticsData instantiates a new AdminAnalyticsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminAnalyticsDataWithDefaults

`func NewAdminAnalyticsDataWithDefaults() *AdminAnalyticsData`

NewAdminAnalyticsDataWithDefaults instantiates a new AdminAnalyticsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *AdminAnalyticsData) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *AdminAnalyticsData) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *AdminAnalyticsData) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *AdminAnalyticsData) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetInterval

`func (o *AdminAnalyticsData) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *AdminAnalyticsData) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *AdminAnalyticsData) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *AdminAnalyticsData) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *AdminAnalyticsData) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *AdminAnalyticsData) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *AdminAnalyticsData) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *AdminAnalyticsData) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetSignups

`func (o *AdminAnalyticsData) GetSignups() []AdminSeriesPoint`

GetSignups returns the Signups field if non-nil, zero value otherwise.

### GetSignupsOk

`func (o *AdminAnalyticsData) GetSignupsOk() (*[]AdminSeriesPoint, bool)`

GetSignupsOk returns a tuple with the Signups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignups

`func (o *AdminAnalyticsData) SetSignups(v []AdminSeriesPoint)`

SetSignups sets Signups field to given value.

### HasSignups

`func (o *AdminAnalyticsData) HasSignups() bool`

HasSignups returns a boolean if a field has been set.

### GetCumulativeCustomers

`func (o *AdminAnalyticsData) GetCumulativeCustomers() []AdminSeriesPoint`

GetCumulativeCustomers returns the CumulativeCustomers field if non-nil, zero value otherwise.

### GetCumulativeCustomersOk

`func (o *AdminAnalyticsData) GetCumulativeCustomersOk() (*[]AdminSeriesPoint, bool)`

GetCumulativeCustomersOk returns a tuple with the CumulativeCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeCustomers

`func (o *AdminAnalyticsData) SetCumulativeCustomers(v []AdminSeriesPoint)`

SetCumulativeCustomers sets CumulativeCustomers field to given value.

### HasCumulativeCustomers

`func (o *AdminAnalyticsData) HasCumulativeCustomers() bool`

HasCumulativeCustomers returns a boolean if a field has been set.

### GetTotalCustomers

`func (o *AdminAnalyticsData) GetTotalCustomers() int32`

GetTotalCustomers returns the TotalCustomers field if non-nil, zero value otherwise.

### GetTotalCustomersOk

`func (o *AdminAnalyticsData) GetTotalCustomersOk() (*int32, bool)`

GetTotalCustomersOk returns a tuple with the TotalCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCustomers

`func (o *AdminAnalyticsData) SetTotalCustomers(v int32)`

SetTotalCustomers sets TotalCustomers field to given value.

### HasTotalCustomers

`func (o *AdminAnalyticsData) HasTotalCustomers() bool`

HasTotalCustomers returns a boolean if a field has been set.

### GetNewCustomers

`func (o *AdminAnalyticsData) GetNewCustomers() int32`

GetNewCustomers returns the NewCustomers field if non-nil, zero value otherwise.

### GetNewCustomersOk

`func (o *AdminAnalyticsData) GetNewCustomersOk() (*int32, bool)`

GetNewCustomersOk returns a tuple with the NewCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCustomers

`func (o *AdminAnalyticsData) SetNewCustomers(v int32)`

SetNewCustomers sets NewCustomers field to given value.

### HasNewCustomers

`func (o *AdminAnalyticsData) HasNewCustomers() bool`

HasNewCustomers returns a boolean if a field has been set.

### GetGrowthRatePct

`func (o *AdminAnalyticsData) GetGrowthRatePct() float64`

GetGrowthRatePct returns the GrowthRatePct field if non-nil, zero value otherwise.

### GetGrowthRatePctOk

`func (o *AdminAnalyticsData) GetGrowthRatePctOk() (*float64, bool)`

GetGrowthRatePctOk returns a tuple with the GrowthRatePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowthRatePct

`func (o *AdminAnalyticsData) SetGrowthRatePct(v float64)`

SetGrowthRatePct sets GrowthRatePct field to given value.

### HasGrowthRatePct

`func (o *AdminAnalyticsData) HasGrowthRatePct() bool`

HasGrowthRatePct returns a boolean if a field has been set.

### GetActiveCustomers

`func (o *AdminAnalyticsData) GetActiveCustomers() []AdminSeriesPoint`

GetActiveCustomers returns the ActiveCustomers field if non-nil, zero value otherwise.

### GetActiveCustomersOk

`func (o *AdminAnalyticsData) GetActiveCustomersOk() (*[]AdminSeriesPoint, bool)`

GetActiveCustomersOk returns a tuple with the ActiveCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCustomers

`func (o *AdminAnalyticsData) SetActiveCustomers(v []AdminSeriesPoint)`

SetActiveCustomers sets ActiveCustomers field to given value.

### HasActiveCustomers

`func (o *AdminAnalyticsData) HasActiveCustomers() bool`

HasActiveCustomers returns a boolean if a field has been set.

### GetDau

`func (o *AdminAnalyticsData) GetDau() int32`

GetDau returns the Dau field if non-nil, zero value otherwise.

### GetDauOk

`func (o *AdminAnalyticsData) GetDauOk() (*int32, bool)`

GetDauOk returns a tuple with the Dau field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDau

`func (o *AdminAnalyticsData) SetDau(v int32)`

SetDau sets Dau field to given value.

### HasDau

`func (o *AdminAnalyticsData) HasDau() bool`

HasDau returns a boolean if a field has been set.

### GetWau

`func (o *AdminAnalyticsData) GetWau() int32`

GetWau returns the Wau field if non-nil, zero value otherwise.

### GetWauOk

`func (o *AdminAnalyticsData) GetWauOk() (*int32, bool)`

GetWauOk returns a tuple with the Wau field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWau

`func (o *AdminAnalyticsData) SetWau(v int32)`

SetWau sets Wau field to given value.

### HasWau

`func (o *AdminAnalyticsData) HasWau() bool`

HasWau returns a boolean if a field has been set.

### GetMau

`func (o *AdminAnalyticsData) GetMau() int32`

GetMau returns the Mau field if non-nil, zero value otherwise.

### GetMauOk

`func (o *AdminAnalyticsData) GetMauOk() (*int32, bool)`

GetMauOk returns a tuple with the Mau field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMau

`func (o *AdminAnalyticsData) SetMau(v int32)`

SetMau sets Mau field to given value.

### HasMau

`func (o *AdminAnalyticsData) HasMau() bool`

HasMau returns a boolean if a field has been set.

### GetRetention

`func (o *AdminAnalyticsData) GetRetention() AdminRetentionGrid`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *AdminAnalyticsData) GetRetentionOk() (*AdminRetentionGrid, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *AdminAnalyticsData) SetRetention(v AdminRetentionGrid)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *AdminAnalyticsData) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetChurn

`func (o *AdminAnalyticsData) GetChurn() []AdminSeriesPoint`

GetChurn returns the Churn field if non-nil, zero value otherwise.

### GetChurnOk

`func (o *AdminAnalyticsData) GetChurnOk() (*[]AdminSeriesPoint, bool)`

GetChurnOk returns a tuple with the Churn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurn

`func (o *AdminAnalyticsData) SetChurn(v []AdminSeriesPoint)`

SetChurn sets Churn field to given value.

### HasChurn

`func (o *AdminAnalyticsData) HasChurn() bool`

HasChurn returns a boolean if a field has been set.

### GetChurnRatePct

`func (o *AdminAnalyticsData) GetChurnRatePct() float64`

GetChurnRatePct returns the ChurnRatePct field if non-nil, zero value otherwise.

### GetChurnRatePctOk

`func (o *AdminAnalyticsData) GetChurnRatePctOk() (*float64, bool)`

GetChurnRatePctOk returns a tuple with the ChurnRatePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurnRatePct

`func (o *AdminAnalyticsData) SetChurnRatePct(v float64)`

SetChurnRatePct sets ChurnRatePct field to given value.

### HasChurnRatePct

`func (o *AdminAnalyticsData) HasChurnRatePct() bool`

HasChurnRatePct returns a boolean if a field has been set.

### GetMrrCents

`func (o *AdminAnalyticsData) GetMrrCents() int64`

GetMrrCents returns the MrrCents field if non-nil, zero value otherwise.

### GetMrrCentsOk

`func (o *AdminAnalyticsData) GetMrrCentsOk() (*int64, bool)`

GetMrrCentsOk returns a tuple with the MrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrrCents

`func (o *AdminAnalyticsData) SetMrrCents(v int64)`

SetMrrCents sets MrrCents field to given value.

### HasMrrCents

`func (o *AdminAnalyticsData) HasMrrCents() bool`

HasMrrCents returns a boolean if a field has been set.

### GetRevenue

`func (o *AdminAnalyticsData) GetRevenue() []AdminSeriesPoint`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *AdminAnalyticsData) GetRevenueOk() (*[]AdminSeriesPoint, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *AdminAnalyticsData) SetRevenue(v []AdminSeriesPoint)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *AdminAnalyticsData) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetArpuCents

`func (o *AdminAnalyticsData) GetArpuCents() int64`

GetArpuCents returns the ArpuCents field if non-nil, zero value otherwise.

### GetArpuCentsOk

`func (o *AdminAnalyticsData) GetArpuCentsOk() (*int64, bool)`

GetArpuCentsOk returns a tuple with the ArpuCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpuCents

`func (o *AdminAnalyticsData) SetArpuCents(v int64)`

SetArpuCents sets ArpuCents field to given value.

### HasArpuCents

`func (o *AdminAnalyticsData) HasArpuCents() bool`

HasArpuCents returns a boolean if a field has been set.

### GetLtvCents

`func (o *AdminAnalyticsData) GetLtvCents() int64`

GetLtvCents returns the LtvCents field if non-nil, zero value otherwise.

### GetLtvCentsOk

`func (o *AdminAnalyticsData) GetLtvCentsOk() (*int64, bool)`

GetLtvCentsOk returns a tuple with the LtvCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLtvCents

`func (o *AdminAnalyticsData) SetLtvCents(v int64)`

SetLtvCents sets LtvCents field to given value.

### HasLtvCents

`func (o *AdminAnalyticsData) HasLtvCents() bool`

HasLtvCents returns a boolean if a field has been set.

### GetNrrPct

`func (o *AdminAnalyticsData) GetNrrPct() float64`

GetNrrPct returns the NrrPct field if non-nil, zero value otherwise.

### GetNrrPctOk

`func (o *AdminAnalyticsData) GetNrrPctOk() (*float64, bool)`

GetNrrPctOk returns a tuple with the NrrPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrrPct

`func (o *AdminAnalyticsData) SetNrrPct(v float64)`

SetNrrPct sets NrrPct field to given value.

### HasNrrPct

`func (o *AdminAnalyticsData) HasNrrPct() bool`

HasNrrPct returns a boolean if a field has been set.

### GetUsage

`func (o *AdminAnalyticsData) GetUsage() []AdminSeriesPoint`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AdminAnalyticsData) GetUsageOk() (*[]AdminSeriesPoint, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AdminAnalyticsData) SetUsage(v []AdminSeriesPoint)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AdminAnalyticsData) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetTopCustomers

`func (o *AdminAnalyticsData) GetTopCustomers() []AdminAnalyticsSlice`

GetTopCustomers returns the TopCustomers field if non-nil, zero value otherwise.

### GetTopCustomersOk

`func (o *AdminAnalyticsData) GetTopCustomersOk() (*[]AdminAnalyticsSlice, bool)`

GetTopCustomersOk returns a tuple with the TopCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopCustomers

`func (o *AdminAnalyticsData) SetTopCustomers(v []AdminAnalyticsSlice)`

SetTopCustomers sets TopCustomers field to given value.

### HasTopCustomers

`func (o *AdminAnalyticsData) HasTopCustomers() bool`

HasTopCustomers returns a boolean if a field has been set.

### GetComputed

`func (o *AdminAnalyticsData) GetComputed() map[string]bool`

GetComputed returns the Computed field if non-nil, zero value otherwise.

### GetComputedOk

`func (o *AdminAnalyticsData) GetComputedOk() (*map[string]bool, bool)`

GetComputedOk returns a tuple with the Computed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputed

`func (o *AdminAnalyticsData) SetComputed(v map[string]bool)`

SetComputed sets Computed field to given value.

### HasComputed

`func (o *AdminAnalyticsData) HasComputed() bool`

HasComputed returns a boolean if a field has been set.

### GetSources

`func (o *AdminAnalyticsData) GetSources() []AdminSourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *AdminAnalyticsData) GetSourcesOk() (*[]AdminSourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *AdminAnalyticsData) SetSources(v []AdminSourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *AdminAnalyticsData) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


