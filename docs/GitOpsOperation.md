# GitOpsOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinishedAt** | Pointer to **string** | FinishedAt is when it ended, RFC 3339. Absent while the phase is Running. | [optional] 
**Message** | Pointer to **string** | Message is CD&#39;s account of the phase — \&quot;successfully synced (all tasks run)\&quot; for a Succeeded operation, the reason it stopped for a Failed one. | [optional] 
**Phase** | Pointer to **string** | Phase is how the last sync operation ended, in CD&#39;s own vocabulary: Running, Succeeded or Failed. It is never empty — an Application whose phase is empty has no operation at all and omits this whole object. | [optional] 
**Revision** | Pointer to **string** | Revision is the commit this operation ATTEMPTED (operationState.syncResult). It differs from the Application&#39;s own revision exactly when the attempt did not land: revision is the last commit CD got applied, this is the last one it tried. | [optional] 
**StartedAt** | Pointer to **string** | StartedAt is when the operation began, RFC 3339. | [optional] 

## Methods

### NewGitOpsOperation

`func NewGitOpsOperation() *GitOpsOperation`

NewGitOpsOperation instantiates a new GitOpsOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitOpsOperationWithDefaults

`func NewGitOpsOperationWithDefaults() *GitOpsOperation`

NewGitOpsOperationWithDefaults instantiates a new GitOpsOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinishedAt

`func (o *GitOpsOperation) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *GitOpsOperation) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *GitOpsOperation) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *GitOpsOperation) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetMessage

`func (o *GitOpsOperation) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GitOpsOperation) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GitOpsOperation) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GitOpsOperation) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPhase

`func (o *GitOpsOperation) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *GitOpsOperation) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *GitOpsOperation) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *GitOpsOperation) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetRevision

`func (o *GitOpsOperation) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *GitOpsOperation) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *GitOpsOperation) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *GitOpsOperation) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetStartedAt

`func (o *GitOpsOperation) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GitOpsOperation) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GitOpsOperation) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GitOpsOperation) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


