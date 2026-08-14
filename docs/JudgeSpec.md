# JudgeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | Pointer to **string** | Criteria is the standard the judge applies, defaulting to a correctness criterion. | [optional] 
**Model** | Pointer to **string** | Model is the model that grades. It defaults to the model under test, so a run with no judge named has the model grade itself. | [optional] 
**Name** | Pointer to **string** | Name is the score name the judge&#39;s grades are filed under, \&quot;llm-judge\&quot; by default. A name that is not a legal handle falls back to that default rather than being stored as sent. | [optional] 

## Methods

### NewJudgeSpec

`func NewJudgeSpec() *JudgeSpec`

NewJudgeSpec instantiates a new JudgeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJudgeSpecWithDefaults

`func NewJudgeSpecWithDefaults() *JudgeSpec`

NewJudgeSpecWithDefaults instantiates a new JudgeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *JudgeSpec) GetCriteria() string`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *JudgeSpec) GetCriteriaOk() (*string, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *JudgeSpec) SetCriteria(v string)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *JudgeSpec) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetModel

`func (o *JudgeSpec) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *JudgeSpec) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *JudgeSpec) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *JudgeSpec) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *JudgeSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JudgeSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JudgeSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JudgeSpec) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


