# O11yO11ySentryPostableProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name is the project&#39;s display name. Required. | 
**Platform** | Pointer to **string** | Platform is the reporting runtime, e.g. go, python, javascript. | [optional] 
**Slug** | Pointer to **string** | Slug is the project&#39;s short name. Server-assigned from Name when empty. | [optional] 

## Methods

### NewO11yO11ySentryPostableProject

`func NewO11yO11ySentryPostableProject(name string, ) *O11yO11ySentryPostableProject`

NewO11yO11ySentryPostableProject instantiates a new O11yO11ySentryPostableProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11ySentryPostableProjectWithDefaults

`func NewO11yO11ySentryPostableProjectWithDefaults() *O11yO11ySentryPostableProject`

NewO11yO11ySentryPostableProjectWithDefaults instantiates a new O11yO11ySentryPostableProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *O11yO11ySentryPostableProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11ySentryPostableProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11ySentryPostableProject) SetName(v string)`

SetName sets Name field to given value.


### GetPlatform

`func (o *O11yO11ySentryPostableProject) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *O11yO11ySentryPostableProject) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *O11yO11ySentryPostableProject) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *O11yO11ySentryPostableProject) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSlug

`func (o *O11yO11ySentryPostableProject) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *O11yO11ySentryPostableProject) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *O11yO11ySentryPostableProject) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *O11yO11ySentryPostableProject) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


