# ProjectProjectsBuildSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brief** | Pointer to **string** | Brief is what the site should be, in plain language. It is the whole input the model gets and it is size-bounded. | [optional] 
**Model** | Pointer to **string** | Model names which model writes the site. Absent takes the deployment&#39;s default — this route spends inference on the caller&#39;s org either way. | [optional] 
**Name** | Pointer to **string** | Name is the site&#39;s display name. Taken from what the model writes when omitted. | [optional] 
**Slug** | Pointer to **string** | Slug is the handle and public host label to publish under. Derived from the name, or from the brief, when omitted. | [optional] 

## Methods

### NewProjectProjectsBuildSite

`func NewProjectProjectsBuildSite() *ProjectProjectsBuildSite`

NewProjectProjectsBuildSite instantiates a new ProjectProjectsBuildSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectsBuildSiteWithDefaults

`func NewProjectProjectsBuildSiteWithDefaults() *ProjectProjectsBuildSite`

NewProjectProjectsBuildSiteWithDefaults instantiates a new ProjectProjectsBuildSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrief

`func (o *ProjectProjectsBuildSite) GetBrief() string`

GetBrief returns the Brief field if non-nil, zero value otherwise.

### GetBriefOk

`func (o *ProjectProjectsBuildSite) GetBriefOk() (*string, bool)`

GetBriefOk returns a tuple with the Brief field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrief

`func (o *ProjectProjectsBuildSite) SetBrief(v string)`

SetBrief sets Brief field to given value.

### HasBrief

`func (o *ProjectProjectsBuildSite) HasBrief() bool`

HasBrief returns a boolean if a field has been set.

### GetModel

`func (o *ProjectProjectsBuildSite) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ProjectProjectsBuildSite) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ProjectProjectsBuildSite) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ProjectProjectsBuildSite) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *ProjectProjectsBuildSite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectProjectsBuildSite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectProjectsBuildSite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectProjectsBuildSite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectProjectsBuildSite) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectProjectsBuildSite) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectProjectsBuildSite) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectProjectsBuildSite) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


