# O11yO11yReductionRuleCreateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | **[]string** | Labels are the label names the rule matches. Required, at least one. | 
**MatchType** | **string** | MatchType is drop or keep: drop the named labels, or keep only them. Required. | 
**MetricName** | **string** | MetricName is the metric the rule governs; one rule per metric. Required. | 

## Methods

### NewO11yO11yReductionRuleCreateIn

`func NewO11yO11yReductionRuleCreateIn(labels []string, matchType string, metricName string, ) *O11yO11yReductionRuleCreateIn`

NewO11yO11yReductionRuleCreateIn instantiates a new O11yO11yReductionRuleCreateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yReductionRuleCreateInWithDefaults

`func NewO11yO11yReductionRuleCreateInWithDefaults() *O11yO11yReductionRuleCreateIn`

NewO11yO11yReductionRuleCreateInWithDefaults instantiates a new O11yO11yReductionRuleCreateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *O11yO11yReductionRuleCreateIn) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *O11yO11yReductionRuleCreateIn) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *O11yO11yReductionRuleCreateIn) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetMatchType

`func (o *O11yO11yReductionRuleCreateIn) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *O11yO11yReductionRuleCreateIn) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *O11yO11yReductionRuleCreateIn) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.


### GetMetricName

`func (o *O11yO11yReductionRuleCreateIn) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *O11yO11yReductionRuleCreateIn) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *O11yO11yReductionRuleCreateIn) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


