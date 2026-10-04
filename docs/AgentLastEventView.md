# AgentLastEventView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** | Actor is who produced the turn, defaulted to the calling principal when the writer named nobody. | [optional] 
**At** | Pointer to **string** | At is when the turn was recorded, RFC 3339 in UTC to the second. | [optional] 
**Kind** | Pointer to **string** | Kind is what the turn was, from the log&#39;s closed six: message, tool-call, spawn, log, status, control. | [optional] 
**Preview** | Pointer to **string** | Preview is the first 240 bytes of the event&#39;s payload, cut without regard for the JSON inside it — it is a string to SHOW, never a value to parse. Read the detail or the stream for the whole payload. | [optional] 
**Seq** | Pointer to **int64** | Seq is that event&#39;s position in the session&#39;s log — monotonic from 1, per session. A reader holding it can ask the detail or stream reads for everything after it, so this doubles as the list&#39;s resume cursor. | [optional] 

## Methods

### NewAgentLastEventView

`func NewAgentLastEventView() *AgentLastEventView`

NewAgentLastEventView instantiates a new AgentLastEventView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentLastEventViewWithDefaults

`func NewAgentLastEventViewWithDefaults() *AgentLastEventView`

NewAgentLastEventViewWithDefaults instantiates a new AgentLastEventView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *AgentLastEventView) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AgentLastEventView) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AgentLastEventView) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *AgentLastEventView) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAt

`func (o *AgentLastEventView) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *AgentLastEventView) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *AgentLastEventView) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *AgentLastEventView) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetKind

`func (o *AgentLastEventView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AgentLastEventView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AgentLastEventView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AgentLastEventView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPreview

`func (o *AgentLastEventView) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *AgentLastEventView) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *AgentLastEventView) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *AgentLastEventView) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetSeq

`func (o *AgentLastEventView) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *AgentLastEventView) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *AgentLastEventView) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *AgentLastEventView) HasSeq() bool`

HasSeq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


