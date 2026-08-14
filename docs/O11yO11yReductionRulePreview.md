# O11yO11yReductionRulePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedAssets** | Pointer to [**[]O11yO11yAffectedAsset**](O11yO11yAffectedAsset.md) | AffectedAssets are the dashboards and alerts the rule would touch. | [optional] 
**CurrentRetainedSeries** | Pointer to **int32** | CurrentRetainedSeries is how many survive the rules in force today. | [optional] 
**DroppedLabels** | Pointer to **[]string** | DroppedLabels are the labels the rule would drop. | [optional] 
**EffectiveFrom** | Pointer to **time.Time** | EffectiveFrom is when the rule would take effect. | [optional] 
**IngestedSeries** | Pointer to **int32** | IngestedSeries is how many series the metric ingests today. | [optional] 
**ReductionPercent** | Pointer to **float32** | ReductionPercent is the estimated reduction, in percent. | [optional] 
**RetainedSeries** | Pointer to **int32** | RetainedSeries is how many would survive with the candidate rule. | [optional] 

## Methods

### NewO11yO11yReductionRulePreview

`func NewO11yO11yReductionRulePreview() *O11yO11yReductionRulePreview`

NewO11yO11yReductionRulePreview instantiates a new O11yO11yReductionRulePreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yReductionRulePreviewWithDefaults

`func NewO11yO11yReductionRulePreviewWithDefaults() *O11yO11yReductionRulePreview`

NewO11yO11yReductionRulePreviewWithDefaults instantiates a new O11yO11yReductionRulePreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedAssets

`func (o *O11yO11yReductionRulePreview) GetAffectedAssets() []O11yO11yAffectedAsset`

GetAffectedAssets returns the AffectedAssets field if non-nil, zero value otherwise.

### GetAffectedAssetsOk

`func (o *O11yO11yReductionRulePreview) GetAffectedAssetsOk() (*[]O11yO11yAffectedAsset, bool)`

GetAffectedAssetsOk returns a tuple with the AffectedAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedAssets

`func (o *O11yO11yReductionRulePreview) SetAffectedAssets(v []O11yO11yAffectedAsset)`

SetAffectedAssets sets AffectedAssets field to given value.

### HasAffectedAssets

`func (o *O11yO11yReductionRulePreview) HasAffectedAssets() bool`

HasAffectedAssets returns a boolean if a field has been set.

### GetCurrentRetainedSeries

`func (o *O11yO11yReductionRulePreview) GetCurrentRetainedSeries() int32`

GetCurrentRetainedSeries returns the CurrentRetainedSeries field if non-nil, zero value otherwise.

### GetCurrentRetainedSeriesOk

`func (o *O11yO11yReductionRulePreview) GetCurrentRetainedSeriesOk() (*int32, bool)`

GetCurrentRetainedSeriesOk returns a tuple with the CurrentRetainedSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRetainedSeries

`func (o *O11yO11yReductionRulePreview) SetCurrentRetainedSeries(v int32)`

SetCurrentRetainedSeries sets CurrentRetainedSeries field to given value.

### HasCurrentRetainedSeries

`func (o *O11yO11yReductionRulePreview) HasCurrentRetainedSeries() bool`

HasCurrentRetainedSeries returns a boolean if a field has been set.

### GetDroppedLabels

`func (o *O11yO11yReductionRulePreview) GetDroppedLabels() []string`

GetDroppedLabels returns the DroppedLabels field if non-nil, zero value otherwise.

### GetDroppedLabelsOk

`func (o *O11yO11yReductionRulePreview) GetDroppedLabelsOk() (*[]string, bool)`

GetDroppedLabelsOk returns a tuple with the DroppedLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDroppedLabels

`func (o *O11yO11yReductionRulePreview) SetDroppedLabels(v []string)`

SetDroppedLabels sets DroppedLabels field to given value.

### HasDroppedLabels

`func (o *O11yO11yReductionRulePreview) HasDroppedLabels() bool`

HasDroppedLabels returns a boolean if a field has been set.

### GetEffectiveFrom

`func (o *O11yO11yReductionRulePreview) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *O11yO11yReductionRulePreview) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *O11yO11yReductionRulePreview) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *O11yO11yReductionRulePreview) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### GetIngestedSeries

`func (o *O11yO11yReductionRulePreview) GetIngestedSeries() int32`

GetIngestedSeries returns the IngestedSeries field if non-nil, zero value otherwise.

### GetIngestedSeriesOk

`func (o *O11yO11yReductionRulePreview) GetIngestedSeriesOk() (*int32, bool)`

GetIngestedSeriesOk returns a tuple with the IngestedSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedSeries

`func (o *O11yO11yReductionRulePreview) SetIngestedSeries(v int32)`

SetIngestedSeries sets IngestedSeries field to given value.

### HasIngestedSeries

`func (o *O11yO11yReductionRulePreview) HasIngestedSeries() bool`

HasIngestedSeries returns a boolean if a field has been set.

### GetReductionPercent

`func (o *O11yO11yReductionRulePreview) GetReductionPercent() float32`

GetReductionPercent returns the ReductionPercent field if non-nil, zero value otherwise.

### GetReductionPercentOk

`func (o *O11yO11yReductionRulePreview) GetReductionPercentOk() (*float32, bool)`

GetReductionPercentOk returns a tuple with the ReductionPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReductionPercent

`func (o *O11yO11yReductionRulePreview) SetReductionPercent(v float32)`

SetReductionPercent sets ReductionPercent field to given value.

### HasReductionPercent

`func (o *O11yO11yReductionRulePreview) HasReductionPercent() bool`

HasReductionPercent returns a boolean if a field has been set.

### GetRetainedSeries

`func (o *O11yO11yReductionRulePreview) GetRetainedSeries() int32`

GetRetainedSeries returns the RetainedSeries field if non-nil, zero value otherwise.

### GetRetainedSeriesOk

`func (o *O11yO11yReductionRulePreview) GetRetainedSeriesOk() (*int32, bool)`

GetRetainedSeriesOk returns a tuple with the RetainedSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainedSeries

`func (o *O11yO11yReductionRulePreview) SetRetainedSeries(v int32)`

SetRetainedSeries sets RetainedSeries field to given value.

### HasRetainedSeries

`func (o *O11yO11yReductionRulePreview) HasRetainedSeries() bool`

HasRetainedSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


