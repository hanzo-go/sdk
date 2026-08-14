# ProjectsDeployStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | Pointer to **string** | Commit is the git sha this build was produced from, recorded on the deployment so a released site can be traced back to its source. Optional.  It is the ONLY field here, and deliberately: the predecessor also accepted &#x60;source&#x60; and &#x60;branch&#x60;. &#x60;source&#x60; was the Content-Type discriminator this split removed. &#x60;branch&#x60; was accepted and DISCARDED — there is no branch column on a deployment, and the lifecycle event derives the branch from the project&#39;s own linked one — so declaring it would publish a settable field that does nothing into the document, every generated SDK and the MCP input schema. A field that is read by nothing is not described as if it were.  &#x60;url:\&quot;-\&quot;&#x60; because zip binds the query string OVER a decoded body, so without it a &#x60;?commit&#x3D;&#x60; the caller never sent would outrank the one it did. | [optional] 
**Slug** | Pointer to **string** | Slug is the site to deploy, from the path. | [optional] 

## Methods

### NewProjectsDeployStart

`func NewProjectsDeployStart() *ProjectsDeployStart`

NewProjectsDeployStart instantiates a new ProjectsDeployStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsDeployStartWithDefaults

`func NewProjectsDeployStartWithDefaults() *ProjectsDeployStart`

NewProjectsDeployStartWithDefaults instantiates a new ProjectsDeployStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *ProjectsDeployStart) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ProjectsDeployStart) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ProjectsDeployStart) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *ProjectsDeployStart) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectsDeployStart) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsDeployStart) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsDeployStart) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectsDeployStart) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


