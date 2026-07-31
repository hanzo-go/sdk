# CloudRevenueData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArpuCents** | Pointer to **int32** |  | [optional] 
**Customers** | Pointer to **int32** |  | [optional] 
**GeneratedAt** | Pointer to **string** |  | [optional] 
**MrrCents** | Pointer to **int32** |  | [optional] 
**PayingCustomers** | Pointer to **int32** |  | [optional] 
**PerCustomer** | Pointer to [**[]CloudRevenueCustomer**](CloudRevenueCustomer.md) |  | [optional] 
**Sources** | Pointer to [**[]CloudSourceStatus**](CloudSourceStatus.md) |  | [optional] 
**SpendTrend** | Pointer to [**[]CloudSeriesPoint**](CloudSeriesPoint.md) |  | [optional] 
**TotalBalancesCents** | Pointer to **int32** |  | [optional] 
**TotalSpendCents** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudRevenueData

`func NewCloudRevenueData() *CloudRevenueData`

NewCloudRevenueData instantiates a new CloudRevenueData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRevenueDataWithDefaults

`func NewCloudRevenueDataWithDefaults() *CloudRevenueData`

NewCloudRevenueDataWithDefaults instantiates a new CloudRevenueData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpuCents

`func (o *CloudRevenueData) GetArpuCents() int32`

GetArpuCents returns the ArpuCents field if non-nil, zero value otherwise.

### GetArpuCentsOk

`func (o *CloudRevenueData) GetArpuCentsOk() (*int32, bool)`

GetArpuCentsOk returns a tuple with the ArpuCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpuCents

`func (o *CloudRevenueData) SetArpuCents(v int32)`

SetArpuCents sets ArpuCents field to given value.

### HasArpuCents

`func (o *CloudRevenueData) HasArpuCents() bool`

HasArpuCents returns a boolean if a field has been set.

### GetCustomers

`func (o *CloudRevenueData) GetCustomers() int32`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *CloudRevenueData) GetCustomersOk() (*int32, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *CloudRevenueData) SetCustomers(v int32)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *CloudRevenueData) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *CloudRevenueData) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *CloudRevenueData) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *CloudRevenueData) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *CloudRevenueData) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetMrrCents

`func (o *CloudRevenueData) GetMrrCents() int32`

GetMrrCents returns the MrrCents field if non-nil, zero value otherwise.

### GetMrrCentsOk

`func (o *CloudRevenueData) GetMrrCentsOk() (*int32, bool)`

GetMrrCentsOk returns a tuple with the MrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrrCents

`func (o *CloudRevenueData) SetMrrCents(v int32)`

SetMrrCents sets MrrCents field to given value.

### HasMrrCents

`func (o *CloudRevenueData) HasMrrCents() bool`

HasMrrCents returns a boolean if a field has been set.

### GetPayingCustomers

`func (o *CloudRevenueData) GetPayingCustomers() int32`

GetPayingCustomers returns the PayingCustomers field if non-nil, zero value otherwise.

### GetPayingCustomersOk

`func (o *CloudRevenueData) GetPayingCustomersOk() (*int32, bool)`

GetPayingCustomersOk returns a tuple with the PayingCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayingCustomers

`func (o *CloudRevenueData) SetPayingCustomers(v int32)`

SetPayingCustomers sets PayingCustomers field to given value.

### HasPayingCustomers

`func (o *CloudRevenueData) HasPayingCustomers() bool`

HasPayingCustomers returns a boolean if a field has been set.

### GetPerCustomer

`func (o *CloudRevenueData) GetPerCustomer() []CloudRevenueCustomer`

GetPerCustomer returns the PerCustomer field if non-nil, zero value otherwise.

### GetPerCustomerOk

`func (o *CloudRevenueData) GetPerCustomerOk() (*[]CloudRevenueCustomer, bool)`

GetPerCustomerOk returns a tuple with the PerCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerCustomer

`func (o *CloudRevenueData) SetPerCustomer(v []CloudRevenueCustomer)`

SetPerCustomer sets PerCustomer field to given value.

### HasPerCustomer

`func (o *CloudRevenueData) HasPerCustomer() bool`

HasPerCustomer returns a boolean if a field has been set.

### GetSources

`func (o *CloudRevenueData) GetSources() []CloudSourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CloudRevenueData) GetSourcesOk() (*[]CloudSourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CloudRevenueData) SetSources(v []CloudSourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *CloudRevenueData) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetSpendTrend

`func (o *CloudRevenueData) GetSpendTrend() []CloudSeriesPoint`

GetSpendTrend returns the SpendTrend field if non-nil, zero value otherwise.

### GetSpendTrendOk

`func (o *CloudRevenueData) GetSpendTrendOk() (*[]CloudSeriesPoint, bool)`

GetSpendTrendOk returns a tuple with the SpendTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendTrend

`func (o *CloudRevenueData) SetSpendTrend(v []CloudSeriesPoint)`

SetSpendTrend sets SpendTrend field to given value.

### HasSpendTrend

`func (o *CloudRevenueData) HasSpendTrend() bool`

HasSpendTrend returns a boolean if a field has been set.

### GetTotalBalancesCents

`func (o *CloudRevenueData) GetTotalBalancesCents() int32`

GetTotalBalancesCents returns the TotalBalancesCents field if non-nil, zero value otherwise.

### GetTotalBalancesCentsOk

`func (o *CloudRevenueData) GetTotalBalancesCentsOk() (*int32, bool)`

GetTotalBalancesCentsOk returns a tuple with the TotalBalancesCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBalancesCents

`func (o *CloudRevenueData) SetTotalBalancesCents(v int32)`

SetTotalBalancesCents sets TotalBalancesCents field to given value.

### HasTotalBalancesCents

`func (o *CloudRevenueData) HasTotalBalancesCents() bool`

HasTotalBalancesCents returns a boolean if a field has been set.

### GetTotalSpendCents

`func (o *CloudRevenueData) GetTotalSpendCents() int32`

GetTotalSpendCents returns the TotalSpendCents field if non-nil, zero value otherwise.

### GetTotalSpendCentsOk

`func (o *CloudRevenueData) GetTotalSpendCentsOk() (*int32, bool)`

GetTotalSpendCentsOk returns a tuple with the TotalSpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpendCents

`func (o *CloudRevenueData) SetTotalSpendCents(v int32)`

SetTotalSpendCents sets TotalSpendCents field to given value.

### HasTotalSpendCents

`func (o *CloudRevenueData) HasTotalSpendCents() bool`

HasTotalSpendCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


