# ProjectProjectsRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Active is whether this is the release the site is SERVING right now. Exactly one release of a site is active; the others are kept so they can be activated again, until retention reclaims them. | [optional] 
**Bytes** | Pointer to **int64** | Bytes is their total size in bytes. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the release was cut, as Unix seconds — not when it was last activated. | [optional] 
**Objects** | Pointer to **int64** | Objects is how many files the release holds. | [optional] 
**ReleaseId** | Pointer to **string** | ReleaseID is derived from a DIGEST of the release&#39;s own manifest, so identical content is the same release and a release can never be confused with another one. Activating an older id IS the rollback. | [optional] 
**Slug** | Pointer to **string** | Slug is the site this release belongs to. | [optional] 
**Source** | Pointer to **string** | Source is what the release was cut from — the build output or upload it was promoted out of. | [optional] 
**Url** | Pointer to **string** | URL is where the site serves. Present only on the ACTIVE release, since an inactive one is not answering anywhere. | [optional] 

## Methods

### NewProjectProjectsRelease

`func NewProjectProjectsRelease() *ProjectProjectsRelease`

NewProjectProjectsRelease instantiates a new ProjectProjectsRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectsReleaseWithDefaults

`func NewProjectProjectsReleaseWithDefaults() *ProjectProjectsRelease`

NewProjectProjectsReleaseWithDefaults instantiates a new ProjectProjectsRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ProjectProjectsRelease) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ProjectProjectsRelease) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ProjectProjectsRelease) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ProjectProjectsRelease) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBytes

`func (o *ProjectProjectsRelease) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ProjectProjectsRelease) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ProjectProjectsRelease) SetBytes(v int64)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *ProjectProjectsRelease) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProjectProjectsRelease) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectProjectsRelease) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectProjectsRelease) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectProjectsRelease) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetObjects

`func (o *ProjectProjectsRelease) GetObjects() int64`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ProjectProjectsRelease) GetObjectsOk() (*int64, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ProjectProjectsRelease) SetObjects(v int64)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ProjectProjectsRelease) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetReleaseId

`func (o *ProjectProjectsRelease) GetReleaseId() string`

GetReleaseId returns the ReleaseId field if non-nil, zero value otherwise.

### GetReleaseIdOk

`func (o *ProjectProjectsRelease) GetReleaseIdOk() (*string, bool)`

GetReleaseIdOk returns a tuple with the ReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseId

`func (o *ProjectProjectsRelease) SetReleaseId(v string)`

SetReleaseId sets ReleaseId field to given value.

### HasReleaseId

`func (o *ProjectProjectsRelease) HasReleaseId() bool`

HasReleaseId returns a boolean if a field has been set.

### GetSlug

`func (o *ProjectProjectsRelease) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProjectProjectsRelease) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProjectProjectsRelease) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProjectProjectsRelease) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSource

`func (o *ProjectProjectsRelease) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProjectProjectsRelease) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProjectProjectsRelease) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ProjectProjectsRelease) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUrl

`func (o *ProjectProjectsRelease) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProjectProjectsRelease) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProjectProjectsRelease) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProjectProjectsRelease) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


