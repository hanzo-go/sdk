# RunSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgScore** | Pointer to **float64** | AvgScore is the mean over the scored examples, 0 when none scored. | [optional] 
**Dataset** | Pointer to **string** | Dataset is the set that was scored. | [optional] 
**Items** | Pointer to **int64** | Items is how many examples the run attempted. | [optional] 
**JudgeModel** | Pointer to **string** | JudgeModel is the model that graded. | [optional] 
**Model** | Pointer to **string** | Model is the model under test. | [optional] 
**Results** | Pointer to [**[]ItemResult**](ItemResult.md) | Results is one row per attempted example. | [optional] 
**RunName** | Pointer to **string** | RunName is the run&#39;s label, which scores and traces are filed under. | [optional] 
**Scored** | Pointer to **int64** | Scored is how many produced a real score. It counts successes only, so a partial run is honest about what it achieved. | [optional] 

## Methods

### NewRunSummary

`func NewRunSummary() *RunSummary`

NewRunSummary instantiates a new RunSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunSummaryWithDefaults

`func NewRunSummaryWithDefaults() *RunSummary`

NewRunSummaryWithDefaults instantiates a new RunSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgScore

`func (o *RunSummary) GetAvgScore() float64`

GetAvgScore returns the AvgScore field if non-nil, zero value otherwise.

### GetAvgScoreOk

`func (o *RunSummary) GetAvgScoreOk() (*float64, bool)`

GetAvgScoreOk returns a tuple with the AvgScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgScore

`func (o *RunSummary) SetAvgScore(v float64)`

SetAvgScore sets AvgScore field to given value.

### HasAvgScore

`func (o *RunSummary) HasAvgScore() bool`

HasAvgScore returns a boolean if a field has been set.

### GetDataset

`func (o *RunSummary) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *RunSummary) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *RunSummary) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *RunSummary) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetItems

`func (o *RunSummary) GetItems() int64`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RunSummary) GetItemsOk() (*int64, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RunSummary) SetItems(v int64)`

SetItems sets Items field to given value.

### HasItems

`func (o *RunSummary) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetJudgeModel

`func (o *RunSummary) GetJudgeModel() string`

GetJudgeModel returns the JudgeModel field if non-nil, zero value otherwise.

### GetJudgeModelOk

`func (o *RunSummary) GetJudgeModelOk() (*string, bool)`

GetJudgeModelOk returns a tuple with the JudgeModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeModel

`func (o *RunSummary) SetJudgeModel(v string)`

SetJudgeModel sets JudgeModel field to given value.

### HasJudgeModel

`func (o *RunSummary) HasJudgeModel() bool`

HasJudgeModel returns a boolean if a field has been set.

### GetModel

`func (o *RunSummary) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RunSummary) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RunSummary) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *RunSummary) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetResults

`func (o *RunSummary) GetResults() []ItemResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RunSummary) GetResultsOk() (*[]ItemResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RunSummary) SetResults(v []ItemResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *RunSummary) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetRunName

`func (o *RunSummary) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *RunSummary) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *RunSummary) SetRunName(v string)`

SetRunName sets RunName field to given value.

### HasRunName

`func (o *RunSummary) HasRunName() bool`

HasRunName returns a boolean if a field has been set.

### GetScored

`func (o *RunSummary) GetScored() int64`

GetScored returns the Scored field if non-nil, zero value otherwise.

### GetScoredOk

`func (o *RunSummary) GetScoredOk() (*int64, bool)`

GetScoredOk returns a tuple with the Scored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScored

`func (o *RunSummary) SetScored(v int64)`

SetScored sets Scored field to given value.

### HasScored

`func (o *RunSummary) HasScored() bool`

HasScored returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


