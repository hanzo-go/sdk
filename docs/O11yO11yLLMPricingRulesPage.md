# O11yO11yLLMPricingRulesPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]O11yO11yLLMPricingRule**](O11yO11yLLMPricingRule.md) | Items are the rules. | [optional] 
**Limit** | Pointer to **int32** | Limit is the page cap the read ran with. | [optional] 
**Offset** | Pointer to **int32** | Offset is the row offset this page started at. | [optional] 
**Total** | Pointer to **int32** | Total is how many rules match, across all pages. | [optional] 

## Methods

### NewO11yO11yLLMPricingRulesPage

`func NewO11yO11yLLMPricingRulesPage() *O11yO11yLLMPricingRulesPage`

NewO11yO11yLLMPricingRulesPage instantiates a new O11yO11yLLMPricingRulesPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMPricingRulesPageWithDefaults

`func NewO11yO11yLLMPricingRulesPageWithDefaults() *O11yO11yLLMPricingRulesPage`

NewO11yO11yLLMPricingRulesPageWithDefaults instantiates a new O11yO11yLLMPricingRulesPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *O11yO11yLLMPricingRulesPage) GetItems() []O11yO11yLLMPricingRule`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *O11yO11yLLMPricingRulesPage) GetItemsOk() (*[]O11yO11yLLMPricingRule, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *O11yO11yLLMPricingRulesPage) SetItems(v []O11yO11yLLMPricingRule)`

SetItems sets Items field to given value.

### HasItems

`func (o *O11yO11yLLMPricingRulesPage) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLimit

`func (o *O11yO11yLLMPricingRulesPage) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yO11yLLMPricingRulesPage) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yO11yLLMPricingRulesPage) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yO11yLLMPricingRulesPage) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *O11yO11yLLMPricingRulesPage) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *O11yO11yLLMPricingRulesPage) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *O11yO11yLLMPricingRulesPage) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *O11yO11yLLMPricingRulesPage) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTotal

`func (o *O11yO11yLLMPricingRulesPage) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yO11yLLMPricingRulesPage) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yO11yLLMPricingRulesPage) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yO11yLLMPricingRulesPage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


