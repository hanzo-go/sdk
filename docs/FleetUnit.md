# FleetUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host is the unit&#39;s hostname. Empty for a unit that is not one host: a cluster row has no hostname to report. | [optional] 
**Kind** | Pointer to **string** | Kind is what the unit IS: laptop, cloud, gpu, cluster, machine or worker. | [optional] 
**Label** | Pointer to **string** | Label is the name to show a human — a target&#39;s label, a worker&#39;s hostname, a machine&#39;s display name. Empty when the source has none to give. | [optional] 
**Metrics** | Pointer to [**FleetMetrics**](FleetMetrics.md) | Metrics is the unit&#39;s latest utilization: its own live snapshot when it keeps one (a run-target&#39;s heartbeat wins), else the newest sample from the series for the SAME source. Absent means nothing is known about this unit&#39;s load — which is deliberately not the same as a reading of zero. | [optional] 
**Queued** | Pointer to **int64** | Queued is how many renders are waiting on THIS GPU&#39;s own lane in the org&#39;s gpu-jobs queue. BYO units only — an agent run-target dispatches, it does not queue — and omitted when nothing is waiting. | [optional] 
**Running** | Pointer to **int64** | Running is what the unit is executing right now: agent sessions in flight for a run-target, claimed renders for a BYO GPU. | [optional] 
**Sessions** | Pointer to **int64** | Sessions is how many agent sessions are open on this unit. Always present, and 0 for a source that cannot host agent sessions at all — a fact about that plane, not a gap in the reading. | [optional] 
**Source** | Pointer to **string** | Source is the plane this row came from: \&quot;agent\&quot; (a linked run-target), \&quot;byo\&quot; (a worker or cluster the org dialed in) or \&quot;visor\&quot; (a machine Hanzo provisioned). It is half the row&#39;s identity, and it says which face owns the unit — /v1/agent/targets, /v1/compute/fleet/workers, /v1/compute/machines. | [optional] 
**Spec** | Pointer to [**FleetSpec**](FleetSpec.md) | Spec is the unit&#39;s static capability. Absent when the source reported none — unknown capability, never a zeroed one. | [optional] 
**Status** | Pointer to **string** | Status is liveness in the SOURCE&#39;s own vocabulary, because each plane decides it differently: a run-target&#39;s is derived from its heartbeat, a BYO worker&#39;s is online/offline on the 90s window, a BYO cluster&#39;s is \&quot;attached\&quot;, and a Visor machine&#39;s is the provider&#39;s word for its lifecycle state. | [optional] 
**Unit** | Pointer to **string** | Unit is the SOURCE&#39;s own id for this unit — a run-target id, a BYO worker id, a Visor machine name — so a row links straight back to the face that owns it. It is unique within a source, not across them: two planes may mint the same id, which is why (source, unit) together is the identity. | [optional] 

## Methods

### NewFleetUnit

`func NewFleetUnit() *FleetUnit`

NewFleetUnit instantiates a new FleetUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetUnitWithDefaults

`func NewFleetUnitWithDefaults() *FleetUnit`

NewFleetUnitWithDefaults instantiates a new FleetUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *FleetUnit) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *FleetUnit) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *FleetUnit) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *FleetUnit) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKind

`func (o *FleetUnit) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FleetUnit) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FleetUnit) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FleetUnit) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *FleetUnit) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FleetUnit) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FleetUnit) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FleetUnit) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMetrics

`func (o *FleetUnit) GetMetrics() FleetMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *FleetUnit) GetMetricsOk() (*FleetMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *FleetUnit) SetMetrics(v FleetMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *FleetUnit) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetQueued

`func (o *FleetUnit) GetQueued() int64`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *FleetUnit) GetQueuedOk() (*int64, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *FleetUnit) SetQueued(v int64)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *FleetUnit) HasQueued() bool`

HasQueued returns a boolean if a field has been set.

### GetRunning

`func (o *FleetUnit) GetRunning() int64`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *FleetUnit) GetRunningOk() (*int64, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *FleetUnit) SetRunning(v int64)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *FleetUnit) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetSessions

`func (o *FleetUnit) GetSessions() int64`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *FleetUnit) GetSessionsOk() (*int64, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *FleetUnit) SetSessions(v int64)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *FleetUnit) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetSource

`func (o *FleetUnit) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FleetUnit) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FleetUnit) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *FleetUnit) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSpec

`func (o *FleetUnit) GetSpec() FleetSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *FleetUnit) GetSpecOk() (*FleetSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *FleetUnit) SetSpec(v FleetSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *FleetUnit) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *FleetUnit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FleetUnit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FleetUnit) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FleetUnit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUnit

`func (o *FleetUnit) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *FleetUnit) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *FleetUnit) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *FleetUnit) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


