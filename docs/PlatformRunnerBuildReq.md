# PlatformRunnerBuildReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | Arch is the target architecture for the artifact lane. | [optional] 
**Args** | Pointer to **map[string]string** | Args are --build-arg values. They are what lets several images off ONE Dockerfile mean different things — the sandbox classes are three entries differing only by STAGE. Validated at the k8s choke point, with VERSION and REVISION taking precedence: those are receipts the builder derives from the tag and the commit, and a caller that could overwrite them could make an image lie about which commit it is. | [optional] 
**Binaries** | Pointer to [**[]PlatformBinarySpec**](PlatformBinarySpec.md) | Binaries selects the ARTIFACT lane (artifact.go): build what the repo&#39;s hanzo.yml &#x60;binaries:&#x60; block declares — a Go binary, an npm tarball, a Rust binary — and publish it to hanzoai/s3 instead of pushing an image. It is the same recipe hanzoai/ci reads, sent verbatim, so &#x60;image&#x60; is meaningless here and must be absent. | [optional] 
**Branch** | Pointer to **string** | Branch is the branch to build when no SHA or Ref is given. | [optional] 
**Bucket** | Pointer to **string** | Bucket mirrors hanzo.yml&#39;s &#x60;bucket:&#x60; — where the artifact lane publishes. | [optional] 
**Context** | Pointer to **string** | Context is the build context path within the repo. | [optional] 
**DockerTarget** | Pointer to **string** | DockerTarget is the multi-stage build target to stop at. | [optional] 
**Dockerfile** | Pointer to **string** | Dockerfile is the path to build from; empty uses the zero-config frontend. | [optional] 
**Image** | Pointer to **string** | Image is the output image ref to push. Required on the image lane, and it must target a registry namespace the caller&#39;s org owns. | [optional] 
**Os** | Pointer to **string** | OS is the target operating system for the artifact lane. | [optional] 
**Platforms** | Pointer to **[]string** | Platforms are the &#x60;&lt;os&gt;/&lt;arch&gt;&#x60; pairs the image is built for. EMPTY MEANS EVERY ARCHITECTURE THE FLEET HAS — each built on a node of that architecture, joined by ONE manifest index, so a single tag serves both. Name one to build only that one. | [optional] 
**Ref** | Pointer to **string** | Ref is the git ref to build when no SHA is given. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository clone URL to build. Required on the image lane. | [optional] 
**Sha** | Pointer to **string** | SHA is the commit to pin; it wins over Ref and Branch. | [optional] 
**Tag** | Pointer to **string** | Tag is the publish path segment, so both entry points write ONE index at ONE URL. It defaults to the pinned ref, and must be named explicitly for a branch. | [optional] 
**Tags** | Pointer to **[]string** | Tags are more tags of Image&#39;s own repository, pushed onto the same manifest by the same export, so every one names the digest Image does — by construction, not by a second push agreeing with the first. Bare tags (&#x60;v1.2.3&#x60;, &#x60;latest&#x60;), at most 8, and only beside a tag Image: a digest names an image that already exists. Duplicates and Image&#39;s own tag are dropped; the answer&#39;s &#x60;tags&#x60; lists what will be written. | [optional] 

## Methods

### NewPlatformRunnerBuildReq

`func NewPlatformRunnerBuildReq() *PlatformRunnerBuildReq`

NewPlatformRunnerBuildReq instantiates a new PlatformRunnerBuildReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformRunnerBuildReqWithDefaults

`func NewPlatformRunnerBuildReqWithDefaults() *PlatformRunnerBuildReq`

NewPlatformRunnerBuildReqWithDefaults instantiates a new PlatformRunnerBuildReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *PlatformRunnerBuildReq) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *PlatformRunnerBuildReq) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *PlatformRunnerBuildReq) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *PlatformRunnerBuildReq) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetArgs

`func (o *PlatformRunnerBuildReq) GetArgs() map[string]string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *PlatformRunnerBuildReq) GetArgsOk() (*map[string]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *PlatformRunnerBuildReq) SetArgs(v map[string]string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *PlatformRunnerBuildReq) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetBinaries

`func (o *PlatformRunnerBuildReq) GetBinaries() []PlatformBinarySpec`

GetBinaries returns the Binaries field if non-nil, zero value otherwise.

### GetBinariesOk

`func (o *PlatformRunnerBuildReq) GetBinariesOk() (*[]PlatformBinarySpec, bool)`

GetBinariesOk returns a tuple with the Binaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaries

`func (o *PlatformRunnerBuildReq) SetBinaries(v []PlatformBinarySpec)`

SetBinaries sets Binaries field to given value.

### HasBinaries

`func (o *PlatformRunnerBuildReq) HasBinaries() bool`

HasBinaries returns a boolean if a field has been set.

### GetBranch

`func (o *PlatformRunnerBuildReq) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *PlatformRunnerBuildReq) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *PlatformRunnerBuildReq) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *PlatformRunnerBuildReq) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetBucket

`func (o *PlatformRunnerBuildReq) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *PlatformRunnerBuildReq) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *PlatformRunnerBuildReq) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *PlatformRunnerBuildReq) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetContext

`func (o *PlatformRunnerBuildReq) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PlatformRunnerBuildReq) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PlatformRunnerBuildReq) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *PlatformRunnerBuildReq) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetDockerTarget

`func (o *PlatformRunnerBuildReq) GetDockerTarget() string`

GetDockerTarget returns the DockerTarget field if non-nil, zero value otherwise.

### GetDockerTargetOk

`func (o *PlatformRunnerBuildReq) GetDockerTargetOk() (*string, bool)`

GetDockerTargetOk returns a tuple with the DockerTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerTarget

`func (o *PlatformRunnerBuildReq) SetDockerTarget(v string)`

SetDockerTarget sets DockerTarget field to given value.

### HasDockerTarget

`func (o *PlatformRunnerBuildReq) HasDockerTarget() bool`

HasDockerTarget returns a boolean if a field has been set.

### GetDockerfile

`func (o *PlatformRunnerBuildReq) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *PlatformRunnerBuildReq) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *PlatformRunnerBuildReq) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *PlatformRunnerBuildReq) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### GetImage

`func (o *PlatformRunnerBuildReq) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PlatformRunnerBuildReq) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PlatformRunnerBuildReq) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PlatformRunnerBuildReq) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetOs

`func (o *PlatformRunnerBuildReq) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *PlatformRunnerBuildReq) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *PlatformRunnerBuildReq) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *PlatformRunnerBuildReq) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPlatforms

`func (o *PlatformRunnerBuildReq) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *PlatformRunnerBuildReq) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *PlatformRunnerBuildReq) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *PlatformRunnerBuildReq) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetRef

`func (o *PlatformRunnerBuildReq) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *PlatformRunnerBuildReq) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *PlatformRunnerBuildReq) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *PlatformRunnerBuildReq) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRepo

`func (o *PlatformRunnerBuildReq) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *PlatformRunnerBuildReq) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *PlatformRunnerBuildReq) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *PlatformRunnerBuildReq) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSha

`func (o *PlatformRunnerBuildReq) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *PlatformRunnerBuildReq) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *PlatformRunnerBuildReq) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *PlatformRunnerBuildReq) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetTag

`func (o *PlatformRunnerBuildReq) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PlatformRunnerBuildReq) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PlatformRunnerBuildReq) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PlatformRunnerBuildReq) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTags

`func (o *PlatformRunnerBuildReq) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PlatformRunnerBuildReq) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PlatformRunnerBuildReq) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PlatformRunnerBuildReq) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


