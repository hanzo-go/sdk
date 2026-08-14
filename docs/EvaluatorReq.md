# EvaluatorReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | Pointer to **string** | Criteria is the written standard the judge applies; over 64 KiB is refused. | [optional] 
**Model** | Pointer to **string** | Model is the model that will do the grading. | [optional] 
**Name** | **string** | Name is the judge&#39;s org-unique handle, matching ^[A-Za-z0-9][A-Za-z0-9._-]{0,63}$. Re-posting a name edits that judge rather than adding a second one. | 
**ScoreName** | Pointer to **string** | ScoreName is the name the resulting scores are filed under. It defaults to the judge&#39;s own name and must match the same pattern. | [optional] 

## Methods

### NewEvaluatorReq

`func NewEvaluatorReq(name string, ) *EvaluatorReq`

NewEvaluatorReq instantiates a new EvaluatorReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluatorReqWithDefaults

`func NewEvaluatorReqWithDefaults() *EvaluatorReq`

NewEvaluatorReqWithDefaults instantiates a new EvaluatorReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *EvaluatorReq) GetCriteria() string`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *EvaluatorReq) GetCriteriaOk() (*string, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *EvaluatorReq) SetCriteria(v string)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *EvaluatorReq) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetModel

`func (o *EvaluatorReq) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EvaluatorReq) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EvaluatorReq) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EvaluatorReq) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *EvaluatorReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvaluatorReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvaluatorReq) SetName(v string)`

SetName sets Name field to given value.


### GetScoreName

`func (o *EvaluatorReq) GetScoreName() string`

GetScoreName returns the ScoreName field if non-nil, zero value otherwise.

### GetScoreNameOk

`func (o *EvaluatorReq) GetScoreNameOk() (*string, bool)`

GetScoreNameOk returns a tuple with the ScoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreName

`func (o *EvaluatorReq) SetScoreName(v string)`

SetScoreName sets ScoreName field to given value.

### HasScoreName

`func (o *EvaluatorReq) HasScoreName() bool`

HasScoreName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


