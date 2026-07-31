# ProjectsUpdateProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Must not be blank when supplied. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Framework** | Pointer to **string** |  | [optional] 
**CacheControl** | Pointer to **string** | Per-project HTML/document Cache-Control policy; must not contain newlines. | [optional] 
**Repo** | Pointer to [**PlatformCreateAppReqRepo**](PlatformCreateAppReqRepo.md) |  | [optional] 

## Methods

### NewProjectsUpdateProjectRequest

`func NewProjectsUpdateProjectRequest() *ProjectsUpdateProjectRequest`

NewProjectsUpdateProjectRequest instantiates a new ProjectsUpdateProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsUpdateProjectRequestWithDefaults

`func NewProjectsUpdateProjectRequestWithDefaults() *ProjectsUpdateProjectRequest`

NewProjectsUpdateProjectRequestWithDefaults instantiates a new ProjectsUpdateProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectsUpdateProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsUpdateProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsUpdateProjectRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectsUpdateProjectRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectsUpdateProjectRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectsUpdateProjectRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectsUpdateProjectRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectsUpdateProjectRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFramework

`func (o *ProjectsUpdateProjectRequest) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *ProjectsUpdateProjectRequest) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *ProjectsUpdateProjectRequest) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *ProjectsUpdateProjectRequest) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetCacheControl

`func (o *ProjectsUpdateProjectRequest) GetCacheControl() string`

GetCacheControl returns the CacheControl field if non-nil, zero value otherwise.

### GetCacheControlOk

`func (o *ProjectsUpdateProjectRequest) GetCacheControlOk() (*string, bool)`

GetCacheControlOk returns a tuple with the CacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheControl

`func (o *ProjectsUpdateProjectRequest) SetCacheControl(v string)`

SetCacheControl sets CacheControl field to given value.

### HasCacheControl

`func (o *ProjectsUpdateProjectRequest) HasCacheControl() bool`

HasCacheControl returns a boolean if a field has been set.

### GetRepo

`func (o *ProjectsUpdateProjectRequest) GetRepo() PlatformCreateAppReqRepo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ProjectsUpdateProjectRequest) GetRepoOk() (*PlatformCreateAppReqRepo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ProjectsUpdateProjectRequest) SetRepo(v PlatformCreateAppReqRepo)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ProjectsUpdateProjectRequest) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


