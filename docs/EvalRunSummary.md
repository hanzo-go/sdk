# EvalRunSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgScore** | Pointer to **float64** | AvgScore is the mean over the scored examples, 0 when none scored. | [optional] 
**Dataset** | Pointer to **string** | Dataset is the set that was scored. | [optional] 
**Items** | Pointer to **int64** | Items is how many examples the run attempted. | [optional] 
**JudgeModel** | Pointer to **string** | JudgeModel is the model that graded. | [optional] 
**Model** | Pointer to **string** | Model is the model under test. | [optional] 
**Results** | Pointer to [**[]EvalItemResult**](EvalItemResult.md) | Results is one row per attempted example. | [optional] 
**RunName** | Pointer to **string** | RunName is the run&#39;s label, which scores and traces are filed under. | [optional] 
**Scored** | Pointer to **int64** | Scored is how many produced a real score. It counts successes only, so a partial run is honest about what it achieved. | [optional] 

## Methods

### NewEvalRunSummary

`func NewEvalRunSummary() *EvalRunSummary`

NewEvalRunSummary instantiates a new EvalRunSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalRunSummaryWithDefaults

`func NewEvalRunSummaryWithDefaults() *EvalRunSummary`

NewEvalRunSummaryWithDefaults instantiates a new EvalRunSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgScore

`func (o *EvalRunSummary) GetAvgScore() float64`

GetAvgScore returns the AvgScore field if non-nil, zero value otherwise.

### GetAvgScoreOk

`func (o *EvalRunSummary) GetAvgScoreOk() (*float64, bool)`

GetAvgScoreOk returns a tuple with the AvgScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgScore

`func (o *EvalRunSummary) SetAvgScore(v float64)`

SetAvgScore sets AvgScore field to given value.

### HasAvgScore

`func (o *EvalRunSummary) HasAvgScore() bool`

HasAvgScore returns a boolean if a field has been set.

### GetDataset

`func (o *EvalRunSummary) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *EvalRunSummary) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *EvalRunSummary) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *EvalRunSummary) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetItems

`func (o *EvalRunSummary) GetItems() int64`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EvalRunSummary) GetItemsOk() (*int64, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EvalRunSummary) SetItems(v int64)`

SetItems sets Items field to given value.

### HasItems

`func (o *EvalRunSummary) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetJudgeModel

`func (o *EvalRunSummary) GetJudgeModel() string`

GetJudgeModel returns the JudgeModel field if non-nil, zero value otherwise.

### GetJudgeModelOk

`func (o *EvalRunSummary) GetJudgeModelOk() (*string, bool)`

GetJudgeModelOk returns a tuple with the JudgeModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeModel

`func (o *EvalRunSummary) SetJudgeModel(v string)`

SetJudgeModel sets JudgeModel field to given value.

### HasJudgeModel

`func (o *EvalRunSummary) HasJudgeModel() bool`

HasJudgeModel returns a boolean if a field has been set.

### GetModel

`func (o *EvalRunSummary) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EvalRunSummary) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EvalRunSummary) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EvalRunSummary) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetResults

`func (o *EvalRunSummary) GetResults() []EvalItemResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *EvalRunSummary) GetResultsOk() (*[]EvalItemResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *EvalRunSummary) SetResults(v []EvalItemResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *EvalRunSummary) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetRunName

`func (o *EvalRunSummary) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *EvalRunSummary) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *EvalRunSummary) SetRunName(v string)`

SetRunName sets RunName field to given value.

### HasRunName

`func (o *EvalRunSummary) HasRunName() bool`

HasRunName returns a boolean if a field has been set.

### GetScored

`func (o *EvalRunSummary) GetScored() int64`

GetScored returns the Scored field if non-nil, zero value otherwise.

### GetScoredOk

`func (o *EvalRunSummary) GetScoredOk() (*int64, bool)`

GetScoredOk returns a tuple with the Scored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScored

`func (o *EvalRunSummary) SetScored(v int64)`

SetScored sets Scored field to given value.

### HasScored

`func (o *EvalRunSummary) HasScored() bool`

HasScored returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


