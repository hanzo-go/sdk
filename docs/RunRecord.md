# RunRecord

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

### NewRunRecord

`func NewRunRecord() *RunRecord`

NewRunRecord instantiates a new RunRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunRecordWithDefaults

`func NewRunRecordWithDefaults() *RunRecord`

NewRunRecordWithDefaults instantiates a new RunRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgScore

`func (o *RunRecord) GetAvgScore() float64`

GetAvgScore returns the AvgScore field if non-nil, zero value otherwise.

### GetAvgScoreOk

`func (o *RunRecord) GetAvgScoreOk() (*float64, bool)`

GetAvgScoreOk returns a tuple with the AvgScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgScore

`func (o *RunRecord) SetAvgScore(v float64)`

SetAvgScore sets AvgScore field to given value.

### HasAvgScore

`func (o *RunRecord) HasAvgScore() bool`

HasAvgScore returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RunRecord) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RunRecord) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RunRecord) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RunRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDataset

`func (o *RunRecord) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *RunRecord) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *RunRecord) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *RunRecord) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetItems

`func (o *RunRecord) GetItems() int64`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RunRecord) GetItemsOk() (*int64, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RunRecord) SetItems(v int64)`

SetItems sets Items field to given value.

### HasItems

`func (o *RunRecord) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetJudgeModel

`func (o *RunRecord) GetJudgeModel() string`

GetJudgeModel returns the JudgeModel field if non-nil, zero value otherwise.

### GetJudgeModelOk

`func (o *RunRecord) GetJudgeModelOk() (*string, bool)`

GetJudgeModelOk returns a tuple with the JudgeModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeModel

`func (o *RunRecord) SetJudgeModel(v string)`

SetJudgeModel sets JudgeModel field to given value.

### HasJudgeModel

`func (o *RunRecord) HasJudgeModel() bool`

HasJudgeModel returns a boolean if a field has been set.

### GetModel

`func (o *RunRecord) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RunRecord) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RunRecord) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *RunRecord) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRunName

`func (o *RunRecord) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *RunRecord) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *RunRecord) SetRunName(v string)`

SetRunName sets RunName field to given value.

### HasRunName

`func (o *RunRecord) HasRunName() bool`

HasRunName returns a boolean if a field has been set.

### GetScored

`func (o *RunRecord) GetScored() int64`

GetScored returns the Scored field if non-nil, zero value otherwise.

### GetScoredOk

`func (o *RunRecord) GetScoredOk() (*int64, bool)`

GetScoredOk returns a tuple with the Scored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScored

`func (o *RunRecord) SetScored(v int64)`

SetScored sets Scored field to given value.

### HasScored

`func (o *RunRecord) HasScored() bool`

HasScored returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RunRecord) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RunRecord) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RunRecord) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RunRecord) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


