# AiDecisionsChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | [**map[string]AiDecisionsChoiceCriteriaValue**](AiDecisionsChoiceCriteriaValue.md) |  | 
**Instructions** | Pointer to [**AiDecisionSidesFalse**](AiDecisionSidesFalse.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewAiDecisionsChoice

`func NewAiDecisionsChoice(criteria map[string]AiDecisionsChoiceCriteriaValue, type_ string, ) *AiDecisionsChoice`

NewAiDecisionsChoice instantiates a new AiDecisionsChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiDecisionsChoiceWithDefaults

`func NewAiDecisionsChoiceWithDefaults() *AiDecisionsChoice`

NewAiDecisionsChoiceWithDefaults instantiates a new AiDecisionsChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *AiDecisionsChoice) GetCriteria() map[string]AiDecisionsChoiceCriteriaValue`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *AiDecisionsChoice) GetCriteriaOk() (*map[string]AiDecisionsChoiceCriteriaValue, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *AiDecisionsChoice) SetCriteria(v map[string]AiDecisionsChoiceCriteriaValue)`

SetCriteria sets Criteria field to given value.


### GetInstructions

`func (o *AiDecisionsChoice) GetInstructions() AiDecisionSidesFalse`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AiDecisionsChoice) GetInstructionsOk() (*AiDecisionSidesFalse, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AiDecisionsChoice) SetInstructions(v AiDecisionSidesFalse)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AiDecisionsChoice) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetType

`func (o *AiDecisionsChoice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiDecisionsChoice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiDecisionsChoice) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


