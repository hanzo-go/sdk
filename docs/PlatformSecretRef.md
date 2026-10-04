# PlatformSecretRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to **string** | Env is the KMS environment slug. | [optional] 
**Keys** | Pointer to **[]string** | Keys are the names read from the path; empty means the whole path. | [optional] 
**Path** | Pointer to **string** | Path is the secret path, always absolute. | [optional] 
**Project** | Pointer to **string** | Project is the KMS project slug the path lives in. | [optional] 

## Methods

### NewPlatformSecretRef

`func NewPlatformSecretRef() *PlatformSecretRef`

NewPlatformSecretRef instantiates a new PlatformSecretRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformSecretRefWithDefaults

`func NewPlatformSecretRefWithDefaults() *PlatformSecretRef`

NewPlatformSecretRefWithDefaults instantiates a new PlatformSecretRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *PlatformSecretRef) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *PlatformSecretRef) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *PlatformSecretRef) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *PlatformSecretRef) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetKeys

`func (o *PlatformSecretRef) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *PlatformSecretRef) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *PlatformSecretRef) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *PlatformSecretRef) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetPath

`func (o *PlatformSecretRef) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PlatformSecretRef) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PlatformSecretRef) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PlatformSecretRef) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProject

`func (o *PlatformSecretRef) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PlatformSecretRef) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PlatformSecretRef) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PlatformSecretRef) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


