# EvalsRunSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**JudgeModel** | Pointer to **string** |  | [optional] 
**RunName** | Pointer to **string** |  | [optional] 
**Items** | Pointer to **int32** |  | [optional] 
**Scored** | Pointer to **int32** |  | [optional] 
**AvgScore** | Pointer to **float32** |  | [optional] 
**Results** | Pointer to [**[]EvalsItemResult**](EvalsItemResult.md) |  | [optional] 

## Methods

### NewEvalsRunSummary

`func NewEvalsRunSummary() *EvalsRunSummary`

NewEvalsRunSummary instantiates a new EvalsRunSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalsRunSummaryWithDefaults

`func NewEvalsRunSummaryWithDefaults() *EvalsRunSummary`

NewEvalsRunSummaryWithDefaults instantiates a new EvalsRunSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *EvalsRunSummary) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *EvalsRunSummary) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *EvalsRunSummary) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *EvalsRunSummary) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetModel

`func (o *EvalsRunSummary) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EvalsRunSummary) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EvalsRunSummary) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EvalsRunSummary) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetJudgeModel

`func (o *EvalsRunSummary) GetJudgeModel() string`

GetJudgeModel returns the JudgeModel field if non-nil, zero value otherwise.

### GetJudgeModelOk

`func (o *EvalsRunSummary) GetJudgeModelOk() (*string, bool)`

GetJudgeModelOk returns a tuple with the JudgeModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeModel

`func (o *EvalsRunSummary) SetJudgeModel(v string)`

SetJudgeModel sets JudgeModel field to given value.

### HasJudgeModel

`func (o *EvalsRunSummary) HasJudgeModel() bool`

HasJudgeModel returns a boolean if a field has been set.

### GetRunName

`func (o *EvalsRunSummary) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *EvalsRunSummary) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *EvalsRunSummary) SetRunName(v string)`

SetRunName sets RunName field to given value.

### HasRunName

`func (o *EvalsRunSummary) HasRunName() bool`

HasRunName returns a boolean if a field has been set.

### GetItems

`func (o *EvalsRunSummary) GetItems() int32`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EvalsRunSummary) GetItemsOk() (*int32, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EvalsRunSummary) SetItems(v int32)`

SetItems sets Items field to given value.

### HasItems

`func (o *EvalsRunSummary) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetScored

`func (o *EvalsRunSummary) GetScored() int32`

GetScored returns the Scored field if non-nil, zero value otherwise.

### GetScoredOk

`func (o *EvalsRunSummary) GetScoredOk() (*int32, bool)`

GetScoredOk returns a tuple with the Scored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScored

`func (o *EvalsRunSummary) SetScored(v int32)`

SetScored sets Scored field to given value.

### HasScored

`func (o *EvalsRunSummary) HasScored() bool`

HasScored returns a boolean if a field has been set.

### GetAvgScore

`func (o *EvalsRunSummary) GetAvgScore() float32`

GetAvgScore returns the AvgScore field if non-nil, zero value otherwise.

### GetAvgScoreOk

`func (o *EvalsRunSummary) GetAvgScoreOk() (*float32, bool)`

GetAvgScoreOk returns a tuple with the AvgScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgScore

`func (o *EvalsRunSummary) SetAvgScore(v float32)`

SetAvgScore sets AvgScore field to given value.

### HasAvgScore

`func (o *EvalsRunSummary) HasAvgScore() bool`

HasAvgScore returns a boolean if a field has been set.

### GetResults

`func (o *EvalsRunSummary) GetResults() []EvalsItemResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *EvalsRunSummary) GetResultsOk() (*[]EvalsItemResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *EvalsRunSummary) SetResults(v []EvalsItemResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *EvalsRunSummary) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


