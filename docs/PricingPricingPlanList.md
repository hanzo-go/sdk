# PricingPricingPlanList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plans** | Pointer to **[]map[string]interface{}** | Plans are the plans in this section, each an opaque object exactly as the pricing source emits it — typically id, name, description, price and a feature list. | [optional] 

## Methods

### NewPricingPricingPlanList

`func NewPricingPricingPlanList() *PricingPricingPlanList`

NewPricingPricingPlanList instantiates a new PricingPricingPlanList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingPricingPlanListWithDefaults

`func NewPricingPricingPlanListWithDefaults() *PricingPricingPlanList`

NewPricingPricingPlanListWithDefaults instantiates a new PricingPricingPlanList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlans

`func (o *PricingPricingPlanList) GetPlans() []map[string]interface{}`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *PricingPricingPlanList) GetPlansOk() (*[]map[string]interface{}, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *PricingPricingPlanList) SetPlans(v []map[string]interface{})`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *PricingPricingPlanList) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


