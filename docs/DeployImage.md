# DeployImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | Pointer to **string** | Digest is the sha256 of the bytes, when the reference is pinned by one. | [optional] 
**Ref** | Pointer to **string** | Ref is the reference exactly as the pods run it. | [optional] 
**Repository** | Pointer to **string** | Repository is the reference without its tag or digest. | [optional] 
**Revision** | Pointer to **string** | Revision is the commit the image was built from — its org.opencontainers.image.revision label — where the registry answers it: a digest-pinned image in one of our registries that this deployment can read. Absent means unknown, never \&quot;built from nothing\&quot;. | [optional] 
**Tag** | Pointer to **string** | Tag is the release name, when the reference carries one. | [optional] 

## Methods

### NewDeployImage

`func NewDeployImage() *DeployImage`

NewDeployImage instantiates a new DeployImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployImageWithDefaults

`func NewDeployImageWithDefaults() *DeployImage`

NewDeployImageWithDefaults instantiates a new DeployImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *DeployImage) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *DeployImage) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *DeployImage) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *DeployImage) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetRef

`func (o *DeployImage) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DeployImage) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DeployImage) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DeployImage) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRepository

`func (o *DeployImage) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DeployImage) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DeployImage) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *DeployImage) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRevision

`func (o *DeployImage) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DeployImage) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DeployImage) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *DeployImage) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetTag

`func (o *DeployImage) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DeployImage) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DeployImage) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *DeployImage) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


