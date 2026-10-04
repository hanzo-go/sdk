# ProjectProjectsDeploySite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]ProjectProjectsFile**](ProjectProjectsFile.md) | Files is the whole site, inline — every file it consists of. It REPLACES what is there rather than merging, so an omitted file is a deleted one. | [optional] 
**Name** | Pointer to **string** | Name is the site&#39;s display name. | [optional] 
**Slug** | Pointer to **string** | Slug is the handle and public host label to publish under. | [optional] 

## Methods

### NewProjectProjectsDeploySite

`func NewProjectProjectsDeploySite() *ProjectProjectsDeploySite`

NewProjectProjectsDeploySite instantiates a new ProjectProjectsDeploySite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectsDeploySiteWithDefaults

`func NewProjectProjectsDeploySiteWithDefaults() *ProjectProjectsDeploySite`

NewProjectProjectsDeploySiteWithDefaults instantiates a new ProjectProjectsDeploySite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *ProjectProjectsDeploySite) GetFiles() []ProjectProjectsFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ProjectProjectsDeploySite) GetFilesOk() (*[]ProjectProjectsFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ProjectProjectsDeploySite) SetFiles(v []ProjectProjectsFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ProjectProjectsDeploySite) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetName

`func (o *ProjectProjectsDeploySite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectProjectsDeploySite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectProjectsDeploySite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectProjectsDeploySite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectProjectsDeploySite) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectProjectsDeploySite) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectProjectsDeploySite) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectProjectsDeploySite) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


