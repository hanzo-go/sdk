# BillingTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingTransaction

`func NewBillingTransaction() *BillingTransaction`

NewBillingTransaction instantiates a new BillingTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingTransactionWithDefaults

`func NewBillingTransactionWithDefaults() *BillingTransaction`

NewBillingTransactionWithDefaults instantiates a new BillingTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BillingTransaction) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingTransaction) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingTransaction) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingTransaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BillingTransaction) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingTransaction) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingTransaction) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BillingTransaction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingTransaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingTransaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingTransaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingTransaction) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExpiresAt

`func (o *BillingTransaction) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *BillingTransaction) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *BillingTransaction) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *BillingTransaction) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *BillingTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingTransaction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *BillingTransaction) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BillingTransaction) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BillingTransaction) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BillingTransaction) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *BillingTransaction) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BillingTransaction) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetNotes

`func (o *BillingTransaction) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BillingTransaction) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BillingTransaction) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BillingTransaction) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTags

`func (o *BillingTransaction) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BillingTransaction) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BillingTransaction) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BillingTransaction) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *BillingTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingTransaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BillingTransaction) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


