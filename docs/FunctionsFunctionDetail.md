# FunctionsFunctionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**EnvCount** | Pointer to **int32** |  | [optional] 
**TimeoutSec** | Pointer to **int32** |  | [optional] 
**MemoryLimit** | Pointer to **string** |  | [optional] 
**Invocations7d** | Pointer to **int32** |  | [optional] 
**SuccessRate** | Pointer to **float32** |  | [optional] 
**AvgDurationMs** | Pointer to **float32** |  | [optional] 
**Errors7d** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**LastDeployedAt** | Pointer to **string** |  | [optional] 
**Triggers** | Pointer to [**[]FunctionsTrigger**](FunctionsTrigger.md) |  | [optional] 
**RecentInvocations** | Pointer to [**[]FunctionsInvocation**](FunctionsInvocation.md) |  | [optional] 
**Secrets** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFunctionsFunctionDetail

`func NewFunctionsFunctionDetail() *FunctionsFunctionDetail`

NewFunctionsFunctionDetail instantiates a new FunctionsFunctionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionsFunctionDetailWithDefaults

`func NewFunctionsFunctionDetailWithDefaults() *FunctionsFunctionDetail`

NewFunctionsFunctionDetailWithDefaults instantiates a new FunctionsFunctionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FunctionsFunctionDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionsFunctionDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionsFunctionDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FunctionsFunctionDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *FunctionsFunctionDetail) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *FunctionsFunctionDetail) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *FunctionsFunctionDetail) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *FunctionsFunctionDetail) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetEnvironment

`func (o *FunctionsFunctionDetail) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FunctionsFunctionDetail) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FunctionsFunctionDetail) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FunctionsFunctionDetail) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetStatus

`func (o *FunctionsFunctionDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FunctionsFunctionDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FunctionsFunctionDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FunctionsFunctionDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImage

`func (o *FunctionsFunctionDetail) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *FunctionsFunctionDetail) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *FunctionsFunctionDetail) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *FunctionsFunctionDetail) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetEndpoint

`func (o *FunctionsFunctionDetail) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *FunctionsFunctionDetail) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *FunctionsFunctionDetail) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *FunctionsFunctionDetail) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetEnvCount

`func (o *FunctionsFunctionDetail) GetEnvCount() int32`

GetEnvCount returns the EnvCount field if non-nil, zero value otherwise.

### GetEnvCountOk

`func (o *FunctionsFunctionDetail) GetEnvCountOk() (*int32, bool)`

GetEnvCountOk returns a tuple with the EnvCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvCount

`func (o *FunctionsFunctionDetail) SetEnvCount(v int32)`

SetEnvCount sets EnvCount field to given value.

### HasEnvCount

`func (o *FunctionsFunctionDetail) HasEnvCount() bool`

HasEnvCount returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *FunctionsFunctionDetail) GetTimeoutSec() int32`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *FunctionsFunctionDetail) GetTimeoutSecOk() (*int32, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *FunctionsFunctionDetail) SetTimeoutSec(v int32)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *FunctionsFunctionDetail) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *FunctionsFunctionDetail) GetMemoryLimit() string`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *FunctionsFunctionDetail) GetMemoryLimitOk() (*string, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *FunctionsFunctionDetail) SetMemoryLimit(v string)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *FunctionsFunctionDetail) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetInvocations7d

`func (o *FunctionsFunctionDetail) GetInvocations7d() int32`

GetInvocations7d returns the Invocations7d field if non-nil, zero value otherwise.

### GetInvocations7dOk

`func (o *FunctionsFunctionDetail) GetInvocations7dOk() (*int32, bool)`

GetInvocations7dOk returns a tuple with the Invocations7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocations7d

`func (o *FunctionsFunctionDetail) SetInvocations7d(v int32)`

SetInvocations7d sets Invocations7d field to given value.

### HasInvocations7d

`func (o *FunctionsFunctionDetail) HasInvocations7d() bool`

