# ArtifactOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **bool** | Created is false when these exact bytes were already recorded — the write is a no-op. | [optional] 
**Ref** | Pointer to **string** | Ref is the content address, \&quot;sha256:&lt;hash&gt;\&quot;. | [optional] 
**RolledUp** | Pointer to **bool** | RolledUp is false when the OLAP roll-up was skipped; the SQLite write still stands. | [optional] 
**Sha256** | Pointer to **string** | SHA256 is the SERVER&#39;s hash of the bytes — the artifact&#39;s identity. | [optional] 

## Methods

### NewArtifactOut

`func NewArtifactOut() *ArtifactOut`

NewArtifactOut instantiates a new ArtifactOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactOutWithDefaults

`func NewArtifactOutWithDefaults() *ArtifactOut`

NewArtifactOutWithDefaults instantiates a new ArtifactOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ArtifactOut) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ArtifactOut) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ArtifactOut) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ArtifactOut) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetRef

`func (o *ArtifactOut) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ArtifactOut) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ArtifactOut) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ArtifactOut) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRolledUp

`func (o *ArtifactOut) GetRolledUp() bool`

GetRolledUp returns the RolledUp field if non-nil, zero value otherwise.

### GetRolledUpOk

`func (o *ArtifactOut) GetRolledUpOk() (*bool, bool)`

GetRolledUpOk returns a tuple with the RolledUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolledUp

`func (o *ArtifactOut) SetRolledUp(v bool)`

SetRolledUp sets RolledUp field to given value.

### HasRolledUp

`func (o *ArtifactOut) HasRolledUp() bool`

HasRolledUp returns a boolean if a field has been set.

### GetSha256

`func (o *ArtifactOut) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ArtifactOut) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ArtifactOut) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *ArtifactOut) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


