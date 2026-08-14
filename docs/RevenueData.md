# RevenueData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArpuCents** | Pointer to **int32** |  | [optional] 
**Customers** | Pointer to **int32** |  | [optional] 
**GeneratedAt** | Pointer to **string** |  | [optional] 
**MrrCents** | Pointer to **int32** |  | [optional] 
**PayingCustomers** | Pointer to **int32** |  | [optional] 
**PerCustomer** | Pointer to [**[]RevenueCustomer**](RevenueCustomer.md) |  | [optional] 
**Sources** | Pointer to [**[]SourceStatus**](SourceStatus.md) |  | [optional] 
**SpendTrend** | Pointer to [**[]SeriesPoint**](SeriesPoint.md) |  | [optional] 
**TotalBalancesCents** | Pointer to **int32** |  | [optional] 
**TotalSpendCents** | Pointer to **int32** |  | [optional] 

## Methods

### NewRevenueData

`func NewRevenueData() *RevenueData`

NewRevenueData instantiates a new RevenueData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevenueDataWithDefaults

`func NewRevenueDataWithDefaults() *RevenueData`

NewRevenueDataWithDefaults instantiates a new RevenueData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpuCents

`func (o *RevenueData) GetArpuCents() int32`

GetArpuCents returns the ArpuCents field if non-nil, zero value otherwise.

### GetArpuCentsOk

`func (o *RevenueData) GetArpuCentsOk() (*int32, bool)`

GetArpuCentsOk returns a tuple with the ArpuCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpuCents

`func (o *RevenueData) SetArpuCents(v int32)`

SetArpuCents sets ArpuCents field to given value.

### HasArpuCents

`func (o *RevenueData) HasArpuCents() bool`

HasArpuCents returns a boolean if a field has been set.

### GetCustomers

`func (o *RevenueData) GetCustomers() int32`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *RevenueData) GetCustomersOk() (*int32, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *RevenueData) SetCustomers(v int32)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *RevenueData) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *RevenueData) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *RevenueData) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *RevenueData) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *RevenueData) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetMrrCents

`func (o *RevenueData) GetMrrCents() int32`

GetMrrCents returns the MrrCents field if non-nil, zero value otherwise.

### GetMrrCentsOk

`func (o *RevenueData) GetMrrCentsOk() (*int32, bool)`

GetMrrCentsOk returns a tuple with the MrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrrCents

`func (o *RevenueData) SetMrrCents(v int32)`

SetMrrCents sets MrrCents field to given value.

### HasMrrCents

`func (o *RevenueData) HasMrrCents() bool`

HasMrrCents returns a boolean if a field has been set.

### GetPayingCustomers

`func (o *RevenueData) GetPayingCustomers() int32`

GetPayingCustomers returns the PayingCustomers field if non-nil, zero value otherwise.

### GetPayingCustomersOk

`func (o *RevenueData) GetPayingCustomersOk() (*int32, bool)`

GetPayingCustomersOk returns a tuple with the PayingCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayingCustomers

`func (o *RevenueData) SetPayingCustomers(v int32)`

SetPayingCustomers sets PayingCustomers field to given value.

### HasPayingCustomers

`func (o *RevenueData) HasPayingCustomers() bool`

HasPayingCustomers returns a boolean if a field has been set.

### GetPerCustomer

`func (o *RevenueData) GetPerCustomer() []RevenueCustomer`

GetPerCustomer returns the PerCustomer field if non-nil, zero value otherwise.

### GetPerCustomerOk

`func (o *RevenueData) GetPerCustomerOk() (*[]RevenueCustomer, bool)`

GetPerCustomerOk returns a tuple with the PerCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerCustomer

`func (o *RevenueData) SetPerCustomer(v []RevenueCustomer)`

SetPerCustomer sets PerCustomer field to given value.

### HasPerCustomer

`func (o *RevenueData) HasPerCustomer() bool`

HasPerCustomer returns a boolean if a field has been set.

### GetSources

`func (o *RevenueData) GetSources() []SourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *RevenueData) GetSourcesOk() (*[]SourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *RevenueData) SetSources(v []SourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *RevenueData) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetSpendTrend

`func (o *RevenueData) GetSpendTrend() []SeriesPoint`

GetSpendTrend returns the SpendTrend field if non-nil, zero value otherwise.

### GetSpendTrendOk

`func (o *RevenueData) GetSpendTrendOk() (*[]SeriesPoint, bool)`

GetSpendTrendOk returns a tuple with the SpendTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendTrend

`func (o *RevenueData) SetSpendTrend(v []SeriesPoint)`

SetSpendTrend sets SpendTrend field to given value.

### HasSpendTrend

`func (o *RevenueData) HasSpendTrend() bool`

HasSpendTrend returns a boolean if a field has been set.

### GetTotalBalancesCents

`func (o *RevenueData) GetTotalBalancesCents() int32`

GetTotalBalancesCents returns the TotalBalancesCents field if non-nil, zero value otherwise.

### GetTotalBalancesCentsOk

`func (o *RevenueData) GetTotalBalancesCentsOk() (*int32, bool)`

GetTotalBalancesCentsOk returns a tuple with the TotalBalancesCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBalancesCents

`func (o *RevenueData) SetTotalBalancesCents(v int32)`

SetTotalBalancesCents sets TotalBalancesCents field to given value.

### HasTotalBalancesCents

`func (o *RevenueData) HasTotalBalancesCents() bool`

HasTotalBalancesCents returns a boolean if a field has been set.

### GetTotalSpendCents

`func (o *RevenueData) GetTotalSpendCents() int32`

GetTotalSpendCents returns the TotalSpendCents field if non-nil, zero value otherwise.

### GetTotalSpendCentsOk

`func (o *RevenueData) GetTotalSpendCentsOk() (*int32, bool)`

GetTotalSpendCentsOk returns a tuple with the TotalSpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpendCents

`func (o *RevenueData) SetTotalSpendCents(v int32)`

SetTotalSpendCents sets TotalSpendCents field to given value.

### HasTotalSpendCents

`func (o *RevenueData) HasTotalSpendCents() bool`

HasTotalSpendCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


