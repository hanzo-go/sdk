# CiArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewCiArtifact

`func NewCiArtifact() *CiArtifact`

NewCiArtifact instantiates a new CiArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiArtifactWithDefaults

`func NewCiArtifactWithDefaults() *CiArtifact`

NewCiArtifactWithDefaults instantiates a new CiArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *CiArtifact) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *CiArtifact) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *CiArtifact) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *CiArtifact) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetTag

`func (o *CiArtifact) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CiArtifact) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CiArtifact) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CiArtifact) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


