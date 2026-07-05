# RegistryArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Digest** | Pointer to **string** | SHA256 content digest | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**MediaType** | Pointer to **string** |  | [optional] 
**ManifestMediaType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]RegistryTag**](RegistryTag.md) |  | [optional] 
**PushTime** | Pointer to **time.Time** |  | [optional] 
**PullTime** | Pointer to **time.Time** |  | [optional] 
**ScanOverview** | Pointer to [**map[string]RegistryScanOverview**](RegistryScanOverview.md) |  | [optional] 

## Methods

### NewRegistryArtifact

`func NewRegistryArtifact() *RegistryArtifact`

NewRegistryArtifact instantiates a new RegistryArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryArtifactWithDefaults

`func NewRegistryArtifactWithDefaults() *RegistryArtifact`

NewRegistryArtifactWithDefaults instantiates a new RegistryArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegistryArtifact) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistryArtifact) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistryArtifact) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RegistryArtifact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDigest

`func (o *RegistryArtifact) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *RegistryArtifact) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *RegistryArtifact) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *RegistryArtifact) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetSize

`func (o *RegistryArtifact) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RegistryArtifact) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RegistryArtifact) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *RegistryArtifact) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMediaType

`func (o *RegistryArtifact) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *RegistryArtifact) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *RegistryArtifact) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *RegistryArtifact) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetManifestMediaType

`func (o *RegistryArtifact) GetManifestMediaType() string`

GetManifestMediaType returns the ManifestMediaType field if non-nil, zero value otherwise.

### GetManifestMediaTypeOk

`func (o *RegistryArtifact) GetManifestMediaTypeOk() (*string, bool)`

GetManifestMediaTypeOk returns a tuple with the ManifestMediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestMediaType

`func (o *RegistryArtifact) SetManifestMediaType(v string)`

SetManifestMediaType sets ManifestMediaType field to given value.

### HasManifestMediaType

`func (o *RegistryArtifact) HasManifestMediaType() bool`

HasManifestMediaType returns a boolean if a field has been set.

### GetType

`func (o *RegistryArtifact) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegistryArtifact) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegistryArtifact) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RegistryArtifact) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTags

`func (o *RegistryArtifact) GetTags() []RegistryTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RegistryArtifact) GetTagsOk() (*[]RegistryTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RegistryArtifact) SetTags(v []RegistryTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RegistryArtifact) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPushTime

`func (o *RegistryArtifact) GetPushTime() time.Time`

GetPushTime returns the PushTime field if non-nil, zero value otherwise.

### GetPushTimeOk

`func (o *RegistryArtifact) GetPushTimeOk() (*time.Time, bool)`

GetPushTimeOk returns a tuple with the PushTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushTime

`func (o *RegistryArtifact) SetPushTime(v time.Time)`

SetPushTime sets PushTime field to given value.

### HasPushTime

`func (o *RegistryArtifact) HasPushTime() bool`

HasPushTime returns a boolean if a field has been set.

### GetPullTime

`func (o *RegistryArtifact) GetPullTime() time.Time`

GetPullTime returns the PullTime field if non-nil, zero value otherwise.

### GetPullTimeOk

`func (o *RegistryArtifact) GetPullTimeOk() (*time.Time, bool)`

GetPullTimeOk returns a tuple with the PullTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullTime

`func (o *RegistryArtifact) SetPullTime(v time.Time)`

SetPullTime sets PullTime field to given value.

### HasPullTime

`func (o *RegistryArtifact) HasPullTime() bool`

HasPullTime returns a boolean if a field has been set.

### GetScanOverview

`func (o *RegistryArtifact) GetScanOverview() map[string]RegistryScanOverview`

GetScanOverview returns the ScanOverview field if non-nil, zero value otherwise.

### GetScanOverviewOk

`func (o *RegistryArtifact) GetScanOverviewOk() (*map[string]RegistryScanOverview, bool)`

GetScanOverviewOk returns a tuple with the ScanOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanOverview

`func (o *RegistryArtifact) SetScanOverview(v map[string]RegistryScanOverview)`

SetScanOverview sets ScanOverview field to given value.

### HasScanOverview

`func (o *RegistryArtifact) HasScanOverview() bool`

HasScanOverview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


