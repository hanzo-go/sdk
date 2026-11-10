# AiDecisionsBoolean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | Pointer to [**AiDecisionSides**](AiDecisionSides.md) |  | [optional] 
**Instructions** | Pointer to **interface{}** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewAiDecisionsBoolean

`func NewAiDecisionsBoolean(type_ string, ) *AiDecisionsBoolean`

NewAiDecisionsBoolean instantiates a new AiDecisionsBoolean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiDecisionsBooleanWithDefaults

`func NewAiDecisionsBooleanWithDefaults() *AiDecisionsBoolean`

NewAiDecisionsBooleanWithDefaults instantiates a new AiDecisionsBoolean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *AiDecisionsBoolean) GetCriteria() AiDecisionSides`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *AiDecisionsBoolean) GetCriteriaOk() (*AiDecisionSides, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *AiDecisionsBoolean) SetCriteria(v AiDecisionSides)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *AiDecisionsBoolean) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetInstructions

`func (o *AiDecisionsBoolean) GetInstructions() interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AiDecisionsBoolean) GetInstructionsOk() (*interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AiDecisionsBoolean) SetInstructions(v interface{})`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AiDecisionsBoolean) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *AiDecisionsBoolean) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *AiDecisionsBoolean) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetLabels

`func (o *AiDecisionsBoolean) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AiDecisionsBoolean) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AiDecisionsBoolean) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AiDecisionsBoolean) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *AiDecisionsBoolean) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiDecisionsBoolean) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiDecisionsBoolean) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


