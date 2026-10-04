# PricingPricingRegionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | Pointer to **[]map[string]interface{}** | Regions are the regions cloud instances can be placed in, each an opaque object exactly as the pricing source emits it — typically id, name and location. | [optional] 

## Methods

### NewPricingPricingRegionList

`func NewPricingPricingRegionList() *PricingPricingRegionList`

NewPricingPricingRegionList instantiates a new PricingPricingRegionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingPricingRegionListWithDefaults

`func NewPricingPricingRegionListWithDefaults() *PricingPricingRegionList`

NewPricingPricingRegionListWithDefaults instantiates a new PricingPricingRegionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *PricingPricingRegionList) GetRegions() []map[string]interface{}`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *PricingPricingRegionList) GetRegionsOk() (*[]map[string]interface{}, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *PricingPricingRegionList) SetRegions(v []map[string]interface{})`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *PricingPricingRegionList) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


