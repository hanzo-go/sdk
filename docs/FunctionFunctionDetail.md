# FunctionFunctionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgDurationMs** | Pointer to **float64** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**EnvCount** | Pointer to **int64** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Errors7d** | Pointer to **int64** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Invocations7d** | Pointer to **int64** |  | [optional] 
**LastDeployedAt** | Pointer to **string** |  | [optional] 
**MemoryLimit** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**RecentInvocations** | Pointer to [**[]FunctionInvocationView**](FunctionInvocationView.md) | RecentInvocations is its twenty most recent runs, newest first. | [optional] 
**Secrets** | Pointer to **[]string** | Secrets are the NAMES it mounts. Values are never read or returned. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SuccessRate** | Pointer to **float64** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**TimeoutSec** | Pointer to **int64** |  | [optional] 
**Triggers** | Pointer to [**[]FunctionTriggerView**](FunctionTriggerView.md) | Triggers is how this function is reached. | [optional] 

## Methods

### NewFunctionFunctionDetail

`func NewFunctionFunctionDetail() *FunctionFunctionDetail`

NewFunctionFunctionDetail instantiates a new FunctionFunctionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionFunctionDetailWithDefaults

`func NewFunctionFunctionDetailWithDefaults() *FunctionFunctionDetail`

NewFunctionFunctionDetailWithDefaults instantiates a new FunctionFunctionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgDurationMs

`func (o *FunctionFunctionDetail) GetAvgDurationMs() float64`

GetAvgDurationMs returns the AvgDurationMs field if non-nil, zero value otherwise.

### GetAvgDurationMsOk

`func (o *FunctionFunctionDetail) GetAvgDurationMsOk() (*float64, bool)`

GetAvgDurationMsOk returns a tuple with the AvgDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDurationMs

`func (o *FunctionFunctionDetail) SetAvgDurationMs(v float64)`

SetAvgDurationMs sets AvgDurationMs field to given value.

### HasAvgDurationMs

`func (o *FunctionFunctionDetail) HasAvgDurationMs() bool`

HasAvgDurationMs returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FunctionFunctionDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FunctionFunctionDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FunctionFunctionDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FunctionFunctionDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEndpoint

`func (o *FunctionFunctionDetail) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *FunctionFunctionDetail) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *FunctionFunctionDetail) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *FunctionFunctionDetail) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetEnvCount

`func (o *FunctionFunctionDetail) GetEnvCount() int64`

GetEnvCount returns the EnvCount field if non-nil, zero value otherwise.

### GetEnvCountOk

`func (o *FunctionFunctionDetail) GetEnvCountOk() (*int64, bool)`

GetEnvCountOk returns a tuple with the EnvCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvCount

`func (o *FunctionFunctionDetail) SetEnvCount(v int64)`

SetEnvCount sets EnvCount field to given value.

### HasEnvCount

`func (o *FunctionFunctionDetail) HasEnvCount() bool`

HasEnvCount returns a boolean if a field has been set.

### GetEnvironment

`func (o *FunctionFunctionDetail) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FunctionFunctionDetail) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FunctionFunctionDetail) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FunctionFunctionDetail) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetErrors7d

`func (o *FunctionFunctionDetail) GetErrors7d() int64`

GetErrors7d returns the Errors7d field if non-nil, zero value otherwise.

### GetErrors7dOk

`func (o *FunctionFunctionDetail) GetErrors7dOk() (*int64, bool)`

GetErrors7dOk returns a tuple with the Errors7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors7d

`func (o *FunctionFunctionDetail) SetErrors7d(v int64)`

SetErrors7d sets Errors7d field to given value.

### HasErrors7d

`func (o *FunctionFunctionDetail) HasErrors7d() bool`

HasErrors7d returns a boolean if a field has been set.

### GetImage

