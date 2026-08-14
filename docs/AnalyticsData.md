# AnalyticsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveCustomers** | Pointer to [**[]SeriesPoint**](SeriesPoint.md) | Active customers — from the usage ledger. | [optional] 
**ArpuCents** | Pointer to **int32** |  | [optional] 
**Churn** | Pointer to [**[]SeriesPoint**](SeriesPoint.md) | Churn — logo churn (count) + rate. | [optional] 
**ChurnRatePct** | Pointer to **float32** |  | [optional] 
**Computed** | Pointer to **map[string]bool** | Transparency: which metrics are backed by real data vs honest-empty. | [optional] 
**CumulativeCustomers** | Pointer to [**[]SeriesPoint**](SeriesPoint.md) |  | [optional] 
**Dau** | Pointer to **int32** |  | [optional] 
**GeneratedAt** | Pointer to **string** |  | [optional] 
**GrowthRatePct** | Pointer to **float32** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] 
**LtvCents** | Pointer to **int32** | null until churn is observed | [optional] 
**Mau** | Pointer to **int32** |  | [optional] 
**MrrCents** | Pointer to **int32** | Revenue analytics. | [optional] 
**NewCustomers** | Pointer to **int32** |  | [optional] 
**NrrPct** | Pointer to **float32** | null — needs MRR history | [optional] 
**Range** | Pointer to **string** |  | [optional] 
**Retention** | Pointer to [**RetentionGrid**](RetentionGrid.md) | Retention triangle — signup cohort × active period. | [optional] 
**Revenue** | Pointer to [**[]SeriesPoint**](SeriesPoint.md) |  | [optional] 
**Signups** | Pointer to [**[]SeriesPoint**](SeriesPoint.md) | Growth — from IAM createdTime (always real). | [optional] 
**Sources** | Pointer to [**[]SourceStatus**](SourceStatus.md) |  | [optional] 
**TopCustomers** | Pointer to [**[]AnalyticsSlice**](AnalyticsSlice.md) |  | [optional] 
**TotalCustomers** | Pointer to **int32** |  | [optional] 
**Usage** | Pointer to [**[]SeriesPoint**](SeriesPoint.md) | Usage analytics. | [optional] 
**Wau** | Pointer to **int32** |  | [optional] 

## Methods

### NewAnalyticsData

`func NewAnalyticsData() *AnalyticsData`

NewAnalyticsData instantiates a new AnalyticsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsDataWithDefaults

`func NewAnalyticsDataWithDefaults() *AnalyticsData`

NewAnalyticsDataWithDefaults instantiates a new AnalyticsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveCustomers

`func (o *AnalyticsData) GetActiveCustomers() []SeriesPoint`

GetActiveCustomers returns the ActiveCustomers field if non-nil, zero value otherwise.

### GetActiveCustomersOk

`func (o *AnalyticsData) GetActiveCustomersOk() (*[]SeriesPoint, bool)`

GetActiveCustomersOk returns a tuple with the ActiveCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCustomers

`func (o *AnalyticsData) SetActiveCustomers(v []SeriesPoint)`

SetActiveCustomers sets ActiveCustomers field to given value.

### HasActiveCustomers

`func (o *AnalyticsData) HasActiveCustomers() bool`

HasActiveCustomers returns a boolean if a field has been set.

### GetArpuCents

`func (o *AnalyticsData) GetArpuCents() int32`

GetArpuCents returns the ArpuCents field if non-nil, zero value otherwise.

### GetArpuCentsOk

`func (o *AnalyticsData) GetArpuCentsOk() (*int32, bool)`

GetArpuCentsOk returns a tuple with the ArpuCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpuCents

`func (o *AnalyticsData) SetArpuCents(v int32)`

SetArpuCents sets ArpuCents field to given value.

### HasArpuCents

`func (o *AnalyticsData) HasArpuCents() bool`

HasArpuCents returns a boolean if a field has been set.

### GetChurn

`func (o *AnalyticsData) GetChurn() []SeriesPoint`

GetChurn returns the Churn field if non-nil, zero value otherwise.

### GetChurnOk

`func (o *AnalyticsData) GetChurnOk() (*[]SeriesPoint, bool)`

GetChurnOk returns a tuple with the Churn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurn

`func (o *AnalyticsData) SetChurn(v []SeriesPoint)`

SetChurn sets Churn field to given value.

### HasChurn

`func (o *AnalyticsData) HasChurn() bool`

HasChurn returns a boolean if a field has been set.

### GetChurnRatePct

`func (o *AnalyticsData) GetChurnRatePct() float32`

GetChurnRatePct returns the ChurnRatePct field if non-nil, zero value otherwise.

### GetChurnRatePctOk

`func (o *AnalyticsData) GetChurnRatePctOk() (*float32, bool)`

