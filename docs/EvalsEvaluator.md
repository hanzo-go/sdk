# EvalsEvaluator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Metric** | Pointer to [**EvalsMetric**](EvalsMetric.md) |  | [optional] 
**Rubric** | Pointer to **string** | Custom rubric (when not a pre-built metric) | [optional] 
**JudgeModel** | Pointer to **string** |  | [optional] 

## Methods

### NewEvalsEvaluator

`func NewEvalsEvaluator(name string, ) *EvalsEvaluator`

NewEvalsEvaluator instantiates a new EvalsEvaluator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalsEvaluatorWithDefaults

`func NewEvalsEvaluatorWithDefaults() *EvalsEvaluator`

NewEvalsEvaluatorWithDefaults instantiates a new EvalsEvaluator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EvalsEvaluator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvalsEvaluator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvalsEvaluator) SetName(v string)`

SetName sets Name field to given value.


### GetMetric

`func (o *EvalsEvaluator) GetMetric() EvalsMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *EvalsEvaluator) GetMetricOk() (*EvalsMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *EvalsEvaluator) SetMetric(v EvalsMetric)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *EvalsEvaluator) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetRubric

`func (o *EvalsEvaluator) GetRubric() string`

GetRubric returns the Rubric field if non-nil, zero value otherwise.

### GetRubricOk

`func (o *EvalsEvaluator) GetRubricOk() (*string, bool)`

GetRubricOk returns a tuple with the Rubric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRubric

`func (o *EvalsEvaluator) SetRubric(v string)`

SetRubric sets Rubric field to given value.

### HasRubric

`func (o *EvalsEvaluator) HasRubric() bool`

HasRubric returns a boolean if a field has been set.

### GetJudgeModel

`func (o *EvalsEvaluator) GetJudgeModel() string`

GetJudgeModel returns the JudgeModel field if non-nil, zero value otherwise.

### GetJudgeModelOk

`func (o *EvalsEvaluator) GetJudgeModelOk() (*string, bool)`

GetJudgeModelOk returns a tuple with the JudgeModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeModel

`func (o *EvalsEvaluator) SetJudgeModel(v string)`

SetJudgeModel sets JudgeModel field to given value.

### HasJudgeModel

`func (o *EvalsEvaluator) HasJudgeModel() bool`

HasJudgeModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


