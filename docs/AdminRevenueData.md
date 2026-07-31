# AdminRevenueData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalBalancesCents** | Pointer to **int64** |  | [optional] 
**TotalSpendCents** | Pointer to **int64** |  | [optional] 
**MrrCents** | Pointer to **int64** |  | [optional] 
**Customers** | Pointer to **int32** |  | [optional] 
**PayingCustomers** | Pointer to **int32** |  | [optional] 
**ArpuCents** | Pointer to **int64** |  | [optional] 
**PerCustomer** | Pointer to [**[]AdminRevenueCustomer**](AdminRevenueCustomer.md) |  | [optional] 
**SpendTrend** | Pointer to [**[]AdminSeriesPoint**](AdminSeriesPoint.md) |  | [optional] 
**GeneratedAt** | Pointer to **string** |  | [optional] 
**Sources** | Pointer to [**[]AdminSourceStatus**](AdminSourceStatus.md) |  | [optional] 

## Methods

### NewAdminRevenueData

`func NewAdminRevenueData() *AdminRevenueData`

NewAdminRevenueData instantiates a new AdminRevenueData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminRevenueDataWithDefaults

`func NewAdminRevenueDataWithDefaults() *AdminRevenueData`

NewAdminRevenueDataWithDefaults instantiates a new AdminRevenueData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalBalancesCents

`func (o *AdminRevenueData) GetTotalBalancesCents() int64`

GetTotalBalancesCents returns the TotalBalancesCents field if non-nil, zero value otherwise.

### GetTotalBalancesCentsOk

`func (o *AdminRevenueData) GetTotalBalancesCentsOk() (*int64, bool)`

GetTotalBalancesCentsOk returns a tuple with the TotalBalancesCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBalancesCents

`func (o *AdminRevenueData) SetTotalBalancesCents(v int64)`

SetTotalBalancesCents sets TotalBalancesCents field to given value.

### HasTotalBalancesCents

`func (o *AdminRevenueData) HasTotalBalancesCents() bool`

HasTotalBalancesCents returns a boolean if a field has been set.

### GetTotalSpendCents

`func (o *AdminRevenueData) GetTotalSpendCents() int64`

GetTotalSpendCents returns the TotalSpendCents field if non-nil, zero value otherwise.

### GetTotalSpendCentsOk

`func (o *AdminRevenueData) GetTotalSpendCentsOk() (*int64, bool)`

GetTotalSpendCentsOk returns a tuple with the TotalSpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpendCents

`func (o *AdminRevenueData) SetTotalSpendCents(v int64)`

SetTotalSpendCents sets TotalSpendCents field to given value.

### HasTotalSpendCents

`func (o *AdminRevenueData) HasTotalSpendCents() bool`

HasTotalSpendCents returns a boolean if a field has been set.

### GetMrrCents

`func (o *AdminRevenueData) GetMrrCents() int64`

GetMrrCents returns the MrrCents field if non-nil, zero value otherwise.

### GetMrrCentsOk

`func (o *AdminRevenueData) GetMrrCentsOk() (*int64, bool)`

GetMrrCentsOk returns a tuple with the MrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrrCents

`func (o *AdminRevenueData) SetMrrCents(v int64)`

SetMrrCents sets MrrCents field to given value.

### HasMrrCents

`func (o *AdminRevenueData) HasMrrCents() bool`

HasMrrCents returns a boolean if a field has been set.

### GetCustomers

`func (o *AdminRevenueData) GetCustomers() int32`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *AdminRevenueData) GetCustomersOk() (*int32, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *AdminRevenueData) SetCustomers(v int32)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *AdminRevenueData) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetPayingCustomers

`func (o *AdminRevenueData) GetPayingCustomers() int32`

GetPayingCustomers returns the PayingCustomers field if non-nil, zero value otherwise.

### GetPayingCustomersOk

`func (o *AdminRevenueData) GetPayingCustomersOk() (*int32, bool)`

GetPayingCustomersOk returns a tuple with the PayingCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayingCustomers

`func (o *AdminRevenueData) SetPayingCustomers(v int32)`

SetPayingCustomers sets PayingCustomers field to given value.

### HasPayingCustomers

`func (o *AdminRevenueData) HasPayingCustomers() bool`

HasPayingCustomers returns a boolean if a field has been set.

### GetArpuCents

`func (o *AdminRevenueData) GetArpuCents() int64`

GetArpuCents returns the ArpuCents field if non-nil, zero value otherwise.

### GetArpuCentsOk

`func (o *AdminRevenueData) GetArpuCentsOk() (*int64, bool)`

GetArpuCentsOk returns a tuple with the ArpuCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpuCents

`func (o *AdminRevenueData) SetArpuCents(v int64)`

SetArpuCents sets ArpuCents field to given value.

### HasArpuCents

`func (o *AdminRevenueData) HasArpuCents() bool`

HasArpuCents returns a boolean if a field has been set.

### GetPerCustomer

`func (o *AdminRevenueData) GetPerCustomer() []AdminRevenueCustomer`

GetPerCustomer returns the PerCustomer field if non-nil, zero value otherwise.

### GetPerCustomerOk

`func (o *AdminRevenueData) GetPerCustomerOk() (*[]AdminRevenueCustomer, bool)`

GetPerCustomerOk returns a tuple with the PerCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerCustomer

`func (o *AdminRevenueData) SetPerCustomer(v []AdminRevenueCustomer)`

SetPerCustomer sets PerCustomer field to given value.

### HasPerCustomer

`func (o *AdminRevenueData) HasPerCustomer() bool`

HasPerCustomer returns a boolean if a field has been set.

### GetSpendTrend

`func (o *AdminRevenueData) GetSpendTrend() []AdminSeriesPoint`

GetSpendTrend returns the SpendTrend field if non-nil, zero value otherwise.

### GetSpendTrendOk

`func (o *AdminRevenueData) GetSpendTrendOk() (*[]AdminSeriesPoint, bool)`

GetSpendTrendOk returns a tuple with the SpendTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendTrend

`func (o *AdminRevenueData) SetSpendTrend(v []AdminSeriesPoint)`

SetSpendTrend sets SpendTrend field to given value.

### HasSpendTrend

`func (o *AdminRevenueData) HasSpendTrend() bool`

HasSpendTrend returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *AdminRevenueData) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *AdminRevenueData) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *AdminRevenueData) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *AdminRevenueData) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetSources

`func (o *AdminRevenueData) GetSources() []AdminSourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *AdminRevenueData) GetSourcesOk() (*[]AdminSourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *AdminRevenueData) SetSources(v []AdminSourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *AdminRevenueData) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


