# CloudMetricsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOf** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Customers** | Pointer to [**[]CloudSaaSCustomer**](CloudSaaSCustomer.md) |  | [optional] 
**Gaps** | Pointer to **[]string** |  | [optional] 
**GeneratedAt** | Pointer to **string** |  | [optional] 
**Orgs** | Pointer to **int32** |  | [optional] 
**Revenue** | Pointer to [**CloudSaaSRevenue**](CloudSaaSRevenue.md) |  | [optional] 
**Sources** | Pointer to [**[]CloudSourceStatus**](CloudSourceStatus.md) |  | [optional] 
**Subscriptions** | Pointer to [**CloudSaaSSubs**](CloudSaaSSubs.md) |  | [optional] 
**Usage** | Pointer to [**CloudSaaSUsage**](CloudSaaSUsage.md) |  | [optional] 
**Window** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudMetricsData

`func NewCloudMetricsData() *CloudMetricsData`

NewCloudMetricsData instantiates a new CloudMetricsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMetricsDataWithDefaults

`func NewCloudMetricsDataWithDefaults() *CloudMetricsData`

NewCloudMetricsDataWithDefaults instantiates a new CloudMetricsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsOf

`func (o *CloudMetricsData) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *CloudMetricsData) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *CloudMetricsData) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *CloudMetricsData) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetCurrency

`func (o *CloudMetricsData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudMetricsData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudMetricsData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudMetricsData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomers

`func (o *CloudMetricsData) GetCustomers() []CloudSaaSCustomer`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *CloudMetricsData) GetCustomersOk() (*[]CloudSaaSCustomer, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *CloudMetricsData) SetCustomers(v []CloudSaaSCustomer)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *CloudMetricsData) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetGaps

`func (o *CloudMetricsData) GetGaps() []string`

GetGaps returns the Gaps field if non-nil, zero value otherwise.

### GetGapsOk

`func (o *CloudMetricsData) GetGapsOk() (*[]string, bool)`

GetGapsOk returns a tuple with the Gaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaps

`func (o *CloudMetricsData) SetGaps(v []string)`

SetGaps sets Gaps field to given value.

### HasGaps

`func (o *CloudMetricsData) HasGaps() bool`

HasGaps returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *CloudMetricsData) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *CloudMetricsData) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *CloudMetricsData) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *CloudMetricsData) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetOrgs

`func (o *CloudMetricsData) GetOrgs() int32`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *CloudMetricsData) GetOrgsOk() (*int32, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *CloudMetricsData) SetOrgs(v int32)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *CloudMetricsData) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetRevenue

`func (o *CloudMetricsData) GetRevenue() CloudSaaSRevenue`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CloudMetricsData) GetRevenueOk() (*CloudSaaSRevenue, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CloudMetricsData) SetRevenue(v CloudSaaSRevenue)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *CloudMetricsData) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetSources

`func (o *CloudMetricsData) GetSources() []CloudSourceStatus`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CloudMetricsData) GetSourcesOk() (*[]CloudSourceStatus, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CloudMetricsData) SetSources(v []CloudSourceStatus)`

SetSources sets Sources field to given value.

### HasSources

`func (o *CloudMetricsData) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetSubscriptions

`func (o *CloudMetricsData) GetSubscriptions() CloudSaaSSubs`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *CloudMetricsData) GetSubscriptionsOk() (*CloudSaaSSubs, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *CloudMetricsData) SetSubscriptions(v CloudSaaSSubs)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *CloudMetricsData) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetUsage

`func (o *CloudMetricsData) GetUsage() CloudSaaSUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *CloudMetricsData) GetUsageOk() (*CloudSaaSUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *CloudMetricsData) SetUsage(v CloudSaaSUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *CloudMetricsData) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetWindow

`func (o *CloudMetricsData) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *CloudMetricsData) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *CloudMetricsData) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *CloudMetricsData) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


