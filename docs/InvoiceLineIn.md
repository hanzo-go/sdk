# InvoiceLineIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | Amount is the line total in whole cents (250000 is $2,500.00). | [optional] 
**Description** | Pointer to **string** | Description is the human-readable line, e.g. \&quot;Advisory retainer — August\&quot;. | [optional] 
**Quantity** | Pointer to **int32** | Quantity is the number of units, when the line is metered. Optional. | [optional] 
**UnitPrice** | Pointer to **int32** | UnitPrice is the per-unit price in cents, when the line is metered. Optional. | [optional] 

## Methods

### NewInvoiceLineIn

`func NewInvoiceLineIn() *InvoiceLineIn`

NewInvoiceLineIn instantiates a new InvoiceLineIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineInWithDefaults

`func NewInvoiceLineInWithDefaults() *InvoiceLineIn`

NewInvoiceLineInWithDefaults instantiates a new InvoiceLineIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InvoiceLineIn) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceLineIn) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceLineIn) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InvoiceLineIn) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDescription

`func (o *InvoiceLineIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceLineIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceLineIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceLineIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuantity

`func (o *InvoiceLineIn) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceLineIn) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceLineIn) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InvoiceLineIn) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitPrice

`func (o *InvoiceLineIn) GetUnitPrice() int32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *InvoiceLineIn) GetUnitPriceOk() (*int32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *InvoiceLineIn) SetUnitPrice(v int32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *InvoiceLineIn) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


