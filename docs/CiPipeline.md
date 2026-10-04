# CiPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behind** | Pointer to **int64** | Behind counts commits after the one that last produced an image whose own build has FINISHED without producing one, and Since is when the oldest of them landed. A commit still building is not counted, so a push in flight is not drift and a service appears here only once something has actually stopped without shipping. How long that has stood is the number worth acting on; that it is true says nothing about whether anyone should move. | [optional] 
**Built** | Pointer to [**CiArtifact**](CiArtifact.md) |  | [optional] 
**Declared** | Pointer to [**CiArtifact**](CiArtifact.md) |  | [optional] 
**Drift** | Pointer to **[]string** |  | [optional] 
**Head** | Pointer to [**CiTip**](CiTip.md) |  | [optional] 
**Image** | Pointer to **string** | ghcr.io/hanzoai/cloud | [optional] 
**Name** | Pointer to **string** | cloud | [optional] 
**Namespace** | Pointer to **string** | hanzo | [optional] 
**Org** | Pointer to **string** | Hanzo Git owner; empty when the repo is unresolved | [optional] 
**PinnedAt** | Pointer to **time.Time** |  | [optional] 
**Ready** | Pointer to **int64** |  | [optional] 
**Repo** | Pointer to **string** | hanzo-inc/cloud | [optional] 
**Running** | Pointer to [**CiArtifact**](CiArtifact.md) |  | [optional] 
**Since** | Pointer to **time.Time** |  | [optional] 
**Want** | Pointer to **int64** |  | [optional] 

## Methods

### NewCiPipeline

`func NewCiPipeline() *CiPipeline`

NewCiPipeline instantiates a new CiPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCiPipelineWithDefaults

`func NewCiPipelineWithDefaults() *CiPipeline`

NewCiPipelineWithDefaults instantiates a new CiPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehind

`func (o *CiPipeline) GetBehind() int64`

GetBehind returns the Behind field if non-nil, zero value otherwise.

### GetBehindOk

`func (o *CiPipeline) GetBehindOk() (*int64, bool)`

GetBehindOk returns a tuple with the Behind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehind

`func (o *CiPipeline) SetBehind(v int64)`

SetBehind sets Behind field to given value.

### HasBehind

`func (o *CiPipeline) HasBehind() bool`

HasBehind returns a boolean if a field has been set.

### GetBuilt

`func (o *CiPipeline) GetBuilt() CiArtifact`

GetBuilt returns the Built field if non-nil, zero value otherwise.

### GetBuiltOk

`func (o *CiPipeline) GetBuiltOk() (*CiArtifact, bool)`

GetBuiltOk returns a tuple with the Built field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilt

`func (o *CiPipeline) SetBuilt(v CiArtifact)`

SetBuilt sets Built field to given value.

### HasBuilt

`func (o *CiPipeline) HasBuilt() bool`

HasBuilt returns a boolean if a field has been set.

### GetDeclared

`func (o *CiPipeline) GetDeclared() CiArtifact`

GetDeclared returns the Declared field if non-nil, zero value otherwise.

### GetDeclaredOk

`func (o *CiPipeline) GetDeclaredOk() (*CiArtifact, bool)`

GetDeclaredOk returns a tuple with the Declared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclared

`func (o *CiPipeline) SetDeclared(v CiArtifact)`

SetDeclared sets Declared field to given value.

### HasDeclared

`func (o *CiPipeline) HasDeclared() bool`

HasDeclared returns a boolean if a field has been set.

### GetDrift

`func (o *CiPipeline) GetDrift() []string`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *CiPipeline) GetDriftOk() (*[]string, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *CiPipeline) SetDrift(v []string)`

SetDrift sets Drift field to given value.

### HasDrift

`func (o *CiPipeline) HasDrift() bool`

HasDrift returns a boolean if a field has been set.

### GetHead

`func (o *CiPipeline) GetHead() CiTip`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *CiPipeline) GetHeadOk() (*CiTip, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *CiPipeline) SetHead(v CiTip)`

SetHead sets Head field to given value.

### HasHead

`func (o *CiPipeline) HasHead() bool`

HasHead returns a boolean if a field has been set.

### GetImage

`func (o *CiPipeline) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CiPipeline) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CiPipeline) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CiPipeline) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *CiPipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CiPipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CiPipeline) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CiPipeline) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *CiPipeline) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CiPipeline) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CiPipeline) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CiPipeline) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOrg

`func (o *CiPipeline) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CiPipeline) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CiPipeline) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CiPipeline) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPinnedAt

`func (o *CiPipeline) GetPinnedAt() time.Time`

GetPinnedAt returns the PinnedAt field if non-nil, zero value otherwise.

### GetPinnedAtOk

`func (o *CiPipeline) GetPinnedAtOk() (*time.Time, bool)`

GetPinnedAtOk returns a tuple with the PinnedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedAt

`func (o *CiPipeline) SetPinnedAt(v time.Time)`

SetPinnedAt sets PinnedAt field to given value.

### HasPinnedAt

`func (o *CiPipeline) HasPinnedAt() bool`

HasPinnedAt returns a boolean if a field has been set.

### GetReady

`func (o *CiPipeline) GetReady() int64`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *CiPipeline) GetReadyOk() (*int64, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *CiPipeline) SetReady(v int64)`

SetReady sets Ready field to given value.

### HasReady

`func (o *CiPipeline) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetRepo

`func (o *CiPipeline) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CiPipeline) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CiPipeline) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CiPipeline) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRunning

`func (o *CiPipeline) GetRunning() CiArtifact`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *CiPipeline) GetRunningOk() (*CiArtifact, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *CiPipeline) SetRunning(v CiArtifact)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *CiPipeline) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetSince

`func (o *CiPipeline) GetSince() time.Time`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *CiPipeline) GetSinceOk() (*time.Time, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *CiPipeline) SetSince(v time.Time)`

SetSince sets Since field to given value.

### HasSince

`func (o *CiPipeline) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetWant

`func (o *CiPipeline) GetWant() int64`

GetWant returns the Want field if non-nil, zero value otherwise.

### GetWantOk

`func (o *CiPipeline) GetWantOk() (*int64, bool)`

GetWantOk returns a tuple with the Want field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWant

`func (o *CiPipeline) SetWant(v int64)`

SetWant sets Want field to given value.

### HasWant

`func (o *CiPipeline) HasWant() bool`

HasWant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


