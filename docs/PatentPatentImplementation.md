# PatentPatentImplementation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactSha256** | Pointer to **string** | ArtifactSHA256 is the digest of the recorded artifact — a diff, a file, a build — when the caller has one and wants it addressed by content. | [optional] 
**AuthoredBy** | Pointer to **string** | AuthoredBy is the validated principal who recorded this. Authorship of the IDEA is PatentIdea.AuthoredBy — this field records who filed the implementation record, which is not the same question. | [optional] 
**Canonical** | Pointer to **bool** | Canonical is DERIVED on read. | [optional] 
**ContentHash** | Pointer to **string** | ContentHash is the server&#39;s digest of the recorded content. | [optional] 
**Generated** | Pointer to **bool** | Generated says the recorded artifact was produced with a model&#39;s assistance rather than typed by hand. | [optional] 
**GitBranch** | Pointer to **string** | GitBranch is the branch that commit was on. | [optional] 
**GitDirty** | Pointer to **bool** | GitDirty says the working tree carried uncommitted changes when this was recorded. It is recorded because an exact commit claim that is not true is worse than no commit claim. | [optional] 
**GitSha** | Pointer to **string** | GitSHA is the commit. Together with Seq it is what makes the claim checkable outside this app: a reader can go read the code. | [optional] 
**HumanReview** | Pointer to **string** | HumanReview is what a person did with the output: accepted, edited, rewritten, or none. &#x60;none&#x60; on a novel mechanism is the disposition to know about. | [optional] 
**Id** | Pointer to **string** | ID is this implementation&#39;s stable id — usually the artifact or change it records, e.g. router-v2-generated. Required. | [optional] 
**IdeaIds** | Pointer to **[]string** | IdeaIDs names the conceptions this implements. REQUIRED and non-empty: an implementation citing nothing is the inverted order the guidance is written against, and filing it silently would be the worst thing this surface could do. | [optional] 
**Model** | Pointer to **string** | Model names the system that produced the code, as it was addressed. This is the field that answers \&quot;how much of this was a tool\&quot; without a guess. | [optional] 
**Notes** | Pointer to **string** | Notes is anything a later reader needs. | [optional] 
**Project** | Pointer to **string** | Project is the SERVER&#39;s scope value. | [optional] 
**PromptSha256** | Pointer to **string** | PromptSHA256 is the digest of the prompt, not the prompt. The text often carries third-party material; the digest lets a later reader prove the same request was made. | [optional] 
**RecordedAt** | Pointer to **string** | RecordedAt is the server&#39;s clock. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository the work landed in. | [optional] 
**Revision** | Pointer to **string** | Revision is original, corrected or retracted. | [optional] 
**Seq** | Pointer to **int64** | Seq is the append order, which is what GET /v1/patent/proof compares against the cited ideas. | [optional] 
**Visibility** | Pointer to **string** | Visibility is private, org or public; forced private on a write. | [optional] 

## Methods

### NewPatentPatentImplementation

`func NewPatentPatentImplementation() *PatentPatentImplementation`

NewPatentPatentImplementation instantiates a new PatentPatentImplementation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentPatentImplementationWithDefaults

`func NewPatentPatentImplementationWithDefaults() *PatentPatentImplementation`

NewPatentPatentImplementationWithDefaults instantiates a new PatentPatentImplementation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactSha256

`func (o *PatentPatentImplementation) GetArtifactSha256() string`

GetArtifactSha256 returns the ArtifactSha256 field if non-nil, zero value otherwise.

### GetArtifactSha256Ok

`func (o *PatentPatentImplementation) GetArtifactSha256Ok() (*string, bool)`

GetArtifactSha256Ok returns a tuple with the ArtifactSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactSha256

`func (o *PatentPatentImplementation) SetArtifactSha256(v string)`

SetArtifactSha256 sets ArtifactSha256 field to given value.

### HasArtifactSha256

`func (o *PatentPatentImplementation) HasArtifactSha256() bool`

HasArtifactSha256 returns a boolean if a field has been set.

### GetAuthoredBy

`func (o *PatentPatentImplementation) GetAuthoredBy() string`

GetAuthoredBy returns the AuthoredBy field if non-nil, zero value otherwise.

### GetAuthoredByOk

`func (o *PatentPatentImplementation) GetAuthoredByOk() (*string, bool)`

GetAuthoredByOk returns a tuple with the AuthoredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoredBy

`func (o *PatentPatentImplementation) SetAuthoredBy(v string)`

SetAuthoredBy sets AuthoredBy field to given value.

### HasAuthoredBy

`func (o *PatentPatentImplementation) HasAuthoredBy() bool`

HasAuthoredBy returns a boolean if a field has been set.

### GetCanonical

`func (o *PatentPatentImplementation) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *PatentPatentImplementation) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *PatentPatentImplementation) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *PatentPatentImplementation) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### GetContentHash

`func (o *PatentPatentImplementation) GetContentHash() string`

GetContentHash returns the ContentHash field if non-nil, zero value otherwise.

### GetContentHashOk

`func (o *PatentPatentImplementation) GetContentHashOk() (*string, bool)`

GetContentHashOk returns a tuple with the ContentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHash

`func (o *PatentPatentImplementation) SetContentHash(v string)`

SetContentHash sets ContentHash field to given value.

### HasContentHash

`func (o *PatentPatentImplementation) HasContentHash() bool`

HasContentHash returns a boolean if a field has been set.

### GetGenerated

`func (o *PatentPatentImplementation) GetGenerated() bool`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *PatentPatentImplementation) GetGeneratedOk() (*bool, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *PatentPatentImplementation) SetGenerated(v bool)`

SetGenerated sets Generated field to given value.

### HasGenerated

