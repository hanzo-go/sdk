# EvalEvaluatorList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]EvalEvaluatorView**](EvalEvaluatorView.md) | Data is the caller org&#39;s judges, bounded by limit. | [optional] 

## Methods

### NewEvalEvaluatorList

`func NewEvalEvaluatorList() *EvalEvaluatorList`

NewEvalEvaluatorList instantiates a new EvalEvaluatorList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalEvaluatorListWithDefaults

`func NewEvalEvaluatorListWithDefaults() *EvalEvaluatorList`

NewEvalEvaluatorListWithDefaults instantiates a new EvalEvaluatorList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EvalEvaluatorList) GetData() []EvalEvaluatorView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EvalEvaluatorList) GetDataOk() (*[]EvalEvaluatorView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EvalEvaluatorList) SetData(v []EvalEvaluatorView)`

SetData sets Data field to given value.

### HasData

`func (o *EvalEvaluatorList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


