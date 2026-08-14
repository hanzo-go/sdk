# UnreconciledOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Questions** | Pointer to [**[]BankQuestion**](BankQuestion.md) | Questions is the open clarifying question per unmatched inflow. | [optional] 
**Transactions** | Pointer to [**[]BankTxnRow**](BankTxnRow.md) | Transactions is every bank row still unmatched against the ledger. | [optional] 

## Methods

### NewUnreconciledOut

`func NewUnreconciledOut() *UnreconciledOut`

NewUnreconciledOut instantiates a new UnreconciledOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnreconciledOutWithDefaults

`func NewUnreconciledOutWithDefaults() *UnreconciledOut`

NewUnreconciledOutWithDefaults instantiates a new UnreconciledOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestions

`func (o *UnreconciledOut) GetQuestions() []BankQuestion`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *UnreconciledOut) GetQuestionsOk() (*[]BankQuestion, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *UnreconciledOut) SetQuestions(v []BankQuestion)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *UnreconciledOut) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### GetTransactions

`func (o *UnreconciledOut) GetTransactions() []BankTxnRow`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *UnreconciledOut) GetTransactionsOk() (*[]BankTxnRow, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *UnreconciledOut) SetTransactions(v []BankTxnRow)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *UnreconciledOut) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


