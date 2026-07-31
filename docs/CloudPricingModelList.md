# CloudPricingModelList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | Pointer to **[]map[string]map[string]interface{}** | Models are the catalog entries visible to the caller, each an opaque object exactly as the pricing source emits it, with any admin override merged on top. An admin additionally sees hidden entries, each annotated under \&quot;_overlay\&quot;. | [optional] 
**Total** | Pointer to **int32** | Total is how many models this answer carries — recounted over the visible set, not the catalog&#39;s own total. | [optional] 
**Updated** | Pointer to **map[string]interface{}** | Updated is when the catalog was last refreshed, as the pricing source recorded it. | [optional] 

## Methods

### NewCloudPricingModelList

`func NewCloudPricingModelList() *CloudPricingModelList`

NewCloudPricingModelList instantiates a new CloudPricingModelList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPricingModelListWithDefaults

`func NewCloudPricingModelListWithDefaults() *CloudPricingModelList`

NewCloudPricingModelListWithDefaults instantiates a new CloudPricingModelList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *CloudPricingModelList) GetModels() []map[string]map[string]interface{}`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *CloudPricingModelList) GetModelsOk() (*[]map[string]map[string]interface{}, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *CloudPricingModelList) SetModels(v []map[string]map[string]interface{})`

SetModels sets Models field to given value.

### HasModels

`func (o *CloudPricingModelList) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetTotal

`func (o *CloudPricingModelList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CloudPricingModelList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CloudPricingModelList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CloudPricingModelList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUpdated

`func (o *CloudPricingModelList) GetUpdated() map[string]interface{}`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CloudPricingModelList) GetUpdatedOk() (*map[string]interface{}, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CloudPricingModelList) SetUpdated(v map[string]interface{})`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CloudPricingModelList) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


