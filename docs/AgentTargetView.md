# AgentTargetView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **string** | Capacity is a human summary of what the machine has (\&quot;8 vCPU / 32G\&quot;, \&quot;1× GB10\&quot;), up to 256 characters. Prose for a card — Spec is the same thing in a form a scheduler can read, and nothing derives one from the other. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when the machine was first registered, RFC 3339 in UTC. A re-link refreshes the row and leaves this alone, so it dates the machine and not the connection. | [optional] 
**Host** | Pointer to **string** | Host is the hostname sessions on this machine report, and it is a JOIN KEY, not a label: a session naming this host counts against the load below even when it names no target id, and a re-link of the same (org, host, owner) refreshes this row instead of creating a second. Empty means the machine is addressable only by ID. | [optional] 
**Id** | Pointer to **string** | ID is the machine&#39;s handle, minted as \&quot;tgt_\&quot; + 32 hex characters. It is what a session records to say it ran here, and what every later patch, claim or delete addresses. | [optional] 
**Kind** | Pointer to **string** | Kind is what sort of destination this is, from a closed five: laptop | cloud | gpu | cluster | machine. A register that named none is a &#x60;machine&#x60;. | [optional] 
**Label** | Pointer to **string** | Label is the name a person gave the machine (\&quot;workshop\&quot;), up to 128 characters. Required at register, free text, and the only field here meant for reading rather than matching. | [optional] 
**Metrics** | Pointer to [**AgentMetrics**](AgentMetrics.md) | Metrics is what the machine was DOING at its last heartbeat — loadavg, memory, accelerator utilization. Absent when it has never beaten. It is a SNAPSHOT: the series over time lives in the fleet samples, not here. | [optional] 
**MetricsAt** | Pointer to **string** | MetricsAt is when that heartbeat was recorded, RFC 3339 in UTC, and the SERVER stamps it — a client cannot backdate or forge the staleness clock. Absent means never beaten, which is exactly the case where Status is taken at its word. | [optional] 
**Running** | Pointer to **int64** | Running is how many of those are in &#x60;running&#x60; right now — the number a dispatcher weighs against Capacity. paused sessions are in Sessions and not here. | [optional] 
**Sessions** | Pointer to **int64** | Sessions is how many of the org&#39;s sessions are mapped to this machine, by target id OR by matching Host. All of them, whatever their status. | [optional] 
**Spec** | Pointer to [**AgentSpec**](AgentSpec.md) | Spec is what the machine IS — os, arch, cores, RAM, accelerators — the static half, changed only when something reports it again. Absent when nothing has ever been reported, and a scheduler reads absence as \&quot;cannot satisfy a floor\&quot; rather than as \&quot;no limits\&quot;. | [optional] 
**Status** | Pointer to **string** | Status is the EFFECTIVE liveness — online | offline | draining — not the stored one. offline and draining are operator INTENT and are reported as they stand; &#x60;online&#x60; is checked against the heartbeat, and a machine that has beaten before but not in the last 90 seconds reports offline whatever its row says. A target that has NEVER beaten keeps its stored status, because a hand-registered destination has no fact to check. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is the last write to the row, same format — which for a beating machine is its last heartbeat, since a heartbeat IS a write. | [optional] 

## Methods

### NewAgentTargetView

`func NewAgentTargetView() *AgentTargetView`

NewAgentTargetView instantiates a new AgentTargetView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTargetViewWithDefaults

`func NewAgentTargetViewWithDefaults() *AgentTargetView`

NewAgentTargetViewWithDefaults instantiates a new AgentTargetView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *AgentTargetView) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *AgentTargetView) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *AgentTargetView) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *AgentTargetView) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentTargetView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentTargetView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentTargetView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentTargetView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHost

`func (o *AgentTargetView) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AgentTargetView) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AgentTargetView) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AgentTargetView) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *AgentTargetView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentTargetView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentTargetView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentTargetView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *AgentTargetView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AgentTargetView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AgentTargetView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AgentTargetView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *AgentTargetView) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AgentTargetView) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AgentTargetView) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AgentTargetView) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMetrics

`func (o *AgentTargetView) GetMetrics() AgentMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *AgentTargetView) GetMetricsOk() (*AgentMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *AgentTargetView) SetMetrics(v AgentMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *AgentTargetView) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMetricsAt

`func (o *AgentTargetView) GetMetricsAt() string`

GetMetricsAt returns the MetricsAt field if non-nil, zero value otherwise.

### GetMetricsAtOk

`func (o *AgentTargetView) GetMetricsAtOk() (*string, bool)`

GetMetricsAtOk returns a tuple with the MetricsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsAt

`func (o *AgentTargetView) SetMetricsAt(v string)`

SetMetricsAt sets MetricsAt field to given value.

### HasMetricsAt

`func (o *AgentTargetView) HasMetricsAt() bool`

HasMetricsAt returns a boolean if a field has been set.

### GetRunning

`func (o *AgentTargetView) GetRunning() int64`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *AgentTargetView) GetRunningOk() (*int64, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *AgentTargetView) SetRunning(v int64)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *AgentTargetView) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetSessions

`func (o *AgentTargetView) GetSessions() int64`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *AgentTargetView) GetSessionsOk() (*int64, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *AgentTargetView) SetSessions(v int64)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *AgentTargetView) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetSpec

`func (o *AgentTargetView) GetSpec() AgentSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AgentTargetView) GetSpecOk() (*AgentSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AgentTargetView) SetSpec(v AgentSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AgentTargetView) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *AgentTargetView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentTargetView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentTargetView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentTargetView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AgentTargetView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentTargetView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentTargetView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AgentTargetView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


