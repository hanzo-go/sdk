# CloudPricingRegionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | Pointer to **[]map[string]map[string]interface{}** | Regions are the regions cloud instances can be placed in, each an opaque object exactly as the pricing source emits it — typically id, name and location. | [optional] 

## Methods

### NewCloudPricingRegionList

`func NewCloudPricingRegionList() *CloudPricingRegionList`

NewCloudPricingRegionList instantiates a new CloudPricingRegionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPricingRegionListWithDefaults

`func NewCloudPricingRegionListWithDefaults() *CloudPricingRegionList`

NewCloudPricingRegionListWithDefaults instantiates a new CloudPricingRegionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *CloudPricingRegionList) GetRegions() []map[string]map[string]interface{}`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *CloudPricingRegionList) GetRegionsOk() (*[]map[string]map[string]interface{}, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *CloudPricingRegionList) SetRegions(v []map[string]map[string]interface{})`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *CloudPricingRegionList) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


