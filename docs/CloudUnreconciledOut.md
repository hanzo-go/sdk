# CloudUnreconciledOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Questions** | Pointer to [**[]CloudBankQuestion**](CloudBankQuestion.md) | Questions is the open clarifying question per unmatched inflow. | [optional] 
**Transactions** | Pointer to [**[]CloudBankTxnRow**](CloudBankTxnRow.md) | Transactions is every bank row still unmatched against the ledger. | [optional] 

## Methods

### NewCloudUnreconciledOut

`func NewCloudUnreconciledOut() *CloudUnreconciledOut`

NewCloudUnreconciledOut instantiates a new CloudUnreconciledOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudUnreconciledOutWithDefaults

`func NewCloudUnreconciledOutWithDefaults() *CloudUnreconciledOut`

NewCloudUnreconciledOutWithDefaults instantiates a new CloudUnreconciledOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestions

`func (o *CloudUnreconciledOut) GetQuestions() []CloudBankQuestion`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *CloudUnreconciledOut) GetQuestionsOk() (*[]CloudBankQuestion, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *CloudUnreconciledOut) SetQuestions(v []CloudBankQuestion)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *CloudUnreconciledOut) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### GetTransactions

`func (o *CloudUnreconciledOut) GetTransactions() []CloudBankTxnRow`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *CloudUnreconciledOut) GetTransactionsOk() (*[]CloudBankTxnRow, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *CloudUnreconciledOut) SetTransactions(v []CloudBankTxnRow)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *CloudUnreconciledOut) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


