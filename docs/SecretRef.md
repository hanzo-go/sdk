# SecretRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to **string** |  | [optional] 
**Keys** | Pointer to **[]string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 

## Methods

### NewSecretRef

`func NewSecretRef() *SecretRef`

NewSecretRef instantiates a new SecretRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretRefWithDefaults

`func NewSecretRefWithDefaults() *SecretRef`

NewSecretRefWithDefaults instantiates a new SecretRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *SecretRef) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *SecretRef) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *SecretRef) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *SecretRef) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetKeys

`func (o *SecretRef) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *SecretRef) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *SecretRef) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *SecretRef) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetPath

`func (o *SecretRef) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SecretRef) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SecretRef) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SecretRef) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProject

`func (o *SecretRef) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SecretRef) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SecretRef) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *SecretRef) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


