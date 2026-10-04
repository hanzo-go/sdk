# BillingInvoiceLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | Amount is the line total in whole cents (250000 is $2,500.00). | [optional] 
**Description** | Pointer to **string** | Description is the human-readable line, e.g. \&quot;Advisory retainer — August\&quot;. | [optional] 
**Quantity** | Pointer to **int64** | Quantity is the number of units, when the line is metered. Optional. | [optional] 
**UnitPrice** | Pointer to **int64** | UnitPrice is the per-unit price in cents, when the line is metered. Optional. | [optional] 

## Methods

### NewBillingInvoiceLine

`func NewBillingInvoiceLine() *BillingInvoiceLine`

NewBillingInvoiceLine instantiates a new BillingInvoiceLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInvoiceLineWithDefaults

`func NewBillingInvoiceLineWithDefaults() *BillingInvoiceLine`

NewBillingInvoiceLineWithDefaults instantiates a new BillingInvoiceLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BillingInvoiceLine) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingInvoiceLine) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingInvoiceLine) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingInvoiceLine) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDescription

`func (o *BillingInvoiceLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingInvoiceLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingInvoiceLine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingInvoiceLine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuantity

`func (o *BillingInvoiceLine) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillingInvoiceLine) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillingInvoiceLine) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BillingInvoiceLine) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitPrice

`func (o *BillingInvoiceLine) GetUnitPrice() int64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *BillingInvoiceLine) GetUnitPriceOk() (*int64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *BillingInvoiceLine) SetUnitPrice(v int64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *BillingInvoiceLine) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


