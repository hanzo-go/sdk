# AiDecisionsScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | [**[]AiDecisionSidesFalse**](AiDecisionSidesFalse.md) |  | 
**Instructions** | Pointer to [**AiDecisionSidesFalse**](AiDecisionSidesFalse.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewAiDecisionsScore

`func NewAiDecisionsScore(criteria []AiDecisionSidesFalse, type_ string, ) *AiDecisionsScore`

NewAiDecisionsScore instantiates a new AiDecisionsScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiDecisionsScoreWithDefaults

`func NewAiDecisionsScoreWithDefaults() *AiDecisionsScore`

NewAiDecisionsScoreWithDefaults instantiates a new AiDecisionsScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *AiDecisionsScore) GetCriteria() []AiDecisionSidesFalse`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *AiDecisionsScore) GetCriteriaOk() (*[]AiDecisionSidesFalse, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *AiDecisionsScore) SetCriteria(v []AiDecisionSidesFalse)`

SetCriteria sets Criteria field to given value.


### GetInstructions

`func (o *AiDecisionsScore) GetInstructions() AiDecisionSidesFalse`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AiDecisionsScore) GetInstructionsOk() (*AiDecisionSidesFalse, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AiDecisionsScore) SetInstructions(v AiDecisionSidesFalse)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AiDecisionsScore) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetType

`func (o *AiDecisionsScore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiDecisionsScore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiDecisionsScore) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


