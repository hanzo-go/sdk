# CloudAnalyticsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveCustomers** | Pointer to [**[]CloudSeriesPoint**](CloudSeriesPoint.md) | Active customers — from the usage ledger. | [optional] 
**ArpuCents** | Pointer to **int32** |  | [optional] 
**Churn** | Pointer to [**[]CloudSeriesPoint**](CloudSeriesPoint.md) | Churn — logo churn (count) + rate. | [optional] 
**ChurnRatePct** | Pointer to **float32** |  | [optional] 
**Computed** | Pointer to **map[string]bool** | Transparency: which metrics are backed by real data vs honest-empty. | [optional] 
**CumulativeCustomers** | Pointer to [**[]CloudSeriesPoint**](CloudSeriesPoint.md) |  | [optional] 
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
**Retention** | Pointer to [**CloudRetentionGrid**](CloudRetentionGrid.md) | Retention triangle — signup cohort × active period. | [optional] 
**Revenue** | Pointer to [**[]CloudSeriesPoint**](CloudSeriesPoint.md) |  | [optional] 
**Signups** | Pointer to [**[]CloudSeriesPoint**](CloudSeriesPoint.md) | Growth — from IAM createdTime (always real). | [optional] 
**Sources** | Pointer to [**[]CloudSourceStatus**](CloudSourceStatus.md) |  | [optional] 
**TopCustomers** | Pointer to [**[]CloudAnalyticsSlice**](CloudAnalyticsSlice.md) |  | [optional] 
**TotalCustomers** | Pointer to **int32** |  | [optional] 
**Usage** | Pointer to [**[]CloudSeriesPoint**](CloudSeriesPoint.md) | Usage analytics. | [optional] 
**Wau** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudAnalyticsData

`func NewCloudAnalyticsData() *CloudAnalyticsData`

NewCloudAnalyticsData instantiates a new CloudAnalyticsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAnalyticsDataWithDefaults

`func NewCloudAnalyticsDataWithDefaults() *CloudAnalyticsData`

NewCloudAnalyticsDataWithDefaults instantiates a new CloudAnalyticsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveCustomers

`func (o *CloudAnalyticsData) GetActiveCustomers() []CloudSeriesPoint`

GetActiveCustomers returns the ActiveCustomers field if non-nil, zero value otherwise.

### GetActiveCustomersOk

`func (o *CloudAnalyticsData) GetActiveCustomersOk() (*[]CloudSeriesPoint, bool)`

GetActiveCustomersOk returns a tuple with the ActiveCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCustomers

`func (o *CloudAnalyticsData) SetActiveCustomers(v []CloudSeriesPoint)`

SetActiveCustomers sets ActiveCustomers field to given value.

### HasActiveCustomers

`func (o *CloudAnalyticsData) HasActiveCustomers() bool`

HasActiveCustomers returns a boolean if a field has been set.

### GetArpuCents

`func (o *CloudAnalyticsData) GetArpuCents() int32`

GetArpuCents returns the ArpuCents field if non-nil, zero value otherwise.

### GetArpuCentsOk

`func (o *CloudAnalyticsData) GetArpuCentsOk() (*int32, bool)`

GetArpuCentsOk returns a tuple with the ArpuCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpuCents

`func (o *CloudAnalyticsData) SetArpuCents(v int32)`

SetArpuCents sets ArpuCents field to given value.

### HasArpuCents

`func (o *CloudAnalyticsData) HasArpuCents() bool`

HasArpuCents returns a boolean if a field has been set.

### GetChurn

`func (o *CloudAnalyticsData) GetChurn() []CloudSeriesPoint`

GetChurn returns the Churn field if non-nil, zero value otherwise.

### GetChurnOk

`func (o *CloudAnalyticsData) GetChurnOk() (*[]CloudSeriesPoint, bool)`

GetChurnOk returns a tuple with the Churn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurn

`func (o *CloudAnalyticsData) SetChurn(v []CloudSeriesPoint)`

SetChurn sets Churn field to given value.

### HasChurn

`func (o *CloudAnalyticsData) HasChurn() bool`

HasChurn returns a boolean if a field has been set.

### GetChurnRatePct

`func (o *CloudAnalyticsData) GetChurnRatePct() float32`

GetChurnRatePct returns the ChurnRatePct field if non-nil, zero value otherwise.

### GetChurnRatePctOk

`func (o *CloudAnalyticsData) GetChurnRatePctOk() (*float32, bool)`

GetChurnRatePctOk returns a tuple with the ChurnRatePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurnRatePct

`func (o *CloudAnalyticsData) SetChurnRatePct(v float32)`

SetChurnRatePct sets ChurnRatePct field to given value.

### HasChurnRatePct

`func (o *CloudAnalyticsData) HasChurnRatePct() bool`

HasChurnRatePct returns a boolean if a field has been set.

### GetComputed

`func (o *CloudAnalyticsData) GetComputed() map[string]bool`

