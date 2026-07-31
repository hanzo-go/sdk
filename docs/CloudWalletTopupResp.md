# CloudWalletTopupResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to **int32** | Balance is the org&#39;s new USD-ledger balance in cents. Best-effort: a read failure reports 0, and the credit has already landed either way. | [optional] 
**CreditedCents** | Pointer to **int32** | CreditedCents is the USD credit recorded, derived from the ON-CHAIN value using the token&#39;s own decimals — never a client-supplied number. | [optional] 
**Status** | Pointer to **string** | Status is how commerce recorded the payment. | [optional] 
**TxHash** | Pointer to **string** | TxHash is the transfer that was credited. | [optional] 

## Methods

### NewCloudWalletTopupResp

`func NewCloudWalletTopupResp() *CloudWalletTopupResp`

NewCloudWalletTopupResp instantiates a new CloudWalletTopupResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudWalletTopupRespWithDefaults

`func NewCloudWalletTopupRespWithDefaults() *CloudWalletTopupResp`

NewCloudWalletTopupRespWithDefaults instantiates a new CloudWalletTopupResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *CloudWalletTopupResp) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CloudWalletTopupResp) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CloudWalletTopupResp) SetBalance(v int32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *CloudWalletTopupResp) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCreditedCents

`func (o *CloudWalletTopupResp) GetCreditedCents() int32`

GetCreditedCents returns the CreditedCents field if non-nil, zero value otherwise.

### GetCreditedCentsOk

`func (o *CloudWalletTopupResp) GetCreditedCentsOk() (*int32, bool)`

GetCreditedCentsOk returns a tuple with the CreditedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditedCents

`func (o *CloudWalletTopupResp) SetCreditedCents(v int32)`

SetCreditedCents sets CreditedCents field to given value.

### HasCreditedCents

`func (o *CloudWalletTopupResp) HasCreditedCents() bool`

HasCreditedCents returns a boolean if a field has been set.

### GetStatus

`func (o *CloudWalletTopupResp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudWalletTopupResp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudWalletTopupResp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudWalletTopupResp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTxHash

`func (o *CloudWalletTopupResp) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *CloudWalletTopupResp) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *CloudWalletTopupResp) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *CloudWalletTopupResp) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


