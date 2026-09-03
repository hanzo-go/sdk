# Remittance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int64** | AmountCents is the amount disbursed, in cents. It was reserved against pending commission atomically when recorded, so it never exceeds what was owed. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the payout was recorded, Unix seconds UTC — when the balance moved, not necessarily when the cash landed. | [optional] 
**Id** | Pointer to **string** | ID is the payout row&#39;s server-minted handle, \&quot;apo_\&quot;-prefixed. | [optional] 
**Method** | Pointer to **string** | Method is how it was settled. \&quot;credits\&quot; issued a commerce grant into the affiliate org&#39;s own wallet; any other value (wire, paypal, check, …) is a RECORD of cash a human moved out of band. | [optional] 
**Reference** | Pointer to **string** | Reference is the operator&#39;s settlement note — a bank id, a ledger ref. Free text, absent when none was given. | [optional] 
**Txn** | Pointer to **string** | Txn is the commerce ledger transaction id, set ONLY where a \&quot;credits\&quot; payout actually issued the grant. Absent for cash methods, which write no ledger row. | [optional] 

## Methods

### NewRemittance

`func NewRemittance() *Remittance`

NewRemittance instantiates a new Remittance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemittanceWithDefaults

`func NewRemittanceWithDefaults() *Remittance`

NewRemittanceWithDefaults instantiates a new Remittance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *Remittance) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *Remittance) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *Remittance) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *Remittance) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Remittance) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Remittance) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Remittance) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Remittance) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *Remittance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Remittance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Remittance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Remittance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *Remittance) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Remittance) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Remittance) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Remittance) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetReference

`func (o *Remittance) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Remittance) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Remittance) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *Remittance) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetTxn

`func (o *Remittance) GetTxn() string`

GetTxn returns the Txn field if non-nil, zero value otherwise.

### GetTxnOk

`func (o *Remittance) GetTxnOk() (*string, bool)`

GetTxnOk returns a tuple with the Txn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxn

`func (o *Remittance) SetTxn(v string)`

SetTxn sets Txn field to given value.

### HasTxn

`func (o *Remittance) HasTxn() bool`

HasTxn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


