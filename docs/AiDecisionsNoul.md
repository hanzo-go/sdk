# AiDecisionsNoul

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | Pointer to [**AiDecisionSides**](AiDecisionSides.md) |  | [optional] 
**Instructions** | Pointer to [**AiDecisionSidesFalse**](AiDecisionSidesFalse.md) |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewAiDecisionsNoul

`func NewAiDecisionsNoul(type_ string, ) *AiDecisionsNoul`

NewAiDecisionsNoul instantiates a new AiDecisionsNoul object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiDecisionsNoulWithDefaults

`func NewAiDecisionsNoulWithDefaults() *AiDecisionsNoul`

NewAiDecisionsNoulWithDefaults instantiates a new AiDecisionsNoul object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *AiDecisionsNoul) GetCriteria() AiDecisionSides`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *AiDecisionsNoul) GetCriteriaOk() (*AiDecisionSides, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *AiDecisionsNoul) SetCriteria(v AiDecisionSides)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *AiDecisionsNoul) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetInstructions

`func (o *AiDecisionsNoul) GetInstructions() AiDecisionSidesFalse`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AiDecisionsNoul) GetInstructionsOk() (*AiDecisionSidesFalse, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AiDecisionsNoul) SetInstructions(v AiDecisionSidesFalse)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AiDecisionsNoul) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetLabels

`func (o *AiDecisionsNoul) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AiDecisionsNoul) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AiDecisionsNoul) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AiDecisionsNoul) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *AiDecisionsNoul) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiDecisionsNoul) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiDecisionsNoul) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


