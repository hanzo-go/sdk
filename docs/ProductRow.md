# ProductRow

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

### NewProductRow

`func NewProductRow() *ProductRow`

NewProductRow instantiates a new ProductRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductRowWithDefaults

`func NewProductRowWithDefaults() *ProductRow`

NewProductRowWithDefaults instantiates a new ProductRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *ProductRow) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ProductRow) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ProductRow) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ProductRow) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDeclaredTag

`func (o *ProductRow) GetDeclaredTag() string`

GetDeclaredTag returns the DeclaredTag field if non-nil, zero value otherwise.

### GetDeclaredTagOk

`func (o *ProductRow) GetDeclaredTagOk() (*string, bool)`

GetDeclaredTagOk returns a tuple with the DeclaredTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredTag

`func (o *ProductRow) SetDeclaredTag(v string)`

SetDeclaredTag sets DeclaredTag field to given value.

### HasDeclaredTag

`func (o *ProductRow) HasDeclaredTag() bool`

HasDeclaredTag returns a boolean if a field has been set.

### GetDrift

`func (o *ProductRow) GetDrift() bool`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *ProductRow) GetDriftOk() (*bool, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *ProductRow) SetDrift(v bool)`

SetDrift sets Drift field to given value.

### HasDrift

`func (o *ProductRow) HasDrift() bool`

HasDrift returns a boolean if a field has been set.

### GetDriftSeverity

`func (o *ProductRow) GetDriftSeverity() string`

GetDriftSeverity returns the DriftSeverity field if non-nil, zero value otherwise.

### GetDriftSeverityOk

`func (o *ProductRow) GetDriftSeverityOk() (*string, bool)`

GetDriftSeverityOk returns a tuple with the DriftSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriftSeverity

`func (o *ProductRow) SetDriftSeverity(v string)`

SetDriftSeverity sets DriftSeverity field to given value.

### HasDriftSeverity

`func (o *ProductRow) HasDriftSeverity() bool`

HasDriftSeverity returns a boolean if a field has been set.

### GetEnv

`func (o *ProductRow) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ProductRow) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ProductRow) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ProductRow) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetHealth

`func (o *ProductRow) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProductRow) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProductRow) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ProductRow) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetKind

`func (o *ProductRow) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ProductRow) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ProductRow) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ProductRow) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLatestTag

`func (o *ProductRow) GetLatestTag() string`

GetLatestTag returns the LatestTag field if non-nil, zero value otherwise.

### GetLatestTagOk

`func (o *ProductRow) GetLatestTagOk() (*string, bool)`

GetLatestTagOk returns a tuple with the LatestTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTag

`func (o *ProductRow) SetLatestTag(v string)`

SetLatestTag sets LatestTag field to given value.

### HasLatestTag

`func (o *ProductRow) HasLatestTag() bool`

HasLatestTag returns a boolean if a field has been set.

### GetName

`func (o *ProductRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *ProductRow) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ProductRow) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ProductRow) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ProductRow) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOrg

`func (o *ProductRow) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *ProductRow) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *ProductRow) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *ProductRow) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPhase

`func (o *ProductRow) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ProductRow) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ProductRow) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *ProductRow) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetRepo

`func (o *ProductRow) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ProductRow) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ProductRow) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ProductRow) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRunningTag

`func (o *ProductRow) GetRunningTag() string`

GetRunningTag returns the RunningTag field if non-nil, zero value otherwise.

### GetRunningTagOk

`func (o *ProductRow) GetRunningTagOk() (*string, bool)`

GetRunningTagOk returns a tuple with the RunningTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningTag

`func (o *ProductRow) SetRunningTag(v string)`

SetRunningTag sets RunningTag field to given value.

### HasRunningTag

`func (o *ProductRow) HasRunningTag() bool`

HasRunningTag returns a boolean if a field has been set.

### GetTier

`func (o *ProductRow) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *ProductRow) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *ProductRow) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *ProductRow) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetUpdated

`func (o *ProductRow) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ProductRow) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ProductRow) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ProductRow) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


