# PlatformCreateDatabaseInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AppName** | Pointer to **string** |  | [optional] 
**EnvironmentId** | **string** |  | 
**ServerId** | Pointer to **string** |  | [optional] 
**DockerImage** | Pointer to **string** |  | [optional] 
**DatabaseName** | Pointer to **string** |  | [optional] 
**DatabaseUser** | Pointer to **string** |  | [optional] 
**DatabasePassword** | Pointer to **string** |  | [optional] 
**DatabaseRootPassword** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPlatformCreateDatabaseInput

`func NewPlatformCreateDatabaseInput(name string, environmentId string, ) *PlatformCreateDatabaseInput`

NewPlatformCreateDatabaseInput instantiates a new PlatformCreateDatabaseInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformCreateDatabaseInputWithDefaults

`func NewPlatformCreateDatabaseInputWithDefaults() *PlatformCreateDatabaseInput`

NewPlatformCreateDatabaseInputWithDefaults instantiates a new PlatformCreateDatabaseInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlatformCreateDatabaseInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformCreateDatabaseInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformCreateDatabaseInput) SetName(v string)`

SetName sets Name field to given value.


### GetAppName

`func (o *PlatformCreateDatabaseInput) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *PlatformCreateDatabaseInput) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *PlatformCreateDatabaseInput) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *PlatformCreateDatabaseInput) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *PlatformCreateDatabaseInput) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *PlatformCreateDatabaseInput) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *PlatformCreateDatabaseInput) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetServerId

`func (o *PlatformCreateDatabaseInput) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *PlatformCreateDatabaseInput) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *PlatformCreateDatabaseInput) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *PlatformCreateDatabaseInput) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetDockerImage

`func (o *PlatformCreateDatabaseInput) GetDockerImage() string`

GetDockerImage returns the DockerImage field if non-nil, zero value otherwise.

### GetDockerImageOk

`func (o *PlatformCreateDatabaseInput) GetDockerImageOk() (*string, bool)`

GetDockerImageOk returns a tuple with the DockerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImage

`func (o *PlatformCreateDatabaseInput) SetDockerImage(v string)`

SetDockerImage sets DockerImage field to given value.

### HasDockerImage

`func (o *PlatformCreateDatabaseInput) HasDockerImage() bool`

HasDockerImage returns a boolean if a field has been set.

### GetDatabaseName

`func (o *PlatformCreateDatabaseInput) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *PlatformCreateDatabaseInput) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *PlatformCreateDatabaseInput) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *PlatformCreateDatabaseInput) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetDatabaseUser

`func (o *PlatformCreateDatabaseInput) GetDatabaseUser() string`

GetDatabaseUser returns the DatabaseUser field if non-nil, zero value otherwise.

### GetDatabaseUserOk

`func (o *PlatformCreateDatabaseInput) GetDatabaseUserOk() (*string, bool)`

GetDatabaseUserOk returns a tuple with the DatabaseUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUser

`func (o *PlatformCreateDatabaseInput) SetDatabaseUser(v string)`

SetDatabaseUser sets DatabaseUser field to given value.

### HasDatabaseUser

`func (o *PlatformCreateDatabaseInput) HasDatabaseUser() bool`

HasDatabaseUser returns a boolean if a field has been set.

### GetDatabasePassword

`func (o *PlatformCreateDatabaseInput) GetDatabasePassword() string`

GetDatabasePassword returns the DatabasePassword field if non-nil, zero value otherwise.

### GetDatabasePasswordOk

`func (o *PlatformCreateDatabaseInput) GetDatabasePasswordOk() (*string, bool)`

GetDatabasePasswordOk returns a tuple with the DatabasePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasePassword

`func (o *PlatformCreateDatabaseInput) SetDatabasePassword(v string)`

SetDatabasePassword sets DatabasePassword field to given value.

### HasDatabasePassword

`func (o *PlatformCreateDatabaseInput) HasDatabasePassword() bool`

HasDatabasePassword returns a boolean if a field has been set.

### GetDatabaseRootPassword

`func (o *PlatformCreateDatabaseInput) GetDatabaseRootPassword() string`

GetDatabaseRootPassword returns the DatabaseRootPassword field if non-nil, zero value otherwise.

### GetDatabaseRootPasswordOk

`func (o *PlatformCreateDatabaseInput) GetDatabaseRootPasswordOk() (*string, bool)`

GetDatabaseRootPasswordOk returns a tuple with the DatabaseRootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseRootPassword

`func (o *PlatformCreateDatabaseInput) SetDatabaseRootPassword(v string)`

SetDatabaseRootPassword sets DatabaseRootPassword field to given value.

### HasDatabaseRootPassword

`func (o *PlatformCreateDatabaseInput) HasDatabaseRootPassword() bool`

HasDatabaseRootPassword returns a boolean if a field has been set.

### GetDescription

`func (o *PlatformCreateDatabaseInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformCreateDatabaseInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformCreateDatabaseInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformCreateDatabaseInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


