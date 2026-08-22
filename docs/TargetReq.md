# TargetReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **string** | Capacity is a human summary of the machine&#39;s size, up to 256 characters. Prose only; a scheduler reads Spec. | [optional] 
**Host** | Pointer to **string** | Host is the hostname sessions on this machine will report. It is what makes a re-link IDEMPOTENT: the same (org, host, owner) refreshes the existing row and answers 200, while a request with no host always creates a new target and answers 201. It never adopts a row owned by somebody else. | [optional] 
**Kind** | Pointer to **string** | Kind is laptop | cloud | gpu | cluster | machine. Empty registers a &#x60;machine&#x60;; anything outside the five is a 400. | [optional] 
**Label** | Pointer to **string** | Label is the name to show for this machine, up to 128 characters. REQUIRED — it is the only field here a person reads. | [optional] 
**Metrics** | Pointer to [**Metrics**](Metrics.md) | Metrics is a live sample, and sending one IS A HEARTBEAT: it refreshes the row and starts the 90-second liveness window, and it is appended to the fleet series as one point. Its own &#x60;at&#x60; is ignored — the server stamps the time, so a client can never age or backdate its own machine. Omit it to register a machine without claiming it is alive. | [optional] 
**Spec** | Pointer to [**Spec**](Spec.md) | Spec is the machine&#39;s static capability — os, arch, cores, RAM, accelerators. Every field is bounded on write and at most 32 accelerators are accepted, so what comes back may be clamped. Omit it for a destination nothing probes. | [optional] 
**Status** | Pointer to **string** | Status is online | offline | draining. Empty registers &#x60;online&#x60;. It states INTENT — a heartbeat is what decides whether an online machine is actually reachable, so declaring online does not make it so. | [optional] 

## Methods

### NewTargetReq

`func NewTargetReq() *TargetReq`

NewTargetReq instantiates a new TargetReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetReqWithDefaults

`func NewTargetReqWithDefaults() *TargetReq`

NewTargetReqWithDefaults instantiates a new TargetReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *TargetReq) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *TargetReq) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *TargetReq) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *TargetReq) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetHost

`func (o *TargetReq) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TargetReq) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TargetReq) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TargetReq) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKind

`func (o *TargetReq) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TargetReq) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TargetReq) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TargetReq) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *TargetReq) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TargetReq) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TargetReq) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TargetReq) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMetrics

`func (o *TargetReq) GetMetrics() Metrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *TargetReq) GetMetricsOk() (*Metrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *TargetReq) SetMetrics(v Metrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *TargetReq) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetSpec

`func (o *TargetReq) GetSpec() Spec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *TargetReq) GetSpecOk() (*Spec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *TargetReq) SetSpec(v Spec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *TargetReq) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *TargetReq) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TargetReq) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TargetReq) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TargetReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


