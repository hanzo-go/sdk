# Ruleset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to **int32** | Rules is how many detection rules the engine holds. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; whenever the findings store opened. | [optional] 

## Methods

### NewRuleset

`func NewRuleset() *Ruleset`

NewRuleset instantiates a new Ruleset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesetWithDefaults

`func NewRulesetWithDefaults() *Ruleset`

NewRulesetWithDefaults instantiates a new Ruleset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *Ruleset) GetRules() int32`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Ruleset) GetRulesOk() (*int32, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Ruleset) SetRules(v int32)`

SetRules sets Rules field to given value.

### HasRules

`func (o *Ruleset) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetStatus

`func (o *Ruleset) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Ruleset) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Ruleset) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Ruleset) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


