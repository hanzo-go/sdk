# TaxW9List

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsPayee** | Pointer to [**[]TaxW9**](TaxW9.md) | AsPayee are the requests for the caller&#39;s own W-9. | [optional] 
**AsPayer** | Pointer to [**[]TaxW9**](TaxW9.md) | AsPayer are the W-9s the caller asked for. | [optional] 

## Methods

### NewTaxW9List

`func NewTaxW9List() *TaxW9List`

NewTaxW9List instantiates a new TaxW9List object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxW9ListWithDefaults

`func NewTaxW9ListWithDefaults() *TaxW9List`

NewTaxW9ListWithDefaults instantiates a new TaxW9List object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsPayee

`func (o *TaxW9List) GetAsPayee() []TaxW9`

GetAsPayee returns the AsPayee field if non-nil, zero value otherwise.

### GetAsPayeeOk

`func (o *TaxW9List) GetAsPayeeOk() (*[]TaxW9, bool)`

GetAsPayeeOk returns a tuple with the AsPayee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsPayee

`func (o *TaxW9List) SetAsPayee(v []TaxW9)`

SetAsPayee sets AsPayee field to given value.

### HasAsPayee

`func (o *TaxW9List) HasAsPayee() bool`

HasAsPayee returns a boolean if a field has been set.

### GetAsPayer

`func (o *TaxW9List) GetAsPayer() []TaxW9`

GetAsPayer returns the AsPayer field if non-nil, zero value otherwise.

### GetAsPayerOk

`func (o *TaxW9List) GetAsPayerOk() (*[]TaxW9, bool)`

GetAsPayerOk returns a tuple with the AsPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsPayer

`func (o *TaxW9List) SetAsPayer(v []TaxW9)`

SetAsPayer sets AsPayer field to given value.

### HasAsPayer

`func (o *TaxW9List) HasAsPayer() bool`

HasAsPayer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


