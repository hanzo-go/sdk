# TaxMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int64** | At is when the payer recorded it, unix seconds. | [optional] 
**By** | Pointer to **string** | By is the admin who recorded it. | [optional] 
**Result** | Pointer to **string** | Result is what IRS TIN Matching answered: matched or mismatched. | [optional] 

## Methods

### NewTaxMatch

`func NewTaxMatch() *TaxMatch`

NewTaxMatch instantiates a new TaxMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxMatchWithDefaults

`func NewTaxMatchWithDefaults() *TaxMatch`

NewTaxMatchWithDefaults instantiates a new TaxMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *TaxMatch) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *TaxMatch) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *TaxMatch) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *TaxMatch) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBy

`func (o *TaxMatch) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *TaxMatch) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *TaxMatch) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *TaxMatch) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetResult

`func (o *TaxMatch) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TaxMatch) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TaxMatch) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *TaxMatch) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


