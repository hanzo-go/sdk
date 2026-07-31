# ProjectsCreateProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name (required). | 
**Slug** | Pointer to **string** | Org-unique handle; derived from name when omitted. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Framework** | Pointer to **string** | Build hint; defaults to \&quot;static\&quot;. | [optional] 
**Repo** | Pointer to [**PlatformCreateAppReqRepo**](PlatformCreateAppReqRepo.md) |  | [optional] 

## Methods

### NewProjectsCreateProjectRequest

`func NewProjectsCreateProjectRequest(name string, ) *ProjectsCreateProjectRequest`

NewProjectsCreateProjectRequest instantiates a new ProjectsCreateProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsCreateProjectRequestWithDefaults

`func NewProjectsCreateProjectRequestWithDefaults() *ProjectsCreateProjectRequest`

NewProjectsCreateProjectRequestWithDefaults instantiates a new ProjectsCreateProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectsCreateProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsCreateProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsCreateProjectRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *ProjectsCreateProjectRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsCreateProjectRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsCreateProjectRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectsCreateProjectRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectsCreateProjectRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectsCreateProjectRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectsCreateProjectRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectsCreateProjectRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFramework

`func (o *ProjectsCreateProjectRequest) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *ProjectsCreateProjectRequest) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *ProjectsCreateProjectRequest) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *ProjectsCreateProjectRequest) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetRepo

`func (o *ProjectsCreateProjectRequest) GetRepo() PlatformCreateAppReqRepo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ProjectsCreateProjectRequest) GetRepoOk() (*PlatformCreateAppReqRepo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ProjectsCreateProjectRequest) SetRepo(v PlatformCreateAppReqRepo)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ProjectsCreateProjectRequest) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


