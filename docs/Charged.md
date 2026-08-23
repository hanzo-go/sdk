# Charged

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceCents** | Pointer to **int32** | BalanceCents is the subject&#39;s balance AFTER the charge settled, in cents, so a caller does not have to re-read to show the new number. | [optional] 
**ProcessorRef** | Pointer to **string** | ProcessorRef is the payment processor&#39;s own reference. It is the only field that proves money moved at the GATEWAY rather than merely in our ledger, which is why it is answered and not only logged. Absent where the processor returned none. | [optional] 
**Status** | Pointer to **string** | Status is how the charge ended. Read it rather than inferring success from the HTTP status: the call succeeded whenever this field is present, and what the PROCESSOR did is what this says. | [optional] 
**Test** | Pointer to **bool** | Test states which bucket was credited — sandbox money or real money — so no reader has to guess whether a receipt is real. Sandbox and live funds are physically separate ledgers, and a reader that conflates them restates the company&#39;s revenue. | [optional] 
**TransactionId** | Pointer to **string** | TransactionID is the ledger entry this charge created. It is the handle a later read or a refund names, and it is minted by the ledger rather than by the caller. | [optional] 

## Methods

### NewCharged

`func NewCharged() *Charged`

NewCharged instantiates a new Charged object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargedWithDefaults

`func NewChargedWithDefaults() *Charged`

NewChargedWithDefaults instantiates a new Charged object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceCents

`func (o *Charged) GetBalanceCents() int32`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *Charged) GetBalanceCentsOk() (*int32, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *Charged) SetBalanceCents(v int32)`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *Charged) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### GetProcessorRef

`func (o *Charged) GetProcessorRef() string`

GetProcessorRef returns the ProcessorRef field if non-nil, zero value otherwise.

### GetProcessorRefOk

`func (o *Charged) GetProcessorRefOk() (*string, bool)`

GetProcessorRefOk returns a tuple with the ProcessorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorRef

`func (o *Charged) SetProcessorRef(v string)`

SetProcessorRef sets ProcessorRef field to given value.

### HasProcessorRef

`func (o *Charged) HasProcessorRef() bool`

HasProcessorRef returns a boolean if a field has been set.

### GetStatus

`func (o *Charged) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Charged) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Charged) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Charged) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTest

`func (o *Charged) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *Charged) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *Charged) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *Charged) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetTransactionId

`func (o *Charged) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Charged) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Charged) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *Charged) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


