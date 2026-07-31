# CloudPayoutView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**Settlement** | Pointer to **string** | Settlement discloses treasury-vs-wallet-vs-cash on every payout, to the author and to the admin mirror alike — the disclosure that keeps a first-party settlement legible as internal accounting. | [optional] 
**Txn** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudPayoutView

`func NewCloudPayoutView() *CloudPayoutView`

NewCloudPayoutView instantiates a new CloudPayoutView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPayoutViewWithDefaults

`func NewCloudPayoutViewWithDefaults() *CloudPayoutView`

NewCloudPayoutViewWithDefaults instantiates a new CloudPayoutView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *CloudPayoutView) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CloudPayoutView) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CloudPayoutView) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *CloudPayoutView) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudPayoutView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudPayoutView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudPayoutView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudPayoutView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CloudPayoutView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudPayoutView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudPayoutView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudPayoutView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *CloudPayoutView) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CloudPayoutView) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CloudPayoutView) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CloudPayoutView) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetReference

`func (o *CloudPayoutView) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CloudPayoutView) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CloudPayoutView) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CloudPayoutView) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSettlement

`func (o *CloudPayoutView) GetSettlement() string`

GetSettlement returns the Settlement field if non-nil, zero value otherwise.

### GetSettlementOk

`func (o *CloudPayoutView) GetSettlementOk() (*string, bool)`

GetSettlementOk returns a tuple with the Settlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlement

`func (o *CloudPayoutView) SetSettlement(v string)`

SetSettlement sets Settlement field to given value.

### HasSettlement

`func (o *CloudPayoutView) HasSettlement() bool`

HasSettlement returns a boolean if a field has been set.

### GetTxn

`func (o *CloudPayoutView) GetTxn() string`

GetTxn returns the Txn field if non-nil, zero value otherwise.

### GetTxnOk

`func (o *CloudPayoutView) GetTxnOk() (*string, bool)`

GetTxnOk returns a tuple with the Txn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxn

`func (o *CloudPayoutView) SetTxn(v string)`

SetTxn sets Txn field to given value.

### HasTxn

`func (o *CloudPayoutView) HasTxn() bool`

HasTxn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


