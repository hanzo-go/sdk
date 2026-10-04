# BillingBillingInvoice

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
**LineItems** | Pointer to [**[]BillingInvoiceLineItem**](BillingInvoiceLineItem.md) | LineItems carries no omitempty and is never allocated empty, because the wire it reproduces sends &#x60;null&#x60; for an invoice with no lines. An empty array there would be a different answer to \&quot;were there lines\&quot;. | [optional] 
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

### NewBillingBillingInvoice

`func NewBillingBillingInvoice() *BillingBillingInvoice`

NewBillingBillingInvoice instantiates a new BillingBillingInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingBillingInvoiceWithDefaults

`func NewBillingBillingInvoiceWithDefaults() *BillingBillingInvoice`

NewBillingBillingInvoiceWithDefaults instantiates a new BillingBillingInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountDue

`func (o *BillingBillingInvoice) GetAmountDue() int64`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *BillingBillingInvoice) GetAmountDueOk() (*int64, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *BillingBillingInvoice) SetAmountDue(v int64)`

SetAmountDue sets AmountDue field to given value.

### HasAmountDue

`func (o *BillingBillingInvoice) HasAmountDue() bool`

HasAmountDue returns a boolean if a field has been set.

### GetAmountPaid

`func (o *BillingBillingInvoice) GetAmountPaid() int64`

GetAmountPaid returns the AmountPaid field if non-nil, zero value otherwise.

### GetAmountPaidOk

`func (o *BillingBillingInvoice) GetAmountPaidOk() (*int64, bool)`

GetAmountPaidOk returns a tuple with the AmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaid

`func (o *BillingBillingInvoice) SetAmountPaid(v int64)`

SetAmountPaid sets AmountPaid field to given value.

### HasAmountPaid

`func (o *BillingBillingInvoice) HasAmountPaid() bool`

HasAmountPaid returns a boolean if a field has been set.

### GetAttemptCount

`func (o *BillingBillingInvoice) GetAttemptCount() int64`

GetAttemptCount returns the AttemptCount field if non-nil, zero value otherwise.

### GetAttemptCountOk

`func (o *BillingBillingInvoice) GetAttemptCountOk() (*int64, bool)`

GetAttemptCountOk returns a tuple with the AttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCount

`func (o *BillingBillingInvoice) SetAttemptCount(v int64)`

SetAttemptCount sets AttemptCount field to given value.

### HasAttemptCount

`func (o *BillingBillingInvoice) HasAttemptCount() bool`

HasAttemptCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BillingBillingInvoice) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingBillingInvoice) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingBillingInvoice) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BillingBillingInvoice) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreditApplied

`func (o *BillingBillingInvoice) GetCreditApplied() int64`

GetCreditApplied returns the CreditApplied field if non-nil, zero value otherwise.

### GetCreditAppliedOk

`func (o *BillingBillingInvoice) GetCreditAppliedOk() (*int64, bool)`

GetCreditAppliedOk returns a tuple with the CreditApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditApplied

`func (o *BillingBillingInvoice) SetCreditApplied(v int64)`

SetCreditApplied sets CreditApplied field to given value.

### HasCreditApplied

`func (o *BillingBillingInvoice) HasCreditApplied() bool`

HasCreditApplied returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingBillingInvoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingBillingInvoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingBillingInvoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingBillingInvoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *BillingBillingInvoice) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *BillingBillingInvoice) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *BillingBillingInvoice) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *BillingBillingInvoice) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetDiscount

`func (o *BillingBillingInvoice) GetDiscount() int64`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *BillingBillingInvoice) GetDiscountOk() (*int64, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *BillingBillingInvoice) SetDiscount(v int64)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *BillingBillingInvoice) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDueDate

`func (o *BillingBillingInvoice) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *BillingBillingInvoice) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *BillingBillingInvoice) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *BillingBillingInvoice) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetId

`func (o *BillingBillingInvoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingBillingInvoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingBillingInvoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingBillingInvoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLineItems

`func (o *BillingBillingInvoice) GetLineItems() []BillingInvoiceLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *BillingBillingInvoice) GetLineItemsOk() (*[]BillingInvoiceLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *BillingBillingInvoice) SetLineItems(v []BillingInvoiceLineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *BillingBillingInvoice) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetNumber

