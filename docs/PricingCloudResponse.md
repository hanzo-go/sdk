# PricingCloudResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plans** | Pointer to [**[]PricingCloudPlan**](PricingCloudPlan.md) |  | [optional] 
**Regions** | Pointer to [**[]PricingCloudRegion**](PricingCloudRegion.md) |  | [optional] 
**BlockStorage** | Pointer to [**PricingBlockStoragePricing**](PricingBlockStoragePricing.md) |  | [optional] 

## Methods

### NewPricingCloudResponse

`func NewPricingCloudResponse() *PricingCloudResponse`

NewPricingCloudResponse instantiates a new PricingCloudResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingCloudResponseWithDefaults

`func NewPricingCloudResponseWithDefaults() *PricingCloudResponse`

NewPricingCloudResponseWithDefaults instantiates a new PricingCloudResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlans

`func (o *PricingCloudResponse) GetPlans() []PricingCloudPlan`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *PricingCloudResponse) GetPlansOk() (*[]PricingCloudPlan, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *PricingCloudResponse) SetPlans(v []PricingCloudPlan)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *PricingCloudResponse) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetRegions

`func (o *PricingCloudResponse) GetRegions() []PricingCloudRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *PricingCloudResponse) GetRegionsOk() (*[]PricingCloudRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *PricingCloudResponse) SetRegions(v []PricingCloudRegion)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *PricingCloudResponse) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetBlockStorage

`func (o *PricingCloudResponse) GetBlockStorage() PricingBlockStoragePricing`

GetBlockStorage returns the BlockStorage field if non-nil, zero value otherwise.

### GetBlockStorageOk

`func (o *PricingCloudResponse) GetBlockStorageOk() (*PricingBlockStoragePricing, bool)`

GetBlockStorageOk returns a tuple with the BlockStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockStorage

`func (o *PricingCloudResponse) SetBlockStorage(v PricingBlockStoragePricing)`

SetBlockStorage sets BlockStorage field to given value.

### HasBlockStorage

`func (o *PricingCloudResponse) HasBlockStorage() bool`

HasBlockStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


