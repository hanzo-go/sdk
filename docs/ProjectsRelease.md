# ProjectsRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Active is whether this is the release the site is SERVING right now. Exactly one release of a site is active; the others are kept so they can be activated again, until retention reclaims them. | [optional] 
**Bytes** | Pointer to **int32** | Bytes is their total size in bytes. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the release was cut, as Unix seconds — not when it was last activated. | [optional] 
**Objects** | Pointer to **int32** | Objects is how many files the release holds. | [optional] 
**ReleaseId** | Pointer to **string** | ReleaseID is derived from a DIGEST of the release&#39;s own manifest, so identical content is the same release and a release can never be confused with another one. Activating an older id IS the rollback. | [optional] 
**Slug** | Pointer to **string** | Slug is the site this release belongs to. | [optional] 
**Source** | Pointer to **string** | Source is what the release was cut from — the build output or upload it was promoted out of. | [optional] 
**Url** | Pointer to **string** | URL is where the site serves. Present only on the ACTIVE release, since an inactive one is not answering anywhere. | [optional] 

## Methods

### NewProjectsRelease

`func NewProjectsRelease() *ProjectsRelease`

NewProjectsRelease instantiates a new ProjectsRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsReleaseWithDefaults

`func NewProjectsReleaseWithDefaults() *ProjectsRelease`

NewProjectsReleaseWithDefaults instantiates a new ProjectsRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ProjectsRelease) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ProjectsRelease) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ProjectsRelease) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ProjectsRelease) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBytes

`func (o *ProjectsRelease) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ProjectsRelease) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ProjectsRelease) SetBytes(v int32)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *ProjectsRelease) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProjectsRelease) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectsRelease) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectsRelease) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectsRelease) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetObjects

`func (o *ProjectsRelease) GetObjects() int32`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ProjectsRelease) GetObjectsOk() (*int32, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ProjectsRelease) SetObjects(v int32)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ProjectsRelease) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetReleaseId

`func (o *ProjectsRelease) GetReleaseId() string`

GetReleaseId returns the ReleaseId field if non-nil, zero value otherwise.

### GetReleaseIdOk

`func (o *ProjectsRelease) GetReleaseIdOk() (*string, bool)`

GetReleaseIdOk returns a tuple with the ReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseId

`func (o *ProjectsRelease) SetReleaseId(v string)`

SetReleaseId sets ReleaseId field to given value.

### HasReleaseId

`func (o *ProjectsRelease) HasReleaseId() bool`

HasReleaseId returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectsRelease) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectsRelease) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectsRelease) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectsRelease) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSource

`func (o *ProjectsRelease) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProjectsRelease) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProjectsRelease) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ProjectsRelease) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUrl

`func (o *ProjectsRelease) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProjectsRelease) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProjectsRelease) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProjectsRelease) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


