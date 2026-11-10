# AiDecision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decision** | Pointer to **interface{}** |  | [optional] 
**Receipt** | Pointer to [**AiReceipt**](AiReceipt.md) | Receipt says which model served the decision and who paid for it. | [optional] 

## Methods

### NewAiDecision

`func NewAiDecision() *AiDecision`

NewAiDecision instantiates a new AiDecision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiDecisionWithDefaults

`func NewAiDecisionWithDefaults() *AiDecision`

NewAiDecisionWithDefaults instantiates a new AiDecision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecision

`func (o *AiDecision) GetDecision() interface{}`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *AiDecision) GetDecisionOk() (*interface{}, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *AiDecision) SetDecision(v interface{})`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *AiDecision) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### SetDecisionNil

`func (o *AiDecision) SetDecisionNil(b bool)`

 SetDecisionNil sets the value for Decision to be an explicit nil

### UnsetDecision
`func (o *AiDecision) UnsetDecision()`

UnsetDecision ensures that no value is present for Decision, not even an explicit nil
### GetReceipt

`func (o *AiDecision) GetReceipt() AiReceipt`

GetReceipt returns the Receipt field if non-nil, zero value otherwise.

### GetReceiptOk

`func (o *AiDecision) GetReceiptOk() (*AiReceipt, bool)`

GetReceiptOk returns a tuple with the Receipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipt

`func (o *AiDecision) SetReceipt(v AiReceipt)`

SetReceipt sets Receipt field to given value.

### HasReceipt

`func (o *AiDecision) HasReceipt() bool`

HasReceipt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


