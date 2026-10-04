# AgentTargetReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **string** | Capacity is a human summary of the machine&#39;s size, up to 256 characters. Prose only; a scheduler reads Spec. | [optional] 
**Host** | Pointer to **string** | Host is the hostname sessions on this machine will report. It is what makes a re-link IDEMPOTENT: the same (org, host, owner) refreshes the existing row and answers 200, while a request with no host always creates a new target and answers 201. It never adopts a row owned by somebody else. | [optional] 
**Kind** | Pointer to **string** | Kind is laptop | cloud | gpu | cluster | machine. Empty registers a &#x60;machine&#x60;; anything outside the five is a 400. | [optional] 
**Label** | Pointer to **string** | Label is the name to show for this machine, up to 128 characters. REQUIRED — it is the only field here a person reads. | [optional] 
**Metrics** | Pointer to [**AgentMetrics**](AgentMetrics.md) | Metrics is a live sample, and sending one IS A HEARTBEAT: it refreshes the row and starts the 90-second liveness window, and it is appended to the fleet series as one point. Its own &#x60;at&#x60; is ignored — the server stamps the time, so a client can never age or backdate its own machine. Omit it to register a machine without claiming it is alive. | [optional] 
**Spec** | Pointer to [**AgentSpec**](AgentSpec.md) | Spec is the machine&#39;s static capability — os, arch, cores, RAM, accelerators. Every field is bounded on write and at most 32 accelerators are accepted, so what comes back may be clamped. Omit it for a destination nothing probes. | [optional] 
**Status** | Pointer to **string** | Status is online | offline | draining. Empty registers &#x60;online&#x60;. It states INTENT — a heartbeat is what decides whether an online machine is actually reachable, so declaring online does not make it so. | [optional] 

## Methods

### NewAgentTargetReq

`func NewAgentTargetReq() *AgentTargetReq`

NewAgentTargetReq instantiates a new AgentTargetReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTargetReqWithDefaults

`func NewAgentTargetReqWithDefaults() *AgentTargetReq`

NewAgentTargetReqWithDefaults instantiates a new AgentTargetReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *AgentTargetReq) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *AgentTargetReq) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *AgentTargetReq) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *AgentTargetReq) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetHost

`func (o *AgentTargetReq) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AgentTargetReq) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AgentTargetReq) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AgentTargetReq) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKind

`func (o *AgentTargetReq) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AgentTargetReq) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AgentTargetReq) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AgentTargetReq) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *AgentTargetReq) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AgentTargetReq) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AgentTargetReq) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AgentTargetReq) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMetrics

`func (o *AgentTargetReq) GetMetrics() AgentMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *AgentTargetReq) GetMetricsOk() (*AgentMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *AgentTargetReq) SetMetrics(v AgentMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *AgentTargetReq) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetSpec

`func (o *AgentTargetReq) GetSpec() AgentSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AgentTargetReq) GetSpecOk() (*AgentSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AgentTargetReq) SetSpec(v AgentSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AgentTargetReq) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *AgentTargetReq) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentTargetReq) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentTargetReq) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentTargetReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