GetChurnRatePctOk returns a tuple with the ChurnRatePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurnRatePct

`func (o *AnalyticsData) SetChurnRatePct(v float32)`

SetChurnRatePct sets ChurnRatePct field to given value.

### HasChurnRatePct

`func (o *AnalyticsData) HasChurnRatePct() bool`

HasChurnRatePct returns a boolean if a field has been set.

### GetComputed

`func (o *AnalyticsData) GetComputed() map[string]bool`

GetComputed returns the Computed field if non-nil, zero value otherwise.

### GetComputedOk

`func (o *AnalyticsData) GetComputedOk() (*map[string]bool, bool)`

GetComputedOk returns a tuple with the Computed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputed

`func (o *AnalyticsData) SetComputed(v map[string]bool)`

SetComputed sets Computed field to given value.

### HasComputed

`func (o *AnalyticsData) HasComputed() bool`

HasComputed returns a boolean if a field has been set.

### GetCumulativeCustomers

`func (o *AnalyticsData) GetCumulativeCustomers() []SeriesPoint`

GetCumulativeCustomers returns the CumulativeCustomers field if non-nil, zero value otherwise.

### GetCumulativeCustomersOk

`func (o *AnalyticsData) GetCumulativeCustomersOk() (*[]SeriesPoint, bool)`

GetCumulativeCustomersOk returns a tuple with the CumulativeCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeCustomers

`func (o *AnalyticsData) SetCumulativeCustomers(v []SeriesPoint)`

SetCumulativeCustomers sets CumulativeCustomers field to given value.

### HasCumulativeCustomers

`func (o *AnalyticsData) HasCumulativeCustomers() bool`

HasCumulativeCustomers returns a boolean if a field has been set.

### GetDau

`func (o *AnalyticsData) GetDau() int32`

GetDau returns the Dau field if non-nil, zero value otherwise.

### GetDauOk

`func (o *AnalyticsData) GetDauOk() (*int32, bool)`

GetDauOk returns a tuple with the Dau field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDau

`func (o *AnalyticsData) SetDau(v int32)`

SetDau sets Dau field to given value.

### HasDau

`func (o *AnalyticsData) HasDau() bool`

HasDau returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *AnalyticsData) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *AnalyticsData) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *AnalyticsData) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *AnalyticsData) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetGrowthRatePct

`func (o *AnalyticsData) GetGrowthRatePct() float32`

GetGrowthRatePct returns the GrowthRatePct field if non-nil, zero value otherwise.

### GetGrowthRatePctOk

`func (o *AnalyticsData) GetGrowthRatePctOk() (*float32, bool)`

GetGrowthRatePctOk returns a tuple with the GrowthRatePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowthRatePct

`func (o *AnalyticsData) SetGrowthRatePct(v float32)`

SetGrowthRatePct sets GrowthRatePct field to given value.

### HasGrowthRatePct

`func (o *AnalyticsData) HasGrowthRatePct() bool`

HasGrowthRatePct returns a boolean if a field has been set.

### GetInterval

`func (o *AnalyticsData) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *AnalyticsData) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *AnalyticsData) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *AnalyticsData) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLtvCents

`func (o *AnalyticsData) GetLtvCents() int32`

GetLtvCents returns the LtvCents field if non-nil, zero value otherwise.

### GetLtvCentsOk

`func (o *AnalyticsData) GetLtvCentsOk() (*int32, bool)`

GetLtvCentsOk returns a tuple with the LtvCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLtvCents

`func (o *AnalyticsData) SetLtvCents(v int32)`

SetLtvCents sets LtvCents field to given value.

### HasLtvCents

`func (o *AnalyticsData) HasLtvCents() bool`

HasLtvCents returns a boolean if a field has been set.

### GetMau

`func (o *AnalyticsData) GetMau() int32`

GetMau returns the Mau field if non-nil, zero value otherwise.

### GetMauOk

`func (o *AnalyticsData) GetMauOk() (*int32, bool)`

GetMauOk returns a tuple with the Mau field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMau

`func (o *AnalyticsData) SetMau(v int32)`

SetMau sets Mau field to given value.

### HasMau

`func (o *AnalyticsData) HasMau() bool`

HasMau returns a boolean if a field has been set.

### GetMrrCents

`func (o *AnalyticsData) GetMrrCents() int32`

GetMrrCents returns the MrrCents field if non-nil, zero value otherwise.

### GetMrrCentsOk

`func (o *AnalyticsData) GetMrrCentsOk() (*int32, bool)`

GetMrrCentsOk returns a tuple with the MrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrrCents

`func (o *AnalyticsData) SetMrrCents(v int32)`

SetMrrCents sets MrrCents field to given value.

### HasMrrCents

`func (o *AnalyticsData) HasMrrCents() bool`

HasMrrCents returns a boolean if a field has been set.

### GetNewCustomers

`func (o *AnalyticsData) GetNewCustomers() int32`

GetNewCustomers returns the NewCustomers field if non-nil, zero value otherwise.

### GetNewCustomersOk

`func (o *AnalyticsData) GetNewCustomersOk() (*int32, bool)`

GetNewCustomersOk returns a tuple with the NewCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCustomers

`func (o *AnalyticsData) SetNewCustomers(v int32)`

SetNewCustomers sets NewCustomers field to given value.

### HasNewCustomers

`func (o *AnalyticsData) HasNewCustomers() bool`

HasNewCustomers returns a boolean if a field has been set.

### GetNrrPct

`func (o *AnalyticsData) GetNrrPct() float32`

GetNrrPct returns the NrrPct field if non-nil, zero value otherwise.

### GetNrrPctOk

`func (o *AnalyticsData) GetNrrPctOk() (*float32, bool)`

GetNrrPctOk returns a tuple with the NrrPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrrPct

`func (o *AnalyticsData) SetNrrPct(v float32)`

SetNrrPct sets NrrPct field to given value.

### HasNrrPct

`func (o *AnalyticsData) HasNrrPct() bool`

HasNrrPct returns a boolean if a field has been set.

### GetRange

`func (o *AnalyticsData) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *AnalyticsData) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *AnalyticsData) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *AnalyticsData) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRetention

