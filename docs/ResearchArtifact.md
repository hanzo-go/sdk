# ResearchArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | base64 bytes on WRITE — the server hashes + stores them (never returned on read) | 
**Sha256** | Pointer to **string** | SERVER-derived identity (read); on write it is optional and must match the content bytes | [optional] 
**Kind** | **string** |  | 
**Ref** | Pointer to **string** | server-derived content address (sha256:&lt;hash&gt;); a client ref is ignored | [optional] 
**RunId** | Pointer to **string** | the run/snapshot id it belongs to | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** | Read-only on ingest; set via a grant. | [optional] [default to "private"]
**RetentionClass** | Pointer to **string** |  | [optional] [readonly] [default to "raw-artifact"]
**GitSha** | Pointer to **string** |  | [optional] 
**GitBranch** | Pointer to **string** |  | [optional] 
**GitDirty** | Pointer to **bool** |  | [optional] 
**LibVersions** | Pointer to **map[string]string** |  | [optional] 
**Ts** | Pointer to **int64** |  | [optional] 

## Methods

### NewResearchArtifact

`func NewResearchArtifact(content string, kind string, ) *ResearchArtifact`

NewResearchArtifact instantiates a new ResearchArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResearchArtifactWithDefaults

`func NewResearchArtifactWithDefaults() *ResearchArtifact`

NewResearchArtifactWithDefaults instantiates a new ResearchArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ResearchArtifact) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ResearchArtifact) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ResearchArtifact) SetContent(v string)`

SetContent sets Content field to given value.


### GetSha256

`func (o *ResearchArtifact) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ResearchArtifact) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ResearchArtifact) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *ResearchArtifact) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetKind

`func (o *ResearchArtifact) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ResearchArtifact) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ResearchArtifact) SetKind(v string)`

SetKind sets Kind field to given value.


### GetRef

`func (o *ResearchArtifact) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ResearchArtifact) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ResearchArtifact) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ResearchArtifact) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRunId

`func (o *ResearchArtifact) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *ResearchArtifact) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *ResearchArtifact) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *ResearchArtifact) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetProject

`func (o *ResearchArtifact) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ResearchArtifact) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ResearchArtifact) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ResearchArtifact) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetVisibility

`func (o *ResearchArtifact) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ResearchArtifact) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ResearchArtifact) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ResearchArtifact) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetRetentionClass

`func (o *ResearchArtifact) GetRetentionClass() string`

GetRetentionClass returns the RetentionClass field if non-nil, zero value otherwise.

### GetRetentionClassOk

`func (o *ResearchArtifact) GetRetentionClassOk() (*string, bool)`

GetRetentionClassOk returns a tuple with the RetentionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionClass

`func (o *ResearchArtifact) SetRetentionClass(v string)`

SetRetentionClass sets RetentionClass field to given value.

### HasRetentionClass

`func (o *ResearchArtifact) HasRetentionClass() bool`

HasRetentionClass returns a boolean if a field has been set.

### GetGitSha

`func (o *ResearchArtifact) GetGitSha() string`

GetGitSha returns the GitSha field if non-nil, zero value otherwise.

### GetGitShaOk

`func (o *ResearchArtifact) GetGitShaOk() (*string, bool)`

GetGitShaOk returns a tuple with the GitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitSha

`func (o *ResearchArtifact) SetGitSha(v string)`

SetGitSha sets GitSha field to given value.

### HasGitSha

`func (o *ResearchArtifact) HasGitSha() bool`

HasGitSha returns a boolean if a field has been set.

### GetGitBranch

`func (o *ResearchArtifact) GetGitBranch() string`

GetGitBranch returns the GitBranch field if non-nil, zero value otherwise.

### GetGitBranchOk

`func (o *ResearchArtifact) GetGitBranchOk() (*string, bool)`

GetGitBranchOk returns a tuple with the GitBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitBranch

`func (o *ResearchArtifact) SetGitBranch(v string)`

SetGitBranch sets GitBranch field to given value.

### HasGitBranch

`func (o *ResearchArtifact) HasGitBranch() bool`

HasGitBranch returns a boolean if a field has been set.

### GetGitDirty

`func (o *ResearchArtifact) GetGitDirty() bool`

GetGitDirty returns the GitDirty field if non-nil, zero value otherwise.

### GetGitDirtyOk

`func (o *ResearchArtifact) GetGitDirtyOk() (*bool, bool)`

GetGitDirtyOk returns a tuple with the GitDirty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitDirty

`func (o *ResearchArtifact) SetGitDirty(v bool)`

SetGitDirty sets GitDirty field to given value.

### HasGitDirty

`func (o *ResearchArtifact) HasGitDirty() bool`

HasGitDirty returns a boolean if a field has been set.

### GetLibVersions

`func (o *ResearchArtifact) GetLibVersions() map[string]string`

GetLibVersions returns the LibVersions field if non-nil, zero value otherwise.

### GetLibVersionsOk

`func (o *ResearchArtifact) GetLibVersionsOk() (*map[string]string, bool)`

GetLibVersionsOk returns a tuple with the LibVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibVersions

`func (o *ResearchArtifact) SetLibVersions(v map[string]string)`

SetLibVersions sets LibVersions field to given value.

### HasLibVersions

`func (o *ResearchArtifact) HasLibVersions() bool`

HasLibVersions returns a boolean if a field has been set.

### GetTs

`func (o *ResearchArtifact) GetTs() int64`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ResearchArtifact) GetTsOk() (*int64, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ResearchArtifact) SetTs(v int64)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ResearchArtifact) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


