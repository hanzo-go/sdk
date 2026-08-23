# Declared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** |  | [optional] 
**Automated** | Pointer to **bool** |  | [optional] 
**Cd** | Pointer to [**CDApp**](CDApp.md) |  | [optional] 
**Digest** | Pointer to **string** |  | [optional] 
**Env** | Pointer to [**[]DeclareEnv**](DeclareEnv.md) |  | [optional] 
**Hosts** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**Repository** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewDeclared

`func NewDeclared() *Declared`

NewDeclared instantiates a new Declared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeclaredWithDefaults

`func NewDeclaredWithDefaults() *Declared`

NewDeclaredWithDefaults instantiates a new Declared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *Declared) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *Declared) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *Declared) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *Declared) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetAutomated

`func (o *Declared) GetAutomated() bool`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *Declared) GetAutomatedOk() (*bool, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *Declared) SetAutomated(v bool)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *Declared) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetCd

`func (o *Declared) GetCd() CDApp`

GetCd returns the Cd field if non-nil, zero value otherwise.

### GetCdOk

`func (o *Declared) GetCdOk() (*CDApp, bool)`

GetCdOk returns a tuple with the Cd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCd

`func (o *Declared) SetCd(v CDApp)`

SetCd sets Cd field to given value.

### HasCd

`func (o *Declared) HasCd() bool`

HasCd returns a boolean if a field has been set.

### GetDigest

`func (o *Declared) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *Declared) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *Declared) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *Declared) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetEnv

`func (o *Declared) GetEnv() []DeclareEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *Declared) GetEnvOk() (*[]DeclareEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *Declared) SetEnv(v []DeclareEnv)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *Declared) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetHosts

`func (o *Declared) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *Declared) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *Declared) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *Declared) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetName

`func (o *Declared) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Declared) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Declared) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Declared) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *Declared) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Declared) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Declared) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Declared) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPath

`func (o *Declared) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Declared) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Declared) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Declared) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProject

`func (o *Declared) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Declared) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Declared) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Declared) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReplicas

`func (o *Declared) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *Declared) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *Declared) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *Declared) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRepository

`func (o *Declared) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *Declared) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *Declared) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *Declared) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetTag

`func (o *Declared) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Declared) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Declared) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Declared) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


