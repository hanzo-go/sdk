# TaxPayeeYear

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category is what the payer states it bought from this payee; services when it has stated nothing. | [optional] 
**Lines** | Pointer to [**[]TaxLine**](TaxLine.md) | Lines are the payments, oldest first. | [optional] 
**Payee** | Pointer to **string** | Payee is the org paid. | [optional] 
**Totals** | Pointer to [**[]TaxTotal**](TaxTotal.md) | Totals are the reportable sums by box. | [optional] 
**W9** | Pointer to **string** | W9 is where the payer&#39;s W-9 request to this payee stands: none, requested, granted, declined or revoked. | [optional] 

## Methods

### NewTaxPayeeYear

`func NewTaxPayeeYear() *TaxPayeeYear`

NewTaxPayeeYear instantiates a new TaxPayeeYear object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxPayeeYearWithDefaults

`func NewTaxPayeeYearWithDefaults() *TaxPayeeYear`

NewTaxPayeeYearWithDefaults instantiates a new TaxPayeeYear object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *TaxPayeeYear) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TaxPayeeYear) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TaxPayeeYear) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TaxPayeeYear) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetLines

`func (o *TaxPayeeYear) GetLines() []TaxLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *TaxPayeeYear) GetLinesOk() (*[]TaxLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *TaxPayeeYear) SetLines(v []TaxLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *TaxPayeeYear) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetPayee

`func (o *TaxPayeeYear) GetPayee() string`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *TaxPayeeYear) GetPayeeOk() (*string, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *TaxPayeeYear) SetPayee(v string)`

SetPayee sets Payee field to given value.

### HasPayee

`func (o *TaxPayeeYear) HasPayee() bool`

HasPayee returns a boolean if a field has been set.

### GetTotals

`func (o *TaxPayeeYear) GetTotals() []TaxTotal`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *TaxPayeeYear) GetTotalsOk() (*[]TaxTotal, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *TaxPayeeYear) SetTotals(v []TaxTotal)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *TaxPayeeYear) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetW9

`func (o *TaxPayeeYear) GetW9() string`

GetW9 returns the W9 field if non-nil, zero value otherwise.

### GetW9Ok

`func (o *TaxPayeeYear) GetW9Ok() (*string, bool)`

GetW9Ok returns a tuple with the W9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW9

`func (o *TaxPayeeYear) SetW9(v string)`

SetW9 sets W9 field to given value.

### HasW9

`func (o *TaxPayeeYear) HasW9() bool`

HasW9 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


