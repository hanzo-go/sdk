# AgentBuildTurn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** | Actor is who took the turn. A deploy turn&#39;s actor is the literal \&quot;deploy\&quot;, because nobody took it. | [optional] 
**At** | Pointer to **string** | At is when the turn was recorded, RFC 3339 in UTC to the second. | [optional] 
**Body** | Pointer to **string** | Body is the readable text of the turn, taken from the stored event&#39;s &#x60;text&#x60;. Empty when the event carried a payload of some other shape — this route reads transcripts and does not invent prose for turns that are not one. | [optional] 
**Commit** | Pointer to **string** | Commit is the full sha this turn produced, empty when the turn changed nothing. It is ECHOED from the transcript, and the authority is the commit itself: it carries the &#x60;Hanzo-Session:&#x60;/&#x60;Hanzo-Turn:&#x60; trailer, or a note under refs/notes/hanzo-provenance saying the same, so the claim is checkable at source with the command in &#x60;verify&#x60;. | [optional] 
**Kind** | Pointer to **string** | Kind is what the turn was, from the log&#39;s closed six: message, tool-call, spawn, log, status, control. A deploy arrives as a &#x60;status&#x60; turn. | [optional] 
**Subject** | Pointer to **string** | Subject is that commit&#39;s subject line, from the same transcript, so a reader sees what the commit says without fetching the repository. | [optional] 
**Turn** | Pointer to **int64** | Seq is this turn&#39;s POSITION in the session&#39;s log — monotonic from 1, per session — and it is what a commit&#39;s &#x60;Hanzo-Turn:&#x60; trailer names. It is not a count of anything: the count is &#x60;turns&#x60; on the summary beside it. | [optional] 

## Methods

### NewAgentBuildTurn

`func NewAgentBuildTurn() *AgentBuildTurn`

NewAgentBuildTurn instantiates a new AgentBuildTurn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentBuildTurnWithDefaults

`func NewAgentBuildTurnWithDefaults() *AgentBuildTurn`

NewAgentBuildTurnWithDefaults instantiates a new AgentBuildTurn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *AgentBuildTurn) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AgentBuildTurn) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AgentBuildTurn) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *AgentBuildTurn) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAt

`func (o *AgentBuildTurn) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *AgentBuildTurn) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *AgentBuildTurn) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *AgentBuildTurn) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBody

`func (o *AgentBuildTurn) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *AgentBuildTurn) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *AgentBuildTurn) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *AgentBuildTurn) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCommit

`func (o *AgentBuildTurn) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *AgentBuildTurn) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *AgentBuildTurn) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *AgentBuildTurn) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetKind

`func (o *AgentBuildTurn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AgentBuildTurn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AgentBuildTurn) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AgentBuildTurn) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSubject

`func (o *AgentBuildTurn) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AgentBuildTurn) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AgentBuildTurn) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AgentBuildTurn) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTurn

`func (o *AgentBuildTurn) GetTurn() int64`

GetTurn returns the Turn field if non-nil, zero value otherwise.

### GetTurnOk

`func (o *AgentBuildTurn) GetTurnOk() (*int64, bool)`

GetTurnOk returns a tuple with the Turn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurn

`func (o *AgentBuildTurn) SetTurn(v int64)`

SetTurn sets Turn field to given value.

### HasTurn

`func (o *AgentBuildTurn) HasTurn() bool`

HasTurn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


