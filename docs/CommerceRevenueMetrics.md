# CommerceRevenueMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MrrCents** | Pointer to **int64** |  | [optional] 
**ArrCents** | Pointer to **int64** |  | [optional] 
**ActiveSubscriptions** | Pointer to **int32** |  | [optional] 
**PayingCustomers** | Pointer to **int32** |  | [optional] 
**Trials** | Pointer to **int32** |  | [optional] 
**NewMrrCents** | Pointer to **int64** |  | [optional] 
**ChurnedMrrCents** | Pointer to **int64** |  | [optional] 
**NetNewMrrCents** | Pointer to **int64** |  | [optional] 
**ByCategory** | Pointer to [**[]CommerceCategoryMRR**](CommerceCategoryMRR.md) |  | [optional] 

## Methods

### NewCommerceRevenueMetrics

`func NewCommerceRevenueMetrics() *CommerceRevenueMetrics`

NewCommerceRevenueMetrics instantiates a new CommerceRevenueMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceRevenueMetricsWithDefaults

`func NewCommerceRevenueMetricsWithDefaults() *CommerceRevenueMetrics`

NewCommerceRevenueMetricsWithDefaults instantiates a new CommerceRevenueMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMrrCents

`func (o *CommerceRevenueMetrics) GetMrrCents() int64`

GetMrrCents returns the MrrCents field if non-nil, zero value otherwise.

### GetMrrCentsOk

`func (o *CommerceRevenueMetrics) GetMrrCentsOk() (*int64, bool)`

GetMrrCentsOk returns a tuple with the MrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrrCents

`func (o *CommerceRevenueMetrics) SetMrrCents(v int64)`

SetMrrCents sets MrrCents field to given value.

### HasMrrCents

`func (o *CommerceRevenueMetrics) HasMrrCents() bool`

HasMrrCents returns a boolean if a field has been set.

### GetArrCents

`func (o *CommerceRevenueMetrics) GetArrCents() int64`

GetArrCents returns the ArrCents field if non-nil, zero value otherwise.

### GetArrCentsOk

`func (o *CommerceRevenueMetrics) GetArrCentsOk() (*int64, bool)`

GetArrCentsOk returns a tuple with the ArrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrCents

`func (o *CommerceRevenueMetrics) SetArrCents(v int64)`

SetArrCents sets ArrCents field to given value.

### HasArrCents

`func (o *CommerceRevenueMetrics) HasArrCents() bool`

HasArrCents returns a boolean if a field has been set.

### GetActiveSubscriptions

`func (o *CommerceRevenueMetrics) GetActiveSubscriptions() int32`

GetActiveSubscriptions returns the ActiveSubscriptions field if non-nil, zero value otherwise.

### GetActiveSubscriptionsOk

`func (o *CommerceRevenueMetrics) GetActiveSubscriptionsOk() (*int32, bool)`

GetActiveSubscriptionsOk returns a tuple with the ActiveSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSubscriptions

`func (o *CommerceRevenueMetrics) SetActiveSubscriptions(v int32)`

SetActiveSubscriptions sets ActiveSubscriptions field to given value.

### HasActiveSubscriptions

`func (o *CommerceRevenueMetrics) HasActiveSubscriptions() bool`

HasActiveSubscriptions returns a boolean if a field has been set.

### GetPayingCustomers

`func (o *CommerceRevenueMetrics) GetPayingCustomers() int32`

GetPayingCustomers returns the PayingCustomers field if non-nil, zero value otherwise.

### GetPayingCustomersOk

`func (o *CommerceRevenueMetrics) GetPayingCustomersOk() (*int32, bool)`

GetPayingCustomersOk returns a tuple with the PayingCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayingCustomers

`func (o *CommerceRevenueMetrics) SetPayingCustomers(v int32)`

SetPayingCustomers sets PayingCustomers field to given value.

### HasPayingCustomers

`func (o *CommerceRevenueMetrics) HasPayingCustomers() bool`

HasPayingCustomers returns a boolean if a field has been set.

### GetTrials

`func (o *CommerceRevenueMetrics) GetTrials() int32`

GetTrials returns the Trials field if non-nil, zero value otherwise.

### GetTrialsOk

`func (o *CommerceRevenueMetrics) GetTrialsOk() (*int32, bool)`

GetTrialsOk returns a tuple with the Trials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrials

`func (o *CommerceRevenueMetrics) SetTrials(v int32)`

SetTrials sets Trials field to given value.

### HasTrials

`func (o *CommerceRevenueMetrics) HasTrials() bool`

HasTrials returns a boolean if a field has been set.

### GetNewMrrCents

`func (o *CommerceRevenueMetrics) GetNewMrrCents() int64`

GetNewMrrCents returns the NewMrrCents field if non-nil, zero value otherwise.

### GetNewMrrCentsOk

`func (o *CommerceRevenueMetrics) GetNewMrrCentsOk() (*int64, bool)`

GetNewMrrCentsOk returns a tuple with the NewMrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMrrCents

`func (o *CommerceRevenueMetrics) SetNewMrrCents(v int64)`

SetNewMrrCents sets NewMrrCents field to given value.

### HasNewMrrCents

`func (o *CommerceRevenueMetrics) HasNewMrrCents() bool`

HasNewMrrCents returns a boolean if a field has been set.

### GetChurnedMrrCents

`func (o *CommerceRevenueMetrics) GetChurnedMrrCents() int64`

GetChurnedMrrCents returns the ChurnedMrrCents field if non-nil, zero value otherwise.

### GetChurnedMrrCentsOk

`func (o *CommerceRevenueMetrics) GetChurnedMrrCentsOk() (*int64, bool)`

GetChurnedMrrCentsOk returns a tuple with the ChurnedMrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurnedMrrCents

`func (o *CommerceRevenueMetrics) SetChurnedMrrCents(v int64)`

SetChurnedMrrCents sets ChurnedMrrCents field to given value.

### HasChurnedMrrCents

`func (o *CommerceRevenueMetrics) HasChurnedMrrCents() bool`

HasChurnedMrrCents returns a boolean if a field has been set.

### GetNetNewMrrCents

`func (o *CommerceRevenueMetrics) GetNetNewMrrCents() int64`

GetNetNewMrrCents returns the NetNewMrrCents field if non-nil, zero value otherwise.

### GetNetNewMrrCentsOk

`func (o *CommerceRevenueMetrics) GetNetNewMrrCentsOk() (*int64, bool)`

GetNetNewMrrCentsOk returns a tuple with the NetNewMrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetNewMrrCents

`func (o *CommerceRevenueMetrics) SetNetNewMrrCents(v int64)`

SetNetNewMrrCents sets NetNewMrrCents field to given value.

### HasNetNewMrrCents

`func (o *CommerceRevenueMetrics) HasNetNewMrrCents() bool`

HasNetNewMrrCents returns a boolean if a field has been set.

### GetByCategory

`func (o *CommerceRevenueMetrics) GetByCategory() []CommerceCategoryMRR`

GetByCategory returns the ByCategory field if non-nil, zero value otherwise.

### GetByCategoryOk

`func (o *CommerceRevenueMetrics) GetByCategoryOk() (*[]CommerceCategoryMRR, bool)`

GetByCategoryOk returns a tuple with the ByCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByCategory

`func (o *CommerceRevenueMetrics) SetByCategory(v []CommerceCategoryMRR)`

SetByCategory sets ByCategory field to given value.

### HasByCategory

`func (o *CommerceRevenueMetrics) HasByCategory() bool`

HasByCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


