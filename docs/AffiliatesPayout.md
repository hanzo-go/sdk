# AffiliatesPayout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AmountCents** | Pointer to **int64** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**Txn** | Pointer to **string** | Commerce transaction id for a credits payout (absent for cash payouts). | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewAffiliatesPayout

`func NewAffiliatesPayout() *AffiliatesPayout`

NewAffiliatesPayout instantiates a new AffiliatesPayout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliatesPayoutWithDefaults

`func NewAffiliatesPayoutWithDefaults() *AffiliatesPayout`

NewAffiliatesPayoutWithDefaults instantiates a new AffiliatesPayout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AffiliatesPayout) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AffiliatesPayout) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AffiliatesPayout) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AffiliatesPayout) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAmountCents

`func (o *AffiliatesPayout) GetAmountCents() int64`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *AffiliatesPayout) GetAmountCentsOk() (*int64, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *AffiliatesPayout) SetAmountCents(v int64)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *AffiliatesPayout) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetMethod

`func (o *AffiliatesPayout) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AffiliatesPayout) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AffiliatesPayout) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *AffiliatesPayout) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetReference

`func (o *AffiliatesPayout) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AffiliatesPayout) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AffiliatesPayout) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AffiliatesPayout) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetTxn

`func (o *AffiliatesPayout) GetTxn() string`

GetTxn returns the Txn field if non-nil, zero value otherwise.

### GetTxnOk

`func (o *AffiliatesPayout) GetTxnOk() (*string, bool)`

GetTxnOk returns a tuple with the Txn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxn

`func (o *AffiliatesPayout) SetTxn(v string)`

SetTxn sets Txn field to given value.

### HasTxn

`func (o *AffiliatesPayout) HasTxn() bool`

HasTxn returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AffiliatesPayout) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AffiliatesPayout) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AffiliatesPayout) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AffiliatesPayout) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


