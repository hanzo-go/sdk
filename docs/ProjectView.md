# ProjectView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to **int32** | Applications is how many platform apps this org has under the project, counted per request. It is the one fact IAM cannot answer about a project. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is IAM&#39;s creation time as unix seconds. 0 when IAM&#39;s timestamp is absent or unparseable — never a fabricated time. | [optional] 
**Description** | Pointer to **string** | Description is IAM&#39;s free text about the project. Nothing derives from it. | [optional] 
**Name** | Pointer to **string** | Name is IAM&#39;s display name, falling back to the slug when the project has none, so this is never empty. | [optional] 
**Org** | Pointer to **string** | Org is the project&#39;s IAM owner, and the tenant every app under it deploys into. It comes from the validated identity, never from the request. | [optional] 
**Slug** | Pointer to **string** | Slug is the project&#39;s IAM name — half of the (org,name) identity, the &#x60;:project&#x60; path segment, and the scope key an app is filed under. It is the project&#39;s address; Name is not. | [optional] 

## Methods

### NewProjectView

`func NewProjectView() *ProjectView`

NewProjectView instantiates a new ProjectView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectViewWithDefaults

`func NewProjectViewWithDefaults() *ProjectView`

NewProjectViewWithDefaults instantiates a new ProjectView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *ProjectView) GetApplications() int32`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ProjectView) GetApplicationsOk() (*int32, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ProjectView) SetApplications(v int32)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *ProjectView) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProjectView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ProjectView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *ProjectView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *ProjectView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *ProjectView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *ProjectView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectView) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectView) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectView) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectView) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


