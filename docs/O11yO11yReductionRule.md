# O11yO11yReductionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Active says whether the rule is in force. | [optional] 
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when the rule was created. | [optional] 
**CreatedBy** | Pointer to **string** | CreatedBy is who created it. | [optional] 
**EffectiveFrom** | Pointer to **time.Time** | EffectiveFrom is when the rule took effect. | [optional] 
**Id** | Pointer to **string** | ID is the rule&#39;s id. | [optional] 
**IngestedSamples** | Pointer to **int32** | IngestedSamples is how many samples arrived while the rule was active. | [optional] 
**IngestedSeries** | Pointer to **int32** | IngestedSeries is how many series arrived while the rule was active. | [optional] 
**Labels** | Pointer to **[]string** | Labels are the label names the rule matches. | [optional] 
**MatchType** | Pointer to **string** | MatchType is drop or keep. | [optional] 
**MetricName** | Pointer to **string** | MetricName is the metric the rule governs. | [optional] 
**RetainedSamples** | Pointer to **int32** | RetainedSamples is how many of them were kept. | [optional] 
**RetainedSeries** | Pointer to **int32** | RetainedSeries is how many of them were kept. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when the rule last changed. | [optional] 
**UpdatedBy** | Pointer to **string** | UpdatedBy is who last changed it. | [optional] 

## Methods

### NewO11yO11yReductionRule

`func NewO11yO11yReductionRule() *O11yO11yReductionRule`

NewO11yO11yReductionRule instantiates a new O11yO11yReductionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yReductionRuleWithDefaults

`func NewO11yO11yReductionRuleWithDefaults() *O11yO11yReductionRule`

NewO11yO11yReductionRuleWithDefaults instantiates a new O11yO11yReductionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *O11yO11yReductionRule) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *O11yO11yReductionRule) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *O11yO11yReductionRule) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *O11yO11yReductionRule) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *O11yO11yReductionRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yReductionRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yReductionRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yReductionRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *O11yO11yReductionRule) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *O11yO11yReductionRule) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *O11yO11yReductionRule) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *O11yO11yReductionRule) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEffectiveFrom

`func (o *O11yO11yReductionRule) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *O11yO11yReductionRule) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *O11yO11yReductionRule) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *O11yO11yReductionRule) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yReductionRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yReductionRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yReductionRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yReductionRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIngestedSamples

`func (o *O11yO11yReductionRule) GetIngestedSamples() int32`

GetIngestedSamples returns the IngestedSamples field if non-nil, zero value otherwise.

### GetIngestedSamplesOk

`func (o *O11yO11yReductionRule) GetIngestedSamplesOk() (*int32, bool)`

GetIngestedSamplesOk returns a tuple with the IngestedSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedSamples

`func (o *O11yO11yReductionRule) SetIngestedSamples(v int32)`

SetIngestedSamples sets IngestedSamples field to given value.

### HasIngestedSamples

`func (o *O11yO11yReductionRule) HasIngestedSamples() bool`

HasIngestedSamples returns a boolean if a field has been set.

### GetIngestedSeries

`func (o *O11yO11yReductionRule) GetIngestedSeries() int32`

GetIngestedSeries returns the IngestedSeries field if non-nil, zero value otherwise.

### GetIngestedSeriesOk

`func (o *O11yO11yReductionRule) GetIngestedSeriesOk() (*int32, bool)`

GetIngestedSeriesOk returns a tuple with the IngestedSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedSeries

`func (o *O11yO11yReductionRule) SetIngestedSeries(v int32)`

SetIngestedSeries sets IngestedSeries field to given value.

### HasIngestedSeries

`func (o *O11yO11yReductionRule) HasIngestedSeries() bool`

HasIngestedSeries returns a boolean if a field has been set.

### GetLabels

`func (o *O11yO11yReductionRule) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *O11yO11yReductionRule) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *O11yO11yReductionRule) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *O11yO11yReductionRule) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMatchType

`func (o *O11yO11yReductionRule) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *O11yO11yReductionRule) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *O11yO11yReductionRule) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.

### HasMatchType

`func (o *O11yO11yReductionRule) HasMatchType() bool`

HasMatchType returns a boolean if a field has been set.

### GetMetricName

`func (o *O11yO11yReductionRule) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *O11yO11yReductionRule) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *O11yO11yReductionRule) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *O11yO11yReductionRule) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetRetainedSamples

`func (o *O11yO11yReductionRule) GetRetainedSamples() int32`

GetRetainedSamples returns the RetainedSamples field if non-nil, zero value otherwise.

### GetRetainedSamplesOk

`func (o *O11yO11yReductionRule) GetRetainedSamplesOk() (*int32, bool)`

GetRetainedSamplesOk returns a tuple with the RetainedSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedSamples

`func (o *O11yO11yReductionRule) SetRetainedSamples(v int32)`

SetRetainedSamples sets RetainedSamples field to given value.

### HasRetainedSamples

`func (o *O11yO11yReductionRule) HasRetainedSamples() bool`

HasRetainedSamples returns a boolean if a field has been set.

### GetRetainedSeries

`func (o *O11yO11yReductionRule) GetRetainedSeries() int32`

GetRetainedSeries returns the RetainedSeries field if non-nil, zero value otherwise.

### GetRetainedSeriesOk

`func (o *O11yO11yReductionRule) GetRetainedSeriesOk() (*int32, bool)`

GetRetainedSeriesOk returns a tuple with the RetainedSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedSeries

`func (o *O11yO11yReductionRule) SetRetainedSeries(v int32)`

SetRetainedSeries sets RetainedSeries field to given value.

### HasRetainedSeries

`func (o *O11yO11yReductionRule) HasRetainedSeries() bool`

HasRetainedSeries returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yReductionRule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yReductionRule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yReductionRule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yReductionRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *O11yO11yReductionRule) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *O11yO11yReductionRule) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *O11yO11yReductionRule) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *O11yO11yReductionRule) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


