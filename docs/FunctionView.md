# FunctionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgDurationMs** | Pointer to **float32** | mean wall-clock of those runs | [optional] 
**CreatedAt** | Pointer to **string** | when it was first published | [optional] 
**Endpoint** | Pointer to **string** | the path that invokes it | [optional] 
**EnvCount** | Pointer to **int32** | how many secret NAMES it mounts; values are never carried | [optional] 
**Environment** | Pointer to **string** | the language it runs under | [optional] 
**Errors7d** | Pointer to **int32** | how many of those runs failed | [optional] 
**Image** | Pointer to **string** | the prebuilt image it runs, when it runs one instead of source | [optional] 
**Invocations7d** | Pointer to **int32** | runs in the last 7 days; ABSENT, never 0, when it has not run | [optional] 
**LastDeployedAt** | Pointer to **string** | when its code last changed | [optional] 
**MemoryLimit** | Pointer to **string** | the memory it runs with, and the multiplier on its compute charge | [optional] 
**Name** | Pointer to **string** | the function&#39;s org-unique handle | [optional] 
**Namespace** | Pointer to **string** | the display group it belongs to; the org is the isolation key | [optional] 
**Status** | Pointer to **string** | whether it is ready to serve | [optional] 
**SuccessRate** | Pointer to **float32** | share of those runs that succeeded, 0..1 | [optional] 
**Target** | Pointer to **string** | where it runs: empty for the sandbox, \&quot;fleet\&quot; for the org&#39;s GPU fleet | [optional] 
**TimeoutSec** | Pointer to **int32** | its per-invocation deadline | [optional] 

## Methods

### NewFunctionView

`func NewFunctionView() *FunctionView`

NewFunctionView instantiates a new FunctionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionViewWithDefaults

`func NewFunctionViewWithDefaults() *FunctionView`

NewFunctionViewWithDefaults instantiates a new FunctionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgDurationMs

`func (o *FunctionView) GetAvgDurationMs() float32`

GetAvgDurationMs returns the AvgDurationMs field if non-nil, zero value otherwise.

### GetAvgDurationMsOk

`func (o *FunctionView) GetAvgDurationMsOk() (*float32, bool)`

GetAvgDurationMsOk returns a tuple with the AvgDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDurationMs

`func (o *FunctionView) SetAvgDurationMs(v float32)`

SetAvgDurationMs sets AvgDurationMs field to given value.

### HasAvgDurationMs

`func (o *FunctionView) HasAvgDurationMs() bool`

HasAvgDurationMs returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FunctionView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FunctionView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FunctionView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FunctionView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEndpoint

`func (o *FunctionView) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *FunctionView) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *FunctionView) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *FunctionView) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetEnvCount

`func (o *FunctionView) GetEnvCount() int32`

GetEnvCount returns the EnvCount field if non-nil, zero value otherwise.

### GetEnvCountOk

`func (o *FunctionView) GetEnvCountOk() (*int32, bool)`

GetEnvCountOk returns a tuple with the EnvCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvCount

`func (o *FunctionView) SetEnvCount(v int32)`

SetEnvCount sets EnvCount field to given value.

### HasEnvCount

`func (o *FunctionView) HasEnvCount() bool`

HasEnvCount returns a boolean if a field has been set.

### GetEnvironment

`func (o *FunctionView) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FunctionView) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FunctionView) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FunctionView) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetErrors7d

`func (o *FunctionView) GetErrors7d() int32`

GetErrors7d returns the Errors7d field if non-nil, zero value otherwise.

### GetErrors7dOk

`func (o *FunctionView) GetErrors7dOk() (*int32, bool)`

GetErrors7dOk returns a tuple with the Errors7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors7d

`func (o *FunctionView) SetErrors7d(v int32)`

SetErrors7d sets Errors7d field to given value.

### HasErrors7d

`func (o *FunctionView) HasErrors7d() bool`

HasErrors7d returns a boolean if a field has been set.

### GetImage

`func (o *FunctionView) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *FunctionView) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *FunctionView) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *FunctionView) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetInvocations7d

`func (o *FunctionView) GetInvocations7d() int32`

GetInvocations7d returns the Invocations7d field if non-nil, zero value otherwise.

### GetInvocations7dOk

`func (o *FunctionView) GetInvocations7dOk() (*int32, bool)`

GetInvocations7dOk returns a tuple with the Invocations7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocations7d

`func (o *FunctionView) SetInvocations7d(v int32)`

SetInvocations7d sets Invocations7d field to given value.

### HasInvocations7d

`func (o *FunctionView) HasInvocations7d() bool`

HasInvocations7d returns a boolean if a field has been set.

### GetLastDeployedAt

`func (o *FunctionView) GetLastDeployedAt() string`

GetLastDeployedAt returns the LastDeployedAt field if non-nil, zero value otherwise.

### GetLastDeployedAtOk

`func (o *FunctionView) GetLastDeployedAtOk() (*string, bool)`

GetLastDeployedAtOk returns a tuple with the LastDeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeployedAt

`func (o *FunctionView) SetLastDeployedAt(v string)`

SetLastDeployedAt sets LastDeployedAt field to given value.

### HasLastDeployedAt

`func (o *FunctionView) HasLastDeployedAt() bool`

HasLastDeployedAt returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *FunctionView) GetMemoryLimit() string`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *FunctionView) GetMemoryLimitOk() (*string, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *FunctionView) SetMemoryLimit(v string)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *FunctionView) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetName

`func (o *FunctionView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FunctionView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *FunctionView) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *FunctionView) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *FunctionView) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *FunctionView) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetStatus

`func (o *FunctionView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FunctionView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FunctionView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FunctionView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccessRate

`func (o *FunctionView) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *FunctionView) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *FunctionView) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.

### HasSuccessRate

`func (o *FunctionView) HasSuccessRate() bool`

HasSuccessRate returns a boolean if a field has been set.

### GetTarget

`func (o *FunctionView) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *FunctionView) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *FunctionView) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *FunctionView) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *FunctionView) GetTimeoutSec() int32`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *FunctionView) GetTimeoutSecOk() (*int32, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *FunctionView) SetTimeoutSec(v int32)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *FunctionView) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


