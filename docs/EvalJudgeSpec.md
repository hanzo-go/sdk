# EvalJudgeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | Pointer to **string** | Criteria is the standard the judge applies, defaulting to a correctness criterion. | [optional] 
**Model** | Pointer to **string** | Model is the model that grades. It defaults to the model under test, so a run with no judge named has the model grade itself. | [optional] 
**Name** | Pointer to **string** | Name is the score name the judge&#39;s grades are filed under, \&quot;llm-judge\&quot; by default. A name that is not a legal handle falls back to that default rather than being stored as sent. | [optional] 

## Methods

### NewEvalJudgeSpec

`func NewEvalJudgeSpec() *EvalJudgeSpec`

NewEvalJudgeSpec instantiates a new EvalJudgeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalJudgeSpecWithDefaults

`func NewEvalJudgeSpecWithDefaults() *EvalJudgeSpec`

NewEvalJudgeSpecWithDefaults instantiates a new EvalJudgeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *EvalJudgeSpec) GetCriteria() string`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *EvalJudgeSpec) GetCriteriaOk() (*string, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *EvalJudgeSpec) SetCriteria(v string)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *EvalJudgeSpec) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetModel

`func (o *EvalJudgeSpec) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EvalJudgeSpec) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EvalJudgeSpec) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EvalJudgeSpec) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *EvalJudgeSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvalJudgeSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvalJudgeSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EvalJudgeSpec) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


