# O11yO11yRetentionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]O11yO11yRetentionMatch**](O11yO11yRetentionMatch.md) | Conditions all have to hold for the rule to match. | [optional] 
**TtlDays** | Pointer to **int64** | TTLDays is the retention applied when it does, in days. | [optional] 

## Methods

### NewO11yO11yRetentionRule

`func NewO11yO11yRetentionRule() *O11yO11yRetentionRule`

NewO11yO11yRetentionRule instantiates a new O11yO11yRetentionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yRetentionRuleWithDefaults

`func NewO11yO11yRetentionRuleWithDefaults() *O11yO11yRetentionRule`

NewO11yO11yRetentionRuleWithDefaults instantiates a new O11yO11yRetentionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *O11yO11yRetentionRule) GetConditions() []O11yO11yRetentionMatch`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *O11yO11yRetentionRule) GetConditionsOk() (*[]O11yO11yRetentionMatch, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *O11yO11yRetentionRule) SetConditions(v []O11yO11yRetentionMatch)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *O11yO11yRetentionRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetTtlDays

`func (o *O11yO11yRetentionRule) GetTtlDays() int64`

GetTtlDays returns the TtlDays field if non-nil, zero value otherwise.

### GetTtlDaysOk

`func (o *O11yO11yRetentionRule) GetTtlDaysOk() (*int64, bool)`

GetTtlDaysOk returns a tuple with the TtlDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlDays

`func (o *O11yO11yRetentionRule) SetTtlDays(v int64)`

SetTtlDays sets TtlDays field to given value.

### HasTtlDays

`func (o *O11yO11yRetentionRule) HasTtlDays() bool`

HasTtlDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


