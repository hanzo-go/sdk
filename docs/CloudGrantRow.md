# CloudGrantRow

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

### NewCloudGrantRow

`func NewCloudGrantRow() *CloudGrantRow`

NewCloudGrantRow instantiates a new CloudGrantRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGrantRowWithDefaults

`func NewCloudGrantRowWithDefaults() *CloudGrantRow`

NewCloudGrantRowWithDefaults instantiates a new CloudGrantRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *CloudGrantRow) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *CloudGrantRow) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *CloudGrantRow) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *CloudGrantRow) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAmountCents

`func (o *CloudGrantRow) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CloudGrantRow) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CloudGrantRow) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *CloudGrantRow) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudGrantRow) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudGrantRow) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudGrantRow) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudGrantRow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *CloudGrantRow) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudGrantRow) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudGrantRow) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudGrantRow) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetOrg

`func (o *CloudGrantRow) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudGrantRow) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudGrantRow) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudGrantRow) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetReason

`func (o *CloudGrantRow) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CloudGrantRow) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CloudGrantRow) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CloudGrantRow) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetResult

`func (o *CloudGrantRow) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CloudGrantRow) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CloudGrantRow) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *CloudGrantRow) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSource

`func (o *CloudGrantRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudGrantRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudGrantRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudGrantRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTransactionId

`func (o *CloudGrantRow) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CloudGrantRow) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CloudGrantRow) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *CloudGrantRow) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


