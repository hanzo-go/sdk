# SaaSRevenue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveSubscriptions** | Pointer to **int32** |  | [optional] 
**ArrCents** | Pointer to **int32** |  | [optional] 
**ByCategory** | Pointer to [**[]SaaSCategory**](SaaSCategory.md) |  | [optional] 
**ChurnedMrrCents** | Pointer to **int32** |  | [optional] 
**MrrCents** | Pointer to **int32** |  | [optional] 
**NetNewMrrCents** | Pointer to **int32** |  | [optional] 
**NewMrrCents** | Pointer to **int32** |  | [optional] 
**PayingCustomers** | Pointer to **int32** |  | [optional] 
**Trials** | Pointer to **int32** |  | [optional] 

## Methods

### NewSaaSRevenue

`func NewSaaSRevenue() *SaaSRevenue`

NewSaaSRevenue instantiates a new SaaSRevenue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaaSRevenueWithDefaults

`func NewSaaSRevenueWithDefaults() *SaaSRevenue`

NewSaaSRevenueWithDefaults instantiates a new SaaSRevenue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveSubscriptions

`func (o *SaaSRevenue) GetActiveSubscriptions() int32`

GetActiveSubscriptions returns the ActiveSubscriptions field if non-nil, zero value otherwise.

### GetActiveSubscriptionsOk

`func (o *SaaSRevenue) GetActiveSubscriptionsOk() (*int32, bool)`

GetActiveSubscriptionsOk returns a tuple with the ActiveSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSubscriptions

`func (o *SaaSRevenue) SetActiveSubscriptions(v int32)`

SetActiveSubscriptions sets ActiveSubscriptions field to given value.

### HasActiveSubscriptions

`func (o *SaaSRevenue) HasActiveSubscriptions() bool`

HasActiveSubscriptions returns a boolean if a field has been set.

### GetArrCents

`func (o *SaaSRevenue) GetArrCents() int32`

GetArrCents returns the ArrCents field if non-nil, zero value otherwise.

### GetArrCentsOk

`func (o *SaaSRevenue) GetArrCentsOk() (*int32, bool)`

GetArrCentsOk returns a tuple with the ArrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrCents

`func (o *SaaSRevenue) SetArrCents(v int32)`

SetArrCents sets ArrCents field to given value.

### HasArrCents

`func (o *SaaSRevenue) HasArrCents() bool`

HasArrCents returns a boolean if a field has been set.

### GetByCategory

`func (o *SaaSRevenue) GetByCategory() []SaaSCategory`

GetByCategory returns the ByCategory field if non-nil, zero value otherwise.

### GetByCategoryOk

`func (o *SaaSRevenue) GetByCategoryOk() (*[]SaaSCategory, bool)`

GetByCategoryOk returns a tuple with the ByCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByCategory

`func (o *SaaSRevenue) SetByCategory(v []SaaSCategory)`

SetByCategory sets ByCategory field to given value.

### HasByCategory

`func (o *SaaSRevenue) HasByCategory() bool`

HasByCategory returns a boolean if a field has been set.

### GetChurnedMrrCents

`func (o *SaaSRevenue) GetChurnedMrrCents() int32`

GetChurnedMrrCents returns the ChurnedMrrCents field if non-nil, zero value otherwise.

### GetChurnedMrrCentsOk

`func (o *SaaSRevenue) GetChurnedMrrCentsOk() (*int32, bool)`

GetChurnedMrrCentsOk returns a tuple with the ChurnedMrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurnedMrrCents

`func (o *SaaSRevenue) SetChurnedMrrCents(v int32)`

SetChurnedMrrCents sets ChurnedMrrCents field to given value.

### HasChurnedMrrCents

`func (o *SaaSRevenue) HasChurnedMrrCents() bool`

HasChurnedMrrCents returns a boolean if a field has been set.

### GetMrrCents

`func (o *SaaSRevenue) GetMrrCents() int32`

GetMrrCents returns the MrrCents field if non-nil, zero value otherwise.

### GetMrrCentsOk

`func (o *SaaSRevenue) GetMrrCentsOk() (*int32, bool)`

GetMrrCentsOk returns a tuple with the MrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrrCents

`func (o *SaaSRevenue) SetMrrCents(v int32)`

SetMrrCents sets MrrCents field to given value.

### HasMrrCents

`func (o *SaaSRevenue) HasMrrCents() bool`

HasMrrCents returns a boolean if a field has been set.

### GetNetNewMrrCents

`func (o *SaaSRevenue) GetNetNewMrrCents() int32`

GetNetNewMrrCents returns the NetNewMrrCents field if non-nil, zero value otherwise.

### GetNetNewMrrCentsOk

`func (o *SaaSRevenue) GetNetNewMrrCentsOk() (*int32, bool)`

GetNetNewMrrCentsOk returns a tuple with the NetNewMrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetNewMrrCents

`func (o *SaaSRevenue) SetNetNewMrrCents(v int32)`

SetNetNewMrrCents sets NetNewMrrCents field to given value.

### HasNetNewMrrCents

`func (o *SaaSRevenue) HasNetNewMrrCents() bool`

HasNetNewMrrCents returns a boolean if a field has been set.

### GetNewMrrCents

`func (o *SaaSRevenue) GetNewMrrCents() int32`

GetNewMrrCents returns the NewMrrCents field if non-nil, zero value otherwise.

### GetNewMrrCentsOk

`func (o *SaaSRevenue) GetNewMrrCentsOk() (*int32, bool)`

GetNewMrrCentsOk returns a tuple with the NewMrrCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMrrCents

`func (o *SaaSRevenue) SetNewMrrCents(v int32)`

SetNewMrrCents sets NewMrrCents field to given value.

### HasNewMrrCents

`func (o *SaaSRevenue) HasNewMrrCents() bool`

HasNewMrrCents returns a boolean if a field has been set.

### GetPayingCustomers

`func (o *SaaSRevenue) GetPayingCustomers() int32`

GetPayingCustomers returns the PayingCustomers field if non-nil, zero value otherwise.

### GetPayingCustomersOk

`func (o *SaaSRevenue) GetPayingCustomersOk() (*int32, bool)`

GetPayingCustomersOk returns a tuple with the PayingCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayingCustomers

`func (o *SaaSRevenue) SetPayingCustomers(v int32)`

SetPayingCustomers sets PayingCustomers field to given value.

### HasPayingCustomers

`func (o *SaaSRevenue) HasPayingCustomers() bool`

HasPayingCustomers returns a boolean if a field has been set.

### GetTrials

`func (o *SaaSRevenue) GetTrials() int32`

GetTrials returns the Trials field if non-nil, zero value otherwise.

### GetTrialsOk

`func (o *SaaSRevenue) GetTrialsOk() (*int32, bool)`

GetTrialsOk returns a tuple with the Trials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrials

`func (o *SaaSRevenue) SetTrials(v int32)`

SetTrials sets Trials field to given value.

### HasTrials

`func (o *SaaSRevenue) HasTrials() bool`

HasTrials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


