# O11yO11yReductionRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]O11yO11yReductionRule**](O11yO11yReductionRule.md) | Rules are the rules. | [optional] 
**Total** | Pointer to **int32** | Total is how many rules matched, across all pages. | [optional] 

## Methods

### NewO11yO11yReductionRules

`func NewO11yO11yReductionRules() *O11yO11yReductionRules`

NewO11yO11yReductionRules instantiates a new O11yO11yReductionRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yReductionRulesWithDefaults

`func NewO11yO11yReductionRulesWithDefaults() *O11yO11yReductionRules`

NewO11yO11yReductionRulesWithDefaults instantiates a new O11yO11yReductionRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *O11yO11yReductionRules) GetRules() []O11yO11yReductionRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *O11yO11yReductionRules) GetRulesOk() (*[]O11yO11yReductionRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *O11yO11yReductionRules) SetRules(v []O11yO11yReductionRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *O11yO11yReductionRules) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTotal

`func (o *O11yO11yReductionRules) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yO11yReductionRules) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yO11yReductionRules) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yO11yReductionRules) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


