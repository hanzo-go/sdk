# ResearchExperiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to **string** | Server-stamped on ingest; a payload cannot forge it. | [optional] 
**Id** | **string** | Stable id — \&quot;&lt;kind&gt;:&lt;subject&gt;:&lt;task&gt;\&quot;. Keys supersession. | 
**Revision** | Pointer to **string** |  | [optional] [default to "original"]
**Status** | Pointer to **string** |  | [optional] [default to "complete"]
**Canonical** | Pointer to **bool** | True when this is the current authoritative version. | [optional] [readonly] 
**Visibility** | Pointer to **string** | Read-only on ingest (forced private); set via a grant. | [optional] [default to "private"]
**Trainable** | Pointer to **bool** | Read-only on ingest; set via a grant. | [optional] [default to false]
**Publishable** | Pointer to **bool** | Read-only on ingest; set via a grant. | [optional] [default to false]
**Kind** | **string** |  | 
**Subject** | **string** | model / kernel / dataset under test | 
**Task** | Pointer to **string** | benchmark id | [optional] 
**Metric** | Pointer to **string** | accuracy | tok/s | loss | net-lift | [optional] 
**Value** | Pointer to **float32** | the headline number | [optional] 
**N** | Pointer to **int32** | answered sample size / coverage | [optional] 
**NTotal** | Pointer to **int32** | item target (done vs remaining &#x3D; n_total - n) | [optional] 
**CostUsd** | Pointer to **float32** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** | nonpersonal provenance (permanent retention class) | [optional] 
**GitSha** | Pointer to **string** | producing repo commit at run time (indexed) | [optional] 
**GitBranch** | Pointer to **string** | producing repo branch at run time (indexed) | [optional] 
**GitDirty** | Pointer to **bool** | tree was dirty at run time | [optional] 
**LibVersions** | Pointer to **map[string]string** | {lib: version} at run time | [optional] 
**Ts** | Pointer to **int64** | run/observation clock (unix seconds); orders versions | [optional] 
**Endpoint** | Pointer to **string** | OPTIONAL BYO arm URL; SSRF-gated at ingest (https | [optional] 

## Methods

### NewResearchExperiment

`func NewResearchExperiment(id string, kind string, subject string, ) *ResearchExperiment`

NewResearchExperiment instantiates a new ResearchExperiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResearchExperimentWithDefaults

`func NewResearchExperimentWithDefaults() *ResearchExperiment`

NewResearchExperimentWithDefaults instantiates a new ResearchExperiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ResearchExperiment) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ResearchExperiment) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ResearchExperiment) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ResearchExperiment) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetId

`func (o *ResearchExperiment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResearchExperiment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResearchExperiment) SetId(v string)`

SetId sets Id field to given value.


### GetRevision

`func (o *ResearchExperiment) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ResearchExperiment) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ResearchExperiment) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ResearchExperiment) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetStatus

`func (o *ResearchExperiment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResearchExperiment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResearchExperiment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResearchExperiment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCanonical

`func (o *ResearchExperiment) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *ResearchExperiment) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *ResearchExperiment) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *ResearchExperiment) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### GetVisibility

`func (o *ResearchExperiment) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ResearchExperiment) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ResearchExperiment) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ResearchExperiment) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTrainable

`func (o *ResearchExperiment) GetTrainable() bool`

GetTrainable returns the Trainable field if non-nil, zero value otherwise.

### GetTrainableOk

`func (o *ResearchExperiment) GetTrainableOk() (*bool, bool)`

GetTrainableOk returns a tuple with the Trainable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainable

`func (o *ResearchExperiment) SetTrainable(v bool)`

SetTrainable sets Trainable field to given value.

### HasTrainable

`func (o *ResearchExperiment) HasTrainable() bool`

HasTrainable returns a boolean if a field has been set.

### GetPublishable

`func (o *ResearchExperiment) GetPublishable() bool`

GetPublishable returns the Publishable field if non-nil, zero value otherwise.

### GetPublishableOk

`func (o *ResearchExperiment) GetPublishableOk() (*bool, bool)`

GetPublishableOk returns a tuple with the Publishable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishable

`func (o *ResearchExperiment) SetPublishable(v bool)`

SetPublishable sets Publishable field to given value.

### HasPublishable

`func (o *ResearchExperiment) HasPublishable() bool`

HasPublishable returns a boolean if a field has been set.

### GetKind

`func (o *ResearchExperiment) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ResearchExperiment) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ResearchExperiment) SetKind(v string)`

SetKind sets Kind field to given value.


### GetSubject

