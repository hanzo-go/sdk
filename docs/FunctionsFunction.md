# FunctionsFunction

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

## Methods

### NewFunctionsFunction

`func NewFunctionsFunction() *FunctionsFunction`

NewFunctionsFunction instantiates a new FunctionsFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionsFunctionWithDefaults

`func NewFunctionsFunctionWithDefaults() *FunctionsFunction`

NewFunctionsFunctionWithDefaults instantiates a new FunctionsFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FunctionsFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionsFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionsFunction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FunctionsFunction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *FunctionsFunction) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *FunctionsFunction) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *FunctionsFunction) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *FunctionsFunction) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetEnvironment

`func (o *FunctionsFunction) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FunctionsFunction) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FunctionsFunction) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FunctionsFunction) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetStatus

`func (o *FunctionsFunction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FunctionsFunction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FunctionsFunction) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FunctionsFunction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImage

`func (o *FunctionsFunction) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *FunctionsFunction) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *FunctionsFunction) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *FunctionsFunction) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetEndpoint

`func (o *FunctionsFunction) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *FunctionsFunction) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *FunctionsFunction) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *FunctionsFunction) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetEnvCount

`func (o *FunctionsFunction) GetEnvCount() int32`

GetEnvCount returns the EnvCount field if non-nil, zero value otherwise.

### GetEnvCountOk

`func (o *FunctionsFunction) GetEnvCountOk() (*int32, bool)`

GetEnvCountOk returns a tuple with the EnvCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvCount

`func (o *FunctionsFunction) SetEnvCount(v int32)`

SetEnvCount sets EnvCount field to given value.

### HasEnvCount

`func (o *FunctionsFunction) HasEnvCount() bool`

HasEnvCount returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *FunctionsFunction) GetTimeoutSec() int32`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *FunctionsFunction) GetTimeoutSecOk() (*int32, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *FunctionsFunction) SetTimeoutSec(v int32)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *FunctionsFunction) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *FunctionsFunction) GetMemoryLimit() string`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *FunctionsFunction) GetMemoryLimitOk() (*string, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *FunctionsFunction) SetMemoryLimit(v string)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *FunctionsFunction) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetInvocations7d

`func (o *FunctionsFunction) GetInvocations7d() int32`

GetInvocations7d returns the Invocations7d field if non-nil, zero value otherwise.

### GetInvocations7dOk

`func (o *FunctionsFunction) GetInvocations7dOk() (*int32, bool)`

GetInvocations7dOk returns a tuple with the Invocations7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocations7d

`func (o *FunctionsFunction) SetInvocations7d(v int32)`

SetInvocations7d sets Invocations7d field to given value.

### HasInvocations7d

`func (o *FunctionsFunction) HasInvocations7d() bool`

HasInvocations7d returns a boolean if a field has been set.

### GetSuccessRate

`func (o *FunctionsFunction) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *FunctionsFunction) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *FunctionsFunction) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.

### HasSuccessRate

`func (o *FunctionsFunction) HasSuccessRate() bool`

HasSuccessRate returns a boolean if a field has been set.

### GetAvgDurationMs

`func (o *FunctionsFunction) GetAvgDurationMs() float32`

GetAvgDurationMs returns the AvgDurationMs field if non-nil, zero value otherwise.

### GetAvgDurationMsOk

`func (o *FunctionsFunction) GetAvgDurationMsOk() (*float32, bool)`

GetAvgDurationMsOk returns a tuple with the AvgDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDurationMs

`func (o *FunctionsFunction) SetAvgDurationMs(v float32)`

SetAvgDurationMs sets AvgDurationMs field to given value.

### HasAvgDurationMs

`func (o *FunctionsFunction) HasAvgDurationMs() bool`

HasAvgDurationMs returns a boolean if a field has been set.

### GetErrors7d

`func (o *FunctionsFunction) GetErrors7d() int32`

GetErrors7d returns the Errors7d field if non-nil, zero value otherwise.

### GetErrors7dOk

`func (o *FunctionsFunction) GetErrors7dOk() (*int32, bool)`

GetErrors7dOk returns a tuple with the Errors7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors7d

`func (o *FunctionsFunction) SetErrors7d(v int32)`

SetErrors7d sets Errors7d field to given value.

### HasErrors7d

`func (o *FunctionsFunction) HasErrors7d() bool`

HasErrors7d returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FunctionsFunction) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FunctionsFunction) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FunctionsFunction) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FunctionsFunction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastDeployedAt

`func (o *FunctionsFunction) GetLastDeployedAt() string`

GetLastDeployedAt returns the LastDeployedAt field if non-nil, zero value otherwise.

### GetLastDeployedAtOk

`func (o *FunctionsFunction) GetLastDeployedAtOk() (*string, bool)`

GetLastDeployedAtOk returns a tuple with the LastDeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeployedAt

`func (o *FunctionsFunction) SetLastDeployedAt(v string)`

SetLastDeployedAt sets LastDeployedAt field to given value.

### HasLastDeployedAt

`func (o *FunctionsFunction) HasLastDeployedAt() bool`

HasLastDeployedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