GetComputed returns the Computed field if non-nil, zero value otherwise.

### GetComputedOk

`func (o *CloudAnalyticsData) GetComputedOk() (*map[string]bool, bool)`

GetComputedOk returns a tuple with the Computed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputed

`func (o *CloudAnalyticsData) SetComputed(v map[string]bool)`

SetComputed sets Computed field to given value.

### HasComputed

`func (o *CloudAnalyticsData) HasComputed() bool`

HasComputed returns a boolean if a field has been set.

### GetCumulativeCustomers

`func (o *CloudAnalyticsData) GetCumulativeCustomers() []CloudSeriesPoint`

GetCumulativeCustomers returns the CumulativeCustomers field if non-nil, zero value otherwise.

### GetCumulativeCustomersOk

`func (o *CloudAnalyticsData) GetCumulativeCustomersOk() (*[]CloudSeriesPoint, bool)`

GetCumulativeCustomersOk returns a tuple with the CumulativeCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeCustomers

`func (o *CloudAnalyticsData) SetCumulativeCustomers(v []CloudSeriesPoint)`

SetCumulativeCustomers sets CumulativeCustomers field to given value.

### HasCumulativeCustomers

`func (o *CloudAnalyticsData) HasCumulativeCustomers() bool`

HasCumulativeCustomers returns a boolean if a field has been set.

### GetDau

`func (o *CloudAnalyticsData) GetDau() int32`

GetDau returns the Dau field if non-nil, zero value otherwise.

### GetDauOk

`func (o *CloudAnalyticsData) GetDauOk() (*int32, bool)`

GetDauOk returns a tuple with the Dau field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDau

`func (o *CloudAnalyticsData) SetDau(v int32)`

SetDau sets Dau field to given value.

### HasDau

`func (o *CloudAnalyticsData) HasDau() bool`

HasDau returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *CloudAnalyticsData) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *CloudAnalyticsData) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *CloudAnalyticsData) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *CloudAnalyticsData) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetGrowthRatePct

`func (o *CloudAnalyticsData) GetGrowthRatePct() float32`

GetGrowthRatePct returns the GrowthRatePct field if non-nil, zero value otherwise.

### GetGrowthRatePctOk

`func (o *CloudAnalyticsData) GetGrowthRatePctOk() (*float32, bool)`

GetGrowthRatePctOk returns a tuple with the GrowthRatePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowthRatePct

`func (o *CloudAnalyticsData) SetGrowthRatePct(v float32)`

SetGrowthRatePct sets GrowthRatePct field to given value.

### HasGrowthRatePct

`func (o *CloudAnalyticsData) HasGrowthRatePct() bool`

HasGrowthRatePct returns a boolean if a field has been set.

### GetInterval

`func (o *CloudAnalyticsData) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CloudAnalyticsData) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CloudAnalyticsData) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CloudAnalyticsData) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLtvCents

`func (o *CloudAnalyticsData) GetLtvCents() int32`

GetLtvCents returns the LtvCents field if non-nil, zero value otherwise.

### GetLtvCentsOk

`func (o *CloudAnalyticsData) GetLtvCentsOk() (*int32, bool)`

GetLtvCentsOk returns a tuple with the LtvCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLtvCents

`func (o *CloudAnalyticsData) SetLtvCents(v int32)`

SetLtvCents sets LtvCents field to given value.

### HasLtvCents

`func (o *CloudAnalyticsData) HasLtvCents() bool`

HasLtvCents returns a boolean if a field has been set.

### GetMau

`func (o *CloudAnalyticsData) GetMau() int32`

GetMau returns the Mau field if non-nil, zero value otherwise.

### GetMauOk

`func (o *CloudAnalyticsData) GetMauOk() (*int32, bool)`

GetMauOk returns a tuple with the Mau field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMau

`func (o *CloudAnalyticsData) SetMau(v int32)`

SetMau sets Mau field to given value.

### HasMau

`func (o *CloudAnalyticsData) HasMau() bool`

HasMau returns a boolean if a field has been set.

### GetMrrCents

`func (o *CloudAnalyticsData) GetMrrCents() int32`

GetMrrCents returns the MrrCents field if non-nil, zero value otherwise.

### GetMrrCentsOk

`func (o *CloudAnalyticsData) GetMrrCentsOk() (*int32, bool)`

GetMrrCentsOk returns a tuple with the MrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrrCents

`func (o *CloudAnalyticsData) SetMrrCents(v int32)`

SetMrrCents sets MrrCents field to given value.

### HasMrrCents

`func (o *CloudAnalyticsData) HasMrrCents() bool`

HasMrrCents returns a boolean if a field has been set.

### GetNewCustomers

`func (o *CloudAnalyticsData) GetNewCustomers() int32`

GetNewCustomers returns the NewCustomers field if non-nil, zero value otherwise.

### GetNewCustomersOk

`func (o *CloudAnalyticsData) GetNewCustomersOk() (*int32, bool)`

GetNewCustomersOk returns a tuple with the NewCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCustomers

`func (o *CloudAnalyticsData) SetNewCustomers(v int32)`

SetNewCustomers sets NewCustomers field to given value.

### HasNewCustomers

`func (o *CloudAnalyticsData) HasNewCustomers() bool`

HasNewCustomers returns a boolean if a field has been set.

### GetNrrPct

`func (o *CloudAnalyticsData) GetNrrPct() float32`

GetNrrPct returns the NrrPct field if non-nil, zero value otherwise.

### GetNrrPctOk

`func (o *CloudAnalyticsData) GetNrrPctOk() (*float32, bool)`

GetNrrPctOk returns a tuple with the NrrPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrrPct

`func (o *CloudAnalyticsData) SetNrrPct(v float32)`

SetNrrPct sets NrrPct field to given value.

### HasNrrPct

`func (o *CloudAnalyticsData) HasNrrPct() bool`

HasNrrPct returns a boolean if a field has been set.

### GetRange

`func (o *CloudAnalyticsData) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CloudAnalyticsData) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CloudAnalyticsData) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *CloudAnalyticsData) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRetention

