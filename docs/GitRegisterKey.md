# GitRegisterKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Human label; falls back to the key comment | [optional] 
**PublicKey** | **string** | OpenSSH authorized-key line | 

## Methods

### NewGitRegisterKey

`func NewGitRegisterKey(publicKey string, ) *GitRegisterKey`

NewGitRegisterKey instantiates a new GitRegisterKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRegisterKeyWithDefaults

`func NewGitRegisterKeyWithDefaults() *GitRegisterKey`

NewGitRegisterKeyWithDefaults instantiates a new GitRegisterKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *GitRegisterKey) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GitRegisterKey) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GitRegisterKey) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GitRegisterKey) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPublicKey

`func (o *GitRegisterKey) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *GitRegisterKey) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *GitRegisterKey) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


