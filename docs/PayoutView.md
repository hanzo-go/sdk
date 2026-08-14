# PayoutView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** | AmountCents is the amount RESERVED against pending royalty, in integer USD cents, always positive. The reservation is atomic and can never exceed accrued − paid, so this is owed money moved out of pending — not money moved. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is unix seconds when the payout was RECORDED — the moment the amount left pending, not the moment a human moved the money. | [optional] 
**Id** | Pointer to **string** | ID is the payout row&#39;s server-minted handle, \&quot;apo_\&quot;-prefixed. A caller never supplies it; it is what an operator quotes when reconciling a settlement. | [optional] 
**Method** | Pointer to **string** | Method is how the operator says this settles, lowercased as recorded. \&quot;credits\&quot; is the one method that means the author&#39;s own wallet; anything else — wire, paypal, check — is a cash disbursement a human performs. Recording it pays nobody either way. | [optional] 
**Reference** | Pointer to **string** | Reference is the operator&#39;s external handle for the settlement: a wire confirmation, a PayPal transaction id. Absent when none was given. | [optional] 
**Settlement** | Pointer to **string** | Settlement discloses treasury-vs-wallet-vs-cash on every payout, to the author and to the admin mirror alike — the disclosure that keeps a first-party settlement legible as internal accounting. | [optional] 
**Txn** | Pointer to **string** | Txn is the commerce ledger transaction id of a SETTLED credits payout, and it is absent on every payout this service records. Recording moves no money, and authors asks the money plane exactly one question — what has this org spent? — with no write to answer it with, so there is no receipt to carry. It fills in only when a settlement stamps its transaction back onto the row. | [optional] 

## Methods

### NewPayoutView

`func NewPayoutView() *PayoutView`

NewPayoutView instantiates a new PayoutView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutViewWithDefaults

`func NewPayoutViewWithDefaults() *PayoutView`

NewPayoutViewWithDefaults instantiates a new PayoutView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *PayoutView) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *PayoutView) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *PayoutView) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *PayoutView) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PayoutView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PayoutView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PayoutView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PayoutView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *PayoutView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PayoutView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PayoutView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PayoutView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *PayoutView) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PayoutView) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PayoutView) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PayoutView) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetReference

`func (o *PayoutView) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PayoutView) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PayoutView) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PayoutView) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSettlement

`func (o *PayoutView) GetSettlement() string`

GetSettlement returns the Settlement field if non-nil, zero value otherwise.

### GetSettlementOk

`func (o *PayoutView) GetSettlementOk() (*string, bool)`

GetSettlementOk returns a tuple with the Settlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlement

`func (o *PayoutView) SetSettlement(v string)`

SetSettlement sets Settlement field to given value.

### HasSettlement

`func (o *PayoutView) HasSettlement() bool`

HasSettlement returns a boolean if a field has been set.

### GetTxn

`func (o *PayoutView) GetTxn() string`

GetTxn returns the Txn field if non-nil, zero value otherwise.

### GetTxnOk

`func (o *PayoutView) GetTxnOk() (*string, bool)`

GetTxnOk returns a tuple with the Txn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxn

`func (o *PayoutView) SetTxn(v string)`

SetTxn sets Txn field to given value.

### HasTxn

`func (o *PayoutView) HasTxn() bool`

HasTxn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


