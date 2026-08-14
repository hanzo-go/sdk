# Experiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Canonical** | Pointer to **bool** |  | [optional] 
**CostUsd** | Pointer to **float32** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**GitBranch** | Pointer to **string** |  | [optional] 
**GitDirty** | Pointer to **bool** |  | [optional] 
**GitSha** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**LibVersions** | Pointer to **interface{}** |  | [optional] 
**Meta** | Pointer to **interface{}** |  | [optional] 
**Metric** | Pointer to **string** |  | [optional] 
**N** | Pointer to **int32** |  | [optional] 
**NTotal** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Publishable** | Pointer to **bool** |  | [optional] 
**Revision** | Pointer to **string** | original | corrected | retracted | [optional] 
**Status** | Pointer to **string** | planning | running | complete | faulted | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Task** | Pointer to **string** |  | [optional] 
**Trainable** | Pointer to **bool** |  | [optional] 
**Ts** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewExperiment

`func NewExperiment() *Experiment`

NewExperiment instantiates a new Experiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentWithDefaults

`func NewExperimentWithDefaults() *Experiment`

NewExperimentWithDefaults instantiates a new Experiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonical

`func (o *Experiment) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *Experiment) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *Experiment) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *Experiment) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### GetCostUsd

`func (o *Experiment) GetCostUsd() float32`

GetCostUsd returns the CostUsd field if non-nil, zero value otherwise.

### GetCostUsdOk

`func (o *Experiment) GetCostUsdOk() (*float32, bool)`

GetCostUsdOk returns a tuple with the CostUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostUsd

`func (o *Experiment) SetCostUsd(v float32)`

SetCostUsd sets CostUsd field to given value.

### HasCostUsd

`func (o *Experiment) HasCostUsd() bool`

HasCostUsd returns a boolean if a field has been set.

### GetEndpoint

`func (o *Experiment) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *Experiment) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *Experiment) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *Experiment) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetGitBranch

`func (o *Experiment) GetGitBranch() string`

GetGitBranch returns the GitBranch field if non-nil, zero value otherwise.

### GetGitBranchOk

`func (o *Experiment) GetGitBranchOk() (*string, bool)`

GetGitBranchOk returns a tuple with the GitBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitBranch

`func (o *Experiment) SetGitBranch(v string)`

SetGitBranch sets GitBranch field to given value.

### HasGitBranch

`func (o *Experiment) HasGitBranch() bool`

HasGitBranch returns a boolean if a field has been set.

### GetGitDirty

`func (o *Experiment) GetGitDirty() bool`

GetGitDirty returns the GitDirty field if non-nil, zero value otherwise.

### GetGitDirtyOk

`func (o *Experiment) GetGitDirtyOk() (*bool, bool)`

GetGitDirtyOk returns a tuple with the GitDirty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitDirty

`func (o *Experiment) SetGitDirty(v bool)`

SetGitDirty sets GitDirty field to given value.

### HasGitDirty

`func (o *Experiment) HasGitDirty() bool`

HasGitDirty returns a boolean if a field has been set.

### GetGitSha

`func (o *Experiment) GetGitSha() string`

GetGitSha returns the GitSha field if non-nil, zero value otherwise.

### GetGitShaOk

`func (o *Experiment) GetGitShaOk() (*string, bool)`

GetGitShaOk returns a tuple with the GitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitSha

`func (o *Experiment) SetGitSha(v string)`

SetGitSha sets GitSha field to given value.

### HasGitSha

`func (o *Experiment) HasGitSha() bool`

HasGitSha returns a boolean if a field has been set.

### GetId

`func (o *Experiment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Experiment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Experiment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Experiment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *Experiment) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Experiment) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Experiment) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Experiment) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLibVersions

`func (o *Experiment) GetLibVersions() interface{}`

GetLibVersions returns the LibVersions field if non-nil, zero value otherwise.

### GetLibVersionsOk

`func (o *Experiment) GetLibVersionsOk() (*interface{}, bool)`

GetLibVersionsOk returns a tuple with the LibVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibVersions

`func (o *Experiment) SetLibVersions(v interface{})`

SetLibVersions sets LibVersions field to given value.

### HasLibVersions

`func (o *Experiment) HasLibVersions() bool`

HasLibVersions returns a boolean if a field has been set.

### SetLibVersionsNil

`func (o *Experiment) SetLibVersionsNil(b bool)`

 SetLibVersionsNil sets the value for LibVersions to be an explicit nil

### UnsetLibVersions
`func (o *Experiment) UnsetLibVersions()`

UnsetLibVersions ensures that no value is present for LibVersions, not even an explicit nil
### GetMeta

`func (o *Experiment) GetMeta() interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Experiment) GetMetaOk() (*interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Experiment) SetMeta(v interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Experiment) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *Experiment) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *Experiment) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetMetric

`func (o *Experiment) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *Experiment) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *Experiment) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *Experiment) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetN

`func (o *Experiment) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *Experiment) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *Experiment) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *Experiment) HasN() bool`

HasN returns a boolean if a field has been set.

### GetNTotal

`func (o *Experiment) GetNTotal() int32`

GetNTotal returns the NTotal field if non-nil, zero value otherwise.

### GetNTotalOk

`func (o *Experiment) GetNTotalOk() (*int32, bool)`

GetNTotalOk returns a tuple with the NTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTotal

`func (o *Experiment) SetNTotal(v int32)`

SetNTotal sets NTotal field to given value.

### HasNTotal

`func (o *Experiment) HasNTotal() bool`

HasNTotal returns a boolean if a field has been set.

### GetProject

`func (o *Experiment) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Experiment) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Experiment) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Experiment) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPublishable

`func (o *Experiment) GetPublishable() bool`

GetPublishable returns the Publishable field if non-nil, zero value otherwise.

### GetPublishableOk

`func (o *Experiment) GetPublishableOk() (*bool, bool)`

GetPublishableOk returns a tuple with the Publishable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishable

`func (o *Experiment) SetPublishable(v bool)`

SetPublishable sets Publishable field to given value.

### HasPublishable

`func (o *Experiment) HasPublishable() bool`

HasPublishable returns a boolean if a field has been set.

### GetRevision

`func (o *Experiment) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Experiment) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Experiment) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *Experiment) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetStatus

`func (o *Experiment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Experiment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Experiment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Experiment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *Experiment) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Experiment) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Experiment) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Experiment) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTask

`func (o *Experiment) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *Experiment) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *Experiment) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *Experiment) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTrainable

`func (o *Experiment) GetTrainable() bool`

GetTrainable returns the Trainable field if non-nil, zero value otherwise.

### GetTrainableOk

`func (o *Experiment) GetTrainableOk() (*bool, bool)`

GetTrainableOk returns a tuple with the Trainable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainable

`func (o *Experiment) SetTrainable(v bool)`

SetTrainable sets Trainable field to given value.

### HasTrainable

`func (o *Experiment) HasTrainable() bool`

HasTrainable returns a boolean if a field has been set.

### GetTs

`func (o *Experiment) GetTs() int32`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *Experiment) GetTsOk() (*int32, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *Experiment) SetTs(v int32)`

SetTs sets Ts field to given value.

### HasTs

`func (o *Experiment) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetValue

`func (o *Experiment) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Experiment) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Experiment) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *Experiment) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVisibility

`func (o *Experiment) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Experiment) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Experiment) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Experiment) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


