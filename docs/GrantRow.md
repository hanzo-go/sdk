# GrantRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** | staff email (or sub) who issued it | [optional] 
**AmountCents** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **string** | success | error | [optional] 
**Source** | Pointer to **string** | \&quot;trial\&quot; | \&quot;prepaid\&quot; | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 

## Methods

### NewGrantRow

`func NewGrantRow() *GrantRow`

NewGrantRow instantiates a new GrantRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantRowWithDefaults

`func NewGrantRowWithDefaults() *GrantRow`

NewGrantRowWithDefaults instantiates a new GrantRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *GrantRow) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *GrantRow) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *GrantRow) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *GrantRow) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAmountCents

`func (o *GrantRow) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *GrantRow) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *GrantRow) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *GrantRow) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GrantRow) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GrantRow) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GrantRow) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GrantRow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *GrantRow) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GrantRow) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GrantRow) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GrantRow) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetOrg

`func (o *GrantRow) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *GrantRow) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *GrantRow) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *GrantRow) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetReason

`func (o *GrantRow) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GrantRow) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GrantRow) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GrantRow) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetResult

`func (o *GrantRow) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GrantRow) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GrantRow) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *GrantRow) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSource

`func (o *GrantRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GrantRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GrantRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GrantRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTransactionId

`func (o *GrantRow) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *GrantRow) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *GrantRow) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *GrantRow) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


