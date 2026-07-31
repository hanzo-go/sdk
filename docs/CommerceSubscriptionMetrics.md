# CommerceSubscriptionMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByPlan** | Pointer to [**[]CommercePlanBreakdown**](CommercePlanBreakdown.md) |  | [optional] 
**TrialsActive** | Pointer to **int32** |  | [optional] 
**New** | Pointer to **int32** |  | [optional] 
**Canceled** | Pointer to **int32** |  | [optional] 
**Upgrades** | Pointer to **int32** | null — plan-change events are not instrumented | [optional] 
**Downgrades** | Pointer to **int32** | null — plan-change events are not instrumented | [optional] 
**Recent** | Pointer to [**[]CommerceSubEvent**](CommerceSubEvent.md) |  | [optional] 

## Methods

### NewCommerceSubscriptionMetrics

`func NewCommerceSubscriptionMetrics() *CommerceSubscriptionMetrics`

NewCommerceSubscriptionMetrics instantiates a new CommerceSubscriptionMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceSubscriptionMetricsWithDefaults

`func NewCommerceSubscriptionMetricsWithDefaults() *CommerceSubscriptionMetrics`

NewCommerceSubscriptionMetricsWithDefaults instantiates a new CommerceSubscriptionMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByPlan

`func (o *CommerceSubscriptionMetrics) GetByPlan() []CommercePlanBreakdown`

GetByPlan returns the ByPlan field if non-nil, zero value otherwise.

### GetByPlanOk

`func (o *CommerceSubscriptionMetrics) GetByPlanOk() (*[]CommercePlanBreakdown, bool)`

GetByPlanOk returns a tuple with the ByPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByPlan

`func (o *CommerceSubscriptionMetrics) SetByPlan(v []CommercePlanBreakdown)`

SetByPlan sets ByPlan field to given value.

### HasByPlan

`func (o *CommerceSubscriptionMetrics) HasByPlan() bool`

HasByPlan returns a boolean if a field has been set.

### GetTrialsActive

`func (o *CommerceSubscriptionMetrics) GetTrialsActive() int32`

GetTrialsActive returns the TrialsActive field if non-nil, zero value otherwise.

### GetTrialsActiveOk

`func (o *CommerceSubscriptionMetrics) GetTrialsActiveOk() (*int32, bool)`

GetTrialsActiveOk returns a tuple with the TrialsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialsActive

`func (o *CommerceSubscriptionMetrics) SetTrialsActive(v int32)`

SetTrialsActive sets TrialsActive field to given value.

### HasTrialsActive

`func (o *CommerceSubscriptionMetrics) HasTrialsActive() bool`

HasTrialsActive returns a boolean if a field has been set.

### GetNew

`func (o *CommerceSubscriptionMetrics) GetNew() int32`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *CommerceSubscriptionMetrics) GetNewOk() (*int32, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *CommerceSubscriptionMetrics) SetNew(v int32)`

SetNew sets New field to given value.

### HasNew

`func (o *CommerceSubscriptionMetrics) HasNew() bool`

HasNew returns a boolean if a field has been set.

### GetCanceled

`func (o *CommerceSubscriptionMetrics) GetCanceled() int32`

GetCanceled returns the Canceled field if non-nil, zero value otherwise.

### GetCanceledOk

`func (o *CommerceSubscriptionMetrics) GetCanceledOk() (*int32, bool)`

GetCanceledOk returns a tuple with the Canceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceled

`func (o *CommerceSubscriptionMetrics) SetCanceled(v int32)`

SetCanceled sets Canceled field to given value.

### HasCanceled

`func (o *CommerceSubscriptionMetrics) HasCanceled() bool`

HasCanceled returns a boolean if a field has been set.

### GetUpgrades

`func (o *CommerceSubscriptionMetrics) GetUpgrades() int32`

GetUpgrades returns the Upgrades field if non-nil, zero value otherwise.

### GetUpgradesOk

`func (o *CommerceSubscriptionMetrics) GetUpgradesOk() (*int32, bool)`

GetUpgradesOk returns a tuple with the Upgrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrades

`func (o *CommerceSubscriptionMetrics) SetUpgrades(v int32)`

SetUpgrades sets Upgrades field to given value.

### HasUpgrades

`func (o *CommerceSubscriptionMetrics) HasUpgrades() bool`

HasUpgrades returns a boolean if a field has been set.

### GetDowngrades

`func (o *CommerceSubscriptionMetrics) GetDowngrades() int32`

GetDowngrades returns the Downgrades field if non-nil, zero value otherwise.

### GetDowngradesOk

`func (o *CommerceSubscriptionMetrics) GetDowngradesOk() (*int32, bool)`

GetDowngradesOk returns a tuple with the Downgrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowngrades

`func (o *CommerceSubscriptionMetrics) SetDowngrades(v int32)`

SetDowngrades sets Downgrades field to given value.

### HasDowngrades

`func (o *CommerceSubscriptionMetrics) HasDowngrades() bool`

HasDowngrades returns a boolean if a field has been set.

### GetRecent

`func (o *CommerceSubscriptionMetrics) GetRecent() []CommerceSubEvent`

GetRecent returns the Recent field if non-nil, zero value otherwise.

### GetRecentOk

`func (o *CommerceSubscriptionMetrics) GetRecentOk() (*[]CommerceSubEvent, bool)`

GetRecentOk returns a tuple with the Recent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecent

`func (o *CommerceSubscriptionMetrics) SetRecent(v []CommerceSubEvent)`

SetRecent sets Recent field to given value.

### HasRecent

`func (o *CommerceSubscriptionMetrics) HasRecent() bool`

HasRecent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