HasInvocations7d returns a boolean if a field has been set.

### GetSuccessRate

`func (o *FunctionsFunctionDetail) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *FunctionsFunctionDetail) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *FunctionsFunctionDetail) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.

### HasSuccessRate

`func (o *FunctionsFunctionDetail) HasSuccessRate() bool`

HasSuccessRate returns a boolean if a field has been set.

### GetAvgDurationMs

`func (o *FunctionsFunctionDetail) GetAvgDurationMs() float32`

GetAvgDurationMs returns the AvgDurationMs field if non-nil, zero value otherwise.

### GetAvgDurationMsOk

`func (o *FunctionsFunctionDetail) GetAvgDurationMsOk() (*float32, bool)`

GetAvgDurationMsOk returns a tuple with the AvgDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDurationMs

`func (o *FunctionsFunctionDetail) SetAvgDurationMs(v float32)`

SetAvgDurationMs sets AvgDurationMs field to given value.

### HasAvgDurationMs

`func (o *FunctionsFunctionDetail) HasAvgDurationMs() bool`

HasAvgDurationMs returns a boolean if a field has been set.

### GetErrors7d

`func (o *FunctionsFunctionDetail) GetErrors7d() int32`

GetErrors7d returns the Errors7d field if non-nil, zero value otherwise.

### GetErrors7dOk

`func (o *FunctionsFunctionDetail) GetErrors7dOk() (*int32, bool)`

GetErrors7dOk returns a tuple with the Errors7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors7d

`func (o *FunctionsFunctionDetail) SetErrors7d(v int32)`

SetErrors7d sets Errors7d field to given value.

### HasErrors7d

`func (o *FunctionsFunctionDetail) HasErrors7d() bool`

HasErrors7d returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FunctionsFunctionDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FunctionsFunctionDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FunctionsFunctionDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FunctionsFunctionDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastDeployedAt

`func (o *FunctionsFunctionDetail) GetLastDeployedAt() string`

GetLastDeployedAt returns the LastDeployedAt field if non-nil, zero value otherwise.

### GetLastDeployedAtOk

`func (o *FunctionsFunctionDetail) GetLastDeployedAtOk() (*string, bool)`

GetLastDeployedAtOk returns a tuple with the LastDeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeployedAt

`func (o *FunctionsFunctionDetail) SetLastDeployedAt(v string)`

SetLastDeployedAt sets LastDeployedAt field to given value.

### HasLastDeployedAt

`func (o *FunctionsFunctionDetail) HasLastDeployedAt() bool`

HasLastDeployedAt returns a boolean if a field has been set.

### GetTriggers

`func (o *FunctionsFunctionDetail) GetTriggers() []FunctionsTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *FunctionsFunctionDetail) GetTriggersOk() (*[]FunctionsTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *FunctionsFunctionDetail) SetTriggers(v []FunctionsTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *FunctionsFunctionDetail) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetRecentInvocations

`func (o *FunctionsFunctionDetail) GetRecentInvocations() []FunctionsInvocation`

GetRecentInvocations returns the RecentInvocations field if non-nil, zero value otherwise.

### GetRecentInvocationsOk

`func (o *FunctionsFunctionDetail) GetRecentInvocationsOk() (*[]FunctionsInvocation, bool)`

GetRecentInvocationsOk returns a tuple with the RecentInvocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentInvocations

`func (o *FunctionsFunctionDetail) SetRecentInvocations(v []FunctionsInvocation)`

SetRecentInvocations sets RecentInvocations field to given value.

### HasRecentInvocations

`func (o *FunctionsFunctionDetail) HasRecentInvocations() bool`

HasRecentInvocations returns a boolean if a field has been set.

### GetSecrets

`func (o *FunctionsFunctionDetail) GetSecrets() []string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *FunctionsFunctionDetail) GetSecretsOk() (*[]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *FunctionsFunctionDetail) SetSecrets(v []string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *FunctionsFunctionDetail) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


