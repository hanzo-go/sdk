# Pipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behind** | Pointer to **int64** | Behind counts commits after the one that last produced an image whose own build has FINISHED without producing one, and Since is when the oldest of them landed. A commit still building is not counted, so a push in flight is not drift and a service appears here only once something has actually stopped without shipping. How long that has stood is the number worth acting on; that it is true says nothing about whether anyone should move. | [optional] 
**Built** | Pointer to [**Artifact**](Artifact.md) |  | [optional] 
**Declared** | Pointer to [**Artifact**](Artifact.md) |  | [optional] 
**Drift** | Pointer to **[]string** |  | [optional] 
**Head** | Pointer to [**Tip**](Tip.md) |  | [optional] 
**Image** | Pointer to **string** | ghcr.io/hanzoai/cloud | [optional] 
**Name** | Pointer to **string** | cloud | [optional] 
**Namespace** | Pointer to **string** | hanzo | [optional] 
**Org** | Pointer to **string** | Hanzo Git owner; empty when the repo is unresolved | [optional] 
**PinnedAt** | Pointer to **time.Time** |  | [optional] 
**Ready** | Pointer to **int64** |  | [optional] 
**Repo** | Pointer to **string** | hanzo-inc/cloud | [optional] 
**Running** | Pointer to [**Artifact**](Artifact.md) |  | [optional] 
**Since** | Pointer to **time.Time** |  | [optional] 
**Want** | Pointer to **int64** |  | [optional] 

## Methods

### NewPipeline

`func NewPipeline() *Pipeline`

NewPipeline instantiates a new Pipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineWithDefaults

`func NewPipelineWithDefaults() *Pipeline`

NewPipelineWithDefaults instantiates a new Pipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehind

`func (o *Pipeline) GetBehind() int64`

GetBehind returns the Behind field if non-nil, zero value otherwise.

### GetBehindOk

`func (o *Pipeline) GetBehindOk() (*int64, bool)`

GetBehindOk returns a tuple with the Behind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehind

`func (o *Pipeline) SetBehind(v int64)`

SetBehind sets Behind field to given value.

### HasBehind

`func (o *Pipeline) HasBehind() bool`

HasBehind returns a boolean if a field has been set.

### GetBuilt

`func (o *Pipeline) GetBuilt() Artifact`

GetBuilt returns the Built field if non-nil, zero value otherwise.

### GetBuiltOk

`func (o *Pipeline) GetBuiltOk() (*Artifact, bool)`

GetBuiltOk returns a tuple with the Built field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilt

`func (o *Pipeline) SetBuilt(v Artifact)`

SetBuilt sets Built field to given value.

### HasBuilt

`func (o *Pipeline) HasBuilt() bool`

HasBuilt returns a boolean if a field has been set.

### GetDeclared

`func (o *Pipeline) GetDeclared() Artifact`

GetDeclared returns the Declared field if non-nil, zero value otherwise.

### GetDeclaredOk

`func (o *Pipeline) GetDeclaredOk() (*Artifact, bool)`

GetDeclaredOk returns a tuple with the Declared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclared

`func (o *Pipeline) SetDeclared(v Artifact)`

SetDeclared sets Declared field to given value.

### HasDeclared

`func (o *Pipeline) HasDeclared() bool`

HasDeclared returns a boolean if a field has been set.

### GetDrift

`func (o *Pipeline) GetDrift() []string`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *Pipeline) GetDriftOk() (*[]string, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *Pipeline) SetDrift(v []string)`

SetDrift sets Drift field to given value.

### HasDrift

`func (o *Pipeline) HasDrift() bool`

HasDrift returns a boolean if a field has been set.

### GetHead

`func (o *Pipeline) GetHead() Tip`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *Pipeline) GetHeadOk() (*Tip, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *Pipeline) SetHead(v Tip)`

SetHead sets Head field to given value.

### HasHead

`func (o *Pipeline) HasHead() bool`

HasHead returns a boolean if a field has been set.

### GetImage

`func (o *Pipeline) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Pipeline) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Pipeline) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Pipeline) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *Pipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pipeline) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Pipeline) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *Pipeline) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Pipeline) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Pipeline) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Pipeline) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOrg

`func (o *Pipeline) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Pipeline) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Pipeline) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Pipeline) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPinnedAt

`func (o *Pipeline) GetPinnedAt() time.Time`

GetPinnedAt returns the PinnedAt field if non-nil, zero value otherwise.

### GetPinnedAtOk

`func (o *Pipeline) GetPinnedAtOk() (*time.Time, bool)`

GetPinnedAtOk returns a tuple with the PinnedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedAt

`func (o *Pipeline) SetPinnedAt(v time.Time)`

SetPinnedAt sets PinnedAt field to given value.

### HasPinnedAt

`func (o *Pipeline) HasPinnedAt() bool`

HasPinnedAt returns a boolean if a field has been set.

### GetReady

`func (o *Pipeline) GetReady() int64`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *Pipeline) GetReadyOk() (*int64, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *Pipeline) SetReady(v int64)`

SetReady sets Ready field to given value.

### HasReady

`func (o *Pipeline) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetRepo

`func (o *Pipeline) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *Pipeline) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *Pipeline) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *Pipeline) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRunning

`func (o *Pipeline) GetRunning() Artifact`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *Pipeline) GetRunningOk() (*Artifact, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *Pipeline) SetRunning(v Artifact)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *Pipeline) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetSince

`func (o *Pipeline) GetSince() time.Time`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *Pipeline) GetSinceOk() (*time.Time, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *Pipeline) SetSince(v time.Time)`

SetSince sets Since field to given value.

### HasSince

`func (o *Pipeline) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetWant

`func (o *Pipeline) GetWant() int64`

GetWant returns the Want field if non-nil, zero value otherwise.

### GetWantOk

`func (o *Pipeline) GetWantOk() (*int64, bool)`

GetWantOk returns a tuple with the Want field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWant

`func (o *Pipeline) SetWant(v int64)`

SetWant sets Want field to given value.

### HasWant

`func (o *Pipeline) HasWant() bool`

HasWant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


