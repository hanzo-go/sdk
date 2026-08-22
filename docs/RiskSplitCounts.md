# RiskSplitCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Judged** | Pointer to **int32** | Judged is how many rows carry a disposition. It is zero until a label plane writes one, and reporting it plainly is what lets a model plane refuse to rank rather than name a winner it cannot justify. | [optional] 
**Productive** | Pointer to **int32** | Productive is how many judged rows carry the one disposition. | [optional] 
**Rows** | Pointer to **int32** | Rows is how many rows the version holds across every split. It is the size of the version, not of the source window — the horizon, the cuts and the row cap all bind before this number. | [optional] 
**Subjects** | Pointer to **int32** | Subjects is how many distinct subjects the rows belong to. Every row of one subject is in ONE split, so this is the real sample size — the row count flatters it whenever a subject is active. | [optional] 
**Test** | Pointer to **int32** | Test is how many fall after the second cut — the LATEST slice, and the only one a score is honest about, since the split is temporal. | [optional] 
**Train** | Pointer to **int32** | Train is how many rows fall before the first cut — the EARLIEST slice of the window, which is what a model is fitted on. | [optional] 
**Unproductive** | Pointer to **int32** | Unproductive is how many carry the other. With Productive it accounts for Judged, so the class imbalance is visible before anyone trains on it; both stay 0 while Judged is 0. | [optional] 
**Val** | Pointer to **int32** | Val is how many fall between the two cuts, held out for tuning. | [optional] 

## Methods

### NewRiskSplitCounts

`func NewRiskSplitCounts() *RiskSplitCounts`

NewRiskSplitCounts instantiates a new RiskSplitCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskSplitCountsWithDefaults

`func NewRiskSplitCountsWithDefaults() *RiskSplitCounts`

NewRiskSplitCountsWithDefaults instantiates a new RiskSplitCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJudged

`func (o *RiskSplitCounts) GetJudged() int32`

GetJudged returns the Judged field if non-nil, zero value otherwise.

### GetJudgedOk

`func (o *RiskSplitCounts) GetJudgedOk() (*int32, bool)`

GetJudgedOk returns a tuple with the Judged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudged

`func (o *RiskSplitCounts) SetJudged(v int32)`

SetJudged sets Judged field to given value.

### HasJudged

`func (o *RiskSplitCounts) HasJudged() bool`

HasJudged returns a boolean if a field has been set.

### GetProductive

`func (o *RiskSplitCounts) GetProductive() int32`

GetProductive returns the Productive field if non-nil, zero value otherwise.

### GetProductiveOk

`func (o *RiskSplitCounts) GetProductiveOk() (*int32, bool)`

GetProductiveOk returns a tuple with the Productive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductive

`func (o *RiskSplitCounts) SetProductive(v int32)`

SetProductive sets Productive field to given value.

### HasProductive

`func (o *RiskSplitCounts) HasProductive() bool`

HasProductive returns a boolean if a field has been set.

### GetRows

`func (o *RiskSplitCounts) GetRows() int32`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *RiskSplitCounts) GetRowsOk() (*int32, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *RiskSplitCounts) SetRows(v int32)`

SetRows sets Rows field to given value.

### HasRows

`func (o *RiskSplitCounts) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetSubjects

`func (o *RiskSplitCounts) GetSubjects() int32`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *RiskSplitCounts) GetSubjectsOk() (*int32, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *RiskSplitCounts) SetSubjects(v int32)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *RiskSplitCounts) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.

### GetTest

`func (o *RiskSplitCounts) GetTest() int32`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *RiskSplitCounts) GetTestOk() (*int32, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *RiskSplitCounts) SetTest(v int32)`

SetTest sets Test field to given value.

### HasTest

`func (o *RiskSplitCounts) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetTrain

`func (o *RiskSplitCounts) GetTrain() int32`

GetTrain returns the Train field if non-nil, zero value otherwise.

### GetTrainOk

`func (o *RiskSplitCounts) GetTrainOk() (*int32, bool)`

GetTrainOk returns a tuple with the Train field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrain

`func (o *RiskSplitCounts) SetTrain(v int32)`

SetTrain sets Train field to given value.

### HasTrain

`func (o *RiskSplitCounts) HasTrain() bool`

HasTrain returns a boolean if a field has been set.

### GetUnproductive

`func (o *RiskSplitCounts) GetUnproductive() int32`

GetUnproductive returns the Unproductive field if non-nil, zero value otherwise.

### GetUnproductiveOk

`func (o *RiskSplitCounts) GetUnproductiveOk() (*int32, bool)`

GetUnproductiveOk returns a tuple with the Unproductive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnproductive

`func (o *RiskSplitCounts) SetUnproductive(v int32)`

SetUnproductive sets Unproductive field to given value.

### HasUnproductive

`func (o *RiskSplitCounts) HasUnproductive() bool`

HasUnproductive returns a boolean if a field has been set.

### GetVal

`func (o *RiskSplitCounts) GetVal() int32`

GetVal returns the Val field if non-nil, zero value otherwise.

### GetValOk

`func (o *RiskSplitCounts) GetValOk() (*int32, bool)`

GetValOk returns a tuple with the Val field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVal

`func (o *RiskSplitCounts) SetVal(v int32)`

SetVal sets Val field to given value.

### HasVal

`func (o *RiskSplitCounts) HasVal() bool`

HasVal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


