# ProjectsSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the site&#39;s display name. | [optional] 
**Slug** | Pointer to **string** | Slug is the site&#39;s handle — also the label of the host it serves at. | [optional] 
**Status** | Pointer to **string** | Status is the project&#39;s state behind the site — whether it is serving, still building, or failed its last build. A site that is listed is not necessarily one that answers. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is when the project last changed, as Unix seconds. | [optional] 
**Url** | Pointer to **string** | URL is the pretty address readers use, not the object-store path behind it. | [optional] 

## Methods

### NewProjectsSite

`func NewProjectsSite() *ProjectsSite`

NewProjectsSite instantiates a new ProjectsSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsSiteWithDefaults

`func NewProjectsSiteWithDefaults() *ProjectsSite`

NewProjectsSiteWithDefaults instantiates a new ProjectsSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectsSite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsSite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsSite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectsSite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectsSite) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsSite) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsSite) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectsSite) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStatus

`func (o *ProjectsSite) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectsSite) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectsSite) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProjectsSite) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectsSite) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectsSite) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectsSite) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectsSite) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *ProjectsSite) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProjectsSite) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProjectsSite) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProjectsSite) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


