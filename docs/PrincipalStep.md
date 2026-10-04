# PrincipalStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the step&#39;s stable name, for a caller that branches on it. | [optional] 
**Rule** | Pointer to [**PrincipalRule**](PrincipalRule.md) | Rule is why. | [optional] 
**What** | Pointer to **string** | What must happen, in words. | [optional] 
**Where** | Pointer to **string** | Where it is done, when it is done on this platform. | [optional] 
**Who** | Pointer to **string** | Who does it: payer, payee, or \&quot;Hanzo platform reviewer\&quot;. | [optional] 

## Methods

### NewPrincipalStep

`func NewPrincipalStep() *PrincipalStep`

NewPrincipalStep instantiates a new PrincipalStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalStepWithDefaults

`func NewPrincipalStepWithDefaults() *PrincipalStep`

NewPrincipalStepWithDefaults instantiates a new PrincipalStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PrincipalStep) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PrincipalStep) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PrincipalStep) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PrincipalStep) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRule

`func (o *PrincipalStep) GetRule() PrincipalRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *PrincipalStep) GetRuleOk() (*PrincipalRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *PrincipalStep) SetRule(v PrincipalRule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *PrincipalStep) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetWhat

`func (o *PrincipalStep) GetWhat() string`

GetWhat returns the What field if non-nil, zero value otherwise.

### GetWhatOk

`func (o *PrincipalStep) GetWhatOk() (*string, bool)`

GetWhatOk returns a tuple with the What field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhat

`func (o *PrincipalStep) SetWhat(v string)`

SetWhat sets What field to given value.

### HasWhat

`func (o *PrincipalStep) HasWhat() bool`

HasWhat returns a boolean if a field has been set.

### GetWhere

`func (o *PrincipalStep) GetWhere() string`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *PrincipalStep) GetWhereOk() (*string, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *PrincipalStep) SetWhere(v string)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *PrincipalStep) HasWhere() bool`

HasWhere returns a boolean if a field has been set.

### GetWho

`func (o *PrincipalStep) GetWho() string`

GetWho returns the Who field if non-nil, zero value otherwise.

### GetWhoOk

`func (o *PrincipalStep) GetWhoOk() (*string, bool)`

GetWhoOk returns a tuple with the Who field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWho

`func (o *PrincipalStep) SetWho(v string)`

SetWho sets Who field to given value.

### HasWho

`func (o *PrincipalStep) HasWho() bool`

HasWho returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


