# CloudArtifactOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **bool** | Created is false when these exact bytes were already recorded — the write is a no-op. | [optional] 
**Ref** | Pointer to **string** | Ref is the content address, \&quot;sha256:&lt;hash&gt;\&quot;. | [optional] 
**RolledUp** | Pointer to **bool** | RolledUp is false when the OLAP roll-up was skipped; the SQLite write still stands. | [optional] 
**Sha256** | Pointer to **string** | SHA256 is the SERVER&#39;s hash of the bytes — the artifact&#39;s identity. | [optional] 

## Methods

### NewCloudArtifactOut

`func NewCloudArtifactOut() *CloudArtifactOut`

NewCloudArtifactOut instantiates a new CloudArtifactOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudArtifactOutWithDefaults

`func NewCloudArtifactOutWithDefaults() *CloudArtifactOut`

NewCloudArtifactOutWithDefaults instantiates a new CloudArtifactOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CloudArtifactOut) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CloudArtifactOut) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CloudArtifactOut) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CloudArtifactOut) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetRef

`func (o *CloudArtifactOut) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CloudArtifactOut) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CloudArtifactOut) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CloudArtifactOut) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRolledUp

`func (o *CloudArtifactOut) GetRolledUp() bool`

GetRolledUp returns the RolledUp field if non-nil, zero value otherwise.

### GetRolledUpOk

`func (o *CloudArtifactOut) GetRolledUpOk() (*bool, bool)`

GetRolledUpOk returns a tuple with the RolledUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolledUp

`func (o *CloudArtifactOut) SetRolledUp(v bool)`

SetRolledUp sets RolledUp field to given value.

### HasRolledUp

`func (o *CloudArtifactOut) HasRolledUp() bool`

HasRolledUp returns a boolean if a field has been set.

### GetSha256

`func (o *CloudArtifactOut) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *CloudArtifactOut) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *CloudArtifactOut) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *CloudArtifactOut) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