`func (o *FunctionFunctionDetail) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *FunctionFunctionDetail) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *FunctionFunctionDetail) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *FunctionFunctionDetail) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetInvocations7d

`func (o *FunctionFunctionDetail) GetInvocations7d() int64`

GetInvocations7d returns the Invocations7d field if non-nil, zero value otherwise.

### GetInvocations7dOk

`func (o *FunctionFunctionDetail) GetInvocations7dOk() (*int64, bool)`

GetInvocations7dOk returns a tuple with the Invocations7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocations7d

`func (o *FunctionFunctionDetail) SetInvocations7d(v int64)`

SetInvocations7d sets Invocations7d field to given value.

### HasInvocations7d

`func (o *FunctionFunctionDetail) HasInvocations7d() bool`

HasInvocations7d returns a boolean if a field has been set.

### GetLastDeployedAt

`func (o *FunctionFunctionDetail) GetLastDeployedAt() string`

GetLastDeployedAt returns the LastDeployedAt field if non-nil, zero value otherwise.

### GetLastDeployedAtOk

`func (o *FunctionFunctionDetail) GetLastDeployedAtOk() (*string, bool)`

GetLastDeployedAtOk returns a tuple with the LastDeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeployedAt

`func (o *FunctionFunctionDetail) SetLastDeployedAt(v string)`

SetLastDeployedAt sets LastDeployedAt field to given value.

### HasLastDeployedAt

`func (o *FunctionFunctionDetail) HasLastDeployedAt() bool`

HasLastDeployedAt returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *FunctionFunctionDetail) GetMemoryLimit() string`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *FunctionFunctionDetail) GetMemoryLimitOk() (*string, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *FunctionFunctionDetail) SetMemoryLimit(v string)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *FunctionFunctionDetail) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetName

`func (o *FunctionFunctionDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionFunctionDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionFunctionDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FunctionFunctionDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *FunctionFunctionDetail) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *FunctionFunctionDetail) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *FunctionFunctionDetail) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *FunctionFunctionDetail) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetRecentInvocations

`func (o *FunctionFunctionDetail) GetRecentInvocations() []FunctionInvocationView`

GetRecentInvocations returns the RecentInvocations field if non-nil, zero value otherwise.

### GetRecentInvocationsOk

`func (o *FunctionFunctionDetail) GetRecentInvocationsOk() (*[]FunctionInvocationView, bool)`

GetRecentInvocationsOk returns a tuple with the RecentInvocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentInvocations

`func (o *FunctionFunctionDetail) SetRecentInvocations(v []FunctionInvocationView)`

SetRecentInvocations sets RecentInvocations field to given value.

### HasRecentInvocations

`func (o *FunctionFunctionDetail) HasRecentInvocations() bool`

HasRecentInvocations returns a boolean if a field has been set.

### GetSecrets

`func (o *FunctionFunctionDetail) GetSecrets() []string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *FunctionFunctionDetail) GetSecretsOk() (*[]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *FunctionFunctionDetail) SetSecrets(v []string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *FunctionFunctionDetail) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetStatus

`func (o *FunctionFunctionDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FunctionFunctionDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FunctionFunctionDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FunctionFunctionDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccessRate

`func (o *FunctionFunctionDetail) GetSuccessRate() float64`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *FunctionFunctionDetail) GetSuccessRateOk() (*float64, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *FunctionFunctionDetail) SetSuccessRate(v float64)`

SetSuccessRate sets SuccessRate field to given value.

### HasSuccessRate

`func (o *FunctionFunctionDetail) HasSuccessRate() bool`

HasSuccessRate returns a boolean if a field has been set.

### GetTarget

`func (o *FunctionFunctionDetail) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *FunctionFunctionDetail) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *FunctionFunctionDetail) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *FunctionFunctionDetail) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *FunctionFunctionDetail) GetTimeoutSec() int64`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *FunctionFunctionDetail) GetTimeoutSecOk() (*int64, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *FunctionFunctionDetail) SetTimeoutSec(v int64)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *FunctionFunctionDetail) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.

### GetTriggers

`func (o *FunctionFunctionDetail) GetTriggers() []FunctionTriggerView`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *FunctionFunctionDetail) GetTriggersOk() (*[]FunctionTriggerView, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *FunctionFunctionDetail) SetTriggers(v []FunctionTriggerView)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *FunctionFunctionDetail) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


