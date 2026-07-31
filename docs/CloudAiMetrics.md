# CloudAiMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** |  | [optional] 
**EvalRuns** | Pointer to [**[]CloudAimRunStat**](CloudAimRunStat.md) | recent eval runs (progress) | [optional] 
**Evals** | Pointer to [**CloudAimEvals**](CloudAimEvals.md) |  | [optional] 
**O11yAi** | Pointer to [**CloudAimO11yAI**](CloudAimO11yAI.md) |  | [optional] 
**O11yAiModels** | Pointer to [**[]CloudAimLfModelStat**](CloudAimLfModelStat.md) | o11y_ai per-model (honest-empty today) | [optional] 
**Range** | Pointer to **string** |  | [optional] 
**ScoreNames** | Pointer to [**[]CloudAimScoreStat**](CloudAimScoreStat.md) | eval_scores per score-name | [optional] 
**ScoreSeries** | Pointer to [**[]CloudAimScorePoint**](CloudAimScorePoint.md) | avg eval score over time (progress trend) | [optional] 
**Start** | Pointer to **string** |  | [optional] 
**TopModels** | Pointer to [**[]CloudAimModelStat**](CloudAimModelStat.md) | cloud_usage per-model (populated today) | [optional] 
**Usage** | Pointer to [**CloudAimUsage**](CloudAimUsage.md) |  | [optional] 

## Methods

### NewCloudAiMetrics

`func NewCloudAiMetrics() *CloudAiMetrics`

NewCloudAiMetrics instantiates a new CloudAiMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAiMetricsWithDefaults

`func NewCloudAiMetricsWithDefaults() *CloudAiMetrics`

NewCloudAiMetricsWithDefaults instantiates a new CloudAiMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *CloudAiMetrics) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CloudAiMetrics) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CloudAiMetrics) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CloudAiMetrics) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetEvalRuns

`func (o *CloudAiMetrics) GetEvalRuns() []CloudAimRunStat`

GetEvalRuns returns the EvalRuns field if non-nil, zero value otherwise.

### GetEvalRunsOk

`func (o *CloudAiMetrics) GetEvalRunsOk() (*[]CloudAimRunStat, bool)`

GetEvalRunsOk returns a tuple with the EvalRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvalRuns

`func (o *CloudAiMetrics) SetEvalRuns(v []CloudAimRunStat)`

SetEvalRuns sets EvalRuns field to given value.

### HasEvalRuns

`func (o *CloudAiMetrics) HasEvalRuns() bool`

HasEvalRuns returns a boolean if a field has been set.

### GetEvals

`func (o *CloudAiMetrics) GetEvals() CloudAimEvals`

GetEvals returns the Evals field if non-nil, zero value otherwise.

### GetEvalsOk

`func (o *CloudAiMetrics) GetEvalsOk() (*CloudAimEvals, bool)`

GetEvalsOk returns a tuple with the Evals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvals

`func (o *CloudAiMetrics) SetEvals(v CloudAimEvals)`

SetEvals sets Evals field to given value.

### HasEvals

`func (o *CloudAiMetrics) HasEvals() bool`

HasEvals returns a boolean if a field has been set.

### GetO11yAi

`func (o *CloudAiMetrics) GetO11yAi() CloudAimO11yAI`

GetO11yAi returns the O11yAi field if non-nil, zero value otherwise.

### GetO11yAiOk

`func (o *CloudAiMetrics) GetO11yAiOk() (*CloudAimO11yAI, bool)`

GetO11yAiOk returns a tuple with the O11yAi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO11yAi

`func (o *CloudAiMetrics) SetO11yAi(v CloudAimO11yAI)`

SetO11yAi sets O11yAi field to given value.

### HasO11yAi

`func (o *CloudAiMetrics) HasO11yAi() bool`

HasO11yAi returns a boolean if a field has been set.

### GetO11yAiModels

`func (o *CloudAiMetrics) GetO11yAiModels() []CloudAimLfModelStat`

GetO11yAiModels returns the O11yAiModels field if non-nil, zero value otherwise.

### GetO11yAiModelsOk

`func (o *CloudAiMetrics) GetO11yAiModelsOk() (*[]CloudAimLfModelStat, bool)`

GetO11yAiModelsOk returns a tuple with the O11yAiModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO11yAiModels

`func (o *CloudAiMetrics) SetO11yAiModels(v []CloudAimLfModelStat)`

SetO11yAiModels sets O11yAiModels field to given value.

### HasO11yAiModels

`func (o *CloudAiMetrics) HasO11yAiModels() bool`

HasO11yAiModels returns a boolean if a field has been set.

### GetRange

`func (o *CloudAiMetrics) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CloudAiMetrics) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CloudAiMetrics) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *CloudAiMetrics) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScoreNames

`func (o *CloudAiMetrics) GetScoreNames() []CloudAimScoreStat`

GetScoreNames returns the ScoreNames field if non-nil, zero value otherwise.

### GetScoreNamesOk

`func (o *CloudAiMetrics) GetScoreNamesOk() (*[]CloudAimScoreStat, bool)`

GetScoreNamesOk returns a tuple with the ScoreNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreNames

`func (o *CloudAiMetrics) SetScoreNames(v []CloudAimScoreStat)`

SetScoreNames sets ScoreNames field to given value.

### HasScoreNames

`func (o *CloudAiMetrics) HasScoreNames() bool`

HasScoreNames returns a boolean if a field has been set.

### GetScoreSeries

`func (o *CloudAiMetrics) GetScoreSeries() []CloudAimScorePoint`

GetScoreSeries returns the ScoreSeries field if non-nil, zero value otherwise.

### GetScoreSeriesOk

`func (o *CloudAiMetrics) GetScoreSeriesOk() (*[]CloudAimScorePoint, bool)`

GetScoreSeriesOk returns a tuple with the ScoreSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreSeries

`func (o *CloudAiMetrics) SetScoreSeries(v []CloudAimScorePoint)`

SetScoreSeries sets ScoreSeries field to given value.

### HasScoreSeries

`func (o *CloudAiMetrics) HasScoreSeries() bool`

HasScoreSeries returns a boolean if a field has been set.

### GetStart

`func (o *CloudAiMetrics) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CloudAiMetrics) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CloudAiMetrics) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *CloudAiMetrics) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTopModels

`func (o *CloudAiMetrics) GetTopModels() []CloudAimModelStat`

GetTopModels returns the TopModels field if non-nil, zero value otherwise.

### GetTopModelsOk

`func (o *CloudAiMetrics) GetTopModelsOk() (*[]CloudAimModelStat, bool)`

GetTopModelsOk returns a tuple with the TopModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopModels

`func (o *CloudAiMetrics) SetTopModels(v []CloudAimModelStat)`

SetTopModels sets TopModels field to given value.

### HasTopModels

`func (o *CloudAiMetrics) HasTopModels() bool`

HasTopModels returns a boolean if a field has been set.

### GetUsage

`func (o *CloudAiMetrics) GetUsage() CloudAimUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *CloudAiMetrics) GetUsageOk() (*CloudAimUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *CloudAiMetrics) SetUsage(v CloudAimUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *CloudAiMetrics) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


