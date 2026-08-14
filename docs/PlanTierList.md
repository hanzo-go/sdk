# PlanTierList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tiers** | Pointer to **[]interface{}** | Tiers are the rentable GPU configurations, each an opaque object exactly as the catalog emits it — typically id, name, GPU count and model, VRAM, vCPUs, host memory and hourly price. | [optional] 

## Methods

### NewPlanTierList

`func NewPlanTierList() *PlanTierList`

NewPlanTierList instantiates a new PlanTierList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanTierListWithDefaults

`func NewPlanTierListWithDefaults() *PlanTierList`

NewPlanTierListWithDefaults instantiates a new PlanTierList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTiers

`func (o *PlanTierList) GetTiers() []interface{}`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *PlanTierList) GetTiersOk() (*[]interface{}, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *PlanTierList) SetTiers(v []interface{})`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *PlanTierList) HasTiers() bool`

HasTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


