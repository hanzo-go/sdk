# ProjectProjectsPublish

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | Pointer to **string** | Slug is the site to publish, from the path. | [optional] 
**Source** | Pointer to **string** | Source is the build output to promote, as a path RELATIVE to your org&#39;s own storage space — never a URL and never a bucket. The org segment is prepended server-side from the validated principal, so the worst a hostile source can address is something your own org already owns. | [optional] 

## Methods

### NewProjectProjectsPublish

`func NewProjectProjectsPublish() *ProjectProjectsPublish`

NewProjectProjectsPublish instantiates a new ProjectProjectsPublish object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectsPublishWithDefaults

`func NewProjectProjectsPublishWithDefaults() *ProjectProjectsPublish`

NewProjectProjectsPublishWithDefaults instantiates a new ProjectProjectsPublish object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *ProjectProjectsPublish) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectProjectsPublish) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectProjectsPublish) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectProjectsPublish) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSource

`func (o *ProjectProjectsPublish) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProjectProjectsPublish) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProjectProjectsPublish) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ProjectProjectsPublish) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


