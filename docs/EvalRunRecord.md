# EvalRunRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgScore** | Pointer to **float64** | AvgScore is the mean over the scored examples. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when the run first landed. | [optional] 
**Dataset** | Pointer to **string** | Dataset is the set that was scored. | [optional] 
**Items** | Pointer to **int64** | Items is how many examples were attempted. | [optional] 
**JudgeModel** | Pointer to **string** | JudgeModel is the model that graded. | [optional] 
**Model** | Pointer to **string** | Model is the model under test. | [optional] 
**RunName** | Pointer to **string** | RunName is the run&#39;s label. | [optional] 
**Scored** | Pointer to **int64** | Scored is how many produced a real score. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when the record last changed. | [optional] 

## Methods

### NewEvalRunRecord

`func NewEvalRunRecord() *EvalRunRecord`

NewEvalRunRecord instantiates a new EvalRunRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalRunRecordWithDefaults

`func NewEvalRunRecordWithDefaults() *EvalRunRecord`

NewEvalRunRecordWithDefaults instantiates a new EvalRunRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgScore

`func (o *EvalRunRecord) GetAvgScore() float64`

GetAvgScore returns the AvgScore field if non-nil, zero value otherwise.

### GetAvgScoreOk

`func (o *EvalRunRecord) GetAvgScoreOk() (*float64, bool)`

GetAvgScoreOk returns a tuple with the AvgScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgScore

`func (o *EvalRunRecord) SetAvgScore(v float64)`

SetAvgScore sets AvgScore field to given value.

### HasAvgScore

`func (o *EvalRunRecord) HasAvgScore() bool`

HasAvgScore returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EvalRunRecord) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EvalRunRecord) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EvalRunRecord) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EvalRunRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDataset

`func (o *EvalRunRecord) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *EvalRunRecord) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *EvalRunRecord) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *EvalRunRecord) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetItems

`func (o *EvalRunRecord) GetItems() int64`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EvalRunRecord) GetItemsOk() (*int64, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EvalRunRecord) SetItems(v int64)`

SetItems sets Items field to given value.

### HasItems

`func (o *EvalRunRecord) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetJudgeModel

`func (o *EvalRunRecord) GetJudgeModel() string`

GetJudgeModel returns the JudgeModel field if non-nil, zero value otherwise.

### GetJudgeModelOk

`func (o *EvalRunRecord) GetJudgeModelOk() (*string, bool)`

GetJudgeModelOk returns a tuple with the JudgeModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeModel

`func (o *EvalRunRecord) SetJudgeModel(v string)`

SetJudgeModel sets JudgeModel field to given value.

### HasJudgeModel

`func (o *EvalRunRecord) HasJudgeModel() bool`

HasJudgeModel returns a boolean if a field has been set.

### GetModel

`func (o *EvalRunRecord) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EvalRunRecord) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EvalRunRecord) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EvalRunRecord) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRunName

`func (o *EvalRunRecord) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *EvalRunRecord) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *EvalRunRecord) SetRunName(v string)`

SetRunName sets RunName field to given value.

### HasRunName

`func (o *EvalRunRecord) HasRunName() bool`

HasRunName returns a boolean if a field has been set.

### GetScored

`func (o *EvalRunRecord) GetScored() int64`

GetScored returns the Scored field if non-nil, zero value otherwise.

### GetScoredOk

`func (o *EvalRunRecord) GetScoredOk() (*int64, bool)`

GetScoredOk returns a tuple with the Scored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScored

`func (o *EvalRunRecord) SetScored(v int64)`

SetScored sets Scored field to given value.

### HasScored

`func (o *EvalRunRecord) HasScored() bool`

HasScored returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EvalRunRecord) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EvalRunRecord) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EvalRunRecord) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EvalRunRecord) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


