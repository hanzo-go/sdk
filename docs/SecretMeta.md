# SecretMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to **string** | Env is the environment the secret belongs to. It is part of the storage key, so the same name in two environments is two secrets. | [optional] 
**Name** | Pointer to **string** | Name is the secret&#39;s name within its path and environment. | [optional] 
**Path** | Pointer to **string** | Path is the subpath the secret is stored under, beneath the org root. | [optional] 
**Scheme** | Pointer to **string** | Scheme names how the value is sealed at rest, so a caller can tell a migrated record from a current one without opening it. | [optional] 

## Methods

### NewSecretMeta

`func NewSecretMeta() *SecretMeta`

NewSecretMeta instantiates a new SecretMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretMetaWithDefaults

`func NewSecretMetaWithDefaults() *SecretMeta`

NewSecretMetaWithDefaults instantiates a new SecretMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *SecretMeta) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *SecretMeta) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *SecretMeta) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *SecretMeta) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetName

`func (o *SecretMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecretMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecretMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecretMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *SecretMeta) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SecretMeta) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SecretMeta) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SecretMeta) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetScheme

`func (o *SecretMeta) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *SecretMeta) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *SecretMeta) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *SecretMeta) HasScheme() bool`

HasScheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


