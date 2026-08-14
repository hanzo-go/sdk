# RunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | **string** | Dataset is the set to score, which must belong to the caller&#39;s org and hold at least one ACTIVE example. | 
**Judge** | Pointer to [**JudgeSpec**](JudgeSpec.md) | Judge is the judge to grade with. Omitted, the model under test grades itself against a default correctness criterion under the score name \&quot;llm-judge\&quot;. | [optional] 
**Limit** | Pointer to **int32** | Limit caps how many examples this run scores. It defaults to 20, and anything above 100 falls back to that default. | [optional] 
**Model** | **string** | Model is the model under test. | 
**RunName** | Pointer to **string** | RunName labels the run and is generated from the clock when omitted. It must match ^[A-Za-z0-9][A-Za-z0-9._-]{0,63}$. | [optional] 

## Methods

### NewRunRequest

`func NewRunRequest(dataset string, model string, ) *RunRequest`

NewRunRequest instantiates a new RunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunRequestWithDefaults

`func NewRunRequestWithDefaults() *RunRequest`

NewRunRequestWithDefaults instantiates a new RunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *RunRequest) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *RunRequest) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *RunRequest) SetDataset(v string)`

SetDataset sets Dataset field to given value.


### GetJudge

`func (o *RunRequest) GetJudge() JudgeSpec`

GetJudge returns the Judge field if non-nil, zero value otherwise.

### GetJudgeOk

`func (o *RunRequest) GetJudgeOk() (*JudgeSpec, bool)`

GetJudgeOk returns a tuple with the Judge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudge

`func (o *RunRequest) SetJudge(v JudgeSpec)`

SetJudge sets Judge field to given value.

### HasJudge

`func (o *RunRequest) HasJudge() bool`

HasJudge returns a boolean if a field has been set.

### GetLimit

`func (o *RunRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RunRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RunRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RunRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetModel

`func (o *RunRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RunRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RunRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetRunName

`func (o *RunRequest) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *RunRequest) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *RunRequest) SetRunName(v string)`

SetRunName sets RunName field to given value.

### HasRunName

`func (o *RunRequest) HasRunName() bool`

HasRunName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


