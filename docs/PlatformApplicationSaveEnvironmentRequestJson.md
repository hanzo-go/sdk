# PlatformApplicationSaveEnvironmentRequestJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** |  | 
**Env** | Pointer to **string** | KEY&#x3D;VALUE newline-separated | [optional] 
**BuildArgs** | Pointer to **string** |  | [optional] 

## Methods

### NewPlatformApplicationSaveEnvironmentRequestJson

`func NewPlatformApplicationSaveEnvironmentRequestJson(applicationId string, ) *PlatformApplicationSaveEnvironmentRequestJson`

NewPlatformApplicationSaveEnvironmentRequestJson instantiates a new PlatformApplicationSaveEnvironmentRequestJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformApplicationSaveEnvironmentRequestJsonWithDefaults

`func NewPlatformApplicationSaveEnvironmentRequestJsonWithDefaults() *PlatformApplicationSaveEnvironmentRequestJson`

NewPlatformApplicationSaveEnvironmentRequestJsonWithDefaults instantiates a new PlatformApplicationSaveEnvironmentRequestJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *PlatformApplicationSaveEnvironmentRequestJson) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PlatformApplicationSaveEnvironmentRequestJson) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PlatformApplicationSaveEnvironmentRequestJson) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetEnv

`func (o *PlatformApplicationSaveEnvironmentRequestJson) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *PlatformApplicationSaveEnvironmentRequestJson) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *PlatformApplicationSaveEnvironmentRequestJson) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *PlatformApplicationSaveEnvironmentRequestJson) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetBuildArgs

`func (o *PlatformApplicationSaveEnvironmentRequestJson) GetBuildArgs() string`

GetBuildArgs returns the BuildArgs field if non-nil, zero value otherwise.

### GetBuildArgsOk

`func (o *PlatformApplicationSaveEnvironmentRequestJson) GetBuildArgsOk() (*string, bool)`

GetBuildArgsOk returns a tuple with the BuildArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildArgs

`func (o *PlatformApplicationSaveEnvironmentRequestJson) SetBuildArgs(v string)`

SetBuildArgs sets BuildArgs field to given value.

### HasBuildArgs

`func (o *PlatformApplicationSaveEnvironmentRequestJson) HasBuildArgs() bool`

HasBuildArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