`func (o *BillingBillingInvoice) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *BillingBillingInvoice) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *BillingBillingInvoice) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *BillingBillingInvoice) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetNumberStr

`func (o *BillingBillingInvoice) GetNumberStr() string`

GetNumberStr returns the NumberStr field if non-nil, zero value otherwise.

### GetNumberStrOk

`func (o *BillingBillingInvoice) GetNumberStrOk() (*string, bool)`

GetNumberStrOk returns a tuple with the NumberStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberStr

`func (o *BillingBillingInvoice) SetNumberStr(v string)`

SetNumberStr sets NumberStr field to given value.

### HasNumberStr

`func (o *BillingBillingInvoice) HasNumberStr() bool`

HasNumberStr returns a boolean if a field has been set.

### GetPaidAt

`func (o *BillingBillingInvoice) GetPaidAt() string`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *BillingBillingInvoice) GetPaidAtOk() (*string, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *BillingBillingInvoice) SetPaidAt(v string)`

SetPaidAt sets PaidAt field to given value.

### HasPaidAt

`func (o *BillingBillingInvoice) HasPaidAt() bool`

HasPaidAt returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *BillingBillingInvoice) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *BillingBillingInvoice) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *BillingBillingInvoice) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *BillingBillingInvoice) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentRef

`func (o *BillingBillingInvoice) GetPaymentRef() string`

GetPaymentRef returns the PaymentRef field if non-nil, zero value otherwise.

### GetPaymentRefOk

`func (o *BillingBillingInvoice) GetPaymentRefOk() (*string, bool)`

GetPaymentRefOk returns a tuple with the PaymentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRef

`func (o *BillingBillingInvoice) SetPaymentRef(v string)`

SetPaymentRef sets PaymentRef field to given value.

### HasPaymentRef

`func (o *BillingBillingInvoice) HasPaymentRef() bool`

HasPaymentRef returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *BillingBillingInvoice) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *BillingBillingInvoice) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *BillingBillingInvoice) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *BillingBillingInvoice) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *BillingBillingInvoice) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *BillingBillingInvoice) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *BillingBillingInvoice) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *BillingBillingInvoice) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetStatus

`func (o *BillingBillingInvoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingBillingInvoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingBillingInvoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingBillingInvoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *BillingBillingInvoice) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *BillingBillingInvoice) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *BillingBillingInvoice) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *BillingBillingInvoice) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubtotal

`func (o *BillingBillingInvoice) GetSubtotal() int64`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *BillingBillingInvoice) GetSubtotalOk() (*int64, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *BillingBillingInvoice) SetSubtotal(v int64)`

SetSubtotal sets Subtotal field to given value.

### HasSubtotal

`func (o *BillingBillingInvoice) HasSubtotal() bool`

HasSubtotal returns a boolean if a field has been set.

### GetTax

`func (o *BillingBillingInvoice) GetTax() int64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *BillingBillingInvoice) GetTaxOk() (*int64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *BillingBillingInvoice) SetTax(v int64)`

SetTax sets Tax field to given value.

### HasTax

`func (o *BillingBillingInvoice) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BillingBillingInvoice) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingBillingInvoice) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingBillingInvoice) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BillingBillingInvoice) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *BillingBillingInvoice) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BillingBillingInvoice) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BillingBillingInvoice) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BillingBillingInvoice) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVoidedAt

`func (o *BillingBillingInvoice) GetVoidedAt() string`

GetVoidedAt returns the VoidedAt field if non-nil, zero value otherwise.

### GetVoidedAtOk

`func (o *BillingBillingInvoice) GetVoidedAtOk() (*string, bool)`

GetVoidedAtOk returns a tuple with the VoidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedAt

`func (o *BillingBillingInvoice) SetVoidedAt(v string)`

SetVoidedAt sets VoidedAt field to given value.

### HasVoidedAt

`func (o *BillingBillingInvoice) HasVoidedAt() bool`

HasVoidedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


