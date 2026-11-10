# AiDecisionsQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | **[]interface{}** |  | 
**Instructions** | Pointer to **interface{}** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewAiDecisionsQuestion

`func NewAiDecisionsQuestion(criteria []interface{}, type_ string, ) *AiDecisionsQuestion`

NewAiDecisionsQuestion instantiates a new AiDecisionsQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiDecisionsQuestionWithDefaults

`func NewAiDecisionsQuestionWithDefaults() *AiDecisionsQuestion`

NewAiDecisionsQuestionWithDefaults instantiates a new AiDecisionsQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *AiDecisionsQuestion) GetCriteria() []interface{}`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *AiDecisionsQuestion) GetCriteriaOk() (*[]interface{}, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *AiDecisionsQuestion) SetCriteria(v []interface{})`

SetCriteria sets Criteria field to given value.


### GetInstructions

`func (o *AiDecisionsQuestion) GetInstructions() interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AiDecisionsQuestion) GetInstructionsOk() (*interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AiDecisionsQuestion) SetInstructions(v interface{})`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AiDecisionsQuestion) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *AiDecisionsQuestion) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *AiDecisionsQuestion) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetLabels

`func (o *AiDecisionsQuestion) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AiDecisionsQuestion) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AiDecisionsQuestion) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AiDecisionsQuestion) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *AiDecisionsQuestion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiDecisionsQuestion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiDecisionsQuestion) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


