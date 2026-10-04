# PlatformDeclared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** |  | [optional] 
**Automated** | Pointer to **bool** |  | [optional] 
**Cd** | Pointer to [**PlatformCDApp**](PlatformCDApp.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] 
**Digest** | Pointer to **string** |  | [optional] 
**Env** | Pointer to [**[]PlatformDeclareEnv**](PlatformDeclareEnv.md) |  | [optional] 
**Hosts** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**PartOf** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to **int64** |  | [optional] 
**Repository** | Pointer to **string** |  | [optional] 
**Secrets** | Pointer to [**[]PlatformSecretRef**](PlatformSecretRef.md) |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewPlatformDeclared

`func NewPlatformDeclared() *PlatformDeclared`

NewPlatformDeclared instantiates a new PlatformDeclared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformDeclaredWithDefaults

`func NewPlatformDeclaredWithDefaults() *PlatformDeclared`

NewPlatformDeclaredWithDefaults instantiates a new PlatformDeclared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *PlatformDeclared) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *PlatformDeclared) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *PlatformDeclared) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *PlatformDeclared) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetAutomated

`func (o *PlatformDeclared) GetAutomated() bool`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *PlatformDeclared) GetAutomatedOk() (*bool, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *PlatformDeclared) SetAutomated(v bool)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *PlatformDeclared) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetCd

`func (o *PlatformDeclared) GetCd() PlatformCDApp`

GetCd returns the Cd field if non-nil, zero value otherwise.

### GetCdOk

`func (o *PlatformDeclared) GetCdOk() (*PlatformCDApp, bool)`

GetCdOk returns a tuple with the Cd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCd

`func (o *PlatformDeclared) SetCd(v PlatformCDApp)`

SetCd sets Cd field to given value.

### HasCd

`func (o *PlatformDeclared) HasCd() bool`

HasCd returns a boolean if a field has been set.

### GetComponent

`func (o *PlatformDeclared) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *PlatformDeclared) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *PlatformDeclared) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *PlatformDeclared) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetDigest

`func (o *PlatformDeclared) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *PlatformDeclared) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *PlatformDeclared) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *PlatformDeclared) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetEnv

`func (o *PlatformDeclared) GetEnv() []PlatformDeclareEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *PlatformDeclared) GetEnvOk() (*[]PlatformDeclareEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *PlatformDeclared) SetEnv(v []PlatformDeclareEnv)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *PlatformDeclared) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetHosts

`func (o *PlatformDeclared) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *PlatformDeclared) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *PlatformDeclared) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *PlatformDeclared) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetName

`func (o *PlatformDeclared) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformDeclared) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformDeclared) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformDeclared) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *PlatformDeclared) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PlatformDeclared) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PlatformDeclared) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PlatformDeclared) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPartOf

`func (o *PlatformDeclared) GetPartOf() string`

GetPartOf returns the PartOf field if non-nil, zero value otherwise.

### GetPartOfOk

`func (o *PlatformDeclared) GetPartOfOk() (*string, bool)`

GetPartOfOk returns a tuple with the PartOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartOf

`func (o *PlatformDeclared) SetPartOf(v string)`

SetPartOf sets PartOf field to given value.

### HasPartOf

`func (o *PlatformDeclared) HasPartOf() bool`

HasPartOf returns a boolean if a field has been set.

### GetPath

`func (o *PlatformDeclared) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PlatformDeclared) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PlatformDeclared) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PlatformDeclared) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProject

`func (o *PlatformDeclared) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PlatformDeclared) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PlatformDeclared) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PlatformDeclared) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReplicas

`func (o *PlatformDeclared) GetReplicas() int64`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *PlatformDeclared) GetReplicasOk() (*int64, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *PlatformDeclared) SetReplicas(v int64)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *PlatformDeclared) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRepository

`func (o *PlatformDeclared) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PlatformDeclared) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PlatformDeclared) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PlatformDeclared) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetSecrets

`func (o *PlatformDeclared) GetSecrets() []PlatformSecretRef`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *PlatformDeclared) GetSecretsOk() (*[]PlatformSecretRef, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *PlatformDeclared) SetSecrets(v []PlatformSecretRef)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *PlatformDeclared) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetTag

`func (o *PlatformDeclared) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PlatformDeclared) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PlatformDeclared) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PlatformDeclared) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


