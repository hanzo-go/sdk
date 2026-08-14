# PaymentOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceCents** | Pointer to **int32** | BalanceCents is the org&#39;s balance AFTER this payment, read back from the same key just credited so it matches what the balance endpoint reports. | [optional] 
**Id** | Pointer to **string** | ID is the ledger transaction id for the credit. It is what getPayment reads back, and the customer-visible receipt for the money. | [optional] 
**ProcessorRef** | Pointer to **string** | ProcessorRef is the payment processor&#39;s own reference for the charge (Square&#39;s payment id). It is the field that proves money actually moved at the gateway rather than only in our ledger — the thing to quote when reconciling against a processor dashboard. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; on a settled charge. A charge that did not settle is an error with the processor&#39;s reason, never a status field to inspect. | [optional] 
**Test** | Pointer to **bool** | Test reports which bucket this credited: true is a SANDBOX charge crediting the test balance, false is live money. It is always stated so a receipt can never be mistaken for the other kind. | [optional] 

## Methods

### NewPaymentOut

`func NewPaymentOut() *PaymentOut`

NewPaymentOut instantiates a new PaymentOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentOutWithDefaults

`func NewPaymentOutWithDefaults() *PaymentOut`

NewPaymentOutWithDefaults instantiates a new PaymentOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceCents

`func (o *PaymentOut) GetBalanceCents() int32`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *PaymentOut) GetBalanceCentsOk() (*int32, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *PaymentOut) SetBalanceCents(v int32)`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *PaymentOut) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.

### GetId

`func (o *PaymentOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentOut) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentOut) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProcessorRef

`func (o *PaymentOut) GetProcessorRef() string`

GetProcessorRef returns the ProcessorRef field if non-nil, zero value otherwise.

### GetProcessorRefOk

`func (o *PaymentOut) GetProcessorRefOk() (*string, bool)`

GetProcessorRefOk returns a tuple with the ProcessorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorRef

`func (o *PaymentOut) SetProcessorRef(v string)`

SetProcessorRef sets ProcessorRef field to given value.

### HasProcessorRef

`func (o *PaymentOut) HasProcessorRef() bool`

HasProcessorRef returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTest

`func (o *PaymentOut) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *PaymentOut) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *PaymentOut) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *PaymentOut) HasTest() bool`

HasTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


