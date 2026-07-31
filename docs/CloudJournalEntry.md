# CloudJournalEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**Postings** | Pointer to [**[]CloudPosting**](CloudPosting.md) |  | [optional] 
**Program** | Pointer to **string** | referral|affiliate|author for payouts; \&quot;\&quot; otherwise | [optional] 
**Ref** | Pointer to **string** | idempotency ref (unique within Kind+Program) | [optional] 

## Methods

### NewCloudJournalEntry

`func NewCloudJournalEntry() *CloudJournalEntry`

NewCloudJournalEntry instantiates a new CloudJournalEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudJournalEntryWithDefaults

`func NewCloudJournalEntryWithDefaults() *CloudJournalEntry`

NewCloudJournalEntryWithDefaults instantiates a new CloudJournalEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CloudJournalEntry) GetAmount() interface{}`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CloudJournalEntry) GetAmountOk() (*interface{}, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CloudJournalEntry) SetAmount(v interface{})`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CloudJournalEntry) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *CloudJournalEntry) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *CloudJournalEntry) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCreatedAt

`func (o *CloudJournalEntry) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudJournalEntry) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudJournalEntry) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudJournalEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CloudJournalEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudJournalEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudJournalEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudJournalEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CloudJournalEntry) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudJournalEntry) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudJournalEntry) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudJournalEntry) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMemo

`func (o *CloudJournalEntry) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *CloudJournalEntry) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *CloudJournalEntry) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *CloudJournalEntry) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetPostings

`func (o *CloudJournalEntry) GetPostings() []CloudPosting`

GetPostings returns the Postings field if non-nil, zero value otherwise.

### GetPostingsOk

`func (o *CloudJournalEntry) GetPostingsOk() (*[]CloudPosting, bool)`

GetPostingsOk returns a tuple with the Postings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostings

`func (o *CloudJournalEntry) SetPostings(v []CloudPosting)`

SetPostings sets Postings field to given value.

### HasPostings

`func (o *CloudJournalEntry) HasPostings() bool`

HasPostings returns a boolean if a field has been set.

### GetProgram

`func (o *CloudJournalEntry) GetProgram() string`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *CloudJournalEntry) GetProgramOk() (*string, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *CloudJournalEntry) SetProgram(v string)`

SetProgram sets Program field to given value.

### HasProgram

`func (o *CloudJournalEntry) HasProgram() bool`

HasProgram returns a boolean if a field has been set.

### GetRef

`func (o *CloudJournalEntry) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CloudJournalEntry) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CloudJournalEntry) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CloudJournalEntry) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


