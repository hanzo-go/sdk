# InvoiceOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountDueCents** | Pointer to **int32** | AmountDueCents is what remains collectible. | [optional] 
**AmountPaidCents** | Pointer to **int32** | AmountPaidCents is what has been collected so far. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when the draft was raised, RFC3339. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO 4217 code. | [optional] 
**CustomerEmail** | Pointer to **string** | CustomerEmail is where it is sent. | [optional] 
**Id** | Pointer to **string** | ID is the invoice id — what the issue, collect and void ops address. | [optional] 
**Lines** | Pointer to [**[]InvoiceLineIn**](InvoiceLineIn.md) | Lines are the charges on the invoice. | [optional] 
**Number** | Pointer to **string** | Number is the human-facing invoice number, e.g. \&quot;INV-0042\&quot;. A draft has none; issuing assigns it. | [optional] 
**PaymentRef** | Pointer to **string** | PaymentRef is the processor reference for the collection, once paid. | [optional] 
**Status** | Pointer to **string** | Status is draft, open, paid, void or uncollectible. A draft is not collectible; issuing moves it to open. | [optional] 
**SubtotalCents** | Pointer to **int32** | SubtotalCents is the sum of the lines. | [optional] 
**UserId** | Pointer to **string** | UserID is the customer billed. | [optional] 

## Methods

### NewInvoiceOut

`func NewInvoiceOut() *InvoiceOut`

NewInvoiceOut instantiates a new InvoiceOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceOutWithDefaults

`func NewInvoiceOutWithDefaults() *InvoiceOut`

NewInvoiceOutWithDefaults instantiates a new InvoiceOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountDueCents

`func (o *InvoiceOut) GetAmountDueCents() int32`

GetAmountDueCents returns the AmountDueCents field if non-nil, zero value otherwise.

### GetAmountDueCentsOk

`func (o *InvoiceOut) GetAmountDueCentsOk() (*int32, bool)`

GetAmountDueCentsOk returns a tuple with the AmountDueCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDueCents

`func (o *InvoiceOut) SetAmountDueCents(v int32)`

SetAmountDueCents sets AmountDueCents field to given value.

### HasAmountDueCents

`func (o *InvoiceOut) HasAmountDueCents() bool`

HasAmountDueCents returns a boolean if a field has been set.

### GetAmountPaidCents

`func (o *InvoiceOut) GetAmountPaidCents() int32`

GetAmountPaidCents returns the AmountPaidCents field if non-nil, zero value otherwise.

### GetAmountPaidCentsOk

`func (o *InvoiceOut) GetAmountPaidCentsOk() (*int32, bool)`

GetAmountPaidCentsOk returns a tuple with the AmountPaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaidCents

`func (o *InvoiceOut) SetAmountPaidCents(v int32)`

SetAmountPaidCents sets AmountPaidCents field to given value.

### HasAmountPaidCents

`func (o *InvoiceOut) HasAmountPaidCents() bool`

HasAmountPaidCents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InvoiceOut) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceOut) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceOut) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InvoiceOut) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *InvoiceOut) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceOut) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceOut) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InvoiceOut) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *InvoiceOut) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *InvoiceOut) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *InvoiceOut) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *InvoiceOut) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetId

`func (o *InvoiceOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceOut) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceOut) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLines

`func (o *InvoiceOut) GetLines() []InvoiceLineIn`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *InvoiceOut) GetLinesOk() (*[]InvoiceLineIn, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *InvoiceOut) SetLines(v []InvoiceLineIn)`

SetLines sets Lines field to given value.

### HasLines

`func (o *InvoiceOut) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetNumber

`func (o *InvoiceOut) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InvoiceOut) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InvoiceOut) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InvoiceOut) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPaymentRef

`func (o *InvoiceOut) GetPaymentRef() string`

GetPaymentRef returns the PaymentRef field if non-nil, zero value otherwise.

### GetPaymentRefOk

`func (o *InvoiceOut) GetPaymentRefOk() (*string, bool)`

GetPaymentRefOk returns a tuple with the PaymentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRef

`func (o *InvoiceOut) SetPaymentRef(v string)`

SetPaymentRef sets PaymentRef field to given value.

### HasPaymentRef

`func (o *InvoiceOut) HasPaymentRef() bool`

HasPaymentRef returns a boolean if a field has been set.

### GetStatus

`func (o *InvoiceOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InvoiceOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubtotalCents

`func (o *InvoiceOut) GetSubtotalCents() int32`

GetSubtotalCents returns the SubtotalCents field if non-nil, zero value otherwise.

### GetSubtotalCentsOk

`func (o *InvoiceOut) GetSubtotalCentsOk() (*int32, bool)`

GetSubtotalCentsOk returns a tuple with the SubtotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalCents

`func (o *InvoiceOut) SetSubtotalCents(v int32)`

SetSubtotalCents sets SubtotalCents field to given value.

### HasSubtotalCents

`func (o *InvoiceOut) HasSubtotalCents() bool`

HasSubtotalCents returns a boolean if a field has been set.

### GetUserId

`func (o *InvoiceOut) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *InvoiceOut) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *InvoiceOut) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *InvoiceOut) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


