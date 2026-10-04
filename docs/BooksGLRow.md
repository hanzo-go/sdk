# BooksGLRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the chart-of-accounts number this leg posts to. | [optional] 
**Against** | Pointer to **string** | Against names the OTHER accounts in the same voucher — the contra side of this leg — so a single row reads as an entry rather than as half of one. | [optional] 
**Credit** | Pointer to **int64** | Credit is the amount credited to that account, in whole cents. | [optional] 
**Debit** | Pointer to **int64** | Debit is the amount debited to that account, in whole cents. Exactly one of debit and credit is non-zero on a leg; a negative amount is never used to mean the other side. | [optional] 
**Id** | Pointer to **int64** | ID is the entry&#39;s position in the ledger. The ledger is append-only, so ids ascend with posting order and a higher id is a later entry. | [optional] 
**PostingAt** | Pointer to **string** | PostingAt is the accounting date this entry belongs to — what the reports window on, which need not be when the row was written. | [optional] 
**Remarks** | Pointer to **string** | Remarks is the memo carried onto the entry, for a human reading the ledger. | [optional] 
**SourceId** | Pointer to **string** | SourceID identifies that originating record within its kind. | [optional] 
**SourceKind** | Pointer to **string** | SourceKind is what caused the posting: a bank line, a scanned document, a commerce sale. With sourceId it traces the entry back to the thing that produced it. | [optional] 

## Methods

### NewBooksGLRow

`func NewBooksGLRow() *BooksGLRow`

NewBooksGLRow instantiates a new BooksGLRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksGLRowWithDefaults

`func NewBooksGLRowWithDefaults() *BooksGLRow`

NewBooksGLRowWithDefaults instantiates a new BooksGLRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *BooksGLRow) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *BooksGLRow) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *BooksGLRow) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *BooksGLRow) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAgainst

`func (o *BooksGLRow) GetAgainst() string`

GetAgainst returns the Against field if non-nil, zero value otherwise.

### GetAgainstOk

`func (o *BooksGLRow) GetAgainstOk() (*string, bool)`

GetAgainstOk returns a tuple with the Against field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgainst

`func (o *BooksGLRow) SetAgainst(v string)`

SetAgainst sets Against field to given value.

### HasAgainst

`func (o *BooksGLRow) HasAgainst() bool`

HasAgainst returns a boolean if a field has been set.

### GetCredit

`func (o *BooksGLRow) GetCredit() int64`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *BooksGLRow) GetCreditOk() (*int64, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *BooksGLRow) SetCredit(v int64)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *BooksGLRow) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetDebit

`func (o *BooksGLRow) GetDebit() int64`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *BooksGLRow) GetDebitOk() (*int64, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *BooksGLRow) SetDebit(v int64)`

SetDebit sets Debit field to given value.

### HasDebit

`func (o *BooksGLRow) HasDebit() bool`

HasDebit returns a boolean if a field has been set.

### GetId

`func (o *BooksGLRow) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BooksGLRow) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BooksGLRow) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BooksGLRow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPostingAt

`func (o *BooksGLRow) GetPostingAt() string`

GetPostingAt returns the PostingAt field if non-nil, zero value otherwise.

### GetPostingAtOk

`func (o *BooksGLRow) GetPostingAtOk() (*string, bool)`

GetPostingAtOk returns a tuple with the PostingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingAt

`func (o *BooksGLRow) SetPostingAt(v string)`

SetPostingAt sets PostingAt field to given value.

### HasPostingAt

`func (o *BooksGLRow) HasPostingAt() bool`

HasPostingAt returns a boolean if a field has been set.

### GetRemarks

`func (o *BooksGLRow) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *BooksGLRow) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *BooksGLRow) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *BooksGLRow) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetSourceId

`func (o *BooksGLRow) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BooksGLRow) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BooksGLRow) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *BooksGLRow) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceKind

`func (o *BooksGLRow) GetSourceKind() string`

GetSourceKind returns the SourceKind field if non-nil, zero value otherwise.

### GetSourceKindOk

`func (o *BooksGLRow) GetSourceKindOk() (*string, bool)`

GetSourceKindOk returns a tuple with the SourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKind

`func (o *BooksGLRow) SetSourceKind(v string)`

SetSourceKind sets SourceKind field to given value.

### HasSourceKind

`func (o *BooksGLRow) HasSourceKind() bool`

HasSourceKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


