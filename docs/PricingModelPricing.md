# PricingModelPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to **float32** | Cost per 1M input tokens (USD) | [optional] 
**Output** | Pointer to **float32** | Cost per 1M output tokens (USD) | [optional] 
**CacheRead** | Pointer to **float32** | Cost per 1M cached read tokens (USD) | [optional] 
**CacheWrite** | Pointer to **float32** | Cost per 1M cache write tokens (USD) | [optional] 

## Methods

### NewPricingModelPricing

`func NewPricingModelPricing() *PricingModelPricing`

NewPricingModelPricing instantiates a new PricingModelPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingModelPricingWithDefaults

`func NewPricingModelPricingWithDefaults() *PricingModelPricing`

NewPricingModelPricingWithDefaults instantiates a new PricingModelPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *PricingModelPricing) GetInput() float32`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *PricingModelPricing) GetInputOk() (*float32, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *PricingModelPricing) SetInput(v float32)`

SetInput sets Input field to given value.

### HasInput

`func (o *PricingModelPricing) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *PricingModelPricing) GetOutput() float32`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *PricingModelPricing) GetOutputOk() (*float32, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *PricingModelPricing) SetOutput(v float32)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *PricingModelPricing) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetCacheRead

`func (o *PricingModelPricing) GetCacheRead() float32`

GetCacheRead returns the CacheRead field if non-nil, zero value otherwise.

### GetCacheReadOk

`func (o *PricingModelPricing) GetCacheReadOk() (*float32, bool)`

GetCacheReadOk returns a tuple with the CacheRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheRead

`func (o *PricingModelPricing) SetCacheRead(v float32)`

SetCacheRead sets CacheRead field to given value.

### HasCacheRead

`func (o *PricingModelPricing) HasCacheRead() bool`

HasCacheRead returns a boolean if a field has been set.

### GetCacheWrite

`func (o *PricingModelPricing) GetCacheWrite() float32`

GetCacheWrite returns the CacheWrite field if non-nil, zero value otherwise.

### GetCacheWriteOk

`func (o *PricingModelPricing) GetCacheWriteOk() (*float32, bool)`

GetCacheWriteOk returns a tuple with the CacheWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWrite

`func (o *PricingModelPricing) SetCacheWrite(v float32)`

SetCacheWrite sets CacheWrite field to given value.

### HasCacheWrite

`func (o *PricingModelPricing) HasCacheWrite() bool`

HasCacheWrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


