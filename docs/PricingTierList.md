# PricingTierList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tiers** | Pointer to **[]map[string]map[string]interface{}** | Tiers are the rentable GPU configurations, each an opaque object exactly as the pricing source emits it — typically id, name, accelerator count and model, VRAM, vCPU, memory and hourly price. | [optional] 

## Methods

### NewPricingTierList

`func NewPricingTierList() *PricingTierList`

NewPricingTierList instantiates a new PricingTierList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingTierListWithDefaults

`func NewPricingTierListWithDefaults() *PricingTierList`

NewPricingTierListWithDefaults instantiates a new PricingTierList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTiers

`func (o *PricingTierList) GetTiers() []map[string]map[string]interface{}`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *PricingTierList) GetTiersOk() (*[]map[string]map[string]interface{}, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *PricingTierList) SetTiers(v []map[string]map[string]interface{})`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *PricingTierList) HasTiers() bool`

HasTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


