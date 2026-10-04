# KmsKmsPut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to **string** | Env is the environment to write under. REQUIRED, with no default: it is part of the storage key, so a silently defaulted write lands in a bucket the readers that resolve project, environment and path never look in, and the stale value keeps being served. | [optional] 
**Name** | Pointer to **string** | Name is the secret&#39;s name. Required. | [optional] 
**Path** | Pointer to **string** | Path is an optional subpath beneath the org root, e.g. \&quot;/ci\&quot;. | [optional] 
**Value** | Pointer to **string** | Value is the secret itself. It is sealed under a fresh per-secret data key before storage, so plaintext never reaches disk, and it is never echoed back, logged, or carried in an error. | [optional] 

## Methods

### NewKmsKmsPut

`func NewKmsKmsPut() *KmsKmsPut`

NewKmsKmsPut instantiates a new KmsKmsPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsKmsPutWithDefaults

`func NewKmsKmsPutWithDefaults() *KmsKmsPut`

NewKmsKmsPutWithDefaults instantiates a new KmsKmsPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *KmsKmsPut) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *KmsKmsPut) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *KmsKmsPut) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *KmsKmsPut) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetName

`func (o *KmsKmsPut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsKmsPut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsKmsPut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KmsKmsPut) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *KmsKmsPut) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *KmsKmsPut) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *KmsKmsPut) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *KmsKmsPut) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetValue

`func (o *KmsKmsPut) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KmsKmsPut) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KmsKmsPut) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *KmsKmsPut) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


