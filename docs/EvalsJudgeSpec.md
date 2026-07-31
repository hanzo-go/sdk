# EvalsJudgeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | Judge model (defaults to the model-under-test) | [optional] 
**Criteria** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] [default to "llm-judge"]

## Methods

### NewEvalsJudgeSpec

`func NewEvalsJudgeSpec() *EvalsJudgeSpec`

NewEvalsJudgeSpec instantiates a new EvalsJudgeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalsJudgeSpecWithDefaults

`func NewEvalsJudgeSpecWithDefaults() *EvalsJudgeSpec`

NewEvalsJudgeSpecWithDefaults instantiates a new EvalsJudgeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *EvalsJudgeSpec) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EvalsJudgeSpec) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EvalsJudgeSpec) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EvalsJudgeSpec) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetCriteria

`func (o *EvalsJudgeSpec) GetCriteria() string`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *EvalsJudgeSpec) GetCriteriaOk() (*string, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *EvalsJudgeSpec) SetCriteria(v string)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *EvalsJudgeSpec) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetName

`func (o *EvalsJudgeSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvalsJudgeSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvalsJudgeSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EvalsJudgeSpec) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


