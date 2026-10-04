# BooksUnreconciledOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Questions** | Pointer to [**[]BooksBankQuestion**](BooksBankQuestion.md) | Questions is the open clarifying question per unmatched inflow. | [optional] 
**Transactions** | Pointer to [**[]BooksBankTxnRow**](BooksBankTxnRow.md) | Transactions is every bank row still unmatched against the ledger. | [optional] 

## Methods

### NewBooksUnreconciledOut

`func NewBooksUnreconciledOut() *BooksUnreconciledOut`

NewBooksUnreconciledOut instantiates a new BooksUnreconciledOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksUnreconciledOutWithDefaults

`func NewBooksUnreconciledOutWithDefaults() *BooksUnreconciledOut`

NewBooksUnreconciledOutWithDefaults instantiates a new BooksUnreconciledOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestions

`func (o *BooksUnreconciledOut) GetQuestions() []BooksBankQuestion`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *BooksUnreconciledOut) GetQuestionsOk() (*[]BooksBankQuestion, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *BooksUnreconciledOut) SetQuestions(v []BooksBankQuestion)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *BooksUnreconciledOut) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### GetTransactions

`func (o *BooksUnreconciledOut) GetTransactions() []BooksBankTxnRow`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *BooksUnreconciledOut) GetTransactionsOk() (*[]BooksBankTxnRow, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *BooksUnreconciledOut) SetTransactions(v []BooksBankTxnRow)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *BooksUnreconciledOut) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


