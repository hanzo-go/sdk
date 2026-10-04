# PlatformProjectApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** |  | [optional] 
**Automated** | Pointer to **bool** |  | [optional] 
**Cd** | Pointer to [**PlatformCDApp**](PlatformCDApp.md) | CD is the Application reconciling it; null when CD has none for this file. | [optional] 
**Component** | Pointer to **string** |  | [optional] 
**Digest** | Pointer to **string** |  | [optional] 
**Drift** | Pointer to [**PlatformVerdict**](PlatformVerdict.md) | Drift is declared versus reconciled versus running versus latest, as flags. | [optional] 
**Env** | Pointer to [**[]PlatformDeclareEnv**](PlatformDeclareEnv.md) |  | [optional] 
**Hosts** | Pointer to **[]string** |  | [optional] 
**Image** | Pointer to **string** | Image is the image the declaration names: repository@digest when it pins a digest, else repository:tag. Empty for a declaration that owns no workload. | [optional] 
**Latest** | Pointer to **string** | Latest is the newest semver tag the registry publishes for its repository, \&quot;\&quot; when that cannot be read. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**PartOf** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to **int64** |  | [optional] 
**Repository** | Pointer to **string** |  | [optional] 
**Running** | Pointer to [**PlatformRunning**](PlatformRunning.md) | Running is what the cluster runs for it; null when no pod carries its selector or the cluster could not be read. | [optional] 
**Secrets** | Pointer to [**[]PlatformSecretRef**](PlatformSecretRef.md) |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewPlatformProjectApp

`func NewPlatformProjectApp() *PlatformProjectApp`

NewPlatformProjectApp instantiates a new PlatformProjectApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformProjectAppWithDefaults

`func NewPlatformProjectAppWithDefaults() *PlatformProjectApp`

NewPlatformProjectAppWithDefaults instantiates a new PlatformProjectApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *PlatformProjectApp) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *PlatformProjectApp) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *PlatformProjectApp) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *PlatformProjectApp) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetAutomated

`func (o *PlatformProjectApp) GetAutomated() bool`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *PlatformProjectApp) GetAutomatedOk() (*bool, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *PlatformProjectApp) SetAutomated(v bool)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *PlatformProjectApp) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetCd

`func (o *PlatformProjectApp) GetCd() PlatformCDApp`

GetCd returns the Cd field if non-nil, zero value otherwise.

### GetCdOk

`func (o *PlatformProjectApp) GetCdOk() (*PlatformCDApp, bool)`

GetCdOk returns a tuple with the Cd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCd

`func (o *PlatformProjectApp) SetCd(v PlatformCDApp)`

SetCd sets Cd field to given value.

### HasCd

`func (o *PlatformProjectApp) HasCd() bool`

HasCd returns a boolean if a field has been set.

### GetComponent

`func (o *PlatformProjectApp) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *PlatformProjectApp) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *PlatformProjectApp) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *PlatformProjectApp) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetDigest

`func (o *PlatformProjectApp) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *PlatformProjectApp) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *PlatformProjectApp) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *PlatformProjectApp) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetDrift

`func (o *PlatformProjectApp) GetDrift() PlatformVerdict`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *PlatformProjectApp) GetDriftOk() (*PlatformVerdict, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *PlatformProjectApp) SetDrift(v PlatformVerdict)`

SetDrift sets Drift field to given value.

### HasDrift

`func (o *PlatformProjectApp) HasDrift() bool`

HasDrift returns a boolean if a field has been set.

### GetEnv

`func (o *PlatformProjectApp) GetEnv() []PlatformDeclareEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *PlatformProjectApp) GetEnvOk() (*[]PlatformDeclareEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *PlatformProjectApp) SetEnv(v []PlatformDeclareEnv)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *PlatformProjectApp) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetHosts

`func (o *PlatformProjectApp) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *PlatformProjectApp) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *PlatformProjectApp) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *PlatformProjectApp) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetImage

`func (o *PlatformProjectApp) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PlatformProjectApp) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PlatformProjectApp) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PlatformProjectApp) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLatest

`func (o *PlatformProjectApp) GetLatest() string`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *PlatformProjectApp) GetLatestOk() (*string, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *PlatformProjectApp) SetLatest(v string)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *PlatformProjectApp) HasLatest() bool`

HasLatest returns a boolean if a field has been set.

### GetName

`func (o *PlatformProjectApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformProjectApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformProjectApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformProjectApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *PlatformProjectApp) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PlatformProjectApp) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PlatformProjectApp) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PlatformProjectApp) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPartOf

`func (o *PlatformProjectApp) GetPartOf() string`

GetPartOf returns the PartOf field if non-nil, zero value otherwise.

### GetPartOfOk

`func (o *PlatformProjectApp) GetPartOfOk() (*string, bool)`

GetPartOfOk returns a tuple with the PartOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartOf

`func (o *PlatformProjectApp) SetPartOf(v string)`

SetPartOf sets PartOf field to given value.

### HasPartOf

`func (o *PlatformProjectApp) HasPartOf() bool`

HasPartOf returns a boolean if a field has been set.

### GetPath

`func (o *PlatformProjectApp) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PlatformProjectApp) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PlatformProjectApp) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PlatformProjectApp) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProject

`func (o *PlatformProjectApp) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PlatformProjectApp) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PlatformProjectApp) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PlatformProjectApp) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReplicas

`func (o *PlatformProjectApp) GetReplicas() int64`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *PlatformProjectApp) GetReplicasOk() (*int64, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *PlatformProjectApp) SetReplicas(v int64)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *PlatformProjectApp) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRepository

`func (o *PlatformProjectApp) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PlatformProjectApp) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PlatformProjectApp) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PlatformProjectApp) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRunning

`func (o *PlatformProjectApp) GetRunning() PlatformRunning`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *PlatformProjectApp) GetRunningOk() (*PlatformRunning, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *PlatformProjectApp) SetRunning(v PlatformRunning)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *PlatformProjectApp) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetSecrets

`func (o *PlatformProjectApp) GetSecrets() []PlatformSecretRef`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *PlatformProjectApp) GetSecretsOk() (*[]PlatformSecretRef, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *PlatformProjectApp) SetSecrets(v []PlatformSecretRef)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *PlatformProjectApp) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetTag

`func (o *PlatformProjectApp) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PlatformProjectApp) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PlatformProjectApp) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PlatformProjectApp) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


