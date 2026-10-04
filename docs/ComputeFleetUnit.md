# ComputeFleetUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host is the unit&#39;s hostname. Empty for a unit that is not one host: a cluster row has no hostname to report. | [optional] 
**Kind** | Pointer to **string** | Kind is what the unit IS: laptop, cloud, gpu, cluster, machine or worker. | [optional] 
**Label** | Pointer to **string** | Label is the name to show a human — a target&#39;s label, a worker&#39;s hostname, a machine&#39;s display name. Empty when the source has none to give. | [optional] 
**Metrics** | Pointer to [**ComputeFleetMetrics**](ComputeFleetMetrics.md) | Metrics is the unit&#39;s latest utilization: its own live snapshot when it keeps one (a run-target&#39;s heartbeat wins), else the newest sample from the series for the SAME source. Absent means nothing is known about this unit&#39;s load — which is deliberately not the same as a reading of zero. | [optional] 
**Queued** | Pointer to **int64** | Queued is how many renders are waiting on THIS GPU&#39;s own lane in the org&#39;s gpu-jobs queue. BYO units only — an agent run-target dispatches, it does not queue — and omitted when nothing is waiting. | [optional] 
**Running** | Pointer to **int64** | Running is what the unit is executing right now: agent sessions in flight for a run-target, claimed renders for a BYO GPU. | [optional] 
**Sessions** | Pointer to **int64** | Sessions is how many agent sessions are open on this unit. Always present, and 0 for a source that cannot host agent sessions at all — a fact about that plane, not a gap in the reading. | [optional] 
**Source** | Pointer to **string** | Source is the plane this row came from: \&quot;agent\&quot; (a linked run-target), \&quot;byo\&quot; (a worker or cluster the org dialed in) or \&quot;visor\&quot; (a machine Hanzo provisioned). It is half the row&#39;s identity, and it says which face owns the unit — /v1/agent/targets, /v1/compute/fleet/workers, /v1/compute/machines. | [optional] 
**Spec** | Pointer to [**ComputeFleetSpec**](ComputeFleetSpec.md) | Spec is the unit&#39;s static capability. Absent when the source reported none — unknown capability, never a zeroed one. | [optional] 
**Status** | Pointer to **string** | Status is liveness in the SOURCE&#39;s own vocabulary, because each plane decides it differently: a run-target&#39;s is derived from its heartbeat, a BYO worker&#39;s is online/offline on the 90s window, a BYO cluster&#39;s is \&quot;attached\&quot;, and a Visor machine&#39;s is the provider&#39;s word for its lifecycle state. | [optional] 
**Unit** | Pointer to **string** | Unit is the SOURCE&#39;s own id for this unit — a run-target id, a BYO worker id, a Visor machine name — so a row links straight back to the face that owns it. It is unique within a source, not across them: two planes may mint the same id, which is why (source, unit) together is the identity. | [optional] 

## Methods

### NewComputeFleetUnit

`func NewComputeFleetUnit() *ComputeFleetUnit`

NewComputeFleetUnit instantiates a new ComputeFleetUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeFleetUnitWithDefaults

`func NewComputeFleetUnitWithDefaults() *ComputeFleetUnit`

NewComputeFleetUnitWithDefaults instantiates a new ComputeFleetUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ComputeFleetUnit) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ComputeFleetUnit) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ComputeFleetUnit) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ComputeFleetUnit) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKind

`func (o *ComputeFleetUnit) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ComputeFleetUnit) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ComputeFleetUnit) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ComputeFleetUnit) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *ComputeFleetUnit) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ComputeFleetUnit) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ComputeFleetUnit) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ComputeFleetUnit) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMetrics

`func (o *ComputeFleetUnit) GetMetrics() ComputeFleetMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ComputeFleetUnit) GetMetricsOk() (*ComputeFleetMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ComputeFleetUnit) SetMetrics(v ComputeFleetMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *ComputeFleetUnit) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetQueued

`func (o *ComputeFleetUnit) GetQueued() int64`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *ComputeFleetUnit) GetQueuedOk() (*int64, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *ComputeFleetUnit) SetQueued(v int64)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *ComputeFleetUnit) HasQueued() bool`

HasQueued returns a boolean if a field has been set.

### GetRunning

`func (o *ComputeFleetUnit) GetRunning() int64`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *ComputeFleetUnit) GetRunningOk() (*int64, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *ComputeFleetUnit) SetRunning(v int64)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *ComputeFleetUnit) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetSessions

`func (o *ComputeFleetUnit) GetSessions() int64`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *ComputeFleetUnit) GetSessionsOk() (*int64, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *ComputeFleetUnit) SetSessions(v int64)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *ComputeFleetUnit) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetSource

`func (o *ComputeFleetUnit) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ComputeFleetUnit) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ComputeFleetUnit) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ComputeFleetUnit) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSpec

`func (o *ComputeFleetUnit) GetSpec() ComputeFleetSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ComputeFleetUnit) GetSpecOk() (*ComputeFleetSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ComputeFleetUnit) SetSpec(v ComputeFleetSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ComputeFleetUnit) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *ComputeFleetUnit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComputeFleetUnit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComputeFleetUnit) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComputeFleetUnit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUnit

`func (o *ComputeFleetUnit) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ComputeFleetUnit) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ComputeFleetUnit) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ComputeFleetUnit) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


