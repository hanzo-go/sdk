# Definition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the source to run, capped so one function cannot amplify the store. | [optional] 
**EnvNames** | Pointer to **[]string** | EnvNames are the secret NAMES to mount. Values live in the secret store and are never carried here. | [optional] 
**Environment** | Pointer to **string** | Environment is a second spelling of runtime, accepted so a console that says \&quot;environment\&quot; needs no translation. | [optional] 
**Handler** | Pointer to **string** | Handler is the entry point within the code. | [optional] 
**Image** | Pointer to **string** | Image names a prebuilt image to run instead of source. | [optional] 
**MemoryLimit** | Pointer to **string** | MemoryLimit is the memory the function runs with, defaulting to 256Mi. It is also the multiplier on the GB-seconds compute charge. | [optional] 
**Name** | **string** | Name is the function&#39;s org-unique handle and the segment that addresses it, matching ^[A-Za-z0-9][A-Za-z0-9._-]{0,63}$. The names that would shadow a collection route are reserved. | 
**Namespace** | Pointer to **string** | Namespace groups functions for display. It is cosmetic — the org is the isolation key — and is normalised to a DNS-safe label. | [optional] 
**Runtime** | Pointer to **string** | Runtime is the language the code runs under: node, python or deno. | [optional] 
**Target** | Pointer to **string** | Target is where the function runs: sandbox (the default) or fleet, the org&#39;s own GPU fleet. fleet supports runtime&#x3D;python only. | [optional] 
**TimeoutSec** | Pointer to **int64** | TimeoutSec is the per-invocation deadline, defaulting to 30 and clamped at 900 — a larger value is capped rather than reset to the default. | [optional] 

## Methods

### NewDefinition

`func NewDefinition(name string, ) *Definition`

NewDefinition instantiates a new Definition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefinitionWithDefaults

`func NewDefinitionWithDefaults() *Definition`

NewDefinitionWithDefaults instantiates a new Definition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Definition) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Definition) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Definition) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Definition) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetEnvNames

`func (o *Definition) GetEnvNames() []string`

GetEnvNames returns the EnvNames field if non-nil, zero value otherwise.

### GetEnvNamesOk

`func (o *Definition) GetEnvNamesOk() (*[]string, bool)`

GetEnvNamesOk returns a tuple with the EnvNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvNames

`func (o *Definition) SetEnvNames(v []string)`

SetEnvNames sets EnvNames field to given value.

### HasEnvNames

`func (o *Definition) HasEnvNames() bool`

HasEnvNames returns a boolean if a field has been set.

### GetEnvironment

`func (o *Definition) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Definition) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Definition) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *Definition) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetHandler

`func (o *Definition) GetHandler() string`

GetHandler returns the Handler field if non-nil, zero value otherwise.

### GetHandlerOk

`func (o *Definition) GetHandlerOk() (*string, bool)`

GetHandlerOk returns a tuple with the Handler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandler

`func (o *Definition) SetHandler(v string)`

SetHandler sets Handler field to given value.

### HasHandler

`func (o *Definition) HasHandler() bool`

HasHandler returns a boolean if a field has been set.

### GetImage

`func (o *Definition) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Definition) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Definition) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Definition) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *Definition) GetMemoryLimit() string`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *Definition) GetMemoryLimitOk() (*string, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *Definition) SetMemoryLimit(v string)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *Definition) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetName

`func (o *Definition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Definition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Definition) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *Definition) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Definition) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Definition) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Definition) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetRuntime

`func (o *Definition) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *Definition) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *Definition) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *Definition) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetTarget

`func (o *Definition) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Definition) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Definition) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *Definition) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *Definition) GetTimeoutSec() int64`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *Definition) GetTimeoutSecOk() (*int64, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *Definition) SetTimeoutSec(v int64)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *Definition) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


