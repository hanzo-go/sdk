# O11yO11yAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | Alias is the aggregation&#39;s alias. | [optional] 
**AnomalyScores** | Pointer to [**[]O11yO11yMetricSeries**](O11yO11yMetricSeries.md) | AnomalyScores are anomaly overlays. | [optional] 
**Index** | Pointer to **int64** | Index is the aggregation&#39;s position in the query. | [optional] 
**LowerBoundSeries** | Pointer to [**[]O11yO11yMetricSeries**](O11yO11yMetricSeries.md) | LowerBoundSeries are forecast lower bounds. | [optional] 
**Meta** | Pointer to [**O11yO11yAggregationMeta**](O11yO11yAggregationMeta.md) | Meta describes the aggregation. | [optional] 
**PredictedSeries** | Pointer to [**[]O11yO11yMetricSeries**](O11yO11yMetricSeries.md) | PredictedSeries are forecast overlays, when the query asked for them. | [optional] 
**Series** | Pointer to [**[]O11yO11yMetricSeries**](O11yO11yMetricSeries.md) | Series are the aggregated time series. | [optional] 
**UpperBoundSeries** | Pointer to [**[]O11yO11yMetricSeries**](O11yO11yMetricSeries.md) | UpperBoundSeries are forecast upper bounds. | [optional] 

## Methods

### NewO11yO11yAggregation

`func NewO11yO11yAggregation() *O11yO11yAggregation`

NewO11yO11yAggregation instantiates a new O11yO11yAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yAggregationWithDefaults

`func NewO11yO11yAggregationWithDefaults() *O11yO11yAggregation`

NewO11yO11yAggregationWithDefaults instantiates a new O11yO11yAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *O11yO11yAggregation) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *O11yO11yAggregation) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *O11yO11yAggregation) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *O11yO11yAggregation) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAnomalyScores

`func (o *O11yO11yAggregation) GetAnomalyScores() []O11yO11yMetricSeries`

GetAnomalyScores returns the AnomalyScores field if non-nil, zero value otherwise.

### GetAnomalyScoresOk

`func (o *O11yO11yAggregation) GetAnomalyScoresOk() (*[]O11yO11yMetricSeries, bool)`

GetAnomalyScoresOk returns a tuple with the AnomalyScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyScores

`func (o *O11yO11yAggregation) SetAnomalyScores(v []O11yO11yMetricSeries)`

SetAnomalyScores sets AnomalyScores field to given value.

### HasAnomalyScores

`func (o *O11yO11yAggregation) HasAnomalyScores() bool`

HasAnomalyScores returns a boolean if a field has been set.

### GetIndex

`func (o *O11yO11yAggregation) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *O11yO11yAggregation) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *O11yO11yAggregation) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *O11yO11yAggregation) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetLowerBoundSeries

`func (o *O11yO11yAggregation) GetLowerBoundSeries() []O11yO11yMetricSeries`

GetLowerBoundSeries returns the LowerBoundSeries field if non-nil, zero value otherwise.

### GetLowerBoundSeriesOk

`func (o *O11yO11yAggregation) GetLowerBoundSeriesOk() (*[]O11yO11yMetricSeries, bool)`

GetLowerBoundSeriesOk returns a tuple with the LowerBoundSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundSeries

`func (o *O11yO11yAggregation) SetLowerBoundSeries(v []O11yO11yMetricSeries)`

SetLowerBoundSeries sets LowerBoundSeries field to given value.

### HasLowerBoundSeries

`func (o *O11yO11yAggregation) HasLowerBoundSeries() bool`

HasLowerBoundSeries returns a boolean if a field has been set.

### GetMeta

`func (o *O11yO11yAggregation) GetMeta() O11yO11yAggregationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yO11yAggregation) GetMetaOk() (*O11yO11yAggregationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yO11yAggregation) SetMeta(v O11yO11yAggregationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yO11yAggregation) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPredictedSeries

`func (o *O11yO11yAggregation) GetPredictedSeries() []O11yO11yMetricSeries`

GetPredictedSeries returns the PredictedSeries field if non-nil, zero value otherwise.

### GetPredictedSeriesOk

`func (o *O11yO11yAggregation) GetPredictedSeriesOk() (*[]O11yO11yMetricSeries, bool)`

GetPredictedSeriesOk returns a tuple with the PredictedSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedSeries

`func (o *O11yO11yAggregation) SetPredictedSeries(v []O11yO11yMetricSeries)`

SetPredictedSeries sets PredictedSeries field to given value.

### HasPredictedSeries

`func (o *O11yO11yAggregation) HasPredictedSeries() bool`

HasPredictedSeries returns a boolean if a field has been set.

### GetSeries

`func (o *O11yO11yAggregation) GetSeries() []O11yO11yMetricSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *O11yO11yAggregation) GetSeriesOk() (*[]O11yO11yMetricSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *O11yO11yAggregation) SetSeries(v []O11yO11yMetricSeries)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *O11yO11yAggregation) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetUpperBoundSeries

`func (o *O11yO11yAggregation) GetUpperBoundSeries() []O11yO11yMetricSeries`

GetUpperBoundSeries returns the UpperBoundSeries field if non-nil, zero value otherwise.

### GetUpperBoundSeriesOk

`func (o *O11yO11yAggregation) GetUpperBoundSeriesOk() (*[]O11yO11yMetricSeries, bool)`

GetUpperBoundSeriesOk returns a tuple with the UpperBoundSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundSeries

`func (o *O11yO11yAggregation) SetUpperBoundSeries(v []O11yO11yMetricSeries)`

SetUpperBoundSeries sets UpperBoundSeries field to given value.

### HasUpperBoundSeries

`func (o *O11yO11yAggregation) HasUpperBoundSeries() bool`

HasUpperBoundSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


