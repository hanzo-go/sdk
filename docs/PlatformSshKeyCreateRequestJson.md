# PlatformSshKeyCreateRequestJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**PrivateKey** | **string** |  | 
**PublicKey** | Pointer to **string** |  | [optional] 

## Methods

### NewPlatformSshKeyCreateRequestJson

`func NewPlatformSshKeyCreateRequestJson(name string, privateKey string, ) *PlatformSshKeyCreateRequestJson`

NewPlatformSshKeyCreateRequestJson instantiates a new PlatformSshKeyCreateRequestJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformSshKeyCreateRequestJsonWithDefaults

`func NewPlatformSshKeyCreateRequestJsonWithDefaults() *PlatformSshKeyCreateRequestJson`

NewPlatformSshKeyCreateRequestJsonWithDefaults instantiates a new PlatformSshKeyCreateRequestJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlatformSshKeyCreateRequestJson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformSshKeyCreateRequestJson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformSshKeyCreateRequestJson) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PlatformSshKeyCreateRequestJson) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformSshKeyCreateRequestJson) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformSshKeyCreateRequestJson) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformSshKeyCreateRequestJson) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrivateKey

`func (o *PlatformSshKeyCreateRequestJson) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PlatformSshKeyCreateRequestJson) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PlatformSshKeyCreateRequestJson) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetPublicKey

`func (o *PlatformSshKeyCreateRequestJson) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *PlatformSshKeyCreateRequestJson) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *PlatformSshKeyCreateRequestJson) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *PlatformSshKeyCreateRequestJson) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


