# O11yO11yLLMRulePricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to [**O11yO11yLLMPricingCacheCosts**](O11yO11yLLMPricingCacheCosts.md) | Cache is the cost of cached tokens, when the model prices them. | [optional] 
**Input** | Pointer to **float64** | Input is the cost per unit of input tokens. | [optional] 
**Output** | Pointer to **float64** | Output is the cost per unit of output tokens. | [optional] 

## Methods

### NewO11yO11yLLMRulePricing

`func NewO11yO11yLLMRulePricing() *O11yO11yLLMRulePricing`

NewO11yO11yLLMRulePricing instantiates a new O11yO11yLLMRulePricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMRulePricingWithDefaults

`func NewO11yO11yLLMRulePricingWithDefaults() *O11yO11yLLMRulePricing`

NewO11yO11yLLMRulePricingWithDefaults instantiates a new O11yO11yLLMRulePricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *O11yO11yLLMRulePricing) GetCache() O11yO11yLLMPricingCacheCosts`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *O11yO11yLLMRulePricing) GetCacheOk() (*O11yO11yLLMPricingCacheCosts, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *O11yO11yLLMRulePricing) SetCache(v O11yO11yLLMPricingCacheCosts)`

SetCache sets Cache field to given value.

### HasCache

`func (o *O11yO11yLLMRulePricing) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetInput

`func (o *O11yO11yLLMRulePricing) GetInput() float64`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *O11yO11yLLMRulePricing) GetInputOk() (*float64, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *O11yO11yLLMRulePricing) SetInput(v float64)`

SetInput sets Input field to given value.

### HasInput

`func (o *O11yO11yLLMRulePricing) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *O11yO11yLLMRulePricing) GetOutput() float64`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *O11yO11yLLMRulePricing) GetOutputOk() (*float64, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *O11yO11yLLMRulePricing) SetOutput(v float64)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *O11yO11yLLMRulePricing) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


