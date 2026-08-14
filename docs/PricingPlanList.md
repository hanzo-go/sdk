# PricingPlanList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plans** | Pointer to **[]map[string]map[string]interface{}** | Plans are the plans in this section, each an opaque object exactly as the pricing source emits it — typically id, name, description, price and a feature list. | [optional] 

## Methods

### NewPricingPlanList

`func NewPricingPlanList() *PricingPlanList`

NewPricingPlanList instantiates a new PricingPlanList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingPlanListWithDefaults

`func NewPricingPlanListWithDefaults() *PricingPlanList`

NewPricingPlanListWithDefaults instantiates a new PricingPlanList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlans

`func (o *PricingPlanList) GetPlans() []map[string]map[string]interface{}`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *PricingPlanList) GetPlansOk() (*[]map[string]map[string]interface{}, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *PricingPlanList) SetPlans(v []map[string]map[string]interface{})`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *PricingPlanList) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


