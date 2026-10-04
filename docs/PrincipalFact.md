# PrincipalFact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocks** | Pointer to **bool** | Blocks is whether the payment waits on it. | [optional] 
**Code** | Pointer to **string** | Code is the fact&#39;s stable name. | [optional] 
**Question** | Pointer to **string** | Question is what must be answered. | [optional] 
**Rule** | Pointer to [**PrincipalRule**](PrincipalRule.md) | Rule is the rule that turns on it. | [optional] 

## Methods

### NewPrincipalFact

`func NewPrincipalFact() *PrincipalFact`

NewPrincipalFact instantiates a new PrincipalFact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalFactWithDefaults

`func NewPrincipalFactWithDefaults() *PrincipalFact`

NewPrincipalFactWithDefaults instantiates a new PrincipalFact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocks

`func (o *PrincipalFact) GetBlocks() bool`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *PrincipalFact) GetBlocksOk() (*bool, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *PrincipalFact) SetBlocks(v bool)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *PrincipalFact) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetCode

`func (o *PrincipalFact) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PrincipalFact) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PrincipalFact) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PrincipalFact) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetQuestion

`func (o *PrincipalFact) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *PrincipalFact) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *PrincipalFact) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *PrincipalFact) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetRule

`func (o *PrincipalFact) GetRule() PrincipalRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *PrincipalFact) GetRuleOk() (*PrincipalRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *PrincipalFact) SetRule(v PrincipalRule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *PrincipalFact) HasRule() bool`

HasRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


