# BillingInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountDue** | Pointer to **int64** |  | [optional] 
**AmountPaid** | Pointer to **int64** |  | [optional] 
**AttemptCount** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreditApplied** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CustomerEmail** | Pointer to **string** |  | [optional] 
**Discount** | Pointer to **int64** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LineItems** | Pointer to [**[]InvoiceLineItem**](InvoiceLineItem.md) | LineItems carries no omitempty and is never allocated empty, because the wire it reproduces sends &#x60;null&#x60; for an invoice with no lines. An empty array there would be a different answer to \&quot;were there lines\&quot;. | [optional] 
**Number** | Pointer to **int64** |  | [optional] 
**NumberStr** | Pointer to **string** |  | [optional] 
**PaidAt** | Pointer to **string** |  | [optional] 
**PaymentMethod** | Pointer to **string** |  | [optional] 
**PaymentRef** | Pointer to **string** |  | [optional] 
**PeriodEnd** | Pointer to **string** |  | [optional] 
**PeriodStart** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**Subtotal** | Pointer to **int64** |  | [optional] 
**Tax** | Pointer to **int64** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**VoidedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingInvoice

`func NewBillingInvoice() *BillingInvoice`

NewBillingInvoice instantiates a new BillingInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInvoiceWithDefaults

`func NewBillingInvoiceWithDefaults() *BillingInvoice`

NewBillingInvoiceWithDefaults instantiates a new BillingInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountDue

`func (o *BillingInvoice) GetAmountDue() int64`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *BillingInvoice) GetAmountDueOk() (*int64, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *BillingInvoice) SetAmountDue(v int64)`

SetAmountDue sets AmountDue field to given value.

### HasAmountDue

`func (o *BillingInvoice) HasAmountDue() bool`

HasAmountDue returns a boolean if a field has been set.

### GetAmountPaid

`func (o *BillingInvoice) GetAmountPaid() int64`

GetAmountPaid returns the AmountPaid field if non-nil, zero value otherwise.

### GetAmountPaidOk

`func (o *BillingInvoice) GetAmountPaidOk() (*int64, bool)`

GetAmountPaidOk returns a tuple with the AmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaid

`func (o *BillingInvoice) SetAmountPaid(v int64)`

SetAmountPaid sets AmountPaid field to given value.

### HasAmountPaid

`func (o *BillingInvoice) HasAmountPaid() bool`

HasAmountPaid returns a boolean if a field has been set.

### GetAttemptCount

`func (o *BillingInvoice) GetAttemptCount() int64`

GetAttemptCount returns the AttemptCount field if non-nil, zero value otherwise.

### GetAttemptCountOk

`func (o *BillingInvoice) GetAttemptCountOk() (*int64, bool)`

GetAttemptCountOk returns a tuple with the AttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCount

`func (o *BillingInvoice) SetAttemptCount(v int64)`

SetAttemptCount sets AttemptCount field to given value.

### HasAttemptCount

`func (o *BillingInvoice) HasAttemptCount() bool`

HasAttemptCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BillingInvoice) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingInvoice) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingInvoice) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BillingInvoice) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreditApplied

`func (o *BillingInvoice) GetCreditApplied() int64`

GetCreditApplied returns the CreditApplied field if non-nil, zero value otherwise.

### GetCreditAppliedOk

`func (o *BillingInvoice) GetCreditAppliedOk() (*int64, bool)`

GetCreditAppliedOk returns a tuple with the CreditApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditApplied

`func (o *BillingInvoice) SetCreditApplied(v int64)`

SetCreditApplied sets CreditApplied field to given value.

### HasCreditApplied

`func (o *BillingInvoice) HasCreditApplied() bool`

HasCreditApplied returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingInvoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingInvoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingInvoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingInvoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *BillingInvoice) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *BillingInvoice) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *BillingInvoice) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *BillingInvoice) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetDiscount

`func (o *BillingInvoice) GetDiscount() int64`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *BillingInvoice) GetDiscountOk() (*int64, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *BillingInvoice) SetDiscount(v int64)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *BillingInvoice) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDueDate

`func (o *BillingInvoice) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *BillingInvoice) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *BillingInvoice) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *BillingInvoice) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetId

