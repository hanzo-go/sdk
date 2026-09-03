# Declaration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** | Application is the CD Application name the generator mints: &lt;org&gt;-&lt;name&gt;. It is the join key against /v1/platform/cd. | [optional] 
**Automated** | Pointer to **bool** | Automated is cd.automated: false means the Application reports drift and NOTHING moves. It is off by default for a new file on purpose. | [optional] 
**Digest** | Pointer to **string** | image.digest — wins over tag | [optional] 
**Env** | Pointer to [**[]DeclareEnv**](DeclareEnv.md) | Env is the declared container environment, as the chart&#39;s list of {name,value}. It is read back so a re-declare of an identical body is a no-op rather than a refusal — idempotency is what makes a retry safe. | [optional] 
**Hosts** | Pointer to **[]string** | ingress.hosts, both shapes flattened | [optional] 
**Name** | Pointer to **string** | the Helm release name — the file&#39;s basename | [optional] 
**Org** | Pointer to **string** | Org is the owner. It is ALSO the values directory and the destination namespace, because those are one value under one name — see the header. | [optional] 
**Path** | Pointer to **string** | Path is the file, relative to the repository root. | [optional] 
**Project** | Pointer to **string** | Project is the AppProject the sync is admitted under, derived from the directory exactly as the ApplicationSet derives it. It differs from Org for a reserved directory, which syncs under the platform fence. | [optional] 
**Replicas** | Pointer to **int64** |  | [optional] 
**Repository** | Pointer to **string** | image.repository | [optional] 
**Tag** | Pointer to **string** | image.tag | [optional] 

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

`func (o *Declaration) GetReplicas() int64`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *Declaration) GetReplicasOk() (*int64, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *Declaration) SetReplicas(v int64)`

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


