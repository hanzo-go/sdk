# CloudPricingPresetList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Presets** | Pointer to **[]map[string]map[string]interface{}** | Presets are the named compute sizes, each an opaque object exactly as the pricing source emits it — typically id, name, provider slug, vCPU, memory, disk and price. | [optional] 

## Methods

### NewCloudPricingPresetList

`func NewCloudPricingPresetList() *CloudPricingPresetList`

NewCloudPricingPresetList instantiates a new CloudPricingPresetList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPricingPresetListWithDefaults

`func NewCloudPricingPresetListWithDefaults() *CloudPricingPresetList`

NewCloudPricingPresetListWithDefaults instantiates a new CloudPricingPresetList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPresets

`func (o *CloudPricingPresetList) GetPresets() []map[string]map[string]interface{}`

GetPresets returns the Presets field if non-nil, zero value otherwise.

### GetPresetsOk

`func (o *CloudPricingPresetList) GetPresetsOk() (*[]map[string]map[string]interface{}, bool)`

GetPresetsOk returns a tuple with the Presets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresets

`func (o *CloudPricingPresetList) SetPresets(v []map[string]map[string]interface{})`

SetPresets sets Presets field to given value.

### HasPresets

`func (o *CloudPricingPresetList) HasPresets() bool`

HasPresets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


