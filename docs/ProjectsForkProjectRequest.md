# ProjectsForkProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** | Template slug to fork (required). | 
**Name** | Pointer to **string** | Target project name; defaults to the template title. | [optional] 
**Target** | Pointer to **string** | Overrides the derived project slug; defaults to the template slug. | [optional] 

## Methods

### NewProjectsForkProjectRequest

`func NewProjectsForkProjectRequest(slug string, ) *ProjectsForkProjectRequest`

NewProjectsForkProjectRequest instantiates a new ProjectsForkProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsForkProjectRequestWithDefaults

`func NewProjectsForkProjectRequestWithDefaults() *ProjectsForkProjectRequest`

NewProjectsForkProjectRequestWithDefaults instantiates a new ProjectsForkProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *ProjectsForkProjectRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsForkProjectRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsForkProjectRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *ProjectsForkProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsForkProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsForkProjectRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectsForkProjectRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTarget

`func (o *ProjectsForkProjectRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ProjectsForkProjectRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ProjectsForkProjectRequest) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ProjectsForkProjectRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


