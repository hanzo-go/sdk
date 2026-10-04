# PlatformRunnerBuildResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildJobId** | Pointer to **string** | BuildJobID is the queued build&#39;s id, and what its progress is read by. | [optional] 
**Image** | Pointer to **string** | Image is the ref the image lane will push. | [optional] 
**Index** | Pointer to **string** | Index is the binaries.json URL the artifact lane will publish. | [optional] 
**Platforms** | Pointer to **[]string** | Platforms are the architectures the image lane will publish, echoed back. | [optional] 
**Reason** | Pointer to **string** | Reason is why a failed build failed, as the cluster said it: the solve&#39;s error line (&#x60;error: failed to solve: ...&#x60;), a container waiting on what it cannot have, or the deadline. | [optional] 
**RunnerPool** | Pointer to **string** | RunnerPool is the runner class the build was placed on. | [optional] 
**Status** | Pointer to **string** | Status is &#x60;queued&#x60; until the build finishes, then &#x60;succeeded&#x60; or &#x60;failed&#x60;. The POST answers &#x60;queued&#x60;: a build that is accepted has not finished. | [optional] 
**Tags** | Pointer to **[]string** | Tags are the extra tags the image lane will write beside Image, onto the same manifest: the request&#39;s, deduplicated. This is the promise — a tag absent here will not be written. | [optional] 
**Target** | Pointer to **string** | Target is the multi-stage build target, echoed back. | [optional] 

## Methods

### NewPlatformRunnerBuildResp

`func NewPlatformRunnerBuildResp() *PlatformRunnerBuildResp`

NewPlatformRunnerBuildResp instantiates a new PlatformRunnerBuildResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformRunnerBuildRespWithDefaults

`func NewPlatformRunnerBuildRespWithDefaults() *PlatformRunnerBuildResp`

NewPlatformRunnerBuildRespWithDefaults instantiates a new PlatformRunnerBuildResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildJobId

`func (o *PlatformRunnerBuildResp) GetBuildJobId() string`

GetBuildJobId returns the BuildJobId field if non-nil, zero value otherwise.

### GetBuildJobIdOk

`func (o *PlatformRunnerBuildResp) GetBuildJobIdOk() (*string, bool)`

GetBuildJobIdOk returns a tuple with the BuildJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildJobId

`func (o *PlatformRunnerBuildResp) SetBuildJobId(v string)`

SetBuildJobId sets BuildJobId field to given value.

### HasBuildJobId

`func (o *PlatformRunnerBuildResp) HasBuildJobId() bool`

HasBuildJobId returns a boolean if a field has been set.

### GetImage

`func (o *PlatformRunnerBuildResp) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PlatformRunnerBuildResp) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PlatformRunnerBuildResp) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PlatformRunnerBuildResp) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetIndex

`func (o *PlatformRunnerBuildResp) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *PlatformRunnerBuildResp) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *PlatformRunnerBuildResp) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *PlatformRunnerBuildResp) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetPlatforms

`func (o *PlatformRunnerBuildResp) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *PlatformRunnerBuildResp) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *PlatformRunnerBuildResp) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *PlatformRunnerBuildResp) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetReason

`func (o *PlatformRunnerBuildResp) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PlatformRunnerBuildResp) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PlatformRunnerBuildResp) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PlatformRunnerBuildResp) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRunnerPool

`func (o *PlatformRunnerBuildResp) GetRunnerPool() string`

GetRunnerPool returns the RunnerPool field if non-nil, zero value otherwise.

### GetRunnerPoolOk

`func (o *PlatformRunnerBuildResp) GetRunnerPoolOk() (*string, bool)`

GetRunnerPoolOk returns a tuple with the RunnerPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerPool

`func (o *PlatformRunnerBuildResp) SetRunnerPool(v string)`

SetRunnerPool sets RunnerPool field to given value.

### HasRunnerPool

`func (o *PlatformRunnerBuildResp) HasRunnerPool() bool`

HasRunnerPool returns a boolean if a field has been set.

### GetStatus

`func (o *PlatformRunnerBuildResp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlatformRunnerBuildResp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlatformRunnerBuildResp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlatformRunnerBuildResp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *PlatformRunnerBuildResp) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PlatformRunnerBuildResp) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PlatformRunnerBuildResp) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PlatformRunnerBuildResp) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTarget

`func (o *PlatformRunnerBuildResp) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PlatformRunnerBuildResp) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PlatformRunnerBuildResp) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *PlatformRunnerBuildResp) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


