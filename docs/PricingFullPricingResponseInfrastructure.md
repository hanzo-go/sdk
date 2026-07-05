# PricingFullPricingResponseInfrastructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to **map[string]interface{}** |  | [optional] 
**Gpu** | Pointer to [**[]PricingGpuTier**](PricingGpuTier.md) |  | [optional] 

## Methods

### NewPricingFullPricingResponseInfrastructure

`func NewPricingFullPricingResponseInfrastructure() *PricingFullPricingResponseInfrastructure`

NewPricingFullPricingResponseInfrastructure instantiates a new PricingFullPricingResponseInfrastructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingFullPricingResponseInfrastructureWithDefaults

`func NewPricingFullPricingResponseInfrastructureWithDefaults() *PricingFullPricingResponseInfrastructure`

NewPricingFullPricingResponseInfrastructureWithDefaults instantiates a new PricingFullPricingResponseInfrastructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *PricingFullPricingResponseInfrastructure) GetCompute() map[string]interface{}`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *PricingFullPricingResponseInfrastructure) GetComputeOk() (*map[string]interface{}, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *PricingFullPricingResponseInfrastructure) SetCompute(v map[string]interface{})`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *PricingFullPricingResponseInfrastructure) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetGpu

`func (o *PricingFullPricingResponseInfrastructure) GetGpu() []PricingGpuTier`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *PricingFullPricingResponseInfrastructure) GetGpuOk() (*[]PricingGpuTier, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *PricingFullPricingResponseInfrastructure) SetGpu(v []PricingGpuTier)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *PricingFullPricingResponseInfrastructure) HasGpu() bool`

HasGpu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


