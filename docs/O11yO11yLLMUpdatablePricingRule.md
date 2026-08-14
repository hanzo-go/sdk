# O11yO11yLLMUpdatablePricingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enabled turns the rule on. | [optional] 
**Id** | Pointer to **string** | ID matches an existing rule by its id. | [optional] 
**IsOverride** | Pointer to **bool** | IsOverride pins the rule so the sync job skips it. Omit to leave a matched override untouched. | [optional] 
**ModelName** | Pointer to **string** | Model is the model the rule prices. Required. | [optional] 
**ModelPattern** | Pointer to **[]string** | ModelPattern are the model-name globs the rule matches. Required. | [optional] 
**Pricing** | Pointer to [**O11yO11yLLMRulePricing**](O11yO11yLLMRulePricing.md) | Pricing is the per-unit cost. Required. | [optional] 
**Provider** | Pointer to **string** | Provider is the model&#39;s provider. Required. | [optional] 
**SourceId** | Pointer to **string** | SourceID matches an existing rule by its upstream source id. | [optional] 
**Unit** | Pointer to **string** | Unit is the pricing unit, e.g. per_million_tokens. Required. | [optional] 

## Methods

### NewO11yO11yLLMUpdatablePricingRule

`func NewO11yO11yLLMUpdatablePricingRule() *O11yO11yLLMUpdatablePricingRule`

NewO11yO11yLLMUpdatablePricingRule instantiates a new O11yO11yLLMUpdatablePricingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMUpdatablePricingRuleWithDefaults

`func NewO11yO11yLLMUpdatablePricingRuleWithDefaults() *O11yO11yLLMUpdatablePricingRule`

NewO11yO11yLLMUpdatablePricingRuleWithDefaults instantiates a new O11yO11yLLMUpdatablePricingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *O11yO11yLLMUpdatablePricingRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *O11yO11yLLMUpdatablePricingRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *O11yO11yLLMUpdatablePricingRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *O11yO11yLLMUpdatablePricingRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yLLMUpdatablePricingRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yLLMUpdatablePricingRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yLLMUpdatablePricingRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yLLMUpdatablePricingRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsOverride

`func (o *O11yO11yLLMUpdatablePricingRule) GetIsOverride() bool`

GetIsOverride returns the IsOverride field if non-nil, zero value otherwise.

### GetIsOverrideOk

`func (o *O11yO11yLLMUpdatablePricingRule) GetIsOverrideOk() (*bool, bool)`

GetIsOverrideOk returns a tuple with the IsOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverride

`func (o *O11yO11yLLMUpdatablePricingRule) SetIsOverride(v bool)`

SetIsOverride sets IsOverride field to given value.

### HasIsOverride

`func (o *O11yO11yLLMUpdatablePricingRule) HasIsOverride() bool`

HasIsOverride returns a boolean if a field has been set.

### GetModelName

`func (o *O11yO11yLLMUpdatablePricingRule) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *O11yO11yLLMUpdatablePricingRule) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *O11yO11yLLMUpdatablePricingRule) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *O11yO11yLLMUpdatablePricingRule) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetModelPattern

`func (o *O11yO11yLLMUpdatablePricingRule) GetModelPattern() []string`

GetModelPattern returns the ModelPattern field if non-nil, zero value otherwise.

### GetModelPatternOk

`func (o *O11yO11yLLMUpdatablePricingRule) GetModelPatternOk() (*[]string, bool)`

GetModelPatternOk returns a tuple with the ModelPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelPattern

`func (o *O11yO11yLLMUpdatablePricingRule) SetModelPattern(v []string)`

SetModelPattern sets ModelPattern field to given value.

### HasModelPattern

`func (o *O11yO11yLLMUpdatablePricingRule) HasModelPattern() bool`

HasModelPattern returns a boolean if a field has been set.

### GetPricing

`func (o *O11yO11yLLMUpdatablePricingRule) GetPricing() O11yO11yLLMRulePricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *O11yO11yLLMUpdatablePricingRule) GetPricingOk() (*O11yO11yLLMRulePricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *O11yO11yLLMUpdatablePricingRule) SetPricing(v O11yO11yLLMRulePricing)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *O11yO11yLLMUpdatablePricingRule) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetProvider

`func (o *O11yO11yLLMUpdatablePricingRule) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *O11yO11yLLMUpdatablePricingRule) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *O11yO11yLLMUpdatablePricingRule) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *O11yO11yLLMUpdatablePricingRule) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSourceId

`func (o *O11yO11yLLMUpdatablePricingRule) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *O11yO11yLLMUpdatablePricingRule) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *O11yO11yLLMUpdatablePricingRule) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *O11yO11yLLMUpdatablePricingRule) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetUnit

`func (o *O11yO11yLLMUpdatablePricingRule) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *O11yO11yLLMUpdatablePricingRule) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *O11yO11yLLMUpdatablePricingRule) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *O11yO11yLLMUpdatablePricingRule) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


