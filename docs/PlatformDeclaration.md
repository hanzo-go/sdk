# PlatformDeclaration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** | Application is the CD Application name the generator mints from the path — the directory joined to the file name, less any words the two share (applicationName). It is how cd.hanzo.ai names the app. | [optional] 
**Automated** | Pointer to **bool** | Automated is cd.automated: false means the Application reports drift and NOTHING moves. It is off by default for a new file on purpose. | [optional] 
**Component** | Pointer to **string** | Component is the app&#39;s role within its project (&#x60;component&#x60;), \&quot;\&quot; when unnamed. | [optional] 
**Digest** | Pointer to **string** | image.digest — wins over tag | [optional] 
**Env** | Pointer to [**[]PlatformDeclareEnv**](PlatformDeclareEnv.md) | Env is the declared container environment, as the chart&#39;s list of {name,value}. It is read back so a re-declare of an identical body is a no-op rather than a refusal — idempotency is what makes a retry safe. | [optional] 
**Hosts** | Pointer to **[]string** | ingress.hosts, both shapes flattened | [optional] 
**Name** | Pointer to **string** | the Helm release name — the file&#39;s basename | [optional] 
**Org** | Pointer to **string** | Org is the owner. It is ALSO the values directory and the destination namespace, because those are one value under one name — see the header. | [optional] 
**PartOf** | Pointer to **string** | PartOf is the platform project this app belongs to: the values file&#39;s own &#x60;partOf&#x60; key, which the chart also stamps on every object it renders as app.kubernetes.io/part-of. It is the ONE declaration of membership — there is no project table — and \&quot;\&quot; means the file names none. | [optional] 
**Path** | Pointer to **string** | Path is the file, relative to the repository root. | [optional] 
**Project** | Pointer to **string** | Project is the AppProject the sync is admitted under, derived from the directory exactly as the ApplicationSet derives it. It differs from Org for a reserved directory, which syncs under the platform fence. | [optional] 
**Replicas** | Pointer to **int64** |  | [optional] 
**Repository** | Pointer to **string** | image.repository | [optional] 
**Secrets** | Pointer to [**[]PlatformSecretRef**](PlatformSecretRef.md) | Secrets are the KMS coordinates the app reads — kmsSecrets and secrets.paths — as paths and key names. A values file never carries a secret value, so neither does this. | [optional] 
**Tag** | Pointer to **string** | image.tag | [optional] 

## Methods

### NewPlatformDeclaration

`func NewPlatformDeclaration() *PlatformDeclaration`

NewPlatformDeclaration instantiates a new PlatformDeclaration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformDeclarationWithDefaults

`func NewPlatformDeclarationWithDefaults() *PlatformDeclaration`

NewPlatformDeclarationWithDefaults instantiates a new PlatformDeclaration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *PlatformDeclaration) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *PlatformDeclaration) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *PlatformDeclaration) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *PlatformDeclaration) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetAutomated

`func (o *PlatformDeclaration) GetAutomated() bool`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *PlatformDeclaration) GetAutomatedOk() (*bool, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *PlatformDeclaration) SetAutomated(v bool)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *PlatformDeclaration) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetComponent

`func (o *PlatformDeclaration) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *PlatformDeclaration) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *PlatformDeclaration) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *PlatformDeclaration) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetDigest

`func (o *PlatformDeclaration) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *PlatformDeclaration) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *PlatformDeclaration) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *PlatformDeclaration) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetEnv

`func (o *PlatformDeclaration) GetEnv() []PlatformDeclareEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *PlatformDeclaration) GetEnvOk() (*[]PlatformDeclareEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *PlatformDeclaration) SetEnv(v []PlatformDeclareEnv)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *PlatformDeclaration) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetHosts

`func (o *PlatformDeclaration) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *PlatformDeclaration) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *PlatformDeclaration) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *PlatformDeclaration) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetName

`func (o *PlatformDeclaration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformDeclaration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformDeclaration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformDeclaration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *PlatformDeclaration) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PlatformDeclaration) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PlatformDeclaration) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PlatformDeclaration) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPartOf

`func (o *PlatformDeclaration) GetPartOf() string`

GetPartOf returns the PartOf field if non-nil, zero value otherwise.

### GetPartOfOk

`func (o *PlatformDeclaration) GetPartOfOk() (*string, bool)`

GetPartOfOk returns a tuple with the PartOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartOf

`func (o *PlatformDeclaration) SetPartOf(v string)`

SetPartOf sets PartOf field to given value.

### HasPartOf

`func (o *PlatformDeclaration) HasPartOf() bool`

HasPartOf returns a boolean if a field has been set.

### GetPath

`func (o *PlatformDeclaration) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PlatformDeclaration) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PlatformDeclaration) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PlatformDeclaration) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProject

`func (o *PlatformDeclaration) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PlatformDeclaration) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PlatformDeclaration) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PlatformDeclaration) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReplicas

`func (o *PlatformDeclaration) GetReplicas() int64`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *PlatformDeclaration) GetReplicasOk() (*int64, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *PlatformDeclaration) SetReplicas(v int64)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *PlatformDeclaration) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRepository

`func (o *PlatformDeclaration) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PlatformDeclaration) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PlatformDeclaration) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PlatformDeclaration) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetSecrets

`func (o *PlatformDeclaration) GetSecrets() []PlatformSecretRef`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *PlatformDeclaration) GetSecretsOk() (*[]PlatformSecretRef, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *PlatformDeclaration) SetSecrets(v []PlatformSecretRef)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *PlatformDeclaration) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetTag

`func (o *PlatformDeclaration) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PlatformDeclaration) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PlatformDeclaration) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PlatformDeclaration) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


