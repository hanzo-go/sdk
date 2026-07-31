# ProjectsDeploySiteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | [**[]ProjectsSiteFile**](ProjectsSiteFile.md) | Site files (required, non-empty); index.html must be present at the root. | 
**Slug** | Pointer to **string** | Target project slug; derived from name (or minted \&quot;site-&lt;token&gt;\&quot;) when omitted. | [optional] 
**Name** | Pointer to **string** | Display name; defaults to \&quot;Site\&quot;. | [optional] 

## Methods

### NewProjectsDeploySiteRequest

`func NewProjectsDeploySiteRequest(files []ProjectsSiteFile, ) *ProjectsDeploySiteRequest`

NewProjectsDeploySiteRequest instantiates a new ProjectsDeploySiteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsDeploySiteRequestWithDefaults

`func NewProjectsDeploySiteRequestWithDefaults() *ProjectsDeploySiteRequest`

NewProjectsDeploySiteRequestWithDefaults instantiates a new ProjectsDeploySiteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *ProjectsDeploySiteRequest) GetFiles() []ProjectsSiteFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ProjectsDeploySiteRequest) GetFilesOk() (*[]ProjectsSiteFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ProjectsDeploySiteRequest) SetFiles(v []ProjectsSiteFile)`

SetFiles sets Files field to given value.


### GetSlug

`func (o *ProjectsDeploySiteRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsDeploySiteRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsDeploySiteRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectsDeploySiteRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetName

`func (o *ProjectsDeploySiteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsDeploySiteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsDeploySiteRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectsDeploySiteRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


