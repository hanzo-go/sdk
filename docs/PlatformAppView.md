# PlatformAppView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | service / CR name, e.g. iam | [optional] 
**Cluster** | Pointer to **string** | always \&quot;hanzo-k8s\&quot;; cross-cluster federation is a follow-up phase | [optional] 
**DeclaredTag** | Pointer to **string** | spec.image.tag — the tag the CR SAYS to run; anything but vX.Y.Z is red drift | [optional] 
**Drift** | Pointer to [**PlatformVerdict**](PlatformVerdict.md) | declared vs running vs latest as flags, plus their rolled-up severity | [optional] 
**Endpoints** | Pointer to **[]string** | status.endpoints the operator published for this service; empty until it reconciles | [optional] 
**Env** | Pointer to **string** | main|test|dev | [optional] 
**Health** | Pointer to **string** | green|yellow|red|\&quot;\&quot; (unknown) | [optional] 
**Id** | Pointer to **string** | &lt;org&gt;/&lt;app&gt;/&lt;env&gt;, e.g. hanzoai/iam/main | [optional] 
**LatestTag** | Pointer to **string** | newest released tag; \&quot;\&quot; until the GH release reader lands, so \&quot;stale\&quot; cannot fire yet | [optional] 
**Namespace** | Pointer to **string** | namespace the row was scanned from: a platform one (hanzo, hanzo-testnet, …) or a tenant-&lt;org&gt; | [optional] 
**Org** | Pointer to **string** | image namespace, e.g. hanzoai | [optional] 
**Phase** | Pointer to **string** | operator status.phase (Running/…) | [optional] 
**Registry** | Pointer to **string** | spec.image.repository verbatim (ghcr.io/hanzoai/iam); Org and Repo are read off it | [optional] 
**Repo** | Pointer to **string** | owner/repo, e.g. hanzoai/iam | [optional] 
**Role** | Pointer to **string** | operator spec.role (sql|kv|generic|ingress|…) or \&quot;\&quot; — the one declared class field | [optional] 
**RunningTag** | Pointer to **string** | the tag the live Deployment actually runs; \&quot;\&quot; when unreadable — unknown, not a guess | [optional] 

## Methods

### NewPlatformAppView

`func NewPlatformAppView() *PlatformAppView`

NewPlatformAppView instantiates a new PlatformAppView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformAppViewWithDefaults

`func NewPlatformAppViewWithDefaults() *PlatformAppView`

NewPlatformAppViewWithDefaults instantiates a new PlatformAppView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *PlatformAppView) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *PlatformAppView) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *PlatformAppView) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *PlatformAppView) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetCluster

`func (o *PlatformAppView) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PlatformAppView) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PlatformAppView) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PlatformAppView) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDeclaredTag

`func (o *PlatformAppView) GetDeclaredTag() string`

GetDeclaredTag returns the DeclaredTag field if non-nil, zero value otherwise.

### GetDeclaredTagOk

`func (o *PlatformAppView) GetDeclaredTagOk() (*string, bool)`

GetDeclaredTagOk returns a tuple with the DeclaredTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredTag

`func (o *PlatformAppView) SetDeclaredTag(v string)`

SetDeclaredTag sets DeclaredTag field to given value.

### HasDeclaredTag

`func (o *PlatformAppView) HasDeclaredTag() bool`

HasDeclaredTag returns a boolean if a field has been set.

### GetDrift

`func (o *PlatformAppView) GetDrift() PlatformVerdict`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *PlatformAppView) GetDriftOk() (*PlatformVerdict, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *PlatformAppView) SetDrift(v PlatformVerdict)`

SetDrift sets Drift field to given value.

### HasDrift

`func (o *PlatformAppView) HasDrift() bool`

HasDrift returns a boolean if a field has been set.

### GetEndpoints

`func (o *PlatformAppView) GetEndpoints() []string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *PlatformAppView) GetEndpointsOk() (*[]string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *PlatformAppView) SetEndpoints(v []string)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *PlatformAppView) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetEnv

`func (o *PlatformAppView) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *PlatformAppView) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *PlatformAppView) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *PlatformAppView) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetHealth

`func (o *PlatformAppView) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *PlatformAppView) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *PlatformAppView) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *PlatformAppView) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetId

`func (o *PlatformAppView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlatformAppView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlatformAppView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlatformAppView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatestTag

`func (o *PlatformAppView) GetLatestTag() string`

GetLatestTag returns the LatestTag field if non-nil, zero value otherwise.

### GetLatestTagOk

`func (o *PlatformAppView) GetLatestTagOk() (*string, bool)`

GetLatestTagOk returns a tuple with the LatestTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTag

`func (o *PlatformAppView) SetLatestTag(v string)`

SetLatestTag sets LatestTag field to given value.

### HasLatestTag

`func (o *PlatformAppView) HasLatestTag() bool`

HasLatestTag returns a boolean if a field has been set.

### GetNamespace

`func (o *PlatformAppView) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PlatformAppView) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PlatformAppView) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PlatformAppView) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOrg

`func (o *PlatformAppView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PlatformAppView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PlatformAppView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PlatformAppView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPhase

`func (o *PlatformAppView) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *PlatformAppView) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *PlatformAppView) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *PlatformAppView) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetRegistry

`func (o *PlatformAppView) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *PlatformAppView) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *PlatformAppView) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *PlatformAppView) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetRepo

`func (o *PlatformAppView) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *PlatformAppView) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *PlatformAppView) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *PlatformAppView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRole

`func (o *PlatformAppView) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PlatformAppView) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PlatformAppView) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PlatformAppView) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRunningTag

`func (o *PlatformAppView) GetRunningTag() string`

GetRunningTag returns the RunningTag field if non-nil, zero value otherwise.

### GetRunningTagOk

`func (o *PlatformAppView) GetRunningTagOk() (*string, bool)`

GetRunningTagOk returns a tuple with the RunningTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningTag

`func (o *PlatformAppView) SetRunningTag(v string)`

SetRunningTag sets RunningTag field to given value.

### HasRunningTag

`func (o *PlatformAppView) HasRunningTag() bool`

HasRunningTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


