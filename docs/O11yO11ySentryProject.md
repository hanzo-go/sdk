# O11yO11ySentryProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when the project was created. | [optional] 
**Dsn** | Pointer to **string** | DSN is the project&#39;s freshly-derived ingest DSN. | [optional] 
**Id** | Pointer to **string** | ID is the project id. | [optional] 
**Name** | Pointer to **string** | Name is the project&#39;s display name. | [optional] 
**Platform** | Pointer to **string** | Platform is the reporting runtime, e.g. go, python, javascript. | [optional] 
**Slug** | Pointer to **string** | Slug is the project&#39;s short name. | [optional] 
**Status** | Pointer to **string** | Status is the project&#39;s lifecycle state: active or disabled. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when the project last changed. | [optional] 

## Methods

### NewO11yO11ySentryProject

`func NewO11yO11ySentryProject() *O11yO11ySentryProject`

NewO11yO11ySentryProject instantiates a new O11yO11ySentryProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11ySentryProjectWithDefaults

`func NewO11yO11ySentryProjectWithDefaults() *O11yO11ySentryProject`

NewO11yO11ySentryProjectWithDefaults instantiates a new O11yO11ySentryProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yO11ySentryProject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11ySentryProject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11ySentryProject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11ySentryProject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDsn

`func (o *O11yO11ySentryProject) GetDsn() string`

GetDsn returns the Dsn field if non-nil, zero value otherwise.

### GetDsnOk

`func (o *O11yO11ySentryProject) GetDsnOk() (*string, bool)`

GetDsnOk returns a tuple with the Dsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsn

`func (o *O11yO11ySentryProject) SetDsn(v string)`

SetDsn sets Dsn field to given value.

### HasDsn

`func (o *O11yO11ySentryProject) HasDsn() bool`

HasDsn returns a boolean if a field has been set.

### GetId

`func (o *O11yO11ySentryProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11ySentryProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11ySentryProject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11ySentryProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *O11yO11ySentryProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11ySentryProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11ySentryProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11ySentryProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *O11yO11ySentryProject) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *O11yO11ySentryProject) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *O11yO11ySentryProject) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *O11yO11ySentryProject) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSlug

`func (o *O11yO11ySentryProject) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *O11yO11ySentryProject) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *O11yO11ySentryProject) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *O11yO11ySentryProject) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11ySentryProject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11ySentryProject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11ySentryProject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11ySentryProject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11ySentryProject) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11ySentryProject) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11ySentryProject) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11ySentryProject) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


