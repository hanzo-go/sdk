# TaxVerdict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Box** | Pointer to [**TaxBox**](TaxBox.md) | Box is where it is reported, when it is. | [optional] 
**Reportable** | Pointer to **bool** |  | [optional] 
**Rules** | Pointer to [**[]TaxRule**](TaxRule.md) | Rules are every rule that applied, in the order applied — the last one is the verdict&#39;s reason. | [optional] 

## Methods

### NewTaxVerdict

`func NewTaxVerdict() *TaxVerdict`

NewTaxVerdict instantiates a new TaxVerdict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxVerdictWithDefaults

`func NewTaxVerdictWithDefaults() *TaxVerdict`

NewTaxVerdictWithDefaults instantiates a new TaxVerdict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBox

`func (o *TaxVerdict) GetBox() TaxBox`

GetBox returns the Box field if non-nil, zero value otherwise.

### GetBoxOk

`func (o *TaxVerdict) GetBoxOk() (*TaxBox, bool)`

GetBoxOk returns a tuple with the Box field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBox

`func (o *TaxVerdict) SetBox(v TaxBox)`

SetBox sets Box field to given value.

### HasBox

`func (o *TaxVerdict) HasBox() bool`

HasBox returns a boolean if a field has been set.

### GetReportable

`func (o *TaxVerdict) GetReportable() bool`

GetReportable returns the Reportable field if non-nil, zero value otherwise.

### GetReportableOk

`func (o *TaxVerdict) GetReportableOk() (*bool, bool)`

GetReportableOk returns a tuple with the Reportable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportable

`func (o *TaxVerdict) SetReportable(v bool)`

SetReportable sets Reportable field to given value.

### HasReportable

`func (o *TaxVerdict) HasReportable() bool`

HasReportable returns a boolean if a field has been set.

### GetRules

`func (o *TaxVerdict) GetRules() []TaxRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *TaxVerdict) GetRulesOk() (*[]TaxRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *TaxVerdict) SetRules(v []TaxRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *TaxVerdict) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


