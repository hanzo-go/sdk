# CloudExperiment

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

### NewCloudExperiment

`func NewCloudExperiment() *CloudExperiment`

NewCloudExperiment instantiates a new CloudExperiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudExperimentWithDefaults

`func NewCloudExperimentWithDefaults() *CloudExperiment`

NewCloudExperimentWithDefaults instantiates a new CloudExperiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonical

`func (o *CloudExperiment) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *CloudExperiment) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *CloudExperiment) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *CloudExperiment) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### GetCostUsd

`func (o *CloudExperiment) GetCostUsd() float32`

GetCostUsd returns the CostUsd field if non-nil, zero value otherwise.

### GetCostUsdOk

`func (o *CloudExperiment) GetCostUsdOk() (*float32, bool)`

GetCostUsdOk returns a tuple with the CostUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostUsd

`func (o *CloudExperiment) SetCostUsd(v float32)`

SetCostUsd sets CostUsd field to given value.

### HasCostUsd

`func (o *CloudExperiment) HasCostUsd() bool`

HasCostUsd returns a boolean if a field has been set.

### GetEndpoint

`func (o *CloudExperiment) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CloudExperiment) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CloudExperiment) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *CloudExperiment) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetGitBranch

`func (o *CloudExperiment) GetGitBranch() string`

GetGitBranch returns the GitBranch field if non-nil, zero value otherwise.

### GetGitBranchOk

`func (o *CloudExperiment) GetGitBranchOk() (*string, bool)`

GetGitBranchOk returns a tuple with the GitBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitBranch

`func (o *CloudExperiment) SetGitBranch(v string)`

SetGitBranch sets GitBranch field to given value.

### HasGitBranch

`func (o *CloudExperiment) HasGitBranch() bool`

HasGitBranch returns a boolean if a field has been set.

### GetGitDirty

`func (o *CloudExperiment) GetGitDirty() bool`

GetGitDirty returns the GitDirty field if non-nil, zero value otherwise.

### GetGitDirtyOk

`func (o *CloudExperiment) GetGitDirtyOk() (*bool, bool)`

GetGitDirtyOk returns a tuple with the GitDirty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitDirty

`func (o *CloudExperiment) SetGitDirty(v bool)`

SetGitDirty sets GitDirty field to given value.

### HasGitDirty

`func (o *CloudExperiment) HasGitDirty() bool`

HasGitDirty returns a boolean if a field has been set.

### GetGitSha

`func (o *CloudExperiment) GetGitSha() string`

GetGitSha returns the GitSha field if non-nil, zero value otherwise.

### GetGitShaOk

`func (o *CloudExperiment) GetGitShaOk() (*string, bool)`

GetGitShaOk returns a tuple with the GitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitSha

`func (o *CloudExperiment) SetGitSha(v string)`

SetGitSha sets GitSha field to given value.

### HasGitSha

`func (o *CloudExperiment) HasGitSha() bool`

HasGitSha returns a boolean if a field has been set.

### GetId

`func (o *CloudExperiment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudExperiment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudExperiment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudExperiment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *CloudExperiment) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudExperiment) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudExperiment) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudExperiment) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLibVersions

`func (o *CloudExperiment) GetLibVersions() interface{}`

GetLibVersions returns the LibVersions field if non-nil, zero value otherwise.

### GetLibVersionsOk

`func (o *CloudExperiment) GetLibVersionsOk() (*interface{}, bool)`

GetLibVersionsOk returns a tuple with the LibVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibVersions

`func (o *CloudExperiment) SetLibVersions(v interface{})`

SetLibVersions sets LibVersions field to given value.

### HasLibVersions

`func (o *CloudExperiment) HasLibVersions() bool`

HasLibVersions returns a boolean if a field has been set.

### SetLibVersionsNil

`func (o *CloudExperiment) SetLibVersionsNil(b bool)`

 SetLibVersionsNil sets the value for LibVersions to be an explicit nil

### UnsetLibVersions
`func (o *CloudExperiment) UnsetLibVersions()`

UnsetLibVersions ensures that no value is present for LibVersions, not even an explicit nil
### GetMeta

`func (o *CloudExperiment) GetMeta() interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CloudExperiment) GetMetaOk() (*interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CloudExperiment) SetMeta(v interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CloudExperiment) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *CloudExperiment) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *CloudExperiment) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetMetric

`func (o *CloudExperiment) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *CloudExperiment) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *CloudExperiment) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *CloudExperiment) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetN

`func (o *CloudExperiment) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *CloudExperiment) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *CloudExperiment) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *CloudExperiment) HasN() bool`

HasN returns a boolean if a field has been set.

### GetNTotal

`func (o *CloudExperiment) GetNTotal() int32`

GetNTotal returns the NTotal field if non-nil, zero value otherwise.

### GetNTotalOk

`func (o *CloudExperiment) GetNTotalOk() (*int32, bool)`

GetNTotalOk returns a tuple with the NTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNTotal

`func (o *CloudExperiment) SetNTotal(v int32)`

SetNTotal sets NTotal field to given value.

### HasNTotal

`func (o *CloudExperiment) HasNTotal() bool`

HasNTotal returns a boolean if a field has been set.

### GetProject

`func (o *CloudExperiment) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudExperiment) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudExperiment) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudExperiment) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPublishable

`func (o *CloudExperiment) GetPublishable() bool`

GetPublishable returns the Publishable field if non-nil, zero value otherwise.

### GetPublishableOk

`func (o *CloudExperiment) GetPublishableOk() (*bool, bool)`

GetPublishableOk returns a tuple with the Publishable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishable

`func (o *CloudExperiment) SetPublishable(v bool)`

SetPublishable sets Publishable field to given value.

### HasPublishable

`func (o *CloudExperiment) HasPublishable() bool`

HasPublishable returns a boolean if a field has been set.

### GetRevision

`func (o *CloudExperiment) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *CloudExperiment) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *CloudExperiment) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *CloudExperiment) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetStatus

`func (o *CloudExperiment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudExperiment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudExperiment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudExperiment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *CloudExperiment) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CloudExperiment) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CloudExperiment) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CloudExperiment) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTask

`func (o *CloudExperiment) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *CloudExperiment) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *CloudExperiment) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *CloudExperiment) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTrainable

`func (o *CloudExperiment) GetTrainable() bool`

GetTrainable returns the Trainable field if non-nil, zero value otherwise.

### GetTrainableOk

`func (o *CloudExperiment) GetTrainableOk() (*bool, bool)`

GetTrainableOk returns a tuple with the Trainable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainable

`func (o *CloudExperiment) SetTrainable(v bool)`

SetTrainable sets Trainable field to given value.

### HasTrainable

`func (o *CloudExperiment) HasTrainable() bool`

HasTrainable returns a boolean if a field has been set.

### GetTs

`func (o *CloudExperiment) GetTs() int32`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *CloudExperiment) GetTsOk() (*int32, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *CloudExperiment) SetTs(v int32)`

SetTs sets Ts field to given value.

### HasTs

`func (o *CloudExperiment) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetValue

`func (o *CloudExperiment) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CloudExperiment) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CloudExperiment) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *CloudExperiment) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVisibility

`func (o *CloudExperiment) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CloudExperiment) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CloudExperiment) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CloudExperiment) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


