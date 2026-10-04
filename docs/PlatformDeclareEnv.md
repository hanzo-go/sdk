# PlatformDeclareEnv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Public** | Pointer to **bool** | Public marks a value that may be WRITTEN INTO GIT. Absent, it is false, and the value is sealed into KMS and referenced.  ★ THE DEFAULT IS SECRET, AND THE POLARITY IS THE WHOLE DESIGN. This lane&#39;s output is a commit in a repository replicated to every clone, so a value written there stays there. Sealing by default means the worst outcome of a wrong guess is a config value an operator reads back from KMS rather than from git, and nothing sensitive ever depends on recognising its shape.  It is also the only rule that needs no list: a password, a kubeconfig or an MFA seed is sealed whatever it looks like, because only an explicit public mark writes a value into git. | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewPlatformDeclareEnv

`func NewPlatformDeclareEnv() *PlatformDeclareEnv`

NewPlatformDeclareEnv instantiates a new PlatformDeclareEnv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformDeclareEnvWithDefaults

`func NewPlatformDeclareEnvWithDefaults() *PlatformDeclareEnv`

NewPlatformDeclareEnvWithDefaults instantiates a new PlatformDeclareEnv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlatformDeclareEnv) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformDeclareEnv) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformDeclareEnv) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformDeclareEnv) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublic

`func (o *PlatformDeclareEnv) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *PlatformDeclareEnv) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *PlatformDeclareEnv) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *PlatformDeclareEnv) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetValue

`func (o *PlatformDeclareEnv) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PlatformDeclareEnv) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PlatformDeclareEnv) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PlatformDeclareEnv) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


