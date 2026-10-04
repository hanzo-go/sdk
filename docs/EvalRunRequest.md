# EvalRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | **string** | Dataset is the set to score, which must belong to the caller&#39;s org and hold at least one ACTIVE example. | 
**Judge** | Pointer to [**EvalJudgeSpec**](EvalJudgeSpec.md) | Judge is the judge to grade with. Omitted, the model under test grades itself against a default correctness criterion under the score name \&quot;llm-judge\&quot;. | [optional] 
**Limit** | Pointer to **int64** | Limit caps how many examples this run scores. It defaults to 20, and anything above 100 falls back to that default. | [optional] 
**Model** | **string** | Model is the model under test. | 
**RunName** | Pointer to **string** | RunName labels the run and is generated from the clock when omitted. It must match ^[A-Za-z0-9][A-Za-z0-9._-]{0,63}$. | [optional] 

## Methods

### NewEvalRunRequest

`func NewEvalRunRequest(dataset string, model string, ) *EvalRunRequest`

NewEvalRunRequest instantiates a new EvalRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalRunRequestWithDefaults

`func NewEvalRunRequestWithDefaults() *EvalRunRequest`

NewEvalRunRequestWithDefaults instantiates a new EvalRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *EvalRunRequest) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *EvalRunRequest) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *EvalRunRequest) SetDataset(v string)`

SetDataset sets Dataset field to given value.


### GetJudge

`func (o *EvalRunRequest) GetJudge() EvalJudgeSpec`

GetJudge returns the Judge field if non-nil, zero value otherwise.

### GetJudgeOk

`func (o *EvalRunRequest) GetJudgeOk() (*EvalJudgeSpec, bool)`

GetJudgeOk returns a tuple with the Judge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudge

`func (o *EvalRunRequest) SetJudge(v EvalJudgeSpec)`

SetJudge sets Judge field to given value.

### HasJudge

`func (o *EvalRunRequest) HasJudge() bool`

HasJudge returns a boolean if a field has been set.

### GetLimit

`func (o *EvalRunRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *EvalRunRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *EvalRunRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *EvalRunRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetModel

`func (o *EvalRunRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EvalRunRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EvalRunRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetRunName

`func (o *EvalRunRequest) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *EvalRunRequest) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *EvalRunRequest) SetRunName(v string)`

SetRunName sets RunName field to given value.

### HasRunName

`func (o *EvalRunRequest) HasRunName() bool`

HasRunName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


