# TaxTotal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int64** | AmountCents is the reportable payments in this box, summed exactly and then rounded to cents. | [optional] 
**Box** | Pointer to [**TaxBox**](TaxBox.md) | Box is the form and box. | [optional] 
**Payments** | Pointer to **[]string** | Payments are the ids that sum into it. | [optional] 
**Reportable** | Pointer to **bool** | Reportable is true when the total is at or above the threshold. | [optional] 
**Rule** | Pointer to [**TaxRule**](TaxRule.md) | Rule states the comparison and its source. | [optional] 
**ThresholdCents** | Pointer to **int64** | ThresholdCents is the year&#39;s threshold for this box. | [optional] 

## Methods

### NewTaxTotal

`func NewTaxTotal() *TaxTotal`

NewTaxTotal instantiates a new TaxTotal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxTotalWithDefaults

`func NewTaxTotalWithDefaults() *TaxTotal`

NewTaxTotalWithDefaults instantiates a new TaxTotal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *TaxTotal) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *TaxTotal) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *TaxTotal) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *TaxTotal) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetBox

`func (o *TaxTotal) GetBox() TaxBox`

GetBox returns the Box field if non-nil, zero value otherwise.

### GetBoxOk

`func (o *TaxTotal) GetBoxOk() (*TaxBox, bool)`

GetBoxOk returns a tuple with the Box field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBox

`func (o *TaxTotal) SetBox(v TaxBox)`

SetBox sets Box field to given value.

### HasBox

`func (o *TaxTotal) HasBox() bool`

HasBox returns a boolean if a field has been set.

### GetPayments

`func (o *TaxTotal) GetPayments() []string`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *TaxTotal) GetPaymentsOk() (*[]string, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *TaxTotal) SetPayments(v []string)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *TaxTotal) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetReportable

`func (o *TaxTotal) GetReportable() bool`

GetReportable returns the Reportable field if non-nil, zero value otherwise.

### GetReportableOk

`func (o *TaxTotal) GetReportableOk() (*bool, bool)`

GetReportableOk returns a tuple with the Reportable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportable

`func (o *TaxTotal) SetReportable(v bool)`

SetReportable sets Reportable field to given value.

### HasReportable

`func (o *TaxTotal) HasReportable() bool`

HasReportable returns a boolean if a field has been set.

### GetRule

`func (o *TaxTotal) GetRule() TaxRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *TaxTotal) GetRuleOk() (*TaxRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *TaxTotal) SetRule(v TaxRule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *TaxTotal) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetThresholdCents

`func (o *TaxTotal) GetThresholdCents() int64`

GetThresholdCents returns the ThresholdCents field if non-nil, zero value otherwise.

### GetThresholdCentsOk

`func (o *TaxTotal) GetThresholdCentsOk() (*int64, bool)`

GetThresholdCentsOk returns a tuple with the ThresholdCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdCents

`func (o *TaxTotal) SetThresholdCents(v int64)`

SetThresholdCents sets ThresholdCents field to given value.

### HasThresholdCents

`func (o *TaxTotal) HasThresholdCents() bool`

HasThresholdCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


