# TaxW9Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category is what the caller pays it for: services, attorney, rents, royalties, other or merchandise. Services when omitted. | [optional] 
**Payee** | Pointer to **string** | Payee is the org id of the company asked. The caller must have paid it on Hanzo&#39;s rails. | [optional] 

## Methods

### NewTaxW9Request

`func NewTaxW9Request() *TaxW9Request`

NewTaxW9Request instantiates a new TaxW9Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxW9RequestWithDefaults

`func NewTaxW9RequestWithDefaults() *TaxW9Request`

NewTaxW9RequestWithDefaults instantiates a new TaxW9Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *TaxW9Request) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TaxW9Request) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TaxW9Request) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TaxW9Request) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetPayee

`func (o *TaxW9Request) GetPayee() string`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *TaxW9Request) GetPayeeOk() (*string, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *TaxW9Request) SetPayee(v string)`

SetPayee sets Payee field to given value.

### HasPayee

`func (o *TaxW9Request) HasPayee() bool`

HasPayee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


