# SignIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | Pointer to **string** | Digest is a pre-computed 32-byte digest as hex, with or without the 0x prefix. When present it is signed verbatim and message is ignored. | [optional] 
**Message** | Pointer to **string** | Message is arbitrary text to hash with Keccak256 and sign. Used only when digest is empty. | [optional] 

## Methods

### NewSignIn

`func NewSignIn() *SignIn`

NewSignIn instantiates a new SignIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignInWithDefaults

`func NewSignInWithDefaults() *SignIn`

NewSignInWithDefaults instantiates a new SignIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *SignIn) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *SignIn) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *SignIn) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *SignIn) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetMessage

`func (o *SignIn) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SignIn) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SignIn) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SignIn) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


