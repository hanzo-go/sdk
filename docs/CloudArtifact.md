# CloudArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | base64 bytes on write; the server hashes + stores them (never returned) | [optional] 
**GitBranch** | Pointer to **string** |  | [optional] 
**GitDirty** | Pointer to **bool** |  | [optional] 
**GitSha** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**LibVersions** | Pointer to **interface{}** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to **string** | server-derived content address (sha256:&lt;hash&gt;) | [optional] 
**RetentionClass** | Pointer to **string** |  | [optional] 
**RunId** | Pointer to **string** |  | [optional] 
**Sha256** | Pointer to **string** | SERVER-derived on write; the identity | [optional] 
**Ts** | Pointer to **int32** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudArtifact

`func NewCloudArtifact() *CloudArtifact`

NewCloudArtifact instantiates a new CloudArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudArtifactWithDefaults

`func NewCloudArtifactWithDefaults() *CloudArtifact`

NewCloudArtifactWithDefaults instantiates a new CloudArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CloudArtifact) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CloudArtifact) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CloudArtifact) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CloudArtifact) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetGitBranch

`func (o *CloudArtifact) GetGitBranch() string`

GetGitBranch returns the GitBranch field if non-nil, zero value otherwise.

### GetGitBranchOk

`func (o *CloudArtifact) GetGitBranchOk() (*string, bool)`

GetGitBranchOk returns a tuple with the GitBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitBranch

`func (o *CloudArtifact) SetGitBranch(v string)`

SetGitBranch sets GitBranch field to given value.

### HasGitBranch

`func (o *CloudArtifact) HasGitBranch() bool`

HasGitBranch returns a boolean if a field has been set.

### GetGitDirty

`func (o *CloudArtifact) GetGitDirty() bool`

GetGitDirty returns the GitDirty field if non-nil, zero value otherwise.

### GetGitDirtyOk

`func (o *CloudArtifact) GetGitDirtyOk() (*bool, bool)`

GetGitDirtyOk returns a tuple with the GitDirty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitDirty

`func (o *CloudArtifact) SetGitDirty(v bool)`

SetGitDirty sets GitDirty field to given value.

### HasGitDirty

`func (o *CloudArtifact) HasGitDirty() bool`

HasGitDirty returns a boolean if a field has been set.

### GetGitSha

`func (o *CloudArtifact) GetGitSha() string`

GetGitSha returns the GitSha field if non-nil, zero value otherwise.

### GetGitShaOk

`func (o *CloudArtifact) GetGitShaOk() (*string, bool)`

GetGitShaOk returns a tuple with the GitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitSha

`func (o *CloudArtifact) SetGitSha(v string)`

SetGitSha sets GitSha field to given value.

### HasGitSha

`func (o *CloudArtifact) HasGitSha() bool`

HasGitSha returns a boolean if a field has been set.

### GetKind

`func (o *CloudArtifact) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudArtifact) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudArtifact) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudArtifact) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLibVersions

`func (o *CloudArtifact) GetLibVersions() interface{}`

GetLibVersions returns the LibVersions field if non-nil, zero value otherwise.

### GetLibVersionsOk

`func (o *CloudArtifact) GetLibVersionsOk() (*interface{}, bool)`

GetLibVersionsOk returns a tuple with the LibVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibVersions

`func (o *CloudArtifact) SetLibVersions(v interface{})`

SetLibVersions sets LibVersions field to given value.

### HasLibVersions

`func (o *CloudArtifact) HasLibVersions() bool`

HasLibVersions returns a boolean if a field has been set.

### SetLibVersionsNil

`func (o *CloudArtifact) SetLibVersionsNil(b bool)`

 SetLibVersionsNil sets the value for LibVersions to be an explicit nil

### UnsetLibVersions
`func (o *CloudArtifact) UnsetLibVersions()`

UnsetLibVersions ensures that no value is present for LibVersions, not even an explicit nil
### GetProject

`func (o *CloudArtifact) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudArtifact) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudArtifact) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudArtifact) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRef

`func (o *CloudArtifact) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CloudArtifact) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CloudArtifact) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CloudArtifact) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRetentionClass

`func (o *CloudArtifact) GetRetentionClass() string`

GetRetentionClass returns the RetentionClass field if non-nil, zero value otherwise.

### GetRetentionClassOk

`func (o *CloudArtifact) GetRetentionClassOk() (*string, bool)`

GetRetentionClassOk returns a tuple with the RetentionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionClass

`func (o *CloudArtifact) SetRetentionClass(v string)`

SetRetentionClass sets RetentionClass field to given value.

### HasRetentionClass

`func (o *CloudArtifact) HasRetentionClass() bool`

HasRetentionClass returns a boolean if a field has been set.

### GetRunId

`func (o *CloudArtifact) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *CloudArtifact) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *CloudArtifact) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *CloudArtifact) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetSha256

`func (o *CloudArtifact) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *CloudArtifact) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *CloudArtifact) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *CloudArtifact) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetTs

`func (o *CloudArtifact) GetTs() int32`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *CloudArtifact) GetTsOk() (*int32, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *CloudArtifact) SetTs(v int32)`

SetTs sets Ts field to given value.

### HasTs

`func (o *CloudArtifact) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetVisibility

`func (o *CloudArtifact) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CloudArtifact) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CloudArtifact) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CloudArtifact) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


