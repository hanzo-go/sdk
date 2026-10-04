# AgentEventView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** | Actor is who produced the turn. A write that names nobody takes the calling principal, so this is rarely empty in practice. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when the turn was recorded, RFC 3339 in UTC to the second. Seconds are coarse enough that two turns can share one, which is why Seq and not this is the order. | [optional] 
**Id** | Pointer to **string** | ID is the event&#39;s own handle, minted as \&quot;evt_\&quot; + 32 hex characters. It identifies the turn; Seq is what ORDERS it. | [optional] 
**Kind** | Pointer to **string** | Kind is what the turn IS, from a closed six: message (a model turn), tool-call, spawn (a subagent started), log, status, control (a steering command the running surface consumes). Anything else is refused at the write. | [optional] 
**Payload** | Pointer to **interface{}** |  | [optional] 
**Seq** | Pointer to **int64** | Seq is the turn&#39;s position in this session&#39;s log: monotonic from 1, assigned by the store inside the insert, and unique PER SESSION rather than globally. It is the cursor a reader resumes from after a reconnect — ask for everything after your last-seen seq. | [optional] 
**SessionId** | Pointer to **string** | SessionID is the session this turn belongs to. Carried on every event so a stream frame stands alone — a subscriber watching a whole tree gets turns from several sessions down one connection. | [optional] 

## Methods

### NewAgentEventView

`func NewAgentEventView() *AgentEventView`

NewAgentEventView instantiates a new AgentEventView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentEventViewWithDefaults

`func NewAgentEventViewWithDefaults() *AgentEventView`

NewAgentEventViewWithDefaults instantiates a new AgentEventView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *AgentEventView) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AgentEventView) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AgentEventView) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *AgentEventView) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentEventView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentEventView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentEventView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentEventView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *AgentEventView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentEventView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentEventView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentEventView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *AgentEventView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AgentEventView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AgentEventView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AgentEventView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPayload

`func (o *AgentEventView) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *AgentEventView) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *AgentEventView) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *AgentEventView) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *AgentEventView) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *AgentEventView) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetSeq

`func (o *AgentEventView) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *AgentEventView) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *AgentEventView) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *AgentEventView) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetSessionId

`func (o *AgentEventView) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *AgentEventView) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *AgentEventView) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *AgentEventView) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


