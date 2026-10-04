# BillingCollected

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceUsedCents** | Pointer to **int64** | BalanceUsedCents is how much was covered by prepaid balance. | [optional] 
**CardChargedCents** | Pointer to **int64** | CardChargedCents is how much was charged to the card on file. | [optional] 
**CreditUsedCents** | Pointer to **int64** | CreditUsedCents is how much was covered by credit grants. | [optional] 
**Invoice** | Pointer to [**BillingInvoice**](BillingInvoice.md) | Invoice is the invoice AFTER the attempt — its status is the authority on what happened, not this struct&#39;s other fields. | [optional] 
**Paid** | Pointer to **bool** | Paid reports whether the invoice is now settled in full. A false here with no error is a DECLINE: the invoice stays open and may be collected again. | [optional] 
**ProcessorRef** | Pointer to **string** | ProcessorRef is the processor&#39;s reference for any card charge — the field that proves money moved at the gateway rather than only in our ledger. | [optional] 
**Reason** | Pointer to **string** | Reason explains a decline or partial collection. Empty on success. | [optional] 

## Methods

### NewBillingCollected

`func NewBillingCollected() *BillingCollected`

NewBillingCollected instantiates a new BillingCollected object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCollectedWithDefaults

`func NewBillingCollectedWithDefaults() *BillingCollected`

NewBillingCollectedWithDefaults instantiates a new BillingCollected object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceUsedCents

`func (o *BillingCollected) GetBalanceUsedCents() int64`

GetBalanceUsedCents returns the BalanceUsedCents field if non-nil, zero value otherwise.

### GetBalanceUsedCentsOk

`func (o *BillingCollected) GetBalanceUsedCentsOk() (*int64, bool)`

GetBalanceUsedCentsOk returns a tuple with the BalanceUsedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceUsedCents

`func (o *BillingCollected) SetBalanceUsedCents(v int64)`

SetBalanceUsedCents sets BalanceUsedCents field to given value.

### HasBalanceUsedCents

`func (o *BillingCollected) HasBalanceUsedCents() bool`

HasBalanceUsedCents returns a boolean if a field has been set.

### GetCardChargedCents

`func (o *BillingCollected) GetCardChargedCents() int64`

GetCardChargedCents returns the CardChargedCents field if non-nil, zero value otherwise.

### GetCardChargedCentsOk

`func (o *BillingCollected) GetCardChargedCentsOk() (*int64, bool)`

GetCardChargedCentsOk returns a tuple with the CardChargedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardChargedCents

`func (o *BillingCollected) SetCardChargedCents(v int64)`

SetCardChargedCents sets CardChargedCents field to given value.

### HasCardChargedCents

`func (o *BillingCollected) HasCardChargedCents() bool`

HasCardChargedCents returns a boolean if a field has been set.

### GetCreditUsedCents

`func (o *BillingCollected) GetCreditUsedCents() int64`

GetCreditUsedCents returns the CreditUsedCents field if non-nil, zero value otherwise.

### GetCreditUsedCentsOk

`func (o *BillingCollected) GetCreditUsedCentsOk() (*int64, bool)`

GetCreditUsedCentsOk returns a tuple with the CreditUsedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditUsedCents

`func (o *BillingCollected) SetCreditUsedCents(v int64)`

SetCreditUsedCents sets CreditUsedCents field to given value.

### HasCreditUsedCents

`func (o *BillingCollected) HasCreditUsedCents() bool`

HasCreditUsedCents returns a boolean if a field has been set.

### GetInvoice

`func (o *BillingCollected) GetInvoice() BillingInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *BillingCollected) GetInvoiceOk() (*BillingInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *BillingCollected) SetInvoice(v BillingInvoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *BillingCollected) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetPaid

`func (o *BillingCollected) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *BillingCollected) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *BillingCollected) SetPaid(v bool)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *BillingCollected) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetProcessorRef

`func (o *BillingCollected) GetProcessorRef() string`

GetProcessorRef returns the ProcessorRef field if non-nil, zero value otherwise.

### GetProcessorRefOk

`func (o *BillingCollected) GetProcessorRefOk() (*string, bool)`

GetProcessorRefOk returns a tuple with the ProcessorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorRef

`func (o *BillingCollected) SetProcessorRef(v string)`

SetProcessorRef sets ProcessorRef field to given value.

### HasProcessorRef

`func (o *BillingCollected) HasProcessorRef() bool`

HasProcessorRef returns a boolean if a field has been set.

### GetReason

`func (o *BillingCollected) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BillingCollected) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BillingCollected) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BillingCollected) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


