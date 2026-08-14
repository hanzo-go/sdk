# O11yO11yLLMPricingCacheCosts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Mode is how cached tokens are counted — subtract (inside input_tokens, OpenAI-style), additive (reported separately, Anthropic-style) or unknown. | [optional] 
**Read** | Pointer to **float32** | Read is the cost per unit of cache-read tokens. | [optional] 
**Write** | Pointer to **float32** | Write is the cost per unit of cache-write tokens. | [optional] 

## Methods

### NewO11yO11yLLMPricingCacheCosts

`func NewO11yO11yLLMPricingCacheCosts() *O11yO11yLLMPricingCacheCosts`

NewO11yO11yLLMPricingCacheCosts instantiates a new O11yO11yLLMPricingCacheCosts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMPricingCacheCostsWithDefaults

`func NewO11yO11yLLMPricingCacheCostsWithDefaults() *O11yO11yLLMPricingCacheCosts`

NewO11yO11yLLMPricingCacheCostsWithDefaults instantiates a new O11yO11yLLMPricingCacheCosts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *O11yO11yLLMPricingCacheCosts) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *O11yO11yLLMPricingCacheCosts) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *O11yO11yLLMPricingCacheCosts) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *O11yO11yLLMPricingCacheCosts) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetRead

`func (o *O11yO11yLLMPricingCacheCosts) GetRead() float32`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *O11yO11yLLMPricingCacheCosts) GetReadOk() (*float32, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *O11yO11yLLMPricingCacheCosts) SetRead(v float32)`

SetRead sets Read field to given value.

### HasRead

`func (o *O11yO11yLLMPricingCacheCosts) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetWrite

`func (o *O11yO11yLLMPricingCacheCosts) GetWrite() float32`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *O11yO11yLLMPricingCacheCosts) GetWriteOk() (*float32, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *O11yO11yLLMPricingCacheCosts) SetWrite(v float32)`

SetWrite sets Write field to given value.

### HasWrite

`func (o *O11yO11yLLMPricingCacheCosts) HasWrite() bool`

HasWrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


