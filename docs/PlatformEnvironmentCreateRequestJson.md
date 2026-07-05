# PlatformEnvironmentCreateRequestJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**Name** | **string** | Cannot be \&quot;production\&quot; (reserved) | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPlatformEnvironmentCreateRequestJson

`func NewPlatformEnvironmentCreateRequestJson(projectId string, name string, ) *PlatformEnvironmentCreateRequestJson`

NewPlatformEnvironmentCreateRequestJson instantiates a new PlatformEnvironmentCreateRequestJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformEnvironmentCreateRequestJsonWithDefaults

`func NewPlatformEnvironmentCreateRequestJsonWithDefaults() *PlatformEnvironmentCreateRequestJson`

NewPlatformEnvironmentCreateRequestJsonWithDefaults instantiates a new PlatformEnvironmentCreateRequestJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *PlatformEnvironmentCreateRequestJson) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PlatformEnvironmentCreateRequestJson) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PlatformEnvironmentCreateRequestJson) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *PlatformEnvironmentCreateRequestJson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformEnvironmentCreateRequestJson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformEnvironmentCreateRequestJson) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PlatformEnvironmentCreateRequestJson) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformEnvironmentCreateRequestJson) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformEnvironmentCreateRequestJson) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformEnvironmentCreateRequestJson) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


