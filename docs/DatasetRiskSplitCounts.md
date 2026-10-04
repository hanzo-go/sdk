# DatasetRiskSplitCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Judged** | Pointer to **int64** | Judged is how many rows carry a disposition. It is zero until a label plane writes one, and reporting it plainly is what lets a model plane refuse to rank rather than name a winner it cannot justify. | [optional] 
**Productive** | Pointer to **int64** | Productive is how many judged rows carry the one disposition. | [optional] 
**Rows** | Pointer to **int64** | Rows is how many rows the version holds across every split. It is the size of the version, not of the source window — the horizon, the cuts and the row cap all bind before this number. | [optional] 
**Subjects** | Pointer to **int64** | Subjects is how many distinct subjects the rows belong to. Every row of one subject is in ONE split, so this is the real sample size — the row count flatters it whenever a subject is active. | [optional] 
**Test** | Pointer to **int64** | Test is how many fall after the second cut — the LATEST slice, and the only one a score is honest about, since the split is temporal. | [optional] 
**Train** | Pointer to **int64** | Train is how many rows fall before the first cut — the EARLIEST slice of the window, which is what a model is fitted on. | [optional] 
**Unproductive** | Pointer to **int64** | Unproductive is how many carry the other. With Productive it accounts for Judged, so the class imbalance is visible before anyone trains on it; both stay 0 while Judged is 0. | [optional] 
**Val** | Pointer to **int64** | Val is how many fall between the two cuts, held out for tuning. | [optional] 

## Methods

### NewDatasetRiskSplitCounts

`func NewDatasetRiskSplitCounts() *DatasetRiskSplitCounts`

NewDatasetRiskSplitCounts instantiates a new DatasetRiskSplitCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetRiskSplitCountsWithDefaults

`func NewDatasetRiskSplitCountsWithDefaults() *DatasetRiskSplitCounts`

NewDatasetRiskSplitCountsWithDefaults instantiates a new DatasetRiskSplitCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJudged

`func (o *DatasetRiskSplitCounts) GetJudged() int64`

GetJudged returns the Judged field if non-nil, zero value otherwise.

### GetJudgedOk

`func (o *DatasetRiskSplitCounts) GetJudgedOk() (*int64, bool)`

GetJudgedOk returns a tuple with the Judged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudged

`func (o *DatasetRiskSplitCounts) SetJudged(v int64)`

SetJudged sets Judged field to given value.

### HasJudged

`func (o *DatasetRiskSplitCounts) HasJudged() bool`

HasJudged returns a boolean if a field has been set.

### GetProductive

`func (o *DatasetRiskSplitCounts) GetProductive() int64`

GetProductive returns the Productive field if non-nil, zero value otherwise.

### GetProductiveOk

`func (o *DatasetRiskSplitCounts) GetProductiveOk() (*int64, bool)`

GetProductiveOk returns a tuple with the Productive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductive

`func (o *DatasetRiskSplitCounts) SetProductive(v int64)`

SetProductive sets Productive field to given value.

### HasProductive

`func (o *DatasetRiskSplitCounts) HasProductive() bool`

HasProductive returns a boolean if a field has been set.

### GetRows

`func (o *DatasetRiskSplitCounts) GetRows() int64`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *DatasetRiskSplitCounts) GetRowsOk() (*int64, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *DatasetRiskSplitCounts) SetRows(v int64)`

SetRows sets Rows field to given value.

### HasRows

`func (o *DatasetRiskSplitCounts) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetSubjects

`func (o *DatasetRiskSplitCounts) GetSubjects() int64`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *DatasetRiskSplitCounts) GetSubjectsOk() (*int64, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *DatasetRiskSplitCounts) SetSubjects(v int64)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *DatasetRiskSplitCounts) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.

### GetTest

`func (o *DatasetRiskSplitCounts) GetTest() int64`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *DatasetRiskSplitCounts) GetTestOk() (*int64, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *DatasetRiskSplitCounts) SetTest(v int64)`

SetTest sets Test field to given value.

### HasTest

`func (o *DatasetRiskSplitCounts) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetTrain

`func (o *DatasetRiskSplitCounts) GetTrain() int64`

GetTrain returns the Train field if non-nil, zero value otherwise.

### GetTrainOk

`func (o *DatasetRiskSplitCounts) GetTrainOk() (*int64, bool)`

GetTrainOk returns a tuple with the Train field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrain

`func (o *DatasetRiskSplitCounts) SetTrain(v int64)`

SetTrain sets Train field to given value.

### HasTrain

`func (o *DatasetRiskSplitCounts) HasTrain() bool`

HasTrain returns a boolean if a field has been set.

### GetUnproductive

`func (o *DatasetRiskSplitCounts) GetUnproductive() int64`

GetUnproductive returns the Unproductive field if non-nil, zero value otherwise.

### GetUnproductiveOk

`func (o *DatasetRiskSplitCounts) GetUnproductiveOk() (*int64, bool)`

GetUnproductiveOk returns a tuple with the Unproductive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnproductive

`func (o *DatasetRiskSplitCounts) SetUnproductive(v int64)`

SetUnproductive sets Unproductive field to given value.

### HasUnproductive

`func (o *DatasetRiskSplitCounts) HasUnproductive() bool`

HasUnproductive returns a boolean if a field has been set.

### GetVal

`func (o *DatasetRiskSplitCounts) GetVal() int64`

GetVal returns the Val field if non-nil, zero value otherwise.

### GetValOk

`func (o *DatasetRiskSplitCounts) GetValOk() (*int64, bool)`

GetValOk returns a tuple with the Val field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVal

`func (o *DatasetRiskSplitCounts) SetVal(v int64)`

SetVal sets Val field to given value.

### HasVal

`func (o *DatasetRiskSplitCounts) HasVal() bool`

HasVal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


