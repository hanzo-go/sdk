# SecurityRuleset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to **int64** | Rules is how many detection rules the engine holds. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; whenever the findings store opened. | [optional] 

## Methods

### NewSecurityRuleset

`func NewSecurityRuleset() *SecurityRuleset`

NewSecurityRuleset instantiates a new SecurityRuleset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityRulesetWithDefaults

`func NewSecurityRulesetWithDefaults() *SecurityRuleset`

NewSecurityRulesetWithDefaults instantiates a new SecurityRuleset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *SecurityRuleset) GetRules() int64`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SecurityRuleset) GetRulesOk() (*int64, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SecurityRuleset) SetRules(v int64)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SecurityRuleset) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityRuleset) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityRuleset) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityRuleset) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityRuleset) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


