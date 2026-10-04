# TaxAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Box** | Pointer to **string** | Box is the box number, e.g. \&quot;1a\&quot;. | [optional] 
**Cents** | Pointer to **int64** | Cents is the amount in cents. | [optional] 
**Label** | Pointer to **string** | Label is the box&#39;s caption on the form. | [optional] 

## Methods

### NewTaxAmount

`func NewTaxAmount() *TaxAmount`

NewTaxAmount instantiates a new TaxAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxAmountWithDefaults

`func NewTaxAmountWithDefaults() *TaxAmount`

NewTaxAmountWithDefaults instantiates a new TaxAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBox

`func (o *TaxAmount) GetBox() string`

GetBox returns the Box field if non-nil, zero value otherwise.

### GetBoxOk

`func (o *TaxAmount) GetBoxOk() (*string, bool)`

GetBoxOk returns a tuple with the Box field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBox

`func (o *TaxAmount) SetBox(v string)`

SetBox sets Box field to given value.

### HasBox

`func (o *TaxAmount) HasBox() bool`

HasBox returns a boolean if a field has been set.

### GetCents

`func (o *TaxAmount) GetCents() int64`

GetCents returns the Cents field if non-nil, zero value otherwise.

### GetCentsOk

`func (o *TaxAmount) GetCentsOk() (*int64, bool)`

GetCentsOk returns a tuple with the Cents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCents

`func (o *TaxAmount) SetCents(v int64)`

SetCents sets Cents field to given value.

### HasCents

`func (o *TaxAmount) HasCents() bool`

HasCents returns a boolean if a field has been set.

### GetLabel

`func (o *TaxAmount) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TaxAmount) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TaxAmount) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TaxAmount) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


