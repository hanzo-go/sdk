# O11yAlertRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Evaluator** | Pointer to [**O11yAlertRuleConditionEvaluator**](O11yAlertRuleConditionEvaluator.md) |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yAlertRuleCondition

`func NewO11yAlertRuleCondition() *O11yAlertRuleCondition`

NewO11yAlertRuleCondition instantiates a new O11yAlertRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAlertRuleConditionWithDefaults

`func NewO11yAlertRuleConditionWithDefaults() *O11yAlertRuleCondition`

NewO11yAlertRuleConditionWithDefaults instantiates a new O11yAlertRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvaluator

`func (o *O11yAlertRuleCondition) GetEvaluator() O11yAlertRuleConditionEvaluator`

GetEvaluator returns the Evaluator field if non-nil, zero value otherwise.

### GetEvaluatorOk

`func (o *O11yAlertRuleCondition) GetEvaluatorOk() (*O11yAlertRuleConditionEvaluator, bool)`

GetEvaluatorOk returns a tuple with the Evaluator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluator

`func (o *O11yAlertRuleCondition) SetEvaluator(v O11yAlertRuleConditionEvaluator)`

SetEvaluator sets Evaluator field to given value.

### HasEvaluator

`func (o *O11yAlertRuleCondition) HasEvaluator() bool`

HasEvaluator returns a boolean if a field has been set.

### GetOperator

`func (o *O11yAlertRuleCondition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *O11yAlertRuleCondition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *O11yAlertRuleCondition) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *O11yAlertRuleCondition) HasOperator() bool`

HasOperator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


