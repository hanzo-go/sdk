# CommerceTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**DestinationId** | Pointer to **string** |  | [optional] 
**DestinationKind** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Test** | Pointer to **bool** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**SourceKind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewCommerceTransaction

`func NewCommerceTransaction() *CommerceTransaction`

NewCommerceTransaction instantiates a new CommerceTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceTransactionWithDefaults

`func NewCommerceTransactionWithDefaults() *CommerceTransaction`

NewCommerceTransactionWithDefaults instantiates a new CommerceTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommerceTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommerceTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommerceTransaction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommerceTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDestinationId

`func (o *CommerceTransaction) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *CommerceTransaction) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *CommerceTransaction) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *CommerceTransaction) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetDestinationKind

`func (o *CommerceTransaction) GetDestinationKind() string`

GetDestinationKind returns the DestinationKind field if non-nil, zero value otherwise.

### GetDestinationKindOk

`func (o *CommerceTransaction) GetDestinationKindOk() (*string, bool)`

GetDestinationKindOk returns a tuple with the DestinationKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationKind

`func (o *CommerceTransaction) SetDestinationKind(v string)`

SetDestinationKind sets DestinationKind field to given value.

### HasDestinationKind

`func (o *CommerceTransaction) HasDestinationKind() bool`

HasDestinationKind returns a boolean if a field has been set.

### GetCurrency

`func (o *CommerceTransaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CommerceTransaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CommerceTransaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CommerceTransaction) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmount

`func (o *CommerceTransaction) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CommerceTransaction) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CommerceTransaction) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CommerceTransaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetType

`func (o *CommerceTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommerceTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommerceTransaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CommerceTransaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTest

`func (o *CommerceTransaction) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *CommerceTransaction) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *CommerceTransaction) SetTest(v bool)`

SetTest sets Test field to given value.

### HasTest

`func (o *CommerceTransaction) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetNotes

`func (o *CommerceTransaction) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CommerceTransaction) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CommerceTransaction) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CommerceTransaction) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetSourceId

`func (o *CommerceTransaction) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CommerceTransaction) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CommerceTransaction) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *CommerceTransaction) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceKind

`func (o *CommerceTransaction) GetSourceKind() string`

GetSourceKind returns the SourceKind field if non-nil, zero value otherwise.

### GetSourceKindOk

`func (o *CommerceTransaction) GetSourceKindOk() (*string, bool)`

GetSourceKindOk returns a tuple with the SourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKind

`func (o *CommerceTransaction) SetSourceKind(v string)`

SetSourceKind sets SourceKind field to given value.

### HasSourceKind

`func (o *CommerceTransaction) HasSourceKind() bool`

HasSourceKind returns a boolean if a field has been set.

### GetMetadata

`func (o *CommerceTransaction) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CommerceTransaction) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CommerceTransaction) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CommerceTransaction) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommerceTransaction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommerceTransaction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommerceTransaction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommerceTransaction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CommerceTransaction) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CommerceTransaction) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CommerceTransaction) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CommerceTransaction) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


