# BillingCharged

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceCents** | Pointer to **int64** | BalanceCents is the subject&#39;s balance AFTER the charge settled, in cents, so a caller does not have to re-read to show the new number. | [optional] 
**ProcessorRef** | Pointer to **string** | ProcessorRef is the payment processor&#39;s own reference. It is the only field that proves money moved at the GATEWAY rather than merely in our ledger, which is why it is answered and not only logged. Absent where the processor returned none. | [optional] 
**Status** | Pointer to **string** | Status is how the charge ended. Read it rather than inferring success from the HTTP status: the call succeeded whenever this field is present, and what the PROCESSOR did is what this says. | [optional] 
**Test** | Pointer to **bool** | Test states which bucket was credited — sandbox money or real money — so no reader has to guess whether a receipt is real. Sandbox and live funds are physically separate ledgers, and a reader that conflates them restates the company&#39;s revenue. | [optional] 
**TransactionId** | Pointer to **string** | TransactionID is the ledger entry this charge created. It is the handle a later read or a refund names, and it is minted by the ledger rather than by the caller. | [optional] 

## Methods

### NewBillingCharged

`func NewBillingCharged() *BillingCharged`

NewBillingCharged instantiates a new BillingCharged object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingChargedWithDefaults

`func NewBillingChargedWithDefaults() *BillingCharged`

NewBillingChargedWithDefaults instantiates a new BillingCharged object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceCents

`func (o *BillingCharged) GetBalanceCents() int64`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *BillingCharged) GetBalanceCentsOk() (*int64, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *BillingCharged) SetBalanceCents(v int64)`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *BillingCharged) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### GetProcessorRef

`func (o *BillingCharged) GetProcessorRef() string`

GetProcessorRef returns the ProcessorRef field if non-nil, zero value otherwise.

### GetProcessorRefOk

`func (o *BillingCharged) GetProcessorRefOk() (*string, bool)`

GetProcessorRefOk returns a tuple with the ProcessorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorRef

`func (o *BillingCharged) SetProcessorRef(v string)`

SetProcessorRef sets ProcessorRef field to given value.

### HasProcessorRef

`func (o *BillingCharged) HasProcessorRef() bool`

HasProcessorRef returns a boolean if a field has been set.

### GetStatus

`func (o *BillingCharged) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingCharged) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingCharged) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingCharged) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTest

`func (o *BillingCharged) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *BillingCharged) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *BillingCharged) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *BillingCharged) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetTransactionId

`func (o *BillingCharged) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *BillingCharged) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *BillingCharged) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *BillingCharged) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


