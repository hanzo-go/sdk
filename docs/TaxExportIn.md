# TaxExportIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Kind is 1099-NEC or 1099-MISC. | [optional] 
**Year** | Pointer to **int64** | Year is the tax year. | [optional] 

## Methods

### NewTaxExportIn

`func NewTaxExportIn() *TaxExportIn`

NewTaxExportIn instantiates a new TaxExportIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxExportInWithDefaults

`func NewTaxExportInWithDefaults() *TaxExportIn`

NewTaxExportInWithDefaults instantiates a new TaxExportIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TaxExportIn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TaxExportIn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TaxExportIn) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TaxExportIn) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetYear

`func (o *TaxExportIn) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *TaxExportIn) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *TaxExportIn) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *TaxExportIn) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


