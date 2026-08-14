# MetricsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOf** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Customers** | Pointer to [**[]SaaSCustomer**](SaaSCustomer.md) |  | [optional] 
**Gaps** | Pointer to **[]string** |  | [optional] 
**GeneratedAt** | Pointer to **string** |  | [optional] 
**Orgs** | Pointer to **int32** |  | [optional] 
**Revenue** | Pointer to [**SaaSRevenue**](SaaSRevenue.md) |  | [optional] 
**Sources** | Pointer to [**[]SourceStatus**](SourceStatus.md) |  | [optional] 
**Subscriptions** | Pointer to [**SaaSSubs**](SaaSSubs.md) |  | [optional] 
**Usage** | Pointer to [**SaaSUsage**](SaaSUsage.md) |  | [optional] 
**Window** | Pointer to **string** |  | [optional] 

## Methods

### NewMetricsData

`func NewMetricsData() *MetricsData`

NewMetricsData instantiates a new MetricsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsDataWithDefaults

`func NewMetricsDataWithDefaults() *MetricsData`

NewMetricsDataWithDefaults instantiates a new MetricsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsOf

`func (o *MetricsData) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *MetricsData) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *MetricsData) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *MetricsData) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetCurrency

`func (o *MetricsData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MetricsData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MetricsData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MetricsData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomers

`func (o *MetricsData) GetCustomers() []SaaSCustomer`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *MetricsData) GetCustomersOk() (*[]SaaSCustomer, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *MetricsData) SetCustomers(v []SaaSCustomer)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *MetricsData) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetGaps

`func (o *MetricsData) GetGaps() []string`

GetGaps returns the Gaps field if non-nil, zero value otherwise.

### GetGapsOk

`func (o *MetricsData) GetGapsOk() (*[]string, bool)`

GetGapsOk returns a tuple with the Gaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaps

`func (o *MetricsData) SetGaps(v []string)`

SetGaps sets Gaps field to given value.

### HasGaps

`func (o *MetricsData) HasGaps() bool`

HasGaps returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *MetricsData) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *MetricsData) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *MetricsData) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *MetricsData) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetOrgs

`func (o *MetricsData) GetOrgs() int32`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *MetricsData) GetOrgsOk() (*int32, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *MetricsData) SetOrgs(v int32)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *MetricsData) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetRevenue

`func (o *MetricsData) GetRevenue() SaaSRevenue`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *MetricsData) GetRevenueOk() (*SaaSRevenue, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *MetricsData) SetRevenue(v SaaSRevenue)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *MetricsData) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetSources

`func (o *MetricsData) GetSources() []SourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *MetricsData) GetSourcesOk() (*[]SourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *MetricsData) SetSources(v []SourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *MetricsData) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetSubscriptions

`func (o *MetricsData) GetSubscriptions() SaaSSubs`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *MetricsData) GetSubscriptionsOk() (*SaaSSubs, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *MetricsData) SetSubscriptions(v SaaSSubs)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *MetricsData) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetUsage

`func (o *MetricsData) GetUsage() SaaSUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *MetricsData) GetUsageOk() (*SaaSUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *MetricsData) SetUsage(v SaaSUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *MetricsData) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetWindow

`func (o *MetricsData) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *MetricsData) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *MetricsData) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *MetricsData) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


