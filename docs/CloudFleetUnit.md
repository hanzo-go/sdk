# CloudFleetUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to [**CloudFleetMetrics**](CloudFleetMetrics.md) |  | [optional] 
**Queued** | Pointer to **int32** |  | [optional] 
**Running** | Pointer to **int32** | Running is what the unit is actively executing: agent sessions for a run-target, in-flight renders for a BYO GPU. Queued is the gpu-jobs backlog on this GPU&#39;s lane (BYO units only; an agent unit does not queue). Both come from the org&#39;s gpu-jobs queue for BYO units, overlaid in listFleet. | [optional] 
**Sessions** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Spec** | Pointer to [**CloudFleetSpec**](CloudFleetSpec.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudFleetUnit

`func NewCloudFleetUnit() *CloudFleetUnit`

NewCloudFleetUnit instantiates a new CloudFleetUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFleetUnitWithDefaults

`func NewCloudFleetUnitWithDefaults() *CloudFleetUnit`

NewCloudFleetUnitWithDefaults instantiates a new CloudFleetUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *CloudFleetUnit) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CloudFleetUnit) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CloudFleetUnit) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CloudFleetUnit) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKind

`func (o *CloudFleetUnit) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudFleetUnit) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudFleetUnit) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudFleetUnit) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *CloudFleetUnit) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CloudFleetUnit) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CloudFleetUnit) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CloudFleetUnit) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMetrics

`func (o *CloudFleetUnit) GetMetrics() CloudFleetMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *CloudFleetUnit) GetMetricsOk() (*CloudFleetMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *CloudFleetUnit) SetMetrics(v CloudFleetMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *CloudFleetUnit) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetQueued

`func (o *CloudFleetUnit) GetQueued() int32`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *CloudFleetUnit) GetQueuedOk() (*int32, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *CloudFleetUnit) SetQueued(v int32)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *CloudFleetUnit) HasQueued() bool`

HasQueued returns a boolean if a field has been set.

### GetRunning

`func (o *CloudFleetUnit) GetRunning() int32`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *CloudFleetUnit) GetRunningOk() (*int32, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *CloudFleetUnit) SetRunning(v int32)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *CloudFleetUnit) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetSessions

`func (o *CloudFleetUnit) GetSessions() int32`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *CloudFleetUnit) GetSessionsOk() (*int32, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *CloudFleetUnit) SetSessions(v int32)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *CloudFleetUnit) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetSource

`func (o *CloudFleetUnit) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudFleetUnit) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudFleetUnit) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudFleetUnit) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSpec

`func (o *CloudFleetUnit) GetSpec() CloudFleetSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CloudFleetUnit) GetSpecOk() (*CloudFleetSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CloudFleetUnit) SetSpec(v CloudFleetSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CloudFleetUnit) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *CloudFleetUnit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudFleetUnit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudFleetUnit) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudFleetUnit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUnit

`func (o *CloudFleetUnit) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CloudFleetUnit) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CloudFleetUnit) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CloudFleetUnit) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