`func (o *ResearchExperiment) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ResearchExperiment) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ResearchExperiment) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetTask

`func (o *ResearchExperiment) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *ResearchExperiment) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *ResearchExperiment) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *ResearchExperiment) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetMetric

`func (o *ResearchExperiment) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ResearchExperiment) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ResearchExperiment) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *ResearchExperiment) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetValue

`func (o *ResearchExperiment) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResearchExperiment) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResearchExperiment) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ResearchExperiment) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetN

`func (o *ResearchExperiment) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *ResearchExperiment) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *ResearchExperiment) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *ResearchExperiment) HasN() bool`

HasN returns a boolean if a field has been set.

### GetNTotal

`func (o *ResearchExperiment) GetNTotal() int32`

GetNTotal returns the NTotal field if non-nil, zero value otherwise.

### GetNTotalOk

`func (o *ResearchExperiment) GetNTotalOk() (*int32, bool)`

GetNTotalOk returns a tuple with the NTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTotal

`func (o *ResearchExperiment) SetNTotal(v int32)`

SetNTotal sets NTotal field to given value.

### HasNTotal

`func (o *ResearchExperiment) HasNTotal() bool`

HasNTotal returns a boolean if a field has been set.

### GetCostUsd

`func (o *ResearchExperiment) GetCostUsd() float32`

GetCostUsd returns the CostUsd field if non-nil, zero value otherwise.

### GetCostUsdOk

`func (o *ResearchExperiment) GetCostUsdOk() (*float32, bool)`

GetCostUsdOk returns a tuple with the CostUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostUsd

`func (o *ResearchExperiment) SetCostUsd(v float32)`

SetCostUsd sets CostUsd field to given value.

### HasCostUsd

`func (o *ResearchExperiment) HasCostUsd() bool`

HasCostUsd returns a boolean if a field has been set.

### GetMeta

`func (o *ResearchExperiment) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ResearchExperiment) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ResearchExperiment) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ResearchExperiment) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetGitSha

`func (o *ResearchExperiment) GetGitSha() string`

GetGitSha returns the GitSha field if non-nil, zero value otherwise.

### GetGitShaOk

`func (o *ResearchExperiment) GetGitShaOk() (*string, bool)`

GetGitShaOk returns a tuple with the GitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitSha

`func (o *ResearchExperiment) SetGitSha(v string)`

SetGitSha sets GitSha field to given value.

### HasGitSha

`func (o *ResearchExperiment) HasGitSha() bool`

HasGitSha returns a boolean if a field has been set.

### GetGitBranch

`func (o *ResearchExperiment) GetGitBranch() string`

GetGitBranch returns the GitBranch field if non-nil, zero value otherwise.

### GetGitBranchOk

`func (o *ResearchExperiment) GetGitBranchOk() (*string, bool)`

GetGitBranchOk returns a tuple with the GitBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitBranch

`func (o *ResearchExperiment) SetGitBranch(v string)`

SetGitBranch sets GitBranch field to given value.

### HasGitBranch

`func (o *ResearchExperiment) HasGitBranch() bool`

HasGitBranch returns a boolean if a field has been set.

### GetGitDirty

`func (o *ResearchExperiment) GetGitDirty() bool`

GetGitDirty returns the GitDirty field if non-nil, zero value otherwise.

### GetGitDirtyOk

`func (o *ResearchExperiment) GetGitDirtyOk() (*bool, bool)`

GetGitDirtyOk returns a tuple with the GitDirty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitDirty

`func (o *ResearchExperiment) SetGitDirty(v bool)`

SetGitDirty sets GitDirty field to given value.

### HasGitDirty

`func (o *ResearchExperiment) HasGitDirty() bool`

HasGitDirty returns a boolean if a field has been set.

### GetLibVersions

`func (o *ResearchExperiment) GetLibVersions() map[string]string`

GetLibVersions returns the LibVersions field if non-nil, zero value otherwise.

### GetLibVersionsOk

`func (o *ResearchExperiment) GetLibVersionsOk() (*map[string]string, bool)`

GetLibVersionsOk returns a tuple with the LibVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibVersions

`func (o *ResearchExperiment) SetLibVersions(v map[string]string)`

SetLibVersions sets LibVersions field to given value.

### HasLibVersions

`func (o *ResearchExperiment) HasLibVersions() bool`

HasLibVersions returns a boolean if a field has been set.

### GetTs

`func (o *ResearchExperiment) GetTs() int64`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ResearchExperiment) GetTsOk() (*int64, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ResearchExperiment) SetTs(v int64)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ResearchExperiment) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetEndpoint

`func (o *ResearchExperiment) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ResearchExperiment) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ResearchExperiment) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ResearchExperiment) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


