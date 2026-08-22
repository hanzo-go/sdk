# Invoices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Invoices** | Pointer to [**[]BillingInvoice**](BillingInvoice.md) |  | [optional] 

## Methods

### NewInvoices

`func NewInvoices() *Invoices`

NewInvoices instantiates a new Invoices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicesWithDefaults

`func NewInvoicesWithDefaults() *Invoices`

NewInvoicesWithDefaults instantiates a new Invoices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *Invoices) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Invoices) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Invoices) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Invoices) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetInvoices

`func (o *Invoices) GetInvoices() []BillingInvoice`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *Invoices) GetInvoicesOk() (*[]BillingInvoice, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *Invoices) SetInvoices(v []BillingInvoice)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *Invoices) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


