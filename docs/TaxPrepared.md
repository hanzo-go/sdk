# TaxPrepared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Forms** | Pointer to [**[]TaxForm**](TaxForm.md) | Forms are the drafts made or refreshed. | [optional] 
**Skipped** | Pointer to **[]string** | Skipped says, per payee and form, why nothing was prepared — a form already furnished is never re-prepared, only corrected. | [optional] 
**Year** | Pointer to **int64** | Year is the tax year. | [optional] 

## Methods

### NewTaxPrepared

`func NewTaxPrepared() *TaxPrepared`

NewTaxPrepared instantiates a new TaxPrepared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxPreparedWithDefaults

`func NewTaxPreparedWithDefaults() *TaxPrepared`

NewTaxPreparedWithDefaults instantiates a new TaxPrepared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForms

`func (o *TaxPrepared) GetForms() []TaxForm`

GetForms returns the Forms field if non-nil, zero value otherwise.

### GetFormsOk

`func (o *TaxPrepared) GetFormsOk() (*[]TaxForm, bool)`

GetFormsOk returns a tuple with the Forms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForms

`func (o *TaxPrepared) SetForms(v []TaxForm)`

SetForms sets Forms field to given value.

### HasForms

`func (o *TaxPrepared) HasForms() bool`

HasForms returns a boolean if a field has been set.

### GetSkipped

`func (o *TaxPrepared) GetSkipped() []string`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *TaxPrepared) GetSkippedOk() (*[]string, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *TaxPrepared) SetSkipped(v []string)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *TaxPrepared) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetYear

`func (o *TaxPrepared) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *TaxPrepared) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *TaxPrepared) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *TaxPrepared) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


