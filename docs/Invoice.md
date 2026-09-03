# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountDueCents** | Pointer to **int64** | AmountDueCents is what remains collectible. | [optional] 
**AmountPaidCents** | Pointer to **int64** | AmountPaidCents is what has been collected so far. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when the draft was raised, RFC3339. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO 4217 code. | [optional] 
**CustomerEmail** | Pointer to **string** | CustomerEmail is where it is sent. | [optional] 
**Id** | Pointer to **string** | ID is the invoice id — what the issue, collect and void ops address. | [optional] 
**Lines** | Pointer to [**[]InvoiceLine**](InvoiceLine.md) | Lines are the charges on the invoice. | [optional] 
**Number** | Pointer to **string** | Number is the human-facing invoice number, e.g. \&quot;INV-0042\&quot;. A draft has none; issuing assigns it. | [optional] 
**PaymentRef** | Pointer to **string** | PaymentRef is the processor reference for the collection, once paid. | [optional] 
**Status** | Pointer to **string** | Status is draft, open, paid, void or uncollectible. A draft is not collectible; issuing moves it to open. | [optional] 
**SubtotalCents** | Pointer to **int64** | SubtotalCents is the sum of the lines. | [optional] 
**UserId** | Pointer to **string** | UserID is the customer billed. | [optional] 

## Methods

### NewInvoice

`func NewInvoice() *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountDueCents

`func (o *Invoice) GetAmountDueCents() int64`

GetAmountDueCents returns the AmountDueCents field if non-nil, zero value otherwise.

### GetAmountDueCentsOk

`func (o *Invoice) GetAmountDueCentsOk() (*int64, bool)`

GetAmountDueCentsOk returns a tuple with the AmountDueCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDueCents

`func (o *Invoice) SetAmountDueCents(v int64)`

SetAmountDueCents sets AmountDueCents field to given value.

### HasAmountDueCents

`func (o *Invoice) HasAmountDueCents() bool`

HasAmountDueCents returns a boolean if a field has been set.

### GetAmountPaidCents

`func (o *Invoice) GetAmountPaidCents() int64`

GetAmountPaidCents returns the AmountPaidCents field if non-nil, zero value otherwise.

### GetAmountPaidCentsOk

`func (o *Invoice) GetAmountPaidCentsOk() (*int64, bool)`

GetAmountPaidCentsOk returns a tuple with the AmountPaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaidCents

`func (o *Invoice) SetAmountPaidCents(v int64)`

SetAmountPaidCents sets AmountPaidCents field to given value.

### HasAmountPaidCents

`func (o *Invoice) HasAmountPaidCents() bool`

HasAmountPaidCents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Invoice) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Invoice) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Invoice) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Invoice) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *Invoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Invoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Invoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Invoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *Invoice) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *Invoice) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *Invoice) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *Invoice) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetId

`func (o *Invoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Invoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLines

`func (o *Invoice) GetLines() []InvoiceLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *Invoice) GetLinesOk() (*[]InvoiceLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *Invoice) SetLines(v []InvoiceLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *Invoice) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetNumber

`func (o *Invoice) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Invoice) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Invoice) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Invoice) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPaymentRef

`func (o *Invoice) GetPaymentRef() string`

GetPaymentRef returns the PaymentRef field if non-nil, zero value otherwise.

### GetPaymentRefOk

`func (o *Invoice) GetPaymentRefOk() (*string, bool)`

GetPaymentRefOk returns a tuple with the PaymentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRef

`func (o *Invoice) SetPaymentRef(v string)`

SetPaymentRef sets PaymentRef field to given value.

### HasPaymentRef

`func (o *Invoice) HasPaymentRef() bool`

HasPaymentRef returns a boolean if a field has been set.

### GetStatus

`func (o *Invoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Invoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Invoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubtotalCents

`func (o *Invoice) GetSubtotalCents() int64`

GetSubtotalCents returns the SubtotalCents field if non-nil, zero value otherwise.

### GetSubtotalCentsOk

`func (o *Invoice) GetSubtotalCentsOk() (*int64, bool)`

GetSubtotalCentsOk returns a tuple with the SubtotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalCents

`func (o *Invoice) SetSubtotalCents(v int64)`

SetSubtotalCents sets SubtotalCents field to given value.

### HasSubtotalCents

`func (o *Invoice) HasSubtotalCents() bool`

HasSubtotalCents returns a boolean if a field has been set.

### GetUserId

`func (o *Invoice) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Invoice) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Invoice) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Invoice) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


