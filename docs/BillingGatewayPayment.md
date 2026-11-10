# BillingGatewayPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Credited** | Pointer to **bool** |  | [optional] 
**Deferred** | Pointer to **bool** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Payment** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TxHash** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingGatewayPayment

`func NewBillingGatewayPayment() *BillingGatewayPayment`

NewBillingGatewayPayment instantiates a new BillingGatewayPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingGatewayPaymentWithDefaults

`func NewBillingGatewayPaymentWithDefaults() *BillingGatewayPayment`

NewBillingGatewayPaymentWithDefaults instantiates a new BillingGatewayPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *BillingGatewayPayment) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *BillingGatewayPayment) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *BillingGatewayPayment) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *BillingGatewayPayment) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BillingGatewayPayment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingGatewayPayment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingGatewayPayment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BillingGatewayPayment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredited

`func (o *BillingGatewayPayment) GetCredited() bool`

GetCredited returns the Credited field if non-nil, zero value otherwise.

### GetCreditedOk

`func (o *BillingGatewayPayment) GetCreditedOk() (*bool, bool)`

GetCreditedOk returns a tuple with the Credited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredited

`func (o *BillingGatewayPayment) SetCredited(v bool)`

SetCredited sets Credited field to given value.

### HasCredited

`func (o *BillingGatewayPayment) HasCredited() bool`

HasCredited returns a boolean if a field has been set.

### GetDeferred

`func (o *BillingGatewayPayment) GetDeferred() bool`

GetDeferred returns the Deferred field if non-nil, zero value otherwise.

### GetDeferredOk

`func (o *BillingGatewayPayment) GetDeferredOk() (*bool, bool)`

GetDeferredOk returns a tuple with the Deferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferred

`func (o *BillingGatewayPayment) SetDeferred(v bool)`

SetDeferred sets Deferred field to given value.

### HasDeferred

`func (o *BillingGatewayPayment) HasDeferred() bool`

HasDeferred returns a boolean if a field has been set.

### GetFrom

`func (o *BillingGatewayPayment) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BillingGatewayPayment) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BillingGatewayPayment) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *BillingGatewayPayment) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *BillingGatewayPayment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingGatewayPayment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingGatewayPayment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingGatewayPayment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetwork

`func (o *BillingGatewayPayment) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *BillingGatewayPayment) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *BillingGatewayPayment) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *BillingGatewayPayment) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPayment

`func (o *BillingGatewayPayment) GetPayment() string`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *BillingGatewayPayment) GetPaymentOk() (*string, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *BillingGatewayPayment) SetPayment(v string)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *BillingGatewayPayment) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetReason

`func (o *BillingGatewayPayment) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BillingGatewayPayment) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BillingGatewayPayment) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BillingGatewayPayment) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *BillingGatewayPayment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingGatewayPayment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingGatewayPayment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingGatewayPayment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTxHash

`func (o *BillingGatewayPayment) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *BillingGatewayPayment) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *BillingGatewayPayment) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *BillingGatewayPayment) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BillingGatewayPayment) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingGatewayPayment) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingGatewayPayment) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BillingGatewayPayment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValue

`func (o *BillingGatewayPayment) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BillingGatewayPayment) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BillingGatewayPayment) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BillingGatewayPayment) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


