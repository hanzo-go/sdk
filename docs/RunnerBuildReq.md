# RunnerBuildReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | Arch is the target architecture for the artifact lane. | [optional] 
**Args** | Pointer to **map[string]string** | Args are --build-arg values. They are what lets several images off ONE Dockerfile mean different things — the sandbox classes are three entries differing only by STAGE. Validated at the k8s choke point, with VERSION and REVISION taking precedence: those are receipts the builder derives from the tag and the commit, and a caller that could overwrite them could make an image lie about which commit it is. | [optional] 
**Binaries** | Pointer to [**[]BinarySpec**](BinarySpec.md) | Binaries selects the ARTIFACT lane (artifact.go): build what the repo&#39;s hanzo.yml &#x60;binaries:&#x60; block declares — a Go binary, an npm tarball, a Rust binary — and publish it to hanzoai/s3 instead of pushing an image. It is the same recipe hanzoai/ci reads, sent verbatim, so &#x60;image&#x60; is meaningless here and must be absent. | [optional] 
**Branch** | Pointer to **string** | Branch is the branch to build when no SHA or Ref is given. | [optional] 
**Bucket** | Pointer to **string** | Bucket mirrors hanzo.yml&#39;s &#x60;bucket:&#x60; — where the artifact lane publishes. | [optional] 
**Context** | Pointer to **string** | Context is the build context path within the repo. | [optional] 
**DockerTarget** | Pointer to **string** | DockerTarget is the multi-stage build target to stop at. | [optional] 
**Dockerfile** | Pointer to **string** | Dockerfile is the path to build from; empty uses the zero-config frontend. | [optional] 
**Image** | Pointer to **string** | Image is the output image ref to push. Required on the image lane, and it must target a registry namespace the caller&#39;s org owns. | [optional] 
**Os** | Pointer to **string** | OS is the target operating system for the artifact lane. | [optional] 
**Ref** | Pointer to **string** | Ref is the git ref to build when no SHA is given. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository clone URL to build. Required on the image lane. | [optional] 
**Sha** | Pointer to **string** | SHA is the commit to pin; it wins over Ref and Branch. | [optional] 
**Tag** | Pointer to **string** | Tag is the publish path segment, so both entry points write ONE index at ONE URL. It defaults to the pinned ref, and must be named explicitly for a branch. | [optional] 

## Methods

### NewRunnerBuildReq

`func NewRunnerBuildReq() *RunnerBuildReq`

NewRunnerBuildReq instantiates a new RunnerBuildReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerBuildReqWithDefaults

`func NewRunnerBuildReqWithDefaults() *RunnerBuildReq`

NewRunnerBuildReqWithDefaults instantiates a new RunnerBuildReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *RunnerBuildReq) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *RunnerBuildReq) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *RunnerBuildReq) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *RunnerBuildReq) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetArgs

`func (o *RunnerBuildReq) GetArgs() map[string]string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *RunnerBuildReq) GetArgsOk() (*map[string]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *RunnerBuildReq) SetArgs(v map[string]string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *RunnerBuildReq) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetBinaries

`func (o *RunnerBuildReq) GetBinaries() []BinarySpec`

GetBinaries returns the Binaries field if non-nil, zero value otherwise.

### GetBinariesOk

`func (o *RunnerBuildReq) GetBinariesOk() (*[]BinarySpec, bool)`

GetBinariesOk returns a tuple with the Binaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaries

`func (o *RunnerBuildReq) SetBinaries(v []BinarySpec)`

SetBinaries sets Binaries field to given value.

### HasBinaries

`func (o *RunnerBuildReq) HasBinaries() bool`

HasBinaries returns a boolean if a field has been set.

### GetBranch

`func (o *RunnerBuildReq) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *RunnerBuildReq) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *RunnerBuildReq) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *RunnerBuildReq) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetBucket

`func (o *RunnerBuildReq) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *RunnerBuildReq) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *RunnerBuildReq) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *RunnerBuildReq) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetContext

`func (o *RunnerBuildReq) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RunnerBuildReq) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RunnerBuildReq) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *RunnerBuildReq) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetDockerTarget

`func (o *RunnerBuildReq) GetDockerTarget() string`

GetDockerTarget returns the DockerTarget field if non-nil, zero value otherwise.

### GetDockerTargetOk

`func (o *RunnerBuildReq) GetDockerTargetOk() (*string, bool)`

GetDockerTargetOk returns a tuple with the DockerTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerTarget

`func (o *RunnerBuildReq) SetDockerTarget(v string)`

SetDockerTarget sets DockerTarget field to given value.

### HasDockerTarget

`func (o *RunnerBuildReq) HasDockerTarget() bool`

HasDockerTarget returns a boolean if a field has been set.

### GetDockerfile

`func (o *RunnerBuildReq) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *RunnerBuildReq) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *RunnerBuildReq) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *RunnerBuildReq) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### GetImage

`func (o *RunnerBuildReq) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *RunnerBuildReq) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *RunnerBuildReq) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *RunnerBuildReq) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetOs

`func (o *RunnerBuildReq) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *RunnerBuildReq) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *RunnerBuildReq) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *RunnerBuildReq) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetRef

`func (o *RunnerBuildReq) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RunnerBuildReq) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RunnerBuildReq) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RunnerBuildReq) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRepo

`func (o *RunnerBuildReq) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *RunnerBuildReq) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *RunnerBuildReq) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *RunnerBuildReq) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSha

`func (o *RunnerBuildReq) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *RunnerBuildReq) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *RunnerBuildReq) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *RunnerBuildReq) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetTag

`func (o *RunnerBuildReq) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *RunnerBuildReq) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *RunnerBuildReq) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *RunnerBuildReq) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


