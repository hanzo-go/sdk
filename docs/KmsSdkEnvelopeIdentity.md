# KmsSdkEnvelopeIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scheme** | **int32** | NodeID scheme. 0x42 &#x3D; ML-DSA-65. | 
**Node** | **string** | 20-byte canonical NodeID (cb58), the SHAKE256-384 digest prefix. | 
**Digest** | **string** | 48-byte SHAKE256-384 FullDigest commitment (base64). | 
**Path** | **string** | BIP-44 service path the identity was derived from (e.g. hanzo/kms-operator). | 
**Pubkey** | **string** | ML-DSA-65 public key (base64), for offline verification. | 

## Methods

### NewKmsSdkEnvelopeIdentity

`func NewKmsSdkEnvelopeIdentity(scheme int32, node string, digest string, path string, pubkey string, ) *KmsSdkEnvelopeIdentity`

NewKmsSdkEnvelopeIdentity instantiates a new KmsSdkEnvelopeIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsSdkEnvelopeIdentityWithDefaults

`func NewKmsSdkEnvelopeIdentityWithDefaults() *KmsSdkEnvelopeIdentity`

NewKmsSdkEnvelopeIdentityWithDefaults instantiates a new KmsSdkEnvelopeIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheme

`func (o *KmsSdkEnvelopeIdentity) GetScheme() int32`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *KmsSdkEnvelopeIdentity) GetSchemeOk() (*int32, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *KmsSdkEnvelopeIdentity) SetScheme(v int32)`

SetScheme sets Scheme field to given value.


### GetNode

`func (o *KmsSdkEnvelopeIdentity) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *KmsSdkEnvelopeIdentity) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *KmsSdkEnvelopeIdentity) SetNode(v string)`

SetNode sets Node field to given value.


### GetDigest

`func (o *KmsSdkEnvelopeIdentity) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *KmsSdkEnvelopeIdentity) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *KmsSdkEnvelopeIdentity) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetPath

`func (o *KmsSdkEnvelopeIdentity) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *KmsSdkEnvelopeIdentity) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *KmsSdkEnvelopeIdentity) SetPath(v string)`

SetPath sets Path field to given value.


### GetPubkey

`func (o *KmsSdkEnvelopeIdentity) GetPubkey() string`

GetPubkey returns the Pubkey field if non-nil, zero value otherwise.

### GetPubkeyOk

`func (o *KmsSdkEnvelopeIdentity) GetPubkeyOk() (*string, bool)`

GetPubkeyOk returns a tuple with the Pubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubkey

`func (o *KmsSdkEnvelopeIdentity) SetPubkey(v string)`

SetPubkey sets Pubkey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


