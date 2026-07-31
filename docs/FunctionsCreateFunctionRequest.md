# FunctionsCreateFunctionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Environment** | Pointer to **string** |  | [optional] 
**Runtime** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Handler** | Pointer to **string** |  | [optional] 
**TimeoutSec** | Pointer to **int32** |  | [optional] 
**MemoryLimit** | Pointer to **string** |  | [optional] 
**EnvNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFunctionsCreateFunctionRequest

`func NewFunctionsCreateFunctionRequest(name string, ) *FunctionsCreateFunctionRequest`

NewFunctionsCreateFunctionRequest instantiates a new FunctionsCreateFunctionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionsCreateFunctionRequestWithDefaults

`func NewFunctionsCreateFunctionRequestWithDefaults() *FunctionsCreateFunctionRequest`

NewFunctionsCreateFunctionRequestWithDefaults instantiates a new FunctionsCreateFunctionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FunctionsCreateFunctionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionsCreateFunctionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionsCreateFunctionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnvironment

`func (o *FunctionsCreateFunctionRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FunctionsCreateFunctionRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FunctionsCreateFunctionRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FunctionsCreateFunctionRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetRuntime

`func (o *FunctionsCreateFunctionRequest) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *FunctionsCreateFunctionRequest) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *FunctionsCreateFunctionRequest) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *FunctionsCreateFunctionRequest) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetNamespace

`func (o *FunctionsCreateFunctionRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *FunctionsCreateFunctionRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *FunctionsCreateFunctionRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *FunctionsCreateFunctionRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetImage

`func (o *FunctionsCreateFunctionRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *FunctionsCreateFunctionRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *FunctionsCreateFunctionRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *FunctionsCreateFunctionRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetCode

`func (o *FunctionsCreateFunctionRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FunctionsCreateFunctionRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FunctionsCreateFunctionRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FunctionsCreateFunctionRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetHandler

`func (o *FunctionsCreateFunctionRequest) GetHandler() string`

GetHandler returns the Handler field if non-nil, zero value otherwise.

### GetHandlerOk

`func (o *FunctionsCreateFunctionRequest) GetHandlerOk() (*string, bool)`

GetHandlerOk returns a tuple with the Handler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandler

`func (o *FunctionsCreateFunctionRequest) SetHandler(v string)`

SetHandler sets Handler field to given value.

### HasHandler

`func (o *FunctionsCreateFunctionRequest) HasHandler() bool`

HasHandler returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *FunctionsCreateFunctionRequest) GetTimeoutSec() int32`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *FunctionsCreateFunctionRequest) GetTimeoutSecOk() (*int32, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *FunctionsCreateFunctionRequest) SetTimeoutSec(v int32)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *FunctionsCreateFunctionRequest) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *FunctionsCreateFunctionRequest) GetMemoryLimit() string`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *FunctionsCreateFunctionRequest) GetMemoryLimitOk() (*string, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *FunctionsCreateFunctionRequest) SetMemoryLimit(v string)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *FunctionsCreateFunctionRequest) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetEnvNames

`func (o *FunctionsCreateFunctionRequest) GetEnvNames() []string`

GetEnvNames returns the EnvNames field if non-nil, zero value otherwise.

### GetEnvNamesOk

`func (o *FunctionsCreateFunctionRequest) GetEnvNamesOk() (*[]string, bool)`

GetEnvNamesOk returns a tuple with the EnvNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvNames

`func (o *FunctionsCreateFunctionRequest) SetEnvNames(v []string)`

SetEnvNames sets EnvNames field to given value.

### HasEnvNames

`func (o *FunctionsCreateFunctionRequest) HasEnvNames() bool`

HasEnvNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


