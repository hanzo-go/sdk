# BooksBankTally

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ingested** | Pointer to **int64** | transactions seen | [optional] 
**Posted** | Pointer to **int64** | vouchers newly posted (outflow + reconciled) | [optional] 
**Questions** | Pointer to **int64** | unmatched inflows that raised a question | [optional] 
**Reconciled** | Pointer to **int64** | inflows cleared against Square-clearing | [optional] 
**Skipped** | Pointer to **int64** | already-processed idempotent no-ops | [optional] 
**Transfers** | Pointer to **int64** | own-account moves recorded (no P&amp;L) | [optional] 

## Methods

### NewBooksBankTally

`func NewBooksBankTally() *BooksBankTally`

NewBooksBankTally instantiates a new BooksBankTally object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksBankTallyWithDefaults

`func NewBooksBankTallyWithDefaults() *BooksBankTally`

NewBooksBankTallyWithDefaults instantiates a new BooksBankTally object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngested

`func (o *BooksBankTally) GetIngested() int64`

GetIngested returns the Ingested field if non-nil, zero value otherwise.

### GetIngestedOk

`func (o *BooksBankTally) GetIngestedOk() (*int64, bool)`

GetIngestedOk returns a tuple with the Ingested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngested

`func (o *BooksBankTally) SetIngested(v int64)`

SetIngested sets Ingested field to given value.

### HasIngested

`func (o *BooksBankTally) HasIngested() bool`

HasIngested returns a boolean if a field has been set.

### GetPosted

`func (o *BooksBankTally) GetPosted() int64`

GetPosted returns the Posted field if non-nil, zero value otherwise.

### GetPostedOk

`func (o *BooksBankTally) GetPostedOk() (*int64, bool)`

GetPostedOk returns a tuple with the Posted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosted

`func (o *BooksBankTally) SetPosted(v int64)`

SetPosted sets Posted field to given value.

### HasPosted

`func (o *BooksBankTally) HasPosted() bool`

HasPosted returns a boolean if a field has been set.

### GetQuestions

`func (o *BooksBankTally) GetQuestions() int64`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *BooksBankTally) GetQuestionsOk() (*int64, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *BooksBankTally) SetQuestions(v int64)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *BooksBankTally) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### GetReconciled

`func (o *BooksBankTally) GetReconciled() int64`

GetReconciled returns the Reconciled field if non-nil, zero value otherwise.

### GetReconciledOk

`func (o *BooksBankTally) GetReconciledOk() (*int64, bool)`

GetReconciledOk returns a tuple with the Reconciled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciled

`func (o *BooksBankTally) SetReconciled(v int64)`

SetReconciled sets Reconciled field to given value.

### HasReconciled

`func (o *BooksBankTally) HasReconciled() bool`

HasReconciled returns a boolean if a field has been set.

### GetSkipped

`func (o *BooksBankTally) GetSkipped() int64`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *BooksBankTally) GetSkippedOk() (*int64, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *BooksBankTally) SetSkipped(v int64)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *BooksBankTally) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetTransfers

`func (o *BooksBankTally) GetTransfers() int64`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *BooksBankTally) GetTransfersOk() (*int64, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *BooksBankTally) SetTransfers(v int64)`

SetTransfers sets Transfers field to given value.

### HasTransfers

`func (o *BooksBankTally) HasTransfers() bool`

HasTransfers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


