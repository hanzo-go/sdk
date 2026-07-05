# PricingGpuTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Gpu** | Pointer to **string** |  | [optional] 
**Vram** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** | Price per hour (USD) | [optional] 

## Methods

### NewPricingGpuTier

`func NewPricingGpuTier() *PricingGpuTier`

NewPricingGpuTier instantiates a new PricingGpuTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingGpuTierWithDefaults

`func NewPricingGpuTierWithDefaults() *PricingGpuTier`

NewPricingGpuTierWithDefaults instantiates a new PricingGpuTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PricingGpuTier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PricingGpuTier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PricingGpuTier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PricingGpuTier) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGpu

`func (o *PricingGpuTier) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *PricingGpuTier) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *PricingGpuTier) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *PricingGpuTier) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetVram

`func (o *PricingGpuTier) GetVram() string`

GetVram returns the Vram field if non-nil, zero value otherwise.

### GetVramOk

`func (o *PricingGpuTier) GetVramOk() (*string, bool)`

GetVramOk returns a tuple with the Vram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVram

`func (o *PricingGpuTier) SetVram(v string)`

SetVram sets Vram field to given value.

### HasVram

`func (o *PricingGpuTier) HasVram() bool`

HasVram returns a boolean if a field has been set.

### GetPrice

`func (o *PricingGpuTier) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PricingGpuTier) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PricingGpuTier) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PricingGpuTier) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


