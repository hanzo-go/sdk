# AiMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** |  | [optional] 
**EvalRuns** | Pointer to [**[]AimRunStat**](AimRunStat.md) | recent eval runs (progress) | [optional] 
**Evals** | Pointer to [**AimEvals**](AimEvals.md) |  | [optional] 
**O11yAi** | Pointer to [**AimO11yAI**](AimO11yAI.md) |  | [optional] 
**O11yAiModels** | Pointer to [**[]AimLfModelStat**](AimLfModelStat.md) | gen_ai spans per-model | [optional] 
**Range** | Pointer to **string** |  | [optional] 
**ScoreNames** | Pointer to [**[]AimScoreStat**](AimScoreStat.md) | eval_scores per score-name | [optional] 
**ScoreSeries** | Pointer to [**[]AimScorePoint**](AimScorePoint.md) | avg eval score over time (progress trend) | [optional] 
**Start** | Pointer to **string** |  | [optional] 
**TopActors** | Pointer to [**[]AimActorStat**](AimActorStat.md) | TopActors is per-PRINCIPAL spend from the same ledger — whose bill it is, which the per-model board cannot answer. | [optional] 
**TopModels** | Pointer to [**[]AimModelStat**](AimModelStat.md) | cloud_usage per-model (populated today) | [optional] 
**Usage** | Pointer to [**AimUsage**](AimUsage.md) |  | [optional] 

## Methods

### NewAiMetrics

`func NewAiMetrics() *AiMetrics`

NewAiMetrics instantiates a new AiMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiMetricsWithDefaults

`func NewAiMetricsWithDefaults() *AiMetrics`

NewAiMetricsWithDefaults instantiates a new AiMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *AiMetrics) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AiMetrics) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AiMetrics) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *AiMetrics) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetEvalRuns

`func (o *AiMetrics) GetEvalRuns() []AimRunStat`

GetEvalRuns returns the EvalRuns field if non-nil, zero value otherwise.

### GetEvalRunsOk

`func (o *AiMetrics) GetEvalRunsOk() (*[]AimRunStat, bool)`

GetEvalRunsOk returns a tuple with the EvalRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvalRuns

`func (o *AiMetrics) SetEvalRuns(v []AimRunStat)`

SetEvalRuns sets EvalRuns field to given value.

### HasEvalRuns

`func (o *AiMetrics) HasEvalRuns() bool`

HasEvalRuns returns a boolean if a field has been set.

### GetEvals

`func (o *AiMetrics) GetEvals() AimEvals`

GetEvals returns the Evals field if non-nil, zero value otherwise.

### GetEvalsOk

`func (o *AiMetrics) GetEvalsOk() (*AimEvals, bool)`

GetEvalsOk returns a tuple with the Evals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvals

`func (o *AiMetrics) SetEvals(v AimEvals)`

SetEvals sets Evals field to given value.

### HasEvals

`func (o *AiMetrics) HasEvals() bool`

HasEvals returns a boolean if a field has been set.

### GetO11yAi

`func (o *AiMetrics) GetO11yAi() AimO11yAI`

GetO11yAi returns the O11yAi field if non-nil, zero value otherwise.

### GetO11yAiOk

`func (o *AiMetrics) GetO11yAiOk() (*AimO11yAI, bool)`

GetO11yAiOk returns a tuple with the O11yAi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO11yAi

`func (o *AiMetrics) SetO11yAi(v AimO11yAI)`

SetO11yAi sets O11yAi field to given value.

### HasO11yAi

`func (o *AiMetrics) HasO11yAi() bool`

HasO11yAi returns a boolean if a field has been set.

### GetO11yAiModels

`func (o *AiMetrics) GetO11yAiModels() []AimLfModelStat`

GetO11yAiModels returns the O11yAiModels field if non-nil, zero value otherwise.

### GetO11yAiModelsOk

`func (o *AiMetrics) GetO11yAiModelsOk() (*[]AimLfModelStat, bool)`

GetO11yAiModelsOk returns a tuple with the O11yAiModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO11yAiModels

`func (o *AiMetrics) SetO11yAiModels(v []AimLfModelStat)`

SetO11yAiModels sets O11yAiModels field to given value.

### HasO11yAiModels

`func (o *AiMetrics) HasO11yAiModels() bool`

HasO11yAiModels returns a boolean if a field has been set.

### GetRange

`func (o *AiMetrics) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *AiMetrics) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *AiMetrics) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *AiMetrics) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScoreNames

`func (o *AiMetrics) GetScoreNames() []AimScoreStat`

GetScoreNames returns the ScoreNames field if non-nil, zero value otherwise.

### GetScoreNamesOk

`func (o *AiMetrics) GetScoreNamesOk() (*[]AimScoreStat, bool)`

GetScoreNamesOk returns a tuple with the ScoreNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreNames

`func (o *AiMetrics) SetScoreNames(v []AimScoreStat)`

SetScoreNames sets ScoreNames field to given value.

### HasScoreNames

`func (o *AiMetrics) HasScoreNames() bool`

HasScoreNames returns a boolean if a field has been set.

### GetScoreSeries

`func (o *AiMetrics) GetScoreSeries() []AimScorePoint`

GetScoreSeries returns the ScoreSeries field if non-nil, zero value otherwise.

### GetScoreSeriesOk

`func (o *AiMetrics) GetScoreSeriesOk() (*[]AimScorePoint, bool)`

GetScoreSeriesOk returns a tuple with the ScoreSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreSeries

`func (o *AiMetrics) SetScoreSeries(v []AimScorePoint)`

SetScoreSeries sets ScoreSeries field to given value.

### HasScoreSeries

`func (o *AiMetrics) HasScoreSeries() bool`

HasScoreSeries returns a boolean if a field has been set.

### GetStart

`func (o *AiMetrics) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AiMetrics) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AiMetrics) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *AiMetrics) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTopActors

`func (o *AiMetrics) GetTopActors() []AimActorStat`

GetTopActors returns the TopActors field if non-nil, zero value otherwise.

### GetTopActorsOk

`func (o *AiMetrics) GetTopActorsOk() (*[]AimActorStat, bool)`

GetTopActorsOk returns a tuple with the TopActors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopActors

`func (o *AiMetrics) SetTopActors(v []AimActorStat)`

SetTopActors sets TopActors field to given value.

### HasTopActors

`func (o *AiMetrics) HasTopActors() bool`

HasTopActors returns a boolean if a field has been set.

### GetTopModels

`func (o *AiMetrics) GetTopModels() []AimModelStat`

GetTopModels returns the TopModels field if non-nil, zero value otherwise.

### GetTopModelsOk

`func (o *AiMetrics) GetTopModelsOk() (*[]AimModelStat, bool)`

GetTopModelsOk returns a tuple with the TopModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopModels

`func (o *AiMetrics) SetTopModels(v []AimModelStat)`

SetTopModels sets TopModels field to given value.

### HasTopModels

`func (o *AiMetrics) HasTopModels() bool`

HasTopModels returns a boolean if a field has been set.

### GetUsage

`func (o *AiMetrics) GetUsage() AimUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AiMetrics) GetUsageOk() (*AimUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AiMetrics) SetUsage(v AimUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AiMetrics) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


