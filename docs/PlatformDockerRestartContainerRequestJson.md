# PlatformDockerRestartContainerRequestJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerId** | **string** |  | 
**ServerId** | Pointer to **string** |  | [optional] 

## Methods

### NewPlatformDockerRestartContainerRequestJson

`func NewPlatformDockerRestartContainerRequestJson(containerId string, ) *PlatformDockerRestartContainerRequestJson`

NewPlatformDockerRestartContainerRequestJson instantiates a new PlatformDockerRestartContainerRequestJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformDockerRestartContainerRequestJsonWithDefaults

`func NewPlatformDockerRestartContainerRequestJsonWithDefaults() *PlatformDockerRestartContainerRequestJson`

NewPlatformDockerRestartContainerRequestJsonWithDefaults instantiates a new PlatformDockerRestartContainerRequestJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerId

`func (o *PlatformDockerRestartContainerRequestJson) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *PlatformDockerRestartContainerRequestJson) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *PlatformDockerRestartContainerRequestJson) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.


### GetServerId

`func (o *PlatformDockerRestartContainerRequestJson) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *PlatformDockerRestartContainerRequestJson) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *PlatformDockerRestartContainerRequestJson) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *PlatformDockerRestartContainerRequestJson) HasServerId() bool`

HasServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


