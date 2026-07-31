# CloudBankTally

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ingested** | Pointer to **int32** | transactions seen | [optional] 
**Posted** | Pointer to **int32** | vouchers newly posted (outflow + reconciled) | [optional] 
**Questions** | Pointer to **int32** | unmatched inflows that raised a question | [optional] 
**Reconciled** | Pointer to **int32** | inflows cleared against Square-clearing | [optional] 
**Skipped** | Pointer to **int32** | already-processed idempotent no-ops | [optional] 
**Transfers** | Pointer to **int32** | own-account moves recorded (no P&amp;L) | [optional] 

## Methods

### NewCloudBankTally

`func NewCloudBankTally() *CloudBankTally`

NewCloudBankTally instantiates a new CloudBankTally object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBankTallyWithDefaults

`func NewCloudBankTallyWithDefaults() *CloudBankTally`

NewCloudBankTallyWithDefaults instantiates a new CloudBankTally object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngested

`func (o *CloudBankTally) GetIngested() int32`

GetIngested returns the Ingested field if non-nil, zero value otherwise.

### GetIngestedOk

`func (o *CloudBankTally) GetIngestedOk() (*int32, bool)`

GetIngestedOk returns a tuple with the Ingested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngested

`func (o *CloudBankTally) SetIngested(v int32)`

SetIngested sets Ingested field to given value.

### HasIngested

`func (o *CloudBankTally) HasIngested() bool`

HasIngested returns a boolean if a field has been set.

### GetPosted

`func (o *CloudBankTally) GetPosted() int32`

GetPosted returns the Posted field if non-nil, zero value otherwise.

### GetPostedOk

`func (o *CloudBankTally) GetPostedOk() (*int32, bool)`

GetPostedOk returns a tuple with the Posted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosted

`func (o *CloudBankTally) SetPosted(v int32)`

SetPosted sets Posted field to given value.

### HasPosted

`func (o *CloudBankTally) HasPosted() bool`

HasPosted returns a boolean if a field has been set.

### GetQuestions

`func (o *CloudBankTally) GetQuestions() int32`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *CloudBankTally) GetQuestionsOk() (*int32, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *CloudBankTally) SetQuestions(v int32)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *CloudBankTally) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### GetReconciled

`func (o *CloudBankTally) GetReconciled() int32`

GetReconciled returns the Reconciled field if non-nil, zero value otherwise.

### GetReconciledOk

`func (o *CloudBankTally) GetReconciledOk() (*int32, bool)`

GetReconciledOk returns a tuple with the Reconciled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciled

`func (o *CloudBankTally) SetReconciled(v int32)`

SetReconciled sets Reconciled field to given value.

### HasReconciled

`func (o *CloudBankTally) HasReconciled() bool`

HasReconciled returns a boolean if a field has been set.

### GetSkipped

`func (o *CloudBankTally) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *CloudBankTally) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *CloudBankTally) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *CloudBankTally) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetTransfers

`func (o *CloudBankTally) GetTransfers() int32`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *CloudBankTally) GetTransfersOk() (*int32, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *CloudBankTally) SetTransfers(v int32)`

SetTransfers sets Transfers field to given value.

### HasTransfers

`func (o *CloudBankTally) HasTransfers() bool`

HasTransfers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


