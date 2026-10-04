# BillingInvoices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Cursor** | Pointer to **string** |  | [optional] 
**Invoices** | Pointer to [**[]BillingBillingInvoice**](BillingBillingInvoice.md) |  | [optional] 

## Methods

### NewBillingInvoices

`func NewBillingInvoices() *BillingInvoices`

NewBillingInvoices instantiates a new BillingInvoices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInvoicesWithDefaults

`func NewBillingInvoicesWithDefaults() *BillingInvoices`

NewBillingInvoicesWithDefaults instantiates a new BillingInvoices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *BillingInvoices) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BillingInvoices) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BillingInvoices) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *BillingInvoices) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCursor

`func (o *BillingInvoices) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *BillingInvoices) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *BillingInvoices) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *BillingInvoices) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetInvoices

`func (o *BillingInvoices) GetInvoices() []BillingBillingInvoice`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *BillingInvoices) GetInvoicesOk() (*[]BillingBillingInvoice, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *BillingInvoices) SetInvoices(v []BillingBillingInvoice)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *BillingInvoices) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