`func (o *PatentPatentImplementation) HasGenerated() bool`

HasGenerated returns a boolean if a field has been set.

### GetGitBranch

`func (o *PatentPatentImplementation) GetGitBranch() string`

GetGitBranch returns the GitBranch field if non-nil, zero value otherwise.

### GetGitBranchOk

`func (o *PatentPatentImplementation) GetGitBranchOk() (*string, bool)`

GetGitBranchOk returns a tuple with the GitBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitBranch

`func (o *PatentPatentImplementation) SetGitBranch(v string)`

SetGitBranch sets GitBranch field to given value.

### HasGitBranch

`func (o *PatentPatentImplementation) HasGitBranch() bool`

HasGitBranch returns a boolean if a field has been set.

### GetGitDirty

`func (o *PatentPatentImplementation) GetGitDirty() bool`

GetGitDirty returns the GitDirty field if non-nil, zero value otherwise.

### GetGitDirtyOk

`func (o *PatentPatentImplementation) GetGitDirtyOk() (*bool, bool)`

GetGitDirtyOk returns a tuple with the GitDirty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitDirty

`func (o *PatentPatentImplementation) SetGitDirty(v bool)`

SetGitDirty sets GitDirty field to given value.

### HasGitDirty

`func (o *PatentPatentImplementation) HasGitDirty() bool`

HasGitDirty returns a boolean if a field has been set.

### GetGitSha

`func (o *PatentPatentImplementation) GetGitSha() string`

GetGitSha returns the GitSha field if non-nil, zero value otherwise.

### GetGitShaOk

`func (o *PatentPatentImplementation) GetGitShaOk() (*string, bool)`

GetGitShaOk returns a tuple with the GitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitSha

`func (o *PatentPatentImplementation) SetGitSha(v string)`

SetGitSha sets GitSha field to given value.

### HasGitSha

`func (o *PatentPatentImplementation) HasGitSha() bool`

HasGitSha returns a boolean if a field has been set.

### GetHumanReview

`func (o *PatentPatentImplementation) GetHumanReview() string`

GetHumanReview returns the HumanReview field if non-nil, zero value otherwise.

### GetHumanReviewOk

`func (o *PatentPatentImplementation) GetHumanReviewOk() (*string, bool)`

GetHumanReviewOk returns a tuple with the HumanReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanReview

`func (o *PatentPatentImplementation) SetHumanReview(v string)`

SetHumanReview sets HumanReview field to given value.

### HasHumanReview

`func (o *PatentPatentImplementation) HasHumanReview() bool`

HasHumanReview returns a boolean if a field has been set.

### GetId

`func (o *PatentPatentImplementation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatentPatentImplementation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatentPatentImplementation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatentPatentImplementation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdeaIds

`func (o *PatentPatentImplementation) GetIdeaIds() []string`

GetIdeaIds returns the IdeaIds field if non-nil, zero value otherwise.

### GetIdeaIdsOk

`func (o *PatentPatentImplementation) GetIdeaIdsOk() (*[]string, bool)`

GetIdeaIdsOk returns a tuple with the IdeaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdeaIds

`func (o *PatentPatentImplementation) SetIdeaIds(v []string)`

SetIdeaIds sets IdeaIds field to given value.

### HasIdeaIds

`func (o *PatentPatentImplementation) HasIdeaIds() bool`

HasIdeaIds returns a boolean if a field has been set.

### GetModel

`func (o *PatentPatentImplementation) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PatentPatentImplementation) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PatentPatentImplementation) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *PatentPatentImplementation) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNotes

`func (o *PatentPatentImplementation) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PatentPatentImplementation) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PatentPatentImplementation) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PatentPatentImplementation) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetProject

`func (o *PatentPatentImplementation) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatentPatentImplementation) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatentPatentImplementation) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatentPatentImplementation) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPromptSha256

`func (o *PatentPatentImplementation) GetPromptSha256() string`

GetPromptSha256 returns the PromptSha256 field if non-nil, zero value otherwise.

### GetPromptSha256Ok

`func (o *PatentPatentImplementation) GetPromptSha256Ok() (*string, bool)`

GetPromptSha256Ok returns a tuple with the PromptSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptSha256

`func (o *PatentPatentImplementation) SetPromptSha256(v string)`

SetPromptSha256 sets PromptSha256 field to given value.

### HasPromptSha256

`func (o *PatentPatentImplementation) HasPromptSha256() bool`

HasPromptSha256 returns a boolean if a field has been set.

### GetRecordedAt

`func (o *PatentPatentImplementation) GetRecordedAt() string`

GetRecordedAt returns the RecordedAt field if non-nil, zero value otherwise.

### GetRecordedAtOk

`func (o *PatentPatentImplementation) GetRecordedAtOk() (*string, bool)`

GetRecordedAtOk returns a tuple with the RecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAt

`func (o *PatentPatentImplementation) SetRecordedAt(v string)`

SetRecordedAt sets RecordedAt field to given value.

### HasRecordedAt

`func (o *PatentPatentImplementation) HasRecordedAt() bool`

HasRecordedAt returns a boolean if a field has been set.

### GetRepo

`func (o *PatentPatentImplementation) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *PatentPatentImplementation) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *PatentPatentImplementation) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *PatentPatentImplementation) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRevision

`func (o *PatentPatentImplementation) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PatentPatentImplementation) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PatentPatentImplementation) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *PatentPatentImplementation) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSeq

`func (o *PatentPatentImplementation) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *PatentPatentImplementation) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *PatentPatentImplementation) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *PatentPatentImplementation) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetVisibility

`func (o *PatentPatentImplementation) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *PatentPatentImplementation) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *PatentPatentImplementation) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *PatentPatentImplementation) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


