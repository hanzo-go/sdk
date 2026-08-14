# CollectOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceUsedCents** | Pointer to **int32** | BalanceUsedCents is how much was covered by prepaid balance. | [optional] 
**CardChargedCents** | Pointer to **int32** | CardChargedCents is how much was charged to the card on file. | [optional] 
**CreditUsedCents** | Pointer to **int32** | CreditUsedCents is how much was covered by credit grants. | [optional] 
**Invoice** | Pointer to [**InvoiceOut**](InvoiceOut.md) | Invoice is the invoice AFTER the attempt — its status is the authority on what happened, not this struct&#39;s other fields. | [optional] 
**Paid** | Pointer to **bool** | Paid reports whether the invoice is now settled in full. A false here with no error is a DECLINE: the invoice stays open and may be collected again. | [optional] 
**ProcessorRef** | Pointer to **string** | ProcessorRef is the processor&#39;s reference for any card charge — the field that proves money moved at the gateway rather than only in our ledger. | [optional] 
**Reason** | Pointer to **string** | Reason explains a decline or partial collection. Empty on success. | [optional] 

## Methods

### NewCollectOut

`func NewCollectOut() *CollectOut`

NewCollectOut instantiates a new CollectOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectOutWithDefaults

`func NewCollectOutWithDefaults() *CollectOut`

NewCollectOutWithDefaults instantiates a new CollectOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceUsedCents

`func (o *CollectOut) GetBalanceUsedCents() int32`

GetBalanceUsedCents returns the BalanceUsedCents field if non-nil, zero value otherwise.

### GetBalanceUsedCentsOk

`func (o *CollectOut) GetBalanceUsedCentsOk() (*int32, bool)`

GetBalanceUsedCentsOk returns a tuple with the BalanceUsedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceUsedCents

`func (o *CollectOut) SetBalanceUsedCents(v int32)`

SetBalanceUsedCents sets BalanceUsedCents field to given value.

### HasBalanceUsedCents

`func (o *CollectOut) HasBalanceUsedCents() bool`

HasBalanceUsedCents returns a boolean if a field has been set.

### GetCardChargedCents

`func (o *CollectOut) GetCardChargedCents() int32`

GetCardChargedCents returns the CardChargedCents field if non-nil, zero value otherwise.

### GetCardChargedCentsOk

`func (o *CollectOut) GetCardChargedCentsOk() (*int32, bool)`

GetCardChargedCentsOk returns a tuple with the CardChargedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardChargedCents

`func (o *CollectOut) SetCardChargedCents(v int32)`

SetCardChargedCents sets CardChargedCents field to given value.

### HasCardChargedCents

`func (o *CollectOut) HasCardChargedCents() bool`

HasCardChargedCents returns a boolean if a field has been set.

### GetCreditUsedCents

`func (o *CollectOut) GetCreditUsedCents() int32`

GetCreditUsedCents returns the CreditUsedCents field if non-nil, zero value otherwise.

### GetCreditUsedCentsOk

`func (o *CollectOut) GetCreditUsedCentsOk() (*int32, bool)`

GetCreditUsedCentsOk returns a tuple with the CreditUsedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditUsedCents

`func (o *CollectOut) SetCreditUsedCents(v int32)`

SetCreditUsedCents sets CreditUsedCents field to given value.

### HasCreditUsedCents

`func (o *CollectOut) HasCreditUsedCents() bool`

HasCreditUsedCents returns a boolean if a field has been set.

### GetInvoice

`func (o *CollectOut) GetInvoice() InvoiceOut`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *CollectOut) GetInvoiceOk() (*InvoiceOut, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *CollectOut) SetInvoice(v InvoiceOut)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *CollectOut) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetPaid

`func (o *CollectOut) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *CollectOut) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *CollectOut) SetPaid(v bool)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *CollectOut) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetProcessorRef

`func (o *CollectOut) GetProcessorRef() string`

GetProcessorRef returns the ProcessorRef field if non-nil, zero value otherwise.

### GetProcessorRefOk

`func (o *CollectOut) GetProcessorRefOk() (*string, bool)`

GetProcessorRefOk returns a tuple with the ProcessorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorRef

`func (o *CollectOut) SetProcessorRef(v string)`

SetProcessorRef sets ProcessorRef field to given value.

### HasProcessorRef

`func (o *CollectOut) HasProcessorRef() bool`

HasProcessorRef returns a boolean if a field has been set.

### GetReason

`func (o *CollectOut) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CollectOut) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CollectOut) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CollectOut) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


