# PricingModelList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | Pointer to **[]map[string]map[string]interface{}** | Models are the catalog entries visible to the caller, each an opaque object exactly as the pricing source emits it, with any admin override merged on top. An admin additionally sees hidden entries, each annotated under \&quot;_overlay\&quot;. | [optional] 
**Total** | Pointer to **int64** | Total is how many models this answer carries — recounted over the visible set, not the catalog&#39;s own total. | [optional] 
**Updated** | Pointer to **map[string]interface{}** | Updated is when the catalog was last refreshed, as the pricing source recorded it. | [optional] 

## Methods

### NewPricingModelList

`func NewPricingModelList() *PricingModelList`

NewPricingModelList instantiates a new PricingModelList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingModelListWithDefaults

`func NewPricingModelListWithDefaults() *PricingModelList`

NewPricingModelListWithDefaults instantiates a new PricingModelList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *PricingModelList) GetModels() []map[string]map[string]interface{}`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *PricingModelList) GetModelsOk() (*[]map[string]map[string]interface{}, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *PricingModelList) SetModels(v []map[string]map[string]interface{})`

SetModels sets Models field to given value.

### HasModels

`func (o *PricingModelList) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetTotal

`func (o *PricingModelList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PricingModelList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PricingModelList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PricingModelList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUpdated

`func (o *PricingModelList) GetUpdated() map[string]interface{}`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PricingModelList) GetUpdatedOk() (*map[string]interface{}, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PricingModelList) SetUpdated(v map[string]interface{})`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PricingModelList) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


