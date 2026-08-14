# RaiseInvoiceIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Currency is the ISO 4217 code, lower-cased. Empty means usd. | [optional] 
**CustomerEmail** | Pointer to **string** | CustomerEmail is where the invoice is sent. Optional. | [optional] 
**Lines** | Pointer to [**[]InvoiceLineIn**](InvoiceLineIn.md) | Lines are the charges. The invoice subtotal and amount due are COMPUTED from these — there is no total field to send, because a total that disagreed with its own lines would bill a number nobody could derive. | [optional] 
**UserId** | Pointer to **string** | UserID identifies the customer being billed, within the caller&#39;s own org. Required — an invoice with no addressee is not an invoice. | [optional] 

## Methods

### NewRaiseInvoiceIn

`func NewRaiseInvoiceIn() *RaiseInvoiceIn`

NewRaiseInvoiceIn instantiates a new RaiseInvoiceIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRaiseInvoiceInWithDefaults

`func NewRaiseInvoiceInWithDefaults() *RaiseInvoiceIn`

NewRaiseInvoiceInWithDefaults instantiates a new RaiseInvoiceIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *RaiseInvoiceIn) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RaiseInvoiceIn) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RaiseInvoiceIn) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *RaiseInvoiceIn) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *RaiseInvoiceIn) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *RaiseInvoiceIn) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *RaiseInvoiceIn) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *RaiseInvoiceIn) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetLines

`func (o *RaiseInvoiceIn) GetLines() []InvoiceLineIn`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *RaiseInvoiceIn) GetLinesOk() (*[]InvoiceLineIn, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *RaiseInvoiceIn) SetLines(v []InvoiceLineIn)`

SetLines sets Lines field to given value.

### HasLines

`func (o *RaiseInvoiceIn) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetUserId

`func (o *RaiseInvoiceIn) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RaiseInvoiceIn) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RaiseInvoiceIn) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RaiseInvoiceIn) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