`func (o *BillingInvoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingInvoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingInvoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingInvoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLineItems

`func (o *BillingInvoice) GetLineItems() []InvoiceLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *BillingInvoice) GetLineItemsOk() (*[]InvoiceLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *BillingInvoice) SetLineItems(v []InvoiceLineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *BillingInvoice) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetNumber

`func (o *BillingInvoice) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *BillingInvoice) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *BillingInvoice) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *BillingInvoice) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetNumberStr

`func (o *BillingInvoice) GetNumberStr() string`

GetNumberStr returns the NumberStr field if non-nil, zero value otherwise.

### GetNumberStrOk

`func (o *BillingInvoice) GetNumberStrOk() (*string, bool)`

GetNumberStrOk returns a tuple with the NumberStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberStr

`func (o *BillingInvoice) SetNumberStr(v string)`

SetNumberStr sets NumberStr field to given value.

### HasNumberStr

`func (o *BillingInvoice) HasNumberStr() bool`

HasNumberStr returns a boolean if a field has been set.

### GetPaidAt

`func (o *BillingInvoice) GetPaidAt() string`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *BillingInvoice) GetPaidAtOk() (*string, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *BillingInvoice) SetPaidAt(v string)`

SetPaidAt sets PaidAt field to given value.

### HasPaidAt

`func (o *BillingInvoice) HasPaidAt() bool`

HasPaidAt returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *BillingInvoice) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *BillingInvoice) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *BillingInvoice) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *BillingInvoice) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentRef

`func (o *BillingInvoice) GetPaymentRef() string`

GetPaymentRef returns the PaymentRef field if non-nil, zero value otherwise.

### GetPaymentRefOk

`func (o *BillingInvoice) GetPaymentRefOk() (*string, bool)`

GetPaymentRefOk returns a tuple with the PaymentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRef

`func (o *BillingInvoice) SetPaymentRef(v string)`

SetPaymentRef sets PaymentRef field to given value.

### HasPaymentRef

`func (o *BillingInvoice) HasPaymentRef() bool`

HasPaymentRef returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *BillingInvoice) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *BillingInvoice) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *BillingInvoice) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *BillingInvoice) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *BillingInvoice) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *BillingInvoice) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *BillingInvoice) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *BillingInvoice) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetStatus

`func (o *BillingInvoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingInvoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingInvoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingInvoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *BillingInvoice) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *BillingInvoice) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *BillingInvoice) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *BillingInvoice) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubtotal

`func (o *BillingInvoice) GetSubtotal() int64`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *BillingInvoice) GetSubtotalOk() (*int64, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *BillingInvoice) SetSubtotal(v int64)`

SetSubtotal sets Subtotal field to given value.

### HasSubtotal

`func (o *BillingInvoice) HasSubtotal() bool`

HasSubtotal returns a boolean if a field has been set.

### GetTax

`func (o *BillingInvoice) GetTax() int64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *BillingInvoice) GetTaxOk() (*int64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *BillingInvoice) SetTax(v int64)`

SetTax sets Tax field to given value.

### HasTax

`func (o *BillingInvoice) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BillingInvoice) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingInvoice) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingInvoice) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BillingInvoice) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *BillingInvoice) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BillingInvoice) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BillingInvoice) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BillingInvoice) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVoidedAt

`func (o *BillingInvoice) GetVoidedAt() string`

GetVoidedAt returns the VoidedAt field if non-nil, zero value otherwise.

### GetVoidedAtOk

`func (o *BillingInvoice) GetVoidedAtOk() (*string, bool)`

GetVoidedAtOk returns a tuple with the VoidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedAt

`func (o *BillingInvoice) SetVoidedAt(v string)`

SetVoidedAt sets VoidedAt field to given value.

### HasVoidedAt

`func (o *BillingInvoice) HasVoidedAt() bool`

HasVoidedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


