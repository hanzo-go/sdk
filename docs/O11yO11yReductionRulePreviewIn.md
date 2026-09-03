# O11yO11yReductionRulePreviewIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | **[]string** | Labels are the label names the rule would match. Required, at least one. | 
**LookbackMs** | Pointer to **int64** | LookbackMs is how far back to sample when estimating. | [optional] 
**MatchType** | **string** | MatchType is drop or keep. Required. | 
**MetricName** | **string** | MetricName is the metric the rule would govern. Required. | 

## Methods

### NewO11yO11yReductionRulePreviewIn

`func NewO11yO11yReductionRulePreviewIn(labels []string, matchType string, metricName string, ) *O11yO11yReductionRulePreviewIn`

NewO11yO11yReductionRulePreviewIn instantiates a new O11yO11yReductionRulePreviewIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yReductionRulePreviewInWithDefaults

`func NewO11yO11yReductionRulePreviewInWithDefaults() *O11yO11yReductionRulePreviewIn`

NewO11yO11yReductionRulePreviewInWithDefaults instantiates a new O11yO11yReductionRulePreviewIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *O11yO11yReductionRulePreviewIn) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *O11yO11yReductionRulePreviewIn) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *O11yO11yReductionRulePreviewIn) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetLookbackMs

`func (o *O11yO11yReductionRulePreviewIn) GetLookbackMs() int64`

GetLookbackMs returns the LookbackMs field if non-nil, zero value otherwise.

### GetLookbackMsOk

`func (o *O11yO11yReductionRulePreviewIn) GetLookbackMsOk() (*int64, bool)`

GetLookbackMsOk returns a tuple with the LookbackMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookbackMs

`func (o *O11yO11yReductionRulePreviewIn) SetLookbackMs(v int64)`

SetLookbackMs sets LookbackMs field to given value.

### HasLookbackMs

`func (o *O11yO11yReductionRulePreviewIn) HasLookbackMs() bool`

HasLookbackMs returns a boolean if a field has been set.

### GetMatchType

`func (o *O11yO11yReductionRulePreviewIn) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *O11yO11yReductionRulePreviewIn) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *O11yO11yReductionRulePreviewIn) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.


### GetMetricName

`func (o *O11yO11yReductionRulePreviewIn) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *O11yO11yReductionRulePreviewIn) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *O11yO11yReductionRulePreviewIn) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


