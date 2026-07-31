# PlatformSetEnvReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to [**[]PlatformEnvVarJSON**](PlatformEnvVarJSON.md) |  | [optional] 

## Methods

### NewPlatformSetEnvReq

`func NewPlatformSetEnvReq() *PlatformSetEnvReq`

NewPlatformSetEnvReq instantiates a new PlatformSetEnvReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformSetEnvReqWithDefaults

`func NewPlatformSetEnvReqWithDefaults() *PlatformSetEnvReq`

NewPlatformSetEnvReqWithDefaults instantiates a new PlatformSetEnvReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *PlatformSetEnvReq) GetEnv() []PlatformEnvVarJSON`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *PlatformSetEnvReq) GetEnvOk() (*[]PlatformEnvVarJSON, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *PlatformSetEnvReq) SetEnv(v []PlatformEnvVarJSON)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *PlatformSetEnvReq) HasEnv() bool`

HasEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


