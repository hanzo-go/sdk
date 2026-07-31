# CommerceSaaSMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOf** | Pointer to **time.Time** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Window** | Pointer to **string** |  | [optional] 
**Revenue** | Pointer to [**CommerceRevenueMetrics**](CommerceRevenueMetrics.md) |  | [optional] 
**Subscriptions** | Pointer to [**CommerceSubscriptionMetrics**](CommerceSubscriptionMetrics.md) |  | [optional] 
**Usage** | Pointer to [**CommerceUsageMetrics**](CommerceUsageMetrics.md) |  | [optional] 
**Customers** | Pointer to [**[]CommerceCustomerRow**](CommerceCustomerRow.md) |  | [optional] 
**Orgs** | Pointer to **int32** | Total tenant organizations walked | [optional] 
**Gaps** | Pointer to **[]string** | Honest \&quot;not instrumented yet\&quot; notes — never fabricated figures | [optional] 

## Methods

### NewCommerceSaaSMetrics

`func NewCommerceSaaSMetrics() *CommerceSaaSMetrics`

NewCommerceSaaSMetrics instantiates a new CommerceSaaSMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceSaaSMetricsWithDefaults

`func NewCommerceSaaSMetricsWithDefaults() *CommerceSaaSMetrics`

NewCommerceSaaSMetricsWithDefaults instantiates a new CommerceSaaSMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsOf

`func (o *CommerceSaaSMetrics) GetAsOf() time.Time`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *CommerceSaaSMetrics) GetAsOfOk() (*time.Time, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *CommerceSaaSMetrics) SetAsOf(v time.Time)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *CommerceSaaSMetrics) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetCurrency

`func (o *CommerceSaaSMetrics) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CommerceSaaSMetrics) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CommerceSaaSMetrics) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CommerceSaaSMetrics) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetWindow

`func (o *CommerceSaaSMetrics) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *CommerceSaaSMetrics) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *CommerceSaaSMetrics) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *CommerceSaaSMetrics) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetRevenue

`func (o *CommerceSaaSMetrics) GetRevenue() CommerceRevenueMetrics`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CommerceSaaSMetrics) GetRevenueOk() (*CommerceRevenueMetrics, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CommerceSaaSMetrics) SetRevenue(v CommerceRevenueMetrics)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *CommerceSaaSMetrics) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetSubscriptions

`func (o *CommerceSaaSMetrics) GetSubscriptions() CommerceSubscriptionMetrics`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *CommerceSaaSMetrics) GetSubscriptionsOk() (*CommerceSubscriptionMetrics, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *CommerceSaaSMetrics) SetSubscriptions(v CommerceSubscriptionMetrics)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *CommerceSaaSMetrics) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetUsage

`func (o *CommerceSaaSMetrics) GetUsage() CommerceUsageMetrics`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *CommerceSaaSMetrics) GetUsageOk() (*CommerceUsageMetrics, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *CommerceSaaSMetrics) SetUsage(v CommerceUsageMetrics)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *CommerceSaaSMetrics) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetCustomers

`func (o *CommerceSaaSMetrics) GetCustomers() []CommerceCustomerRow`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *CommerceSaaSMetrics) GetCustomersOk() (*[]CommerceCustomerRow, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *CommerceSaaSMetrics) SetCustomers(v []CommerceCustomerRow)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *CommerceSaaSMetrics) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetOrgs

`func (o *CommerceSaaSMetrics) GetOrgs() int32`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *CommerceSaaSMetrics) GetOrgsOk() (*int32, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *CommerceSaaSMetrics) SetOrgs(v int32)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *CommerceSaaSMetrics) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetGaps

`func (o *CommerceSaaSMetrics) GetGaps() []string`

GetGaps returns the Gaps field if non-nil, zero value otherwise.

### GetGapsOk

`func (o *CommerceSaaSMetrics) GetGapsOk() (*[]string, bool)`

GetGapsOk returns a tuple with the Gaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaps

`func (o *CommerceSaaSMetrics) SetGaps(v []string)`

SetGaps sets Gaps field to given value.

### HasGaps

`func (o *CommerceSaaSMetrics) HasGaps() bool`

HasGaps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


