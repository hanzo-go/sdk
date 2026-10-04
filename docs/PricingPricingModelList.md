# PricingPricingModelList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | Pointer to **[]map[string]interface{}** | Models are the catalog entries visible to the caller, each an opaque object exactly as the pricing source emits it, with any admin override merged on top. An admin additionally sees hidden entries, each annotated under \&quot;_overlay\&quot;. | [optional] 
**Total** | Pointer to **int64** | Total is how many models this answer carries — recounted over the visible set, not the catalog&#39;s own total. | [optional] 
**Updated** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPricingPricingModelList

`func NewPricingPricingModelList() *PricingPricingModelList`

NewPricingPricingModelList instantiates a new PricingPricingModelList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingPricingModelListWithDefaults

`func NewPricingPricingModelListWithDefaults() *PricingPricingModelList`

NewPricingPricingModelListWithDefaults instantiates a new PricingPricingModelList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *PricingPricingModelList) GetModels() []map[string]interface{}`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *PricingPricingModelList) GetModelsOk() (*[]map[string]interface{}, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *PricingPricingModelList) SetModels(v []map[string]interface{})`

SetModels sets Models field to given value.

### HasModels

`func (o *PricingPricingModelList) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetTotal

`func (o *PricingPricingModelList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PricingPricingModelList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PricingPricingModelList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PricingPricingModelList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUpdated

`func (o *PricingPricingModelList) GetUpdated() interface{}`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PricingPricingModelList) GetUpdatedOk() (*interface{}, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PricingPricingModelList) SetUpdated(v interface{})`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PricingPricingModelList) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *PricingPricingModelList) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *PricingPricingModelList) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