`func (o *CloudAnalyticsData) GetRetention() CloudRetentionGrid`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *CloudAnalyticsData) GetRetentionOk() (*CloudRetentionGrid, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *CloudAnalyticsData) SetRetention(v CloudRetentionGrid)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *CloudAnalyticsData) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetRevenue

`func (o *CloudAnalyticsData) GetRevenue() []CloudSeriesPoint`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CloudAnalyticsData) GetRevenueOk() (*[]CloudSeriesPoint, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CloudAnalyticsData) SetRevenue(v []CloudSeriesPoint)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *CloudAnalyticsData) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetSignups

`func (o *CloudAnalyticsData) GetSignups() []CloudSeriesPoint`

GetSignups returns the Signups field if non-nil, zero value otherwise.

### GetSignupsOk

`func (o *CloudAnalyticsData) GetSignupsOk() (*[]CloudSeriesPoint, bool)`

GetSignupsOk returns a tuple with the Signups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignups

`func (o *CloudAnalyticsData) SetSignups(v []CloudSeriesPoint)`

SetSignups sets Signups field to given value.

### HasSignups

`func (o *CloudAnalyticsData) HasSignups() bool`

HasSignups returns a boolean if a field has been set.

### GetSources

`func (o *CloudAnalyticsData) GetSources() []CloudSourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CloudAnalyticsData) GetSourcesOk() (*[]CloudSourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CloudAnalyticsData) SetSources(v []CloudSourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *CloudAnalyticsData) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetTopCustomers

`func (o *CloudAnalyticsData) GetTopCustomers() []CloudAnalyticsSlice`

GetTopCustomers returns the TopCustomers field if non-nil, zero value otherwise.

### GetTopCustomersOk

`func (o *CloudAnalyticsData) GetTopCustomersOk() (*[]CloudAnalyticsSlice, bool)`

GetTopCustomersOk returns a tuple with the TopCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopCustomers

`func (o *CloudAnalyticsData) SetTopCustomers(v []CloudAnalyticsSlice)`

SetTopCustomers sets TopCustomers field to given value.

### HasTopCustomers

`func (o *CloudAnalyticsData) HasTopCustomers() bool`

HasTopCustomers returns a boolean if a field has been set.

### GetTotalCustomers

`func (o *CloudAnalyticsData) GetTotalCustomers() int32`

GetTotalCustomers returns the TotalCustomers field if non-nil, zero value otherwise.

### GetTotalCustomersOk

`func (o *CloudAnalyticsData) GetTotalCustomersOk() (*int32, bool)`

GetTotalCustomersOk returns a tuple with the TotalCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCustomers

`func (o *CloudAnalyticsData) SetTotalCustomers(v int32)`

SetTotalCustomers sets TotalCustomers field to given value.

### HasTotalCustomers

`func (o *CloudAnalyticsData) HasTotalCustomers() bool`

HasTotalCustomers returns a boolean if a field has been set.

### GetUsage

`func (o *CloudAnalyticsData) GetUsage() []CloudSeriesPoint`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *CloudAnalyticsData) GetUsageOk() (*[]CloudSeriesPoint, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *CloudAnalyticsData) SetUsage(v []CloudSeriesPoint)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *CloudAnalyticsData) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetWau

`func (o *CloudAnalyticsData) GetWau() int32`

GetWau returns the Wau field if non-nil, zero value otherwise.

### GetWauOk

`func (o *CloudAnalyticsData) GetWauOk() (*int32, bool)`

GetWauOk returns a tuple with the Wau field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWau

`func (o *CloudAnalyticsData) SetWau(v int32)`

SetWau sets Wau field to given value.

### HasWau

`func (o *CloudAnalyticsData) HasWau() bool`

HasWau returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


