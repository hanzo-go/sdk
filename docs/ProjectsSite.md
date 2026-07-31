# ProjectsSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** |  | 
**Url** | **string** | Canonical live URL, https://&lt;slug&gt;.&lt;apex&gt;. | 
**Name** | **string** |  | 
**Status** | **string** | Always \&quot;live\&quot; in this list. | 
**UpdatedAt** | **int64** |  | 

## Methods

### NewProjectsSite

`func NewProjectsSite(slug string, url string, name string, status string, updatedAt int64, ) *ProjectsSite`

NewProjectsSite instantiates a new ProjectsSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsSiteWithDefaults

`func NewProjectsSiteWithDefaults() *ProjectsSite`

NewProjectsSiteWithDefaults instantiates a new ProjectsSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetUpdatedAt

`func (o *ProjectsSite) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectsSite) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectsSite) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


