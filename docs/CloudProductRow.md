# CloudProductRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to **string** | hanzo-k8s | [optional] 
**DeclaredTag** | Pointer to **string** | spec.image.tag on the App CR (declared truth) | [optional] 
**Drift** | Pointer to **bool** | any drift flag present | [optional] 
**DriftSeverity** | Pointer to **string** | ok|yellow|red (rolled-up) | [optional] 
**Env** | Pointer to **string** | main|test|dev (lifecycle namespace) | [optional] 
**Health** | Pointer to **string** | green|yellow|red|unknown | [optional] 
**Kind** | Pointer to **string** | operator App CR spec.role (sql|kv|generic|ingress) or \&quot;\&quot; | [optional] 
**LatestTag** | Pointer to **string** | newest released tag (GH release reader — empty until wired) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** | k8s namespace | [optional] 
**Org** | Pointer to **string** | image namespace (hanzoai|luxfi|docker.io/…) | [optional] 
**Phase** | Pointer to **string** | operator status.phase (Running/Creating/…) | [optional] 
**Repo** | Pointer to **string** | owner/repo image coordinate | [optional] 
**RunningTag** | Pointer to **string** | observed from the live Deployment | [optional] 
**Tier** | Pointer to **string** | derived: cloud|data|edge|daemon|paas|app (grouping) | [optional] 
**Updated** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudProductRow

`func NewCloudProductRow() *CloudProductRow`

NewCloudProductRow instantiates a new CloudProductRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProductRowWithDefaults

`func NewCloudProductRowWithDefaults() *CloudProductRow`

NewCloudProductRowWithDefaults instantiates a new CloudProductRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *CloudProductRow) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *CloudProductRow) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *CloudProductRow) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *CloudProductRow) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDeclaredTag

`func (o *CloudProductRow) GetDeclaredTag() string`

GetDeclaredTag returns the DeclaredTag field if non-nil, zero value otherwise.

### GetDeclaredTagOk

`func (o *CloudProductRow) GetDeclaredTagOk() (*string, bool)`

GetDeclaredTagOk returns a tuple with the DeclaredTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredTag

`func (o *CloudProductRow) SetDeclaredTag(v string)`

SetDeclaredTag sets DeclaredTag field to given value.

### HasDeclaredTag

`func (o *CloudProductRow) HasDeclaredTag() bool`

HasDeclaredTag returns a boolean if a field has been set.

### GetDrift

`func (o *CloudProductRow) GetDrift() bool`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *CloudProductRow) GetDriftOk() (*bool, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *CloudProductRow) SetDrift(v bool)`

SetDrift sets Drift field to given value.

### HasDrift

`func (o *CloudProductRow) HasDrift() bool`

HasDrift returns a boolean if a field has been set.

### GetDriftSeverity

`func (o *CloudProductRow) GetDriftSeverity() string`

GetDriftSeverity returns the DriftSeverity field if non-nil, zero value otherwise.

### GetDriftSeverityOk

`func (o *CloudProductRow) GetDriftSeverityOk() (*string, bool)`

GetDriftSeverityOk returns a tuple with the DriftSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriftSeverity

`func (o *CloudProductRow) SetDriftSeverity(v string)`

SetDriftSeverity sets DriftSeverity field to given value.

### HasDriftSeverity

`func (o *CloudProductRow) HasDriftSeverity() bool`

HasDriftSeverity returns a boolean if a field has been set.

### GetEnv

`func (o *CloudProductRow) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *CloudProductRow) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *CloudProductRow) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *CloudProductRow) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetHealth

`func (o *CloudProductRow) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *CloudProductRow) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *CloudProductRow) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *CloudProductRow) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetKind

`func (o *CloudProductRow) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudProductRow) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudProductRow) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudProductRow) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLatestTag

`func (o *CloudProductRow) GetLatestTag() string`

GetLatestTag returns the LatestTag field if non-nil, zero value otherwise.

### GetLatestTagOk

`func (o *CloudProductRow) GetLatestTagOk() (*string, bool)`

GetLatestTagOk returns a tuple with the LatestTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTag

`func (o *CloudProductRow) SetLatestTag(v string)`

SetLatestTag sets LatestTag field to given value.

### HasLatestTag

`func (o *CloudProductRow) HasLatestTag() bool`

HasLatestTag returns a boolean if a field has been set.

### GetName

`func (o *CloudProductRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudProductRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudProductRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudProductRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *CloudProductRow) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CloudProductRow) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CloudProductRow) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CloudProductRow) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOrg

`func (o *CloudProductRow) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudProductRow) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudProductRow) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudProductRow) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPhase

`func (o *CloudProductRow) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *CloudProductRow) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *CloudProductRow) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *CloudProductRow) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetRepo

`func (o *CloudProductRow) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudProductRow) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudProductRow) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudProductRow) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRunningTag

`func (o *CloudProductRow) GetRunningTag() string`

GetRunningTag returns the RunningTag field if non-nil, zero value otherwise.

### GetRunningTagOk

`func (o *CloudProductRow) GetRunningTagOk() (*string, bool)`

GetRunningTagOk returns a tuple with the RunningTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningTag

`func (o *CloudProductRow) SetRunningTag(v string)`

SetRunningTag sets RunningTag field to given value.

### HasRunningTag

`func (o *CloudProductRow) HasRunningTag() bool`

HasRunningTag returns a boolean if a field has been set.

### GetTier

`func (o *CloudProductRow) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *CloudProductRow) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *CloudProductRow) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *CloudProductRow) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetUpdated

`func (o *CloudProductRow) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CloudProductRow) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CloudProductRow) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CloudProductRow) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


