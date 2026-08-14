# JournalEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**Postings** | Pointer to [**[]Posting**](Posting.md) |  | [optional] 
**Program** | Pointer to **string** | referral|affiliate|author for payouts; \&quot;\&quot; otherwise | [optional] 
**Ref** | Pointer to **string** | idempotency ref (unique within Kind+Program) | [optional] 

## Methods

### NewJournalEntry

`func NewJournalEntry() *JournalEntry`

NewJournalEntry instantiates a new JournalEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalEntryWithDefaults

`func NewJournalEntryWithDefaults() *JournalEntry`

NewJournalEntryWithDefaults instantiates a new JournalEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *JournalEntry) GetAmount() interface{}`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *JournalEntry) GetAmountOk() (*interface{}, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *JournalEntry) SetAmount(v interface{})`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *JournalEntry) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *JournalEntry) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *JournalEntry) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCreatedAt

`func (o *JournalEntry) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *JournalEntry) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *JournalEntry) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *JournalEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *JournalEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JournalEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JournalEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JournalEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *JournalEntry) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *JournalEntry) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *JournalEntry) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *JournalEntry) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMemo

`func (o *JournalEntry) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *JournalEntry) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *JournalEntry) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *JournalEntry) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetPostings

`func (o *JournalEntry) GetPostings() []Posting`

GetPostings returns the Postings field if non-nil, zero value otherwise.

### GetPostingsOk

`func (o *JournalEntry) GetPostingsOk() (*[]Posting, bool)`

GetPostingsOk returns a tuple with the Postings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostings

`func (o *JournalEntry) SetPostings(v []Posting)`

SetPostings sets Postings field to given value.

### HasPostings

`func (o *JournalEntry) HasPostings() bool`

HasPostings returns a boolean if a field has been set.

### GetProgram

`func (o *JournalEntry) GetProgram() string`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *JournalEntry) GetProgramOk() (*string, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *JournalEntry) SetProgram(v string)`

SetProgram sets Program field to given value.

### HasProgram

`func (o *JournalEntry) HasProgram() bool`

HasProgram returns a boolean if a field has been set.

### GetRef

`func (o *JournalEntry) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *JournalEntry) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *JournalEntry) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *JournalEntry) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