`func (o *AnalyticsData) GetRetention() RetentionGrid`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *AnalyticsData) GetRetentionOk() (*RetentionGrid, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *AnalyticsData) SetRetention(v RetentionGrid)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *AnalyticsData) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetRevenue

`func (o *AnalyticsData) GetRevenue() []SeriesPoint`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *AnalyticsData) GetRevenueOk() (*[]SeriesPoint, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *AnalyticsData) SetRevenue(v []SeriesPoint)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *AnalyticsData) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetSignups

`func (o *AnalyticsData) GetSignups() []SeriesPoint`

GetSignups returns the Signups field if non-nil, zero value otherwise.

### GetSignupsOk

`func (o *AnalyticsData) GetSignupsOk() (*[]SeriesPoint, bool)`

GetSignupsOk returns a tuple with the Signups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignups

`func (o *AnalyticsData) SetSignups(v []SeriesPoint)`

SetSignups sets Signups field to given value.

### HasSignups

`func (o *AnalyticsData) HasSignups() bool`

HasSignups returns a boolean if a field has been set.

### GetSources

`func (o *AnalyticsData) GetSources() []SourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *AnalyticsData) GetSourcesOk() (*[]SourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *AnalyticsData) SetSources(v []SourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *AnalyticsData) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetTopCustomers

`func (o *AnalyticsData) GetTopCustomers() []AnalyticsSlice`

GetTopCustomers returns the TopCustomers field if non-nil, zero value otherwise.

### GetTopCustomersOk

`func (o *AnalyticsData) GetTopCustomersOk() (*[]AnalyticsSlice, bool)`

GetTopCustomersOk returns a tuple with the TopCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopCustomers

`func (o *AnalyticsData) SetTopCustomers(v []AnalyticsSlice)`

SetTopCustomers sets TopCustomers field to given value.

### HasTopCustomers

`func (o *AnalyticsData) HasTopCustomers() bool`

HasTopCustomers returns a boolean if a field has been set.

### GetTotalCustomers

`func (o *AnalyticsData) GetTotalCustomers() int32`

GetTotalCustomers returns the TotalCustomers field if non-nil, zero value otherwise.

### GetTotalCustomersOk

`func (o *AnalyticsData) GetTotalCustomersOk() (*int32, bool)`

GetTotalCustomersOk returns a tuple with the TotalCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCustomers

`func (o *AnalyticsData) SetTotalCustomers(v int32)`

SetTotalCustomers sets TotalCustomers field to given value.

### HasTotalCustomers

`func (o *AnalyticsData) HasTotalCustomers() bool`

HasTotalCustomers returns a boolean if a field has been set.

### GetUsage

`func (o *AnalyticsData) GetUsage() []SeriesPoint`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AnalyticsData) GetUsageOk() (*[]SeriesPoint, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AnalyticsData) SetUsage(v []SeriesPoint)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AnalyticsData) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetWau

`func (o *AnalyticsData) GetWau() int32`

GetWau returns the Wau field if non-nil, zero value otherwise.

### GetWauOk

`func (o *AnalyticsData) GetWauOk() (*int32, bool)`

GetWauOk returns a tuple with the Wau field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWau

`func (o *AnalyticsData) SetWau(v int32)`

SetWau sets Wau field to given value.

### HasWau

`func (o *AnalyticsData) HasWau() bool`

HasWau returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


