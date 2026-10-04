# AgentPatchTargetIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **string** | Capacity rewrites the human summary, up to 256 characters. \&quot;\&quot; clears it. | [optional] 
**Host** | Pointer to **string** | Host re-points the hostname sessions are matched by. Moving it moves the load: the session counts follow the new name from the next read. | [optional] 
**Id** | Pointer to **string** | ID is the target to update, from the path. | [optional] 
**Kind** | Pointer to **string** | Kind re-files it under laptop | cloud | gpu | cluster | machine. | [optional] 
**Label** | Pointer to **string** | Label renames the machine, up to 128 characters. Empty STRING is refused — a target with no name is a row nobody can pick out of a fleet. | [optional] 
**Metrics** | Pointer to [**AgentMetrics**](AgentMetrics.md) | Metrics replaces the live sample, and sending one IS A HEARTBEAT: the server stamps the time and appends the point to the fleet series. Sending an all-zero sample CLEARS the heartbeat — the machine goes back to having no liveness fact at all, and its stored status is taken at face value again. | [optional] 
**Spec** | Pointer to [**AgentSpec**](AgentSpec.md) | Spec replaces the static capability whole, sanitized and clamped the same way a register&#39;s is. | [optional] 
**Status** | Pointer to **string** | Status sets operator INTENT: online | offline | draining. Draining is how a machine is taken out of dispatch without ending what is already on it. What comes back may still read offline, because the heartbeat outranks the intent. | [optional] 

## Methods

### NewAgentPatchTargetIn

`func NewAgentPatchTargetIn() *AgentPatchTargetIn`

NewAgentPatchTargetIn instantiates a new AgentPatchTargetIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentPatchTargetInWithDefaults

`func NewAgentPatchTargetInWithDefaults() *AgentPatchTargetIn`

NewAgentPatchTargetInWithDefaults instantiates a new AgentPatchTargetIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *AgentPatchTargetIn) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *AgentPatchTargetIn) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *AgentPatchTargetIn) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *AgentPatchTargetIn) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetHost

`func (o *AgentPatchTargetIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AgentPatchTargetIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AgentPatchTargetIn) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AgentPatchTargetIn) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *AgentPatchTargetIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentPatchTargetIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentPatchTargetIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentPatchTargetIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *AgentPatchTargetIn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AgentPatchTargetIn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AgentPatchTargetIn) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AgentPatchTargetIn) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *AgentPatchTargetIn) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AgentPatchTargetIn) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AgentPatchTargetIn) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AgentPatchTargetIn) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMetrics

`func (o *AgentPatchTargetIn) GetMetrics() AgentMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *AgentPatchTargetIn) GetMetricsOk() (*AgentMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *AgentPatchTargetIn) SetMetrics(v AgentMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *AgentPatchTargetIn) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetSpec

`func (o *AgentPatchTargetIn) GetSpec() AgentSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AgentPatchTargetIn) GetSpecOk() (*AgentSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AgentPatchTargetIn) SetSpec(v AgentSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AgentPatchTargetIn) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *AgentPatchTargetIn) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentPatchTargetIn) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentPatchTargetIn) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentPatchTargetIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


