# Declaration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** |  | [optional] 
**Automated** | Pointer to **bool** |  | [optional] 
**Component** | Pointer to **string** |  | [optional] 
**Digest** | Pointer to **string** |  | [optional] 
**Env** | Pointer to [**[]DeclareEnv**](DeclareEnv.md) |  | [optional] 
**Hosts** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**PartOf** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**Repository** | Pointer to **string** |  | [optional] 
**Secrets** | Pointer to [**[]SecretRef**](SecretRef.md) |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewDeclaration

`func NewDeclaration() *Declaration`

NewDeclaration instantiates a new Declaration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeclarationWithDefaults

`func NewDeclarationWithDefaults() *Declaration`

NewDeclarationWithDefaults instantiates a new Declaration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *Declaration) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *Declaration) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *Declaration) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *Declaration) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetAutomated

`func (o *Declaration) GetAutomated() bool`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *Declaration) GetAutomatedOk() (*bool, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *Declaration) SetAutomated(v bool)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *Declaration) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetComponent

`func (o *Declaration) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *Declaration) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *Declaration) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *Declaration) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetDigest

`func (o *Declaration) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *Declaration) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *Declaration) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *Declaration) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetEnv

`func (o *Declaration) GetEnv() []DeclareEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *Declaration) GetEnvOk() (*[]DeclareEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *Declaration) SetEnv(v []DeclareEnv)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *Declaration) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetHosts

`func (o *Declaration) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *Declaration) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *Declaration) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *Declaration) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetName

`func (o *Declaration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Declaration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Declaration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Declaration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *Declaration) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Declaration) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Declaration) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Declaration) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPartOf

`func (o *Declaration) GetPartOf() string`

GetPartOf returns the PartOf field if non-nil, zero value otherwise.

### GetPartOfOk

`func (o *Declaration) GetPartOfOk() (*string, bool)`

GetPartOfOk returns a tuple with the PartOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartOf

`func (o *Declaration) SetPartOf(v string)`

SetPartOf sets PartOf field to given value.

### HasPartOf

`func (o *Declaration) HasPartOf() bool`

HasPartOf returns a boolean if a field has been set.

### GetPath

`func (o *Declaration) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Declaration) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Declaration) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Declaration) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProject

`func (o *Declaration) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Declaration) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Declaration) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Declaration) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReplicas

`func (o *Declaration) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *Declaration) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *Declaration) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *Declaration) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRepository

`func (o *Declaration) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *Declaration) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *Declaration) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *Declaration) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetSecrets

`func (o *Declaration) GetSecrets() []SecretRef`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *Declaration) GetSecretsOk() (*[]SecretRef, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *Declaration) SetSecrets(v []SecretRef)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *Declaration) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetTag

`func (o *Declaration) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Declaration) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Declaration) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Declaration) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


