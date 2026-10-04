# TaxRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the rule&#39;s stable name, for a caller that branches on it. | [optional] 
**Reason** | Pointer to **string** | Reason is the rule, stated. | [optional] 

## Methods

### NewTaxRule

`func NewTaxRule() *TaxRule`

NewTaxRule instantiates a new TaxRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxRuleWithDefaults

`func NewTaxRuleWithDefaults() *TaxRule`

NewTaxRuleWithDefaults instantiates a new TaxRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TaxRule) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TaxRule) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TaxRule) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TaxRule) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetReason

`func (o *TaxRule) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TaxRule) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TaxRule) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TaxRule) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


