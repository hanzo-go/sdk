# ProjectsBuildSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brief** | Pointer to **string** | Brief is what the site should be, in plain language. It is the whole input the model gets and it is size-bounded. | [optional] 
**Model** | Pointer to **string** | Model names which model writes the site. Absent takes the deployment&#39;s default — this route spends inference on the caller&#39;s org either way. | [optional] 
**Name** | Pointer to **string** | Name is the site&#39;s display name. Taken from what the model writes when omitted. | [optional] 
**Slug** | Pointer to **string** | Slug is the handle and public host label to publish under. Derived from the name, or from the brief, when omitted. | [optional] 

## Methods

### NewProjectsBuildSite

`func NewProjectsBuildSite() *ProjectsBuildSite`

NewProjectsBuildSite instantiates a new ProjectsBuildSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsBuildSiteWithDefaults

`func NewProjectsBuildSiteWithDefaults() *ProjectsBuildSite`

NewProjectsBuildSiteWithDefaults instantiates a new ProjectsBuildSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrief

`func (o *ProjectsBuildSite) GetBrief() string`

GetBrief returns the Brief field if non-nil, zero value otherwise.

### GetBriefOk

`func (o *ProjectsBuildSite) GetBriefOk() (*string, bool)`

GetBriefOk returns a tuple with the Brief field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrief

`func (o *ProjectsBuildSite) SetBrief(v string)`

SetBrief sets Brief field to given value.

### HasBrief

`func (o *ProjectsBuildSite) HasBrief() bool`

HasBrief returns a boolean if a field has been set.

### GetModel

`func (o *ProjectsBuildSite) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ProjectsBuildSite) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ProjectsBuildSite) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ProjectsBuildSite) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *ProjectsBuildSite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsBuildSite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsBuildSite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectsBuildSite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectsBuildSite) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsBuildSite) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsBuildSite) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectsBuildSite) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


