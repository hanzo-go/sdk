# RunnerBuildResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildJobId** | Pointer to **string** | BuildJobID is the queued build&#39;s id, and what its progress is read by. | [optional] 
**Image** | Pointer to **string** | Image is the ref the image lane will push. | [optional] 
**Index** | Pointer to **string** | Index is the binaries.json URL the artifact lane will publish. | [optional] 
**RunnerPool** | Pointer to **string** | RunnerPool is the runner class the build was placed on. | [optional] 
**Status** | Pointer to **string** | Status is &#x60;queued&#x60; — the build was accepted and has not finished. | [optional] 
**Target** | Pointer to **string** | Target is the multi-stage build target, echoed back. | [optional] 

## Methods

### NewRunnerBuildResp

`func NewRunnerBuildResp() *RunnerBuildResp`

NewRunnerBuildResp instantiates a new RunnerBuildResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerBuildRespWithDefaults

`func NewRunnerBuildRespWithDefaults() *RunnerBuildResp`

NewRunnerBuildRespWithDefaults instantiates a new RunnerBuildResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildJobId

`func (o *RunnerBuildResp) GetBuildJobId() string`

GetBuildJobId returns the BuildJobId field if non-nil, zero value otherwise.

### GetBuildJobIdOk

`func (o *RunnerBuildResp) GetBuildJobIdOk() (*string, bool)`

GetBuildJobIdOk returns a tuple with the BuildJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildJobId

`func (o *RunnerBuildResp) SetBuildJobId(v string)`

SetBuildJobId sets BuildJobId field to given value.

### HasBuildJobId

`func (o *RunnerBuildResp) HasBuildJobId() bool`

HasBuildJobId returns a boolean if a field has been set.

### GetImage

`func (o *RunnerBuildResp) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *RunnerBuildResp) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *RunnerBuildResp) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *RunnerBuildResp) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetIndex

`func (o *RunnerBuildResp) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *RunnerBuildResp) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *RunnerBuildResp) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *RunnerBuildResp) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetRunnerPool

`func (o *RunnerBuildResp) GetRunnerPool() string`

GetRunnerPool returns the RunnerPool field if non-nil, zero value otherwise.

### GetRunnerPoolOk

`func (o *RunnerBuildResp) GetRunnerPoolOk() (*string, bool)`

GetRunnerPoolOk returns a tuple with the RunnerPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerPool

`func (o *RunnerBuildResp) SetRunnerPool(v string)`

SetRunnerPool sets RunnerPool field to given value.

### HasRunnerPool

`func (o *RunnerBuildResp) HasRunnerPool() bool`

HasRunnerPool returns a boolean if a field has been set.

### GetStatus

`func (o *RunnerBuildResp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RunnerBuildResp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RunnerBuildResp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RunnerBuildResp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTarget

`func (o *RunnerBuildResp) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RunnerBuildResp) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RunnerBuildResp) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *RunnerBuildResp) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


